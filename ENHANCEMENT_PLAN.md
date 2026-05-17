# FutuAPI4Go SDK Enhancement Plan

> **Version:** v0.8.5 | **Last Updated:** 2026-05-17 | **Status:** ACTIVE

## Current State Assessment

The SDK has reached **~99% API coverage** with all core features implemented across v0.5.0–v0.8.5:

- 76/78 protos wrapped (97.4%)
- ~200+ typed constants/enums
- 40+ public API functions in client package
- Full context support, connection state machine, graceful shutdown
- Rate limiter, circuit breaker, retry wired into request path
- WebSocket transport, auto-reconnect, TLS support
- OpenTelemetry tracing + metrics (opt-in)
- K-Line data cache, order pre-flight validation, audit logging
- Fluent API: `cli.Quote().GetBasicQot()`, `cli.Trade().PlaceOrder()`
- 107 demo examples, typed push callbacks, mock server for testing

## Remaining Work

### 🔴 Phase 1: True Gaps (Medium Priority)

| # | Item | Description | File(s) |
|---|------|-------------|---------|
| 1a | ~~**GetDelayStatistics**~~ ✅ DONE — Convert from raw proto types + `WritePacket`/`ReadResponseContext` to proper typed wrappers with nil guards. `QotPushDelayStatistics`, `ReqReplyDelayStatistics`, `PlaceOrderDelayStatistics`, `DelayStatisticsItem` structs added to `pkg/sys/system.go`. | `pkg/sys/system.go:224-253` |
| 1b | ~~**GetFlowSummaryResponse**~~ ✅ DONE — Replace `[]*trdflowsummary.FlowSummaryInfo` with wrapped `[]*trd.FlowSummaryInfo` in `pkg/trd/position.go:582`. Wrapper type added at `pkg/trd/position.go:580-592`. Nil guards on all 8 fields. | `pkg/trd/position.go:580-592` |
| 1c | **Audit examples 21-99** — 107 examples in `futuapi4go-demo/examples/`, ~6 verified (StaticInfo, CapitalFlow, MarketState, PushTicker, PushRT, KLine Pe). ~60+ need cross-layer proto field trace for unmapped fields or raw proto leakage. | `futuapi4go-demo/examples/` |

### 🟡 Phase 2: Cleanup & Standardization (Low Priority)

| # | Item | Description | File(s) |
|---|------|-------------|---------|
| 2a | ~~**ProtoID constant consolidation**~~ ✅ DONE — `ProtoID_Qot_GetTradeDate (3225)` and `ProtoID_Qot_GetRehab (3102)` added to `pkg/constant/constant.go`. Removed duplicate local constants from `pkg/qot/trade_date.go` and `pkg/qot/quote.go`. References updated in `pkg/qot/trade_date.go` and `pkg/qot/holding.go`. | `pkg/constant/constant.go:136-137`, `pkg/qot/trade_date.go:13`, `pkg/qot/quote.go:93` |
| 2b | ~~**S2C nil wrapError standardization**~~ ✅ DONE — 50+ `fmt.Errorf("FuncName: s2c is nil")` replaced with `wrapError(..., int32(common.RetType_RetType_Unknown), "s2c is nil")` across all pkg files. Removed unused `fmt` import from `pkg/trd/trade.go`. | 17 files across `pkg/qot/`, `pkg/trd/`, `pkg/sys/` |
| 2c | ~~**Replace GetXxx() with direct nil checks**~~ — Already applied in new wrappers (GetDelayStatistics, GetFlowSummary) and client API layer. Not applied to existing legacy code per codebase convention. | Widespread — deferred |
| 2d | ~~**FutureInfo.TradeTimeList**~~ ✅ N/A — `pkg/qot/options.go` already uses `[]*qotgetfutureinfo.TradeTime` (wrapped at client layer). No change needed. | `pkg/qot/options.go:218` |
| 2e | ~~**IpoData.CnExData/HkExData/UsExData**~~ ✅ N/A — `pkg/qot/user.go` already wraps all three extended data types (CNIpoExData, HKIpoExData, USIpoExData) in `IpoData` struct. No change needed. | `pkg/qot/user.go:398-406` |

### 🟢 Phase 3: Non-Blocking Improvements

| # | Item | Description | Status |
|---|------|-------------|--------|
| 3a | ~~**Demo replace directive**~~ ✅ DONE — Removed `replace github.com/shing1211/futuapi4go => ../futuapi4go` from `futuapi4go-demo/go.mod`. `go mod tidy` resolved all dependencies from v0.8.5 release. | Done |
| 3b | ~~**GitHub release automation**~~ ✅ DONE — `make release` now checks for `goreleaser` and falls back to clear instructions for manual `gh release create`. Gracefully handles missing goreleaser. | Done |
| 3c | ~~**Push.KLine raw proto passthrough**~~ ✅ DONE — Added `PushKLine` struct in `pkg/push/qot_push.go` (13 fields). `UpdateKL.KLList` now uses `[]*PushKLine` instead of `[]*qotcommon.KLine`. Example 07 updated. Tests updated. | Done |

---

## Previously Completed (v0.5.0–v0.8.5)

| Category | Items |
|----------|-------|
| **Proto Field Coverage** | BasicQot (25 fields), KLine (13 fields), Ticker (15 fields), RT (11 fields), Broker (+OrderID), Position (+PositionSide), Funds (+SecuritiesAssets/FundAssets/BondAssets), PushTicker (15 fields), PushRT (11 fields), StaticInfo (9 fields) |
| **System Wrappers** | GetHistoryKLPoints (proto 3106), GetUsedQuota (proto 1010), TestCmd (proto 1008), GetTradeDate (proto 3225), GetRehab (proto 3102), SkillWrapAPI (proto 8001) |
| **Resilience** | Rate limiter wired into request path, retry with exponential backoff, circuit breaker integration, WebSocket auto-reconnect |
| **Lifecycle** | Connection state machine (Disconnected→Connecting→Connected→Reconnecting→Closing), graceful Shutdown(timeout), OnStateChange callback |
| **Observability** | OpenTelemetry metrics bridge (8 sync instruments + 3 gauges), structured slog integration, Prometheus metrics |
| **Production Features** | K-Line LRU cache with TTL, order pre-flight validation (market hours + buying power + max qty), audit/compliance logging, connection pool optimization |
| **Testing** | Mock OpenD server (FT-protocol + RSA/AES), 57+ new unit tests, race detection, integration test framework |
| **API Design** | 30+ fluent API methods, typed push callbacks (OnQuote/OnOrder/OnKLine/etc.), OrderBuilder with trailing stop, GetXxx() on all enums |
| **Constants** | 69 ProtoID constants, 36 typed enum categories, 200+ values, String()/Prefix() methods, validation helpers, SensitiveString |
| **Documentation** | Bilingual USAGE.md (EN/CN), trilingual doc.go (EN/zh-CN/zh-TW), GoDoc on all public APIs, AGENTS.md operational guide |
| **Removed** | Unused ProtoID_GetHistoryKLPoints, unused ProtoID_GetMarketSnapshot alias, dead buffer pool code |

---

## Detailed Pending Item Descriptions

### Item 1a: GetDelayStatistics Typed Wrappers

**Current state** (`pkg/sys/system.go:217-328`):
- Uses raw proto types (`*trddelaystatistics.Request`, `*trddelaystatistics.Response`)
- Manually constructs and sends packets via `WritePacket`/`ReadResponseContext`
- No typed request/response structs like other APIs

**Required change**: Convert to proper typed wrappers following the pattern in `pkg/sys/system.go` GetGlobalState/GetUserInfo:
```go
type GetDelayStatisticsRequest struct { ... }
type GetDelayStatisticsResponse struct {
    ReqReplyList  []*ReqReplyStatisticsItem
    PlaceOrderList []*PlaceOrderStatisticsItem
    QotPushList   []*PushStatisticsItem
}
func GetDelayStatistics(ctx context.Context, c *futuapi.Client, req *GetDelayStatisticsRequest) (*GetDelayStatisticsResponse, error)
```

**Challenge**: Proto2 wire encoding complexity — the `GetDelayStatistics` response uses a complex nested structure that's hard to map with typed structs.

### Item 1b: GetFlowSummaryResponse Wrapped Type

**Current state** (`pkg/trd/position.go:582`):
```go
type GetFlowSummaryResponse struct {
    FlowSummaryList []*trdflowsummary.FlowSummaryInfo  // raw proto
}
```

**Required change**: The wrapped `FlowSummaryInfo` type already exists at `client/types.go:662`:
```go
type GetFlowSummaryResponse struct {
    FlowSummaryList []*FlowSummaryInfo  // wrapped type
}
```

### Item 1c: Example Audit

**Scope**: 107 examples in `futuapi4go-demo/examples/`, numbered 00 through 99. ~6 verified against proto fields in v0.8.3 [Unreleased] work.

**Method**: For each example, run a cross-layer trace:
1. Call the wrapper function
2. Check response type fields against proto fields
3. Identify unmapped fields or raw proto leakage
4. Document or fix any gaps found

**Already verified**:
- `59_static_info` — StaticInfo fields (additional Verified)
- `12_capital_flow` — CapitalFlow metadata (Verified)
- `16_market_state` — MarketState Code/Name (Verified)
- `02_ticker` — PushTicker fields (Verified)
- `04_rt` — PushRT fields (Verified)
- `06_kline_single`, `07_kline_multi` — KLine Pe (Verified)

---

*Last updated: 2026-05-17*
