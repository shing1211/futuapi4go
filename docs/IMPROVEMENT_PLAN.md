# futuapi4go Robustness Improvement Plan

> **Generated:** 2026-05-18 | **Updated:** 2026-06-29 | **Version:** v0.14.0 | **Status:** Substantially complete — all HIGH/MED priority fixes done, LOW polish items L01-L14 done
>
> **Related:** [DESIGN.md](DESIGN.md) · [ARCHITECTURE.md](ARCHITECTURE.md) · [CHANGELOG.md](../CHANGELOG.md) · [AGENTS.md](../AGENTS.md)

---

## 1. Executive Summary

A comprehensive robustness audit of the futuapi4go Go SDK was conducted across **8 dimensions**: input validation, error handling, concurrency safety, proto safety, connection robustness, memory/resource leaks, trading safety, and security. The AGENTS.md checklist was used as the baseline standard.

**Overall Assessment:** The codebase demonstrates strong defensive programming. The AGENTS.md guidelines are consistently followed. **No critical (money-losing) bugs found.** The SDK is production-safe for read operations (market data, system queries, subscriptions). Trading operations require care around idempotency (a standard caveat for all trading SDKs).

### Robustness Scores

| Dimension | Score | Key Strength | Key Gap |
|-----------|-------|--------------|---------|
| Input Validation | 8/10 | All qot/trd/sys functions validate nil and required fields | `GetDelayStatistics` bypasses connection checks |
| Error Handling | 8/10 | `FutuError` with categories, `wrapError` used consistently | `CancelAllOrder` broken; retType specificity lost |
| Concurrency Safety | 8/10 | `sync.Mutex`/`RWMutex`/`atomic` consistently used | readLoop goroutine leak on hung TCP |
| Proto Safety | 8/10 | All S2C nil-checked; list iteration nil-guarded | `GetHeader().GetAccID()` chain in ModifyOrder |
| Connection Robustness | 7/10 | Reconnect with backoff; circuit breaker; WebSocket + TCP | Goroutine leak; no read deadline in readOne |
| Resource Leaks | 8/10 | WaitGroup tracked; DrainDispatches on close | readOne goroutine not cancellable |
| Trading Safety | 7/10 | Pre-flight validation; SensitiveString for passwords | Price validation not mandatory; no duplicate order prevention |
| Security | 9/10 | `crypto/rand`; `SensitiveString`; no hardcoded secrets | `RSAPrivateKey` stored as plain string in options |

### Executive Recommendation

**Phase 1 (now):** Fix the 1 HIGH and 6 MEDIUM items. These are concrete bugs/gaps with clear fixes.

**Phase 2 (next release):** Address the 13 LOW items for production-grade polish.

**Phase 3 (future):** Architectural hardening — goroutine lifecycle, read deadline management, structured logging unification.

---

## 2. Priority Fixes

### 2.1 Phase 1 — High Priority

#### FIX-001: `CancelAllOrder` is broken

**Severity:** HIGH | **File:** `client/trade_api.go:86-98` + `pkg/trd/orders.go:244-246`

**Root Cause:** `CancelAllOrder` calls `trd.ModifyOrder()` with `OrderID=0` and `ForAll=true`. The validation in `ModifyOrder` rejects `OrderID==0` unconditionally, even when `ForAll` is set.

**Current code (orders.go:244-246):**
```go
if req.OrderID == 0 && req.OrderIDEx == "" {
    return nil, fmt.Errorf("ModifyOrder: OrderID or OrderIDEx is required")
}
```

**Fix:**
```go
if req.OrderID == 0 && req.OrderIDEx == "" && !req.ForAll {
    return nil, fmt.Errorf("ModifyOrder: OrderID or OrderIDEx is required")
}
```

**Verification:** Run `go run ./examples/54_cancel_all_order` with a mock server to confirm it no longer fails validation.

---

### 2.2 Phase 1 — Medium Priority

#### FIX-002: Proto `GetHeader()` chain violates SDK safety guidelines

**Severity:** MEDIUM | **File:** `pkg/trd/orders.go:326-330, 399-403`

**Root Cause:** `ModifyOrder` and `ReconfirmOrder` call chained proto getters: `s2c.GetHeader().GetAccID()`. `GetHeader()` can return `nil`, and chained `GetXxx()` calls are explicitly forbidden by AGENTS.md.

**Current code:**
```go
AccID: s2c.GetHeader().GetAccID()
```

**Fix:**
```go
header := s2c.GetHeader()
if header == nil {
    return nil, wrapError("ModifyOrder", int32(common.RetType_RetType_Unknown), "header is nil")
}
AccID: header.GetAccID()
```

---

#### FIX-003: Goroutine leak in `readLoop` on hung TCP connection

**Severity:** MEDIUM | **File:** `internal/client/client.go:943-952`

**Root Cause:** `readLoop` spawns a new goroutine per iteration to call `readOne()`. If the TCP connection hangs without RST (e.g., firewall drop), `readOne()` blocks indefinitely on `io.ReadFull`. The parent goroutine has no mechanism to abort it.

**Current code:**
```go
go func() {
    pkt, err := c.conn.readOne()
    // ...
}()
```

**Fix — Add read deadline:**
```go
// In readLoop, before spawning the goroutine:
c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
go func() {
    pkt, err := c.conn.readOne()
    // ...
}()
```

Alternative: pass context to `readOne()` and use `io.ReadFull` with context-aware wrapper. This is the more thorough fix but requires refactoring `readOne`.

---

#### FIX-004: Mandatory price validation before order placement

**Severity:** MEDIUM | **File:** `pkg/trd/orders.go:140-142` + `pkg/trd/validation.go:84-90`

**Root Cause:** `PlaceOrder` sets price only if `req.Price != 0`. A limit order with price=0 passes through to the server. `ValidateOrder()` catches this but is an optional call, not integrated into `PlaceOrder()`.

**Current code (orders.go:140-142):**
```go
if req.Price != 0 {
    c2s.Price = &price
}
```

**Fix — Integrate price check into PlaceOrder:**
```go
// After line 115 (existing validation), add:
if req.OrderType != constant.OrderType_Market && req.OrderType != constant.OrderType_MarketToLimit {
    if req.Price <= 0 {
        return nil, wrapError("PlaceOrder", 
            constant.ErrCodePriceOutOfRange, 
            "price must be positive for limit orders")
    }
}
```

---

#### FIX-005: `GetDelayStatistics` bypasses connection safety checks

**Severity:** MEDIUM | **File:** `pkg/sys/system.go:352-363`

**Root Cause:** `GetDelayStatistics` calls `c.Conn().WritePacket()` and `ReadResponseContext()` directly on the raw `Conn`, bypassing `EnsureConnected()`. If the connection drops between the check and the write, `WritePacket` hits a nil `net.Conn`.

**Current code:**
```go
c.Conn().WritePacket(protoID, serialNo, body)
pkt, err := c.Conn().ReadResponseContext(ctx, serialNo, apiTimeout)
```

**Fix:**
```go
if err := c.EnsureConnected(); err != nil {
    return nil, fmt.Errorf("GetDelayStatistics: %w", err)
}
// Only proceed if connection is alive
c.Conn().WritePacket(protoID, serialNo, body)
```

---

#### FIX-006: Error code granularity lost in `wrapError`

**Severity:** MEDIUM | **Files:** `pkg/qot/quote.go:56-75`, `pkg/trd/trade.go:120-139`, `pkg/sys/system.go:51-61`

**Root Cause:** All three `wrapError` functions map server retType values to a small set of generic codes. Server-specific error codes (e.g., insufficient funds, price out of range, order rejected) all become `ErrCodeUnknown`, forcing callers to string-match error messages.

**Current mapping:**
```go
switch retType {
case -500:
    code = constant.ErrCodeInvalidParams
case -400:
    code = constant.ErrCodeUnknown
// ...
}
```

**Fix — Add specific mappings:**
```go
switch retType {
case -500:
    code = constant.ErrCodeInvalidParams
case -401:
    code = constant.ErrCodeOrderRejected
case -402:
    code = constant.ErrCodePriceOutOfRange
case -403:
    code = constant.ErrCodeInsufficientFunds
case -404:
    code = constant.ErrCodeQtyInvalid
// ...
}
```

**Note:** Requires documenting the server's retType-to-meaning mapping from the Futu API docs.

---

#### FIX-007: `RSAPrivateKey` stored as plain string in `ClientOptions`

**Severity:** MEDIUM | **File:** `internal/client/client.go:152`

**Root Cause:** `ClientOptions.RSAPrivateKey` is a plain `string`. If options are ever logged (e.g., debug logging enabled), the private key appears in plaintext.

**Fix:**
```go
// In ClientOptions struct:
RSAPrivateKey constant.SensitiveString // was: string
```

Update all reads to use `.Raw()` or direct conversion.

---

## 3. Phase 2 — Low Priority Polish

### 3.1 Proto Safety

| # | File:Line | Issue | Fix |
|---|-----------|-------|-----|
| L01 | `pkg/trd/position.go:666` | Missing `if item == nil { continue }` in `GetFlowSummary` list iteration | Add nil guard |
| L02 | `client/push.go:46` | `ParsePushKLine` dereferences `data.KLList[0]` without nil-checking element | Add `if data.KLList[0] == nil { return }` |

### 3.2 Concurrency & Connection

| # | File:Line | Issue | Fix |
|---|-----------|-------|-----|
| L03 | `internal/client/conn.go:262` | `ReadResponse` timer not drained — goroutine may live slightly longer | Add `<-timer.C` drain after `Stop()` |
| L04 | `internal/client/conn.go:129` | TOCTOU race in `Conn.Close()` — two goroutines could pass nil check | Use `sync.Once` or atomic guard |
| L05 | `internal/client/ws.go:92` | WebSocket `SetReadDeadline` error silently discarded | Log the error at debug level |
| L06 | `internal/client/pool.go:321` | Pool health check can shrink to zero clients | Add minimum pool size floor |

### 3.3 Trading Safety

| # | File:Line | Issue | Fix |
|---|-----------|-------|-----|
| L07 | `client/trade_api.go:51` | `client.PlaceOrder` delegates all validation to `trd.PlaceOrder` — no re-validation | Acceptable (chain-of-trust), but document |
| L08 | `pkg/trd/validation.go:84` | Zero price on limit order is a warning, not error | Consider upgrading to error for non-zero-commission stocks |
| L09 | `pkg/trd/queries.go:482` | `GetHistoryOrderFillList` mutates input struct with default values | Document or use defensive copy |

### 3.4 Error Handling

| # | File:Line | Issue | Fix |
|---|-----------|-------|-----|
| L10 | `pkg/sys/system.go:51` | sys `wrapError` uses different logic from qot/trd | Unify to single shared implementation |
| L11 | `pkg/retry/retry.go:69` | Non-recoverable error masks context cancellation | Check `ctx.Err()` before returning non-recoverable error |
| L12 | `internal/client/client.go:688` | `ConnectWithRSA` logs errors via `logInfo` (not `logError`) | Switch to `logError` for error-level messages |

### 3.5 Documentation

| # | Issue | Fix |
|---|-------|-----|
| L13 | `client/client.go:71` — `WithTradeEnv` creates shallow copy sharing connection | Document that both clients share the same connection |
| L14 | No retry safety warning for trading operations | Add AGENTS.md note: "Do not use retry package with PlaceOrder/ModifyOrder" |

---

## 4. Architectural Improvements (Phase 3)

### 4.1 Goroutine Lifecycle Hardening

The readLoop goroutine model has fundamental visibility and cancellation gaps:

```
Current:  readLoop() → spawn internal goroutine → readOne() [blocks forever]
          Cannot cancel the inner goroutine. Context only checked between iterations.

Proposed: readLoop() → pass context → readOne(ctx) → io.ReadFull with context-aware deadline
          Or: SetReadDeadline() periodically, check ctx.Done() on timeout
```

**Impact:** Prevents goroutine leaks when TCP connections hang. Reduces goroutine count from O(blocked-connections) to O(1).

### 4.2 Read Deadline Management

`readOne()` (conn.go:178) blocks on `io.ReadFull(c.conn, header)` with no deadline. A hung connection (e.g., firewall drop, keep-alive timeout) blocks indefinitely.

**Proposed:**
```go
// In readLoop, before each readOne:
c.conn.SetReadDeadline(time.Now().Add(c.opts.APITimeout))
pkt, err := c.conn.readOne()
```

### 4.3 Unified `wrapError`

Three separate `wrapError` implementations (qot, trd, sys) have drifted. Unify into a single shared function in `pkg/constant/`:

```go
// pkg/constant/wrap_error.go
func WrapError(fnName string, retType int32, retMsg string) *FutuError
```

### 4.4 Structured Logging Unification

Current: `logf` (global logger) + `c.logInfo`/`c.logWarn`/`c.logError` (client logger) coexist. Some callers use the wrong function (e.g., `logInfo` for errors).

**Proposed:** All callers use `c.logXxx` methods. Remove global `logf`. Add structured fields (serialNo, protoID) to log context.

---

## 5. Testing Gaps

### 5.1 Missing Test Coverage

| Area | Current State | Needed |
|------|---------------|--------|
| `CancelAllOrder` | No test verifies the `ForAll` path | BVT test against mock server |
| `GetFlowSummary` | No test | Unit test with nil list items |
| `GetDelayStatistics` | No test | Test with disconnected client |
| `reconnect()` | No integration test | Simulated disconnect + reconnect cycle |
| Goroutine leaks | No test | `runtime.NumGoroutine()` checks after close |
| Price=0 order | No test | Reject test |
| `readLoop` error recovery | No test | Simulated TCP hang |
| Retry idempotency | No test | Trade + retry = single order |

### 5.2 Test Infrastructure Improvements

- **Mock server:** The existing `test/util/mock_server.go` has a timing bug where it closes before the client reads the response. Fix the connection lifecycle.
- **Chaos testing:** Add `test/chaos/` with network disruption scenarios (drop packets, delay responses, partial reads).
- **Race detector:** Already enabled (`go test -race`), but add CI enforcement.

---

## 6. Security Hardening

### 6.1 Immediate

| # | Action | Rationale |
|---|--------|-----------|
| S01 | `ClientOptions.RSAPrivateKey` → `SensitiveString` | Prevents accidental logging |
| S02 | Audit all `logInfo` calls for sensitive data | Ensure no account IDs, order IDs logged at info |
| S03 | Add `String()` method to `ClientOptions` that redacts keys | If options are ever `%v` printed |

### 6.2 Future

| # | Action | Rationale |
|---|--------|-----------|
| S04 | Memory zeroing for AES keys on close | Defense-in-depth |
| S05 | Token bucket for failed auth attempts | Rate limit brute force |
| S06 | Audit proto response sizes before unmarshal | Prevent OOM on malicious server |

### 6.3 Non-Issues Confirmed

- `crypto/rand` used for all cryptographic operations ✅
- `SensitiveString` redacts all `fmt` output ✅
- No hardcoded secrets in source ✅
- MD5 only used per protocol requirement (documented) ✅
- AES-ECB only used per Futu protocol requirement (documented) ✅

---

## 7. Implementation Roadmap

### Release v0.9.1 — "Safety Patch" (1-2 days)

```
Phase 1 (HIGH + MEDIUM)
├── FIX-001: CancelAllOrder broken            ← P0, immediate fix
├── FIX-002: Proto GetHeader() nil chain      ← prevents potential panic
├── FIX-003: readLoop goroutine leak           ← prevents resource exhaustion
├── FIX-004: Mandatory price validation        ← trading safety
├── FIX-005: GetDelayStatistics safety         ← prevents panic on nil conn
├── FIX-006: Error code granularity            ← improves debugging
└── FIX-007: RSAPrivateKey SensitiveString     ← security hardening
```

### Release v0.9.2 — "Polish" (3-5 days)

```
Phase 2 (LOW priority)
├── L03-L06: Concurrency hardening
├── L07-L09: Trading polish
├── L10-L12: Error handling polish
├── L01-L02: Proto nil guards
└── L13-L14: Documentation
```

### Release v0.10.0 — "Architecture" (1-2 weeks)

```
Phase 3
├── 4.1: Goroutine lifecycle hardening
├── 4.2: Read deadline management
├── 4.3: Unified wrapError
├── 4.4: Structured logging unification
└── 5.2: Test infrastructure improvements
```

---

## 8. Audit Methodology

This plan was generated through a systematic, multi-layer audit:

### Layer 1: Sub-Agent Code Analysis
Two specialized sub-agents performed deep analysis:
- **Sub-Agent 1:** Input validation, error handling, concurrency, proto safety, connection robustness, resource leaks, trading safety, security — across all source files.
- **Sub-Agent 2:** RetType checking, S2C nil guards, proto list iteration, error wrapping, context usage, timeout handling, retry safety — focused on API response handling.

### Layer 2: Knowledge Graph Query
GitNexus code intelligence provided:
- **762 communities** auto-detected via Leiden clustering
- **300 execution flows** traced from entry points to terminals
- **32,662 relationships** between code symbols

### Layer 3: Manual Code Review
Key files examined:
- `internal/client/client.go` — connection lifecycle, reconnect, keepalive
- `internal/client/conn.go` — packet I/O, dispatch, read/write
- `internal/client/rsa.go`, `aes.go` — cryptographic operations
- `pkg/qot/quote.go`, `kline.go`, `sub.go` — market data
- `pkg/trd/orders.go`, `position.go`, `queries.go` — trading
- `pkg/sys/system.go` — system operations
- `client/trade_api.go`, `quote_api.go` — public API wrappers

### Layer 4: Standards Alignment
All findings evaluated against:
- AGENTS.md checklist (11 categories)
- Go best practices (Effective Go, Go Code Review Comments)
- Futu OpenD protocol specification
- Financial API safety patterns

---

## 9. Quick Reference

| Fix ID | Priority | File | Lines | Summary |
|--------|----------|------|-------|---------|
| FIX-001 | HIGH | `pkg/trd/orders.go` | 244-246 | CancelAllOrder broken |
| FIX-002 | MEDIUM | `pkg/trd/orders.go` | 326,399 | GetHeader() nil chain |
| FIX-003 | MEDIUM | `internal/client/client.go` | 943 | readLoop goroutine leak |
| FIX-004 | MEDIUM | `pkg/trd/orders.go` | 140 | Price validation missing |
| FIX-005 | MEDIUM | `pkg/sys/system.go` | 352 | DelayStats bypass |
| FIX-006 | MEDIUM | `pkg/{qot,trd,sys}/*.go` | multiple | Error code granularity |
| FIX-007 | MEDIUM | `internal/client/client.go` | 152 | RSA key plaintext |
| L01-L14 | LOW | multiple | multiple | Polish items |

---

## 10. Related Documents

| Document | Purpose |
|----------|---------|
| [ARCHITECTURE.md](ARCHITECTURE.md) | System architecture, layers, execution flows |
| [DESIGN.md](DESIGN.md) | Design decisions, API patterns, security model |
| [CHANGELOG.md](../CHANGELOG.md) | Release history (v0.5.7 → v0.9.0) |
| [AGENTS.md](../AGENTS.md) | Development standards, code review checklist |
| [README.md](../README.md) | API reference, installation, quick start |
