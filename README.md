# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.2.0-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.4.6408-blue?style=flat-square" alt="Futu Proto Version">
</p>

> **Go-native. Type-safe. Production-ready.** The most complete and ergonomic Go SDK for Futu OpenAPI — market data, trading, real-time push, and more.

## Why futuapi4go?

- **Compile-time safety** — structs over DataFrames, no runtime surprises
- **Go concurrency** — goroutines, channels, context cancellation baked in
- **No Python dependency** — pure Go, deploy anywhere with `go build`
- **More protos** — 78 protos vs Python's ~50, including futures, flow summaries, and all push types
- **Modern infrastructure** — circuit breaker, structured logging, channel-based push delivery, connection pool with health checks

## Install

```bash
go get github.com/shing1211/futuapi4go@v0.2.0
```

## Your First Trade

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/push"
	chanpkg "github.com/shing1211/futuapi4go/pkg/push/chan"
)

func main() {
	cli := client.New()
	defer cli.Close()

	if err := cli.Connect("127.0.0.1:11111"); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}

	// Note: US stocks require subscription before GetQuote works
	// Get a one-shot quote
	quote, err := client.GetQuote(context.Background(), cli, constant.Market_US, "NVDA")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get quote: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("US.NVDA: price=%.2f open=%.2f high=%.2f low=%.2f vol=%d\n",
		quote.Price, quote.Open, quote.High, quote.Low, quote.Volume)

	// Set up channel listeners for real-time data
	quoteCh := make(chan *push.UpdateBasicQot, 100)
	stopQuote := chanpkg.SubscribeQuote(cli, constant.Market_US, "NVDA", quoteCh)
	defer stopQuote()

	// Set up multiple K-line handlers
	klHandlers := map[constant.KLType]func(*push.UpdateKL){
		constant.KLType_K_1Min: func(kl *push.UpdateKL) {
			for _, bar := range kl.KLList {
				fmt.Printf("1MIN KL: %s C=%.2f V=%d\n",
					*bar.Time, *bar.ClosePrice, *bar.Volume)
			}
		},
		constant.KLType_K_Day: func(kl *push.UpdateKL) {
			for _, bar := range kl.KLList {
				fmt.Printf("DAY KL: %s O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
					*bar.Time, *bar.OpenPrice, *bar.HighPrice,
					*bar.LowPrice, *bar.ClosePrice, *bar.Volume)
			}
		},
	}
	stopKLines := chanpkg.SubscribeKLines(cli, constant.Market_US, "NVDA", klHandlers)
	defer stopKLines()

	// Graceful shutdown on Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case q := <-quoteCh:
			fmt.Printf("QUOTE [%s]: price=%.2f vol=%d\n",
				q.Security.GetCode(), q.CurPrice, q.Volume)
		case <-sig:
			fmt.Println("Shutting down...")
			return
		}
	}
}
```

> **Note:** US stocks require subscribing before `GetQuote` works. HK stocks don't need subscription.

## Package Map

| Package | What it's For |
|---------|--------------|
| `client` | High-level wrappers — the recommended entry point |
| `internal/client` | TCP connection, packet I/O, auto-reconnect, keep-alive |
| `pkg/qot` | All market data APIs (quotes, K-lines, order book, tick data...) |
| `pkg/trd` | All trading APIs (orders, positions, funds, history...) |
| `pkg/sys` | System APIs (global state, user info) |
| `pkg/push` | Parse push notification payloads |
| `pkg/push/chan` | Subscribe to real-time data via Go channels |
| `pkg/breaker` | Circuit breaker — protect trading from cascading failures |
| `pkg/logger` | Structured logging, text + JSON, leveled (Debug/Info/Warn/Error) |
| `pkg/util` | Code parsing (`HK.00700` ↔ market+code), market helpers |
| `pkg/constant` | Python-style constants (`Market_HK`, `TrdSide_Buy`, `KLType_K_Day`...) |
| `pkg/pb/*` | 78 protobuf-generated types for all Futu OpenAPI messages |

## Key Features in Depth

### Channel-Based Real-Time Push

Stop polling. Let data come to you:

#### Single data type (channel-based):

```go
import (
	chanpkg "github.com/shing1211/futuapi4go/pkg/push/chan"
)

// Quote updates stream into the channel
ch := make(chan *push.UpdateBasicQot, 100)
stop := chanpkg.SubscribeQuote(cli, constant.Market_HK, "00700", ch)
defer stop()

for q := range ch {
	fmt.Printf("QUOTE [%s]: price=%.2f vol=%d\n",
		q.Security.GetCode(), q.CurPrice, q.Volume)
}
```

#### Multiple K-line types (callback-based with `SubscribeKLines`):

```go
import (
	chanpkg "github.com/shing1211/futuapi4go/pkg/push/chan"
)

// Subscribe to 1-minute and daily K-lines with separate handlers
handlers := map[constant.KLType]func(*push.UpdateKL){
	constant.KLType_K_1Min: func(kl *push.UpdateKL) {
		for _, bar := range kl.KLList {
			fmt.Printf("1MIN KL: %s C=%.2f V=%d\n",
				*bar.Time, *bar.ClosePrice, *bar.Volume)
		}
	},
	constant.KLType_K_Day: func(kl *push.UpdateKL) {
		for _, bar := range kl.KLList {
			fmt.Printf("DAY KL: %s O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
				*bar.Time, *bar.OpenPrice, *bar.HighPrice,
				*bar.LowPrice, *bar.ClosePrice, *bar.Volume)
		}
	},
}

stop := chanpkg.SubscribeKLines(cli, constant.Market_HK, "00700", handlers)
defer stop()
```

### Circuit Breaker for Trading

Protect your trading bot from cascading failures:

```go
cb := breaker.New(
    breaker.WithThreshold(5),
    breaker.WithCooldown(30*time.Second),
)

result, err := cb.Do(func() (interface{}, error) {
    return client.PlaceOrder(cli, accID, market, "00700", side, orderType, price, qty)
})
if err == breaker.ErrOpen {
    fmt.Println("Trading suspended — too many failures")
}
```

### Structured Logging

```go
l := logger.New(
    logger.WithLevel(logger.LevelDebug),
    logger.WithFormat(logger.FormatJSON),
)
l.Info("connected", "addr", "127.0.0.1:11111", "conn_id", 42)
l.Warn("order rejected", "code", "HK.00700", "reason", "insufficient funds")
```

### Code Helpers

```go
import "github.com/shing1211/futuapi4go/pkg/util"

// "HK.00700" → market=1, code="00700"
mkt, code := util.ParseCode("HK.00700")

// Back again
formatted := util.FormatCode(mkt, code) // "HK.00700"

// Market conversion between quote and trading markets
secMkt := util.MarketToTrdSecMarket[mkt]
```

## Client Options

```go
cli := client.New(
    client.WithDialTimeout(10*time.Second),
    client.WithAPISetTimeout(30*time.Second),
    client.WithKeepAliveInterval(30*time.Second),
    client.WithReconnectInterval(5*time.Second),
    client.WithMaxRetries(3),
    client.WithLogLevel(logger.LevelInfo),
)

// Default to simulate trading (safe by default)
cli = cli.WithTradeEnv(constant.TrdEnv_Simulate)
```

## Full API Reference

Every exported function with working examples and quick-reference tables.

> All examples assume `cli := client.New(); cli.Connect("127.0.0.1:11111")`.

---

### Connection & Client

```go
cli := client.New(
    client.WithDialTimeout(10*time.Second),
    client.WithAPISetTimeout(30*time.Second),
    client.WithKeepAliveInterval(30*time.Second),
    client.WithReconnectInterval(5*time.Second),
    client.WithMaxRetries(3),
)
cli = cli.WithTradeEnv(constant.TrdEnv_Simulate) // safe default
cli.Connect("127.0.0.1:11111")

fmt.Println(cli.GetConnID())      // connection ID
fmt.Println(cli.GetLoginUserID()) // Futu user ID
fmt.Println(cli.GetServerVer())  // OpenD version
fmt.Println(cli.IsEncrypt())      // AES encryption enabled?
```

| Function | Signature | Description |
|---|---|---|
| `New` | `New(opts ...Option) *Client` | Create client; defaults to simulate trading |
| `Connect` | `Connect(addr string) error` | Connect to OpenD |
| `Close` | `Close()` | Disconnect |
| `WithTradeEnv` | `WithTradeEnv(trdEnv int32) *Client` | Set real (`0`) or simulate (`1`) |
| `WithTradeMarket` | `WithTradeMarket(trdMkt int32) *Client` | Set default trading market |
| `RegisterHandler` | `RegisterHandler(protoID uint32, h func(uint32, []byte))` | Register push handler |
| `GetConnID` | `GetConnID() uint64` | Connection ID from OpenD |
| `GetServerVer` | `GetServerVer() int32` | OpenD server version |
| `GetLoginUserID` | `GetLoginUserID() uint64` | Logged-in Futu user ID |
| `IsEncrypt` | `IsEncrypt() bool` | True if connection uses AES encryption |
| `CanSendProto` | `CanSendProto(protoID uint32) bool` | Check if proto is available |
| `EnsureConnected` | `EnsureConnected() error` | Return error if not connected |
| `Inner` | `Inner() *futuapi.Client` | Access internal client (advanced) |
| `WithContext` | `WithContext(ctx context.Context) *Client` | Attach context to client |
| `Context` | `Context() context.Context` | Get client's context |
| `GetConn` | `GetConn() *futuapi.Conn` | Access underlying connection (advanced) |

---

### Market Data — Queries

#### GetQuote — real-time price snapshot

```go
quote, err := client.GetQuote(context.Background(), cli, constant.Market_US, "NVDA")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("NVDA: %.2f (open=%.2f high=%.2f low=%.2f vol=%d)\n",
    quote.Price, quote.Open, quote.High, quote.Low, quote.Volume)
```

#### GetKLines — latest K-line bars

```go
klines, err := client.GetKLines(cli, constant.Market_HK, "00700",
    constant.KLType_K_Day, 100)
for _, kl := range klines {
    fmt.Printf("%s O=%.2f H=%.2f L=%.2f C=%.2f\n",
        kl.Time, kl.Open, kl.High, kl.Low, kl.Close)
}
```

#### GetOrderBook — bid/ask depth

```go
book, err := client.GetOrderBook(cli, constant.Market_HK, "00700", 10)
for i, b := range book.Bids {
    fmt.Printf("Bid[%d]: %.2f x %d\n", i, b.Price, b.Volume)
}
for i, a := range book.Asks {
    fmt.Printf("Ask[%d]: %.2f x %d\n", i, a.Price, a.Volume)
}
```

#### GetSecuritySnapshot — multi-stock snapshot

```go
securities := []*qotcommon.Security{
    {Market: ptrInt32(constant.Market_HK), Code: ptrStr("00700")},
    {Market: ptrInt32(constant.Market_HK), Code: ptrStr("09988")},
}
snapshots, err := client.GetSecuritySnapshot(cli, securities)
for _, s := range snapshots {
    fmt.Printf("%s: %.2f\n", s.Security.GetCode(), s.CurPrice)
}
```

#### GetCapitalFlow / GetCapitalDistribution

```go
flows, err := client.GetCapitalFlow(cli, constant.Market_HK, "00700")
for _, f := range flows {
    fmt.Printf("InFlow=%.2f MainInFlow=%.2f\n", f.InFlow, f.MainInFlow)
}

dist, err := client.GetCapitalDistribution(cli, constant.Market_HK, "00700")
if dist != nil {
    fmt.Printf("MainInflow=%.2f BigInflow=%.2f\n", dist.MainInflow, dist.BigInflow)
}
```

| Function | Signature | Description |
|---|---|---|
| `GetQuote` | `GetQuote(ctx, c, market, code) (*Quote, error)` | Real-time quote for one security |
| `GetKLines` | `GetKLines(c, market, code, klType, num) ([]KLine, error)` | Latest K-line bars (up to `num`) |
| `GetOrderBook` | `GetOrderBook(c, market, code, num) (*OrderBook, error)` | Bid/ask depth, `num` levels per side |
| `GetTicker` | `GetTicker(c, market, code, num) ([]Ticker, error)` | Tick-by-tick trades, last `num` |
| `GetRT` | `GetRT(c, market, code) ([]RT, error)` | Intraday time-share data |
| `GetBroker` | `GetBroker(c, market, code, num) ([]Broker, []Broker, error)` | Broker queue (bid, ask) |
| `GetStaticInfo` | `GetStaticInfo(c, market, code) ([]StaticInfo, error)` | Static security info (name, type, lot size) |
| `GetSecuritySnapshot` | `GetSecuritySnapshot(c, securities) ([]*Snapshot, error)` | Full snapshot for multiple securities |
| `GetMarketState` | `GetMarketState(c, market, code) (int32, error)` | Trading status (open/closed/auction...) |
| `GetCapitalFlow` | `GetCapitalFlow(c, market, code) ([]CapitalFlow, error)` | Capital flow (inflow/outflow) |
| `GetCapitalDistribution` | `GetCapitalDistribution(c, market, code) (*CapitalDistribution, error)` | Capital distribution (super/big/mid/small) |
| `GetOwnerPlate` | `GetOwnerPlate(c, market, code) ([]string, error)` | Plates the security belongs to |
| `GetPlateSet` | `GetPlateSet(c, market) ([]Plate, error)` | List plates (industry/region/concept) |
| `GetPlateSecurity` | `GetPlateSecurity(c, market, plateCode) ([]StaticInfo, error)` | Securities in a plate |
| `GetReference` | `GetReference(c, market, code, refType) ([]StaticInfo, error)` | Related securities (warrants, etc.) |
| `GetIpoList` | `GetIpoList(c, market) ([]IpoData, error)` | Upcoming/ongoing IPOs |
| `GetFutureInfo` | `GetFutureInfo(c, code) ([]FutureInfo, error)` | Futures contract info |
| `GetSuspend` | `GetSuspend(c, securities, begin, end) ([]*SuspendInfo, error)` | Suspension dates |
| `GetCodeChange` | `GetCodeChange(c, securities) ([]*CodeChangeInfo, error)` | Code change history |
| `GetHoldingChangeList` | `GetHoldingChangeList(c, market, code, category, begin, end) ([]*HoldingChangeInfo, error)` | Director/holder changes |
| `GetOptionExpirationDate` | `GetOptionExpirationDate(c, market, code) ([]OptionExpiration, error)` | Option expiry dates |
| `GetOptionChain` | `GetOptionChain(c, market, code, indexType, optType, cond, begin, end) ([]*OptChain, error)` | Full option chain |
| `GetWarrant` | `GetWarrant(c, market, code, begin, num, sort, asc, optType, issuer, status) ([]*WarrantData, error)` | Warrant list |
| `StockFilter` | `StockFilter(c, market, begin, num) ([]*StockFilterResult, error)` | Filter stocks by criteria |
| `GetPriceReminder` | `GetPriceReminder(c, market, code) ([]*PriceReminderInfo, error)` | Get price alerts |
| `SetPriceReminder` | `SetPriceReminder(c, market, code, op, type, freq, value, note) (int64, error)` | Add/update/delete price alert |

---

### Market Data — Historical K-Lines

```go
// Fetch all daily K-lines for a date range (auto-paginated)
klines, err := client.RequestHistoryKL(cli,
    constant.Market_HK, "00700",
    constant.KLType_K_Day,
    "2024-01-01", "2025-01-01",
)

// Fetch with custom page size (max 1000 per page)
client.HistoryKLPaginationDelay = 500 * time.Millisecond
klines, err = client.RequestHistoryKLWithLimit(cli,
    constant.Market_HK, "00700",
    constant.KLType_K_Day,
    "2024-01-01", "2025-01-01",
    500, // max per page
)

// Check quota usage
quota, err := client.RequestHistoryKLQuota(cli)
fmt.Printf("Used=%d Remain=%d\n", quota.UsedQuota, quota.RemainQuota)

// Get rehab (split/dividend) factors
rehab, err := client.RequestRehab(cli, constant.Market_HK, "00700")
```

| Function | Signature | Description |
|---|---|---|
| `RequestHistoryKL` | `RequestHistoryKL(c, mkt, code, klType, start, end) ([]KLine, error)` | Auto-paginated historical K-lines |
| `RequestHistoryKLWithLimit` | `RequestHistoryKLWithLimit(c, mkt, code, klType, start, end, maxPerPage) ([]KLine, error)` | With configurable page size |
| `RequestHistoryKLQuota` | `RequestHistoryKLQuota(c) (*HistoryKLQuotaInfo, error)` | API quota usage |
| `RequestRehab` | `RequestRehab(c, market, code) ([]*RehabInfo, error)` | Rehabilitation (split/dividend) factors |
| `GetTradeDate` | `GetTradeDate(c, market, start, end) ([]string, error)` | Market trade dates |
| `RequestTradeDate` | `RequestTradeDate(c, market, start, end, code) ([]string, error)` | Trade dates for a specific security |

---

### Real-Time Subscriptions

#### Subscribe — receive push data for one or more types

```go
// Subscribe to multiple data types at once
// Note: US stocks require subscription before GetQuote works
err := client.Subscribe(cli, constant.Market_US, "NVDA", []constant.SubType{
    constant.SubType_Quote,
    constant.SubType_Ticker,
    constant.SubType_K_1Min,
})
```

#### Channel-based subscription — `chanpkg` (recommended)

##### Single K-line type (channel-based):

```go
quoteCh   := make(chan *push.UpdateBasicQot, 100)
tickerCh  := make(chan *push.UpdateTicker, 100)
orderBookCh := make(chan *push.UpdateOrderBook, 100)
rtCh      := make(chan *push.UpdateRT, 100)
brokerCh  := make(chan *push.UpdateBroker, 100)
klCh      := make(chan *push.UpdateKL, 100)

stopQuote   := chanpkg.SubscribeQuote(cli, constant.Market_US, "NVDA", quoteCh)
stopTicker  := chanpkg.SubscribeTicker(cli, constant.Market_US, "NVDA", tickerCh)
stopOB      := chanpkg.SubscribeOrderBook(cli, constant.Market_US, "NVDA", orderBookCh)
stopRT      := chanpkg.SubscribeRT(cli, constant.Market_US, "NVDA", rtCh)
stopBroker  := chanpkg.SubscribeBroker(cli, constant.Market_US, "NVDA", brokerCh)
stopKL      := chanpkg.SubscribeKLine(cli, constant.Market_US, "NVDA", constant.KLType_K_1Min, klCh)
defer stopQuote()
defer stopTicker()
defer stopOB()
defer stopRT()
defer stopBroker()
defer stopKL()

sig := make(chan os.Signal, 1)
signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

for {
    select {
    case q := <-quoteCh:
        fmt.Printf("QUOTE [%s]: price=%.2f vol=%d\n",
            q.Security.GetCode(), q.CurPrice, q.Volume)
    case t := <-tickerCh:
        if len(t.TickerList) > 0 {
            fmt.Printf("TICKER: %.2f x %d\n",
                t.TickerList[0].GetPrice(), t.TickerList[0].GetVolume())
        }
    case ob := <-orderBookCh:
        if len(ob.OrderBookBidList) > 0 && len(ob.OrderBookAskList) > 0 {
            fmt.Printf("ORDERBOOK: bid=%.2f ask=%.2f\n",
                ob.OrderBookBidList[0].GetPrice(), ob.OrderBookAskList[0].GetPrice())
        }
    case rt := <-rtCh:
        if len(rt.RTList) > 0 {
            fmt.Printf("RT: %.2f avg=%.2f\n",
                rt.RTList[0].GetPrice(), rt.RTList[0].GetAvgPrice())
        }
    case b := <-brokerCh:
        if len(b.BidBrokerList) > 0 {
            fmt.Printf("BROKER: %s pos=%d\n",
                b.BidBrokerList[0].GetName(), b.BidBrokerList[0].GetPos())
        }
    case kl := <-klCh:
        for _, bar := range kl.KLList {
            fmt.Printf("KL: %s C=%.2f V=%d\n",
                *bar.Time, *bar.ClosePrice, *bar.Volume)
        }
    case <-sig:
        fmt.Println("Shutting down...")
        return
    }
}
```

##### Multiple K-line types (callback-based with `SubscribeKLines`):

```go
// Define handlers for different K-line periods
handlers := map[constant.KLType]func(*push.UpdateKL){
    constant.KLType_K_1Min: func(kl *push.UpdateKL) {
        for _, bar := range kl.KLList {
            fmt.Printf("1MIN KL: %s C=%.2f V=%d\n",
                *bar.Time, *bar.ClosePrice, *bar.Volume)
        }
    },
    constant.KLType_K_5Min: func(kl *push.UpdateKL) {
        for _, bar := range kl.KLList {
            fmt.Printf("5MIN KL: %s C=%.2f V=%d\n",
                *bar.Time, *bar.ClosePrice, *bar.Volume)
        }
    },
    constant.KLType_K_Day: func(kl *push.UpdateKL) {
        for _, bar := range kl.KLList {
            fmt.Printf("DAY KL: %s O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
                *bar.Time, *bar.OpenPrice, *bar.HighPrice,
                *bar.LowPrice, *bar.ClosePrice, *bar.Volume)
        }
    },
}

stopKLines := chanpkg.SubscribeKLines(cli, constant.Market_HK, "00700", handlers)
defer stopKLines()

// Wait for Ctrl+C to exit
sig := make(chan os.Signal, 1)
signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
<-sig
fmt.Println("Shutting down...")
```

| Function | Signature | Description |
|---|---|---|
| `Subscribe` | `Subscribe(c, market, code, []SubType) error` | Subscribe to one or more push types |
| `Unsubscribe` | `Unsubscribe(c, market, code, []int32) error` | Unsubscribe specific types |
| `UnsubscribeAll` | `UnsubscribeAll(c) error` | Unsubscribe everything |
| `QuerySubscription` | `QuerySubscription(c) (*GetSubInfoResponse, error)` | Current subscription status |
| `RegQotPush` | `RegQotPush(c, market, code, subtypes, rehabTypes, isReg, isFirst) error` | Register/unregister push |
| `chanpkg.SubscribeQuote` | `(cli, market, code, ch) stopFunc` | Quote push via channel |
| `chanpkg.SubscribeKLine` | `(cli, market, code, KLType, ch) stopFunc` | K-line push via channel |
| `chanpkg.SubscribeTicker` | `(cli, market, code, ch) stopFunc` | Ticker push via channel |
| `chanpkg.SubscribeOrderBook` | `(cli, market, code, ch) stopFunc` | Order book push via channel |
| `chanpkg.SubscribeRT` | `(cli, market, code, ch) stopFunc` | RT push via channel |
| `chanpkg.SubscribeBroker` | `(cli, market, code, ch) stopFunc` | Broker push via channel |
| `chanpkg.SubscribePriceReminder` | `(cli, ch) stopFunc` | Price reminder push via channel |
| `GetSubInfo` | `GetSubInfo(c) (*SubInfo, error)` | Subscription info (quota, types) |

---

### Push Parsing — decode raw push payloads

```go
// Quote updates
cli.RegisterHandler(constant.ProtoID_Qot_UpdateBasicQot, func(pid uint32, body []byte) {
    pq, err := client.ParsePushQuote(body)
    if err != nil || pq == nil {
        return
    }
    fmt.Printf("[%s] %.2f\n", pq.Code, pq.CurPrice)
})

// K-line updates
cli.RegisterHandler(constant.ProtoID_Qot_UpdateKL, func(pid uint32, body []byte) {
    pk, err := client.ParsePushKLine(body)
    if err != nil || pk == nil {
        return
    }
    fmt.Printf("[%s KL] %.2f\n", pk.Code, pk.Close)
})

// Order updates (requires SubAccPush)
cli.RegisterHandler(constant.ProtoID_Trd_UpdateOrder, func(pid uint32, body []byte) {
    ou, err := client.ParsePushOrderUpdate(body)
    if err != nil || ou == nil {
        return
    }
    fmt.Printf("Order %d: status=%d\n", ou.OrderID, ou.OrderStatus)
})
```

| Function | Signature | Description |
|---|---|---|
| `ParsePushQuote` | `ParsePushQuote(body) (*PushQuote, error)` | Decode quote push (ProtoID 3005) |
| `ParsePushKLine` | `ParsePushKLine(body) (*PushKLine, error)` | Decode K-line push (3007) |
| `ParsePushOrderBook` | `ParsePushOrderBook(body) (*PushOrderBook, error)` | Decode order book push (3013) |
| `ParsePushTicker` | `ParsePushTicker(body) (*PushTicker, error)` | Decode ticker push (3011) |
| `ParsePushRT` | `ParsePushRT(body) (*PushRT, error)` | Decode RT push (3009) |
| `ParsePushBroker` | `ParsePushBroker(body) (*PushBroker, error)` | Decode broker push (3015) |
| `ParsePushOrderUpdate` | `ParsePushOrderUpdate(body) (*PushOrderUpdate, error)` | Decode order update push (2208) |
| `ParsePushOrderFill` | `ParsePushOrderFill(body) (*PushOrderFill, error)` | Decode fill push (2218) |

---

### Trading — Account & Funds

```go
// List all accounts
accounts, err := client.GetAccountList(cli)
for _, acc := range accounts {
    fmt.Printf("AccID=%d Env=%d Markets=%v\n",
        acc.AccID, acc.TrdEnv, acc.TrdMarketAuthList)
}

// Unlock trading (required before placing orders)
if err := client.UnlockTrading(cli, "your_md5_password"); err != nil {
    log.Fatal(err)
}

// Quick funds for first account
funds, err := client.GetFunds(cli, 0)
fmt.Printf("Power=%.2f Cash=%.2f Assets=%.2f\n",
    funds.Power, funds.Cash, funds.TotalAssets)

// Full account info with per-currency and per-market breakdown
funds, err = client.GetAccountInfo(cli, accID, constant.TrdMarket_HK)
for _, ci := range funds.CashInfoList {
    fmt.Printf("Currency=%d Cash=%.2f Available=%.2f\n",
        ci.Currency, ci.Cash, ci.AvailableBalance)
}

// Max tradable quantities before placing an order
max, err := client.GetMaxTrdQtys(cli, accID, constant.TrdMarket_HK,
    "00700", constant.OrderType_Normal, 350.0)
fmt.Printf("MaxCashBuy=%.2f MaxSell=%.2f\n", max.MaxCashBuy, max.MaxPositionSell)

// AccTradingInfo — includes initial margin requirements
info, err := client.GetAccTradingInfo(cli, accID, constant.TrdMarket_HK,
    "00700", constant.OrderType_Normal, 350.0)
fmt.Printf("MaxBuy=%.2f LongIM=%.2f\n", info.MaxCashBuy, info.LongRequiredIM)
```

| Function | Signature | Description |
|---|---|---|
| `GetAccountList` | `GetAccountList(c) ([]Account, error)` | All trading accounts |
| `UnlockTrading` | `UnlockTrading(c, pwdMD5) error` | Unlock trading with MD5-hashed password |
| `GetFunds` | `GetFunds(c, accID) (*Funds, error)` | Quick funds for first account |
| `GetAccountInfo` | `GetAccountInfo(c, accID, market) (*Funds, error)` | Full funds with multi-currency/multi-market breakdown |
| `GetMaxTrdQtys` | `GetMaxTrdQtys(c, accID, market, code, orderType, price) (*MaxTrdQtysInfo, error)` | Maximum buy/sell quantities |
| `GetAccTradingInfo` | `GetAccTradingInfo(c, accID, market, code, orderType, price) (*AccTradingInfo, error)` | Max quantities + initial margin |
| `GetOrderFee` | `GetOrderFee(c, accID, market, orderIDExList) ([]*OrderFeeInfo, error)` | Fee breakdown per order |
| `GetMarginRatio` | `GetMarginRatio(c, accID, market, securities) ([]*MarginRatioInfo, error)` | Margin ratios for securities |
| `SubAccPush` | `SubAccPush(c, accIDList) error` | Subscribe to account push notifications |

---

### Trading — Orders

```go
// Place a buy limit order
result, err := client.PlaceOrder(cli,
    accID,
    constant.TrdMarket_HK, // trading market
    "00700",              // code
    constant.TrdSide_Buy,  // side
    constant.OrderType_Normal, // limit order
    350.0,                // price
    100,                  // quantity
)
fmt.Printf("OrderID=%d OrderIDEx=%s\n", result.OrderID, result.OrderIDEx)

// Modify order price and quantity
_, err = client.ModifyOrder(cli, accID, constant.TrdMarket_HK,
    result.OrderID,             // order ID
    constant.ModifyOrderOp_Normal, // modify (not cancel)
    355.0, 200)                 // new price, new qty

// Cancel all open orders
err = client.CancelAllOrder(cli, accID, constant.TrdMarket_HK, constant.TrdEnv_Real)

// Active orders
orders, err := client.GetOrderList(cli, accID)
for _, o := range orders {
    fmt.Printf("Order %d: %s %s @ %.2f qty=%.0f status=%d\n",
        o.OrderID, o.Code, o.TrdSide, o.Price, o.Qty, o.OrderStatus)
}

// Historical orders
hist, err := client.GetHistoryOrderList(cli, accID, constant.TrdMarket_HK,
    "2024-01-01", "2025-12-31")

// Order fills
fills, err := client.GetOrderFillList(cli, accID)
for _, f := range fills {
    fmt.Printf("Fill %d: %s @ %.2f x %.0f\n", f.FillID, f.Code, f.Price, f.Qty)
}

// Cash flow
flows, err := client.GetFlowSummary(cli, accID, constant.TrdMarket_HK, "", 0)
// date="" means today; direction 0=all, 1=in, 2=out
for _, f := range flows {
    fmt.Printf("%s %s: %.2f\n", f.ClearingDate, f.CashFlowType, f.CashFlowAmount)
}
```

| Function | Signature | Description |
|---|---|---|
| `PlaceOrder` | `PlaceOrder(c, accID, market, code, side, orderType, price, qty) (*PlaceOrderResult, error)` | Place a new order |
| `ModifyOrder` | `ModifyOrder(c, accID, market, orderID, op, price, qty) (*ModifyOrderResponse, error)` | Modify price/qty or cancel |
| `CancelAllOrder` | `CancelAllOrder(c, accID, market, trdEnv) error` | Cancel all open orders |
| `ReconfirmOrder` | `ReconfirmOrder(c, accID, market, orderID, reason) (*ReconfirmOrderResult, error)` | Reconfirm order requiring verification |
| `GetOrderList` | `GetOrderList(c, accID) ([]Order, error)` | Active (open) orders |
| `GetHistoryOrderList` | `GetHistoryOrderList(c, accID, market, start, end) ([]Order, error)` | Historical orders |
| `GetOrderFillList` | `GetOrderFillList(c, accID) ([]OrderFill, error)` | Today's order fills |
| `GetHistoryOrderFillList` | `GetHistoryOrderFillList(c, accID, market) ([]OrderFill, error)` | Historical fills |
| `GetFlowSummary` | `GetFlowSummary(c, accID, market, date, direction) ([]*FlowSummaryInfo, error)` | Cash flow entries |

---

### Trading — Positions

```go
positions, err := client.GetPositionList(cli, accID)
for _, p := range positions {
    fmt.Printf("%s: qty=%.0f cost=%.2f cur=%.2f pnl=%.2f (%.2f%%)\n",
        p.Code, p.Quantity, p.CostPrice, p.CurPrice, p.PnL, p.PnLRate)
}
```

| Function | Signature | Description |
|---|---|---|
| `GetPositionList` | `GetPositionList(c, accID) ([]Position, error)` | Current positions with P&L |

---

### User Security (Watchlist)

```go
// List all watchlist groups
groups, err := client.GetUserSecurityGroup(cli)
for _, g := range groups {
    fmt.Printf("Group: %s (type=%d)\n", g.Name, g.GroupType)
}

// Get securities in a group
infos, err := client.GetUserSecurity(cli, "My Watchlist")

// Add/remove securities from a group
err = client.ModifyUserSecurity(cli, "My Watchlist",
    constant.ModifyUserSecurityOp_Add, // or _Del
    constant.Market_US, []string{"NVDA", "AAPL"})
```

| Function | Signature | Description |
|---|---|---|
| `GetUserSecurityGroup` | `GetUserSecurityGroup(c) ([]UserSecurityGroup, error)` | All watchlist groups |
| `GetUserSecurity` | `GetUserSecurity(c, groupName) ([]StaticInfo, error)` | Securities in a group |
| `ModifyUserSecurity` | `ModifyUserSecurity(c, groupName, op, market, codes) error` | Add/delete securities from group |

---

### System

```go
// Global connection state
state, err := client.GetGlobalState(cli)
fmt.Printf("QotLogined=%v TrdLogined=%v ServerBuild=%d\n",
    state.QotLogined, state.TrdLogined, state.ServerBuildNo)

// User info
user, err := client.GetUserInfo(cli)
fmt.Printf("UserID=%d Nick=%s APILevel=%s\n", user.UserID, user.NickName, user.ApiLevel)
```

| Function | Signature | Description |
|---|---|---|
| `GetGlobalState` | `GetGlobalState(c) (*GlobalState, error)` | OpenD connection and login state |
| `GetUserInfo` | `GetUserInfo(c) (*UserInfo, error)` | Futu account user info |
| `GetDelayStatistics` | `GetDelayStatistics(c) (*DelayStatistics, error)` | Latency stats (known proto2 issue with OpenD serverVer=1003) |

---

### Circuit Breaker

```go
import "github.com/shing1211/futuapi4go/pkg/breaker"

cb := breaker.New(
    breaker.WithThreshold(5),
    breaker.WithCooldown(30*time.Second),
    breaker.WithOnOpen(func() { fmt.Println("Circuit OPENED") }),
)

// Wrap any API call
result, err := cb.Do(func() (interface{}, error) {
    return client.PlaceOrder(...)
})
if err == breaker.ErrOpen {
    fmt.Println("Trading suspended — circuit is open")
}

// Or for void-returning calls
err = cb.DoVoid(func() error {
    return client.PlaceOrder(...)
})

// Manual control
fmt.Printf("State=%s Failures=%d\n", cb.State(), cb.Failures())
cb.Reset() // close the circuit
```

| Function | Signature | Description |
|---|---|---|
| `breaker.New` | `New(opts ...Option) *Breaker` | Create circuit breaker |
| `breaker.Do` | `(b *Breaker) Do(fn func() (interface{}, error)) (interface{}, error)` | Execute with protection |
| `breaker.DoVoid` | `(b *Breaker) DoVoid(fn func() error) error` | Execute void function |
| `breaker.State` | `(b *Breaker) State() State` | Current state (Closed/Open/HalfOpen) |
| `breaker.Allow` | `(b *Breaker) Allow() bool` | Check if request is allowed |
| `breaker.RecordSuccess` | `(b *Breaker) RecordSuccess()` | Record a success |
| `breaker.RecordFailure` | `(b *Breaker) RecordFailure()` | Record a failure |
| `breaker.Reset` | `(b *Breaker) Reset()` | Reset to closed |
| `breaker.Stats` | `(b *Breaker) Stats() Stats` | Diagnostic info |
| `breaker.ErrOpen` | `var` | Error returned when circuit is open |

---

### Structured Logging

```go
import (
    "github.com/shing1211/futuapi4go/pkg/logger"
    futulogger "github.com/shing1211/futuapi4go/pkg/logger"
)

l := futulogger.New(
    futulogger.WithLevel(futulogger.LevelDebug),
    futulogger.WithFormat(futulogger.FormatJSON), // or FormatText
)
l.Info("connected", "addr", "127.0.0.1:11111", "conn_id", 42)
l.Warn("order rejected", "code", "HK.00700", "reason", "insufficient funds")
l.Error("connection lost", "err", err)

// Package-level defaults
logger.SetLevel(logger.LevelInfo)
logger.Info("hello", "key", "value")
```

| Function | Signature | Description |
|---|---|---|
| `logger.New` | `New(opts ...Option) *Logger` | Create logger instance |
| `logger.Info/Debug/Warn/Error` | `(l *Logger) Info(msg, fields...)` | Log at specific level |
| `logger.Fatal` | `(l *Logger) Fatal(msg, fields...)` | Log and exit |
| `logger.SetLevel` | `SetLevel(lvl Level)` | Set global level |
| `logger.SetFormat` | `SetFormat(fmt Format)` | Set text (`FormatText`) or JSON (`FormatJSON`) |
| `logger.SetOutput` | `SetOutput(w io.Writer)` | Set output destination |

---

### Constants Reference

```go
// Markets (quote)
constant.Market_HK  // 1  — Hong Kong
constant.Market_US  // 11 — United States
constant.Market_SH   // 21 — Shanghai A-share
constant.Market_SZ   // 22 — Shenzhen A-share
constant.Market_SG   // 31 — Singapore
constant.Market_JP   // 41 — Japan
constant.Market_AU   // 51 — Australia

// Trading markets
constant.TrdMarket_HK      // 1
constant.TrdMarket_US      // 2
constant.TrdMarket_CN      // 3
constant.TrdMarket_Futures // 5

// Trading environment
constant.TrdEnv_Simulate // 0 (default — safe)
constant.TrdEnv_Real     // 1

// Trading sides
constant.TrdSide_Buy      // 1
constant.TrdSide_Sell     // 2
constant.TrdSide_SellShort // 3
constant.TrdSide_BuyBack  // 4

// Order types
constant.OrderType_Normal      // 1 — limit order (recommended)
constant.OrderType_Market      // 2 — market order
constant.OrderType_Stop       // 10 — stop market
constant.OrderType_StopLimit  // 11 — stop limit

// Modify order operations
constant.ModifyOrderOp_Normal // 1 — modify price/qty
constant.ModifyOrderOp_Cancel // 2 — cancel order

// K-line types (for GetKLines / RequestHistoryKL)
constant.KLType_K_1Min  // 1
constant.KLType_K_5Min  // 2
constant.KLType_K_15Min // 3
constant.KLType_K_30Min // 4
constant.KLType_K_60Min // 5
constant.KLType_K_Day   // 6
constant.KLType_K_Week  // 7
constant.KLType_K_Month // 8

// Subscription types (for Subscribe / chanpkg)
constant.SubType_Quote     // 1
constant.SubType_OrderBook // 2
constant.SubType_Ticker    // 4
constant.SubType_RT        // 5
constant.SubType_Broker    // 14
constant.SubType_K_1Min   // 11
constant.SubType_K_5Min   // 7
constant.SubType_K_15Min  // 8
constant.SubType_K_30Min  // 9
constant.SubType_K_60Min  // 10
constant.SubType_K_Day    // 6
constant.SubType_K_Week   // 12
constant.SubType_K_Month  // 13
constant.SubType_K_Quarter // 15
constant.SubType_K_Year   // 16
constant.SubType_K_3Min   // 17

// Rehab (price adjustment)
constant.RehabType_None    // 0 — no adjustment
constant.RehabType_Forward  // 1 — forward (QFQ)
constant.RehabType_Backward // 2 — backward (BQF)

// Plate set types
constant.PlateSetType_Industry // 1
constant.PlateSetType_Region   // 2
constant.PlateSetType_Concept   // 3

// Market states
constant.MarketState_Morning     // 3  — morning session
constant.MarketState_Afternoon  // 5  — afternoon session
constant.MarketState_Closed     // 6  — market closed
constant.MarketState_PreMarketBegin // 7 — US pre-market

// Price reminder
constant.PriceReminderOpAdd    // 1 — add alert
constant.PriceReminderOpUpdate // 2 — update alert
constant.PriceReminderOpDelete // 3 — delete alert
```


## Build & Test

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test ./...       # Run the full test suite
go test -race ./... # Race detector
```

## Architecture

```
futuapi4go/
├── client/               # Public high-level API (recommended)
├── internal/client/      # TCP connection, packet I/O, reconnect, keep-alive
├── pkg/
│   ├── qot/              # Market data — quotes, K-lines, order book, tick data...
│   ├── trd/              # Trading — orders, positions, funds, history...
│   ├── sys/              # System — global state, user info
│   ├── push/             # Push notification parsers
│   ├── push/chan/         # Channel-based push delivery
│   ├── breaker/           # Circuit breaker pattern
│   ├── logger/            # Structured leveled logging
│   ├── util/              # Code parsing, market helpers
│   ├── constant/          # Python-style constants + String() methods
│   └── pb/               # 78 protobuf-generated types (v10.4.6408)
├── api/proto/            # Original .proto definitions (v10.4.6408)
└── cmd/demo/             # Interactive demo
```

## Python Migration

Coming from the Python `futu-api` SDK? The Python-style constants make the transition feel familiar:

```go
// Python: ft.Market.HK, ft.TrdSide.BUY, ft.KLType.K_DAY
// Go:     constant.Market_HK, constant.TrdSide_Buy, constant.KLType_K_Day
```

See the full [Python Migration Guide](PYTHON_MIGRATION_GUIDE.md) for side-by-side comparisons of every API.

## Known Issues

- **`GetDelayStatistics`** / **`GetTradeDate`** — proto2 wire-format incompatibility with OpenD serverVer=1003. Workaround: these calls are skipped gracefully in the demo.
- Both issues are in the SDK's protobuf marshaling layer and will be fixed in a future release.

## License

Apache License 2.0 — see [LICENSE](LICENSE).

> ⚠️ **Trading Disclaimer**: This SDK is a software utility. Trading financial instruments carries significant risk. Always test thoroughly in simulate mode before using real funds.
