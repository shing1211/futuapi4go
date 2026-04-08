package futuapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pkg/pb/common"
	"gitee.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"gitee.com/shing1211/futuapi4go/pkg/pb/keepalive"
)

var (
	logger *log.Logger
)

func SetLogger(l *log.Logger) {
	logger = l
}

func defaultLogger() *log.Logger {
	return log.Default()
}

func logf(format string, v ...interface{}) {
	if logger == nil {
		logger = defaultLogger()
	}
	logger.Printf(format, v...)
}

const (
	ProtoID_InitConnect    = 1001
	ProtoID_KeepAlive      = 1002
	ProtoID_GetGlobalState = 1004
)

const (
	DefaultTimeout           = 30 * time.Second
	DefaultKeepAliveInterval = 30 * time.Second
	DefaultMaxRetries        = 3
	DefaultReconnectInterval = 3 * time.Second
)

type Client struct {
	conn              *Conn
	mu                sync.RWMutex
	connID            uint64
	aesKey            string
	serverVer         int32
	keepAliveInterval int32
	serialNo          uint32
	serialMu          sync.Mutex
	handlers          map[uint32]Handler
	handlersMu        sync.RWMutex
	ctx               context.Context
	cancel            context.CancelFunc
	wg                sync.WaitGroup
	connected         bool

	addr              string
	maxRetries        int
	reconnectInterval time.Duration
	reconnecting      int32 // atomic flag: 0 = not reconnecting, 1 = reconnecting
}

type Handler func(protoID uint32, body []byte)

func New() *Client {
	ctx, cancel := context.WithCancel(context.Background())
	return &Client{
		conn:              NewConn(nil),
		handlers:          make(map[uint32]Handler),
		ctx:               ctx,
		cancel:            cancel,
		maxRetries:        DefaultMaxRetries,
		reconnectInterval: DefaultReconnectInterval,
	}
}

func NewWithOptions(addr string, maxRetries int, reconnectInterval time.Duration) *Client {
	ctx, cancel := context.WithCancel(context.Background())
	if maxRetries == 0 {
		maxRetries = DefaultMaxRetries
	}
	if reconnectInterval == 0 {
		reconnectInterval = DefaultReconnectInterval
	}
	client := &Client{
		conn:              NewConn(nil),
		handlers:          make(map[uint32]Handler),
		ctx:               ctx,
		cancel:            cancel,
		maxRetries:        maxRetries,
		reconnectInterval: reconnectInterval,
	}
	if addr != "" {
		if err := client.Connect(addr); err != nil {
			return nil
		}
	}
	return client
}

func (c *Client) Connect(addr string) error {
	return c.ConnectWithRSA(addr, "")
}

func (c *Client) ConnectWithRSA(addr string, rsaPublicKeyPEM string) error {
	c.mu.Lock()
	c.addr = addr
	c.mu.Unlock()

	if err := c.conn.Dial(addr); err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	clientVer := int32(10100)
	clientID := "futuapi4go"
	recvNotify := true
	var packetEncAlgo int32 = -1  // Default: no encryption
	programmingLanguage := "Go"

	c2s := &initconnect.C2S{
		ClientVer:           &clientVer,
		ClientID:            &clientID,
		RecvNotify:          &recvNotify,
		PacketEncAlgo:       &packetEncAlgo,
		ProgrammingLanguage: &programmingLanguage,
		PushProtoFmt:        func() *int32 { v := int32(0); return &v }(), // Explicitly set Protobuf format
	}

	pkt := &initconnect.Request{
		C2S: c2s,
	}

	body, err := proto.Marshal(pkt)
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("marshal request: %w", err)
	}

	// If RSA public key is provided, encrypt the body
	if rsaPublicKeyPEM != "" {
		encryptedBody, err := RSAEncrypt(rsaPublicKeyPEM, body)
		if err != nil {
			c.conn.Close()
			return fmt.Errorf("RSA encrypt: %w", err)
		}
		// Set encryption algorithm to FTAES_ECB (0) as per protocol spec
		packetEncAlgo = 0
		c2s.PacketEncAlgo = &packetEncAlgo
		
		// Re-marshal with encryption flag set
		pkt = &initconnect.Request{C2S: c2s}
		body, err = proto.Marshal(pkt)
		if err != nil {
			c.conn.Close()
			return fmt.Errorf("marshal request: %w", err)
		}
		
		// Replace body with encrypted version
		body = encryptedBody
		logf("InitConnect: Using RSA encryption")
	} else {
		logf("InitConnect: No encryption (packetEncAlgo=-1)")
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(ProtoID_InitConnect, serialNo, body); err != nil {
		c.conn.Close()
		return fmt.Errorf("write packet: %w", err)
	}

	c.conn.SetReadDeadline(time.Now().Add(DefaultTimeout))
	respPkt, err := c.conn.ReadPacket()
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("read response: %w", err)
	}

	var rsp initconnect.Response
	if err := proto.Unmarshal(respPkt.Body, &rsp); err != nil {
		c.conn.Close()
		return fmt.Errorf("unmarshal response: %w", err)
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		c.conn.Close()
		return fmt.Errorf("init connect failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		c.conn.Close()
		return errors.New("init connect: s2c is nil")
	}

	c.mu.Lock()
	c.connID = s2c.GetConnID()
	c.aesKey = s2c.GetConnAESKey()
	c.serverVer = s2c.GetServerVer()
	c.keepAliveInterval = s2c.GetKeepAliveInterval()
	c.connected = true
	c.mu.Unlock()

	// Set up push notification dispatcher
	c.conn.SetPushHandler(func(pkt *Packet) {
		c.handlersMu.RLock()
		handler, ok := c.handlers[pkt.Header.ProtoID]
		c.handlersMu.RUnlock()
		if ok {
			handler(pkt.Header.ProtoID, pkt.Body)
		}
	})

	if c.keepAliveInterval > 0 {
		interval := time.Duration(c.keepAliveInterval) * time.Second
		if interval < DefaultKeepAliveInterval {
			interval = DefaultKeepAliveInterval
		}
		c.wg.Add(1)
		go c.keepAliveLoop(interval)
	}

	return nil
}

func (c *Client) keepAliveLoop(interval time.Duration) {
	defer c.wg.Done()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if err := c.keepAlive(); err != nil {
				logf("keepalive error: %v", err)
			}
		}
	}
}

func (c *Client) keepAlive() error {
	req := &keepalive.C2S{}
	pkt := &keepalive.Request{C2S: req}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return err
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(ProtoID_KeepAlive, serialNo, body); err != nil {
		return err
	}

	c.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	respPkt, err := c.conn.ReadPacket()
	if err != nil {
		return err
	}

	var rsp keepalive.Response
	if err := proto.Unmarshal(respPkt.Body, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return fmt.Errorf("keepalive failed: retType=%d", rsp.GetRetType())
	}

	return nil
}

func (c *Client) nextSerialNo() uint32 {
	c.serialMu.Lock()
	c.serialNo++
	no := c.serialNo
	c.serialMu.Unlock()
	return no
}

func (c *Client) readLoop() {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		pkt, err := c.conn.ReadPacket()
		if err != nil {
			c.mu.Lock()
			if c.connected {
				c.connected = false
				logf("connection lost: %v\n", err)
				go c.reconnect()
			}
			c.mu.Unlock()
			continue
		}

		c.handlersMu.RLock()
		handler, ok := c.handlers[pkt.Header.ProtoID]
		c.handlersMu.RUnlock()

		if ok {
			handler(pkt.Header.ProtoID, pkt.Body)
		}
	}
}

func (c *Client) reconnect() {
	// Atomically check and set reconnecting flag to prevent TOCTOU race
	if !atomic.CompareAndSwapInt32(&c.reconnecting, 0, 1) {
		return // Already reconnecting
	}
	defer atomic.StoreInt32(&c.reconnecting, 0)

	for attempt := 1; attempt <= c.maxRetries; attempt++ {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		logf("reconnect attempt %d/%d...\n", attempt, c.maxRetries)
		time.Sleep(c.reconnectInterval)

		if err := c.Connect(c.addr); err != nil {
			logf("reconnect failed: %v\n", err)
			continue
		}

		logf("reconnected successfully\n")
		return
	}

	logf("reconnect failed: max retries exceeded\n")
}

func (c *Client) RegisterHandler(protoID uint32, handler Handler) {
	c.handlersMu.Lock()
	c.handlers[protoID] = handler
	c.handlersMu.Unlock()
}

func (c *Client) Close() error {
	c.cancel()
	c.wg.Wait()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) GetConnID() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connID
}

func (c *Client) GetAESKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.aesKey
}

func (c *Client) GetServerVer() int32 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.serverVer
}

func (c *Client) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connected
}

// EnsureConnected returns an error if the client is not connected.
// This should be called by all public API functions before making requests.
func (c *Client) EnsureConnected() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if !c.connected {
		return ErrNotConnected
	}
	if c.conn == nil {
		return ErrNotConnected
	}
	return nil
}

// SetPushHandler sets a handler for push notifications (asynchronous updates from OpenD).
// Push notifications are packets with serial numbers that don't match any request.
func (c *Client) SetPushHandler(handler PacketHandler) {
	c.conn.SetPushHandler(handler)
}

// Context returns the client's context. Used for cancellation of operations.
func (c *Client) Context() context.Context {
	return c.ctx
}

// WithContext returns a new Client with the given context for cancellation support.
// The original client remains usable. Operations will respect the context's deadline/cancellation.
func (c *Client) WithContext(ctx context.Context) *Client {
	newClient := &Client{
		conn:              c.conn,
		handlers:          c.handlers,
		ctx:               ctx,
		cancel:            func() {}, // Don't cancel parent context
		maxRetries:        c.maxRetries,
		reconnectInterval: c.reconnectInterval,
	}
	newClient.mu.RLock()
	newClient.connID = c.connID
	newClient.aesKey = c.aesKey
	newClient.serverVer = c.serverVer
	newClient.keepAliveInterval = c.keepAliveInterval
	newClient.connected = c.connected
	newClient.mu.RUnlock()
	return newClient
}

func (c *Client) Conn() *Conn {
	return c.conn
}

func (c *Client) NextSerialNo() uint32 {
	return c.nextSerialNo()
}

func (c *Client) Request(protoID uint32, req proto.Message, rsp proto.Message) error {
	if c.conn == nil {
		return ErrNotConnected
	}

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(protoID, serialNo, body); err != nil {
		return err
	}

	c.conn.SetReadDeadline(time.Now().Add(DefaultTimeout))
	pkt, err := c.conn.ReadPacket()
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if err := proto.Unmarshal(pkt.Body, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
