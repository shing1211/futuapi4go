# Algorithmic Trading Examples

This directory contains algorithmic trading strategy examples using FutuAPI4Go.

## IMPORTANT SAFETY NOTICE

**All examples run in DRY RUN mode by default!**

- No real orders are placed unless you explicitly set `dryRun = false`
- Always test thoroughly before enabling live trading

---

## Strategy Index

### Trend Following

| # | Example | Strategy | Risk Management |
|---|---|---|---|
| 1 | [`algo_sma_crossover/`](algo_sma_crossover/) | SMA Crossover | None (signal only) |
| 2 | [`algo_breakout_trading/`](algo_breakout_trading/) | Breakout | Stop-Loss & Take-Profit |

### Market Neutral

| # | Example | Strategy | Risk Management |
|---|---|---|---|
| 3 | [`algo_grid_trading/`](algo_grid_trading/) | Grid Trading | Multiple price levels |
| 4 | [`algo_market_making/`](algo_market_making/) | Market Making | Bid-Ask spread |

### Execution Algorithms

| # | Example | Strategy | Risk Management |
|---|---|---|---|
| 5 | [`algo_vwap_execution/`](algo_vwap_execution/) | VWAP Execution | VWAP benchmark |

---

## Quick Start

### 1. SMA Crossover Strategy

A classic trend-following strategy using Simple Moving Average crossovers.

**How it works:**
- When short-term SMA crosses above long-term SMA → **BUY signal**
- When short-term SMA crosses below long-term SMA → **SELL signal**

```bash
cd algo_sma_crossover
go run main.go [account_id]
```

**Best for:**
- Trending markets
- Medium to long-term trading
- Beginners learning algo trading

---

### 2. Grid Trading Strategy

Places buy and sell orders at predetermined price levels to profit from oscillations.

**How it works:**
- Divide price range into N grids
- Place buy orders below current price
- Place sell orders above current price
- Profit from price oscillations

```bash
cd algo_grid_trading
go run main.go [account_id]
```

**Best for:**
- Ranging/consolidating markets
- High volatility environments
- Systematic trading

---

### 3. Market Making Strategy

Simultaneously places bid and ask orders to profit from the spread.

**How it works:**
- Place buy order slightly below current price
- Place sell order slightly above current price
- Profit from bid-ask spread
- Multi-layer orders for deeper liquidity

```bash
cd algo_market_making
go run main.go [account_id]
```

**Best for:**
- High liquidity markets
- Tight spread environments
- Professional traders

---

### 4. Breakout Trading Strategy

Detects price breakouts from consolidation ranges and enters positions.

**How it works:**
- Calculate support/resistance over N periods
- When price breaks above resistance → BUY
- When price breaks below support → SELL
- Automatic stop-loss and take-profit

```bash
cd algo_breakout_trading
go run main.go [account_id]
```

**Best for:**
- Volatile markets with clear ranges
- Momentum trading
- Short to medium-term trading

---

### 5. VWAP Execution Strategy

Intelligent order execution using Volume-Weighted Average Price benchmark.

**How it works:**
- Calculate VWAP from historical data
- Split large order into smaller child orders
- Execute when price is favorable vs VWAP
- Track execution quality

```bash
cd algo_vwap_execution
go run main.go [account_id] [total_shares]
```

**Best for:**
- Large order execution
- Institutional trading
- Minimizing market impact

---

## Configuration

### Environment Variables

```bash
# Set OpenD address
export FUTU_ADDR=127.0.0.1:11111

# Run strategy
go run main.go [account_id]
```

### Strategy Parameters

Each strategy has configurable parameters at the top of the file:

```go
// Example: SMA Crossover parameters
const (
    ShortPeriod = 5  // Short-term SMA
    LongPeriod  = 20 // Long-term SMA
    TradeQty    = 100 // Shares per trade
)
```

---

## Risk Management

### Dry Run Mode

**ALL strategies default to dry run mode!**

To enable live trading, you must:

1. Open the strategy file
2. Find `dryRun := true`
3. Change to `dryRun := false`
4. **TEST THOROUGHLY FIRST**

### Stop-Loss & Take-Profit

Some strategies include automatic risk management:

- **Breakout Strategy**: Configurable stop-loss and take-profit percentages
- **VWAP Execution**: Benchmark-based execution quality tracking

---

## Performance Monitoring

### Metrics to Track

- **Win Rate**: % of profitable trades
- **Average Profit/Loss**: Per trade average
- **Maximum Drawdown**: Largest peak-to-trough decline
- **Sharpe Ratio**: Risk-adjusted return
- **VWAP Slippage**: Execution quality vs benchmark

### Logging

All strategies output detailed execution logs:

```
Execution Summary:
   Total Filled:     1000/1000 shares (100.0%)
   Average Price:  HK$349.50
   VWAP:            HK$350.00
   Execution Savings: HK$500.00 (0.14% better than VWAP)
```

---

## Warnings

1. **Past performance does not guarantee future results**

2. **Always start with DRY RUN mode**

3. **Test with small amounts first**

4. **Monitor positions continuously**

5. **Set stop-losses for all trades**

---

## Additional Resources

- [docs/USER_GUIDE.md](../../docs/USER_GUIDE.md) - Complete user guide
- [docs/DEVELOPER.md](../../docs/DEVELOPER.md) - Developer guide
- [docs/API_REFERENCE.md](../../docs/API_REFERENCE.md) - API reference

---

## Support

For questions or issues:

1. Check the strategy documentation
2. Review the logs
3. Test with simulator first
4. Open an issue on github

---

**Happy Algorithmic Trading!**
