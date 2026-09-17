# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.18.1-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.10.7008-blue?style=flat-square" alt="Futu Proto Version">
  <a href="https://shing1211.github.io/futuapi4go/"><img src="https://img.shields.io/badge/Docs-GitHub%20Pages-97CAFF?style=flat-square&logo=github" alt="Docs"></a>
</p>

> **⚠️ 正在積極開發中**  
> 本 SDK 正在積極開發。雖然可對接真實的 Futu OpenD 執行個體正常運作，
> 但部分 proto 回應欄位可能尚未對應。API 與型別可能在小版本之間變更。
> 在依賴任何欄位之前，請針對你的特定使用情境審核 `client/types.go`，並對照
> [Futu Proto Reference](https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md)。

> **Go 原生。型別安全。可用於生產。** 最完整且最易用的 [Futu OpenAPI](https://www.futunn.com/en/overview) Go SDK —— 行情資料、交易與即時推送。所有通訊皆透過 TCP 上的 Protocol Buffers 進行。

[English](./README.md) · [简体中文](./README.zh-Hans.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Español](./README.es.md)

> 本文件是英文 [README](./README.md) 的社群翻譯。**英文版本為準。**
> 同步於 / Last synced: c417534

- 184 個 protobuf 型別，涵蓋所有 Futu OpenAPI 服務
- 一行程式碼連線，自動讀取環境變數設定（`NewClientFromEnv`）
- 透過 channel 或具型別的回呼進行即時推送
- 流暢 API：`cli.Quote().GetBasicQot()`、`cli.Trade().PlaceOrder()`
- 使用 OpenTelemetry 的分散式追蹤與指標（透過 `pkg/tracing/otel` 選擇啟用）
- 連線狀態機、優雅關閉與自動重新連線
- 限流器、斷路器與重試已接入每一次 API 呼叫
- K 線資料快取（LRU + TTL）、下單預先檢查、稽核日誌
- 使用 goreleaser 自動化發布

## 目錄

- [安裝](#安裝)
- [快速開始](#快速開始)
- [核心功能](#核心功能)
- [範例](#範例)
- [套件對應表](#套件對應表)
- [常用 API](#常用-api)
- [建置與測試](#建置與測試)
- [架構](#架構)
- [疑難排解](#疑難排解)
- [貢獻](#貢獻)
- [授權條款](#授權條款)

## 安裝

```bash
go get github.com/shing1211/futuapi4go@v0.18.1
```

需要 Go 1.26+，以及一個正在執行的 [Futu OpenD](https://www.futunn.com/en/overview) 執行個體。

## 快速開始

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

> **注意：** 美股需要先訂閱，`GetQuote` 才能運作。港股則不需要。

## 核心功能

### 即時推送

停止輪詢 —— 資料到達時即刻接收。兩種遞送模式：

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

### 行情資料

```go
// One-shot
quote, _ := client.GetQuote(ctx, cli, constant.Market_HK, "00700")
snapshots, _ := client.GetSecuritySnapshot(ctx, cli, securities)

// Auto-paginated historical K-lines
klines, _ := client.RequestHistoryKL(ctx, cli, constant.Market_HK, "00700",
	constant.KLType_K_Day, "2024-01-01", "2025-01-01")
```

### 交易

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

### 工具

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

## 範例

涵蓋所有 API 介面的完整可執行範例 —— 包括即時推送、交易工作流程、歷史資料與策略模式：

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## 套件對應表

| 套件 | 用途 |
|---------|---------|
| `client` | 高階封裝 —— 建議的進入點 |
| `pkg/qot` | 行情資料：報價、K 線、訂單簿、逐筆資料…… |
| `pkg/trd` | 交易：訂單、持倉、資金、歷史…… |
| `pkg/sys` | 系統：全域狀態、使用者資訊 |
| `pkg/push` | 推送通知解析器 |
| `pkg/push/chan` | 基於 channel 的即時推送遞送 |
| `pkg/breaker` | 斷路器模式 |
| `pkg/cache` | K 線資料快取（LRU + TTL） |
| `pkg/logger` | 結構化分級日誌 |
| `pkg/util` | 代碼解析（`ParseCode`、`FormatCode`）、市場輔助函式 |
| `pkg/constant` | 具 `String()` 方法的型別化常數 |
| `pkg/degradation` | 連線中斷時的優雅降級 |
| `pkg/futuapi` | 便利的再匯出 —— `NewClient()`、`NewClientFromEnv()` |
| `pkg/health` | 用於 OpenD 存活/就緒探針的健康檢查 |
| `pkg/history` | 自動分頁的歷史 K 線下載 |
| `pkg/market` | 市場時段、交易日曆、交易時段偵測 |
| `pkg/metrics` | 用戶端效能指標收集 |
| `pkg/option` | 選擇權鏈查詢、代碼解析、希臘字母輔助函式 |
| `pkg/pb/*` | 184 個 protobuf 型別（v10.10.7008） |
| `pkg/ratelimit` | API 限流（依 protoID 的權杖桶） |
| `pkg/retry` | 可設定的指數退避重試 |
| `pkg/trd/audit.go` | 交易稽核日誌 |
| `pkg/trd/validation.go` | 下單預先檢查 |
| `pkg/tracing` | 核心追蹤介面（Tracer、Span、預設無操作） |
| `pkg/tracing/otel` | 基於 OpenTelemetry 的追蹤轉接器（選擇啟用） |

## 常用 API

### 連線

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

### 行情資料

| 函式 | 說明 |
|---|---|
| `GetQuote(ctx, c, market, code)` | 即時報價 |
| `GetKLines(ctx, c, market, code, klType, num)` | 最新 K 線柱 |
| `GetOrderBook(ctx, c, market, code, num)` | 買賣盤深度 |
| `GetTicker(ctx, c, market, code, num)` | 逐筆成交 |
| `GetStaticInfo(ctx, c, market, code)` | 證券名稱、型別、每手股數 |
| `GetSecuritySnapshot(ctx, c, securities)` | 多檔證券的完整快照 |
| `GetCapitalFlow(ctx, c, market, code)` | 資金流向 |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | 歷史 K 線（自動分頁） |
| `RequestHistoryKLQuota(ctx, c)` | API 配額使用情形 |

### 交易

| 函式 | 說明 |
|---|---|
| `GetAccountList(ctx, c)` | 所有交易帳戶 |
| `UnlockTrading(ctx, c, pwdMD5)` | 解鎖交易 |
| `GetFunds(ctx, c, accID)` | 帳戶資金與購買力 |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | 下單 |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | 修改或取消訂單 |
| `GetOrderList(ctx, c, accID)` | 活動中訂單 |
| `GetPositionList(ctx, c, accID)` | 目前持倉與損益 |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | 歷史訂單 |
| `GetOrderFillList(ctx, c, accID)` | 訂單成交明細 |

### 訂閱

| 函式 | 說明 |
|---|---|
| `Subscribe(ctx, c, market, code, []SubType)` | 訂閱推送型別 |
| `Unsubscribe(ctx, c, market, code, []SubType)` | 取消訂閱 |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | 透過 channel 推送報價 |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | 透過 channel 推送單一 K 線 |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | 帶篩選的多 K 線推送 |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | 透過 channel 推送逐筆成交 |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | 透過 channel 推送訂單簿 |

## 建置與測試

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test -race ./... # Full suite with race detector
```

## 架構

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

所有通訊皆透過 TCP 上的 Protocol Buffers 進行。完整架構決策參見 [DESIGN.md](DESIGN.md)，測試中使用的 mock OpenD 伺服器參見 [internal/testutil/mock](internal/testutil/mock/)。

## 疑難排解

| 錯誤 | 可能原因 |
|-------|-------------|
| `connection refused` | OpenD 未執行。檢查 `FUTU_OPEND_ADDR`。 |
| 美股 `GetQuote` 無資料 | 美股市場必須先呼叫 `Subscribe`。港股則不需要。 |
| `The packet body SHA1 signature is incorrect`（非常舊的 OpenD） | 將 OpenD 升級至 v10.5+。SDK 使用 SHA1(ciphertext)，OpenD 可接受。 |
| `解析protobuf协议失败` | 請求主體中缺少必要的 C2S 欄位。 |
| `模拟交易不支持` | 模擬模式下無法使用該功能；請使用 `WithTradeEnv(TrdEnv_Real)`。 |

## 貢獻

1. Fork 本儲存庫。
2. 建立功能分支（`git checkout -b feat/my-change`）。
3. 確保所有現有測試通過：`go test -race ./...`
4. 為任何新功能新增測試。
5. 執行 `go vet ./...` 並修正所有警告。
6. 開啟提取要求。

版本歷史參見 [CHANGELOG.md](CHANGELOG.md)，藍圖參見 [ENHANCEMENT_PLAN.md](ENHANCEMENT_PLAN.md)。

## 另請參閱

- [CHANGELOG](CHANGELOG.md) —— 版本歷史與發布說明
- [Version Map](docs/VERSION_MAP.md) —— 每個 SDK 版本所搭載的 Futu OpenD 協定 / proto 數量 / `clientVer`
- [USAGE Guide](docs/USAGE.md) —— 詳細設定、環境與進階模式
- [DESIGN](DESIGN.md) —— 架構、設計決策、API 模式
- [ENHANCEMENT_PLAN](ENHANCEMENT_PLAN.md) —— 即將推出的功能與藍圖
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) —— 每個功能的可執行範例

## 授權條款

Apache License 2.0 —— 參見 [LICENSE](LICENSE)。

> **交易免責聲明**：交易金融工具具有重大風險。在使用真實資金之前，務必在模擬模式下充分測試。
