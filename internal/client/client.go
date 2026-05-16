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
	"crypto/sha256"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/tracing"

	"github.com/gorilla/websocket"
	"github.com/shing1211/futuapi4go/pkg/breaker"
	"github.com/shing1211/futuapi4go/pkg/metrics"
	"github.com/shing1211/futuapi4go/pkg/ratelimit"
	"github.com/shing1211/futuapi4go/pkg/retry"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
)

var (
	loggerMu sync.RWMutex
	logger   = log.Default()
)

func SetLogger(l *log.Logger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	logger = l
}

func defaultLogger() *log.Logger {
	return log.Default()
}

func logf(format string, v ...interface{}) {
	loggerMu.RLock()
	l := logger
	loggerMu.RUnlock()
	if l == nil {
		return
	}
	l.Printf(format, v...)
}

// logInfo logs at info level if log level allows.
func (c *Client) logInfo(format string, v ...interface{}) {
	if c.opts.LogLevel > LogLevelInfo {
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
	if c.opts.LogLevel > LogLevelWarn {
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
	if c.opts.LogLevel > LogLevelError {
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
	ProtoID_GetGlobalState = 1002
	ProtoID_KeepAlive      = 1004
)

const (
	DefaultTimeout           = 30 * time.Second
	DefaultKeepAliveInterval = 30 * time.Second
	DefaultMaxRetries        = 3
	DefaultReconnectInterval = 3 * time.Second
	DefaultDialTimeout       = 10 * time.Second
)

// LogLevel constants for clarity.
// Higher values suppress more verbose logging.
// LogLevelInfo (0) = all logs, LogLevelSilent (3) = no logs.
const (
	LogLevelInfo   int = 0 // Log info, warnings, and errors
	LogLevelWarn   int = 1 // Log warnings and errors only
	LogLevelError  int = 2 // Log errors only
	LogLevelSilent int = 3 // Suppress all logs
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
	SlogLogger *SlogLogger // Structured logger (nil = use default)
	LogLevel int           // Log level: 0=Info, 1=Warn, 2=Error, 3=Silent. Use LogLevel* constants.

	// WebSocket
	WSSecretKey string // Secret key for WebSocket authentication

	// Push notifications
	PushHandler PacketHandler // Handler for incoming push notifications

	TLSConfig *tls.Config	// TLS configuration (nil = no TLS)

	// RSA
	RSAPublicKey  string // RSA public key in PEM format for InitConnect encryption
	RSAPrivateKey string // RSA private key in PEM format (required for encryption mode)

	// Encryption
	EncryptionEnabled bool // Enable FTAES encryption for all packets after InitConnect

	// Resilience
	RateLimiter  *ratelimit.ProtoLimiter // Rate limiter for API calls
	RetryConfig  retry.Config           // Retry configuration
	Breaker      *breaker.Breaker        // Circuit breaker
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

// WithLogLevel sets the log level.
// Use LogLevelInfo, LogLevelWarn, LogLevelError, or LogLevelSilent.
// Example: WithLogLevel(LogLevelWarn) suppresses info logs.
func WithLogLevel(level int) Option {
	return func(o *ClientOptions) { o.LogLevel = level }
}

// WithPushHandler sets a handler for push notifications.
func WithPushHandler(h PacketHandler) Option {
	return func(o *ClientOptions) { o.PushHandler = h }
}

// WithWSSecretKey sets the secret key for WebSocket authentication.
func WithWSSecretKey(key string) Option {
	return func(o *ClientOptions) { o.WSSecretKey = key }
}

// WithRSAPublicKey sets the RSA public key (PEM format) for encrypted InitConnect.
// When set, Connect() will use RSA encryption during the handshake.
// The PEM MUST be a PKIX/PKCS#8 format "PUBLIC KEY" — not a private key PEM.
// Passing a private key PEM will work (the public key is extracted), but logs a
// warning and is NOT recommended in production.
//
// Use this when connecting to a remote OpenD that has RSA encryption enabled.
//
// Example:
//
//	pubKey, err := os.ReadFile("/etc/futu/keys/opend_pubkey.pem")
//	cli := client.New(client.WithRSAPublicKey(string(pubKey)))
//	if err := cli.Connect("172.18.208.88:11111"); err != nil {
//	    log.Fatal(err)
//	}
func WithRSAPublicKey(pem string) Option {
	return func(o *ClientOptions) { o.RSAPublicKey = pem }
}

// WithRSAPrivateKey sets the RSA private key (PEM) for decrypting InitConnect responses.
// Required when WithEncryption is enabled. The public key is extracted automatically,
// so WithRSAPublicKey is not needed when using a private key PEM.
func WithRSAPrivateKey(pem string) Option {
	return func(o *ClientOptions) { o.RSAPrivateKey = pem }
}

// WithEncryption enables FTAES encryption for all packets after InitConnect.
// Requires WithRSAPrivateKey to be set for decrypting the InitConnect response.
// When enabled, the InitConnect response is RSA-decrypted to extract the AES key,
// and all subsequent communication uses FTAES_ECB encryption with that key.
// Matches the Python SDK's SysConfig.enable_proto_encrypt(True) behavior.
func WithEncryption(enable bool) Option {
	return func(o *ClientOptions) { o.EncryptionEnabled = enable }
}

// SetWSSecretKey sets the WebSocket secret key on an existing client.
func (c *Client) SetWSSecretKey(key string) {
	c.opts.WSSecretKey = key
}

// WithSlog sets a structured logger.
func WithSlog(sl *SlogLogger) Option {
	return func(o *ClientOptions) { o.SlogLogger = sl }
}

func WithTLS(cfg *tls.Config) Option {
	return func(o *ClientOptions) { o.TLSConfig = cfg }
}

func WithRateLimiter(rl *ratelimit.ProtoLimiter) Option {
	return func(o *ClientOptions) { o.RateLimiter = rl }
}

func WithRetryConfig(rc retry.Config) Option {
	return func(o *ClientOptions) { o.RetryConfig = rc }
}

func (c *Client) SetRateLimiter(rl *ratelimit.ProtoLimiter) {
	c.rateLimiter = rl
}

func (c *Client) SetRetryConfig(rc retry.Config) {
	c.retryConfig = &rc
}

// WithBreaker returns an Option that sets the circuit breaker for the client.
func WithBreaker(cb *breaker.Breaker) Option {
	return func(o *ClientOptions) { o.Breaker = cb }
}

// SetBreaker replaces the circuit breaker on the client.
func (c *Client) SetBreaker(cb *breaker.Breaker) {
	c.breaker = cb
}

// GetBreaker returns the current circuit breaker, or nil if none is set.
func (c *Client) GetBreaker() *breaker.Breaker {
	return c.breaker
}

type Client struct {
	conn              ConnInterface
	mu                sync.RWMutex
	opts              *ClientOptions
	connID            uint64
	loginUserID       uint64
	aesKey            string
	aesCBCIV          string
	isEncrypt         int32 // atomic: 1 = encrypted, 0 = not
	serverVer         int32
	keepAliveInterval int32
	serialNo          uint32
	serialMu          sync.Mutex
	handlers          map[uint32]Handler
	handlersMu        sync.RWMutex
	ctx               context.Context
	cancel            context.CancelFunc
	wg                sync.WaitGroup
	connected         int32
	connActive        int32

	addr         string
	reconnecting int32

	rsaKey string

	bufPool *bufferPool

	metrics   *Metrics
	metricsMu sync.RWMutex

	breaker      *breaker.Breaker
	rateLimiter  *ratelimit.ProtoLimiter
	retryConfig  *retry.Config
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

func (c *Client) ts() string {
	return time.Now().Format("15:04:05.000")
}

// GetMetrics returns a copy of current metrics.
func (c *Client) GetMetrics() Metrics {
	c.metricsMu.RLock()
	defer c.metricsMu.RUnlock()
	return *c.metrics
}

// GetReconnectCount returns the number of times the client has reconnected.
func (c *Client) GetReconnectCount() uint64 {
	c.metricsMu.RLock()
	defer c.metricsMu.RUnlock()
	return c.metrics.ReconnectCount
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
	protoStr := fmt.Sprintf("%d", protoID)
	status := "success"
	if err != nil {
		status = "error"
	}
	metrics.RecordAPICall(protoStr, status, duration)
}

func (c *Client) recordReconnect() {
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	c.metrics.ReconnectCount++
	metrics.RecordReconnect("connection_lost")
}

func (c *Client) recordPush() {
	c.metricsMu.Lock()
	defer c.metricsMu.Unlock()
	c.metrics.PushReceived++
	metrics.RecordPushMessage("default")
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
		bufPool:  newBufferPool(8192, 4096),
	}
	client.conn.SetAPITimeout(options.APITimeout)

	if options.RateLimiter != nil {
		client.rateLimiter = options.RateLimiter
	}
	if options.RetryConfig.MaxAttempts != 0 {
		client.retryConfig = &options.RetryConfig
	}
	if options.Breaker != nil {
		client.breaker = options.Breaker
	}

	return client
}

// NewWithOptions creates a Client with legacy parameters (deprecated, use New(With...) instead).
func NewWithOptions(addr string, maxRetries int, reconnectInterval time.Duration) *Client {
	return New(
		WithMaxRetries(maxRetries),
		WithReconnectInterval(reconnectInterval),
	)
}

// Connect connects to Futu OpenD via TCP (default).
// If RSAPublicKey was set via WithRSAPublicKey, RSA encryption is used during InitConnect.
func (c *Client) Connect(addr string) error {
	return c.ConnectWithRSA(addr, c.opts.RSAPublicKey)
}

// ConnectWS connects to Futu OpenD via WebSocket.
func (c *Client) ConnectWS(addr string) error {
	return c.connectWebSocket(addr, false)
}

// ConnectWSS connects to Futu OpenD via WebSocket Secure (TLS).
func (c *Client) ConnectWSS(addr string) error {
	return c.connectWebSocket(addr, true)
}

func (c *Client) connectWebSocket(addr string, tls bool) error {
	_, span := tracing.StartSpan(c.ctx, "futuapi.connect_ws",
		tracing.StringAttr("addr", addr),
		tracing.BoolAttr("tls", tls),
	)
	defer span.End()

	c.mu.Lock()
	c.addr = addr
	c.mu.Unlock()

	secretKey := c.opts.WSSecretKey

	var ws *websocket.Conn
	var err error

	if tls {
		ws, _, err = DialWebSocketSecure(context.Background(), addr, secretKey)
	} else {
		ws, _, err = DialWebSocket(context.Background(), addr, secretKey)
	}

	if err != nil {
		return fmt.Errorf("websocket dial: %w", err)
	}

	c.conn = newWSConn(ws)
	c.conn.SetAPITimeout(c.opts.APITimeout)

	clientVer := int32(1005)
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

	pkt := &initconnect.Request{
		C2S: c2s,
	}

	body, err := proto.Marshal(pkt)
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("marshal request: %w", err)
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(ProtoID_InitConnect, serialNo, body); err != nil {
		c.conn.Close()
		return fmt.Errorf("write request: %w", err)
	}

	rspPkt, err := c.conn.ReadResponse(serialNo, 10*time.Second)
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("read response: %w", err)
	}

	rsp := &initconnect.Response{}
	if err := proto.Unmarshal(rspPkt.Body, rsp); err != nil {
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
	c.loginUserID = s2c.GetLoginUserID()
	c.keepAliveInterval = s2c.GetKeepAliveInterval()
	atomic.StoreInt32(&c.isEncrypt, 0)
	c.serverVer = s2c.GetServerVer()
	atomic.StoreInt32(&c.connected, 1)
	atomic.StoreInt32(&c.connActive, 1)
	c.metricsMu.Lock()
	c.metrics.ConnectedSince = time.Now()
	c.metricsMu.Unlock()
	c.mu.Unlock()

	c.conn.SetPushHandler(func(pkt *Packet) {
		_, span := tracing.StartSpan(c.ctx, "futuapi.push",
			tracing.IntAttr("proto_id", int(pkt.Header.ProtoID)),
		)
		defer span.End()

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

	c.logInfo("connected (WS) to %s (connID=%d, userID=%d, ver=%d)", addr, c.connID, c.loginUserID, c.serverVer)

	keepAliveInterval := c.opts.KeepAliveInterval
	if keepAliveInterval == 0 {
		if c.keepAliveInterval > 0 {
			keepAliveInterval = time.Duration(c.keepAliveInterval) * time.Second
		} else {
			keepAliveInterval = DefaultKeepAliveInterval
		}
	}
	if keepAliveInterval > 0 {
		go c.keepAliveLoop(keepAliveInterval)
	}

	return nil
}

func (c *Client) ConnectWithRSA(addr string, rsaPublicKeyPEM string) error {
	_, span := tracing.StartSpan(c.ctx, "futuapi.connect",
		tracing.StringAttr("addr", addr),
	)
	defer span.End()

	c.mu.Lock()
	c.addr = addr
	c.rsaKey = rsaPublicKeyPEM
	c.mu.Unlock()

	c.logInfo("[%s] ConnectWithRSA: Start, addr=%s, rsaKeyLen=%d", c.ts(), addr, len(rsaPublicKeyPEM))

	// Determine encryption mode: use FTAES if EncryptionEnabled and private key is available
	useEncryption := c.opts.EncryptionEnabled && c.opts.RSAPrivateKey != ""
	if c.opts.EncryptionEnabled && c.opts.RSAPrivateKey == "" {
		return fmt.Errorf("encryption enabled but no RSA private key provided: use WithRSAPrivateKey()")
	}

	// Determine which PEM to use for RSA encrypt
	rsaEncryptPEM := rsaPublicKeyPEM
	if useEncryption && rsaEncryptPEM == "" {
		rsaEncryptPEM = c.opts.RSAPrivateKey // RSAEncrypt extracts public key from private key PEM
	}

	if c.opts.TLSConfig != nil {
		c.conn.SetTLSConfig(c.opts.TLSConfig)
	}

	c.logInfo("[%s] ConnectWithRSA: Dialing...", c.ts())
	dialStart := time.Now()
	if err := c.conn.Dial(addr); err != nil {
		c.logInfo("[%s] ConnectWithRSA: Dial FAILED: %v", c.ts(), err)
		return fmt.Errorf("dial: %w", err)
	}
	c.logInfo("[%s] ConnectWithRSA: Dial OK (%v)", c.ts(), time.Since(dialStart))

	clientVer := int32(1005)
	clientID := "futuapi4go"
	recvNotify := true
	programmingLanguage := "Go"

	var packetEncAlgo int32 = -1 // Default: no encryption
	if useEncryption {
		packetEncAlgo = 0 // FTAES_ECB
	}

	c2s := &initconnect.C2S{
		ClientVer:           &clientVer,
		ClientID:            &clientID,
		RecvNotify:          &recvNotify,
		PacketEncAlgo:       &packetEncAlgo,
		ProgrammingLanguage: &programmingLanguage,
		PushProtoFmt:        func() *int32 { v := int32(0); return &v }(),
	}

	pkt := &initconnect.Request{
		C2S: c2s,
	}

	c.logInfo("[%s] ConnectWithRSA: Marshaling InitConnect request...", c.ts())
	plainBody, err := proto.Marshal(pkt)
	if err != nil {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: Marshal FAILED: %v", c.ts(), err)
		return fmt.Errorf("marshal request: %w", err)
	}
	c.logInfo("[%s] ConnectWithRSA: Marshal OK, plainBody=%d bytes", c.ts(), len(plainBody))

	// Compute SHA1 over the SERIALIZED PLAINTEXT blob.
	// Server decrypts first, then verifies SHA1(plaintext) against this value.
	plainSHA1 := sha1.Sum(plainBody)

	body := plainBody

	// RSA encrypt the InitConnect body if a key is available
	if rsaEncryptPEM != "" {
		c.logInfo("[%s] ConnectWithRSA: RSA encrypting body...", c.ts())
		cryptoStart := time.Now()
		encryptedBody, err := RSAEncrypt(rsaEncryptPEM, plainBody)
		if err != nil {
			c.conn.Close()
			c.logInfo("[%s] ConnectWithRSA: RSA encrypt FAILED: %v", c.ts(), err)
			return fmt.Errorf("RSA encrypt: %w", err)
		}
		c.logInfo("[%s] ConnectWithRSA: RSA encrypt OK (%v): %d bytes -> %d bytes", c.ts(), time.Since(cryptoStart), len(plainBody), len(encryptedBody))
		body = encryptedBody
		c.logInfo("[%s] ConnectWithRSA: Using RSA encryption, final body=%d bytes", c.ts(), len(body))
	} else {
		c.logInfo("[%s] ConnectWithRSA: No encryption (packetEncAlgo=-1), body=%d bytes", c.ts(), len(body))
	}

	serialNo := c.nextSerialNo()
	c.logInfo("[%s] ConnectWithRSA: Writing packet (protoID=1001, serialNo=%d, body=%d bytes)...", c.ts(), serialNo, len(body))
	writeStart := time.Now()
	if err := c.conn.WritePacketWithSHA1(ProtoID_InitConnect, serialNo, body, plainSHA1); err != nil {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: WritePacket FAILED: %v", c.ts(), err)
		return fmt.Errorf("write packet: %w", err)
	}
	c.logInfo("[%s] ConnectWithRSA: WritePacket OK (%v)", c.ts(), time.Since(writeStart))

	c.logInfo("[%s] ConnectWithRSA: Starting readLoop goroutine...", c.ts())
	c.wg.Add(1)
	go c.readLoop()

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	c.logInfo("[%s] ConnectWithRSA: Waiting for response (serialNo=%d, timeout=%v)...", c.ts(), serialNo, apiTimeout)
	respStart := time.Now()
	respPkt, err := c.conn.ReadResponse(serialNo, apiTimeout)
	if err != nil {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: ReadResponse FAILED after %v: %v", c.ts(), time.Since(respStart), err)
		return fmt.Errorf("read response: %w", err)
	}
	c.logInfo("[%s] ConnectWithRSA: ReadResponse OK (%v), body=%d bytes, protoID=%d", c.ts(), time.Since(respStart), len(respPkt.Body), respPkt.Header.ProtoID)

	// Decrypt InitConnect response body if FTAES encryption is requested
	rspBody := respPkt.Body
	if useEncryption {
		c.logInfo("[%s] ConnectWithRSA: RSA decrypting InitConnect response body (%d bytes)...", c.ts(), len(rspBody))
		decryptedBody, err := RSADecrypt(c.opts.RSAPrivateKey, rspBody)
		if err != nil {
			c.conn.Close()
			c.logInfo("[%s] ConnectWithRSA: RSA decrypt FAILED: %v", c.ts(), err)
			return fmt.Errorf("RSA decrypt init connect response: %w", err)
		}
		rspBody = decryptedBody
		c.logInfo("[%s] ConnectWithRSA: RSA decrypt OK: %d bytes -> %d bytes", c.ts(), len(respPkt.Body), len(rspBody))
	}

	var rsp initconnect.Response
	if err := proto.Unmarshal(rspBody, &rsp); err != nil {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: Unmarshal response FAILED: %v", c.ts(), err)
		return fmt.Errorf("unmarshal response: %w", err)
	}
	c.logInfo("[%s] ConnectWithRSA: Response unmarshaled, retType=%d, retMsg=%s", c.ts(), rsp.GetRetType(), rsp.GetRetMsg())

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: Server returned error: retType=%d, retMsg=%s", c.ts(), rsp.GetRetType(), rsp.GetRetMsg())
		return fmt.Errorf("init connect failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		c.conn.Close()
		c.logInfo("[%s] ConnectWithRSA: S2C is nil!", c.ts())
		return errors.New("init connect: s2c is nil")
	}

	c.mu.Lock()
	c.connID = s2c.GetConnID()
	c.loginUserID = s2c.GetLoginUserID()
	c.aesKey = s2c.GetConnAESKey()
	c.aesCBCIV = s2c.GetAesCBCiv()
	isEnc := int32(0)
	if useEncryption {
		isEnc = 1
	}
	atomic.StoreInt32(&c.isEncrypt, isEnc)
	c.serverVer = s2c.GetServerVer()
	c.keepAliveInterval = s2c.GetKeepAliveInterval()
	atomic.StoreInt32(&c.connected, 1)
	atomic.StoreInt32(&c.connActive, 1)
	c.metricsMu.Lock()
	c.metrics.ConnectedSince = time.Now()
	c.metricsMu.Unlock()
	c.mu.Unlock()
	c.logInfo("[%s] ConnectWithRSA: Connected! connID=%d, serverVer=%d, loginUID=%d, encrypt=%v", c.ts(), c.connID, c.serverVer, c.loginUserID, useEncryption)

	metrics.RecordConnection("tcp")
	metrics.RecordOpenDUp(true)

	c.conn.SetPushHandler(func(pkt *Packet) {
		_, span := tracing.StartSpan(c.ctx, "futuapi.push",
			tracing.IntAttr("proto_id", int(pkt.Header.ProtoID)),
		)
		defer span.End()

		c.recordPush()
		protoStr := fmt.Sprintf("%d", pkt.Header.ProtoID)
		metrics.RecordPushMessage(protoStr)
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
	c.logInfo("[%s] readLoop: goroutine started", c.ts())

	for {
		select {
		case <-c.ctx.Done():
			c.drainPendingDispatches()
			c.logInfo("[%s] readLoop: context cancelled, exiting", c.ts())
			return
		default:
		}

		resultCh := make(chan *Packet, 1)
		errCh := make(chan error, 1)
		go func() {
			pkt, err := c.conn.readOne()
			if err != nil {
				errCh <- err
			} else {
				resultCh <- pkt
			}
		}()

		select {
		case <-c.ctx.Done():
			c.drainPendingDispatches()
			c.logInfo("[%s] readLoop: context cancelled, exiting", c.ts())
			return
		case pkt := <-resultCh:
			c.logInfo("[%s] readLoop: got packet protoID=%d serialNo=%d bodyLen=%d dispatching...", c.ts(), pkt.Header.ProtoID, pkt.Header.SerialNo, pkt.Header.BodyLen)
			c.conn.Dispatch(pkt)
		case err := <-errCh:
			c.mu.Lock()
			if atomic.LoadInt32(&c.connected) == 1 {
				atomic.StoreInt32(&c.connected, 0)
				c.logWarn("connection lost: %v\n", err)
				c.mu.Unlock()
				go c.reconnect()
			} else {
				c.mu.Unlock()
			}
			return
		}
	}
}

func (c *Client) drainPendingDispatches() {
	c.conn.DrainDispatches()
}

func (c *Client) reconnect() {
	_, span := tracing.StartSpan(c.ctx, "futuapi.reconnect")
	defer span.End()

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

	if c.conn != nil {
		c.conn.Close()
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

func (c *Client) UnregisterHandler(protoID uint32) {
	c.handlersMu.Lock()
	delete(c.handlers, protoID)
	c.handlersMu.Unlock()
}

func (c *Client) Close() error {
	_, span := tracing.StartSpan(c.ctx, "futuapi.close")
	defer span.End()

	atomic.StoreInt32(&c.connActive, 0)
	atomic.StoreInt32(&c.connected, 0)
	metrics.RecordDisconnect("tcp")
	metrics.RecordOpenDUp(false)
	if c.conn != nil {
		c.conn.DrainDispatches()
		c.conn.Close()
	}
	c.cancel()
	c.wg.Wait()
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

// GetLoginUserID returns the Futu/NiuNiu user ID that logged into OpenD.
func (c *Client) GetLoginUserID() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.loginUserID
}

// IsEncrypt returns true if the connection uses AES encryption.
func (c *Client) IsEncrypt() bool {
	return atomic.LoadInt32(&c.isEncrypt) != 0
}

func (c *Client) getAESKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.aesKey
}

// EncryptRequestBody encrypts the request body using FTAES_ECB when the
// connection is in encrypted mode. Returns the encrypted body (or original if
// not encrypted). InitConnect (protoID 1001) is NEVER encrypted — RSA handles
// that in the handshake itself.
func (c *Client) EncryptRequestBody(protoID uint32, body []byte) ([]byte, error) {
	if atomic.LoadInt32(&c.isEncrypt) == 0 || protoID == ProtoID_InitConnect {
		return body, nil
	}
	if c.aesCBCIV != "" {
		return aesCBCEncrypt([]byte(c.getAESKey()), []byte(c.aesCBCIV), body)
	}
	key256 := sha256.Sum256([]byte(c.getAESKey()))
	return aes256Encrypt(key256[:], body)
}

// DecryptResponseBody decrypts the response body using FTAES_ECB when the
// connection is in encrypted mode. Returns the plaintext (or original if not
// encrypted). InitConnect responses are handled separately via RSA.
// Tries multiple AES modes as fallback for compatibility with different OpenD configurations.
func (c *Client) DecryptResponseBody(protoID uint32, body []byte) ([]byte, error) {
	if atomic.LoadInt32(&c.isEncrypt) == 0 || protoID == ProtoID_InitConnect {
		return body, nil
	}

	plaintext, err := ftaesDecrypt([]byte(c.getAESKey()), body)
	if err == nil {
		return plaintext, nil
	}
	if errors.Is(err, ErrNotEncrypted) {
		return body, nil
	}

	if c.aesCBCIV != "" && len(body) > 16 && len(body)%16 == 0 {
		plaintext, err = aesCBCDecrypt([]byte(c.getAESKey()), []byte(c.aesCBCIV), body)
		if err == nil {
			return plaintext, nil
		}
	}

	key256 := sha256.Sum256([]byte(c.getAESKey()))
	plaintext, err = aes256Decrypt(key256[:], body)
	if err == nil {
		return plaintext, nil
	}

	return plaintext, err
}

// CanSendProto reports whether a request for the given proto ID can be sent,
// based on the current connection state. InitConnect can be sent when connected.
// All other protos require the connection to be fully ready (connected + handshaken).
func (c *Client) CanSendProto(protoID uint32) bool {
	if !c.IsConnected() {
		return false
	}
	if protoID == ProtoID_InitConnect {
		return true
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.serverVer > 0
}

func (c *Client) IsConnected() bool {
	return atomic.LoadInt32(&c.connected) == 1
}

// EnsureConnected returns an error if the client is not connected.
// This should be called by all public API functions before making requests.
func (c *Client) EnsureConnected() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if atomic.LoadInt32(&c.connected) == 0 {
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
// Note: the returned Client shares the handler map with the original — handler registrations
// on the returned Client will also affect the original. For isolated handler registries,
// use client.New(...) with separate connections instead.
func (c *Client) WithContext(ctx context.Context) *Client {
	handlers := make(map[uint32]Handler, len(c.handlers))
	for k, v := range c.handlers {
		handlers[k] = v
	}
	newClient := &Client{
		conn:     c.conn,
		opts:     c.opts,
		handlers: handlers,
		ctx:      ctx,
		cancel:   func() {}, // Don't cancel parent context
	}
	// Deep-copy options so caller mutations don't affect the original
	optsCopy := *c.opts
	newClient.opts = &optsCopy
	newClient.mu.RLock()
	newClient.connID = c.connID
	newClient.loginUserID = c.loginUserID
	newClient.aesKey = c.aesKey
	newClient.isEncrypt = atomic.LoadInt32(&c.isEncrypt)
	newClient.serverVer = c.serverVer
	newClient.keepAliveInterval = c.keepAliveInterval
	newClient.connected = atomic.LoadInt32(&c.connected)
	newClient.mu.RUnlock()
	return newClient
}

func (c *Client) Conn() ConnInterface {
	return c.conn
}

func (c *Client) NextSerialNo() uint32 {
	return c.nextSerialNo()
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
	serialNo := c.NextSerialNo()
	if err := c.conn.WritePacket(protoID, serialNo, body); err != nil {
		return fmt.Errorf("write packet: %w", err)
	}
	pktResp, err := c.conn.ReadResponse(serialNo, c.APITimeout())
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
	var err error
	if c.breaker != nil && !isControlProto(protoID) {
		_, err = c.breaker.Do(func() (interface{}, error) {
			return nil, c.requestInternal(protoID, req, rsp)
		})
	} else {
		err = c.requestInternal(protoID, req, rsp)
	}
	c.recordRequest(protoID, time.Since(start), err)
	return err
}

func (c *Client) RequestContext(ctx context.Context, protoID uint32, req proto.Message, rsp proto.Message) error {
	start := time.Now()
	var err error
	if c.breaker != nil && !isControlProto(protoID) {
		_, err = c.breaker.Do(func() (interface{}, error) {
			return nil, c.requestContextInternal(ctx, protoID, req, rsp)
		})
	} else {
		err = c.requestContextInternal(ctx, protoID, req, rsp)
	}
	c.recordRequest(protoID, time.Since(start), err)
	return err
}

func isControlProto(protoID uint32) bool {
	return protoID == ProtoID_InitConnect || protoID == ProtoID_KeepAlive || protoID == ProtoID_GetGlobalState
}

func (c *Client) requestInternal(protoID uint32, req proto.Message, rsp proto.Message) error {
	if c.conn == nil {
		return ErrNotConnected
	}

	body, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	// Encrypt body if connection is encrypted (skip InitConnect — RSA handles that)
	// SHA1 is computed over ciphertext for historical compatibility (v0.5.15).
	// OpenD accepts both SHA1(plaintext) and SHA1(ciphertext) — see WritePacketEncrypted doc.
	serialNo := c.nextSerialNo()
	if atomic.LoadInt32(&c.isEncrypt) != 0 && protoID != ProtoID_InitConnect {
		encBody, err := c.EncryptRequestBody(protoID, body)
		if err != nil {
			return fmt.Errorf("AES encrypt: %w", err)
		}
		encSHA1 := sha1.Sum(encBody)
		if err := c.conn.WritePacketEncrypted(protoID, serialNo, encBody, encSHA1); err != nil {
			return fmt.Errorf("write packet: %w", err)
		}
	} else {
		if err := c.conn.WritePacket(protoID, serialNo, body); err != nil {
			return fmt.Errorf("write packet: %w", err)
		}
	}

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}
	pkt, err := c.conn.ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// Decrypt body if encrypted (skip InitConnect response — RSA handles that)
	plaintext := pkt.Body
	if atomic.LoadInt32(&c.isEncrypt) != 0 && protoID != ProtoID_InitConnect {
		plaintext, err = c.DecryptResponseBody(protoID, pkt.Body)
		if err != nil {
			return fmt.Errorf("AES decrypt: %w", err)
		}
	}

	if err := proto.Unmarshal(plaintext, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (c *Client) requestContextInternal(ctx context.Context, protoID uint32, req proto.Message, rsp proto.Message) error {
	ctx, span := tracing.StartSpan(ctx, "futuapi.request",
		tracing.IntAttr("proto_id", int(protoID)),
	)
	defer span.End()

	if c.conn == nil {
		return ErrNotConnected
	}

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	// Encrypt body if connection is encrypted (skip InitConnect — RSA handles that)
	// SHA1 is computed over ciphertext for historical compatibility (v0.5.15).
	// OpenD accepts both SHA1(plaintext) and SHA1(ciphertext) — see WritePacketEncrypted doc.
	serialNo := c.nextSerialNo()
	if atomic.LoadInt32(&c.isEncrypt) != 0 && protoID != ProtoID_InitConnect {
		encBody, err := c.EncryptRequestBody(protoID, body)
		if err != nil {
			return fmt.Errorf("AES encrypt: %w", err)
		}
		encSHA1 := sha1.Sum(encBody)
		if err := c.conn.WritePacketEncrypted(protoID, serialNo, encBody, encSHA1); err != nil {
			return fmt.Errorf("write packet: %w", err)
		}
	} else {
		if err := c.conn.WritePacket(protoID, serialNo, body); err != nil {
			return err
		}
	}

	apiTimeout := c.opts.APITimeout
	if apiTimeout == 0 {
		apiTimeout = DefaultTimeout
	}

	deadline, hasDeadline := ctx.Deadline()
	if hasDeadline {
		if timeout := time.Until(deadline); timeout < apiTimeout {
			apiTimeout = timeout
		}
	}

	pkt, err := c.conn.ReadResponseContext(ctx, serialNo, apiTimeout)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// Decrypt body if encrypted (skip InitConnect response — RSA handles that)
	plaintext := pkt.Body
	if atomic.LoadInt32(&c.isEncrypt) != 0 && protoID != ProtoID_InitConnect {
		plaintext, err = c.DecryptResponseBody(protoID, pkt.Body)
		if err != nil {
			return fmt.Errorf("AES decrypt: %w", err)
		}
	}

	if err := proto.Unmarshal(plaintext, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
