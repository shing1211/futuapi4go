# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.9.0-00ADD8?style=flat-square" alt="Version">
</p>

> **Go-native. Type-safe. Production-ready.** The most complete and ergonomic Go SDK for Futu OpenAPI — market data, trading, real-time push, and more.

## Why futuapi4go?

Futu's official SDK is Python-first. futuapi4go is **Go-first** — built from the ground up for Go developers who want:

- **Compile-time safety** — structs over DataFrames, no runtime surprises
- **Go concurrency** — goroutines, channels, context cancellation baked in
- **No Python dependency** — pure Go, deploy anywhere with `go build`
- **More protos** — 78 protos vs Python's ~50, including futures, flow summaries, and all push types
- **Modern infrastructure** — circuit breaker, structured logging, channel-based push delivery, connection pool with health checks

## Install

```bash
go get github.com/shing1211/futuapi4go@v0.9.0
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
        panic(err)
    }

    // US stocks require subscription before GetQuote works
    if err := client.Subscribe(cli, constant.Market_US, "NVDA",
        []int32{int32(constant.SubType_Quote), int32(constant.SubType_K_1Min)}); err != nil {
        panic(err)
    }

    // Real-time quote (one-shot)
    quote, err := client.GetQuote(context.Background(), cli, constant.Market_US, "NVDA")
    if err != nil {
        panic(err)
    }
    fmt.Printf("US.NVDA: price=%.2f open=%.2f high=%.2f low=%.2f vol=%d\n",
        quote.Price, quote.Open, quote.High, quote.Low, quote.Volume)

    // Subscribe to live K-line updates via channel
    klCh := make(chan *push.UpdateKL, 100)
    stop := chanpkg.SubscribeKLine(cli, constant.Market_US, "NVDA", constant.KLType_K_1Min, klCh)
    defer stop()

    // Graceful shutdown on Ctrl+C
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

    for {
        select {
        case kl := <-klCh:
            for _, bar := range kl.KLList {
                fmt.Printf("KL: time=%s O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
                    *bar.Time, *bar.OpenPrice, *bar.HighPrice, *bar.LowPrice, *bar.ClosePrice, *bar.Volume)
            }
        case <-sig:
            fmt.Println("Shutting down...")
            return
        }
    }
}
```

> **Note:** US stocks require subscribing before `GetQuote` works. HK stocks don't have this requirement.

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

```go
import (
    "github.com/shing1211/futuapi4go/pkg/push/chan" as chanpkg
)

// Quote updates stream into the channel
ch := make(chan *push.UpdateBasicQot, 100)
stop := chanpkg.SubscribeQuote(cli, constant.Market_HK, "00700", ch)
defer stop()

for q := range ch {
    fmt.Printf("Bid: %.2f | Ask: %.2f\n", q.BidPrice[0], q.AskPrice[0])
}
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

### Trading
`GetAccountList` · `UnlockTrading` · `PlaceOrder` · `ModifyOrder` · `CancelAllOrder` · `GetPositionList` · `GetAccountInfo` · `GetFunds` · `GetAccTradingInfo` · `GetMaxTrdQtys` · `GetOrderFee` · `GetMarginRatio` · `GetOrderList` · `GetHistoryOrderList` · `GetOrderFillList` · `GetHistoryOrderFillList` · `GetFlowSummary` · `SubAccPush` · `ReconfirmOrder`

### Market Data
`GetQuote` · `GetKLines` · `GetOrderBook` · `GetTicker` · `GetRT` · `GetBroker` · `GetStaticInfo` · `GetTradeDate` · `GetFutureInfo` · `GetPlateSet` · `GetPlateSecurity` · `GetOwnerPlate` · `GetReference` · `GetIpoList` · `GetMarketState` · `GetCapitalFlow` · `GetCapitalDistribution` · `GetSecuritySnapshot` · `GetOptionChain` · `GetOptionExpirationDate` · `GetWarrant` · `StockFilter` · `GetSuspend` · `GetCodeChange` · `GetHoldingChangeList` · `GetUserSecurityGroup` · `ModifyUserSecurity` · `GetPriceReminder` · `SetPriceReminder` · `RequestHistoryKL` · `RequestHistoryKLQuota` · `RequestRehab` · `RequestTradeDate`

### System
`GetGlobalState` · `GetUserInfo` · `GetDelayStatistics` · `GetLoginUserID` · `IsEncrypt` · `CanSendProto`

### Subscriptions
`Subscribe` · `Unsubscribe` · `UnsubscribeAll` · `QuerySubscription` · `RegQotPush`

## Testing Without a Real Account

```bash
# Terminal 1 — mock OpenD server
go run cmd/simulator/main.go

# Terminal 2 — your code
go run ./cmd/demo/main.go
```

The simulator handles all 78 protobufs with realistic mock responses. Perfect for CI/CD.

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
│   └── pb/               # 78 protobuf-generated types
├── api/proto/             # Original .proto definitions
├── cmd/simulator/         # Mock OpenD for testing
└── cmd/demo/              # Interactive demo
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
