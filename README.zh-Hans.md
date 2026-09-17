# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/License-Apache%202.0-green?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/futuapi4go-v0.19.1-00ADD8?style=flat-square" alt="Version">
  <img src="https://img.shields.io/badge/Futu%20Proto-v10.10.7008-blue?style=flat-square" alt="Futu Proto Version">
  <a href="https://shing1211.github.io/futuapi4go/"><img src="https://img.shields.io/badge/Docs-GitHub%20Pages-97CAFF?style=flat-square&logo=github" alt="Docs"></a>
</p>

> **⚠️ 正在积极开发中**  
> 本 SDK 正在积极开发。虽然可以对接真实的 Futu OpenD 实例正常运行，
> 但部分 proto 响应字段可能尚未映射。API 和类型可能在小版本之间发生变化。
> 在依赖任何字段之前，请针对你的具体用例审计 `client/types.go` 并对照
> [Futu Proto Reference](https://openapi.futunn.com/mds/Futu-API-Doc-zh-Proto.md)。

> **Go 原生。类型安全。可用于生产。** 面向 [Futu OpenAPI](https://www.futunn.com/en/overview) 的最完整、最易用的 Go SDK —— 行情数据、交易与实时推送。所有通信均通过 TCP 上的 Protocol Buffers 进行。

[English](./README.md) · [简体中文](./README.zh-Hans.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Español](./README.es.md)

> 本文件是英文 [README](./README.md) 的社区翻译。**英文版本为准。**
> 同步于 / Last synced: c417534

- 184 个 protobuf 类型，覆盖所有 Futu OpenAPI 服务
- 一行代码连接，自动读取环境变量配置（`NewClientFromEnv`）
- 通过 channel 或类型化回调实现实时推送
- 流畅 API：`cli.Quote().GetBasicQot()`、`cli.Trade().PlaceOrder()`
- 基于 OpenTelemetry 的分布式追踪与指标（通过 `pkg/tracing/otel` 选择启用）
- 连接状态机、优雅关闭与自动重连
- 限流器、熔断器与重试已接入每一次 API 调用
- K 线数据缓存（LRU + TTL）、下单预检查、审计日志
- 推送标签即自动创建 GitHub 发布，说明取自 CHANGELOG

## 目录

- [安装](#安装)
- [快速开始](#快速开始)
- [核心特性](#核心特性)
- [示例](#示例)
- [包结构映射](#包结构映射)
- [常用 API](#常用-api)
- [构建与测试](#构建与测试)
- [架构](#架构)
- [故障排查](#故障排查)
- [贡献](#贡献)
- [许可证](#许可证)

## 安装

```bash
go get github.com/shing1211/futuapi4go@v0.19.1
```

需要 Go 1.26+，以及一个正在运行的 [Futu OpenD](https://www.futunn.com/en/overview) 实例。

## 快速开始

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

> **注意：** 美股需要先订阅，`GetQuote` 才能工作。港股则不需要。

## 核心特性

### 实时推送

停止轮询 —— 数据到达时即刻接收。两种投递模式：

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

### 行情数据

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

## 示例

涵盖所有 API 表面的完整可运行示例 —— 包括实时推送、交易工作流、历史数据与策略模式：

**[futuapi4go-demo →](https://github.com/shing1211/futuapi4go-demo)**

## 包结构映射

| 包 | 用途 |
|---------|---------|
| `client` | 高层封装 —— 推荐的入口点 |
| `pkg/qot` | 行情数据：报价、K 线、订单簿、逐笔数据…… |
| `pkg/trd` | 交易：订单、持仓、资金、历史…… |
| `pkg/sys` | 系统：全局状态、用户信息 |
| `pkg/push` | 推送通知解析器 |
| `pkg/push/chan` | 基于 channel 的实时推送投递 |
| `pkg/breaker` | 熔断器模式 |
| `pkg/cache` | K 线数据缓存（LRU + TTL） |
| `pkg/logger` | 结构化分级日志 |
| `pkg/util` | 代码解析（`ParseCode`、`FormatCode`）、市场辅助函数 |
| `pkg/constant` | 带 `String()` 方法的类型化常量 |
| `pkg/degradation` | 连接丢失时的优雅降级 |
| `pkg/futuapi` | 便捷再导出 —— `NewClient()`、`NewClientFromEnv()` |
| `pkg/health` | 用于 OpenD 存活/就绪探针的健康检查 |
| `pkg/history` | 自动分页的历史 K 线下载 |
| `pkg/market` | 市场时段、交易日历、会话检测 |
| `pkg/metrics` | 客户端性能指标采集 |
| `pkg/option` | 期权链查询、代码解析、希腊字母辅助函数 |
| `pkg/pb/*` | 184 个 protobuf 类型（v10.10.7008） |
| `pkg/ratelimit` | API 限流（按 protoID 的令牌桶） |
| `pkg/retry` | 带指数退避的可配置重试 |
| `pkg/trd/audit.go` | 交易审计日志 |
| `pkg/trd/validation.go` | 下单预检查 |
| `pkg/tracing` | 核心追踪接口（Tracer、Span、默认空实现） |
| `pkg/tracing/otel` | 基于 OpenTelemetry 的追踪适配器（选择启用） |

## 常用 API

### 连接

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

### 行情数据

| 函数 | 描述 |
|---|---|
| `GetQuote(ctx, c, market, code)` | 实时报价 |
| `GetKLines(ctx, c, market, code, klType, num)` | 最新 K 线柱 |
| `GetOrderBook(ctx, c, market, code, num)` | 买卖盘深度 |
| `GetTicker(ctx, c, market, code, num)` | 逐笔成交 |
| `GetStaticInfo(ctx, c, market, code)` | 证券名称、类型、每手股数 |
| `GetSecuritySnapshot(ctx, c, securities)` | 多只证券的完整快照 |
| `GetCapitalFlow(ctx, c, market, code)` | 资金流向 |
| `RequestHistoryKL(ctx, c, market, code, klType, start, end)` | 历史 K 线（自动分页） |
| `RequestHistoryKLQuota(ctx, c)` | API 配额使用情况 |

### 交易

| 函数 | 描述 |
|---|---|
| `GetAccountList(ctx, c)` | 所有交易账户 |
| `UnlockTrading(ctx, c, pwdMD5)` | 解锁交易 |
| `GetFunds(ctx, c, accID)` | 账户资金与购买力 |
| `PlaceOrder(ctx, c, accID, market, code, side, orderType, price, qty)` | 下单 |
| `ModifyOrder(ctx, c, accID, market, orderID, op, price, qty)` | 修改或撤销订单 |
| `GetOrderList(ctx, c, accID)` | 活动订单 |
| `GetPositionList(ctx, c, accID)` | 当前持仓及盈亏 |
| `GetHistoryOrderList(ctx, c, accID, market, start, end)` | 历史订单 |
| `GetOrderFillList(ctx, c, accID)` | 订单成交明细 |

### 订阅

| 函数 | 描述 |
|---|---|
| `Subscribe(ctx, c, market, code, []SubType)` | 订阅推送类型 |
| `Unsubscribe(ctx, c, market, code, []SubType)` | 取消订阅 |
| `chanpkg.SubscribeQuote(ctx, cli, market, code, ch)` | 通过 channel 推送报价 |
| `chanpkg.SubscribeKLine(ctx, cli, market, code, klType, ch)` | 通过 channel 推送单条 K 线 |
| `chanpkg.SubscribeKLines(ctx, cli, market, code, []klTypes, ch)` | 带过滤的多 K 线推送 |
| `chanpkg.SubscribeTicker(ctx, cli, market, code, ch)` | 通过 channel 推送逐笔成交 |
| `chanpkg.SubscribeOrderBook(ctx, cli, market, code, ch)` | 通过 channel 推送订单簿 |

## 构建与测试

```bash
go build ./...      # Compile everything
go vet ./...        # Lint
go test -race ./... # Full suite with race detector
```

## 架构

```
Application
  └── client/Client         (public wrappers)
       └── pkg/*            (qot, trd, sys — business logic)
            └── internal/client/Client   (connection, reconnect)
                 └── internal/client/Conn  (TCP I/O, packet framing)
                      └── Futu OpenD (TCP socket)
```

所有通信均通过 TCP 上的 Protocol Buffers 进行。完整架构决策参见 [DESIGN.md](DESIGN.md)，测试中使用的 mock OpenD 服务器参见 [internal/testutil/mock](internal/testutil/mock/)。

## 故障排查

| 错误 | 可能原因 |
|-------|-------------|
| `connection refused` | OpenD 未运行。检查 `FUTU_OPEND_ADDR`。 |
| 美股 `GetQuote` 无数据 | 美股市场必须先调用 `Subscribe`。港股则不需要。 |
| `The packet body SHA1 signature is incorrect`（非常旧的 OpenD） | 将 OpenD 升级到 v10.5+。SDK 使用 SHA1(ciphertext)，OpenD 可接受。 |
| `解析protobuf协议失败` | 请求体中缺少必需的 C2S 字段。 |
| `模拟交易不支持` | 模拟模式下该功能不可用；请使用 `WithTradeEnv(TrdEnv_Real)`。 |

## 贡献

1. Fork 本仓库。
2. 创建功能分支（`git checkout -b feat/my-change`）。
3. 确保所有现有测试通过：`go test -race ./...`
4. 为任何新功能添加测试。
5. 运行 `go vet ./...` 并修复所有警告。
6. 提交拉取请求。

版本历史参见 [CHANGELOG.md](CHANGELOG.md)，路线图参见 [ENHANCEMENT_PLAN.md](docs/IMPLEMENTATION_COMPLETE.md)。

## 参见

- [CHANGELOG](CHANGELOG.md) —— 版本历史与发布说明
- [Version Map](docs/VERSION_MAP.md) —— 每个 SDK 版本所携带的 Futu OpenD 协议 / proto 数量 / `clientVer`
- [USAGE Guide](docs/USAGE.md) —— 详细设置、环境与高级模式
- [DESIGN](DESIGN.md) —— 架构、设计决策、API 模式
- [ENHANCEMENT_PLAN](docs/IMPLEMENTATION_COMPLETE.md) —— 即将推出的功能与路线图
- [futuapi4go-demo](https://github.com/shing1211/futuapi4go-demo) —— 每个特性的可运行示例

## 许可证

Apache License 2.0 —— 参见 [LICENSE](LICENSE)。

> **交易免责声明**：交易金融工具存在重大风险。在使用真实资金之前，务必在模拟模式下充分测试。
