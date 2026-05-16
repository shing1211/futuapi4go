# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.6.2-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.5.6508-blue?style=flat-square" alt="Futu Proto Version">
</p>

> **Go-native. Type-safe. Production-ready.** The most complete and ergonomic Go SDK for [Futu OpenAPI](https://www.futunn.com/en/overview) — market data, trading, and real-time push. All communication via Protocol Buffers over TCP.

- 79 protobuf types covering every Futu OpenAPI service
- One-liner connect with automatic env config (`NewClientFromEnv`)
- Real-time push via channels or typed callbacks
- Fluent API: `cli.Quote().GetBasicQot()`, `cli.Trade().PlaceOrder()`
- Circuit breaker, structured logging, and trading utilities included

## Table of Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [Key Features](#key-features)
- [Examples](#examples)
- [Package Map](#package-map)
- [Common APIs](#common-apis)
- [Build & Test](#build--test)
- [Architecture](#architecture)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

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
	"log"

	futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
	// One-call connect (reads env: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, ...)
	cli, err := futuapi.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	quote, err := cli.GetQuote(ctx, "HK.00700")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: price=%.2f high=%.2f low=%.2f vol=%d\n",
		quote.Code, quote.CurPrice, quote.HighPrice, quote.LowPrice, quote.Volume)
}
```

> **Note:** US stocks require subscribing before `GetQuote` works. HK stocks do not.

## Key Features

### Real-Time Push

Stop polling — receive data as it arrives. Two delivery models:

```go
// Option 1: Channels (streaming)
ch := make(chan *push.UpdateBasicQot, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}

// Option 2: Typed callbacks (chainable on client)
cli.OnQuote(func(q *push.UpdateBasicQot) {
	fmt.Printf("[%s] price=%.2f\n", q.Security.GetCode(), q.CurPrice)
}).OnOrder(func(o *push.TrdUpdateOrder) {
	fmt.Printf("Order %s: status=%d\n", o.GetOrderIDEx(), o.GetOrderStatus())
})
```

### Market Data

```go
// One-shot
quote, _ := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
snapshots, _ := client.GetSecuritySnapshot(ctx, cli, securities)

// Auto-paginated historical K-lines
klines, _ := client.RequestHistoryKL(ctx, cli, constant.Market_HK, "00700",
	constant.KLType_K_Day, "2024-01-01", "2025-01-01")
```

### Trading

```go
accounts, _ := client.GetAccountList(ctx, cli)
accID := accounts[0].AccID

client.UnlockTrading(ctx, cli, "md5_password")
result, _ := client.PlaceOrder(ctx, cli, accID,
	constant.TrdMarket_HK, "00700",
	constant.TrdSide_Buy, constant.OrderType_Normal, 350.0, 100)

// Fluent order builder
order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).At(350.0).Build()
```

### Utilities

```go
// Circuit breaker
cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, _ := cb.Do(func() (interface{}, error) {
	return client.PlaceOrder(ctx, cli, accID, ...)
})

// Structured logging
l := futulogger.New(futulogger.WithLevel(futulogger.LevelDebug))
l.Info("connected", "addr", "127.0.0.1:11111")

// Code helpers
mkt, code := util.ParseCode("HK.00700")  // market=1, code="00700"
s := util.FormatCode(mkt, code)          // "HK.00700"
```

## Examples

For complete, runnable examples covering every API surface — including real-time push, trading workflows, historical data, and strategy patterns:

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

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
).WithTradeEnv(constant.TrdEnv_Simulate)

// From env vars: FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, FUTU_ENCRYPT, FUTU_LOG_LEVEL
cli, _ := client.NewClientFromEnv()

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
go test -race ./... # Full suite with race detector
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

All communication is via Protocol Buffers over TCP. See [DESIGN.md](DESIGN.md) for full architecture decisions and [internal/testutil/mock](internal/testutil/mock/) for the mock OpenD server used in tests.

## Troubleshooting

| Error | Likely Cause |
|-------|-------------|
| `connection refused` | OpenD not running. Check `FUTU_OPEND_ADDR`. |
| no data from `GetQuote` (US stocks) | Must call `Subscribe` first for US market. HK does not need it. |
| `The packet body SHA1 signature is incorrect` (very old OpenD) | Upgrade OpenD to v10.5+. The SDK uses SHA1(ciphertext) which OpenD accepts. |
| `解析protobuf协议失败` | Missing required C2S fields in request body. |
| `模拟交易不支持` | Feature not available in simulate mode; use `WithTradeEnv(TrdEnv_Real)`. |

## Contributing

1. Fork the repository.
2. Create a feature branch (`git checkout -b feat/my-change`).
3. Ensure all existing tests pass: `go test -race ./...`
4. Add tests for any new functionality.
5. Run `go vet ./...` and fix any warnings.
6. Open a pull request.

See [CHANGELOG.md](CHANGELOG.md) for the version history and [ENHANCEMENT_PLAN.md](ENHANCEMENT_PLAN.md) for the roadmap.

## See Also

- [CHANGELOG](CHANGELOG.md) — version history and release notes
- [USAGE Guide](docs/USAGE.md) — detailed setup, environment, and advanced patterns
- [DESIGN](DESIGN.md) — architecture, design decisions, API patterns
- [ENHANCEMENT_PLAN](ENHANCEMENT_PLAN.md) — upcoming features and roadmap
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) — runnable examples for every feature

## License

Apache License 2.0 — see [LICENSE](LICENSE).

> **Trading Disclaimer**: This SDK is a software utility. Trading financial instruments carries significant risk. Always test thoroughly in simulate mode before using real funds.
