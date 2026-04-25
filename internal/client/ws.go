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
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func DialWebSocket(ctx context.Context, addr string, secretKey string) (*websocket.Conn, *http.Response, error) {
	url := "ws://" + addr
	if secretKey != "" {
		url += "?secret_key=" + secretKey
	}
	return websocket.DefaultDialer.DialContext(ctx, url, nil)
}

func DialWebSocketSecure(ctx context.Context, addr string, secretKey string) (*websocket.Conn, *http.Response, error) {
	url := "wss://" + addr
	if secretKey != "" {
		url += "?secret_key=" + secretKey
	}
	return websocket.DefaultDialer.DialContext(ctx, url, nil)
}

func IsWebSocketAddr(addr string) bool {
	return strings.HasPrefix(addr, "ws://") ||
		strings.HasPrefix(addr, "wss://") ||
		strings.HasSuffix(addr, "/ws")
}

type wsConn struct {
	conn   *websocket.Conn
	apiTimeout time.Duration

	dispMu   sync.Mutex
	disp     map[uint32]chan *Packet
	dispSize int

	pushHandler PacketHandler

	readCh  chan *Packet
	closeCh chan struct{}
	wg      sync.WaitGroup
}

func newWSConn(conn *websocket.Conn) *wsConn {
	ws := &wsConn{
		conn:   conn,
		disp:   make(map[uint32]chan *Packet),
		readCh:  make(chan *Packet, 10),
		closeCh: make(chan struct{}),
	}

	ws.wg.Add(1)
	go ws.readPump()

	return ws
}

func (c *wsConn) readPump() {
	defer c.wg.Done()

	for {
		select {
		case <-c.closeCh:
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		var data []byte
		var err error
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("websocket panic: %v", r)
				}
			}()
			_, data, err = c.conn.ReadMessage()
		}()

		if err != nil {
			if c.isTimeout(err) {
				continue
			}
			select {
			case <-c.closeCh:
				return
			default:
			}
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}
			return
		}

		pkt, err := c.unpackPacket(data)
		if err != nil {
			continue
		}

		c.Dispatch(pkt)
	}
}

func (c *wsConn) isTimeout(err error) bool {
	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		return true
	}
	if strings.Contains(err.Error(), "timeout") {
		return true
	}
	return false
}

func (c *wsConn) unpackPacket(data []byte) (*Packet, error) {
	if len(data) < HeaderLen {
		return nil, fmt.Errorf("data too small for header")
	}

	var h Header
	copy(h.Magic[:], data[0:2])
	h.ProtoID = binary.LittleEndian.Uint32(data[2:6])
	h.ProtoFmt = byte(data[6])
	h.ProtoVer = data[7]
	h.SerialNo = binary.LittleEndian.Uint32(data[8:12])
	h.BodyLen = binary.LittleEndian.Uint32(data[12:16])
	copy(h.BodySHA1[:], data[16:36])
	copy(h.Reserved[:], data[36:44])

	if string(h.Magic[:]) != "FT" {
		return nil, fmt.Errorf("invalid magic: % x", h.Magic)
	}

	if h.BodyLen > MaxPacketSize {
		return nil, fmt.Errorf("body too large: %d", h.BodyLen)
	}

	if len(data) < int(HeaderLen+h.BodyLen) {
		return nil, fmt.Errorf("data too small for body")
	}

	body := data[HeaderLen : HeaderLen+h.BodyLen]

	actualSHA1 := sha1.Sum(body)
	if actualSHA1 != h.BodySHA1 {
		return nil, fmt.Errorf("checksum mismatch")
	}

	return &Packet{Header: h, Body: body}, nil
}

func (c *wsConn) packPacket(protoID uint32, serialNo uint32, body []byte) []byte {
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

	pkt := make([]byte, HeaderLen+len(body))
	copy(pkt, header)
	copy(pkt[HeaderLen:], body)
	return pkt
}

func (c *wsConn) SetPushHandler(handler PacketHandler) {
	c.pushHandler = handler
}

func (c *wsConn) APITimeout() time.Duration {
	return c.apiTimeout
}

func (c *wsConn) SetAPITimeout(timeout time.Duration) {
	c.apiTimeout = timeout
}

func (c *wsConn) Close() error {
	select {
	case <-c.closeCh:
		// already closed
	default:
		close(c.closeCh)
	}
	c.wg.Wait()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *wsConn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	if len(body) > MaxPacketSize {
		return fmt.Errorf("body too large: %d", len(body))
	}
	if len(body) == 0 {
		return fmt.Errorf("empty body")
	}

	pkt := c.packPacket(protoID, serialNo, body)
	return c.conn.WriteMessage(websocket.BinaryMessage, pkt)
}

func (c *wsConn) ReadPacket() (*Packet, error) {
	select {
	case pkt, ok := <-c.readCh:
		if !ok {
			return nil, fmt.Errorf("connection closed")
		}
		return pkt, nil
	case <-c.closeCh:
		return nil, fmt.Errorf("connection closed")
	}
}

func (c *wsConn) ReadResponse(serialNo uint32, timeout time.Duration) (*Packet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return c.ReadResponseContext(ctx, serialNo, timeout)
}

func (c *wsConn) ReadResponseContext(ctx context.Context, serialNo uint32, timeout time.Duration) (*Packet, error) {
	ch := make(chan *Packet, 1)

	c.dispMu.Lock()
	c.disp[serialNo] = ch
	c.dispSize = len(c.disp)
	c.dispMu.Unlock()

	defer func() {
		c.dispMu.Lock()
		delete(c.disp, serialNo)
		c.dispMu.Unlock()
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case pkt := <-ch:
		return pkt, nil
	}
}

func (c *wsConn) Dispatch(pkt *Packet) {
	c.dispMu.Lock()
	defer c.dispMu.Unlock()

	ch, ok := c.disp[pkt.Header.SerialNo]
	if ok {
		select {
		case ch <- pkt:
		default:
		}
		return
	}

	if c.pushHandler != nil {
		c.pushHandler(pkt)
	}
}

func (c *wsConn) Dial(addr string) error {
	return fmt.Errorf("Dial not supported for WebSocket")
}

func (c *wsConn) readOne() (*Packet, error) {
	return c.ReadPacket()
}