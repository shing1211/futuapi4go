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
	HeaderLen     = 44  // Fixed: Official Futu protocol spec
	MagicBytes    = "FT"
	ProtoVersion  = 0   // Fixed: Protocol version is 0, not 1
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
	ProtoFmt byte              // 1 byte on wire
	ProtoVer byte              // 1 byte on wire  
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
	conn           net.Conn
	mu             sync.Mutex
	sem            chan struct{}
	expectedSerial uint32        // Serial number we're expecting
	pushHandler    PacketHandler // Handler for push notifications (packets with unexpected serial)
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
		sem:  make(chan struct{}, 1),
	}
}

// SetPushHandler sets a callback for handling push notifications.
// Packets with serial numbers that don't match expected requests will be dispatched here.
func (c *Conn) SetPushHandler(handler PacketHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pushHandler = handler
}

// SetExpectedSerial sets the serial number we're expecting for the next response.
func (c *Conn) SetExpectedSerial(serial uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.expectedSerial = serial
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

func (c *Conn) ReadPacket() (*Packet, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("read packet: %w", ErrNotConnected)
	}

	c.mu.Lock()
	expectedSerial := c.expectedSerial
	pushHandler := c.pushHandler
	c.mu.Unlock()

	// Read packets until we get one with the expected serial number.
	// Packets with other serials are push notifications and dispatched to the handler.
	for {
		// Read 44-byte header
		header := make([]byte, HeaderLen)
		n, err := io.ReadFull(c.conn, header)
		if err != nil {
			return nil, fmt.Errorf("read header (%d/%d bytes): %w", n, HeaderLen, err)
		}

		// Parse header fields manually (no struct padding)
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

		pkt := &Packet{Header: h, Body: body}

		// Check if this packet matches the expected serial
		if expectedSerial > 0 && h.SerialNo == expectedSerial {
			return pkt, nil
		}

		// No pending request (expectedSerial == 0) or serial mismatch
		// This is a push notification - dispatch to handler
		if pushHandler != nil {
			pushHandler(pkt)
		}
		// If expectedSerial == 0, we keep reading for push notifications forever
		// If expectedSerial > 0, we continue until we find the matching serial
	}
}

func (c *Conn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	if c.conn == nil {
		return fmt.Errorf("write packet: %w", ErrNotConnected)
	}

	c.mu.Lock()
	// Set expected serial for the next ReadPacket call
	// This enables serial matching to prevent push notifications from being consumed as responses
	c.expectedSerial = serialNo
	defer c.mu.Unlock()

	// Manually encode header per official Futu protocol spec (44 bytes)
	// Reference: https://openapi.futunn.com/futu-api-doc/en/ftapi/protocol.html
	header := make([]byte, HeaderLen)
	
	// Byte 0-1: Magic "FT" (2 bytes)
	header[0] = 'F'
	header[1] = 'T'
	
	// Byte 2-5: ProtoID (4 bytes, little-endian)
	binary.LittleEndian.PutUint32(header[2:], protoID)
	
	// Byte 6: ProtoFmt (1 byte) - 0=Protobuf, 1=JSON
	header[6] = 0  // Protobuf format (byte value)
	
	// Byte 7: ProtoVer (1 byte) - currently 0
	header[7] = ProtoVersion
	
	// Byte 8-11: SerialNo (4 bytes, little-endian)
	binary.LittleEndian.PutUint32(header[8:], serialNo)
	
	// Byte 12-15: BodyLen (4 bytes, little-endian)
	binary.LittleEndian.PutUint32(header[12:], uint32(len(body)))
	
	// Byte 16-35: BodySHA1 (20 bytes)
	sha1Hash := sha1.Sum(body)
	copy(header[16:36], sha1Hash[:])
	
	// Byte 36-43: Reserved (8 bytes) - zeros

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
