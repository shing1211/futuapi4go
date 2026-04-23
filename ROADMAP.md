# futuapi4go Roadmap

> **Last updated: 2026-04-23 — v0.9.0 published** 🎉

---

## Vision

futuapi4go is the most reliable, well-tested, and ergonomic Go SDK for the Futu OpenAPI — powering production trading systems with institutional-grade observability, structured error handling, and comprehensive test coverage.

The Go SDK has already surpassed the Python SDK in proto coverage (78 vs ~50), type safety, and performance. The mission now is production polish and developer experience.

---

## v0.9.0 — Feature Parity Achieved ✅

Released 2026-04-23. Full parity with the Python SDK plus Go-native extras.

**New methods:**
- `GetLoginUserID()` — Futu/NiuNiu user ID from OpenD
- `IsEncrypt()` — whether the connection uses AES encryption
- `CanSendProto()` — connection state check before sending a proto

**New packages:**
- `pkg/breaker` — circuit breaker pattern
- `pkg/logger` — structured leveled logging (text + JSON)
- `pkg/push/chan` — channel-based push delivery
- `pkg/util` — code parsing (`ParseCode`, `FormatCode`, market helpers)
- `pkg/constant` — Python-style `String()` methods on all enums

**SDK enhancements:**
- `GetAccountInfo` — full account info with multi-currency cash and per-market assets
- `GetFlowSummary` — account cash flow entries
- `GetAccTradingInfo` — max tradable quantities + margin info
- Extended `Funds` struct with 16 new fields

**Tests:**
- Unit tests for all new packages (`pkg/util`, `pkg/constant`, `pkg/logger`, `pkg/breaker`)

---

## Feature Matrix: Go vs Python

| Category | Python SDK | Go SDK | Winner |
|----------|-----------|--------|--------|
| Protobufs | ~50 | **78** | 🏆 Go (+28 extras) |
| Quote/Market APIs | ~48 | **~60** | 🏆 Go |
| Trading APIs | ~18 | **~17** | Tie |
| System APIs | ~4 | **~6** | 🏆 Go (GetLoginUserID, IsEncrypt, CanSendProto) |
| Push parsing | ~7 | **~11** | 🏆 Go |
| Channel push | ❌ | **✅** | Go only |
| Circuit breaker | ❌ | **✅** | Go only |
| Structured logging | ❌ | **✅ text+JSON** | Go only |
| Connection pool | ❌ | **✅** | Go only |
| Code parsing helpers | ❌ | **✅** | Go only |

---

## What's Next

### Phase 1 — Production Hardening 🏗️

*SDK is reliable enough for serious trading bots.*

| # | Item | Priority |
|---|------|----------|
| 1 | Prometheus metrics endpoint (`/metrics`) | P1 |
| 2 | OpenTelemetry distributed tracing | P1 |
| 3 | Structured error types (`ErrConnectionFailed`, etc.) | P1 |
| 4 | Connection ping/pong health verification | P1 |
| 5 | Connection chaos tests (mock server failure modes) | P1 |
| 6 | TLS support | P1 |
| 7 | golangci-lint CI configuration | P1 |
| 8 | Example programs audit + index | P1 |

### Phase 2 — Performance & Polish

*Fine-tuning for high-frequency trading workloads.*

| # | Item | Priority |
|---|------|----------|
| 1 | WebSocket transport alternative | P2 |
| 2 | Zero-allocation request hot path | P2 |
| 3 | Connection pool O(1) lookup | P2 |
| 4 | Historical data download utility in SDK | P2 |
| 5 | Option chain helper functions | P2 |
| 6 | HK market hours utility | P2 |
| 7 | Benchmark regression CI | P2 |
| 8 | Performance profiling guide | P2 |
| 9 | Architecture Decision Records (ADRs) | P2 |

### Phase 3 — Ecosystem

*Building a developer community around the SDK.*

| # | Item | Priority |
|---|------|----------|
| 1 | Makefile with standard targets | P3 |
| 2 | Fuzz testing pipeline | P3 |
| 3 | Property-based testing | P3 |
| 4 | GraphQL interface alternative | P3 |
| 5 | Stability report | P3 |
| 6 | Commit convention + commitlint | P3 |

---

## Known Issues

### Proto Compatibility (OpenD serverVer=1003)

| Issue | Status |
|-------|--------|
| `GetDelayStatistics` — proto2 wire-format incompatibility (OpenD rejects packed encoding for `repeated int32`) | Known |
| `GetTradeDate` — all C2S fields required; may also be proto2 issue | Known |

Both are SDK-level protobuf marshaling issues. Workarounds are in place (calls are skipped gracefully). Fix planned for a future release.

---

## Powered By

futuapi4go powers [futugo4bot](https://github.com/shing1211/futu4bot) — a production algorithmic trading bot for HK futures.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

*Last comprehensive review: 2026-04-23*
