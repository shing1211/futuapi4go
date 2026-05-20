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

// Package client provides a public Client type for the Futu OpenD SDK.
// This allows external projects to use the SDK.
package client

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

// ConnState represents the connection state of the client.
type ConnState = futuapi.ConnState

const (
	StateDisconnected ConnState = 0
	StateConnecting   ConnState = 1
	StateConnected    ConnState = 2
	StateReconnecting ConnState = 3
	StateClosing      ConnState = 4
)

// Client is the main client type for connecting to Futu OpenD.
// It wraps the internal client to provide a public API.
type Client struct {
	inner  *futuapi.Client
	trdEnv constant.TrdEnv
	trdMkt constant.TrdMarket
	cbs    *callbackState
}

func New(opts ...Option) *Client {
	futuOpts := make([]futuapi.Option, len(opts))
	for i, o := range opts {
		futuOpts[i] = o
	}
	return &Client{
		inner:  futuapi.New(futuOpts...),
		trdEnv: constant.TrdEnv_Simulate,
		cbs:    &callbackState{},
	}
}

func (c *Client) GetTradeEnv() constant.TrdEnv {
	return c.trdEnv
}

// WithTradeEnv returns a new Client with the specified trading environment.
// Note: The returned client shares the same underlying connection as the original.
// Changing trade environment on one client affects both.
func (c *Client) WithTradeEnv(trdEnv constant.TrdEnv) *Client {
	clone := *c
	clone.trdEnv = trdEnv
	return &clone
}

// WithTradeMarket returns a new Client with the specified trading market.
// Note: The returned client shares the same underlying connection as the original.
func (c *Client) WithTradeMarket(market constant.TrdMarket) *Client {
	clone := *c
	clone.trdMkt = market
	return &clone
}

func (c *Client) GetTradeMarket() constant.TrdMarket {
	return c.trdMkt
}

// FindAccount returns the first account matching the client's trdEnv.
func (c *Client) FindAccount(accounts []Account) *Account {
	for _, acc := range accounts {
		if acc.TrdEnv == int32(c.trdEnv) {
			return &acc
		}
	}
	if len(accounts) > 0 {
		return &accounts[0]
	}
	return nil
}

func (c *Client) Connect(addr string) error {
	return c.inner.Connect(addr)
}

// ConnectWS connects to the Futu OpenD server via WebSocket.
func (c *Client) ConnectWS(addr string, secretKey ...string) error {
	if len(secretKey) > 0 && secretKey[0] != "" {
		c.inner.SetWSSecretKey(secretKey[0])
	}
	return c.inner.ConnectWS(addr)
}

// ConnectWSS connects to the Futu OpenD server via WebSocket Secure (TLS).
func (c *Client) ConnectWSS(addr string, secretKey ...string) error {
	if len(secretKey) > 0 && secretKey[0] != "" {
		c.inner.SetWSSecretKey(secretKey[0])
	}
	return c.inner.ConnectWSS(addr)
}

// ConnectAddr is an alias for Connect.
func (c *Client) ConnectAddr(addr string) error {
	return c.inner.Connect(addr)
}

// Close closes the connection to OpenD.
func (c *Client) Close() {
	c.inner.Close()
}

// Shutdown gracefully drains pending requests then closes the connection.
func (c *Client) Shutdown(timeout time.Duration) error {
	return c.inner.Shutdown(timeout)
}

// IsConnected returns true if the client is currently connected.
func (c *Client) IsConnected() bool {
	return c.inner.IsConnected()
}

// State returns the current connection state.
func (c *Client) State() futuapi.ConnState {
	return c.inner.State()
}

// WaitForSignal blocks until a termination signal is received.
func (c *Client) WaitForSignal(cleanup func()) os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	if cleanup != nil {
		cleanup()
	}
	return sig
}

// CloseOnSignal registers cleanup to be called when the process receives
// termination signals (SIGINT, SIGTERM). Returns a function to unregister
// the handler. The client is automatically closed when the signal is received.
func (c *Client) CloseOnSignal() (unregister func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		c.Close()
	}()

	return func() {
		signal.Stop(sigChan)
	}
}

// GetConnID returns the connection ID assigned by OpenD.
func (c *Client) GetConnID() uint64 {
	return c.inner.GetConnID()
}

// GetServerVer returns the OpenD server version.
func (c *Client) GetServerVer() int32 {
	return c.inner.GetServerVer()
}

// GetLoginUserID returns the Futu/NiuNiu user ID that logged into OpenD.
func (c *Client) GetLoginUserID() uint64 {
	return c.inner.GetLoginUserID()
}

// IsEncrypt returns true if the connection uses AES encryption.
func (c *Client) IsEncrypt() bool {
	return c.inner.IsEncrypt()
}

// CanSendProto reports whether a request for the given proto ID can be sent.
func (c *Client) CanSendProto(protoID uint32) bool {
	return c.inner.CanSendProto(protoID)
}

// EnsureConnected returns an error if the client is not connected.
func (c *Client) EnsureConnected() error {
	return c.inner.EnsureConnected()
}

// Inner returns the underlying internal client (for advanced use).
func (c *Client) Inner() *futuapi.Client {
	return c.inner
}

// WithContext returns a client with the given context.
func (c *Client) WithContext(ctx context.Context) *Client {
	return &Client{inner: c.inner.WithContext(ctx)}
}

// Context returns the client's context.
func (c *Client) Context() context.Context {
	return c.inner.Context()
}

// WithTimeout returns a context with the specified timeout.
func (c *Client) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.inner.Context(), timeout)
}

// WithDeadline returns a context with the specified deadline.
func (c *Client) WithDeadline(deadline time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(c.inner.Context(), deadline)
}

// RegisterHandler registers a handler for push notifications.
func (c *Client) RegisterHandler(protoID uint32, h func(protoID uint32, body []byte)) {
	c.inner.RegisterHandler(protoID, h)
}

// UnregisterHandler unregisters a previously registered push handler.
func (c *Client) UnregisterHandler(protoID uint32) {
	c.inner.UnregisterHandler(protoID)
}

// SetPushHandler registers a handler that receives push notifications for
// specific protoIDs. The handler receives (protoID, rawBody) and should use
// the ParsePush* functions below to decode the body.
func (c *Client) SetPushHandler(protoID uint32, h func(protoID uint32, body []byte)) {
	c.inner.RegisterHandler(protoID, h)
}

// GetConn returns the underlying connection interface (for advanced use).
func (c *Client) GetConn() futuapi.ConnInterface {
	return c.inner.Conn()
}

// Quote returns a QuoteAPI for market data operations.
func (c *Client) Quote() *QuoteAPI {
	return &QuoteAPI{client: c.inner}
}

// Trade returns a TradeAPI for trading operations.
func (c *Client) Trade() *TradeAPI {
	return &TradeAPI{client: c.inner, trdEnv: c.trdEnv, trdMkt: c.trdMkt}
}

// System returns a SystemAPI for system operations.
func (c *Client) System() *SystemAPI {
	return &SystemAPI{client: c.inner}
}

// inferSecMarket extracts the market from a market-prefixed stock code.
func inferSecMarket(code string) int32 {
	if len(code) < 3 {
		return 0
	}
	if len(code) > 3 && code[:3] == "HK." {
		return 1
	}
	if len(code) > 3 && code[:3] == "US." {
		return 2
	}
	if len(code) > 3 && code[:3] == "SH." {
		return 3
	}
	if len(code) > 3 && code[:3] == "SZ." {
		return 3
	}
	if len(code) > 3 && code[len(code)-3:] == ".HK" {
		return 1
	}
	if len(code) > 3 && code[len(code)-3:] == ".US" {
		return 2
	}
	if (len(code) > 3 && code[len(code)-3:] == ".SH") || (len(code) > 3 && code[len(code)-3:] == ".SZ") {
		return 3
	}
	return 0
}

func getStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func getBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func getFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func getInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func getUint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

// ============================================================================
// Options (aliases for internal client options)
type Option = futuapi.Option

// WithDialTimeout sets the connection dial timeout.
func WithDialTimeout(d time.Duration) Option {
	return futuapi.WithDialTimeout(d)
}

// WithAPISetTimeout sets the API request timeout.
func WithAPISetTimeout(d time.Duration) Option {
	return futuapi.WithAPITimeout(d)
}

// WithKeepAliveInterval sets the keep-alive interval.
func WithKeepAliveInterval(d time.Duration) Option {
	return futuapi.WithKeepAliveInterval(d)
}

// WithReconnectInterval sets the base reconnect interval when the connection drops.
func WithReconnectInterval(d time.Duration) Option {
	return futuapi.WithReconnectInterval(d)
}

// WithMaxRetries sets the maximum retry attempts.
func WithMaxRetries(n int) Option {
	return futuapi.WithMaxRetries(n)
}

// WithLogLevel sets the logging level (0=info, 1=warn, 2=error, 3=silent).
func WithLogLevel(level int) Option {
	return futuapi.WithLogLevel(level)
}

// WithRSAPublicKey sets the RSA public key (PEM format) for encrypted InitConnect.
func WithRSAPublicKey(pem string) Option {
	return futuapi.WithRSAPublicKey(pem)
}

// WithRSAPrivateKey sets the RSA private key (PEM) for decrypting InitConnect responses.
func WithRSAPrivateKey(pem string) Option {
	return futuapi.WithRSAPrivateKey(pem)
}

// WithEncryption enables FTAES encryption for all packets after InitConnect.
func WithEncryption(enable bool) Option {
	return futuapi.WithEncryption(enable)
}

// WithOnStateChange sets a callback that is invoked when the connection state changes.
func WithOnStateChange(fn func(oldState, newState futuapi.ConnState)) Option {
	return func(o *futuapi.ClientOptions) { o.OnStateChange = fn }
}

// Env vars read by WithEnvConfig:
//
//	FUTU_OPEND_ADDR          — OpenD address (host:port)
//	FUTU_RSA_PUBLIC_KEY      — RSA public key PEM (file path or inline)
//	FUTU_RSA_PRIVATE_KEY     — RSA private key PEM (file path or inline)
//	FUTU_ENCRYPT             — "1" or "true" to enable encryption
//	FUTU_LOG_LEVEL           — log level (0=info, 1=warn, 2=error, 3=silent)
//	FUTU_TRD_ENV             — "real" or "simulate"
//
// WithEnvConfig returns an Option that reads these variables at call time.
// PEM values are treated as file paths first; if the file does not exist,
// the value is used as the raw PEM string.
func WithEnvConfig() Option {
	return func(o *futuapi.ClientOptions) {
		if v := os.Getenv("FUTU_RSA_PUBLIC_KEY"); v != "" {
			if info, err := os.Stat(v); err == nil && !info.IsDir() {
				data, err := os.ReadFile(v)
				if err != nil {
					slog.Warn("FUTU_RSA_PUBLIC_KEY: file exists but cannot be read", "path", v, "error", err)
					o.RSAPublicKey = v
				} else {
					o.RSAPublicKey = string(data)
				}
			} else {
				o.RSAPublicKey = v
			}
		}
		if v := os.Getenv("FUTU_RSA_PRIVATE_KEY"); v != "" {
			if info, err := os.Stat(v); err == nil && !info.IsDir() {
				data, err := os.ReadFile(v)
				if err != nil {
					slog.Warn("FUTU_RSA_PRIVATE_KEY: file exists but cannot be read", "path", v, "error", err)
					o.RSAPrivateKey = v
				} else {
					o.RSAPrivateKey = string(data)
				}
			} else {
				o.RSAPrivateKey = v
			}
		}
		if v := os.Getenv("FUTU_ENCRYPT"); v == "1" || v == "true" {
			o.EncryptionEnabled = true
		}
		if v := os.Getenv("FUTU_LOG_LEVEL"); v != "" {
			if lvl, err := parseInt(v); err == nil {
				o.LogLevel = lvl
			}
		}
	}
}

func parseInt(s string) (int, error) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("not a number: %s", s)
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}

// Default timeouts and connection limits.
const (
	DefaultDialTimeout      = 10 * time.Second
	DefaultAPITimeout       = 30 * time.Second
	DefaultKeepAlive        = 30 * time.Second
	DefaultMaxRetries       = 3
	DefaultReconnectBackoff = 1.5
)

// Trading environment constants.
const (
	TrdEnv_Real     = int32(1)
	TrdEnv_Simulate = int32(0)
)

// PriceReminderOp constants for price reminder operations.
const (
	PriceReminderOpAdd    = 1
	PriceReminderOpUpdate = 2
	PriceReminderOpDelete = 3
)

// Common market constants.
const (
	Market_HK_Security   = int32(qotcommon.QotMarket_QotMarket_HK_Security)
	Market_HK_Future     = int32(qotcommon.QotMarket_QotMarket_HK_Future)
	Market_US_Security   = int32(qotcommon.QotMarket_QotMarket_US_Security)
	Market_CNSH_Security = int32(qotcommon.QotMarket_QotMarket_CNSH_Security)
	Market_CNSZ_Security = int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)

	Side_Buy  = int32(trdcommon.TrdSide_TrdSide_Buy)
	Side_Sell = int32(trdcommon.TrdSide_TrdSide_Sell)

	OrderType_Normal = int32(trdcommon.OrderType_OrderType_Normal)
	OrderType_Market = int32(trdcommon.OrderType_OrderType_Market)
	OrderType_Stop   = int32(trdcommon.OrderType_OrderType_Stop)

	KLType_Day   = int32(qotcommon.KLType_KLType_Day)
	KLType_1Min  = int32(qotcommon.KLType_KLType_1Min)
	KLType_5Min  = int32(qotcommon.KLType_KLType_5Min)
	KLType_15Min = int32(qotcommon.KLType_KLType_15Min)
	KLType_30Min = int32(qotcommon.KLType_KLType_30Min)
	KLType_60Min = int32(qotcommon.KLType_KLType_60Min)
	KLType_Week  = int32(qotcommon.KLType_KLType_Week)
	KLType_Month = int32(qotcommon.KLType_KLType_Month)

	SubType_Basic     = int32(qot.SubType_Basic)
	SubType_OrderBook = int32(qot.SubType_OrderBook)
	SubType_Ticker    = int32(qot.SubType_Ticker)
	SubType_RT        = int32(qot.SubType_RT)
	SubType_KL        = int32(qot.SubType_KL)
	SubType_KL_1Min   = int32(qot.SubType_KL_1Min)
	SubType_KL_5Min   = int32(qot.SubType_KL_5Min)
	SubType_KL_15Min  = int32(qot.SubType_KL_15Min)
	SubType_KL_30Min  = int32(qot.SubType_KL_30Min)
	SubType_KL_60Min  = int32(qot.SubType_KL_60Min)
	SubType_KL_Day    = int32(qot.SubType_KL_Day)
	SubType_KL_Week   = int32(qot.SubType_KL_Week)
	SubType_KL_Month  = int32(qot.SubType_KL_Month)
	SubType_Broker    = int32(qot.SubType_Broker)
)
