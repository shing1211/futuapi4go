// Package futuapi provides a convenience re-export of the client package
// for quickstart usage. Users who need fine-grained control should
// import github.com/shing1211/futuapi4go/client directly.
//
// Quickstart:
//
//	import futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
//
//	cli, err := futuapi.NewClient("127.0.0.1:11111")
//	if err != nil { log.Fatal(err) }
//	defer cli.Close()
package futuapi

import (
	"os"
	"time"

	"github.com/shing1211/futuapi4go/client"
)

// NewClient creates a new Client, connects to addr, and returns it.
// Equivalent to client.New() + client.Connect(addr) in one call.
func NewClient(addr string, opts ...client.Option) (*client.Client, error) {
	cli := client.New(opts...)
	if err := cli.Connect(addr); err != nil {
		return nil, err
	}
	return cli, nil
}

// NewClientFromEnv creates a Client configured from environment variables
// and connects using the address from FUTU_OPEND_ADDR (or 127.0.0.1:11111).
// See client.WithEnvConfig() for the supported variables.
func NewClientFromEnv() (*client.Client, error) {
	addr := os.Getenv("FUTU_OPEND_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}
	return NewClient(addr, client.WithEnvConfig())
}

// Option is an alias for client.Option, used with NewClient.
type Option = client.Option

// Convenience functions.

// WithEnvConfig configures the client from environment variables.
// See client.WithEnvConfig for details.
func WithEnvConfig() Option {
	return client.WithEnvConfig()
}

// WithDialTimeout sets the connection dial timeout.
func WithDialTimeout(d time.Duration) Option {
	return client.WithDialTimeout(d)
}

// WithLogLevel sets the logging level.
func WithLogLevel(level int) Option {
	return client.WithLogLevel(level)
}

// WithRSAPublicKey sets the RSA public key for encrypted InitConnect.
func WithRSAPublicKey(pem string) Option {
	return client.WithRSAPublicKey(pem)
}

// Convenience constants — re-exported from the client package.

// ProtoIDs for push notifications.
const (
	ProtoID_Qot_UpdateBasicQot      = client.ProtoID_Qot_UpdateBasicQot
	ProtoID_Qot_UpdateKL            = client.ProtoID_Qot_UpdateKL
	ProtoID_Qot_UpdateOrderBook     = client.ProtoID_Qot_UpdateOrderBook
	ProtoID_Qot_UpdateTicker        = client.ProtoID_Qot_UpdateTicker
	ProtoID_Qot_UpdateRT            = client.ProtoID_Qot_UpdateRT
	ProtoID_Qot_UpdateBroker        = client.ProtoID_Qot_UpdateBroker
	ProtoID_Qot_UpdatePriceReminder = client.ProtoID_Qot_UpdatePriceReminder
	ProtoID_Trd_UpdateOrder         = client.ProtoID_Trd_UpdateOrder
	ProtoID_Trd_UpdateOrderFill     = client.ProtoID_Trd_UpdateOrderFill
	ProtoID_Trd_Notify              = client.ProtoID_Trd_Notify
)

// Trading environments.
const (
	TrdEnv_Real     = client.TrdEnv_Real
	TrdEnv_Simulate = client.TrdEnv_Simulate
)

// Order sides.
const (
	Side_Buy  = client.Side_Buy
	Side_Sell = client.Side_Sell
)

// Order types.
const (
	OrderType_Normal = client.OrderType_Normal
	OrderType_Market = client.OrderType_Market
	OrderType_Stop   = client.OrderType_Stop
)

// K-line types.
const (
	KLType_Day   = client.KLType_Day
	KLType_1Min  = client.KLType_1Min
	KLType_5Min  = client.KLType_5Min
	KLType_15Min = client.KLType_15Min
	KLType_30Min = client.KLType_30Min
	KLType_60Min = client.KLType_60Min
	KLType_Week  = client.KLType_Week
	KLType_Month = client.KLType_Month
)

// Markets.
const (
	Market_HK_Security   = client.Market_HK_Security
	Market_HK_Future     = client.Market_HK_Future
	Market_US_Security   = client.Market_US_Security
	Market_CNSH_Security = client.Market_CNSH_Security
	Market_CNSZ_Security = client.Market_CNSZ_Security
)

// Subscription types.
const (
	SubType_Basic     = client.SubType_Basic
	SubType_OrderBook = client.SubType_OrderBook
	SubType_Ticker    = client.SubType_Ticker
	SubType_RT        = client.SubType_RT
	SubType_KL        = client.SubType_KL
	SubType_KL_1Min   = client.SubType_KL_1Min
	SubType_KL_5Min   = client.SubType_KL_5Min
	SubType_KL_15Min  = client.SubType_KL_15Min
	SubType_KL_30Min  = client.SubType_KL_30Min
	SubType_KL_60Min  = client.SubType_KL_60Min
	SubType_KL_Day    = client.SubType_KL_Day
	SubType_KL_Week   = client.SubType_KL_Week
	SubType_KL_Month  = client.SubType_KL_Month
	SubType_Broker    = client.SubType_Broker
)
