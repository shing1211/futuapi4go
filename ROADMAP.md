# futuapi4go Roadmap

> Last Updated: 2026-04-23

---

## Vision

futuapi4go is the most reliable, well-tested, and ergonomic Go SDK for the Futu OpenAPI. It powers production trading systems with institutional-grade observability, structured error handling, and comprehensive test coverage.

**Go SDK surpasses Python SDK** in proto coverage (78 vs ~50 protos), type safety, performance, and modern infrastructure (circuit breaker, structured logging, connection pool, channel-based push).

---

## Version History

### v0.9.0 — Feature Parity (Current)
- [x] All Python SDK APIs implemented (GetAccountInfo, GetFlowSummary, GetAccTradingInfo)
- [x] Full Funds struct with multi-currency cash and per-market assets
- [x] `pkg/breaker` — circuit breaker pattern
- [x] `pkg/logger` — structured leveled logging (text + JSON)
- [x] `pkg/push/chan` — channel-based push delivery
- [x] `pkg/util` — code parsing (ParseCode, FormatCode, market conversion)
- [x] `pkg/constant` — Python-style String() methods on all enums
- [x] Unit tests for `pkg/util`, `pkg/constant`, `pkg/logger`, `pkg/breaker`
- [x] All packages pass `go build` and `go vet`

### v0.8.0 — Production SDK
- [x] Context propagation on public API methods
- [x] Waitable connection pool with context timeout support
- [x] Request-level cancellation support
- [x] Push notification handler API
- [x] ProtoID constants re-exported
- [x] Thread-safe global logger
- [x] Connection pool with health checking
- [x] 100% proto field coverage for all wrapper functions

### v0.7.0
- [x] Full proto field mapping audit
- [x] Proto generation pipeline

---

## Feature Matrix: Go vs Python SDK

| Category | Python | Go | Status |
|----------|--------|----|--------|
| Protobufs | ~50 | 78 | **Go +28 extra** |
| Quote/Market APIs | ~48 | ~60 | ✅ Go >= Python |
| Trading APIs | ~18 | ~17 | ✅ Parity |
| System APIs | ~4 | ~4 | ✅ Tie |
| Push parsing | ~7 | ~11 | ✅ Go ahead |
| Channel push | ❌ | ✅ 7 functions | Go only |
| Circuit breaker | ❌ | ✅ | Go only |
| Structured logging | ❌ | ✅ text+JSON | Go only |
| Connection pool | ❌ | ✅ | Go only |
| Code helpers | ❌ | ✅ ParseCode/FormatCode | Go only |
| Unit tests | pandas DataFrames | struct types | ✅ Tie |

---

## Phases

### Phase 1: Production Quality (P1) 🚧 IN PROGRESS
> *SDK is reliable enough for serious trading bots.*

- [ ] Prometheus metrics endpoint (`/metrics`)
- [ ] OpenTelemetry distributed tracing
- [ ] Structured error types (`ErrConnectionFailed`, etc.)
- [ ] Ping/pong connection health verification
- [ ] Connection chaos tests (mock server failure modes)
- [ ] TLS support
- [ ] Example programs audit + index
- [ ] golangci-lint configuration

### Phase 2: Performance & Polish (P2)
> *Fine-tuning for high-frequency trading workloads.*

- [ ] WebSocket transport alternative
- [ ] Zero-allocation hot request path
- [ ] Connection pool O(1) lookup optimization
- [ ] Historical data download utility in SDK
- [ ] Option chain helper functions
- [ ] HK market hours utility
- [ ] Structured logging (slog)
- [ ] Benchmark regression CI
- [ ] Performance profiling guide
- [ ] Architecture Decision Records (ADRs)

### Phase 3: Ecosystem (P3)
> *Building a developer ecosystem.*

- [ ] Makefile with standard targets
- [ ] Fuzz testing pipeline
- [ ] Property-based testing
- [ ] GraphQL interface alternative
- [ ] Stability report
- [ ] Commit convention + commitlint

---

## Known Issues

### Proto Compatibility (serverVer=1003)
- [ ] `GetDelayStatistics` — proto2 wire-format incompatibility (OpenD rejects packed encoding for `repeated int32`)
- [ ] `GetTradeDate` — all C2S fields required; may also be proto2 issue

### Pending
- [ ] Fix `go vet` remaining warnings
- [ ] Context propagation on remaining APIs
- [ ] Tag proper semver release

---

## Recent Updates (2026-04-23)

- **Account Info**: Added `GetAccountInfo` (accinfo_query equivalent) with full multi-currency cash and per-market assets
- **Cash Flow**: Added `GetFlowSummary` (get_acc_cash_flow equivalent)
- **Trading Info**: Added `GetAccTradingInfo` (acctradinginfo_query equivalent)
- **New Packages**: `pkg/breaker`, `pkg/logger`, `pkg/util`, `pkg/push/chan`
- **Unit Tests**: All new packages have comprehensive test coverage

---

## Relationship to futugo4bot

futuapi4go powers [futugo4bot](https://github.com/shing1211/futu4bot), a production algorithmic trading bot for HK futures.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

---

*Last comprehensive review: 2026-04-23*
