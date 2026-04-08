# futuapi4go

<p align="center">
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Version-0.4.0--dev-orange.svg" alt="Version">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Status-Development-yellow.svg" alt="Status">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Examples-29-brightgreen.svg" alt="Examples">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Progress-44%25-yellow.svg" alt="Progress">
  </a>
</p>

<p align="center">
  <strong>Go 语言实现的富途 OpenD API SDK</strong><br>
  为量化交易者打造的世界级 Golang 交易接口
</p>

---

## 🚀 Quick Start

### 基本用法 / Basic Usage

```go
import futuapi "gitee.com/shing1211/futuapi4go/internal/client"

// 简单连接
cli := futuapi.New()
defer cli.Close()
cli.Connect("127.0.0.1:11111")

// 配置连接
cli := futuapi.New(
    futuapi.WithDialTimeout(5*time.Second),
    futuapi.WithAPITimeout(10*time.Second),
    futuapi.WithMaxRetries(5),
    futuapi.WithReconnectBackoff(2.0),
)

// 连接池
pool := futuapi.NewClientPool(futuapi.DefaultPoolConfig("127.0.0.1:11111"))
defer pool.Close()
pool.StartHealthChecker()
```

---

## ⚠️ Production Status / 生產狀態

**Current Status**: ⚠️ **Development** - Not yet production-ready
**當前狀態**: ⚠️ **開發中** - 尚未達到生產就緒

See [STATUS.md](STATUS.md) for detailed readiness assessment.
查看 [STATUS.md](STATUS.md) 獲取詳細就緒評估。

See [PRODUCTION_PLAN.md](PRODUCTION_PLAN.md) for implementation roadmap.
查看 [PRODUCTION_PLAN.md](PRODUCTION_PLAN.md) 獲取實施路線圖。

---

## 项目状态

| 模块 | 状态 | 说明 |
|------|------|------|
| 核心架构 | ✅ 已完成 | TCP 连接、自动心跳、用户信息 |
| 市场数据 (Qot) | ✅ 已完成 | 37+ APIs 完整实现，安全连接检查 |
| 交易接口 (Trd) | ✅ 已完成 | 16 APIs 完整实现，安全连接检查 |
| 推送通知 | ✅ 已完成 | 实时行情与交易推送，串行匹配 |
| 系统 API | ✅ 已完成 | 全局状态、验证接口 |
| 配置系统 | ✅ 已完成 | 功能选项、超时、重试、日志、连接池 |
| 指标监控 | ✅ 已完成 | Metrics/Instrumentation、健康检查端点 |
| 版本信息 | ✅ 已完成 | GetVersionInfo 实现 |
| OpenD 模拟器 | 🔄 46% 完成 | 29/63 handlers 完整实现 |
| 算法交易示例 | ✅ 已完成 | 5 个策略示例 |
| 测试工具 | ✅ 已完成 | 64 tests passing across 5 packages |

## 已完成阶段 / Completed Phases

### ✅ Phase 1: Critical Bug Fixes (4/4)
- Nil-conn 防护：API 调用前不再 panic
- TOCTOU 竞态修复：重连接线程安全
- 调试日志清理：生产环境无冗余输出
- 日志一致性：统一使用 logf()

### ✅ Phase 2: API Safety Layer (6/6)
- `EnsureConnected()` 辅助方法
- 57 个 API 函数全部包装连接检查
- 串行响应匹配：防止推送通知被误消费
- Context 支持：`Context()`, `WithContext()`
- 推送处理器：`SetPushHandler()`

### ✅ Phase 3: Configuration System (5/5)
- `ClientOptions` 结构体，合理默认值
- 可配置超时：拨号、API、保活
- 重试配置：最大重试、间隔、指数退避
- 日志接口：自定义 logger，4 级日志
- 连接池：`ClientPool` 健康检查，自动重连

### ✅ Phase 4: Testing Infrastructure (4/4)
- 64 tests passing across 5 packages
- Unit tests for core client functionality
- Integration tests with OpenD simulator
- Concurrent access and race condition tests
- Test coverage for error paths and edge cases

### ✅ Phase 6: Push Notifications & Observability (11/11)
- Push notification support with serial matching
- Metrics and instrumentation for monitoring
- Health check endpoint for client pool
- Version information API (GetVersionInfo)
- Release checklist for production readiness
- GetOptionChain implementation (2304)
- GetOptionExpirationDate implementation (2305)
- Bug fixes and protocol improvements

---

## 功能特性

### 市场数据
- 实时行情 (GetBasicQot)
- K线数据 (GetKL, RequestHistoryKL)
- 订单簿 (GetOrderBook)
- 逐笔成交 (GetTicker)
- 分时数据 (GetRT)
- 经纪队列 (GetBroker)
- 板块信息 (GetPlateSet, GetPlateSecurity, GetOwnerPlate)
- 静态信息 (GetStaticInfo)
- 资金流向 (GetCapitalFlow, GetCapitalDistribution)
- 期权数据 (GetOptionChain, GetOptionExpirationDate)
- 涡轮窝轮 (GetWarrant)
- 自选股管理 (GetUserSecurity, ModifyUserSecurity, GetUserSecurityGroup)
- 价格提醒 (SetPriceReminder, GetPriceReminder)
- 股票筛选 (StockFilter)
- 期货信息 (GetFutureInfo)
- 代码变更 (GetCodeChange)
- IPO 列表 (GetIpoList)
- 持仓变化 (GetHoldingChangeList)
- 停牌信息 (GetSuspend)
- 历史 K 线配额 (RequestHistoryKLQuota)
- 复权信息 (RequestRehab)
- 订阅管理 (Subscribe, RegQotPush, GetSubInfo)

### 交易功能
- 账户管理 (GetAccList, UnlockTrade)
- 资金查询 (GetFunds)
- 下单与改单 (PlaceOrder, ModifyOrder)
- 订单管理 (GetOrderList, GetHistoryOrderList)
- 成交记录 (GetOrderFillList, GetHistoryOrderFillList)
- 持仓查询 (GetPositionList)
- 订单费用 (GetOrderFee)
- 保证金率 (GetMarginRatio)
- 最大交易量 (GetMaxTrdQtys)
- 账户推送 (SubAccPush)
- 订单确认 (ReconfirmOrder)
- 资金流动 (GetFlowSummary)

### 推送服务
- 实时行情推送 (Qot_UpdateBasicQot)
- K线推送 (Qot_UpdateKL)
- 订单簿推送 (Qot_UpdateOrderBook)
- 逐笔成交推送 (Qot_UpdateTicker)
- 分时推送 (Qot_UpdateRT)
- 经纪推送 (Qot_UpdateBroker)
- 价格提醒推送 (Qot_UpdatePriceReminder)
- 订单状态推送 (Trd_UpdateOrder)
- 成交推送 (Trd_UpdateOrderFill)
- 交易通知推送 (Trd_Notify)
- 系统通知推送 (Notify)

### 系统功能
- 全局状态 (GetGlobalState)
- 用户信息 (GetUserInfo)
- 延迟统计 (GetDelayStatistics)
- 验证接口 (Verification)

---

## 快速开始

### 安装

```bash
go get gitee.com/shing1211/futuapi4go
```

### 环境要求

| 组件 | 版本 |
|------|------|
| Golang | 1.21+ |
| Futu OpenD | 10.2.6208+ |

### 基础示例

```go
package main

import (
	"fmt"
	"log"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"gitee.com/shing1211/futuapi4go/qot"
	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
)

func main() {
	// 创建客户端
	cli := futuapi.New()

	// 连接 OpenD
	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 查询腾讯控股行情
	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: func() *string { s := "00700"; return &s }()},
	}

	result, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		log.Fatal(err)
	}

	for _, bq := range result {
		fmt.Printf("%s %s: 现价=%.2f 开盘=%.2f 最高=%.2f 最低=%.2f 成交量=%d\n",
			bq.Security.GetCode(), bq.Name, bq.CurPrice, bq.OpenPrice,
			bq.HighPrice, bq.LowPrice, bq.Volume)
	}
}
```

---

## 项目结构

```
futuapi4go/
├── cmd/                      # 应用程序
│   ├── examples/             # 示例程序 (29个)
│   │   ├── qot_*/            # 行情 API 示例 (11个)
│   │   ├── trd_*/            # 交易 API 示例 (7个)
│   │   ├── sys_*/            # 系统 API 示例 (1个)
│   │   └── algo_*/           # 算法交易策略 (5个)
│   └── simulator/            # OpenD 模拟器
├── internal/                 # 私有代码
│   └── client/               # 核心客户端
│       ├── conn.go           # TCP 连接与协议 (44字节头)
│       ├── client.go         # 主客户端与选项配置
│       ├── pool.go           # 连接池管理
│       └── errors.go         # 错误类型
├── pkg/                      # 公共库
│   ├── qot/                  # 行情 API (37 functions)
│   ├── trd/                  # 交易 API (16 functions)
│   ├── sys/                  # 系统 API (4 functions)
│   ├── push/                 # 推送解析 (11 handlers)
│   └── pb/                   # Protobuf 生成代码 (74 packages)
├── api/proto/                # Protobuf 定义 (74 files)
├── docs/                     # 文档
├── scripts/                  # 构建脚本
└── PRODUCTION_PLAN.md        # 实施计划
```

---

## 文档

| 文档 | 说明 |
|------|------|
| [PRODUCTION_PLAN.md](PRODUCTION_PLAN.md) | 综合实施计划与进度追踪 |
| [STATUS.md](STATUS.md) | 生产就绪状态仪表板 |
| [IMPLEMENTATION.md](IMPLEMENTATION.md) | 详细实现清单 |
| [USER_GUIDE.md](USER_GUIDE.md) | 用户使用指南 |
| [SIMULATOR_STATUS.md](SIMULATOR_STATUS.md) | 模拟器状态评估 |
| [CHANGELOG.md](CHANGELOG.md) | 更新日志 |

---

## 示例代码

本项目提供了 **29 个示例程序**，覆盖所有主要功能：

### 📊 行情 API 示例 (11个)
| 示例 | APIs 覆盖 |
|------|----------|
| [qot_get_basic_qot](cmd/examples/qot_get_basic_qot/) | GetBasicQot |
| [qot_get_kl](cmd/examples/qot_get_kl/) | GetKL (日/分/周) |
| [qot_get_order_book](cmd/examples/qot_get_order_book/) | GetOrderBook |
| [qot_get_ticker](cmd/examples/qot_get_ticker/) | GetTicker |
| [qot_get_rt](cmd/examples/qot_get_rt/) | GetRT |
| [qot_get_broker](cmd/examples/qot_get_broker/) | GetBroker |
| [qot_get_capital_flow](cmd/examples/qot_get_capital_flow/) | GetCapitalFlow |
| [qot_get_static_info](cmd/examples/qot_get_static_info/) | GetStaticInfo |
| [qot_get_trade_date](cmd/examples/qot_get_trade_date/) | GetTradeDate |
| [qot_subscribe](cmd/examples/qot_subscribe/) | Subscribe |
| [qot_stock_filter](cmd/examples/qot_stock_filter/) | StockFilter |

### 💰 交易 API 示例 (7个)
| 示例 | APIs 覆盖 |
|------|----------|
| [trd_get_acc_list](cmd/examples/trd_get_acc_list/) | GetAccList |
| [trd_get_funds](cmd/examples/trd_get_funds/) | GetFunds |
| [trd_get_position_list](cmd/examples/trd_get_position_list/) | GetPositionList |
| [trd_unlock_trade](cmd/examples/trd_unlock_trade/) | UnlockTrade |
| [trd_place_order](cmd/examples/trd_place_order/) | PlaceOrder |
| [trd_get_order_list](cmd/examples/trd_get_order_list/) | GetOrderList |
| [trd_modify_order](cmd/examples/trd_modify_order/) | ModifyOrder |

### 🖥️ 系统 API (1个)
| 示例 | APIs 覆盖 |
|------|----------|
| [sys_get_global_state](cmd/examples/sys_get_global_state/) | GetGlobalState |

### 🤖 算法交易策略 (5个)
| 示例 | 策略 |
|------|------|
| [algo_sma_crossover](cmd/examples/algo_sma_crossover/) | SMA 交叉策略 |
| [algo_grid_trading](cmd/examples/algo_grid_trading/) | 网格交易策略 |
| [algo_market_making](cmd/examples/algo_market_making/) | 做市策略 |
| [algo_breakout_trading](cmd/examples/algo_breakout_trading/) | 突破交易策略 |
| [algo_vwap_execution](cmd/examples/algo_vwap_execution/) | VWAP 执行策略 |

详细示例文档请查看: [cmd/examples/EXAMPLES_README.md](cmd/examples/EXAMPLES_README.md)
算法交易文档: [cmd/examples/ALGO_README.md](cmd/examples/ALGO_README.md)

---

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 致谢

- [富途](https://www.futunn.com/) - 提供 OpenAPI
