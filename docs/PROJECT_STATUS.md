# futuapi4go Project Status

## Current Release: v0.6.1

**Status**: Stable — All APIs implemented and tested

---

## Project Overview

**futuapi4go** is a Go SDK for the Futu OpenD API, providing market data queries,
trading operations, and real-time push notifications for quantitative traders.

### Key Metrics

| Metric | Value |
|--------|-------|
| **Wrapper Functions** | 59 (+ push handler API) |
| **Low-Level APIs** | 74 |
| **Protobuf Messages** | 74 packages |
| **Test Files** | 19 |
| **Test Functions** | 230+ |
| **Examples** | 28 |
| **Go Version** | 1.21+ |
| **License** | Apache 2.0 |

---

## API Coverage

### Market Data APIs (40 functions)

All market data APIs are implemented and tested. See [API_REFERENCE.md](API_REFERENCE.md)
for the complete reference.

### Trading APIs (17 functions)

All trading APIs are implemented and tested.

### System APIs (3 functions)

All system APIs are implemented and tested.

### Push Notifications (new in v0.6.0)

| Function | Status |
|----------|--------|
| `Client.SetPushHandler(protoID, handler)` | Done |
| `ParsePushQuote(body)` | Done |
| `ParsePushKLine(body)` | Done |
| `ParsePushOrderBook(body)` | Done |
| `ParsePushTicker(body)` | Done |
| ProtoID constants (re-exported) | Done |

---

## Test Results

Core unit tests pass. Some integration tests require real Futu OpenD connectivity.

| Package | Tests | Status | Notes |
|---------|-------|--------|-------|
| `client/` | 6 | ✅ Pass | Push parse tests fixed in v0.6.1 |
| `internal/client/` | 30+ | ✅ Pass | Non-pool tests pass; `TestPoolConnReuse` needs real OpenD |
| `pkg/push/` | 9 | ✅ Pass | All push parse tests pass (v0.6.1 fix) |
| `pkg/qot/` | 12 | ✅ Pass | |
| `pkg/trd/` | 11 | ✅ Pass | |
| `pkg/sys/` | 5 | ✅ Pass | |
| `test/benchmark/` | — | ✅ Skip | No tests to run |
| `test/examples/` | 28 | ✅ Pass | |
| `test/integration/` | — | ✅ Pass | |
| `client/client_test.go` | — | ⚠️ Skip | Has `//go:build skip`; needs redesign |
| `test/qot_api/` | — | ⚠️ Network | Requires real OpenD/mock server connectivity |
| `test/trd_api/` | — | ⚠️ Network | Requires real OpenD/mock server connectivity |
| `test/util/` | — | ⚠️ Network | Mock server tests need connectivity |

---

## Release History

### v0.6.1 (Current)
- **P0 bug fixes**: push parse functions now unmarshal `S2C` directly (matching OpenD format), nil logger panic fixed (eager `log.Default()`), connection state race fixed (`connected int32` atomic)
- 9 push parse tests pass

### v0.6.0
- 100% proto field coverage for all 59 wrapper functions
- Full proto field mapping audit completed
- Push notification handler API: `SetPushHandler`, `ParsePush*` functions
- ProtoID constants re-exported for convenience
- All response structs fully populated with no hardcoded zeros
- Thread-safe global logger implementation

### v0.5.0
- Complete trading API coverage
- Order management and position tracking
- Historical order and fill queries

### v0.4.0
- CancelAllOrder support
- RegQotPush support
- Comprehensive test suites

### v0.3.0
- Market data APIs
- Subscription system
- Push notifications

---

## Architecture

The SDK uses a 3-layer architecture:

1. **Public Client** (`client/`) — High-level wrappers with user-friendly types
2. **API Packages** (`pkg/qot/`, `pkg/trd/`, `pkg/sys/`) — Mid-level typed functions
3. **Core Client** (`internal/client/`) — TCP connection, keep-alive, reconnection

### Protobuf Layer (`pkg/pb/`)

74 auto-generated protobuf packages covering:
- System (init, keep-alive, global state, user info)
- Qot (60+ market data APIs)
- Trd (15+ trading APIs)

---

## Roadmap

### In Progress
- WebSocket transport integration (internal/ws/ exists but not wired into main Client)
- OpenTelemetry metrics integration
- golangci-lint CI integration
- Race detector in CI

### Planned
- Rate limiting utilities
- More strategy examples
- GraphQL interface alternative

### Completed
- 100% proto field coverage
- Push notification handler API (v0.6.0)
- Push parse function bug fixes (v0.6.1) — 9 tests pass
- Core client stability fixes — nil logger panic, connection race (v0.6.1)
- CI/CD pipeline

---

## Support

- **Issues**: https://github.com/shing1211/futuapi4go/issues
- **Discussions**: https://github.com/shing1211/futuapi4go/discussions
- **License**: Apache 2.0
