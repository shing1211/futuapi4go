# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.6.2-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.5.6508-blue?style=flat-square" alt="Futu Proto Version">
</p>

> **Go-native. Type-safe. Production-ready.** The most complete and ergonomic Go SDK for Futu OpenAPI — market data, trading, real-time push, and more.

## Install

```bash
go get github.com/shing1211/futuapi4go@v0.6.2
```

Requires Go 1.26+ and a running [Futu OpenD](https://www.futunn.com/en/overview) instance.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
)

func main() {
	cli := client.New()
	defer cli.Close()

	// Connect to OpenD (default: simulate trading)
	if err := cli.Connect("127.0.0.1:11111"); err != nil {
		fmt.Fprintf(os.Stderr, "connect failed: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	// Market data
	quote, err := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
	if err != nil {
		fmt.Fprintf(os.Stderr, "quote failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("HK.00700: price=%.2f open=%.2f high=%.2f\n",
		quote.Price, quote.Open, quote.High)

	// Trading — list accounts, unlock, place order
	accounts, _ := client.GetAccountList(ctx, cli)
	if len(accounts) == 0 {
		fmt.Println("no accounts")
		return
	}
	accID := accounts[0].AccID

	_ = client.UnlockTrading(ctx, cli, "your_md5_password") // skip in simulate mode
	result, err := client.PlaceOrder(ctx, cli, accID,
		constant.TrdMarket_HK, "00700",
		constant.TrdSide_Buy, constant.OrderType_Normal,
		350.0, 100)
	if err != nil {
		fmt.Fprintf(os.Stderr, "order failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OrderID=%d OrderIDEx=%s\n", result.OrderID, result.OrderIDEx)
}
```

> **Note:** US stocks require subscribing before `GetQuote` works. HK stocks do not.

## Examples

For more complete, runnable examples covering every API surface — including
real-time push, trading workflows, historical data, and circuit breaker patterns
— see the demo repository:

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## Key Features

### Real-Time Push via Channels

Stop polling — let data come to you.

```go
import (
	chanpkg "github.com/shing1211/futuapi4go/pkg/push/chan"
)

// Quote updates stream into the channel
ch := make(chan *push.UpdateBasicQot, 100)
stop, err := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
if err != nil {
	log.Fatal(err)
}
defer stop()

for q := range ch {
	fmt.Printf("[%s] price=%.2f vol=%d\n", q.Security.GetCode(), q.CurPrice, q.Volume)
}

// Multiple K-line periods with a shared channel filter
ch2 := make(chan *push.UpdateKL, 100)
stop2, err := chanpkg.SubscribeKLines(ctx, cli,
	int32(constant.Market_HK), "00700",
	[]int32{int32(constant.KLType_K_1Min), int32(constant.KLType_K_Day)}, ch2)
if err != nil {
	log.Fatal(err)
}
defer stop2()

for kl := range ch2 {
	switch kl.KlType {
	case int32(constant.KLType_K_1Min):
		fmt.Printf("1MIN %s C=%.2f\n", *kl.KLList[0].Time, *kl.KLList[0].ClosePrice)
	case int32(constant.KLType_K_Day):
		fmt.Printf("DAY %s O=%.2f H=%.2f L=%.2f C=%.2f\n",
			*kl.KLList[0].Time, *kl.KLList[0].OpenPrice,
			*kl.KLList[0].HighPrice, *kl.KLList[0].LowPrice, *kl.KLList[0].ClosePrice)
	}
}
```

### Typed Push Callbacks

Alternative to channels — register callbacks directly on the client for
quote updates, order notifications, fills, K-lines, order book changes,
tickers, and more.

```go
cli.OnQuote(func(q *push.UpdateBasicQot) {
	fmt.Printf("[%s] price=%.2f vol=%d\n", q.Security.GetCode(), q.CurPrice, q.Volume)
}).OnOrder(func(o *push.TrdUpdateOrder) {
	fmt.Printf("Order %s: status=%d\n", o.GetOrderIDEx(), o.GetOrderStatus())
}).OnOrderFill(func(f *push.TrdUpdateOrderFill) {
	fmt.Printf("Fill %s: qty=%d price=%.2f\n", f.GetOrderIDEx(), f.GetQty(), f.GetPrice())
})
```

Callbacks are chainable — each `On*()` returns the `*Client` for fluent setup.

### Circuit Breaker

Protect trading from cascading failures.

```go
import "github.com/shing1211/futuapi4go/pkg/breaker"

cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, err := cb.Do(func() (interface{}, error) {
	return client.PlaceOrder(ctx, cli, accID, constant.TrdMarket_HK, "00700",
		constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100)
})
if err == breaker.ErrOpen {
	fmt.Println("Trading suspended — too many failures")
}
```

### Fluent API

```go
// High-level client wrappers (recommended)
cli.Quote().GetBasicQot(ctx, securities)
cli.Trade().PlaceOrder(ctx, req)
cli.System().GetGlobalState(ctx)

// Order builder
import "github.com/shing1211/futuapi4go/pkg/trd"

order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).
	At(350.0).
	Build()
```

### Historical K-Lines (auto-paginated)

```go
klines, err := client.RequestHistoryKL(ctx, cli,
	constant.Market_HK, "00700",
	constant.KLType_K_Day,
	"2024-01-01", "2025-01-01")

for _, kl := range klines {
	fmt.Printf("%s O=%.2f H=%.2f L=%.2f C=%.2f\n",
		kl.Time, kl.Open, kl.High, kl.Low, kl.Close)
}
```

### Structured Logging

```go
import futulogger "github.com/shing1211/futuapi4go/pkg/logger"

l := futulogger.New(futulogger.WithLevel(futulogger.LevelDebug), futulogger.WithFormat(futulogger.FormatJSON))
l.Info("connected", "addr", "127.0.0.1:11111", "conn_id", 42)
```

### Code Helpers

```go
import "github.com/shing1211/futuapi4go/pkg/util"

// "HK.00700" → market=1, code="00700"
mkt, code := util.ParseCode("HK.00700")
// Back again
s := util.FormatCode(mkt, code) // "HK.00700"
```

## Package Map

| Package | Purpose |
|---------|---------|
| `client` | High-level wrappers — recommended entry point |
| `pkg/qot` | Market data: quotes, K-lines, order book, tick data... |
| `pkg/trd` | Trading: orders, positions, funds, history... |
| `pkg/sys` | System: global state, user info |
| `pkg/push` | Push notification parsers |
| `pkg/push/chan` | Channel-based real-time push delivery |
| `pkg/breaker` | Circuit breaker pattern |
| `pkg/logger` | Structured leveled logging |
| `pkg/util` | Code parsing (`ParseCode`, `FormatCode`), market helpers |
| `pkg/constant` | Typed constants with `String()` methods |
| `pkg/futuapi` | Convenience re-export — `NewClient()`, `NewClientFromEnv()` |
| `pkg/pb/*` | 79 protobuf types (v10.5.6508) |

## Common APIs

### Connection

```go
// Manual config
cli := client.New(
	client.WithDialTimeout(10*time.Second),
	client.WithAPISetTimeout(30*time.Second),
)
cli = cli.WithTradeEnv(constant.TrdEnv_Simulate)

// Or from env vars: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, FUTU_ENCRYPT, FUTU_LOG_LEVEL
cli, err := client.NewClientFromEnv()
if err != nil {
	log.Fatal(err)
}

cli.Connect("127.0.0.1:11111")
// cli.GetConnID(), cli.GetServerVer(), cli.IsEncrypt(), cli.GetLoginUserID()
// cli.CanSendProto(protoID)
```

### Market Data

| Function | Description |
|---|---|
| `GetQuote(ctx, c, market, code)` | Real-time quote |
| `GetKLines(ctx, c, market, code, klType, num)` | Latest K-line bars |
| `GetOrderBook(ctx, c, market, code, num)` | Bid/ask depth |
| `GetTicker(ctx, c, market, code, num)` | Tick-by-tick trades |
| `GetStaticInfo(ctx, c, market, code)` | Security name, type, lot size |
| `GetSecuritySnapshot(ctx, c, securities)` | Full snapshot for multiple securities |
| `GetCapitalFlow(ctx, c, market, code)` | Capital flow |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | Historical K-lines (auto-paginated) |
| `RequestHistoryKLQuota(ctx, c)` | API quota usage |

### Trading

| Function | Description |
|---|---|
| `GetAccountList(ctx, c)` | All trading accounts |
| `UnlockTrading(ctx, c, pwdMD5)` | Unlock trading |
| `GetFunds(ctx, c, accID)` | Account funds and power |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | Place order |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | Modify or cancel order |
| `GetOrderList(ctx, c, accID)` | Active orders |
| `GetPositionList(ctx, c, accID)` | Current positions with P&L |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | Historical orders |
| `GetOrderFillList(ctx, c, accID)` | Order fills |

### Subscriptions

| Function | Description |
|---|---|
| `Subscribe(ctx, c, market, code, []SubType)` | Subscribe to push types |
| `Unsubscribe(ctx, c, market, code, []SubType)` | Unsubscribe |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | Quote push via channel |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | Single K-line push via channel |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | Multi K-line push with filter |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | Ticker push via channel |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | Order book push via channel |

## Build & Test

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test ./...       # Run full suite
go test -race ./... # Race detector
```

## Architecture

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

All communication is via Protocol Buffers over TCP. See [DESIGN.md](DESIGN.md) for full architecture details.

## See Also

- [CHANGELOG](CHANGELOG.md) — version history and release notes
- [DESIGN](DESIGN.md) — architecture, design decisions, and API patterns
- [ENHANCEMENT_PLAN](ENHANCEMENT_PLAN.md) — upcoming features and roadmap
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) — runnable examples for every feature

## License

Apache License 2.0 — see [LICENSE](LICENSE).

> **Trading Disclaimer**: This SDK is a software utility. Trading financial instruments carries significant risk. Always test thoroughly in simulate mode before using real funds.
