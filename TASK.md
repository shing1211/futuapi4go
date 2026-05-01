# futuapi4go Task List

> **Last Updated:** 2026-05-02 | **Version:** v0.5.4

---

## Completed Tasks

### v0.5.4 (Current Release)
- [x] Batch subscription wrappers (SubscribeSymbols/UnsubscribeSymbols)
- [x] Fluent API (client.Quote(), client.Trade(), client.System())
- [x] GetHistoryKLPoints wrapper
- [x] GetUsedQuota wrapper

### v0.5.3
- [x] Enhanced error handling with full message support
- [x] Additional type safety across all APIs
- [x] Performance optimizations

### v0.5.2
- [x] GetHistoryKLPoints wrapper (ProtoID 3106)
- [x] GetUsedQuota wrapper (ProtoID 1010)
- [x] NoDataMode typed enum
- [x] DataStatus typed enum
- [x] Fluent API accessors

### v0.5.1
- [x] Enhanced Error System with FutuError
  - Category, Recovery fields
  - FullMessage() method
  - CodeString() method
  - Is() for errors.Is() compatibility
- [x] Error predicates
  - IsServerError(), IsAPIError()
  - IsConnectionError(), IsTimeoutError(), IsTradingError()
- [x] Error bridge functions
- [x] Circuit breaker on Client via SetBreaker()/GetBreaker()
- [x] Input validation standalones
- [x] OrderBuilder auto-detect market
- [x] Convenience wrappers (GetTodayFills, GetTodayOrders)
- [x] HistoryKLine iterator methods
- [x] Typed enum Int32() method
- [x] TrdMarket.Prefix()
- [x] Client.GetReconnectCount()
- [x] TLS support (WithTLS)
- [x] Rate limiter (ProtoLimiter)
- [x] Retry with backoff
- [x] Health checker
- [x] Performance benchmarks

### v0.5.0
- [x] Context-required API
- [x] Typed Market Constants
- [x] Graceful shutdown helpers

### v0.3.x (Feature Parity)
- [x] All market data APIs (FR2.1-2.16)
- [x] All trading APIs (FR3.1-3.17)
- [x] All system APIs (FR4.1-4.4)
- [x] All push parsers (FR5.3-5.10)
- [x] Channel-based push (FR5.11)
- [x] Connection pool (NR2)
- [x] Buffered I/O (NR1)
- [x] Circuit breaker (NR6)
- [x] Structured logging (NR13)
- [x] Typed enums (NR17)

---

## Upcoming Tasks

### v0.5.5 (Next Patch)
- [ ] GetDelayStatistics fix (known proto2 issue)
- [ ] GetTradeDate compatibility fix
- [ ] Additional input validations
- [ ] More edge case tests

### v0.6.0 (Execution Algorithms)
- [ ] TWAP algorithm implementation
  - [ ] Configurable time slices
  - [ ] Participation rate control
  - [ ] Order splitting logic
- [ ] VWAP algorithm
  - [ ] Volume profile scheduling
  - [ ] Real-time adjustment
- [ ] IS (Implementation Shortfall)
  - [ ] Arrival price tracking
  - [ ] Market impact calculation
  - [ ] Slippage measurement

### v0.7.0 (Risk Engine)
- [ ] VaR/CVaR calculation
  - [ ] Parametric VaR
  - [ ] Monte Carlo simulation
- [ ] Greeks calculator
  - [ ] Delta, Gamma, Theta, Vega
  - [ ] Portfolio-level Greeks
- [ ] Margin requirements
  - [ ] Initial margin
  - [ ] Maintenance margin
- [ ] Margin call warning

### v0.8.0 (Event Framework)
- [ ] Signal subscription system
  - [ ] Market data connectors
  - [ ] Signal processor
- [ ] Strategy interface
  - [ ] OnBar callback
  - [ ] OnTick callback
  - [ ] Lifecycle management
- [ ] Strategy registry
- [ ] Backtesting connector
  - [ ] Historical data iterator
  - [ ] Strategy replay
  - [ ] Performance metrics

### v0.9.0 (Advanced Data)
- [ ] WebSocket streaming
  - [ ] Quote stream
  - [ ] Order book stream
  - [ ] Portfolio stream
- [ ] Historical data pipeline
  - [ ] K-line downloader
  - [ ] Tick aggregator
  - [ ] Data validator
- [ ] Options trading suite
  - [ ] Black-Scholes pricing
  - [ ] IV surface builder
  - [ ] Strategy builder

### v1.0.0 (Major)
- [ ] Full feature parity with Python SDK+
- [ ] Production-hardened
- [ ] Documented examples

---

## Documentation Tasks

### Immediate
- [ ] Update README with latest v0.5.4 examples
- [ ] Add more integration tests
- [ ] Performance benchmarking documentation
- [ ] Error handling guide

### Future
- [ ] Video tutorials
- [ ] Example strategies
- [ ] Best practices guide

---

## Bug Fixes Backlog

### Known Issues

| Issue | Status | Workaround |
|-------|--------|------------|
| GetDelayStatistics proto2 | Known | Call skipped gracefully |
| GetTradeDate compatibility | Known | May require specific OpenD |

### To Investigate

- [ ] Edge cases in push parsing
- [ ] Race conditions in pool
- [ ] Memory leaks in long-running connections

---

## Technical Debt

### High Priority
- [ ] Refactor test structure
- [ ] Add more integration tests
- [ ] CI/CD pipeline optimization

### Medium Priority
- [ ] Deprecation warnings for old APIs
- [ ] Example cleanup
- [ ] Benchmark regression suite

### Low Priority
- [ ] Comments cleanup
- [ ] TODO cleanup

---

## Contribution Guidelines

See [CONTRIBUTING.md](CONTRIBUTING.md)

### Quick Links
- [README.md](README.md) - Quick start
- [AGENTS.md](AGENTS.md) - Developer operations
- [DESIGN.md](DESIGN.md) - Architecture decisions
- [ENHANCEMENT_PLAN.md](ENHANCEMENT_PLAN.md) - Future roadmap