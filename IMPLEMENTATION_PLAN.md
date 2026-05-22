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

## Phase 4: Missing High-Level API Coverage — DONE

### Implemented (25 new quote APIs + 34 new client wrappers)

**9 new pkg/qot/ files** with 25 API implementations:
- `pkg/qot/financials.go`: GetFinancialsStatements (3227), GetFinancialsRevenueBreakdown (3228)
- `pkg/qot/research.go`: GetResearchAnalystConsensus (3229), GetResearchRatingSummary (3230), GetResearchMorningstarReport (3231)
- `pkg/qot/valuation.go`: GetValuationDetail (3232), GetValuationPlateStockList (3233)
- `pkg/qot/corporate.go`: GetCorporateActionsDividends (3234), GetCorporateActionsBuybacks (3235), GetCorporateActionsStockSplits (3236)
- `pkg/qot/shareholders.go`: GetShareholdersOverview (3237), GetShareholdersHoldingChanges (3238), GetShareholdersHolderDetail (3239), GetShareholdersInstitutional (3240)
- `pkg/qot/insider.go`: GetInsiderHolderList (3241), GetInsiderTradeList (3242)
- `pkg/qot/company.go`: GetCompanyProfile (3243), GetCompanyExecutives (3244), GetCompanyExecutiveBackground (3245), GetCompanyOperationalEfficiency (3246)
- `pkg/qot/shortselling.go`: GetTopTenBuySellBrokers (3247), GetDailyShortVolume (3248), GetShortInterest (3249)
- `pkg/qot/option_extra.go`: GetOptionVolatility (3250), GetOptionExerciseProbability (3251)

**Client layer additions**:
- 26 new `client/quote_api.go` wrappers (25 new + 1 missing existing: RequestHistoryKLQuota)
- 26 new `client/fluent_api.go` QuoteAPI methods
- 8 new `client/trade_api.go` convenience wrappers (QuickBuy, QuickSell, QuickMarketBuy, QuickMarketSell, GetPositions, GetTodayFills, GetTodayOrders, GetAccountFunds)

**Deferred** (5 screening APIs — no ProtoID assigned yet):
- Qot_StockScreen.proto, Qot_WarrantScreen.proto, Qot_OptionScreen.proto
- Qot_GetFinancialsEarningsPriceMove.proto, Qot_GetFinancialsEarningsPriceHistory.proto

### Old Phase 4 text (superseded):
- Add `client/` wrappers for `RequestHistoryKLQuota`, `GetUserSecurityGroup`, etc.
- Add fluent API methods for missing functions

## Phase 5: Bug Fixes, Missing APIs & Architecture Hardening — DONE

*Detailed plan: PHASE5_BUGFIX_HARDENING_PLAN.md*

### Step 1: Fix ProtoID Mismatch (P0 — CRITICAL)
- `ProtoID_Qot_GetTradeDate = 3225` was WRONG — 3225 = `Qot_GetFinancialsEarningsPriceMove`
- Removed duplicate `ProtoID_Qot_GetTradeDate`, updated `trade_date.go` to use `ProtoID_Qot_RequestTradeDate` (3219) with official `Qot_RequestTradeDate` proto types
- Added `ProtoID_Qot_GetFinancialsEarningsPriceMove = 3225` and `ProtoID_Qot_GetFinancialsEarningsPriceHistory = 3226`

### Step 2: Implement 2 Missing APIs (P0)
- `GetFinancialsEarningsPriceMove` (ProtoID 3225) — earnings price move data with 15+ fields
- `GetFinancialsEarningsPriceHistory` (ProtoID 3226) — earnings price history data with 20+ fields incl. volatility metrics, option IV crush
- Added proto wrappers, client wrappers, fluent API methods

### Step 3: Fix Proto Safety Violations (P1)
- Fixed 7 violations in `pkg/qot/` (trade_date, option_extra, shortselling, options, market_data)
- Fixed 2 violations in `pkg/push/` (qot_push)
- Fixed 20+ violations in `internal/client/client.go` (InitConnect/keepAlive sections)

### Step 4: Fix Concurrency & Nil Deref Bugs (P1)
- Logger race in `New()` — now uses `SetLogger()` with mutex protection
- `SetTracer()` race — switched to `sync/atomic.Value`
- `Conn.LocalAddr()/RemoteAddr()` — nil dereference when `c.conn == nil`

### Step 5: Fix AES Padding, Breaker, Pool Contention (P2)
- `aesCBCDecrypt` now strips PKCS#7 padding after decryption
- Breaker `halfOpenMax` now enforced in `Allow()` with `halfOpenInFlight` counter
- Pool contention — identified but deferred (complex TCP dial refactor)

### Step 6: Dead Code / Unwired Packages (P2)
- Documented in PHASE5 plan; integration deferred as non-critical

### Step 7: Middleware/Interceptor Pattern (P3)
- Deferred to future phase; RequestContext refactoring non-trivial

### Step 8: Documentation & Release
- Updated CHANGELOG.md with all Phase 5 items
- Updated IMPLEMENTATION_PLAN.md
- Committed and pushed to origin/main and gitee/main

## Phase 6: Architecture Improvements — FUTURE
- PoolType-aware routing
- Error category improvements (Transient vs Permanent)
- Deprecate `Request()` in favor of `RequestContext()`
