# futuapi4go Enhancement Plan

> **Purpose**: Comprehensive production-grade enhancement checklist to make futuapi4go a robust, institutional-quality SDK for Futu OpenD trading.
>
> **Status**: Core functionality complete. Gap to production-ready SDK: well-defined.
>
> **Dependency**: This SDK powers [futugo4bot](https://github.com/shing1211/futugo4bot). Many items below directly improve the trading bot's reliability.
>
> **Priority Key**: P0 = Must fix before production use | P1 = Major improvement | P2 = Important polish | P3 = Nice-to-have

---

## Table of Contents

1. [Critical Bug Fixes](#1-critical-bug-fixes)
2. [SDK Quality & Stability](#2-sdk-quality--stability)
3. [Performance & Scalability](#3-performance--scalability)
4. [API Design & Ergonomics](#4-api-design--ergonomics)
5. [Observability & Debugging](#5-observability--debugging)
6. [Testing & Reliability](#6-testing--reliability)
7. [Security](#7-security)
8. [Documentation](#8-documentation)
9. [Developer Experience](#9-developer-experience)
10. [Advanced Features](#10-advanced-features)
11. [Priority Roadmap](#11-priority-roadmap)

---

## 1. Critical Bug Fixes

> These must be resolved before the SDK is used in any live trading context.

| Priority | Issue | Location | Fix |
|----------|-------|---------|-----|
| **P0** | `client_test.go` references non-exported `Packet`/`PacketHandler` types. Test **does not compile**. | `client/client_test.go:62` | Redesign test to use `test/util/mock_server.go` approach. The `APIConnector` interface doesn't exist in the exported API. Use a mock TCP server instead. |
| **P0** | `push_test.go` marshals `S2C` protobuf directly instead of `Response` wrapper. **7 test failures**. | `pkg/push/push_test.go` | Wrap S2C in `Response{RetType: 0, S2C: ...}` before calling `proto.Marshal()`. All push notifications from Futu OpenD arrive as `Response` messages. |
| **P0** | `logf()` dereferences nil global logger. **Panic in pool tests**. | `internal/client/client.go:50` | Initialize logger eagerly in `New()` — don't rely on lazy `sync.Once`. Ensure `logger != nil` before any `logf()` call. |
| **P0** | Connection state race between `readLoop` and `Close()`. `readLoop` checks bool `c.connected` while `Close()` sets it without stopping the goroutine. | `internal/client/client.go:501` | Use `atomic.LoadInt32(&c.connActive) == 0` consistently in `readLoop` instead of bool `c.connected`. |
| **P1** | `ClientPool` uses `PoolType` parameter but `newClient()` ignores it — all pool types create identical connections. | `internal/client/pool.go:275` | Either implement type-specific connection creation (e.g., market data vs trading subscriptions) or remove the `PoolType` parameter entirely. |
| **P1** | `ClientPool.Put` doesn't validate the returned client belongs to the pool. | `internal/client/pool.go:135` | Add validation: if client was not obtained from this pool, return error or ignore. Prevent cross-pool contamination. |
| **P1** | `logf()` global logger initialization has a race: `sync.Once.Do` called between `RLock` release and `Do` execution. | `internal/client/client.go:36-51` | Initialize logger eagerly in `New()` before any goroutines start. Remove the fragile lazy `sync.Once` pattern. |
| **P2** | `CodePoolExhausted` returned in `pool.go:131` but never defined in `errors.go`. | `internal/client/pool.go`, `errors.go` | Define `CodePoolExhausted` error code constant alongside other `ErrCode*` constants. |
| **P2** | Garbled Korean/Chinese characters in comment. | `internal/client/client.go:212` | Replace with English: `// Metrics tracking`. |

---

## 2. SDK Quality & Stability

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Medium | **Restore full test suite** | `PROJECT_STATUS.md` claims "All 230+ unit tests pass" but `go test ./...` has multiple failures (see Section 1). Fix all test failures. Update status to reflect actual state. |
| **P0** | Low | **Add `go vet` to CI** | `go vet ./...` currently fails. Add `golangci-lint` with standard linters to GitHub Actions. Block merges on lint failures. |
| **P1** | Medium | **Context propagation** | Currently only `WithContext()` creates a context-aware client. Add optional `context.Context` parameter to ALL public API methods (`GetQuote`, `GetKLines`, `PlaceOrder`, etc.). Enable cancellation and timeout at the request level. |
| **P1** | Medium | **Structured error types** | All errors are currently generic. Define structured error types: `ErrConnectionFailed`, `ErrTimeout`, `ErrProtocol`, `ErrServerReject`, `ErrOrderRejected`. Include original error chain and server error code. |
| **P1** | Medium | **Connection keep-alive verification** | The keep-alive ping is sent but there's no verification that it receives a response. Add ping/pong verification: if no pong within timeout, trigger reconnect. |
| **P1** | Medium | **Request timeout per-call** | Currently the timeout is global (`opts.APITimeout`). Allow per-call timeout override: `WithTimeout(time.Second * 5)` option applied to a single request. |
| **P2** | Low | **Graceful shutdown** | `Close()` doesn't drain pending requests. Add graceful shutdown: signal context, wait for in-flight requests with timeout, then close connections. |
| **P2** | Medium | **Protobuf compatibility check** | Futu may update proto definitions. Add a CI check that verifies all proto fields used in code actually exist in the generated `.pb.go` files. |
| **P2** | Low | **API version negotiation** | Log a warning if `ServerVer` from OpenD doesn't match expected range. Some features may not be available on older OpenD versions. |

---

## 3. Performance & Scalability

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P1** | Medium | **Connection pool performance** | Current pool implementation has O(n) lookup in `Put()` and `Get()`. Use `sync.Map` or separate `map[PoolType][]*PoolConn` with proper locking for O(1) amortized lookups. |
| **P1** | Medium | **Zero-allocation request path** | The request/response path allocates slices and maps on every call. Use `sync.Pool` for reused `[]byte` buffers and pre-allocated response structs for hot paths. |
| **P1** | Low | **Batch request optimization** | Some APIs (e.g., `GetBasicQot`, `GetStaticInfo`) accept multiple securities. Ensure the internal serialization is efficient. Benchmark and optimize protobuf marshal/unmarshal for batch payloads. |
| **P2** | High | **WebSocket transport** | `PROJECT_STATUS.md` mentions `internal/ws/` exists but not wired in. Implement WebSocket as an alternative transport for push-heavy workloads (reduces TCP overhead for high-frequency quote streams). |
| **P2** | Medium | **Pre-computed protobuf pools** | Use `proto.Unmarshal` with pre-allocated message structs pooled via `sync.Pool` to reduce GC pressure during high-frequency push handling. |
| **P2** | Low | **Read buffer sizing** | The connection read buffer size should be configurable. Large push payloads (order book, tick streams) need appropriately sized buffers to avoid partial reads. |

---

## 4. API Design & Ergonomics

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Low | **Export `Packet` and `PacketHandler`** | The `client_test.go` failure (C-1) reveals that `Packet` and `PacketHandler` types are needed for testing but not exported. Either export them (`Packet`, `PacketHandler`) or provide test helper functions in a `client/testing` sub-package. |
| **P1** | Medium | **Standardize response types** | Some APIs return raw proto response types, others return custom wrapper types. Standardize: all public APIs return domain-specific types (e.g., `Quote`, `KLine`, `Position`). Auto-generate `FromProtobuf()` converters for all response structs. |
| **P1** | Medium | **Fluent API for Client options** | `WithContext()` doesn't copy all fields (metrics, pendingOrders, disp channels). Either document it as a shallow clone or implement deep copy properly. Add `WithTimeout()`, `WithRetry()` as additional option functions. |
| **P1** | Low | **Named constants for order operations** | Magic number `ModifyOrderOp: 1` for cancel. Define: `ModifyOrderOpModify = 0`, `ModifyOrderOpCancel = 1`. Same for `TrdSide` (Buy=1, Sell=2), `OrderType` constants. |
| **P2** | Medium | **Builder pattern for requests** | Some complex requests (e.g., `StockFilter`) have many optional fields. Add builder pattern: `StockFilterRequest.New().WithMarket(2).WithNum(100).Build()`. |
| **P2** | Low | **Dead code removal** | `requestInternal()` at `client.go:713` is nearly identical to `request()`. Either remove it (use `request` instead) or consolidate logic. |
| **P2** | Low | **API deprecation policy** | Define a process for deprecating functions. Mark old function signatures with `// Deprecated:` comments and point to replacements. |

---

## 5. Observability & Debugging

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | High | **Expose metrics via Prometheus** | The `Metrics` struct exists (`TotalRequests`, `FailedReqs`, `ReconnectCount`, etc.) but isn't exposed via HTTP. Add `/metrics` endpoint with Prometheus counters/gauges: `futuapi_requests_total`, `futuapi_request_duration_seconds`, `futuapi_reconnects_total`, `futuapi_push_messages_total`. |
| **P1** | High | **OpenTelemetry tracing** | Add trace context propagation: `trace.Span` for every request/response cycle. Export to Jaeger or OTLP collector. Key spans: `conn.writePacket`, `conn.readResponse`, `client.request`, push handlers. |
| **P1** | Medium | **Structured logging upgrade** | Current `logf()` uses standard library `log.Logger`. Migrate to `slog` (Go 1.21+) or `zerolog` for structured JSON logging with fields: `conn_id`, `proto_id`, `serial_no`, `latency_ms`. |
| **P1** | Medium | **Request/response logging** | Add a debug mode that logs all requests (ProtoID, serialNo) and responses (serialNo, latency). Make it toggleable via `WithLogLevel()` option. |
| **P2** | Medium | **Connection health dashboard** | Expose connection state as JSON via HTTP: `GET /debug/conn`. Show: `connected`, `conn_id`, `server_ver`, `pending_requests`, `reconnect_count`, `uptime`. |
| **P2** | Low | **Packet hex dump utility** | Add `(*Conn) DumpPacket([]byte) string` for hexdump debugging. Useful for diagnosing protocol issues. |

---

## 6. Testing & Reliability

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Medium | **Fix all broken tests** | `client_test.go`, `push_test.go`, `pool_test.go` all have failures (see Section 1). Fix each one. Target: `go test ./...` passes 100%. |
| **P0** | Medium | **Race detector in CI** | Add `go test -race ./...` to GitHub Actions. Race conditions in connection management are silent failures. Block merges on race detector failures. |
| **P1** | Medium | **Contract tests for protobuf** | Add table-driven tests verifying all proto fields are populated (no hardcoded zeros). Use fixture data to check field-level correctness for every API response type. |
| **P1** | High | **Fuzz testing** | Add Go fuzz tests for: (a) protobuf unmarshal with random bytes, (b) order book parsing with malformed data, (c) K-line data with extreme values. Run via `go test -fuzz`. |
| **P1** | Medium | **Chaos engineering for connections** | Add integration tests that simulate: (a) connection drop mid-request, (b) server rejects with various error codes, (c) slow responses (timeout), (d) partial reads. Use `test/util/mock_server.go` with configurable failure modes. |
| **P1** | Medium | **Benchmark regression CI** | Capture benchmark results in CI. If a PR causes benchmark regression > 10%, fail the build. Store results in JSON, compare against baseline. |
| **P2** | High | **Property-based testing** | Use `go-check.v2` or `testing/quick` for property-based tests: (a) K-line data always satisfies H≥L, (b) order book bid price < ask price, (c) position qty is non-negative. |

---

## 7. Security

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Low | **Secrets in env vars** | `README.md` shows hardcoded password example. Update all examples to load secrets from environment variables. Add `.env.example` file documenting required env vars. |
| **P1** | Low | **TLS support** | Futu OpenD supports TLS connections. Add `WithTLS()` option for secure connections (especially relevant if OpenD is accessed over network rather than localhost). |
| **P1** | Medium | **Input validation hardening** | `feat(validation)` commits exist but validation coverage should be verified. Add fuzz tests (6.4) and audit every public API for injection vectors (though protobuf is safe by default, string fields could contain malicious content). |
| **P1** | Low | **Rate limit awareness** | Document Futu's server-side rate limits (e.g., queries/sec). Add client-side rate limiting utility that respects server limits and returns `ErrRateLimited` with `Retry-After`. |
| **P2** | Medium | **Audit logging** | Log all trading operations (place order, cancel order, unlock) with timestamps, account IDs (masked), and order IDs. Required for security auditing. |

---

## 8. Documentation

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Low | **Fix ROADMAP.md** | Current `ROADMAP.md` is all placeholder TODOs. Replace with real roadmap based on this enhancement plan. |
| **P1** | Medium | **API changelog** | Add `docs/CHANGELOG.md` with auto-generated entries (use `git-cliff` or similar). Every API change must update the changelog. |
| **P1** | Medium | **Architecture decision records (ADRs)** | Document key decisions: (a) why `APIConnector` interface pattern, (b) why protobuf over JSON, (c) connection pool design choices, (d) push notification architecture. Store in `docs/adr/`. |
| **P1** | Low | **Update PROJECT_STATUS.md** | Current status claims all 230+ tests pass. Update to reflect actual test state after fixing broken tests. |
| **P1** | Medium | **Migrate README examples to working code** | `README.md:50` has a TODO note: "Full Environment Variable integration is planned for the next release." Either implement it or remove the misleading comment. |
| **P2** | Medium | **Performance profiling guide** | Add `docs/PERFORMANCE.md`: how to profile the SDK with `pprof`, interpret benchmark results, common bottlenecks and solutions. |
| **P2** | Low | **Stability report** | Add `docs/STABILITY.md`: known limitations, edge cases, OpenD version compatibility matrix. |

---

## 9. Developer Experience

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P0** | Medium | **Fix golangci-lint config** | Add `.golangci.yml` with standard linters (`govet`, `gofmt`, `staticcheck`, `unused`, `errcheck`). Configure `issues.exclude-rules` for generated protobuf code. |
| **P0** | Low | **go.mod upgrade path** | `go.mod` specifies `go 1.21`. Update to `go 1.26` to match `futugo4bot`. Document minimum Go version requirement clearly. |
| **P1** | Medium | **Example programs** | The `cmd/examples/` directory has 28 examples. Audit them: (a) ensure all compile, (b) add comments explaining what each does, (c) add a `README.md` index in `cmd/examples/` listing all examples. |
| **P1** | Low | **Update `go.mod` version** | `futuapi4go` is listed as `v0.0.0` in `futugo4bot/go.mod`. Tag a proper semver release (e.g., `v0.7.0`) so `futugo4bot` can depend on a proper version. |
| **P1** | Medium | **Commit message convention** | Add `.gitmessage` or document conventional commit format (`feat:`, `fix:`, `docs:`, `test:`). Integrate `commitlint` in CI. |
| **P2** | High | **Makefile** | Add `Makefile` with targets: `make test`, `make lint`, `make build`, `make bench`, `make干净` (clean). Consistent developer experience across platforms. |

---

## 10. Advanced Features

| Priority | Effort | Item | Action |
|----------|--------|------|--------|
| **P1** | High | **Historical data downloader** | `futugo4bot` has `cmd/history_download` but it's tightly coupled. Extract a reusable `DownloadHistory(client *Client, req HistoryDownloadRequest) error` function into the SDK. Handle pagination automatically. |
| **P1** | Medium | **Option chain utilities** | `GetOptionChain` returns raw proto types. Add helper functions: `ParseOptionCode(code string) (owner, expiry, strike, optType)`, `GetAtmOption(client, underlying, expiry)` returning ATM call/put. |
| **P1** | Medium | **Market hours utility** | Add `pkg/market/hours.go`: pre-computed HK market hours (9:15-12:00, 13:00-16:30, 17:15-03:00). `IsMarketOpen(t time.Time) bool`, `UntilClose(t time.Time) time.Duration`. Used by `futugo4bot` session filters. |
| **P2** | High | **Streaming API** | For high-frequency quote streams, offer a `Client.SubscribeStream(ctx, securities, callback)` that handles reconnection and backpressure automatically. |
| **P2** | Medium | **Account portfolio snapshot** | Add `GetPortfolioSnapshot(client *Client, accID uint64) (*PortfolioSnapshot, error)`: returns all positions + funds + open orders in a single API call (reduces round-trips on startup). |
| **P2** | Medium | **Webhook support** | Futu supports webhooks for push notifications. Add `WithWebhook(url, secret)` option that forwards push data to an HTTP endpoint. Useful for bridging to other systems. |
| **P3** | High | **GraphQL interface** | `PROJECT_STATUS.md` mentions this. Consider `gqlgen`-based GraphQL layer over the REST-like API for flexible querying. Lower priority — protobuf is already well-structured. |

---

## 11. Priority Roadmap

### Phase 0 — Fix the Foundation (P0 items only)
> *Before any SDK v1.0 release.*

| # | Item | Category | Why |
|---|------|----------|-----|
| 1 | Fix `client_test.go` compilation | Testing | Tests must compile; currently broken |
| 2 | Fix `push_test.go` protobuf types | Testing | 7 failures from wrong proto wrapper |
| 3 | Fix nil logger panic in `logf()` | Bug | Panic during normal operation |
| 4 | Fix connection state race (`readLoop`) | Bug | Silent data race in production |
| 5 | Add race detector to CI | Testing | Prevent future races |
| 6 | Fix `go vet` failures | Quality | Block merges on lint failures |
| 7 | Export `Packet`/`PacketHandler` for testing | API Design | Enables proper SDK testing |
| 8 | Update ROADMAP.md (replace stubs) | Docs | Current roadmap is misleading |
| 9 | Tag proper semver release (`v0.7.0`) | DX | `futugo4bot` depends on `v0.0.0` |
| 10 | Secrets in env vars (update examples) | Security | Current examples show hardcoded passwords |

### Phase 1 — Production Quality (P1 items)
> *SDK is reliable enough for serious trading bots.*

| # | Item | Category |
|---|------|----------|
| 1 | Prometheus metrics endpoint (`/metrics`) | Observability |
| 2 | OpenTelemetry tracing | Observability |
| 3 | Context propagation on all APIs | Quality |
| 4 | Structured error types | Quality |
| 5 | Connection health check (ping/pong) | Quality |
| 6 | Chaos connection tests (mock server failure modes) | Testing |
| 7 | TLS support | Security |
| 8 | Update PROJECT_STATUS.md (actual test state) | Docs |
| 9 | Example programs audit + index | DX |
| 10 | golangci-lint config | DX |

### Phase 2 — Performance & Polish (P2 items)
> *Fine-tuning for high-frequency trading workloads.*

| # | Item | Category |
|---|------|----------|
| 1 | WebSocket transport | Performance |
| 2 | Zero-allocation request path | Performance |
| 3 | Connection pool O(1) lookup | Performance |
| 4 | Historical data downloader in SDK | Advanced |
| 5 | Option chain utilities | Advanced |
| 6 | Market hours utility (`pkg/market/hours.go`) | Advanced |
| 7 | Structured logging (slog) | Observability |
| 8 | Benchmark regression CI | Testing |
| 9 | Performance profiling guide | Docs |
| 10 | Architecture Decision Records | Docs |

### Phase 3 — Ecosystem (P3 items)
> *Building a developer ecosystem around the SDK.*

| # | Item | Category |
|---|------|----------|
| 1 | Makefile with standard targets | DX |
| 2 | Fuzz testing pipeline | Testing |
| 3 | Property-based testing | Testing |
| 4 | GraphQL interface alternative | Advanced |
| 5 | Stability report (`docs/STABILITY.md`) | Docs |
| 6 | Commit convention + commitlint | DX |

---

## Summary

| Phase | Focus | P0 | P1 | P2 | P3 | Total |
|-------|-------|----|----|----|----|-------|
| **Phase 0** | Fix the Foundation | 10 | — | — | — | 10 |
| **Phase 1** | Production Quality | — | 10 | — | — | 10 |
| **Phase 2** | Performance & Polish | — | — | 10 | — | 10 |
| **Phase 3** | Ecosystem | — | — | — | 6 | 6 |

**Total: 36 enhancement items + 8 bug fixes across 8 categories.**

---

## Relationship to futugo4bot

This SDK is a direct dependency of [futugo4bot](https://github.com/shing1211/futugo4bot). The following SDK enhancements directly improve the trading bot:

| futuapi4go Item | Improves futugo4bot |
|----------------|-------------------|
| Structured error types | Better error handling in engine |
| Context propagation | Timeout control on all API calls |
| Prometheus metrics | Grafana dashboard completeness |
| Historical data downloader in SDK | Data pipeline reliability |
| Market hours utility | HK futures session filter accuracy |
| Connection health check | Feed monitoring reliability |
| Option chain utilities | Options strategy support |
| TLS support | Secure remote OpenD connections |
| WebSocket transport | Lower latency for high-frequency quotes |
| OpenTelemetry tracing | Distributed tracing in bot |

See [futugo4bot/docs/ENHANCEMENT_PLAN.md](https://github.com/shing1211/futugo4bot/blob/main/docs/ENHANCEMENT_PLAN.md) for the trading bot's enhancement roadmap.

---

## Related Documentation

- [docs/PROJECT_STATUS.md](PROJECT_STATUS.md) — Current SDK status
- [docs/TESTING.md](TESTING.md) — Testing guide
- [docs/DEVELOPER.md](DEVELOPER.md) — Architecture deep-dive
- [docs/API_REFERENCE.md](API_REFERENCE.md) — API reference
- [ROADMAP.md](../ROADMAP.md) — Project roadmap (update this with the phased plan above)
- [futugo4bot](https://github.com/shing1211/futugo4bot) — Trading bot powered by this SDK
