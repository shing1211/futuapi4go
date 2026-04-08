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

// logInfo logs at info level if log level allows.
func (c *Client) logInfo(format string, v ...interface{}) {
	if c.opts.LogLevel > 0 {
		return
	}
	l := c.opts.Logger
	if l == nil {
		l = defaultLogger()
	}
	l.Printf(format, v...)
}

// logWarn logs at warn level if log level allows.
func (c *Client) logWarn(format string, v ...interface{}) {
	if c.opts.LogLevel > 1 {
		return
	}
	l := c.opts.Logger
	if l == nil {
		l = defaultLogger()
	}
	l.Printf(format, v...)
}

// logError logs at error level if log level allows.
func (c *Client) logError(format string, v ...interface{}) {
	if c.opts.LogLevel > 2 {
		return
	}
	l := c.opts.Logger
	if l == nil {
		l = defaultLogger()
	}
	l.Printf(format, v...)
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
	DefaultDialTimeout       = 10 * time.Second
)

// ClientOptions holds configuration options for the Client.
// Use NewOptions() for sensible defaults, then modify as needed.
type ClientOptions struct {
	// Connection settings
	DialTimeout       time.Duration // Timeout for initial TCP dial
	APITimeout        time.Duration // Default timeout for API calls
	KeepAliveInterval time.Duration // Interval between keepalive pings
	MaxPacketSize     uint32        // Maximum packet size (default 10MB)

	// Reconnection settings
	MaxRetries        int           // Max reconnection attempts
	ReconnectInterval time.Duration // Base interval between reconnect attempts
	ReconnectBackoff  float64       // Multiplier for backoff (1.0 = no backoff)

	// Logging
	Logger   *log.Logger // Custom logger (nil = use default)
	LogLevel int         // Log level: 0=Info, 1=Warn, 2=Error, 3=Silent

	// Push notifications
	PushHandler PacketHandler // Handler for incoming push notifications
}

// NewOptions returns ClientOptions with sensible defaults.
func NewOptions() *ClientOptions {
	return &ClientOptions{
		DialTimeout:       DefaultDialTimeout,
		APITimeout:        DefaultTimeout,
		KeepAliveInterval: DefaultKeepAliveInterval,
		MaxPacketSize:     10 * 1024 * 1024,
		MaxRetries:        DefaultMaxRetries,
		ReconnectInterval: DefaultReconnectInterval,
		ReconnectBackoff:  1.5,
		Logger:            nil,
		LogLevel:          0,
		PushHandler:       nil,
	}
}

// Option is a functional option for configuring Client.
type Option func(*ClientOptions)

// WithDialTimeout sets the TCP dial timeout.
func WithDialTimeout(d time.Duration) Option {
	return func(o *ClientOptions) { o.DialTimeout = d }
}

// WithAPITimeout sets the default API call timeout.
func WithAPITimeout(d time.Duration) Option {
	return func(o *ClientOptions) { o.APITimeout = d }
}

// WithKeepAliveInterval sets the keepalive ping interval.
func WithKeepAliveInterval(d time.Duration) Option {
	return func(o *ClientOptions) { o.KeepAliveInterval = d }
}

// WithMaxRetries sets the maximum reconnection attempts.
func WithMaxRetries(n int) Option {
	return func(o *ClientOptions) { o.MaxRetries = n }
}

// WithReconnectInterval sets the base reconnect interval.
func WithReconnectInterval(d time.Duration) Option {
	return func(o *ClientOptions) { o.ReconnectInterval = d }
}

// WithReconnectBackoff sets the backoff multiplier for reconnection.
func WithReconnectBackoff(m float64) Option {
	return func(o *ClientOptions) { o.ReconnectBackoff = m }
}

// WithLogger sets a custom logger.
func WithLogger(l *log.Logger) Option {
	return func(o *ClientOptions) { o.Logger = l }
}

// WithLogLevel sets the log level (0=Info, 1=Warn, 2=Error, 3=Silent).
func WithLogLevel(level int) Option {
	return func(o *ClientOptions) { o.LogLevel = level }
}

// WithPushHandler sets a handler for push notifications.
func WithPushHandler(h PacketHandler) Option {
	return func(o *ClientOptions) { o.PushHandler = h }
}

type Client struct {
	conn              *Conn
	mu                sync.RWMutex
	opts              *ClientOptions
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
	connActive        int32 // atomic flag: 0 = loops should exit, 1 = loops running

	addr         string
	reconnecting int32 // atomic flag: 0 = not reconnecting, 1 = reconnecting

	// Metrics / 指標
	metrics   *Metrics
	metricsMu sync.RWMutex
}

// Metrics tracks client performance statistics.
type Metrics struct {
	TotalRequests    uint64
	SuccessfulReqs   uint64
	FailedReqs       uint64
	TotalLatencyMs   uint64
	LastRequestTime  time.Time
	LastErrorCode    string
	LastErrorMessage string
	ReconnectCount   uint64
	ConnectedSince   time.Time
	PushReceived     uint64
}

// GetMetrics returns a copy of current metrics.
func (c *Client) GetMetrics() Metrics {
	c.metricsMu.RLock()
	defer c.metricsMu.RUnlock()
	return *c.metrics
}

func (c *Client) recordRequest(protoID uint32, duration time.Duration, err error) {
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	c.metrics.TotalRequests++
	c.metrics.LastRequestTime = time.Now()
	c.metrics.TotalLatencyMs += uint64(duration.Milliseconds())
	if err != nil {
		c.metrics.FailedReqs++
		c.metrics.LastErrorCode = fmt.Sprintf("%d", protoID)
		c.metrics.LastErrorMessage = err.Error()
	} else {
		c.metrics.SuccessfulReqs++
	}
}

func (c *Client) recordReconnect() {
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	c.metrics.ReconnectCount++
}

func (c *Client) recordPush() {
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	c.metrics.PushReceived++
}

type Handler func(protoID uint32, body []byte)

// New creates a Client with default options.
func New(opts ...Option) *Client {
	options := NewOptions()
	for _, opt := range opts {
		opt(options)
	}

	logger = options.Logger

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:     NewConn(nil),
		opts:     options,
		handlers: make(map[uint32]Handler),
		ctx:      ctx,
		cancel:   cancel,
		metrics:  &Metrics{},
	}
	client.conn.apiTimeout = options.APITimeout
	return client
}

// NewWithOptions creates a Client with legacy parameters (deprecated, use New(With...) instead).
func NewWithOptions(addr string, maxRetries int, reconnectInterval time.Duration) *Client {
	return New(
		WithMaxRetries(maxRetries),
		WithReconnectInterval(reconnectInterval),
	)
}

func (c *Client) Connect(addr string) error {
	return c.ConnectWithRSA(addr, "")
}

func (c *Client) ConnectWithRSA(addr string, rsaPublicKeyPEM string) error {
	c.mu.Lock()
	c.addr = addr
	c.mu.Unlock()

	// Dial with configured timeout
	if err := c.conn.Dial(addr); err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	clientVer := int32(10100)
	clientID := "futuapi4go"
	recvNotify := true
	var packetEncAlgo int32 = -1 // Default: no encryption
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

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	respPkt, err := c.conn.ReadResponse(serialNo, apiTimeout)
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
	atomic.StoreInt32(&c.connActive, 1)
	c.metricsMu.Lock()
	c.metrics.ConnectedSince = time.Now()
	c.metricsMu.Unlock()
	c.mu.Unlock()

	// Set up push notification dispatcher and start readLoop
	c.conn.SetPushHandler(func(pkt *Packet) {
		c.recordPush()
		c.handlersMu.RLock()
		handler, ok := c.handlers[pkt.Header.ProtoID]
		c.handlersMu.RUnlock()
		if ok {
			handler(pkt.Header.ProtoID, pkt.Body)
		}
	})

	// Also call user-configured push handler if set
	if c.opts.PushHandler != nil {
		userHandler := c.opts.PushHandler
		c.conn.SetPushHandler(func(pkt *Packet) {
			c.handlersMu.RLock()
			handler, ok := c.handlers[pkt.Header.ProtoID]
			c.handlersMu.RUnlock()
			if ok {
				handler(pkt.Header.ProtoID, pkt.Body)
			}
			userHandler(pkt)
		})
	}

	// Start readLoop for asynchronous push notifications
	c.wg.Add(1)
	go c.readLoop()

	keepAliveInterval := c.opts.KeepAliveInterval
	if keepAliveInterval == 0 {
		if c.keepAliveInterval > 0 {
			keepAliveInterval = time.Duration(c.keepAliveInterval) * time.Second
		} else {
			keepAliveInterval = DefaultKeepAliveInterval
		}
	}
	interval := time.Duration(c.keepAliveInterval) * time.Second
	if interval > 0 && interval < keepAliveInterval {
		interval = keepAliveInterval
	}
	if keepAliveInterval > 0 {
		c.wg.Add(1)
		go c.keepAliveLoop(keepAliveInterval)
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
			if atomic.LoadInt32(&c.connActive) == 0 {
				return
			}
			if err := c.keepAlive(); err != nil {
				c.logWarn("keepalive error: %v\n", err)
			}
		}
	}
}

func (c *Client) keepAlive() error {
	now := time.Now().Unix()
	req := &keepalive.C2S{Time: &now}
	pkt := &keepalive.Request{C2S: req}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return err
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(ProtoID_KeepAlive, serialNo, body); err != nil {
		return err
	}

	respPkt, err := c.conn.ReadResponse(serialNo, 10*time.Second)
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

		if atomic.LoadInt32(&c.connActive) == 0 {
			return
		}

		pkt, err := c.conn.readOne()
		if err != nil {
			c.mu.Lock()
			if c.connected && atomic.LoadInt32(&c.connActive) == 1 {
				c.connected = false
				c.logWarn("connection lost: %v\n", err)
				c.mu.Unlock()
				go c.reconnect()
			} else {
				c.mu.Unlock()
			}
			return
		}

		c.conn.Dispatch(pkt)
	}
}

func (c *Client) reconnect() {
	// Atomically check and set reconnecting flag to prevent TOCTOU race
	if !atomic.CompareAndSwapInt32(&c.reconnecting, 0, 1) {
		return // Already reconnecting
	}
	defer atomic.StoreInt32(&c.reconnecting, 0)
	defer c.recordReconnect()

	maxRetries := c.opts.MaxRetries
	if maxRetries == 0 {
		maxRetries = DefaultMaxRetries
	}
	baseInterval := c.opts.ReconnectInterval
	if baseInterval == 0 {
		baseInterval = DefaultReconnectInterval
	}
	backoff := c.opts.ReconnectBackoff
	if backoff <= 0 {
		backoff = 1.0
	}

	atomic.StoreInt32(&c.connActive, 0)

	interval := baseInterval
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		c.logInfo("reconnect attempt %d/%d...\n", attempt, maxRetries)
		time.Sleep(interval)

		if err := c.Connect(c.addr); err != nil {
			c.logWarn("reconnect failed: %v\n", err)
			interval = time.Duration(float64(interval) * backoff)
			continue
		}

		c.logInfo("reconnected successfully\n")
		return
	}

	c.logError("reconnect failed: max retries exceeded\n")
}

func (c *Client) RegisterHandler(protoID uint32, handler Handler) {
	c.handlersMu.Lock()
	c.handlers[protoID] = handler
	c.handlersMu.Unlock()
}

func (c *Client) Close() error {
	atomic.StoreInt32(&c.connActive, 0)
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
		conn:     c.conn,
		opts:     c.opts,
		handlers: c.handlers,
		ctx:      ctx,
		cancel:   func() {}, // Don't cancel parent context
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
	start := time.Now()
	err := c.requestInternal(protoID, req, rsp)
	c.recordRequest(protoID, time.Since(start), err)
	return err
}

func (c *Client) requestInternal(protoID uint32, req proto.Message, rsp proto.Message) error {
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

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	pkt, err := c.conn.ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if err := proto.Unmarshal(pkt.Body, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
