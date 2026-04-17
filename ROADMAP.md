# futuapi4go Roadmap

> Last Updated: 2026-04-18

---

## Vision

futuapi4go is the most reliable, well-tested, and ergonomic Go SDK for the Futu OpenAPI. It powers production trading systems with institutional-grade observability, structured error handling, and comprehensive test coverage.

---

## Version History

### v0.7.0 — Production SDK (Planned)
- [ ] All P0 bug fixes from ENHANCEMENT_PLAN.md
- [ ] Context propagation on all public API methods
- [ ] Structured error types
- [ ] Prometheus metrics endpoint
- [ ] OpenTelemetry tracing
- [ ] golangci-lint CI integration
- [ ] Race detector in CI
- [ ] Proper semver tagging

### v0.6.0 (Current)
- [x] 100% proto field coverage for all 59 wrapper functions
- [x] Full proto field mapping audit
- [x] Push notification handler API
- [x] ProtoID constants re-exported
- [x] Thread-safe global logger
- [x] Connection pool with health checking

### v0.5.0
- [x] Complete trading API coverage
- [x] Order management and position tracking
- [x] Historical order and fill queries

### v0.4.0
- [x] CancelAllOrder support
- [x] RegQotPush support
- [x] Comprehensive test suites

### v0.3.0
- [x] Market data APIs
- [x] Subscription system
- [x] Push notifications

---

## Phases

### Phase 0: Fix the Foundation (P0)
> *Before any v1.0 release. All items are must-fix.*

- [ ] Fix `client_test.go` compilation (non-exported types, has `//go:build skip`)
- [x] ~~Fix `push_test.go` protobuf wrapper types (7 failures)~~ — Fixed `b6435b4`/`a8c0828`
- [x] ~~Fix nil logger panic in `logf()`~~ — Fixed `b6435b4` (eager `log.Default()`)
- [x] ~~Fix connection state race between `readLoop` and `Close()`~~ — Fixed `b6435b4` (`connected int32` atomic)
- [ ] Add `go test -race` to CI
- [ ] Fix `go vet` failures
- [ ] Export `Packet`/`PacketHandler` types for testing
- [x] ~~Update ROADMAP.md (replace placeholder stubs)~~ — Updated 2026-04-18
- [ ] Tag proper semver release (`v0.7.0`)
- [ ] Update README examples to use env vars for secrets

### Phase 1: Production Quality (P1)
> *SDK is reliable enough for serious trading bots.*

- [ ] Prometheus metrics endpoint (`/metrics`)
- [ ] OpenTelemetry distributed tracing
- [ ] Context propagation on all API methods
- [ ] Structured error types (`ErrConnectionFailed`, etc.)
- [ ] Ping/pong connection health verification
- [ ] Connection chaos tests (mock server failure modes)
- [ ] TLS support
- [ ] Update PROJECT_STATUS.md (actual test state)
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

- [ ] `client_test.go` compilation failure — has `//go:build skip`, pending redesign
- [x] ~~`push_test.go` protobuf type mismatch (7 failures)~~ — Fixed `b6435b4`/`a8c0828`
- [x] ~~nil logger panic in pool tests~~ — Fixed `b6435b4`
- [x] ~~Connection state race in `readLoop`~~ — Fixed `b6435b4`
- [ ] `go vet` failures — Fix pending (Phase 0)
- [x] ~~ROADMAP.md is placeholder~~ — Updated 2026-04-18
- [x] ~~No Prometheus metrics endpoint~~ — Planned for Phase 1
- [x] ~~No OpenTelemetry tracing~~ — Planned for Phase 1
- [x] ~~No context propagation on all APIs~~ — Planned for Phase 1
- [ ] `TestPoolConnReuse` timeout — Pre-existing, requires real OpenD connection
- [ ] `test/qot_api`/`test/trd_api`/`test/util` mock server failures — Pre-existing network issues

---

## Relationship to futugo4bot

futuapi4go powers [futugo4bot](https://github.com/shing1211/futugo4bot), a production algorithmic trading bot for HK futures. See [futugo4bot/docs/ENHANCEMENT_PLAN.md](https://github.com/shing1211/futugo4bot/blob/main/docs/ENHANCEMENT_PLAN.md) for the trading bot's full enhancement roadmap.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

---

*Generated from comprehensive code review and enhancement analysis — 2026-04-15*
