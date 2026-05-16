# FutuAPI4Go SDK - Advanced Enhancement Plan

> **Version**: v0.7.0 | **Date**: 2026-05-16 | **Status**: ACTIVE

---

## Overview

The `futuapi4go` SDK has reached **~99% API coverage** vs the Python SDK with all core trading and market data operations implemented. This document outlines the roadmap for transforming the SDK from a **Futu API wrapper** into a **complete professional trading SDK**.

**Current state**: v0.7.0 — Core APIs complete, Phase 0 bugfixes and Phase 1 production-hardening done.
**Target state**: Production-ready SDK with full API coverage, resilience, observability, and developer experience.

---

## Phase 0: Core SDK Robustness (Priority P0 — Prerequisite)

Before building advanced features, the SDK core must be hardened. All items in this phase are critical bugs, race conditions, or robustness gaps found during the v0.5.12 code audit.

### 0-1: Critical Bugs

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-1a | **AES ECB encrypts only 16 bytes** — `block.Encrypt`/`Decrypt` processes exactly one AES block. Any protobuf body >16 bytes is silently corrupted. | `internal/client/aes.go:36,78` | ✅ Done |
| 0-1b | **Operator precedence in `inferSecMarket`** — `&&` binds tighter than `||`. Code ≤3 chars with `.SZ` suffix panics with index-out-of-bounds. | `client/client.go:495` | ✅ Done |

### 0-2: Race Conditions (Fixed)

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-2a | `isEncrypt` accessed without synchronization in 8 locations (hot paths + encrypt/decrypt helpers). | `internal/client/client.go:300` | ✅ Done — atomics |
| 0-2b | `aesKey` read without mutex in `EncryptRequestBody`/`DecryptResponseBody`. | `internal/client/client.go:974,984` | ✅ Done |
| 0-2c | `MockServer.running` shared between `Stop` (write) and `acceptLoop`/`handleConnection` (read). | `test/util/mock_server.go:54` | ✅ Done — atomics |
| 0-2d | Test goroutines race on `conn.conn` and `serverConn` in 3 conn tests. | `internal/client/conn_test.go` | ✅ Done — channels |

### 0-3: Nil Safety Gaps

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-3a | `GetOrderBook` iterates nil items — no nil check on `ob` or `d` from proto lists. | `pkg/qot/quote.go:350-380` | ✅ Done |
| 0-3b | `GetFunds` — `s2c.GetFunds()` can be nil, panics on first field access. | `pkg/trd/trade.go:323` | ✅ Done |

### 0-4: Buffer Pool & Alloc Cleanup

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-4a | `requestInternal` marshal buffer pool: `GetMarshalBuf` + `PutMarshalBuf` around `proto.Marshal` — pooled buf content is never used (body var retained separately). Remove the dead pool calls. | `internal/client/client.go:1088-1091` | ✅ Done (removed) |
| 0-4b | `requestInternal`/`requestContextInternal` response buffer pool: `respBuf.data = plaintext` just reassigns slice header — no allocation saved. Removed in favor of direct `proto.Unmarshal`. | `internal/client/client.go:1171-1177,1235-1240` | ✅ Done |

### 0-5: Push Handler Registry Races

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-5a | `subscribeOne`'s stop function calls `RegisterHandler(protoID, nil)` — writing nil to handler map races with concurrent push dispatch. Fixed: handler checks a closed stop channel instead. | `pkg/push/chan/chan.go:118-131` | ✅ Done |
| 0-5b | All `Subscribe*` helpers use `context.Background()`. Changed to accept `ctx context.Context` as first parameter. | `pkg/push/chan/chan.go:133,138,143,198,203,208,213,218` | ✅ Done |

### 0-6: SDK Library Race Conditions

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-6a | `WithContext` shares `opts` struct pointer — caller mutations affect the original. Fixed: deep-copy via `optsCopy := *c.opts`. | `internal/client/client.go:1041` | ✅ Done |
| 0-6b | `GetAESKey()` (line 938) and `GetLoginUserID()` (line 953) are exported methods calling lock-protected reads. The `JoinGroup` pattern in `WithContext` accesses raw fields under RLock but shares `aesKey` string — this is safe because strings are immutable in Go, but fragile if `aesKey` is ever reassigned from another goroutine. | `internal/client/client.go:938,953,1047-1055` | ✅ Done — `getAESKey()` mutex accessor added |

### 0-7: Error Handling Inconsistencies

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-7a | `wrapError` in `pkg/sys/system.go` returned `*constant.FutuError` directly (inconsistent with `pkg/qot/quote.go` which returns `error`). Fixed: uses `constant.NewFutuError()` like all other packages. | `pkg/sys/system.go:60-65` | ✅ Done |
| 0-7b | Several `wrapError` calls in `pkg/qot/quote.go` are unused — some functions still use `fmt.Errorf` directly instead of the helper: `GetBasicQot` (line 175), `GetKL` (line 251), `GetOrderBook` (line 333). Inconsistent behavior. | `pkg/qot/quote.go:175,251,333` | ✅ Done |

### 0-8: Test Infrastructure & Coverage

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-8a | **No AES tests** — zero test coverage. Added 19 tests: round-trip (0B to 10KB), trailer validation, edge cases, multi-block padding, public wrappers, benchmarks. | `internal/client/aes_test.go` | ✅ Done |
| 0-8b | **No tests for** `pkg/metrics`, `pkg/ratelimit`, `pkg/retry`, `pkg/degradation`, `pkg/health`, `pkg/history`, `pkg/tracing`. | (multiple) | ✅ Done |
| 0-8c | `test/util/mock_server.go` protocol gap — mock receives InitConnect but doesn't complete the handshake response cycle correctly, causing all `test/qot_api` and `test/trd_api` tests to time out. | `test/util/mock_server.go:152-155` | ✅ Done |

### 0-9: Security & Data Protection

| Item | Description | File | Status |
|------|-------------|------|--------|
| 0-9a | **RSA private key in public key field** — `RSAEncrypt` accepts a private key PEM and extracts the public key from it. This is convenient for testing but dangerous: if a caller passes a private key PEM unintentionally (e.g., from a config file), the private key material is loaded into memory. Document that this is for backward compat only, warn against production use. | `internal/client/rsa.go:29-66` | ✅ Done — runtime logf warning added |
| 0-9b | `nonZeroRandomBytes` allocated a full `n`-byte buffer per iteration. Fixed: use `min(n, 64)` buffer for efficient random reads. | `internal/client/rsa.go:122-137` | ✅ Done |

---

## Phase 1: SDK Maturity & Production Readiness (Priority P0)

The SDK's existing resilience and production-readiness infrastructure must be fully wired and verified. Items in this phase ensure that rate limiting, retry, WebSocket reconnect, and other features are not just present in the codebase but actually functional in the request path.

### 1-1: Correct Wiring of Resilience Components

| Item | Description | File(s) | Status |
|------|-------------|---------|--------|
| 1-1a | **Wire rate limiter into core request path** — `ratelimit.ProtoLimiter` exists in `ClientOptions` but `requestInternal`/`requestContextInternal` never call `rl.Wait(ctx, protoID)`. Dead code today. Add rate limiter invocation at entry of both methods (before `proto.Marshal`). | `internal/client/client.go` | ✅ Done |
| 1-1b | **Wire retry into core request path** — `retry.Config` exists in `ClientOptions` but is never invoked. `Request()`/`RequestContext()` fail immediately instead of retrying transient failures. Wrap `requestInternal`/`requestContextInternal` in `retry.Do()` with recoverable error filtering (timeout + connection errors). | `internal/client/client.go` | ✅ Done |
| 1-1c | **Wire circuit breaker into WS reconnect** — Breaker is used in `Request()`/`RequestContext()` for non-control protos but is never checked in `reconnect()`. Add breaker state check before reconnection attempts to prevent reconnect storms when OpenD is down. | `internal/client/client.go` | ✅ Done |

### 1-2: Transport & API Parity

| Item | Description | File(s) | Status |
|------|-------------|---------|--------|
| 1-2a | **WebSocket auto-reconnect** — `reconnect()` only calls `ConnectWithRSA()` (TCP). WS connections silently die on disconnect. Add `isWebSocket` flag to Client struct, store in `connectWebSocket()`, modify `reconnect()` to call `connectWebSocket()` when WS transport is active. | `internal/client/client.go` | ✅ Done |
| 1-2b | **Add SkillWrapAPI proto wrapper** — `SkillWrapAPI.proto` (proto 8001) has compiled Go protos in `pkg/pb/skillwrapapi/` but zero Go wrapper code. Add `GetTechnicalUnusual()`, `GetFinancialUnusual()`, `GetDerivativeUnusual()` functions in new file `pkg/sys/skill_wrap.go` following existing wrapper patterns. | NEW `pkg/sys/skill_wrap.go`; `client/system_api.go` | ✅ Done |
| 1-2c | **Add Qot_GetTradeDate wrapper** — `Qot_GetTradeDate.proto` has compiled protos in `pkg/pb/qotgettradedate/` but no wrapper. Add `GetTradeDate()` function in new file `pkg/qot/trade_date.go`. (Note: `Qot_RequestTradeDate` IS wrapped as `RequestTradeDate()` in `pkg/qot/misc.go` — these are different protos.) | NEW `pkg/qot/trade_date.go`; `client/quote_api.go` | ✅ Done |
| 1-2d | **Add SetInitConnectConfig for WebSocket push fields** — `connectWebSocket()` doesn't send `SetInitConnectConfig` with push-related fields. Some OpenD versions require this for push delivery over WS. Add optional config send after InitConnect response. | `internal/client/client.go` `connectWebSocket()` | ⚪ Blocked — no `InitConnectConfig` proto in SDK |

### 1-3: Lifecycle & Observability

| Item | Description | File(s) | Status |
|------|-------------|---------|--------|
| 1-3a | **Connection state machine + callbacks** — Replace three atomic flags (`connected`, `connActive`, `reconnecting`) with a single `ConnState` enum (`Disconnected`→`Connecting`→`Connected`→`Reconnecting`→`Closing`). Add `OnStateChange` callback in options. Add public `State()` method. | `internal/client/client.go`; `client/client.go` | ✅ Done |
| 1-3b | **Graceful Shutdown(timeout)** — `Close()` immediately cancels context and kills connections without draining in-flight requests. Add `Shutdown(timeout)` that: (1) sets state to `Closing`, (2) rejects new requests, (3) waits for pending dispatches to drain, (4) calls `Close()`. | `internal/client/client.go` | ✅ Done |
| 1-3c | **Fluent API completion** — Audit all exported functions in `pkg/qot/`, `pkg/trd/`, `pkg/sys/`. Add missing `cli.Quote().GetXxx()` and `cli.Trade().GetXxx()` convenience wrappers. Estimated ~20 new fluent methods. | `client/quote_api.go`, `client/trade_api.go`, `client/system_api.go` | ✅ Done |
| 1-3d | **OpenTelemetry metrics bridge** — Create an OTel adapter for `pkg/metrics` to export counters/histograms via OTLP in addition to Prometheus. Add `SetupOTelMeter()` function that creates `metric.Meter` instruments mirroring existing Prometheus metrics. | NEW `pkg/tracing/otel/metrics.go` | ✅ Done |

### 1-4: Advanced Production Features

| Item | Description | File(s) | Status |
|------|-------------|---------|--------|
| 1-4a | **Historical KL data cache** — In-memory LRU cache for K-line data to reduce redundant API calls. Keyed by `security|klType|rehabType`. Configurable max entries (default 1000) and TTL (default 5min). Thread-safe with `sync.RWMutex`. Provide `GetKLCached()` wrapper that checks cache before calling `GetKL()`. | NEW `pkg/cache/kl_cache.go` | ✅ Done |
| 1-4b | **Structured slog integration** — `SlogLogger` exists in options but ~80% of log calls still use `logInfo`/`logWarn`/`logError` (which use `log.Printf`). Migrate all internal logging to `slog` when `SlogLogger` is set. Add span context (traceID/spanID) to log records when tracing is active. | `internal/client/client.go` (all log call sites) | ✅ Done |
| 1-4c | **Order pre-flight validation** — Before `PlaceOrder`, validate: market hours (`GetMarketState`), buying power (`GetFunds`), order size limits (`GetMaxTrdQtys`). Return `[]ValidationWarning` for non-blocking issues, `ValidationError` for blocking issues. | NEW `pkg/trd/validation.go` | ✅ Done |
| 1-4d | **Audit/compliance logging** — Structured log of all trade operations (PlaceOrder, ModifyOrder, ReconfirmOrder). Records: timestamp, user ID, security, side, qty, price, result, error. Pluggable output (slog, file, or custom writer). | NEW `pkg/trd/audit.go` | ✅ Done |

### Phase 1 Dependency Graph

```
1-2a (WS reconnect)     ← independent
1-2b, 1-2c (wrappers)   → 1-3c (fluent API)   [wrappers needed before fluent methods]
1-3a (state machine)    → 1-3b (graceful shutdown)  [state machine needed for Closing state]
1-3a (state machine)    → 1-1c (breaker in reconnect) [state needed for reconnection guard]
1-3d (OTel metrics)     ← independent
1-4a (KL cache)         ← independent
1-4b (slog)             ← independent
1-4c (validation)       ← independent
1-4d (audit)            ← independent
```

All items can be worked in parallel except where arrows indicate dependencies.

### Phase 1 Verification Strategy

| Item | Verification |
|------|-------------|
| 1-1a | Unit test with mock rate limiter returning `ErrRateLimited` → verify request returns error. Integration test: verify `Wait()` called with correct protoID. |
| 1-1b | Unit test: mock request fails twice then succeeds → verify retry executes 3 times. Verify `ErrRetryExhausted` after all retries fail. |
| 1-1c | Test breaker state check before reconnect by setting breaker to Open state. |
| 1-2a | Mock WS connection disconnect triggers `reconnect()` → verify `connectWebSocket` is called. |
| 1-2b, 1-2c | Table-driven test for each new wrapper: marshal C2S → unmarshal → compare fields. |
| 1-2d | Blocked — no `InitConnectConfig` proto in SDK. |
| 1-3a | Test all state transitions (Disconnected→Connecting→Connected→Reconnecting→Closing). Verify callback invoked for each transition. |
| 1-3b | Test `Shutdown` rejects new requests after state change. Test pending requests drain before close. No goroutine leaks. |
| 1-3c | Integration test: call each new fluent method against mock server. Verify correct proto ID and response parsing. |
| 1-3d | Test OTel meter instruments record correct values. Test dual-write (Prometheus + OTel) produces identical gauge values. |
| 1-4a | Test LRU eviction evicts oldest entry. Cache hit returns stored data (verify by timestamp). Cache miss fetches from real API. TTL expiry forces re-fetch. Thread safety under concurrent access. |
| 1-4b | Test slog output format includes expected fields. Test span context propagates into log records when tracing active. |
| 1-4c | Test with mock market states (open/closed). Test buying power shortfall. Test max qty enforcement. Test non-blocking warnings don't block order. |
| 1-4d | Test audit log records all trade operations. Verify all fields present. Test pluggable writer receives correct records. |

<!-- Phases A–F removed — application-layer features belong in a separate trading application repo -->

*Last updated: 2026-05-16*

