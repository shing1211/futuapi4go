# FutuAPI4Go SDK - Advanced Enhancement Plan

> **Version**: v0.5.x | **Date**: 2026-05-01 | **Status**: ACTIVE

---

## Overview

The `futuapi4go` SDK has reached **~99% API coverage** vs the Python SDK with all core trading and market data operations implemented. This document outlines the roadmap for transforming the SDK from a **Futu API wrapper** into a **complete professional trading SDK**.

**Current state**: v0.5.2 — Core APIs complete, quick wins implemented.
**Target state**: World-class professional trading SDK with execution algorithms, risk engine, and event-driven framework.

---

## Phase A: Execution Algorithms (Priority P1)

### A-1: Execution Algorithm Framework

Implement TWAP/VWAP/IS algorithms as first-class SDK components.

| Item | Description | File | Status |
|------|-------------|------|--------|
| A-1a | TWAP algorithm with configurable time slices and participation rates | `pkg/execution/twap.go` | ⚪ Pending |
| A-1b | VWAP algorithm with volume profile scheduling | `pkg/execution/vwap.go` | ⚪ Pending |
| A-1c | IS (Implementation Shortfall) algorithm | `pkg/execution/is.go` | ⚪ Pending |
| A-1d | Algorithm config struct with schedule, participation rate, risk limits | `pkg/execution/config.go` | ⚪ Pending |

**Algorithm Interface:**
```go
type AlgoExecutor interface {
    Start(ctx context.Context, order *OrderRequest) error
    Pause()
    Resume()
    Cancel() error
    Status() (*AlgoStatus, error)
}

// AlgoStatus reports progress, remaining quantity, avg price, pnl
type AlgoStatus struct {
    AlgoID       string
    State        AlgoState // Running, Paused, Completed, Cancelled
    FilledQty    float64
    RemainingQty float64
    AvgFillPrice float64
    StartTime    time.Time
    EndTime      time.Time
}
```

**IS Algorithm Formula:**
```
IS Cost = slippage + market_impact
slippage = (execution_price - arrival_price) * fill_qty
market_impact = alpha * sigma * sqrt(fill_rate / ADV)
```

---

### A-2: Smart Order Router (SOR)

Aggregate orders and route to optimal venues based on real-time spread, liquidity, and cross-market arbitrage.

| Item | Description | File | Status |
|------|-------------|------|--------|
| A-2a | Venue scoring engine with spread/liquidity weighting | `pkg/router/scorer.go` | ⚪ Pending |
| A-2b | Multi-venue order splitter (HK, US, CN, SG, JP, AU) | `pkg/router/splitter.go` | ⚪ Pending |
| A-2c | Cross-market arbitrage detector | `pkg/router/arbitrage.go` | ⚪ Pending |
| A-2d | SOR configuration with venue priority and fallback | `pkg/router/config.go` | ⚪ Pending |

**SOR Interface:**
```go
type SmartOrderRouter struct {
    venues    []Venue
    scorer    VenueScorer
    fallback  Venue
}

func (r *SmartOrderRouter) Route(ctx context.Context, order *OrderRequest) (*RoutedOrder, error)
```

---

## Phase B: Real-Time Risk Engine (Priority P1)

### B-1: Risk Calculation Engine

Pre-trade and post-trade risk calculations with VaR, Greeks, and margin monitoring.

| Item | Description | File | Status |
|------|-------------|------|--------|
| B-1a | Historical VaR / CVaR calculation (parametric + Monte Carlo) | `pkg/risk/var.go` | ⚪ Pending |
| B-1b | Greeks calculator for options portfolios (Delta, Gamma, Theta, Vega) | `pkg/risk/greeks.go` | ⚪ Pending |
| B-1c | Portfolio-level margin requirement calculator | `pkg/risk/margin.go` | ⚪ Pending |
| B-1d | Margin call early warning system | `pkg/risk/margin_watch.go` | ⚪ Pending |

**Risk Interface:**
```go
type RiskCalculator interface {
    CalculateVaR(positions []Position, confidence float64, horizon time.Duration) (float64, error)
    CalculateCVaR(positions []Position, confidence float64) (float64, error)
    CheckMarginCall(portfolio *Portfolio) (bool, float64, error) // breached, shortfall
    Greeks(position *OptionPosition) (*Greeks, error)
}

type RiskReport struct {
    VaR             float64
    CVaR            float64
    PortfolioBeta   float64
    NetDelta        float64
    NetGamma        float64
    NetVega         float64
    MarginUsed      float64
    MarginAvailable float64
    MarginCall      bool
}
```

---

### B-2: Pre-Trade Risk Checks

Integrated into execution flow before order submission.

| Item | Description | File | Status |
|------|-------------|------|--------|
| B-2a | Per-order risk limits (max position, max loss, max exposure) | `pkg/risk/precheck.go` | ⚪ Pending |
| B-2b | Daily loss limit circuit breaker | `pkg/risk/daily_loss.go` | ⚪ Pending |
| B-2c | Position limit enforcement per symbol/sector | `pkg/risk/position_limit.go` | ⚪ Pending |
| B-2d | Order size validator (lot size, tick price, max qty) | `pkg/risk/order_validate.go` | ⚪ Pending |

---

## Phase C: Event-Driven Framework (Priority P1)

### C-1: Strategy Plug-in Architecture

Signal → Strategy → Risk → Execute → Report pipeline with plug-in strategies.

| Item | Description | File | Status |
|------|-------------|------|--------|
| C-1a | Signal subscription system with market data connectors | `pkg/framework/signal.go` | ⚪ Pending |
| C-1b | Strategy interface with OnBar/OnTick callbacks | `pkg/framework/strategy.go` | ⚪ Pending |
| C-1c | Strategy registry with lifecycle management (Active/Paused/Shutdown) | `pkg/framework/registry.go` | ⚪ Pending |
| C-1d | Trade outcome callback hooks for strategy feedback | `pkg/framework/callbacks.go` | ⚪ Pending |

**Strategy Interface:**
```go
type Strategy interface {
    OnBar(ctx context.Context, bar *Bar) *Signal
    OnTick(ctx context.Context, tick *Tick) *Signal
    Name() string
    Init(ctx context.Context, cfg StrategyConfig) error
    Close() error
}

type Signal struct {
    Symbol    string
    Side      Side // Buy, Sell, Hold
    Price     float64
    Quantity  float64
    Reason    string
    Confidence float64
    FactorScores map[string]float64
}
```

---

### C-2: Backtesting Connector

Strategy replay engine using historical data for validation.

| Item | Description | File | Status |
|------|-------------|------|--------|
| C-2a | Historical data iterator with warmup support | `pkg/backtest/data_iter.go` | ⚪ Pending |
| C-2b | Strategy replay runner with P&L attribution | `pkg/backtest/replay.go` | ⚪ Pending |
| C-2c | Performance metrics (Sharpe, max drawdown, win rate) | `pkg/backtest/metrics.go` | ⚪ Pending |
| C-2d | Walk-forward optimization runner | `pkg/backtest/wfo.go` | ⚪ Pending |

---

### C-3: Alert & Notification System

Built-in alert system with multiple notification channels.

| Item | Description | File | Status |
|------|-------------|------|--------|
| C-3a | Alert condition definitions (price, P&L, risk) | `pkg/alert/condition.go` | ⚪ Pending |
| C-3b | Webhook, Telegram, Discord, DingTalk integrations | `pkg/alert/channels.go` | ⚪ Pending |
| C-3c | Alert scheduler with cooldown and suppression | `pkg/alert/scheduler.go` | ⚪ Pending |

---

## Phase D: Advanced Data Features (Priority P2)

### D-1: WebSocket Streaming API

High-level streaming interfaces with auto-reconnection.

| Item | Description | File | Status |
|------|-------------|------|--------|
| D-1a | Streaming quote client with delta updates | `pkg/stream/quote.go` | ⚪ Pending |
| D-1b | Order book depth streaming with aggregation | `pkg/stream/orderbook.go` | ⚪ Pending |
| D-1c | Multi-symbol ticker aggregation | `pkg/stream/ticker.go` | ⚪ Pending |
| D-1d | Portfolio value streaming | `pkg/stream/portfolio.go` | ⚪ Pending |

**Streaming Interface:**
```go
type QuoteStream interface {
    Subscribe(ctx context.Context, securities []Security) error
    Unsubscribe(ctx context.Context, securities []Security) error
    C() <-chan *QuoteUpdate // channel-based delivery
    Close() error
}
```

---

### D-2: Historical Data Pipeline

Built-in historical data management for backtesting and analysis.

| Item | Description | File | Status |
|------|-------------|------|--------|
| D-2a | K-line downloader with auto-pagination | `pkg/data/kline_downloader.go` | ⚪ Pending |
| D-2b | Tick data aggregator with time binning | `pkg/data/tick_agg.go` | ⚪ Pending |
| D-2c | Data quality validator (gap detection, outlier detection) | `pkg/data/validator.go` | ⚪ Pending |
| D-2d | Bar iterator with warmup for strategy initialization | `pkg/data/bar_iter.go` | ⚪ Pending |

---

### D-3: Options Trading Suite

Professional options analytics built into SDK.

| Item | Description | File | Status |
|------|-------------|------|--------|
| D-3a | Black-Scholes pricing model | `pkg/options/bs.go` | ⚪ Pending |
| D-3b | Greeks calculator (Delta, Gamma, Theta, Vega, Rho) | `pkg/options/greeks.go` | ⚪ Pending |
| D-3c | Implied volatility surface builder | `pkg/options/iv_surface.go` | ⚪ Pending |
| D-3d | Strategy builder (straddles, spreads, butterflies, condors) | `pkg/options/strategies.go` | ⚪ Pending |
| D-3e | Risk graph generator (profit/loss diagrams) | `pkg/options/risk_graph.go` | ⚪ Pending |

**Options Strategy Builder:**
```go
type OptionStrategy struct {
    Type    StrategyType // Straddle, Strangle, Butterfly, IronCondor
    Legs    []OptionLeg  // each leg: code, expiry, strike, side, ratio
}

func BuildStraddle(symbol string, expiry time.Time, strike float64) *OptionStrategy
func BuildButterfly(symbol string, expiry time.Time, strikes []float64) *OptionStrategy
func (s *OptionStrategy) MaxProfit() float64
func (s *OptionStrategy) MaxLoss() float64
func (s *OptionStrategy) Breakeven() []float64
```

---

## Phase E: Portfolio & Multi-Account (Priority P2)

### E-1: Multi-Account Manager

Sub-account allocation and unified management for family offices and prop shops.

| Item | Description | File | Status |
|------|-------------|------|--------|
| E-1a | Sub-account allocation engine (equal, risk-parity, custom weights) | `pkg/multi/allocator.go` | ⚪ Pending |
| E-1b | P&L attribution per sub-account | `pkg/multi/attribution.go` | ⚪ Pending |
| E-1c | Risk isolation between accounts | `pkg/multi/isolation.go` | ⚪ Pending |
| E-1d | Unified account interface for managing 10+ accounts | `pkg/multi/manager.go` | ⚪ Pending |

---

### E-2: Portfolio Optimization Engine

Portfolio optimization algorithms for weight allocation.

| Item | Description | File | Status |
|------|-------------|------|--------|
| E-2a | Mean-variance optimization (Markowitz) | `pkg/portfolio/mean_variance.go` | ⚪ Pending |
| E-2b | Risk parity weighting | `pkg/portfolio/risk_parity.go` | ⚪ Pending |
| E-2c | Black-Litterman model | `pkg/portfolio/black_litterman.go` | ⚪ Pending |
| E-2d | Factor risk model integration | `pkg/portfolio/factor_risk.go` | ⚪ Pending |

---

## Phase F: Persistence & Storage (Priority P3)

### F-1: Data Persistence Layer

Standard storage adapters for common use cases.

| Item | Description | File | Status |
|------|-------------|------|--------|
| F-1a | SQLite adapter for order history and positions | `pkg/store/sqlite.go` | ⚪ Pending |
| F-1b | Redis adapter for warmup data and caching | `pkg/store/redis.go` | ⚪ Pending |
| F-1c | Strategy state serialization | `pkg/store/serializer.go` | ⚪ Pending |

---

## Quick Wins (Already Implemented)

| Item | Status | Evidence |
|------|--------|----------|
| GetAccList futures support | ✅ Done | `client/client.go:3439` — `TradeAPI.GetAccList(ctx, trdCat, isAll)` supports `TrdCategory_Future` (value 2) |
| GetMaxTrdQtys batch | ✅ Done | `client/client.go:1001` — `GetAccTradingInfo` returns batch max quantities |
| GetOrderList filter by status | ✅ Done | `GetOrderListRequest.FilterStatusList` + `FilterConditions` for full filtering |
| SubscribeSymbols batch | ✅ Done | `client/client.go:354-380` — batch subscribe wrapper |
| UnsubscribeSymbols batch | ✅ Done | `client/client.go:382-407` — batch unsubscribe wrapper |
| Market hours IsMarketOpen | ✅ Done | `pkg/market/hours.go:155` — `IsOpen(m, t)` + `IsHKOpen/IsUSOpen/IsCNOpen` |

---

## Progress Dashboard

| Phase | Name | Items | Priority | Status |
|-------|------|-------|----------|--------|
| **A** | Execution Algorithms | 8 | P1 | ⚪ Pending |
| **B** | Real-Time Risk Engine | 7 | P1 | ⚪ Pending |
| **C** | Event-Driven Framework | 11 | P1 | ⚪ Pending |
| **D** | Advanced Data Features | 11 | P2 | ⚪ Pending |
| **E** | Portfolio & Multi-Account | 8 | P2 | ⚪ Pending |
| **F** | Persistence & Storage | 3 | P3 | ⚪ Pending |
| **Total** | | **48** | | |

---

## Implementation Notes

### How to Use This Plan

1. Work in phase order (A → F)
2. Within a phase, prioritize by severity
3. Update status: `⚪ Pending` → `🔄 In Progress` → `✅ Done`
4. All new packages should have:
   - GoDoc comments
   - Unit tests with table-driven tests
   - Example usage in `docs/`
5. Breaking changes require migration guide update

### Version Targets

| Version | Phase | Target |
|---------|-------|--------|
| v0.6.0 | A (Execution Algorithms) | TWAP/VWAP/IS as first-class SDK citizens |
| v0.7.0 | B (Risk Engine) | VaR, Greeks, margin monitoring |
| v0.8.0 | C (Event Framework) | Strategy plug-in, backtesting |
| v0.9.0 | D-F | Streaming, options, portfolio, persistence |

---

## Files to Create/Modify

```
pkg/execution/       — NEW: twap.go, vwap.go, is.go, config.go
pkg/router/          — NEW: scorer.go, splitter.go, arbitrage.go, config.go
pkg/risk/            — NEW: var.go, greeks.go, margin.go, margin_watch.go, precheck.go, daily_loss.go
pkg/framework/       — NEW: signal.go, strategy.go, registry.go, callbacks.go
pkg/backtest/        — NEW: data_iter.go, replay.go, metrics.go, wfo.go
pkg/alert/           — NEW: condition.go, channels.go, scheduler.go
pkg/stream/          — NEW: quote.go, orderbook.go, ticker.go, portfolio.go
pkg/data/            — NEW: kline_downloader.go, tick_agg.go, validator.go, bar_iter.go
pkg/options/         — NEW: bs.go, greeks.go, iv_surface.go, strategies.go, risk_graph.go
pkg/multi/           — NEW: allocator.go, attribution.go, isolation.go, manager.go
pkg/portfolio/       — NEW: mean_variance.go, risk_parity.go, black_litterman.go, factor_risk.go
pkg/store/           — NEW: sqlite.go, redis.go, serializer.go
```

---

*Last updated: 2026-05-01*