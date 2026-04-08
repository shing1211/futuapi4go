# Examples / 示例程式

This directory contains comprehensive examples for every API function in FutuAPI4Go.
本目錄包含 FutuAPI4Go 每個 API 函數的綜合示例。

## 📚 Example Index / 示例索引

### 🎯 Quick Start / 快速開始
| Example / 示例 | API | Description / 描述 |
|---|---|---|
| `debug_test/` | Multiple | Debug test for protocol verification / 協議驗證調試測試 |

### 📊 Market Data APIs / 行情 API

| # | Example / 示例 | API | Description / 描述 |
|---|---|---|---|
| 1 | `qot_get_basic_qot/` | GetBasicQot | Real-time quotes / 實時報價 |
| 2 | `qot_get_kl/` | GetKL | K-line data (daily, minute, weekly) / K線數據 |
| 3 | `qot_get_order_book/` | GetOrderBook | Order book / 買賣盤 |
| 4 | `qot_get_ticker/` | GetTicker | Tick-by-tick trades / 逐筆成交 |
| 5 | `qot_get_rt/` | GetRT | Real-time minute data / 實時分時 |
| 6 | `qot_get_broker/` | GetBroker | Broker queue / 經紀隊列 |
| 7 | `qot_subscribe/` | Subscribe | Real-time subscription / 實時訂閱 |
| 8 | `qot_get_capital_flow/` | GetCapitalFlow | Capital flow analysis / 資金流向 |
| 9 | `qot_stock_filter/` | StockFilter | Stock screening / 股票篩選 |
| 10 | `qot_get_trade_date/` | GetTradeDate | Trading calendar / 交易日曆 |
| 11 | `qot_get_static_info/` | GetStaticInfo | Stock static info / 股票靜態信息 |

### 💰 Trading APIs / 交易 API

| # | Example / 示例 | API | Description / 描述 |
|---|---|---|---|
| 1 | `trd_get_acc_list/` | GetAccList | Account list / 賬戶列表 |
| 2 | `trd_get_funds/` | GetFunds | Account funds / 賬戶資金 |
| 3 | `trd_get_position_list/` | GetPositionList | Holdings / 持倉列表 |
| 4 | `trd_unlock_trade/` | UnlockTrade | Unlock trading / 解鎖交易 |
| 5 | `trd_place_order/` | PlaceOrder | Place order (with dry run) / 下單 |
| 6 | `trd_get_order_list/` | GetOrderList | Order list / 訂單列表 |
| 7 | `trd_modify_order/` | ModifyOrder | Modify order / 修改訂單 |

### 🖥️ System APIs / 系統 API

| # | Example / 示例 | API | Description / 描述 |
|---|---|---|---|
| 1 | `sys_get_global_state/` | GetGlobalState | OpenD global state / OpenD 全局狀態 |

## 🚀 How to Use / 使用方法

### Prerequisites / 前置條件
1. **Futu OpenD must be running** / Futu OpenD 必須運行中
2. **OpenD must be logged in to quote server** / OpenD 必須登錄行情服務器
3. **Default address**: `127.0.0.1:11111`

### Basic Usage / 基本用法

```bash
# Run any example / 運行任何示例
cd cmd/examples/qot_get_basic_qot
go run main.go

# Use custom OpenD address / 使用自定義 OpenD 地址
FUTU_ADDR=192.168.1.100:11111 go run main.go

# Trading examples need account ID / 交易示例需要賬戶ID
go run main.go 123456789  # Replace with your account ID / 替換為您的賬戶ID
```

### Example Output / 示例輸出

```
=== GetBasicQot Example / 獲取實時行情示例 ===
Connecting to 127.0.0.1:11111...
✅ Connected! ConnID=7447516338564503397

Querying quotes for 3 securities...

=== Real-time Quotes / 實時行情 ===

📈 00700 (Tencent)
   Current Price / 現價:  350.50
   Open Price / 開盤價:   348.00
   High Price / 最高價:   352.00
   Low Price / 最低價:    347.50
   Volume / 成交量:       12345678
   ...
```

## 📝 Features / 特性

✅ **Bilingual Documentation** / 雙語文檔
   - All comments in English + Traditional Chinese
   - All 註釋使用英文 + 繁體中文

✅ **Production Ready** / 生產就緒
   - Proper error handling
   - Clear usage instructions
   - Formatted output with emojis

✅ **Safe Defaults** / 安全預設
   - Dry run mode for trading examples
   - No accidental order submissions
   - Clear warnings for destructive operations

## 🛠️ Running All Examples / 運行所有示例

```bash
# Market data examples / 行情示例
for dir in qot_*/; do
  echo "Running $dir..."
  cd "$dir" && go run main.go && cd ..
done

# Trading examples / 交易示例
for dir in trd_*/; do
  echo "Running $dir..."
  cd "$dir" && go run main.go && cd ..
done

# System examples / 系統示例
cd sys_get_global_state && go run main.go
```

## ⚠️ Notes / 注意事項

1. **Protocol Issue / 協議問題**:
   - Some Qot APIs may return parse errors due to proto mismatch
   - See [PROTO_NOTES.md](../PROTO_NOTES.md) for details
   - Subscribe and basic connection APIs work correctly

2. **Trading APIs / 交易API**:
   - Examples use DRY RUN mode by default
   - Remove `dryRun = true` to execute real trades
   - Always test with simulator first

3. **Authentication / 認證**:
   - OpenD must be logged in to quote server
   - Trade unlock requires correct password

## 📖 Documentation / 文檔

- [USER_GUIDE.md](../USER_GUIDE.md) - Complete user guide / 完整用戶指南
- [SIMULATOR.md](../SIMULATOR.md) - OpenD simulator guide / OpenD 模擬器指南
- [PROTO_NOTES.md](../PROTO_NOTES.md) - Protocol implementation notes / 協議實現筆記

---

**Happy Coding! 🚀**
