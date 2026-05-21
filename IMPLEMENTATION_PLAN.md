# Implementation Plan — futuapi4go

*Generated: 2026-05-20 | Updated: 2026-05-21*

## Phase 1: Bug Fixes (High Priority) — DONE

### 1.1 Retry on Trading Operations (CRITICAL) — DONE

### 1.2 `return nil, nil` Ambiguity in Quote APIs — DONE

### 1.3 Silent Error Dropping in Push Channel Handler — DONE

### 1.4 Silent Error Dropping in Push Callbacks — DONE

### 1.5 `WithEnvConfig` File Read Error Masking — DONE

## Phase 2: Input Validation — DONE (Phase 3 complete)

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

## Phase 2: Input Validation — DONE

### 2.1 `client/quote_api.go` Validation (46 functions)
Every function accepting `code` validates `len(code) > 0`. Every function accepting `num` validates `num > 0`. Every function accepting `market` validates it's a known constant. Date range functions validate `startDate`/`endDate` or `beginTime`/`endTime` non-empty. Slice params (`securities`, `subTypes`, `codes`, `times`) validated non-empty. `maxPerPage`/`maxNum` validated > 0.

### 2.2 `client/trade_api.go` Validation (18 functions)
Every function accepting `accID` validates `accID != 0`. `UnlockTrading` validates `pwdMD5` non-empty. `ModifyOrder` validates `orderID != 0`. `SubAccPush` validates `accIDList` non-empty. `GetOrderFee` validates `orderIDExList` non-empty. `GetMarginRatio` validates `securities` non-empty.

### 2.3 `client/system_api.go` Validation
`TestCmd` validates `cmd` non-empty. `Verification` validates `req != nil`.

### 2.4 `pkg/trd/convenience.go` Validation
`QuickBuy/QuickSell` validate `accID != 0`, `code != ""`, `qty > 0`, `price > 0`. `QuickMarketBuy/QuickMarketSell` validate `accID != 0`, `code != ""`, `qty > 0`. `GetPositions/GetTodayFills/GetTodayOrders/GetAccountFunds` validate `accID != 0`.

## Phase 3: Proto Safety Improvements — DONE

- Replaced all `GetXxx()` method calls on protobuf messages in SDK source code with nil-safe direct field access using `util.ProtoStr/ProtoInt32/ProtoFloat64/ProtoBool/ProtoInt64/ProtoUint64/ProtoFloat32` helpers
- Added `pkg/util/proto_helpers.go` with the helper functions
- Processed ~30 files across `client/`, `pkg/trd/`, `pkg/qot/`, `pkg/sys/`, `pkg/push/`
- ~1,200+ GetXxx() calls replaced
- Known exclusions: `c.inner.GetConnID()`, `c.inner.GetServerVer()`, `c.inner.GetLoginUserID()` are NOT proto getters (internal client methods)
- Detail plan: `PHASE3_PROTO_SAFETY_PLAN.md`

## Phase 4: Missing High-Level API Coverage — FUTURE
- Add `client/` wrappers for `RequestHistoryKLQuota`, `GetUserSecurityGroup`, etc.
- Add fluent API methods for missing functions

## Phase 5: Advanced Features — FUTURE
- Connection pool health check improvements
- Middleware/interceptor pattern
- Observability integration (metrics, tracing)
- Request/response logging with sensitive data redaction

## Phase 6: Architecture Improvements — FUTURE
- PoolType-aware routing
- Error category improvements (Transient vs Permanent)
- Deprecate `Request()` in favor of `RequestContext()`
