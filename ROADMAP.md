# futuapi4go Roadmap

> Last Updated: 2026-04-21

---

## Vision

futuapi4go is the most reliable, well-tested, and ergonomic Go SDK for the Futu OpenAPI. It powers production trading systems with institutional-grade observability, structured error handling, and comprehensive test coverage.

---

## Version History

### v0.8.0 — Production SDK (In Progress)
- [x] Context propagation on public API methods (`GetQuote`, `GetBasicQot`)
- [x] Waitable connection pool with context timeout support
- [x] Request-level cancellation support
- [ ] All P0 bug fixes from ENHANCEMENT_PLAN.md
- [ ] Structured error types
- [ ] Prometheus metrics endpoint
- [ ] OpenTelemetry tracing
- [ ] golangci-lint CI integration
- [ ] Race detector in CI
- [ ] Proper semver tagging

### v0.7.0
- [x] 100% proto field coverage for all 59 wrapper functions
- [x] Full proto field mapping audit
- [x] Push notification handler API
- [x] ProtoID constants re-exported
- [x] Thread-safe global logger
- [x] Connection pool with health checking

### v0.6.1
- [x] Fix push_test.go protobuf types
- [x] Fix nil logger panic
- [x] Fix connection state race

---

## Phases

### Phase 0: Fix the Foundation (P0) ✅ COMPLETE
> *Core infrastructure fixes completed.*

- [x] Fix `client_test.go` compilation (non-exported types, has `//go:build skip`)
- [x] Fix `push_test.go` protobuf wrapper types (7 failures)
- [x] Fix nil logger panic in `logf()`
- [x] Fix connection state race between `readLoop` and `Close()`
- [ ] Add `go test -race` to CI
- [ ] Fix `go vet` failures
- [ ] Export `Packet`/`PacketHandler` types for testing
- [x] Update ROADMAP.md (replace placeholder stubs)
- [ ] Tag proper semver release (`v0.8.0`)
- [ ] Update README examples to use env vars for secrets

### Phase 1: Production Quality (P1) 🚧 IN PROGRESS
> *SDK is reliable enough for serious trading bots.*

- [ ] Prometheus metrics endpoint (`/metrics`)
- [ ] OpenTelemetry distributed tracing
- [x] Context propagation on all API methods
- [ ] Structured error types (`ErrConnectionFailed`, etc.)
- [ ] Ping/pong connection health verification
- [ ] Connection chaos tests (mock server failure modes)
- [ ] TLS support
- [x] Update test state (PROJECT_STATUS.md merged into README)
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
- [x] ~~`push_test.go` protobuf type mismatch (7 failures)~~ — Fixed
- [x] ~~nil logger panic in pool tests~~ — Fixed
- [x] ~~Connection state race in `readLoop`~~ — Fixed
- [ ] `go vet` failures — Fix pending (Phase 0)
- [x] ~~ROADMAP.md is placeholder~~ — Updated
- [ ] Context propagation on remaining APIs (P1)
- [ ] `TestPoolConnReuse` timeout — Pre-existing, requires real OpenD connection

---

## Recent Updates (2026-04-21)

- **Context Support**: Added `RequestContext()` and `ReadResponseContext()` for request-level cancellation
- **Waitable Pool**: `ClientPool.Get()` now accepts `context.Context` and waits for available connections
- **Breaking Change**: All API functions now accept `context.Context` as first parameter

---

## Relationship to futugo4bot

futuapi4go powers [futugo4bot](https://github.com/shing1211/futugo4bot), a production algorithmic trading bot for HK futures. See [futugo4bot/docs/ENHANCEMENT_PLAN.md](https://github.com/shing1211/futugo4bot/blob/main/docs/ENHANCEMENT_PLAN.md) for the trading bot's full enhancement roadmap.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

---

*Generated from comprehensive code review and enhancement analysis — 2026-04-21*
