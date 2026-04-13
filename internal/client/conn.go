package futuapi

import (
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	HeaderLen     = 44
	MagicBytes    = "FT"
	ProtoVersion  = 0
	MaxPacketSize = 10 * 1024 * 1024
)

var (
	ErrInvalidHeader  = errors.New("invalid packet header")
	ErrInvalidMagic   = errors.New("invalid magic bytes")
	ErrPacketTooBig   = errors.New("packet too large")
	ErrInvalidBodyLen = errors.New("invalid body length")
)

type Header struct {
	Magic    [2]byte
	ProtoID  uint32
	ProtoFmt byte
	ProtoVer byte
	SerialNo uint32
	BodyLen  uint32
	BodySHA1 [20]byte
	Reserved [8]byte
}

type Packet struct {
	Header Header
	Body   []byte
}

type PacketHandler func(pkt *Packet)

type Conn struct {
	conn net.Conn
	mu   sync.Mutex

	dispMu   sync.Mutex
	disp     map[uint32]chan *Packet
	dispSize int

	pushHandler PacketHandler
	apiTimeout  time.Duration
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
		disp: make(map[uint32]chan *Packet),
	}
}

func (c *Conn) SetPushHandler(handler PacketHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pushHandler = handler
}

func (c *Conn) APITimeout() time.Duration {
	return c.apiTimeout
}

func (c *Conn) Dial(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, 30*time.Second)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Conn) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	if c.conn == nil {
		return fmt.Errorf("set read deadline: %w", ErrNotConnected)
	}
	return c.conn.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	if c.conn == nil {
		return fmt.Errorf("set write deadline: %w", ErrNotConnected)
	}
	return c.conn.SetWriteDeadline(t)
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) Dispatch(pkt *Packet) {
	c.dispMu.Lock()
	ch, ok := c.disp[pkt.Header.SerialNo]
	delete(c.disp, pkt.Header.SerialNo)
	c.dispSize--
	c.dispMu.Unlock()

	if ok {
		select {
		case ch <- pkt:
		default:
		}
		return
	}

	c.mu.Lock()
	h := c.pushHandler
	c.mu.Unlock()
	if h != nil {
		h(pkt)
	}
}

func (c *Conn) readOne() (*Packet, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("read packet: %w", ErrNotConnected)
	}

	header := make([]byte, HeaderLen)
	n, err := io.ReadFull(c.conn, header)
	if err != nil {
		return nil, fmt.Errorf("read header (%d/%d bytes): %w", n, HeaderLen, err)
	}

	var h Header
	copy(h.Magic[:], header[0:2])
	h.ProtoID = binary.LittleEndian.Uint32(header[2:6])
	h.ProtoFmt = byte(header[6])
	h.ProtoVer = header[7]
	h.SerialNo = binary.LittleEndian.Uint32(header[8:12])
	h.BodyLen = binary.LittleEndian.Uint32(header[12:16])
	copy(h.BodySHA1[:], header[16:36])
	copy(h.Reserved[:], header[36:44])

	if string(h.Magic[:]) != "FT" {
		return nil, fmt.Errorf("invalid magic: % x (expected 'FT')", h.Magic)
	}

	if h.BodyLen > MaxPacketSize {
		return nil, fmt.Errorf("body too large: %d bytes", h.BodyLen)
	}

	body := make([]byte, h.BodyLen)
	if h.BodyLen > 0 {
		n, err := io.ReadFull(c.conn, body)
		if err != nil {
			return nil, fmt.Errorf("read body (%d/%d bytes): %w", n, h.BodyLen, err)
		}
	}

	// Verify SHA1 checksum
	actualSHA1 := sha1.Sum(body)
	if actualSHA1 != h.BodySHA1 {
		return nil, fmt.Errorf("body integrity check: %w", ErrChecksumMismatch)
	}

	return &Packet{Header: h, Body: body}, nil
}

func (c *Conn) ReadResponse(serial uint32, timeout time.Duration) (*Packet, error) {
	ch := make(chan *Packet, 1)

	c.dispMu.Lock()
	c.disp[serial] = ch
	c.dispSize++
	c.dispMu.Unlock()

	defer func() {
		c.dispMu.Lock()
		delete(c.disp, serial)
		c.dispMu.Unlock()
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case pkt := <-ch:
		return pkt, nil
	case <-timer.C:
		return nil, fmt.Errorf("read response: i/o timeout")
	}
}

func (c *Conn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	if c.conn == nil {
		return fmt.Errorf("write packet: %w", ErrNotConnected)
	}

	header := make([]byte, HeaderLen)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:], protoID)
	header[6] = 0
	header[7] = ProtoVersion
	binary.LittleEndian.PutUint32(header[8:], serialNo)
	binary.LittleEndian.PutUint32(header[12:], uint32(len(body)))
	sha1Hash := sha1.Sum(body)
	copy(header[16:36], sha1Hash[:])

	if _, err := c.conn.Write(header); err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	if len(body) > 0 {
		if _, err := c.conn.Write(body); err != nil {
			return fmt.Errorf("write body: %w", err)
		}
	}

	return nil
}
