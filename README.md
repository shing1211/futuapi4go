# futuapi4go

> High-performance, production-ready Go SDK for Futu OpenAPI. Covers all market data, trading, and push notification APIs.

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square)](LICENSE)

---

## Features

| Feature | Description |
|---------|-------------|
| **78 protobufs** | All Python SDK protos + 28 extras (futures, flow, all push types) |
| **Type-safe** | Go structs with compile-time safety vs Python DataFrames |
| **Circuit breaker** | Built-in `pkg/breaker` for resilient trading |
| **Structured logging** | `pkg/logger` — text + JSON, leveled (Debug/Info/Warn/Error) |
| **Channel push** | `pkg/push/chan` — goroutine-safe push via Go channels |
| **Connection pool** | `internal/client/pool.go` — reusable connections with health checks |
| **Connection resilience** | Auto-reconnect with configurable backoff |
| **Context support** | Request-level cancellation on all APIs |

---

## Quick Start

```bash
go get github.com/shing1211/futuapi4go
```

```go
package main

import (
    "fmt"
    "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/pkg/constant"
)

func main() {
    cli := client.New()
    defer cli.Close()

    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        panic(err)
    }

    // Get quote
    quote, err := client.GetQuote(nil, cli, constant.Market_HK, "00700")
    if err != nil {
        panic(err)
    }
    fmt.Printf("HK.00700: %.2f\n", quote.Price)

    // Subscribe to real-time data
    client.Subscribe(cli, constant.Market_HK, "00700",
        []int32{constant.SubType_Quote, constant.SubType_KL_1Min})

    // Place a simulate order
    accounts, _ := client.GetAccountList(cli)
    result, err := client.PlaceOrder(cli, accounts[0].AccID, constant.Market_HK,
        "00700", constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Order placed: %d\n", result.OrderID)
}
```

---

## Package Overview

| Package | Purpose |
|---------|---------|
| `client` | High-level public API (recommended) — all wrappers return typed structs |
| `internal/client` | Low-level connection, packet I/O, reconnect, keep-alive |
| `pkg/qot` | Market data APIs (quotes, K-lines, order book, tick data, etc.) |
| `pkg/trd` | Trading APIs (orders, positions, funds, history) |
| `pkg/sys` | System APIs (global state, user info, delay statistics) |
| `pkg/push` | Push notification parsers (quote updates, order fills) |
| `pkg/push/chan` | Channel-based push delivery — subscribe via channels |
| `pkg/breaker` | Circuit breaker pattern — prevents cascading failures |
| `pkg/logger` | Structured leveled logging (text + JSON formats) |
| `pkg/util` | Code parsing, market conversion helpers |
| `pkg/constant` | Python-style constants + String() methods on all enums |
| `pkg/pb/*` | Protobuf-generated types for all Futu OpenAPI protos |

---

## Key Packages

### `pkg/constant` — Python-style Constants

```go
import "github.com/shing1211/futuapi4go/pkg/constant"

// Markets
constant.Market_HK  // 1
constant.Market_US  // 11
constant.Market_SH  // 21
constant.Market_SZ  // 22
constant.Market_SG  // 31
constant.Market_JP  // 41
constant.Market_AU  // 51
constant.Market_MY  // 61
constant.Market_CA  // 71
constant.Market_FX  // 81

// K-Line Types
constant.KLType_K_1Min  // 1
constant.KLType_K_5Min  // 2
constant.KLType_K_15Min  // 3
constant.KLType_K_30Min  // 4
constant.KLType_K_60Min  // 5
constant.KLType_K_Day    // 6
constant.KLType_K_Week   // 7
constant.KLType_K_Month  // 8

// Trading
constant.TrdEnv_Real      // 0
constant.TrdEnv_Simulate   // 1
constant.TrdSide_Buy       // 1
constant.TrdSide_Sell      // 2
constant.OrderType_Normal   // 0
constant.OrderType_Market   // 2

// Rehab
constant.RehabType_None     // 0 (不复权)
constant.RehabType_Forward  // 1 (前复权 QFQ)
constant.RehabType_Backward // 2 (后复权 HFQ)

// SubTypes
constant.SubType_Quote
constant.SubType_K_Day
constant.SubType_K_1Min
constant.SubType_Ticker
constant.SubType_OrderBook
constant.SubType_RT
constant.SubType_Broker
```

All enums have `String()` methods: `constant.Market_HK.String()` → `"Market_HK"`.

### `pkg/push/chan` — Channel-Based Push

```go
import (
    "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/pkg/push/chan"
    "github.com/shing1211/futuapi4go/pkg/push"
)

// Subscribe to real-time quotes via channel
ch := make(chan *push.UpdateBasicQot, 100)
stop := chanpkg.SubscribeQuote(cli, constant.Market_HK, "00700", ch)
defer stop()

for q := range ch {
    fmt.Printf("Quote: %.2f\n", q.CurPrice)
}

// Subscribe to K-line updates
klCh := make(chan *push.UpdateKL, 100)
stopK := chanpkg.SubscribeKLine(cli, constant.Market_HK, "00700", constant.KLType_K_1Min, klCh)
defer stopK()
```

### `pkg/breaker` — Circuit Breaker

```go
import "github.com/shing1211/futuapi4go/pkg/breaker"

cb := breaker.New(
    breaker.WithThreshold(5),
    breaker.WithCooldown(30*time.Second),
)

result, err := cb.Do(func() (interface{}, error) {
    return client.PlaceOrder(cli, accID, market, "00700", side, orderType, price, qty)
})
if err == breaker.ErrOpen {
    fmt.Println("Circuit open — trading suspended")
}
```

### `pkg/logger` — Structured Logging

```go
import "github.com/shing1211/futuapi4go/pkg/logger"

l := logger.New(
    logger.WithLevel(logger.LevelDebug),
    logger.WithFormat(logger.FormatJSON),
    logger.WithOutput(os.Stdout),
)

l.Info("connected", "addr", "127.0.0.1:11111", "conn_id", 42)
l.Warn("order rejected", "code", "HK.00700", "reason", "insufficient funds")
```

---

## Client Options

```go
cli := client.New(
    client.WithDialTimeout(10 * time.Second),
    client.WithAPISetTimeout(30 * time.Second),
    client.WithKeepAliveInterval(30 * time.Second),
    client.WithReconnectInterval(5 * time.Second),
    client.WithMaxRetries(3),
    client.WithLogLevel(1), // 0=info, 1=warn, 2=error
)

// Trading: set default environment
cli = cli.WithTradeEnv(constant.TrdEnv_Simulate)  // default is simulate
```

---

## Code Helpers (`pkg/util`)

```go
import "github.com/shing1211/futuapi4go/pkg/util"

// Parse "HK.00700" → market=1, code="00700"
mkt, code := util.ParseCode("HK.00700")

// Format back → "HK.00700"
formatted := util.FormatCode(mkt, code)

// Market conversion
secMkt := util.MarketToTrdSecMarket[mkt]
qotMkt := util.TrdMarketToQotMarket(secMkt)

// Validate
if util.IsMarketValid(mkt) {
    // ...
}
```

---

## Full API Reference

See [`docs/API_REFERENCE.md`](docs/API_REFERENCE.md) for the complete API documentation.

### Trading APIs

| Function | Description |
|----------|-------------|
| `GetAccountList` | List all trading accounts |
| `UnlockTrading` | Unlock trading with MD5 password |
| `PlaceOrder` | Place buy/sell order |
| `ModifyOrder` | Modify or cancel order |
| `CancelAllOrder` | Cancel all pending orders |
| `GetPositionList` | Get current positions |
| `GetAccountInfo` | Full account info (multi-currency cash, per-market assets) |
| `GetFunds` | Account funds (auto-selects first account) |
| `GetAccTradingInfo` | Max tradable quantities for a security |
| `GetMaxTrdQtys` | Calculate max buy/sell quantities |
| `GetOrderFee` | Calculate order fees |
| `GetMarginRatio` | Margin ratio for short selling |
| `GetOrderList` | Today's orders |
| `GetHistoryOrderList` | Historical orders |
| `GetOrderFillList` | Today's order fills |
| `GetHistoryOrderFillList` | Historical order fills |
| `GetFlowSummary` | Account cash flow entries |
| `SubAccPush` | Subscribe to account push updates |
| `ReconfirmOrder` | Re-confirm a rejected order |

### Market Data APIs

| Function | Description |
|----------|-------------|
| `GetQuote` | Real-time quote |
| `GetKLines` | K-line (candlestick) data |
| `GetOrderBook` | Order book (depth) |
| `GetTicker` | Tick-by-tick trades |
| `GetRT` | Real-time minute (分时) |
| `GetBroker` | Broker queue |
| `GetStaticInfo` | Static security info |
| `GetTradeDate` | Trading dates |
| `GetFutureInfo` | Futures contract info |
| `GetPlateSet` | Available plate sets |
| `GetPlateSecurity` | Securities in a plate |
| `GetOwnerPlate` | Plates a security belongs to |
| `GetReference` | Related securities |
| `GetIpoList` | IPO calendar |
| `GetMarketState` | Market open/close state |
| `GetCapitalFlow` | Capital flow data |
| `GetCapitalDistribution` | Capital distribution |
| `GetSecuritySnapshot` | Multi-security snapshots |
| `GetOptionChain` | Option chain data |
| `GetOptionExpirationDate` | Option expiry dates |
| `GetWarrant` | Warrant data |
| `StockFilter` | Stock screener |
| `GetSuspend` | Suspended securities |
| `GetCodeChange` | Code change info |
| `GetHoldingChangeList` | Director holding changes |
| `GetUserSecurityGroup` | User security groups |
| `ModifyUserSecurity` | Add/remove from groups |
| `GetPriceReminder` | Price alerts |
| `SetPriceReminder` | Set price alert |
| `RequestHistoryKL` | Historical K-lines |
| `RequestHistoryKLQuota` | K-line quota info |
| `RequestRehab` | Rehabilitation (split/bonus) data |
| `RequestTradeDate` | Online trading dates |

### System APIs

| Function | Description |
|----------|-------------|
| `GetGlobalState` | OpenD connection + market statuses |
| `GetUserInfo` | User account info |
| `GetDelayStatistics` | Connection delay stats |

### Subscription APIs

| Function | Description |
|----------|-------------|
| `Subscribe` | Subscribe to real-time data |
| `Unsubscribe` | Unsubscribe |
| `UnsubscribeAll` | Unsubscribe all |
| `QuerySubscription` | Query subscription status |
| `RegQotPush` | Register push notifications |

---

## OpenD Simulator (Testing)

```bash
# Terminal 1: run the simulator (in futuapi4go repo)
go run cmd/simulator/main.go

# Terminal 2: run your demo
go run ./cmd/demo
```

---

## Testing

```bash
go test ./...          # Run all tests
go build ./...         # Build
go vet ./...           # Lint
```

---

## Architecture

```
futuapi4go/
├── client/           # Public high-level API (recommended)
├── internal/client/  # TCP connection, packet I/O, reconnect
├── pkg/
│   ├── qot/          # Market data (quotes, K-lines, etc.)
│   ├── trd/          # Trading (orders, positions, funds)
│   ├── sys/           # System (global state, user info)
│   ├── push/          # Push notification parsers
│   ├── push/chan/    # Channel-based push delivery
│   ├── breaker/       # Circuit breaker pattern
│   ├── logger/        # Structured leveled logging
│   ├── util/          # Code parsing, market helpers
│   ├── constant/      # Python-style constants
│   └── pb/            # Generated protobuf code (78 protos)
├── api/proto/         # Original .proto definitions
├── cmd/simulator/    # Mock OpenD for testing
└── cmd/demo/          # Interactive demo
```

---

## License

Apache License 2.0 — see [LICENSE](LICENSE)

**Disclaimer:** Trading financial instruments carries significant risk. This SDK is a software utility only and does not provide financial advice.
