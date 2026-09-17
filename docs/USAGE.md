# FutuAPI4Go — Quick Start Guide / 快速开始

[English](#english) | [中文](#chinese-中文-简体)

---

## English

### Installation

```bash
go get github.com/shing1211/futuapi4go@latest
```

### Quickstart (Recommended)

```go
package main

import (
	"context"
	"log"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
	cli, err := futuapi.NewClientFromEnv() // reads FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, etc.
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	quote, err := client.GetQuote(ctx, cli, constant.Market_US, "AAPL")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AAPL: %.2f", quote.Price)
}
```

### Quickstart (Manual)

```go
cli := client.New(client.WithLogLevel(3))
if err := cli.Connect("127.0.0.1:11111"); err != nil {
	log.Fatal(err)
}
defer cli.Close()

ctx := context.Background()
quote, err := client.GetQuote(ctx, cli, constant.Market_US, "AAPL")
if err != nil {
	log.Fatal(err)
}
log.Printf("AAPL: %.2f", quote.Price)
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `FUTU_OPEND_ADDR` | OpenD address (`host:port`) | `127.0.0.1:11111` |
| `FUTU_RSA_PUBLIC_KEY` | RSA public key PEM (file path or inline) | — |
| `FUTU_RSA_PRIVATE_KEY` | RSA private key PEM (file path or inline) | — |
| `FUTU_ENCRYPT` | Set to `"1"` or `"true"` to enable encryption | — |
| `FUTU_LOG_LEVEL` | Log level: `0`=info, `1`=warn, `2`=error, `3`=silent | `0` |

The trading environment is **not** read from the environment; set it with the
`WithTradeEnv(trdEnv)` method (defaults to simulate).

### Core Patterns

#### Fluent API

```go
// High-level wrappers
cli.Quote().GetBasicQot(ctx, securities)
cli.Trade().PlaceOrder(ctx, req)
cli.System().GetGlobalState(ctx)
```

#### Channels & Typed Callbacks

```go
// Channel-based (streaming)
ch := make(chan *client.PushQuote, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	log.Printf("%s: %.2f", q.Code, q.CurPrice)
}

// Callback-based (chainable on client)
cli.OnQuote(func(q *client.PushQuote) error {
	log.Printf("%s: %.2f", q.Code, q.CurPrice)
	return nil
}).OnOrder(func(o *client.PushOrderUpdate) error {
	log.Printf("Order %s: status=%d", o.OrderIDEx, o.OrderStatus)
	return nil
})
```

#### Order Builder

```go
import "github.com/shing1211/futuapi4go/pkg/trd"

order, err := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).
	At(350.0).
	Build()
if err != nil {
	log.Fatal(err)
}
```

#### Circuit Breaker

```go
import "github.com/shing1211/futuapi4go/pkg/breaker"

cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, err := cb.Do(func() (interface{}, error) {
	return client.GetQuote(ctx, cli, constant.Market_HK, "00700")
})
if err == breaker.ErrOpen {
	log.Println("Requests suspended — circuit open")
}
```

#### Historical K-Lines (Auto-Paginated)

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

#### Subscribe + Get

```go
client.Subscribe(ctx, cli, constant.Market_US, "AAPL", []constant.SubType{constant.SubType_Basic})

// Then call one-shot API
quote, _ := client.GetQuote(ctx, cli, constant.Market_US, "AAPL")
```

### Troubleshooting

| Issue | Likely Cause |
|-------|-------------|
| `connection refused` | OpenD not running. Check `FUTU_OPEND_ADDR`. |
| No data from `GetQuote` (US stocks) | Need to `Subscribe` first for US market data. |
| `The packet body SHA1 signature is incorrect` | Outdated OpenD (< v10.5). Upgrade OpenD. |
| `没有解锁交易` | Call `UnlockTrading` with your trading password MD5. |
| `模拟交易不支持` | Feature unavailable in simulate mode. Construct the client with `WithTradeEnv(constant.TrdEnv_Real)`. |

### Distributed Tracing

The SDK supports distributed tracing via an optional OpenTelemetry adapter.
When enabled, spans are automatically created for API requests, connection
lifecycle, and push notifications.

```go
import (
    "github.com/shing1211/futuapi4go/pkg/tracing"
    "github.com/shing1211/futuapi4go/pkg/tracing/otel"
)

// Install OpenTelemetry backend
tracing.SetTracer(otel.NewTracer("my-trading-app"))

// All subsequent API calls, connect/disconnect, and push handlers
// will generate spans automatically. Export spans via your preferred
// OTel exporter (stdout, Jaeger, OTLP, etc.).
```

No code changes are required in your application — just set the tracer once
after creating the client. See `examples/97_opentelemetry_tracing` in the
demo repository for a complete example with OTel stdout exporter.

---

### Connection State Machine

Monitor connection lifecycle with the built-in state machine:

```go
// State constants are re-exported by the public client package.
switch cli.State() {
case client.StateConnected:
    fmt.Println("Connected")
case client.StateDisconnected:
    fmt.Println("Disconnected")
case client.StateReconnecting:
    fmt.Println("Reconnecting...")
case client.StateClosing:
    fmt.Println("Closing")
}

// React to state transitions
cli := client.New(
    client.WithOnStateChange(func(old, new client.ConnState) {
        log.Printf("State: %d → %d", old, new)
    }),
)
```

### Graceful Shutdown

Drain in-flight requests before closing with configurable timeout:

```go
if err := cli.Shutdown(5 * time.Second); err != nil {
    log.Printf("shutdown error: %v", err)
}
// New requests return ErrClientClosing during shutdown
```

### K-Line Data Cache

Built-in LRU+TTL cache to avoid redundant API calls:

```go
import "github.com/shing1211/futuapi4go/pkg/cache"

klCache := cache.NewKLCache(
    cache.WithMaxEntries(2000),
    cache.WithTTL(5*time.Minute),
)
cachedCli := cache.NewKLCachedClient(cli.Inner(), klCache)
klines, err := cachedCli.GetKL(ctx, rehabType, klType, security)
```

### Order Pre-Flight Validation

Validate orders before submission to catch common issues:

```go
import "github.com/shing1211/futuapi4go/pkg/trd"

warnings := trd.ValidateOrder(&trd.OrderValidationInput{
    Order:       req,
    MarketOpen:  true,
    BuyingPower: 50000,
    MaxBuyQty:   1000,
    MaxSellQty:  1000,
})
if trd.HasErrors(warnings) {
    for _, w := range warnings {
        log.Printf("Validation: %s", w.Message)
    }
}
```

### Audit Logging

Record all trade operations with structured slog output:

```go
audit := trd.NewAuditLogger(slog.Default())

// trd.PlaceOrder takes the internal client; obtain it via cli.Inner().
resp, err := trd.PlaceOrder(ctx, cli.Inner(), req)
audit.LogPlaceOrder(req, resp.OrderID, err) // JSON: {"op":"PlaceOrder","code":"US.AAPL","success":true}
```

### OpenTelemetry Metrics

Export SDK metrics via OTLP alongside Prometheus:

```go
import "github.com/shing1211/futuapi4go/pkg/tracing/otel"

meter, err := otel.NewOTelMeter()
if err != nil {
    log.Fatal(err)
}
meter.RecordConnection("tcp")
meter.RecordAPICall("3001", "success", 12*time.Millisecond)
```

### Structured Logging

Route SDK events into your own structured logger (levels and attributes are
preserved; the client's `connID`/`userID` are added to every event):

```go
import (
    "log/slog"
    "os"

    "github.com/shing1211/futuapi4go/client"
)

// Any slog.Logger works — a zerolog/zap bridge, or the standard JSON handler.
sl := slog.New(slog.NewJSONHandler(os.Stderr, nil))

cli := client.New(
    client.WithSlogLogger(sl),
    client.WithLogLevelName("warn"), // debug | info | warn | error | silent
)
```

`debug` includes one line per inbound packet (`recv Trd_GetFunds (2101)
serial=535 bytes=121`), which is useful for tracing the wire protocol but noisy
at the account-poll rate; `warn` is the recommended default.

---

## Chinese 中文 (简体)

### 安裝

```bash
go get github.com/shing1211/futuapi4go@latest
```

### 快速開始（推薦）

```go
package main

import (
	"context"
	"log"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
	cli, err := futuapi.NewClientFromEnv() // 自動讀取環境變數
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.Background()
	quote, err := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("00700: %.2f", quote.Price)
}
```

### 快速開始（手動）

```go
cli := client.New(client.WithLogLevel(3))
if err := cli.Connect("127.0.0.1:11111"); err != nil {
	log.Fatal(err)
}
defer cli.Close()

ctx := context.Background()
quote, err := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
if err != nil {
	log.Fatal(err)
}
log.Printf("00700: %.2f", quote.Price)
```

### 環境變數配置

| 變數 | 說明 | 預設值 |
|------|------|--------|
| `FUTU_OPEND_ADDR` | OpenD 位址（`host:port`） | `127.0.0.1:11111` |
| `FUTU_RSA_PUBLIC_KEY` | RSA 公鑰 PEM（檔案路徑或內容） | — |
| `FUTU_RSA_PRIVATE_KEY` | RSA 私鑰 PEM（檔案路徑或內容） | — |
| `FUTU_ENCRYPT` | 設為 `"1"` 或 `"true"` 啟用加密 | — |
| `FUTU_LOG_LEVEL` | 日誌級別：0=資訊, 1=警告, 2=錯誤, 3=靜默 | `0` |

交易環境**不**從環境變數讀取；請使用 `WithTradeEnv(trdEnv)` 方法設定（預設為模擬）。

### 核心模式

#### Fluent API

```go
cli.Quote().GetBasicQot(ctx, securities)
cli.Trade().PlaceOrder(ctx, req)
cli.System().GetGlobalState(ctx)
```

#### 頻道與回呼

```go
// 頻道模式（串流）
ch := make(chan *client.PushQuote, 100)
stop, _ := chanpkg.SubscribeQuote(ctx, cli, constant.Market_HK, "00700", ch)
defer stop()
for q := range ch {
	log.Printf("%s 現價: %.2f", q.Code, q.CurPrice)
}

// 回呼模式（鏈式呼叫）
cli.OnQuote(func(q *client.PushQuote) error {
	log.Printf("%s 現價: %.2f", q.Code, q.CurPrice)
	return nil
}).OnOrder(func(o *client.PushOrderUpdate) error {
	log.Printf("訂單 %s 狀態: %d", o.OrderIDEx, o.OrderStatus)
	return nil
})
```

#### 訂單建立器

```go
order := trd.NewOrder(accID, constant.TrdMarket_HK, constant.TrdEnv_Simulate).
	Buy("00700", 100).
	At(350.0).
	Build()
```

#### 斷路器

```go
cb := breaker.New(breaker.WithThreshold(5), breaker.WithCooldown(30*time.Second))
result, err := cb.Do(func() (interface{}, error) {
	return client.PlaceOrder(ctx, cli, accID, ...)
})
```

#### 自動分頁歷史 K 線

```go
klines, err := client.RequestHistoryKL(ctx, cli,
	constant.Market_HK, "00700",
	constant.KLType_K_Day,
	"2024-01-01", "2025-01-01")
```

#### 訂閱後查詢

```go
client.Subscribe(ctx, cli, constant.Market_US, "AAPL", []constant.SubType{constant.SubType_Basic})
quote, _ := client.GetQuote(ctx, cli, constant.Market_US, "AAPL")
```

### 疑難排解

| 問題 | 可能原因 |
|------|----------|
| `connection refused` | OpenD 未啟動。請檢查 `FUTU_OPEND_ADDR`。 |
| US 股票 `GetQuote` 無資料 | 美股需要先 `Subscribe`。港股不需要。 |
| `没有解锁交易` | 需要先呼叫 `UnlockTrading` 解鎖交易密碼。 |
| `模拟交易不支持` | 模擬模式不支援該功能。請以 `WithTradeEnv(constant.TrdEnv_Real)` 建立客戶端。 |

### 分散式追蹤

SDK 支援透過可選的 OpenTelemetry 適配器進行分散式追蹤。
啟用後，API 請求、連接生命周期和推送通知會自動生成 span。

```go
import (
    "github.com/shing1211/futuapi4go/pkg/tracing"
    "github.com/shing1211/futuapi4go/pkg/tracing/otel"
)

// 安裝 OpenTelemetry 後端
tracing.SetTracer(otel.NewTracer("my-trading-app"))

// 後續的 API 調用、連接/斷開和推送處理器
// 都會自動生成 span。通過您偏好的 OTel 匯出器
//（stdout、Jaeger、OTLP 等）導出 span。
```

只需在創建客戶端後設置一次 tracer——無需修改應用程式代碼。
請參閱 demo 倉庫中的 `examples/97_opentelemetry_tracing` 獲取完整範例。

### 連線狀態機

監控連線生命週期的狀態機：

```go
switch cli.State() {
case client.StateConnected:
    fmt.Println("已連線")
case client.StateDisconnected:
    fmt.Println("未連線")
case client.StateReconnecting:
    fmt.Println("重新連線中...")
}

// 監聽狀態變化
cli := client.New(
    client.WithOnStateChange(func(old, new client.ConnState) {
        log.Printf("狀態: %d → %d", old, new)
    }),
)
```

### 優雅關閉

在關閉前等待進行中的請求完成：

```go
if err := cli.Shutdown(5 * time.Second); err != nil {
    log.Printf("關閉錯誤: %v", err)
}
```

### K 線資料快取

內建 LRU + TTL 快取，減少重複 API 請求：

```go
import "github.com/shing1211/futuapi4go/pkg/cache"

klCache := cache.NewKLCache(cache.WithMaxEntries(2000))
cachedCli := cache.NewKLCachedClient(cli.Inner(), klCache)
klines, err := cachedCli.GetKL(ctx, rehabType, klType, security)
```

### 委託前驗證

送出委託前檢查常見問題：

```go
warnings := trd.ValidateOrder(&trd.OrderValidationInput{
    Order: req, MarketOpen: true, BuyingPower: 50000,
})
if trd.HasErrors(warnings) {
    for _, w := range warnings { log.Printf("驗證: %s", w.Message) }
}
```

### 審計日誌

以結構化日誌記錄所有交易操作：

```go
audit := trd.NewAuditLogger(slog.Default())
resp, err := trd.PlaceOrder(ctx, cli, req)
audit.LogPlaceOrder(req, resp.OrderID, err)
```

### OpenTelemetry 指標

透過 OTLP 匯出 SDK 指標：

```go
meter, err := otel.NewOTelMeter()
meter.RecordAPICall("3001", "success", 12*time.Millisecond)
```

### 結構化日誌

將內部日誌輸出為 JSON 格式：

```go
sl := slog.New(slog.NewJSONHandler(os.Stderr, nil))
cli := client.New(client.WithSlogLogger(sl), client.WithLogLevelName("warn"))
```
