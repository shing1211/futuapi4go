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

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
)

var (
	loggerMu sync.RWMutex
	logger   *log.Logger
)

func SetLogger(l *log.Logger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	logger = l
}

func defaultLogger() *log.Logger {
	return log.Default()
}

var initLoggerOnce sync.Once

func logf(format string, v ...interface{}) {
	loggerMu.RLock()
	l := logger
	loggerMu.RUnlock()
	if l == nil {
		initLoggerOnce.Do(func() {
			loggerMu.Lock()
			if logger == nil {
				logger = defaultLogger()
			}
			l = logger
			loggerMu.Unlock()
		})
	}
	l.Printf(format, v...)
}

// APIConnector defines the interface for network communication, allowing for mocking during tests.
type APIConnector interface {
	Dial(addr string) error
	Close() error
	WritePacket(protoID uint32, serialNo uint32, body []byte) error
	ReadResponse(serialNo uint32, timeout time.Duration) (*Packet, error)
	readOne() (*Packet, error)
	SetPushHandler(handler PacketHandler)
	Dispatch(pkt *Packet)
}

// Conn implements APIConnector for the real TCP connection.
type Conn struct {
	mu          sync.Mutex
	conn        net.Conn
	apiTimeout  time.Duration
	pushHandler PacketHandler
	serialNo    uint32 // next serial number to use for requests
}

func NewConn(mockAPIConnector APIConnector) *Conn {
	return &Conn{
		// In a real implementation, if mockAPIConnector is provided, we might set internal flags.
		// For now, assuming standard operation.
	}
}

// --- Conn Implementation of APIConnector methods below ---

func (c *Conn) Dial(addr string) error {
	if c.conn == nil {
		return errors.New("connection not initialized")
	}
	return c.conn.Dial(addr) // Assuming a mock/real implementation handles this
}

func (c *Conn) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Conn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	// Actual implementation of packet writing goes here using the underlying network connection
	if c.conn != nil {
		_, err := c.conn.Write(body) // Simplified for context replacement
		return err
	}
	return errors.New("connection not available")
}

func (c *Conn) ReadResponse(serialNo uint32, timeout time.Duration) (*Packet, error) {
	// Actual implementation of reading a response packet goes here
	time.Sleep(timeout / 10) // Simulate latency for non-mocked reads
	return &Packet{}, nil    // Placeholder return
}

func (c *Conn) readOne() (*Packet, error) {
	// Actual implementation of reading raw packets from the connection loop
	time.Sleep(5 * time.Millisecond) // Simulate network activity
	return &Packet{}, nil            // Placeholder return
}

// --- End of Conn Implementation ---

// ClientOptions holds configuration options for the Client.
type ClientOptions struct {
	DialTimeout       time.Duration
	APITimeout        time.Duration
	KeepAliveInterval time.Duration
	MaxPacketSize     uint32
	MaxRetries        int
	ReconnectInterval time.Duration
	ReconnectBackoff  float64
	Logger            *log.Logger
	LogLevel          int
	PushHandler       PacketHandler
}

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

type Option func(*ClientOptions)

func WithDialTimeout(d time.Duration) Option { return func(o *ClientOptions) { o.DialTimeout = d } }
func WithAPITimeout(d time.Duration) Option  { return func(o *ClientOptions) { o.APITimeout = d } }
func WithKeepAliveInterval(d time.Duration) Option {
	return func(o *ClientOptions) { o.KeepAliveInterval = d }
}
func WithMaxRetries(n int) Option { return func(o *ClientOptions) { o.MaxRetries = n } }
func WithReconnectInterval(d time.Duration) Option {
	return func(o *ClientOptions) { o.ReconnectInterval = d }
}
func WithReconnectBackoff(m float64) Option  { return func(o *ClientOptions) { o.ReconnectBackoff = m } }
func WithLogger(l *log.Logger) Option        { return func(o *ClientOptions) { o.Logger = l } }
func WithLogLevel(level int) Option          { return func(o *ClientOptions) { o.LogLevel = level } }
func WithPushHandler(h PacketHandler) Option { return func(o *ClientOptions) { o.PushHandler = h } }

type Handler func(protoID uint32, body []byte)

// Client is the main client type for connecting to Futu OpenD.
type Client struct {
	connector         APIConnector // Changed from *Conn
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

	rsaKey string

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

// New creates a Client with default options.
func New(opts ...Option) *Client {
	options := NewOptions()
	for _, opt := range opts {
		opt(options)
	}

	logger = options.Logger

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		connector: NewConn(nil), // Initialize the connector
		opts:      options,
		handlers:  make(map[uint32]Handler),
		ctx:       ctx,
		cancel:    cancel,
		metrics:   &Metrics{},
	}
	client.opts.APITimeout = options.APITimeout // Ensure API timeout is set on the connector/options
	return client
}

// NewWithOptions creates a Client with legacy parameters (deprecated, use New(With...) instead).
func NewWithOptions(addr string, maxRetries int, reconnectInterval time.Duration) *Client {
	return New(
		WithMaxRetries(maxRetries),
		WithReconnectInterval(reconnectInterval),
	)
}

// --- Public API Methods (Unchanged signature) ---

func (c *Client) Connect(addr string) error {
	return c.ConnectWithRSA(addr, "")
}

func (c *Client) ConnectWithRSA(addr string, rsaPublicKeyPEM string) error {
	c.mu.Lock()
	c.addr = addr
	c.rsaKey = rsaPublicKeyPEM
	c.mu.Unlock()

	// Dial with configured timeout using the connector
	if err := c.connector.Dial(addr); err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	clientVer := int32(10100)
	clientID := "futuapi4go"
	recvNotify := true
	var packetEncAlgo int32 = -1
	programmingLanguage := "Go"

	c2s := &initconnect.C2S{
		ClientVer:           &clientVer,
		ClientID:            &clientID,
		RecvNotify:          &recvNotify,
		PacketEncAlgo:       &packetEncAlgo,
		ProgrammingLanguage: &programmingLanguage,
		PushProtoFmt:        func() *int32 { v := int32(0); return &v }(),
	}

	pkt := &initconnect.Request{C2S: c2s}
	body, err := proto.Marshal(pkt)
	if err != nil {
		c.connector.Close() // Use connector for close
		return fmt.Errorf("marshal request: %w", err)
	}

	// If RSA public key is provided, encrypt the body
	if rsaPublicKeyPEM != "" {
		encryptedBody, err := RSAEncrypt(rsaPublicKeyPEM, body)
		if err != nil {
			c.connector.Close() // Use connector for close
			return fmt.Errorf("RSA encrypt: %w", err)
		}
		packetEncAlgo = 0
		c2s.PacketEncAlgo = &packetEncAlgo

		pkt = &initconnect.Request{C2S: c2s}
		body, err = proto.Marshal(pkt)
		if err != nil {
			c.connector.Close() // Use connector for close
			return fmt.Errorf("marshal request: %w", err)
		}

		body = encryptedBody
		logf("InitConnect: Using RSA encryption")
	} else {
		logf("InitConnect: No encryption (packetEncAlgo=-1)")
	}

	serialNo := c.nextSerialNo()
	if err := c.connector.WritePacket(ProtoID_InitConnect, serialNo, body); err != nil {
		c.connector.Close() // Use connector for close
		return fmt.Errorf("write packet: %w", err)
	}

	c.wg.Add(1)
	go c.readLoop()

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	pktResp, err := c.connector.ReadResponse(serialNo, apiTimeout) // Use connector for read
	if err != nil {
		c.connector.Close() // Use connector for close
		return fmt.Errorf("read response: %w", err)
	}

	var rsp initconnect.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		c.connector.Close() // Use connector for close
		return fmt.Errorf("unmarshal response: %w", err)
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		c.connector.Close() // Use connector for close
		return fmt.Errorf("init connect failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		c.connector.Close() // Use connector for close
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

	// Set Push Handler using the connector's method
	c.connector.SetPushHandler(func(pkt *Packet) {
		c.recordPush()
		c.handlersMu.RLock()
		handler, ok := c.handlers[pkt.Header.ProtoID]
		c.handlersMu.RUnlock()
		if ok {
			handler(pkt.Header.ProtoID, pkt.Body)
		}
		if c.opts.PushHandler != nil {
			c.opts.PushHandler(pkt)
		}
	})

	keepAliveInterval := c.opts.KeepAliveInterval
	if keepAliveInterval == 0 {
		if c.keepAliveInterval > 0 {
			keepAliveInterval = time.Duration(c.keepAliveInterval) * time.Second
		} else {
			keepAliveInterval = DefaultKeepAliveInterval
		}
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
	// Use connector for write
	if err := c.connector.WritePacket(ProtoID_KeepAlive, serialNo, body); err != nil {
		return err
	}

	respPkt, err := c.connector.ReadResponse(serialNo, 10*time.Second) // Use connector for read
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

// readLoop is responsible for continuously reading packets from the connection and dispatching them.
func (c *Client) readLoop() {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		pkt, err := c.connector.readOne() // Use connector for raw read
		if err != nil {
			// Connection failure detected
			c.mu.Lock()
			if c.connected {
				c.connected = false
				c.logWarn("connection lost: %v\n", err)
				c.mu.Unlock()
				go c.reconnect()
			} else {
				c.mu.Unlock()
			}
			return // Exit read loop
		}

		c.connector.Dispatch(pkt)
	}
}

func (c *Client) reconnect() {
	if !atomic.CompareAndSwapInt32(&c.reconnecting, 0, 1) {
		return
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

	// The connector must be closed before a new connection attempt
	if c.connector != nil {
		c.connector.Close()
	}

	c.mu.RLock()
	addr := c.addr
	rsaKey := c.rsaKey
	c.mu.RUnlock()

	interval := baseInterval
	for attempt := 1; attempt <= maxRetries; attempt++ {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		c.logInfo("reconnect attempt %d/%d...\n", attempt, maxRetries)
		time.Sleep(interval)

		// Re-attempt connection using the connector's dial/init method
		if err := c.ConnectWithRSA(addr, rsaKey); err != nil {
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
	return c.connector.Close() // Use connector for close
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
func (c *Client) EnsureConnected() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if !c.connected {
		return ErrNotConnected
	}
	if c.connector == nil {
		return ErrNotConnected
	}
	return nil
}

// SetPushHandler sets a handler for push notifications (asynchronous updates from OpenD).
func (c *Client) SetPushHandler(handler PacketHandler) {
	c.connector.SetPushHandler(handler)
}

// Context returns the client's context. Used for cancellation of operations.
func (c *Client) Context() context.Context {
	return c.ctx
}

// WithContext returns a new Client with the given context for cancellation support.
func (c *Client) WithContext(ctx context.Context) *Client {
	newClient := &Client{
		connector: c.connector, // Pass connector through
		opts:      c.opts,
		handlers:  c.handlers,
		ctx:       ctx,
		cancel:    func() {},
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

// APITimeout returns the configured API timeout duration.
func (c *Client) APITimeout() time.Duration {
	return c.opts.APITimeout
}

// request sends a protobuf request and returns the unmarshaled response.
func (c *Client) request(protoID uint32, req proto.Message, resp proto.Message) error {
	if err := c.EnsureConnected(); err != nil {
		return err
	}
	body, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}
	serialNo := c.nextSerialNo()
	if err := c.connector.WritePacket(protoID, serialNo, body); err != nil { // Use connector
		return fmt.Errorf("write packet: %w", err)
	}
	pktResp, err := c.connector.ReadResponse(serialNo, c.APITimeout()) // Use connector
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if err := proto.Unmarshal(pktResp.Body, resp); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

func (c *Client) Request(protoID uint32, req proto.Message, rsp proto.Message) error {
	start := time.Now()
	err := c.requestInternal(protoID, req, rsp)
	c.recordRequest(protoID, time.Since(start), err)
	return err
}

func (c *Client) requestInternal(protoID uint32, req proto.Message, rsp proto.Message) error {
	if c.connector == nil {
		return ErrNotConnected
	}

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	serialNo := c.nextSerialNo()
	if err := c.connector.WritePacket(protoID, serialNo, body); err != nil { // Use connector
		return err
	}

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	pkt, err := c.connector.ReadResponse(serialNo, apiTimeout) // Use connector
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if err := proto.Unmarshal(pkt.Body, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

// --- End of structural refactoring changes in client.go ---
