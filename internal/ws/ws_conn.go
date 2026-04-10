package ws

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

// ============================================================================
// WebSocket Transport for Futu OpenD
// ============================================================================

// WSConn wraps a WebSocket connection to implement the futuapi.ConnInterface
type WSConn struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	disp   map[uint32]chan *futuapi.Packet
	dispMu sync.RWMutex

	pushHandler futuapi.PacketHandler
	pushMu     sync.RWMutex

	apiTimeout time.Duration
	closed     bool
}

// NewWSConn creates a new WebSocket connection
func NewWSConn(wsConn *websocket.Conn, apiTimeout time.Duration) *WSConn {
	return &WSConn{
		conn:       wsConn,
		disp:       make(map[uint32]chan *futuapi.Packet),
		apiTimeout: apiTimeout,
	}
}

// ConnectWS connects to Futu OpenD via WebSocket
func ConnectWS(ctx context.Context, addr string, apiTimeout time.Duration) (*WSConn, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	header := http.Header{}
	// Add any required headers here

	wsConn, _, err := dialer.DialContext(ctx, "ws://"+addr, header)
	if err != nil {
		return nil, fmt.Errorf("WebSocket connection failed: %w", err)
	}

	conn := NewWSConn(wsConn, apiTimeout)

	// Start read loop
	go conn.readLoop()

	return conn, nil
}

// ConnectWSS connects to Futu OpenD via secure WebSocket (wss://)
func ConnectWSS(ctx context.Context, addr string, apiTimeout time.Duration) (*WSConn, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	wsConn, _, err := dialer.DialContext(ctx, "wss://"+addr, nil)
	if err != nil {
		return nil, fmt.Errorf("Secure WebSocket connection failed: %w", err)
	}

	conn := NewWSConn(wsConn, apiTimeout)
	go conn.readLoop()

	return conn, nil
}

// WritePacket writes a packet to the WebSocket
func (w *WSConn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return futuapi.ErrNotConnected
	}

	// Construct binary message (same format as TCP)
	header := make([]byte, 44)
	header[0] = 'F'
	header[1] = 'T'

	// ProtoID (Little Endian)
	header[2] = byte(protoID)
	header[3] = byte(protoID >> 8)
	header[4] = byte(protoID >> 16)
	header[5] = byte(protoID >> 24)

	// ProtoFmt = 0 (Protobuf)
	header[6] = 0

	// ProtoVer = 0
	header[7] = 0

	// SerialNo (Little Endian)
	header[8] = byte(serialNo)
	header[9] = byte(serialNo >> 8)
	header[10] = byte(serialNo >> 16)
	header[11] = byte(serialNo >> 24)

	// BodyLen (Little Endian)
	bodyLen := uint32(len(body))
	header[16] = byte(bodyLen)
	header[17] = byte(bodyLen >> 8)
	header[18] = byte(bodyLen >> 16)
	header[19] = byte(bodyLen >> 24)

	// SHA1 (zeros for now - can be computed if needed)
	// Reserved (zeros)

	// Send as binary message
	message := make([]byte, 44+len(body))
	copy(message[:44], header)
	copy(message[44:], body)

	return w.conn.WriteMessage(websocket.BinaryMessage, message)
}

// ReadResponse waits for a response with the given serial number
func (w *WSConn) ReadResponse(serialNo uint32, timeout time.Duration) (*futuapi.Packet, error) {
	ch := make(chan *futuapi.Packet, 1)

	w.dispMu.Lock()
	w.disp[serialNo] = ch
	w.dispMu.Unlock()

	defer func() {
		w.dispMu.Lock()
		delete(w.disp, serialNo)
		w.dispMu.Unlock()
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case pkt := <-ch:
		return pkt, nil
	case <-timer.C:
		return nil, futuapi.ErrRequestTimeout
	}
}

// SetPushHandler sets the handler for push notifications
func (w *WSConn) SetPushHandler(handler futuapi.PacketHandler) {
	w.pushMu.Lock()
	defer w.pushMu.Unlock()
	w.pushHandler = handler
}

// Close closes the WebSocket connection
func (w *WSConn) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		return nil
	}

	w.closed = true
	return w.conn.Close()
}

// IsClosed returns whether the connection is closed
func (w *WSConn) IsClosed() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.closed
}

// ============================================================================
// Internal methods
// ============================================================================

func (w *WSConn) readLoop() {
	defer func() {
		w.mu.Lock()
		w.closed = true
		w.mu.Unlock()
	}()

	for {
		w.mu.Lock()
		if w.closed {
			w.mu.Unlock()
			return
		}
		w.mu.Unlock()

		msgType, message, err := w.conn.ReadMessage()
		if err != nil {
			return
		}

		if msgType != websocket.BinaryMessage {
			continue
		}

		if len(message) < 44 {
			continue
		}

		// Parse header
		if message[0] != 'F' || message[1] != 'T' {
			continue
		}

		protoID := uint32(message[2]) | uint32(message[3])<<8 | uint32(message[4])<<16 | uint32(message[5])<<24
		serialNo := uint32(message[8]) | uint32(message[9])<<8 | uint32(message[10])<<16 | uint32(message[11])<<24
		bodyLen := uint32(message[16]) | uint32(message[17])<<8 | uint32(message[18])<<16 | uint32(message[19])<<24

		if len(message) < 44+int(bodyLen) {
			continue
		}

		body := message[44 : 44+bodyLen]

		pkt := &futuapi.Packet{
			Header: futuapi.Header{
				ProtoID:  protoID,
				SerialNo: serialNo,
				BodyLen:  bodyLen,
			},
			Body: body,
		}

		// Try to dispatch to waiting reader
		w.dispMu.RLock()
		ch, ok := w.disp[serialNo]
		w.dispMu.RUnlock()

		if ok {
			select {
			case ch <- pkt:
				continue
			default:
			}
		}

		// No waiting reader, treat as push notification
		w.pushMu.RLock()
		handler := w.pushHandler
		w.pushMu.RUnlock()

		if handler != nil {
			handler(pkt)
		}
	}
}
