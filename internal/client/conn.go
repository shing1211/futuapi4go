// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package futuapi

import (
	"bufio"
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/binary"
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

type ConnInterface interface {
	io.Closer
	WritePacket(protoID uint32, serialNo uint32, body []byte) error
	WritePacketWithSHA1(protoID uint32, serialNo uint32, body []byte, bodySHA1 [20]byte) error
	WritePacketEncrypted(protoID uint32, serialNo uint32, encryptedBody []byte, bodySHA1 [20]byte) error
	ReadResponse(serialNo uint32, timeout time.Duration) (*Packet, error)
	ReadResponseContext(ctx context.Context, serialNo uint32, timeout time.Duration) (*Packet, error)
	SetPushHandler(handler PacketHandler)
	Dispatch(pkt *Packet)
	DrainDispatches()
	APITimeout() time.Duration
	SetAPITimeout(time.Duration)
	Dial(addr string) error
	SetTLSConfig(cfg *tls.Config)
	// ExpectEncryptedSHA1 marks inbound bodies as encrypted, so readOne leaves
	// the SHA1 check to the decrypting caller.
	ExpectEncryptedSHA1(bool)
	readOne() (*Packet, error)
}

type Conn struct {
	conn      net.Conn
	reader    *bufio.Reader
	mu        sync.Mutex
	tlsConfig *tls.Config

	dispMu sync.Mutex
	disp   map[uint32]chan *Packet
	early  map[uint32]*Packet

	pushHandler PacketHandler
	apiTimeout  time.Duration

	// expectEncryptedSHA1 records that inbound bodies are encrypted, so the
	// header's SHA1 is over the PLAINTEXT (the protocol says SHA1(plaintext);
	// the Python SDK verifies it after decrypting). readOne cannot check it
	// there, and verifying the ciphertext against a plaintext digest makes every
	// encrypted response look corrupt. Callers verify after decryption instead.
	expectEncryptedSHA1 bool
}

// maxEarlyResponses bounds the buffer of responses that arrive before their
// reader has registered interest (see Dispatch). It is a safety valve; in
// normal operation a reader consumes an early response immediately.
const maxEarlyResponses = 128

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		reader: bufio.NewReaderSize(conn, 64*1024),
		disp:   make(map[uint32]chan *Packet),
		early:  make(map[uint32]*Packet),
	}
}

func (c *Conn) SetPushHandler(handler PacketHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pushHandler = handler
}

func (c *Conn) APITimeout() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.apiTimeout
}

func (c *Conn) SetAPITimeout(timeout time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.apiTimeout = timeout
}

// ExpectEncryptedSHA1 marks whether inbound bodies are encrypted. See the field
// comment on expectEncryptedSHA1.
func (c *Conn) ExpectEncryptedSHA1(v bool) {
	c.mu.Lock()
	c.expectEncryptedSHA1 = v
	c.mu.Unlock()
}

// currentConn returns the underlying connection under the lock. Callers use the
// returned value so a concurrent Close cannot nil it mid-use.
func (c *Conn) currentConn() net.Conn {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn
}

func (c *Conn) Dial(addr string) error {
	if c.tlsConfig != nil {
		conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 30 * time.Second}, "tcp", addr, c.tlsConfig)
		if err != nil {
			return NewErrorWithWrap(CodeTLSHandshakeFailed, "TLS dial failed", err)
		}
		c.conn = conn
	} else {
		conn, err := net.DialTimeout("tcp", addr, 30*time.Second)
		if err != nil {
			return err
		}
		c.conn = conn
	}
	return nil
}

func (c *Conn) SetTLSConfig(cfg *tls.Config) {
	c.tlsConfig = cfg
}

func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	conn := c.currentConn()
	if conn == nil {
		return NewErrorWithWrap(CodeNotConnected, "set read deadline", ErrNotConnected)
	}
	return conn.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	conn := c.currentConn()
	if conn == nil {
		return NewErrorWithWrap(CodeNotConnected, "set write deadline", ErrNotConnected)
	}
	return conn.SetWriteDeadline(t)
}

func (c *Conn) LocalAddr() net.Addr {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return nil
	}
	return conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return nil
	}
	return conn.RemoteAddr()
}

func (c *Conn) Dispatch(pkt *Packet) {
	c.dispMu.Lock()
	ch, ok := c.disp[pkt.Header.SerialNo]
	if ok {
		delete(c.disp, pkt.Header.SerialNo)
		// A nil channel is a reservation placed by writePacketCommon: the
		// request was written but ReadResponse has not registered yet. Buffer
		// the response so it is not lost to the dispatch race.
		if ch == nil {
			if len(c.early) < maxEarlyResponses {
				c.early[pkt.Header.SerialNo] = pkt
			}
			c.dispMu.Unlock()
			return
		}
		c.dispMu.Unlock()

		select {
		case ch <- pkt:
		default:
		}
		return
	}
	c.dispMu.Unlock()

	c.mu.Lock()
	h := c.pushHandler
	c.mu.Unlock()
	if h != nil {
		h(pkt)
	}
}

func (c *Conn) DrainDispatches() {
	c.dispMu.Lock()
	defer c.dispMu.Unlock()
	for serial, ch := range c.disp {
		select {
		case pkt, ok := <-ch:
			if ok {
				c.dispatchToPushHandler(pkt)
			}
		default:
		}
		delete(c.disp, serial)
	}
}

func (c *Conn) dispatchToPushHandler(pkt *Packet) {
	c.mu.Lock()
	h := c.pushHandler
	c.mu.Unlock()
	if h != nil {
		h(pkt)
	}
}

func (c *Conn) readOne() (*Packet, error) {
	conn := c.currentConn()
	if conn == nil {
		return nil, NewErrorWithWrap(CodeNotConnected, "read packet", ErrNotConnected)
	}

	if apiTimeout := c.APITimeout(); apiTimeout > 0 {
		if err := conn.SetReadDeadline(time.Now().Add(apiTimeout)); err != nil {
			return nil, fmt.Errorf("set read deadline: %w", err)
		}
	}

	header := make([]byte, HeaderLen)
	n, err := io.ReadFull(conn, header)
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
		return nil, NewError(CodeInvalidMagic, fmt.Sprintf("invalid magic: % x (expected 'FT')", h.Magic))
	}

	if h.BodyLen > MaxPacketSize {
		return nil, NewError(CodePacketTooBig, fmt.Sprintf("body too large: %d bytes", h.BodyLen))
	}

	body := make([]byte, h.BodyLen)
	if h.BodyLen > 0 {
		n, err := io.ReadFull(conn, body)
		if err != nil {
			return nil, fmt.Errorf("read body (%d/%d bytes): %w", n, h.BodyLen, err)
		}
	}

	// Verify SHA1 checksum. Encrypted bodies carry SHA1(plaintext), so the
	// check belongs after decryption; doing it here would reject every
	// encrypted response. VerifyBodySHA1 is called by the decrypting caller.
	if !c.expectEncryptedSHA1 {
		actualSHA1 := sha1.Sum(body)
		if actualSHA1 != h.BodySHA1 {
			return nil, NewError(CodeChecksumMismatch, "body integrity check failed")
		}
	}

	return &Packet{Header: h, Body: body}, nil
}

// VerifyBodySHA1 checks a decrypted body against the header's SHA1. Callers that
// decrypt a response must use this instead of trusting readOne's check, which is
// skipped for encrypted connections.
func VerifyBodySHA1(pkt *Packet, plaintext []byte) error {
	actual := sha1.Sum(plaintext)
	if actual != pkt.Header.BodySHA1 {
		return NewError(CodeChecksumMismatch, "body integrity check failed")
	}
	return nil
}

func (c *Conn) ReadResponse(serial uint32, timeout time.Duration) (*Packet, error) {
	ch := make(chan *Packet, 1)

	c.dispMu.Lock()
	// The response may have arrived before this reader registered; consume it.
	if pkt, ok := c.early[serial]; ok {
		delete(c.early, serial)
		delete(c.disp, serial)
		c.dispMu.Unlock()
		return pkt, nil
	}
	c.disp[serial] = ch
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

// ReadResponseContext reads a response with context cancellation support.
func (c *Conn) ReadResponseContext(ctx context.Context, serial uint32, timeout time.Duration) (*Packet, error) {
	ch := make(chan *Packet, 1)

	c.dispMu.Lock()
	// The response may have arrived before this reader registered; consume it.
	if pkt, ok := c.early[serial]; ok {
		delete(c.early, serial)
		delete(c.disp, serial)
		c.dispMu.Unlock()
		return pkt, nil
	}
	c.disp[serial] = ch
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
	case <-ctx.Done():
		select {
		case pkt := <-ch:
			return pkt, nil
		default:
			return nil, fmt.Errorf("read response: %w", ctx.Err())
		}
	case <-timer.C:
		select {
		case pkt := <-ch:
			return pkt, nil
		default:
			return nil, fmt.Errorf("read response: i/o timeout")
		}
	}
}

func (c *Conn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	return c.writePacketCommon(protoID, serialNo, body, sha1.Sum(body))
}

func (c *Conn) WritePacketWithSHA1(protoID uint32, serialNo uint32, body []byte, bodySHA1 [20]byte) error {
	return c.writePacketCommon(protoID, serialNo, body, bodySHA1)
}

// WritePacketEncrypted writes a packet where the body is already AES-encrypted.
// bodySHA1 must be the SHA1 of the PLAINTEXT, not of the encrypted body.
//
// Authority: the official Futu Python SDK computes the digest before
// encryption and places it directly in the header:
//
//	futu/common/utils.py :: _joint_head()
//	    sha20 = hashlib.sha1(str_body).digest()   # plaintext
//	    if is_encrypt: str_body = encrypt(str_body)
//
// OpenD decrypts first, then verifies the header SHA1 against the decrypted
// plaintext (futu/common/utils.py :: decrypt_rsp_body). Signing the ciphertext
// makes encrypted requests fail with 「包体数据SHA1签名不正确」.
func (c *Conn) WritePacketEncrypted(protoID uint32, serialNo uint32, encryptedBody []byte, bodySHA1 [20]byte) error {
	return c.writePacketCommon(protoID, serialNo, encryptedBody, bodySHA1)
}

func (c *Conn) writePacketCommon(protoID uint32, serialNo uint32, body []byte, sha1Hash [20]byte) error {
	conn := c.currentConn()
	if conn == nil {
		return fmt.Errorf("write packet: %w", ErrNotConnected)
	}

	if len(body) > MaxPacketSize {
		return NewError(CodePacketTooBig, fmt.Sprintf("body too large: %d bytes (max %d)", len(body), MaxPacketSize))
	}
	if len(body) == 0 {
		return NewError(CodeInvalidPacket, "empty packet body")
	}

	header := make([]byte, HeaderLen)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:], protoID)
	header[6] = 0
	header[7] = ProtoVersion
	binary.LittleEndian.PutUint32(header[8:], serialNo)
	binary.LittleEndian.PutUint32(header[12:], uint32(len(body)))
	copy(header[16:36], sha1Hash[:])

	// Reserve the response slot before the request hits the wire. A fast server
	// can reply before the caller reaches ReadResponse; Dispatch treats this
	// nil reservation as "response expected, reader not yet registered" and
	// buffers the packet instead of dropping it. All writes are paired with a
	// ReadResponse, which removes the reservation.
	c.dispMu.Lock()
	if _, exists := c.disp[serialNo]; !exists {
		c.disp[serialNo] = nil
	}
	c.dispMu.Unlock()

	clearReservation := func() {
		c.dispMu.Lock()
		if ch, ok := c.disp[serialNo]; ok && ch == nil {
			delete(c.disp, serialNo)
		}
		c.dispMu.Unlock()
	}

	if _, err := conn.Write(header); err != nil {
		clearReservation()
		return fmt.Errorf("write header: %w", err)
	}

	if len(body) > 0 {
		if _, err := conn.Write(body); err != nil {
			clearReservation()
			return fmt.Errorf("write body: %w", err)
		}
	}

	return nil
}
