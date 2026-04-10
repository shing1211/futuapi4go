# Examples

This directory contains comprehensive examples for every API function in FutuAPI4Go.

## Example Index

### Quick Start
| Example | API | Description |
|---|---|---|
| `debug_test/` | Multiple | Debug test for protocol verification |

### Market Data APIs

| # | Example | API | Description |
|---|---|---|---|
| 1 | `qot_get_basic_qot/` | GetBasicQot | Real-time quotes |
| 2 | `qot_get_kl/` | GetKL | K-line data (daily, minute, weekly) |
| 3 | `qot_get_order_book/` | GetOrderBook | Order book |
| 4 | `qot_get_ticker/` | GetTicker | Tick-by-tick trades |
| 5 | `qot_get_rt/` | GetRT | Real-time minute data |
| 6 | `qot_get_broker/` | GetBroker | Broker queue |
| 7 | `qot_subscribe/` | Subscribe | Real-time subscription |
| 8 | `qot_get_capital_flow/` | GetCapitalFlow | Capital flow analysis |
| 9 | `qot_stock_filter/` | StockFilter | Stock screening |
| 10 | `qot_get_trade_date/` | GetTradeDate | Trading calendar |
| 11 | `qot_get_static_info/` | GetStaticInfo | Stock static info |

### Trading APIs

| # | Example | API | Description |
|---|---|---|---|
| 1 | `trd_get_acc_list/` | GetAccList | Account list |
| 2 | `trd_get_funds/` | GetFunds | Account funds |
| 3 | `trd_get_position_list/` | GetPositionList | Holdings |
| 4 | `trd_unlock_trade/` | UnlockTrade | Unlock trading |
| 5 | `trd_place_order/` | PlaceOrder | Place order (with dry run) |
| 6 | `trd_get_order_list/` | GetOrderList | Order list |
| 7 | `trd_modify_order/` | ModifyOrder | Modify order |

### System APIs

| # | Example | API | Description |
|---|---|---|---|
| 1 | `sys_get_global_state/` | GetGlobalState | OpenD global state |

## How to Use

### Prerequisites
1. **Futu OpenD must be running**
2. **OpenD must be logged in to quote server**
3. **Default address**: `127.0.0.1:11111`

### Basic Usage

```bash
# Run any example
cd qot_get_basic_qot
go run main.go

# Use custom OpenD address
FUTU_ADDR=192.168.1.100:11111 go run main.go

# Trading examples need account ID
go run main.go 123456789  # Replace with your account ID
```

### Example Output

```
=== GetBasicQot Example ===
Connecting to 127.0.0.1:11111...
Connected! ConnID=7447516338564503397

Querying quotes for 3 securities...

=== Real-time Quotes ===

00700 (Tencent)
   Current Price:  350.50
   Open Price:   348.00
   High Price:   352.00
   Low Price:    347.50
   Volume:       12345678
   ...
```

## Features

- **Production Ready**
   - Proper error handling
   - Clear usage instructions
   - Formatted output with emojis

- **Safe Defaults**
   - Dry run mode for trading examples
   - No accidental order submissions
   - Clear warnings for destructive operations

## Running All Examples

```bash
# Market data examples
for dir in qot_*/; do
  echo "Running $dir..."
  cd "$dir" && go run main.go && cd ..
done

# Trading examples
for dir in trd_*/; do
  echo "Running $dir..."
  cd "$dir" && go run main.go && cd ..
done

# System examples
cd sys_get_global_state && go run main.go
```

## Notes

1. **Trading APIs**:
   - Examples use DRY RUN mode by default
   - Remove `dryRun = true` to execute real trades
   - Always test with simulator first

3. **Authentication**:
   - OpenD must be logged in to quote server
   - Trade unlock requires correct password

## Documentation

- [USER_GUIDE.md](../USER_GUIDE.md) - Complete user guide
- [SIMULATOR.md](../SIMULATOR.md) - OpenD simulator guide

---

**Happy Coding!**
