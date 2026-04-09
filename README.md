# futuapi4go

<p align="center">
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go">
  </a>
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  </a>
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Version-0.4.1-blue.svg" alt="Version">
  </a>
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Status-Production--Ready-brightgreen.svg" alt="Status">
  </a>
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Examples-29-brightgreen.svg" alt="Examples">
  </a>
  <a href="https://github.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Tests-20/20%20PASS-brightgreen.svg" alt="Tests">
  </a>
</p>

<p align="center">
  <strong>Futu OpenD API SDK for Go</strong><br>
  World-class Golang trading interface for quantitative traders
</p>

---

## Installation

```bash
go get github.com/shing1211/futuapi4go
```

### Requirements

| Component | Version |
|-----------|---------|
| Golang | 1.21+ |
| Futu OpenD | 10.2.6208+ |

---

## Quick Start

```go
import (
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
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
		fmt.Printf("%s: CurPrice=%.2f Open=%.2f High=%.2f Low=%.2f\n",
			bq.Security.GetCode(), bq.CurPrice, bq.OpenPrice,
			bq.HighPrice, bq.LowPrice)
	}
}

func ptrStr(s string) *string { return &s }
```

---

## Project Status

**All 20 example compile tests pass with live OpenD (simulated account)**

| Module | Status | Notes |
|---------------|----------------|--------------|
| Core Architecture | Complete | TCP connection, auto heartbeat, user info |
| Market Data | Complete | 37+ APIs, all examples pass |
| Trading | Complete | 16 APIs, all examples pass |
| Push Notifications | Complete | Real-time quotes and trading push |
| System APIs | Complete | Global state, verification |
| Config System | Complete | Options, timeouts, retries, logging |
| Tests | Complete | 20/20 examples + unit + integration all pass |

---

## Examples

**29 example programs**, all compile and pass tests:

### Market Data APIs (11)

| Example | APIs |
|------|------|
| [qot_get_basic_qot](cmd/examples/qot_get_basic_qot/) | GetBasicQot |
| [qot_get_kl](cmd/examples/qot_get_kl/) | GetKL (day/min/week) |
| [qot_get_order_book](cmd/examples/qot_get_order_book/) | GetOrderBook |
| [qot_get_ticker](cmd/examples/qot_get_ticker/) | GetTicker |
| [qot_get_rt](cmd/examples/qot_get_rt/) | GetRT |
| [qot_get_broker](cmd/examples/qot_get_broker/) | GetBroker |
| [qot_get_capital_flow](cmd/examples/qot_get_capital_flow/) | GetCapitalFlow |
| [qot_get_static_info](cmd/examples/qot_get_static_info/) | GetStaticInfo |
| [qot_get_trade_date](cmd/examples/qot_get_trade_date/) | GetTradeDate |
| [qot_subscribe](cmd/examples/qot_subscribe/) | Subscribe |
| [qot_stock_filter](cmd/examples/qot_stock_filter/) | StockFilter |

### Trading APIs (7)

| Example | APIs |
|------|------|
| [trd_get_acc_list](cmd/examples/trd_get_acc_list/) | GetAccList |
| [trd_get_funds](cmd/examples/trd_get_funds/) | GetFunds |
| [trd_get_position_list](cmd/examples/trd_get_position_list/) | GetPositionList |
| [trd_unlock_trade](cmd/examples/trd_unlock_trade/) | UnlockTrade |
| [trd_place_order](cmd/examples/trd_place_order/) | PlaceOrder |
| [trd_get_order_list](cmd/examples/trd_get_order_list/) | GetOrderList |
| [trd_modify_order](cmd/examples/trd_modify_order/) | ModifyOrder |

### Comprehensive (5)

| Example | Description |
|------|------|
| [01_market_data_basic](cmd/examples/01_market_data_basic/) | Basic market data API |
| [02_market_data_advanced](cmd/examples/02_market_data_advanced/) | Advanced market analysis |
| [03_trading_operations](cmd/examples/03_trading_operations/) | Complete trading workflow |
| [04_push_subscriptions](cmd/examples/04_push_subscriptions/) | Real-time push |
| [05_comprehensive_demo](cmd/examples/05_comprehensive_demo/) | Full feature showcase |

### System APIs (1)

| Example | APIs |
|------|------|
| [sys_get_global_state](cmd/examples/sys_get_global_state/) | GetGlobalState |

### Algo Trading (5)

| Example | Strategy |
|------|------|
| [algo_sma_crossover](cmd/examples/algo_sma_crossover/) | SMA Crossover |
| [algo_grid_trading](cmd/examples/algo_grid_trading/) | Grid Trading |
| [algo_market_making](cmd/examples/algo_market_making/) | Market Making |
| [algo_breakout_trading](cmd/examples/algo_breakout_trading/) | Breakout Trading |
| [algo_vwap_execution](cmd/examples/algo_vwap_execution/) | VWAP Execution |

---

## Features

### Market Data
- Real-time quotes, K-lines, order book, tick-by-tick trades, minute data
- Broker queue, sector info, capital flow, options data
- Warrants, watchlist, price alerts, stock screening

### Trading
- Account management, order placement/modification, order management, position queries
- Trade records, order fees, margin ratio, max trade quantity

### Push Notifications
- Real-time quotes, K-lines, order book, tick-by-tick trades, minute data, broker push
- Order status, fills, trade notifications, system notification push

### System
- Global state, user info, latency statistics, verification

---

## Project Structure

```
futuapi4go/
├── cmd/
│   ├── examples/          # 29 example programs
│   └── simulator/          # OpenD simulator
├── internal/client/        # Core client
├── pkg/
│   ├── qot/              # Market Data API (37 functions)
│   ├── trd/              # Trading API (16 functions)
│   ├── sys/              # System API
│   ├── push/             # Push parsing
│   └── pb/               # Protobuf generated code (74 packages)
├── api/proto/            # Protobuf definitions (74 files)
└── test/                 # Integration tests
```

---

## Documentation

| Document | Description |
|------|------|
| [USER_GUIDE.md](USER_GUIDE.md) | User guide |
| [CHANGELOG.md](CHANGELOG.md) | Changelog |
| [PRODUCTION_PLAN.md](PRODUCTION_PLAN.md) | Implementation plan |
| [STATUS.md](STATUS.md) | Production readiness status |

Detailed example docs: [cmd/examples/README.md](cmd/examples/README.md)

---

## License

MIT License
