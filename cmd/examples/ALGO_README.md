# Algorithmic Trading Examples / 算法交易示例

This directory contains algorithmic trading strategy examples using FutuAPI4Go.
本目錄包含使用 FutuAPI4Go 的算法交易策略示例。

## ⚠️ IMPORTANT SAFETY NOTICE / 重要安全提示

**All examples run in DRY RUN mode by default!**
**所有示例預設以模擬運行模式執行！**

- ✅ No real orders are placed unless you explicitly set `dryRun = false`
- ✅ 除非您明確設置 `dryRun = false`，否則不會下達真實訂單
- ✅ Always test thoroughly before enabling live trading
- ✅ 在啟用真實交易前務必充分測試

---

## 📚 Strategy Index / 策略索引

### 🎯 Trend Following / 趨勢跟隨

| # | Example / 示例 | Strategy / 策略 | Risk Management / 風險管理 |
|---|---|---|---|
| 1 | [`algo_sma_crossover/`](algo_sma_crossover/) | SMA Crossover / SMA交叉 | None (signal only) / 無（僅信號） |
| 2 | [`algo_breakout_trading/`](algo_breakout_trading/) | Breakout / 突破交易 | Stop-Loss & Take-Profit / 止損止盈 |

### 📊 Market Neutral / 市場中性

| # | Example / 示例 | Strategy / 策略 | Risk Management / 風險管理 |
|---|---|---|---|
| 3 | [`algo_grid_trading/`](algo_grid_trading/) | Grid Trading / 網格交易 | Multiple price levels / 多價格級別 |
| 4 | [`algo_market_making/`](algo_market_making/) | Market Making / 做市 | Bid-Ask spread / 買賣價差 |

### 📈 Execution Algorithms / 執行算法

| # | Example / 示例 | Strategy / 策略 | Risk Management / 風險管理 |
|---|---|---|---|
| 5 | [`algo_vwap_execution/`](algo_vwap_execution/) | VWAP Execution / VWAP執行 | VWAP benchmark / VWAP基準 |

---

## 🚀 Quick Start / 快速開始

### 1. SMA Crossover Strategy / SMA交叉策略

A classic trend-following strategy using Simple Moving Average crossovers.
使用簡單移動平均線交叉的經典趨勢跟隨策略。

**How it works / 工作原理:**
- When short-term SMA crosses above long-term SMA → **BUY signal** / **買入信號**
- When short-term SMA crosses below long-term SMA → **SELL signal** / **賣出信號**

```bash
cd algo_sma_crossover
go run main.go [account_id]
```

**Best for / 適用於:**
- Trending markets / 趨勢市場
- Medium to long-term trading / 中長期交易
- Beginners learning algo trading / 初學者學習算法交易

---

### 2. Grid Trading Strategy / 網格交易策略

Places buy and sell orders at predetermined price levels to profit from oscillations.
在預設價格級別下買入和賣出訂單，從波動中獲利。

**How it works / 工作原理:**
- Divide price range into N grids / 將價格範圍分為N個網格
- Place buy orders below current price / 在當前價格下方下買入單
- Place sell orders above current price / 在當前價格上方下賣出單
- Profit from price oscillations / 從價格波動中獲利

```bash
cd algo_grid_trading
go run main.go [account_id]
```

**Best for / 適用於:**
- Ranging/consolidating markets / 區間/盤整市場
- High volatility environments / 高波動環境
- Systematic trading / 系統化交易

---

### 3. Market Making Strategy / 做市策略

Simultaneously places bid and ask orders to profit from the spread.
同時下買入和賣出訂單，從價差中獲利。

**How it works / 工作原理:**
- Place buy order slightly below current price / 在當前價格略下方下買入單
- Place sell order slightly above current price / 在當前價格略上方下賣出單
- Profit from bid-ask spread / 從買賣價差中獲利
- Multi-layer orders for deeper liquidity / 多層訂單以獲取更深流動性

```bash
cd algo_market_making
go run main.go [account_id]
```

**Best for / 適用於:**
- High liquidity markets / 高流動性市場
- Tight spread environments / 窄價差環境
- Professional traders / 專業交易者

---

### 4. Breakout Trading Strategy / 突破交易策略

Detects price breakouts from consolidation ranges and enters positions.
檢測價格從盤整區間的突破並建立頭寸。

**How it works / 工作原理:**
- Calculate support/resistance over N periods / 計算N週期的支撐/阻力
- When price breaks above resistance → BUY / 價格突破阻力→買入
- When price breaks below support → SELL / 價格跌破支撐→賣出
- Automatic stop-loss and take-profit / 自動止損止盈

```bash
cd algo_breakout_trading
go run main.go [account_id]
```

**Best for / 適用於:**
- Volatile markets with clear ranges / 有明確區間的波動市場
- Momentum trading / 動能交易
- Short to medium-term trading / 短中期交易

---

### 5. VWAP Execution Strategy / VWAP執行策略

Intelligent order execution using Volume-Weighted Average Price benchmark.
使用成交量加權平均價格基準的智能訂單執行。

**How it works / 工作原理:**
- Calculate VWAP from historical data / 從歷史數據計算VWAP
- Split large order into smaller child orders / 將大單分為小單
- Execute when price is favorable vs VWAP / 當價格相對VWAP有利時執行
- Track execution quality / 追蹤執行質量

```bash
cd algo_vwap_execution
go run main.go [account_id] [total_shares]
```

**Best for / 適用於:**
- Large order execution / 大單執行
- Institutional trading / 機構交易
- Minimizing market impact / 最小化市場影響

---

## ⚙️ Configuration / 配置

### Environment Variables / 環境變數

```bash
# Set OpenD address / 設置OpenD地址
export FUTU_ADDR=127.0.0.1:11111

# Run strategy / 運行策略
go run main.go [account_id]
```

### Strategy Parameters / 策略參數

Each strategy has configurable parameters at the top of the file:
每個策略在文件頂部都有可配置的參數：

```go
// Example: SMA Crossover parameters / SMA交叉參數
const (
    ShortPeriod = 5  // Short-term SMA / 短期SMA
    LongPeriod  = 20 // Long-term SMA / 長期SMA
    TradeQty    = 100 // Shares per trade / 每次交易股數
)
```

---

## 🛡️ Risk Management / 風險管理

### Dry Run Mode / 模擬運行模式

**ALL strategies default to dry run mode!**
**所有策略預設為模擬運行模式！**

To enable live trading, you must:
要啟用真實交易，您必須：

1. Open the strategy file / 打開策略文件
2. Find `dryRun := true` / 找到 `dryRun := true`
3. Change to `dryRun := false` / 改為 `dryRun := false`
4. **TEST THOROUGHLY FIRST** / **先充分測試**

### Stop-Loss & Take-Profit / 止損止盈

Some strategies include automatic risk management:
某些策略包含自動風險管理：

- **Breakout Strategy**: Configurable stop-loss and take-profit percentages
- **突破策略**: 可配置的止損止盈百分比
- **VWAP Execution**: Benchmark-based execution quality tracking
- **VWAP執行**: 基於基準的執行質量追蹤

---

## 📊 Performance Monitoring / 性能監控

### Metrics to Track / 追蹤指標

- **Win Rate** / 勝率: % of profitable trades
- **Average Profit/Loss** / 平均盈虧: Per trade average
- **Maximum Drawdown** / 最大回撤: Largest peak-to-trough decline
- **Sharpe Ratio** / 夏普比率: Risk-adjusted return
- **VWAP Slippage** / VWAP滑點: Execution quality vs benchmark

### Logging / 日誌

All strategies output detailed execution logs:
所有策略都輸出詳細的執行日誌：

```
📊 Execution Summary / 執行摘要:
   Total Filled / 總成交:     1000/1000 shares (100.0%)
   Average Price / 平均價:  HK$349.50
   VWAP / VWAP:            HK$350.00
   ✅ Execution Savings / 執行節省: HK$500.00 (0.14% better than VWAP)
```

---

## ⚠️ Warnings / 警告

1. **Past performance does not guarantee future results**
   **過往表現不保證未來結果**

2. **Always start with DRY RUN mode**
   **務必從模擬運行模式開始**

3. **Test with small amounts first**
   **先以小金額測試**

4. **Monitor positions continuously**
   **持續監控頭寸**

5. **Set stop-losses for all trades**
   **為所有交易設置止損**

---

## 📖 Additional Resources / 額外資源

- [USER_GUIDE.md](../../USER_GUIDE.md) - Complete user guide / 完整用戶指南
- [SIMULATOR.md](../../SIMULATOR.md) - Test with simulator / 使用模擬器測試
- [PROTO_NOTES.md](../../PROTO_NOTES.md) - Protocol notes / 協議說明

---

## 💬 Support / 支持

For questions or issues:
如有問題或困難：

1. Check the strategy documentation / 查看策略文檔
2. Review the logs / 查看日誌
3. Test with simulator first / 先用模擬器測試
4. Open an issue on Gitee / 在Gitee上提交issue

---

**Happy Algorithmic Trading! 🚀**
**祝您算法交易愉快！**
