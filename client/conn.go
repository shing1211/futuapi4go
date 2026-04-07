package futuapi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"gitee.com/shing1211/futuapi4go/pb/common"
)

const (
	HeaderLen     = 48
	MagicBytes    = "FT"
	ProtoVersion  = 1
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
	ProtoFmt common.ProtoFmt
	ProtoVer uint16
	SerialNo uint32
	BodyLen  uint32
	BodySHA1 [20]byte
	Reserved [8]byte
}

type Packet struct {
	Header Header
	Body   []byte
}

type Conn struct {
	conn net.Conn
	mu   sync.Mutex
	sem  chan struct{}
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
		sem:  make(chan struct{}, 1),
	}
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
	return c.conn.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *Conn) ReadPacket() (*Packet, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	header := make([]byte, HeaderLen)
	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	var h Header
	if err := binary.Read(bytes.NewReader(header), binary.LittleEndian, &h); err != nil {
		return nil, fmt.Errorf("decode header: %w", err)
	}

	if string(h.Magic[:]) != MagicBytes {
		return nil, ErrInvalidMagic
	}

	if h.BodyLen > MaxPacketSize {
		return nil, ErrPacketTooBig
	}

	body := make([]byte, h.BodyLen)
	if h.BodyLen > 0 {
		if _, err := io.ReadFull(c.conn, body); err != nil {
			return nil, fmt.Errorf("read body: %w", err)
		}
	}

	return &Packet{Header: h, Body: body}, nil
}

func (c *Conn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	h := Header{
		ProtoFmt: common.ProtoFmt_ProtoFmt_Protobuf,
		ProtoVer: ProtoVersion,
		SerialNo: serialNo,
		BodyLen:  uint32(len(body)),
	}
	copy(h.Magic[:], MagicBytes)
	h.ProtoID = protoID

	header := make([]byte, HeaderLen)
	if err := binary.Write(bytes.NewBuffer(header[:0]), binary.LittleEndian, &h); err != nil {
		return fmt.Errorf("encode header: %w", err)
	}

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

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
