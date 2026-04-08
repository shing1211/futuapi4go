# futuapi4go

<p align="center">
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Version-0.4.1-blue.svg" alt="Version">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Status-Production--Ready-brightgreen.svg" alt="Status">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Examples-29-brightgreen.svg" alt="Examples">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Tests-20/20%20PASS-brightgreen.svg" alt="Tests">
  </a>
</p>

<p align="center">
  <strong>Go 语言实现的富途 OpenD API SDK</strong><br>
  为量化交易者打造的世界级 Golang 交易接口
</p>

---

## 安装 / Installation

```bash
go get gitee.com/shing1211/futuapi4go
```

### 环境要求 / Requirements

| Component | Version |
|-----------|---------|
| Golang | 1.21+ |
| Futu OpenD | 10.2.6208+ |

---

## 快速开始 / Quick Start

```go
import (
	"fmt"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	if err := cli.Connect("127.0.0.1:11111"); err != nil {
		panic(err)
	}

	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: ptrStr("00700")},
	}

	result, err := qot.GetBasicQot(cli, &qot.GetBasicQotRequest{
		SecurityList: securities,
	})
	if err != nil {
		panic(err)
	}

	for _, bq := range result {
		fmt.Printf("%s: 现价=%.2f 开盘=%.2f 最高=%.2f 最低=%.2f\n",
			bq.Security.GetCode(), bq.CurPrice, bq.OpenPrice,
			bq.HighPrice, bq.LowPrice)
	}
}

func ptrStr(s string) *string { return &s }
```

---

## 项目状态 / Project Status

**✅ All 20 example compile tests pass with live OpenD (simulated account)**

| 模块 / Module | 状态 / Status | 说明 / Notes |
|---------------|----------------|--------------|
| 核心架构 / Core | ✅ 完成 | TCP连接、自动心跳、用户信息 |
| 市场数据 / Market Data | ✅ 完成 | 37+ APIs, 所有示例通过 |
| 交易接口 / Trading | ✅ 完成 | 16 APIs, 所有示例通过 |
| 推送通知 / Push | ✅ 完成 | 实时行情与交易推送 |
| 系统 API / System | ✅ 完成 | 全局状态、验证接口 |
| 配置系统 / Config | ✅ 完成 | 功能选项、超时、重试、日志 |
| 测试 / Tests | ✅ 完成 | 20/20 examples + unit + integration 全部通过 |

---

## 示例 / Examples

**29 个示例程序**，全部通过编译测试 / All 29 example programs compile and pass tests:

### 行情 API / Market Data APIs (11)

| 示例 | APIs |
|------|------|
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

### 交易 API / Trading APIs (7)

| 示例 | APIs |
|------|------|
| [trd_get_acc_list](cmd/examples/trd_get_acc_list/) | GetAccList |
| [trd_get_funds](cmd/examples/trd_get_funds/) | GetFunds |
| [trd_get_position_list](cmd/examples/trd_get_position_list/) | GetPositionList |
| [trd_unlock_trade](cmd/examples/trd_unlock_trade/) | UnlockTrade |
| [trd_place_order](cmd/examples/trd_place_order/) | PlaceOrder |
| [trd_get_order_list](cmd/examples/trd_get_order_list/) | GetOrderList |
| [trd_modify_order](cmd/examples/trd_modify_order/) | ModifyOrder |

### 综合示例 / Comprehensive (5)

| 示例 | 说明 |
|------|------|
| [01_market_data_basic](cmd/examples/01_market_data_basic/) | 基础行情 API |
| [02_market_data_advanced](cmd/examples/02_market_data_advanced/) | 高级行情分析 |
| [03_trading_operations](cmd/examples/03_trading_operations/) | 完整交易流程 |
| [04_push_subscriptions](cmd/examples/04_push_subscriptions/) | 实时推送 |
| [05_comprehensive_demo](cmd/examples/05_comprehensive_demo/) | 全功能展示 |

### 系统 API / System (1)

| 示例 | APIs |
|------|------|
| [sys_get_global_state](cmd/examples/sys_get_global_state/) | GetGlobalState |

### 算法交易 / Algo Trading (5)

| 示例 | 策略 |
|------|------|
| [algo_sma_crossover](cmd/examples/algo_sma_crossover/) | SMA 交叉 |
| [algo_grid_trading](cmd/examples/algo_grid_trading/) | 网格交易 |
| [algo_market_making](cmd/examples/algo_market_making/) | 做市策略 |
| [algo_breakout_trading](cmd/examples/algo_breakout_trading/) | 突破交易 |
| [algo_vwap_execution](cmd/examples/algo_vwap_execution/) | VWAP 执行 |

---

## 功能特性 / Features

### 市场数据 / Market Data
- 实时行情、K线、订单簿、逐笔成交、分时数据
- 经纪队列、板块信息、资金流向、期权数据
- 涡轮窝轮、自选股管理、价格提醒、股票筛选

### 交易功能 / Trading
- 账户管理、下单改单、订单管理、持仓查询
- 成交记录、订单费用、保证金率、最大交易量

### 推送服务 / Push Notifications
- 实时行情、K线、订单簿、逐笔成交、分时、经纪推送
- 订单状态、成交、交易通知、系统通知推送

### 系统功能 / System
- 全局状态、用户信息、延迟统计、验证接口

---

## 项目结构 / Project Structure

```
futuapi4go/
├── cmd/
│   ├── examples/          # 29 个示例程序
│   └── simulator/          # OpenD 模拟器
├── internal/client/        # 核心客户端
├── pkg/
│   ├── qot/              # 行情 API (37 functions)
│   ├── trd/              # 交易 API (16 functions)
│   ├── sys/              # 系统 API
│   ├── push/             # 推送解析
│   └── pb/               # Protobuf 生成代码 (74 packages)
├── api/proto/            # Protobuf 定义 (74 files)
└── test/                 # 集成测试
```

---

## 文档 / Documentation

| 文档 | 说明 |
|------|------|
| [USER_GUIDE.md](USER_GUIDE.md) | 用户使用指南 |
| [CHANGELOG.md](CHANGELOG.md) | 更新日志 |
| [PRODUCTION_PLAN.md](PRODUCTION_PLAN.md) | 实施计划 |
| [STATUS.md](STATUS.md) | 生产就绪状态 |

详细示例文档: [cmd/examples/README.md](cmd/examples/README.md)

---

## 许可证 / License

MIT License
