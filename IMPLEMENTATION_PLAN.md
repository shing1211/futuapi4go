# Implementation Plan — futuapi4go

*Generated: 2026-05-20*

## Phase 1: Bug Fixes (High Priority) — CURRENT

### 1.1 Retry on Trading Operations (CRITICAL)
- **File**: `internal/client/client.go:1325-1365`
- **Bug**: `Request()` and `RequestContext()` apply `retry.Do()` to ALL proto requests including `PlaceOrder`, `ModifyOrder`, `CancelOrder` when `retryConfig` is set. This can cause duplicate orders.
- **Fix**: Add `isTradingProto(protoID)` check. Skip retry for trading proto IDs.

### 1.2 `return nil, nil` Ambiguity in Quote APIs
- **Files**: `client/quote_api.go:874`, `client/quote_api.go:934`
- **Bug**: `GetMarketState` and `GetCapitalDistribution` return `(nil, nil)` when data is empty.
- **Fix**: Return empty slice/zero-value struct instead of nil.

### 1.3 Silent Error Dropping in Push Channel Handler
- **File**: `pkg/push/chan/chan.go:145-148`
- **Bug**: Parse errors in push handler are silently dropped.
- **Fix**: Add `slog.Warn` logging when parse fails.

### 1.4 Silent Error Dropping in Push Callbacks
- **File**: `client/push_callbacks.go:119-122`
- **Bug**: Callback errors are swallowed with zero observability.
- **Fix**: Add `slog.Warn` before swallowing.

### 1.5 `WithEnvConfig` File Read Error Masking
- **File**: `client/client.go:409-429`
- **Bug**: If file exists but is unreadable, error is silently swallowed.
- **Fix**: Distinguish file-not-found from read errors.

## Phase 2: Input Validation (Next)
- Add validation to all `client/quote_api.go` functions (28 functions)
- Add validation to all `client/trade_api.go` functions (17 functions)
- Add validation to `pkg/qot/` lower-level functions

## Phase 3: Proto Safety Improvements
- Replace `GetXxx()` with direct field access (1,287 instances)
- Incremental: one file at a time with tests

## Phase 4: Missing High-Level API Coverage
- Add `client/` wrappers for `RequestHistoryKLQuota`, `GetUserSecurityGroup`, etc.
- Add fluent API methods for missing functions

## Phase 5: Advanced Features
- Connection pool health check improvements
- Middleware/interceptor pattern
- Observability integration (metrics, tracing)
- Request/response logging with sensitive data redaction

## Phase 6: Architecture Improvements
- PoolType-aware routing
- Error category improvements (Transient vs Permanent)
- Deprecate `Request()` in favor of `RequestContext()`
