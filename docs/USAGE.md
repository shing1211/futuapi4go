# FutuAPI4Go — Quick Start Guide / 快速开始

[English](#english) | [中文](#chinese)

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
    "log"

    futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
    // One-call connect — reads FUTU_OPEND_ADDR, FUTU_RSA_PUBLIC_KEY, etc.
    // from environment by default, falls back to 127.0.0.1:11111.
    cli, err := futuapi.NewClientFromEnv()
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // ... use cli
}
```

### Quickstart (Manual)

```go
package main

import (
    "context"
    "log"

    "github.com/shing1211/futuapi4go/client"
)

func main() {
    cli := client.New(
        client.WithLogLevel(3), // silent
    )
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // Get a quote
    ctx := context.Background()
    quote, err := cli.GetQuote(ctx, "US.AAPL")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("AAPL: %.2f", quote.CurPrice)
}
```

### Subscribing to Real-Time Push

```go
// Register typed callbacks before connecting.
cli := client.New().
    OnQuote(func(q *client.PushQuote) error {
        log.Printf("%s: %.2f", q.Code, q.CurPrice)
        return nil
    }).
    OnOrder(func(o *client.PushOrderUpdate) error {
        log.Printf("Order %s: status=%d", o.OrderID, o.OrderStatus)
        return nil
    })

cli.Connect("127.0.0.1:11111")

// Then subscribe to symbols:
cli.Subscribe(ctx, []string{"US.AAPL"}, client.SubType_Basic, true)
```

### Trading Example

```go
acc := cli.FindAccount(accounts)
order, err := cli.PlaceOrder(ctx, acc, "US.AAPL",
    client.TrdEnv_Simulate,
    client.Side_Buy,
    client.OrderType_Normal,
    100,     // quantity
    150.0,   // price
    0,       // adjust limit
)
```

### Environment Variables (`WithEnvConfig`)

| Variable | Description |
|----------|-------------|
| `FUTU_OPEND_ADDR` | OpenD address (`host:port`), default `127.0.0.1:11111` |
| `FUTU_RSA_PUBLIC_KEY` | RSA public key PEM (file path or inline) |
| `FUTU_RSA_PRIVATE_KEY` | RSA private key PEM (file path or inline) |
| `FUTU_ENCRYPT` | `"1"` or `"true"` to enable encryption |
| `FUTU_LOG_LEVEL` | `0`=info, `1`=warn, `2`=error, `3`=silent |
| `FUTU_TRD_ENV` | `"real"` or `"simulate"` |

---

## Chinese 中文 (繁體)

### 安裝

```bash
go get github.com/shing1211/futuapi4go@latest
```

### 快速開始（推薦）

```go
package main

import (
    "log"

    futuapi "github.com/shing1211/futuapi4go/pkg/futuapi"
)

func main() {
    // 一鍵連接，自動讀取環境變數
    cli, err := futuapi.NewClientFromEnv()
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()
}
```

### 快速開始（手動）

```go
package main

import (
    "context"
    "log"

    "github.com/shing1211/futuapi4go/client"
)

func main() {
    cli := client.New(client.WithLogLevel(3))
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    ctx := context.Background()
    quote, err := cli.GetQuote(ctx, "US.AAPL")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("AAPL: %.2f", quote.CurPrice)
}
```

### 訂閱即時推送

```go
cli := client.New().
    OnQuote(func(q *client.PushQuote) error {
        log.Printf("%s 現價: %.2f", q.Code, q.CurPrice)
        return nil
    }).
    OnOrder(func(o *client.PushOrderUpdate) error {
        log.Printf("訂單 %s 狀態: %d", o.OrderID, o.OrderStatus)
        return nil
    })

cli.Connect("127.0.0.1:11111")
cli.Subscribe(ctx, []string{"US.AAPL"}, client.SubType_Basic, true)
```

### 交易範例

```go
acc := cli.FindAccount(accounts)
order, err := cli.PlaceOrder(ctx, acc, "US.AAPL",
    client.TrdEnv_Simulate,
    client.Side_Buy,
    client.OrderType_Normal,
    100,    // 數量
    150.0,  // 價格
    0,      // 調整限價
)
```

### 環境變數配置

| 變數 | 說明 |
|------|------|
| `FUTU_OPEND_ADDR` | OpenD 位址（預設 `127.0.0.1:11111`） |
| `FUTU_RSA_PUBLIC_KEY` | RSA 公鑰 PEM（檔案路徑或內容） |
| `FUTU_RSA_PRIVATE_KEY` | RSA 私鑰 PEM（檔案路徑或內容） |
| `FUTU_ENCRYPT` | 設為 `"1"` 或 `"true"` 啟用加密 |
| `FUTU_LOG_LEVEL` | 日誌級別：0=資訊, 1=警告, 2=錯誤, 3=靜默 |
| `FUTU_TRD_ENV` | 交易環境：`"real"` 或 `"simulate"` |
