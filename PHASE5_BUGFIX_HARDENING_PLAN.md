# Phase 5: Bug Fixes, Missing APIs & Architecture Hardening

*Generated: 2026-05-22 | Status: COMPLETE (implemented in v0.12.0+; extended through v0.15.1)*

## Executive Summary

Comprehensive audit against official Futu API v10.8.6808 documentation revealed **1 critical ProtoID mismatch**, **2 missing APIs**, **10+ proto safety violations**, **3 concurrency bugs**, **1 nil dereference risk**, **1 AES padding bug**, **1 breaker logic bug**, and **3 unwired architecture packages**. This phase addresses all findings.

---

## Official Futu API v10.6 Protocol ID Reference

Source: https://openapi.futunn.com/futu-api-doc/en/quote/overview.html

### Quote APIs (ProtoIDs 3001-3251)

| ProtoID | Official API Name | Our Constant | Our Value | Status |
|---------|-------------------|--------------|-----------|--------|
| 3001 | Qot_Sub | ProtoID_Qot_Sub | 3001 | OK |
| 3002 | Qot_RegQotPush | ProtoID_Qot_RegQotPush | 3002 | OK |
| 3003 | Qot_GetSubInfo | ProtoID_Qot_GetSubInfo | 3003 | OK |
| 3004 | Qot_GetBasicQot | ProtoID_Qot_GetBasicQot | 3004 | OK |
| 3005 | Qot_UpdateBasicQot | ProtoID_Qot_UpdateBasicQot | 3005 | OK |
| 3006 | Qot_GetKL | ProtoID_Qot_GetKL | 3006 | OK |
| 3007 | Qot_UpdateKL | ProtoID_Qot_UpdateKL | 3007 | OK |
| 3008 | Qot_GetRT | ProtoID_Qot_GetRT | 3008 | OK |
| 3009 | Qot_UpdateRT | ProtoID_Qot_UpdateRT | 3009 | OK |
| 3010 | Qot_GetTicker | ProtoID_Qot_GetTicker | 3010 | OK |
| 3011 | Qot_UpdateTicker | ProtoID_Qot_UpdateTicker | 3011 | OK |
| 3012 | Qot_GetOrderBook | ProtoID_Qot_GetOrderBook | 3012 | OK |
| 3013 | Qot_UpdateOrderBook | ProtoID_Qot_UpdateOrderBook | 3013 | OK |
| 3014 | Qot_GetBroker | ProtoID_Qot_GetBroker | 3014 | OK |
| 3015 | Qot_UpdateBroker | ProtoID_Qot_UpdateBroker | 3015 | OK |
| 3019 | Qot_UpdatePriceReminder | ProtoID_Qot_UpdatePriceReminder | 3019 | OK |
| 3102 | Qot_RequestRehab | ProtoID_Qot_GetRehab | 3102 | OK (name differs) |
| 3103 | Qot_RequestHistoryKL | ProtoID_Qot_RequestHistoryKL | 3103 | OK |
| 3104 | Qot_RequestHistoryKLQuota | ProtoID_Qot_RequestHistoryKLQuota | 3104 | OK |
| 3105 | Qot_RequestRehab | ProtoID_Qot_RequestRehab | 3105 | **DUPLICATE** - 3102 and 3105 both exist |
| 3106 | Qot_GetHistoryKLPoints | ProtoID_Qot_GetHistoryKLPoints | 3106 | OK |
| 3201 | Qot_GetSuspend | ProtoID_Qot_GetSuspend | 3201 | OK |
| 3202 | Qot_GetStaticInfo | ProtoID_Qot_GetStaticInfo | 3202 | OK |
| 3203 | Qot_GetSecuritySnapshot | ProtoID_Qot_GetSecuritySnapshot | 3203 | OK |
| 3204 | Qot_GetPlateSet | ProtoID_Qot_GetPlateSet | 3204 | OK |
| 3205 | Qot_GetPlateSecurity | ProtoID_Qot_GetPlateSecurity | 3205 | OK |
| 3206 | Qot_GetReference | ProtoID_Qot_GetReference | 3206 | OK |
| 3207 | Qot_GetOwnerPlate | ProtoID_Qot_GetOwnerPlate | 3207 | OK |
| 3208 | Qot_GetHoldingChangeList | ProtoID_Qot_GetHoldingChangeList | 3208 | OK |
| 3209 | Qot_GetOptionChain | ProtoID_Qot_GetOptionChain | 3209 | OK |
| 3210 | Qot_GetWarrant | ProtoID_Qot_GetWarrant | 3210 | OK |
| 3211 | Qot_GetCapitalFlow | ProtoID_Qot_GetCapitalFlow | 3211 | OK |
| 3212 | Qot_GetCapitalDistribution | ProtoID_Qot_GetCapitalDistribution | 3212 | OK |
| 3213 | Qot_GetUserSecurity | ProtoID_Qot_GetUserSecurity | 3213 | OK |
| 3214 | Qot_ModifyUserSecurity | ProtoID_Qot_ModifyUserSecurity | 3214 | OK |
| 3215 | Qot_StockFilter | ProtoID_Qot_StockFilter | 3215 | OK |
| 3216 | Qot_GetCodeChange | ProtoID_Qot_GetCodeChange | 3216 | OK |
| 3217 | Qot_GetIpoList | ProtoID_Qot_GetIpoList | 3217 | OK |
| 3218 | Qot_GetFutureInfo | ProtoID_Qot_GetFutureInfo | 3218 | OK |
| 3219 | Qot_RequestTradeDate | ProtoID_Qot_RequestTradeDate | 3219 | OK |
| 3220 | Qot_SetPriceReminder | ProtoID_Qot_SetPriceReminder | 3220 | OK |
| 3221 | Qot_GetPriceReminder | ProtoID_Qot_GetPriceReminder | 3221 | OK |
| 3222 | Qot_GetUserSecurityGroup | ProtoID_Qot_GetUserSecurityGroup | 3222 | OK |
| 3223 | Qot_GetMarketState | ProtoID_Qot_GetMarketState | 3223 | OK |
| 3224 | Qot_GetOptionExpirationDate | ProtoID_Qot_GetOptionExpirationDate | 3224 | OK |
| **3225** | **Qot_GetFinancialsEarningsPriceMove** | ProtoID_Qot_GetTradeDate | 3225 | **CRITICAL BUG** |
| **3226** | **Qot_GetFinancialsEarningsPriceHistory** | *(missing)* | - | **MISSING** |
| 3227 | Qot_GetFinancialsStatements | ProtoID_Qot_GetFinancialsStatements | 3227 | OK |
| 3228 | Qot_GetFinancialsRevenueBreakdown | ProtoID_Qot_GetFinancialsRevenueBreakdown | 3228 | OK |
| 3229 | Qot_GetResearchAnalystConsensus | ProtoID_Qot_GetResearchAnalystConsensus | 3229 | OK |
| 3230 | Qot_GetResearchRatingSummary | ProtoID_Qot_GetResearchRatingSummary | 3230 | OK |
| 3231 | Qot_GetResearchMorningstarReport | ProtoID_Qot_GetResearchMorningstarReport | 3231 | OK |
| 3232 | Qot_GetValuationDetail | ProtoID_Qot_GetValuationDetail | 3232 | OK |
| 3233 | Qot_GetValuationPlateStockList | ProtoID_Qot_GetValuationPlateStockList | 3233 | OK |
| 3234 | Qot_GetCorporateActionsDividends | ProtoID_Qot_GetCorporateActionsDividends | 3234 | OK |
| 3235 | Qot_GetCorporateActionsBuybacks | ProtoID_Qot_GetCorporateActionsBuybacks | 3235 | OK |
| 3236 | Qot_GetCorporateActionsStockSplits | ProtoID_Qot_GetCorporateActionsStockSplits | 3236 | OK |
| 3237 | Qot_GetShareholdersOverview | ProtoID_Qot_GetShareholdersOverview | 3237 | OK |
| 3238 | Qot_GetShareholdersHoldingChanges | ProtoID_Qot_GetShareholdersHoldingChanges | 3238 | OK |
| 3239 | Qot_GetShareholdersHolderDetail | ProtoID_Qot_GetShareholdersHolderDetail | 3239 | OK |
| 3240 | Qot_GetShareholdersInstitutional | ProtoID_Qot_GetShareholdersInstitutional | 3240 | OK |
| 3241 | Qot_GetInsiderHolderList | ProtoID_Qot_GetInsiderHolderList | 3241 | OK |
| 3242 | Qot_GetInsiderTradeList | ProtoID_Qot_GetInsiderTradeList | 3242 | OK |
| 3243 | Qot_GetCompanyProfile | ProtoID_Qot_GetCompanyProfile | 3243 | OK |
| 3244 | Qot_GetCompanyExecutives | ProtoID_Qot_GetCompanyExecutives | 3244 | OK |
| 3245 | Qot_GetCompanyExecutiveBackground | ProtoID_Qot_GetCompanyExecutiveBackground | 3245 | OK |
| 3246 | Qot_GetCompanyOperationalEfficiency | ProtoID_Qot_GetCompanyOperationalEfficiency | 3246 | OK |
| 3247 | Qot_GetTopTenBuySellBrokers | ProtoID_Qot_GetTopTenBuySellBrokers | 3247 | OK |
| 3248 | Qot_GetDailyShortVolume | ProtoID_Qot_GetDailyShortVolume | 3248 | OK |
| 3249 | Qot_GetShortInterest | ProtoID_Qot_GetShortInterest | 3249 | OK |
| 3250 | Qot_GetOptionVolatility | ProtoID_Qot_GetOptionVolatility | 3250 | OK |
| 3251 | Qot_GetOptionExerciseProbability | ProtoID_Qot_GetOptionExerciseProbability | 3251 | OK |

> **Note:** This audit was conducted against Futu API v10.6. Subsequent protocol upgrades extended the quote-space ProtoID range. The ranges below are now **all wrapped** and tested:
> - **v10.8 (3252–3364)** — 56 new APIs (search, indicators, options analytics, rankings, institutional, chain, heatmap, market fundamentals) implemented in SDK v0.14.0
> - **v10.9 (3434–3456)** — 17 Event Contract / Prediction Market APIs implemented in SDK v0.15.0, plus 2 backfill push parsers (`ParseUpdateOptionEvent` 3310, `ParsePushIndicatorCalc` 3261) and 6 EC chanpkg wrappers in SDK v0.15.1

### Trade APIs (ProtoIDs 2001-2226)

| ProtoID | Official API Name | Our Constant | Status |
|---------|-------------------|--------------|--------|
| 2001 | Trd_GetAccList | ProtoID_Trd_GetAccList | OK |
| 2005 | Trd_UnlockTrade | ProtoID_Trd_UnlockTrade | OK |
| 2008 | Trd_SubAccPush | ProtoID_Trd_SubAccPush | OK |
| 2101 | Trd_GetFunds | ProtoID_Trd_GetFunds | OK |
| 2102 | Trd_GetPositionList | ProtoID_Trd_GetPositionList | OK |
| 2111 | Trd_GetMaxTrdQtys | ProtoID_Trd_GetMaxTrdQtys | OK |
| 2201 | Trd_GetOrderList | ProtoID_Trd_GetOrderList | OK |
| 2202 | Trd_PlaceOrder | ProtoID_Trd_PlaceOrder | OK |
| 2205 | Trd_ModifyOrder | ProtoID_Trd_ModifyOrder | OK |
| 2208 | Trd_UpdateOrder | ProtoID_Trd_UpdateOrder | OK |
| 2211 | Trd_GetOrderFillList | ProtoID_Trd_GetOrderFillList | OK |
| 2218 | Trd_UpdateOrderFill | ProtoID_Trd_UpdateOrderFill | OK |
| 2221 | Trd_GetHistoryOrderList | ProtoID_Trd_GetHistoryOrderList | OK |
| 2222 | Trd_GetHistoryOrderFillList | ProtoID_Trd_GetHistoryOrderFillList | OK |
| 2223 | Trd_GetMarginRatio | ProtoID_Trd_GetMarginRatio | OK |
| 2225 | Trd_GetOrderFee | ProtoID_Trd_GetOrderFee | OK |
| 2226 | Trd_GetAccCashFlow | ProtoID_Trd_FlowSummary | OK (name differs) |

### Basic APIs (ProtoIDs 1001-1010)

| ProtoID | Official API Name | Our Constant | Status |
|---------|-------------------|--------------|--------|
| 1001 | InitConnect | ProtoID_InitConnect | OK |
| 1002 | GetGlobalState | ProtoID_GetGlobalState | OK |
| 1003 | Notify | ProtoID_Notify | OK |
| 1004 | KeepAlive | ProtoID_KeepAlive | OK |
| 1005 | GetUserInfo | ProtoID_GetUserInfo | OK |
| 1006 | Verification | ProtoID_Verification | OK |
| 1007 | GetDelayStatistics | ProtoID_GetDelayStatistics | OK |
| 1008 | TestCmd | ProtoID_TestCmd | OK |
| 1010 | UsedQuota | ProtoID_UsedQuota | OK |

---

## Step 1: Fix ProtoID Mismatch (P0 — CRITICAL)

### 1.1 Problem

`pkg/constant/constant.go:135` defines:
```go
ProtoID_Qot_GetTradeDate = 3225 // 获取交易日
```

But the official Futu API v10.6 assigns ProtoID 3225 to `Qot_GetFinancialsEarningsPriceMove`, NOT to GetTradeDate.

The official API has only `Qot_RequestTradeDate` (ProtoID 3219) for the trading calendar. Our `ProtoID_Qot_GetTradeDate` is a **duplicate with the wrong ProtoID**.

### 1.2 Impact

- `pkg/qot/trade_date.go:42` calls `GetTradeDate()` with ProtoID 3225, which OpenD interprets as `GetFinancialsEarningsPriceMove` — causing protocol mismatch errors or wrong data
- ProtoIDs 3225 and 3226 are not available for the actual `GetFinancialsEarningsPriceMove` and `GetFinancialsEarningsPriceHistory` APIs

### 1.3 Fix

1. **Remove** `ProtoID_Qot_GetTradeDate = 3225` from `pkg/constant/constant.go`
2. **Add** two new constants:
   ```go
   ProtoID_Qot_GetFinancialsEarningsPriceMove    = 3225 // 获取财报价格变动
   ProtoID_Qot_GetFinancialsEarningsPriceHistory = 3226 // 获取财报价格历史
   ```
3. **Update** `pkg/qot/trade_date.go` to use `constant.ProtoID_Qot_RequestTradeDate` (3219) instead of `constant.ProtoID_Qot_GetTradeDate`
4. **Update** `client/quote_api.go` — rename `GetTradeDates` to use `RequestTradeDate` internally
5. **Update** `client/fluent_api.go` — same
6. **Move** `ProtoID_Qot_GetRehab = 3102` to proper sequential position (currently placed after 3225)

### 1.4 Files Changed

- `pkg/constant/constant.go` — remove wrong constant, add 2 new, reorder
- `pkg/qot/trade_date.go` — change ProtoID reference
- `client/quote_api.go` — update GetTradeDates wrapper
- `client/fluent_api.go` — update QuoteAPI method

---

## Step 2: Implement 2 Missing APIs (P0)

### 2.1 GetFinancialsEarningsPriceMove (ProtoID 3225)

**Official docs:** https://openapi.futunn.com/futu-api-doc/en/quote/get-financials-earnings-price-move.html

**Proto definition:**
```protobuf
message C2S {
    required Qot_Common.Security security = 1;
    optional int32 periodCount = 2; // default 10, range [1, 50]
}

message PricePerformanceRow {
    optional int64  tradingDay       = 1;
    optional string tradingDayStr    = 2;
    optional double closePrice       = 3;
    optional double openPrice        = 4;
    optional double highestPrice     = 5;
    optional double lowestPrice      = 6;
    optional double lastClosePrice   = 7;
    optional double optionIV         = 8;
    optional double optionHV         = 9;
}

message ReportCycleQuote {
    optional int32  fiscalYear        = 1;
    optional int32  financialType     = 2;
    optional string periodText        = 3;
    optional int64  pubTradingDay     = 4;
    optional string pubTradingDayStr  = 5;
    optional Qot_Common.EarningsPubTimeType pubType = 6;
    optional int32  priceInfoIndex    = 7;
    repeated PricePerformanceRow itemList = 8;
}

message S2C {
    repeated ReportCycleQuote detailList = 1;
}
```

**Rate limit:** Max 30 requests per 30 seconds. Supports HK and US equities only.

### 2.2 GetFinancialsEarningsPriceHistory (ProtoID 3226)

**Official docs:** https://openapi.futunn.com/futu-api-doc/en/quote/get-financials-earnings-price-history.html

**Proto definition:**
```protobuf
message C2S {
    required Qot_Common.Security security = 1;
}

message PriceInfo {
    optional int64  tradingDay       = 1;
    optional string tradingDayStr    = 2;
    optional double closePrice       = 3;
    optional double openPrice        = 4;
    optional double highestPrice     = 5;
    optional double lowestPrice      = 6;
    optional double lastClosePrice   = 7;
    optional double volume           = 8;
}

message FinScheduleInfo {
    optional int32  delta      = 1;
    optional double closePrice = 2;
}

message PriceHistoryOnEarningsDays {
    optional int32  fiscalYear              = 1;
    optional int32  financialType           = 2;
    optional string periodText              = 3;
    optional bool   isCurrent               = 4;
    optional int64  pubTradingDay           = 5;
    optional string pubTradingDayStr        = 6;
    optional int64  pubTime                 = 7;
    optional string pubTimeStr              = 8;
    optional Qot_Common.EarningsPubTimeType pubType    = 9;
    optional double predictVolaRatioNewest  = 10;
    optional double predictVolaRatioHighest = 11;
    optional double predictVolaValNewest    = 12;
    optional double predictVolaValHighest   = 13;
    optional double optionIVCrush           = 14;
    optional double optionStrikeDateIVCrush = 15;
    optional PriceInfo priceInfo              = 16;
    repeated FinScheduleInfo scheduleInfoList = 17;
}

message S2C {
    repeated PriceHistoryOnEarningsDays detailList = 1;
}
```

**Rate limit:** Max 30 requests per 30 seconds. Supports HK and US equities only.

### 2.3 Implementation Plan

1. Check `api/proto/` for existing `.proto` files for these APIs
2. If missing, create `Qot_GetFinancialsEarningsPriceMove.proto` and `Qot_GetFinancialsEarningsPriceHistory.proto`
3. Run `./scripts/regen-all-protos.sh` to generate Go code
4. Add wrapper functions to `pkg/qot/financials.go`:
   - `GetFinancialsEarningsPriceMove(ctx, client, req) (*GetFinancialsEarningsPriceMoveResponse, error)`
   - `GetFinancialsEarningsPriceHistory(ctx, client, req) (*GetFinancialsEarningsPriceHistoryResponse, error)`
5. Add client wrappers in `client/quote_api.go`
6. Add fluent API methods in `client/fluent_api.go`
7. Add unit tests

### 2.4 Files Changed

- `api/proto/Qot_GetFinancialsEarningsPriceMove.proto` — new
- `api/proto/Qot_GetFinancialsEarningsPriceHistory.proto` — new
- `pkg/pb/qotgetfinancialsearningspricemove/` — generated
- `pkg/pb/qotgetfinancialsearningspricehistory/` — generated
- `pkg/qot/financials.go` — add 2 new functions
- `client/quote_api.go` — add 2 new wrappers
- `client/fluent_api.go` — add 2 new methods

---

## Step 3: Fix Remaining Proto Safety Violations (P1)

### 3.1 Violations in pkg/qot/

| File | Line | Violation | Fix |
|------|------|-----------|-----|
| `trade_date.go` | 65 | `td.GetTradeDateType()` | `util.ProtoInt32(td.TradeDateType)` |
| `option_extra.go` | 87 | `s2c.GetImpvolStatus()` | `util.ProtoInt32(s2c.ImpvolStatus)` |
| `shortselling.go` | 101 | `b.GetBuySellType()` | `util.ProtoInt32(b.BuySellType)` |
| `options.go` | 165 | `s2c.GetOptionChain()` | `s2c.OptionChain` |
| `options.go` | 168 | `s2c.GetOptionChain()` | `s2c.OptionChain` |
| `options.go` | 282 | `fi.GetTradeTime()` | `fi.TradeTime` |
| `market_data.go` | 224 | `t.GetRecvTime()` | `util.ProtoInt64(t.RecvTime)` |

### 3.2 Violations in pkg/push/

| File | Line | Violation | Fix |
|------|------|-----------|-----|
| `trd_push.go` | 64 | `s2c.GetEvent()` | `s2c.Event` (enum-typed, keep as getter) |
| `trd_push.go` | 66 | `s2c.GetConnectStatus()` | `s2c.ConnectStatus` (enum-typed) |
| `trd_push.go` | 68 | `s2c.GetApiLevel()` | `s2c.ApiLevel` (enum-typed) |
| `trd_push.go` | 69 | `s2c.GetApiQuota()` | `s2c.ApiQuota` (enum-typed) |
| `trd_push.go` | 70 | `s2c.GetUsedQuota()` | `s2c.UsedQuota` (enum-typed) |
| `qot_push.go` | 357 | `s2c.GetSetValue()` | `util.ProtoFloat64(s2c.SetValue)` |
| `qot_push.go` | 358 | `s2c.GetCurValue()` | `util.ProtoFloat64(s2c.CurValue)` |

**Note on enum-typed getters:** Per AGENTS.md convention, enum-typed proto getters (GetEvent, GetConnectStatus, GetApiLevel, GetApiQuota, GetUsedQuota) return named types, not dereferenced scalars. These are borderline — they return zero-value for nil fields but the type system prevents most misuse. Decision: **keep enum getters** for now, fix scalar getters only.

### 3.3 Violations in internal/client/client.go

| Line | Violation | Fix |
|------|-----------|-----|
| 598 | `rsp.GetRetType()` | `util.ProtoInt32(rsp.RetType)` |
| 600 | `rsp.GetRetMsg()` | `util.ProtoStr(rsp.RetMsg)` |
| 603 | `rsp.GetS2C()` | `rsp.S2C` |
| 610 | `s2c.GetConnID()` | `util.ProtoUint64(s2c.ConnID)` |
| 611 | `s2c.GetLoginUserID()` | `util.ProtoUint64(s2c.LoginUserID)` |
| 612 | `s2c.GetKeepAliveInterval()` | `util.ProtoInt32(s2c.KeepAliveInterval)` |
| 613 | `s2c.GetServerVer()` | `util.ProtoInt32(s2c.ServerVer)` |
| 801 | `rsp.GetRetType()` | `util.ProtoInt32(rsp.RetType)` |
| 803 | `rsp.GetRetMsg()` | `util.ProtoStr(rsp.RetMsg)` |
| 809 | `rsp.GetS2C()` | `rsp.S2C` |
| 817 | `s2c.GetConnID()` | `util.ProtoUint64(s2c.ConnID)` |
| 818 | `s2c.GetLoginUserID()` | `util.ProtoUint64(s2c.LoginUserID)` |
| 819 | `s2c.GetConnAESKey()` | `s2c.ConnAESKey` (byte slice, safe) |
| 820 | `s2c.GetAesCBCiv()` | `s2c.AesCBCiv` (byte slice, safe) |
| 821 | `s2c.GetServerVer()` | `util.ProtoInt32(s2c.ServerVer)` |
| 822 | `s2c.GetKeepAliveInterval()` | `util.ProtoInt32(s2c.KeepAliveInterval)` |
| 920 | `rsp.GetRetType()` | `util.ProtoInt32(rsp.RetType)` |
| 921 | `rsp.GetRetMsg()` | `util.ProtoStr(rsp.RetMsg)` |

### 3.4 Files Changed

- `pkg/qot/trade_date.go`
- `pkg/qot/option_extra.go`
- `pkg/qot/shortselling.go`
- `pkg/qot/options.go`
- `pkg/qot/market_data.go`
- `pkg/push/qot_push.go`
- `internal/client/client.go`

---

## Step 4: Fix Concurrency & Nil Dereference Bugs (P1)

### 4.1 Logger Race in New()

**File:** `internal/client/client.go:475`

```go
func New(opts ...Option) *Client {
    options := NewOptions()
    for _, opt := range opts {
        opt(options)
    }
    logger = options.Logger  // RACE: writes global without lock
    ...
}
```

**Fix:** Use `SetLogger()` which properly acquires `loggerMu`:
```go
if options.Logger != nil {
    SetLogger(options.Logger)
}
```

### 4.2 SetTracer() Race

**File:** `pkg/tracing/tracing.go:39-43`

```go
var defaultTracer Tracer = NoopTracer{}

func SetTracer(t Tracer) {
    if t != nil {
        defaultTracer = t  // RACE: no mutex, no atomic
    }
}
```

**Fix:** Use `sync/atomic.Value`:
```go
var defaultTracer atomic.Value

func init() {
    defaultTracer.Store(NoopTracer{})
}

func SetTracer(t Tracer) {
    if t != nil {
        defaultTracer.Store(t)
    }
}

func getTracer() Tracer {
    return defaultTracer.Load().(Tracer)
}
```

### 4.3 SetRateLimiter/SetRetryConfig/SetBreaker/SetWSSecretKey Not Thread-Safe

**File:** `internal/client/client.go:314-351`

These setters modify shared state without acquiring `c.mu`. Meanwhile, `RequestContext()` reads these fields without locks.

**Fix:** Acquire `c.mu` in each setter:
```go
func (c *Client) SetRateLimiter(rl *ratelimit.ProtoLimiter) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.rateLimiter = rl
}
```

### 4.4 Conn.LocalAddr()/RemoteAddr() Nil Dereference

**File:** `internal/client/conn.go:153-159`

```go
func (c *Conn) LocalAddr() net.Addr {
    return c.conn.LocalAddr()  // PANIC if c.conn is nil
}
```

**Fix:** Add nil guard:
```go
func (c *Conn) LocalAddr() net.Addr {
    c.mu.Lock()
    conn := c.conn
    c.mu.Unlock()
    if conn == nil {
        return nil
    }
    return conn.LocalAddr()
}
```

### 4.5 Files Changed

- `internal/client/client.go`
- `pkg/tracing/tracing.go`
- `internal/client/conn.go`

---

## Step 5: Fix AES Padding, Breaker, Pool Contention (P2)

### 5.1 AES CBC PKCS#7 Unpadding Missing

**File:** `internal/client/aes.go:123-141`

`aesCBCEncrypt` applies PKCS#7 padding but `aesCBCDecrypt` does NOT strip it.

**Fix:** Add PKCS#7 unpadding after decryption:
```go
func aesCBCDecrypt(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
    ...
    mode.CryptBlocks(plaintext, ciphertext)

    // Strip PKCS#7 padding
    if len(plaintext) > 0 {
        padLen := int(plaintext[len(plaintext)-1])
        if padLen > 0 && padLen <= 16 && padLen <= len(plaintext) {
            plaintext = plaintext[:len(plaintext)-padLen]
        }
    }

    return plaintext, nil
}
```

### 5.2 Breaker halfOpenMax Never Enforced

**File:** `pkg/breaker/breaker.go:168-169`

```go
case StateHalfOpen:
    return true  // Always allows, ignores halfOpenMax
```

**Fix:** Track half-open in-flight count and enforce limit:
```go
type Breaker struct {
    ...
    halfOpenInFlight int32
}

case StateHalfOpen:
    if atomic.LoadInt32(&b.halfOpenInFlight) < int32(b.halfOpenMax) {
        atomic.AddInt32(&b.halfOpenInFlight, 1)
        return true
    }
    return false
```

Also decrement in `RecordSuccess()` and `RecordFailure()` when state is half-open.

### 5.3 Pool Contention: Mutex Held During TCP Dial

**File:** `internal/client/pool.go:103-152`

`Get()` holds `p.mu` while calling `newClientLocked()`, which performs TCP dial + InitConnect handshake (100ms-5s).

**Fix:** Create client outside the lock, then acquire lock only to add to pool:
```go
func (p *ClientPool) Get(ctx context.Context, poolType PoolType) (*Client, error) {
    p.mu.Lock()
    // Check for available connection...
    if found {
        p.mu.Unlock()
        return client, nil
    }
    needNew := len(conns) < p.config.MaxSize
    p.mu.Unlock()

    if needNew {
        client, err := p.newClient()  // Dial outside lock
        if err != nil {
            return nil, err
        }
        p.mu.Lock()
        // Add to pool...
        p.mu.Unlock()
        return client, nil
    }
    ...
}
```

### 5.4 Files Changed

- `internal/client/aes.go`
- `pkg/breaker/breaker.go`
- `internal/client/pool.go`

---

## Step 6: Dead Code / Unwired Packages (P2)

### 6.1 pkg/degradation/ — Zero Integration

Fully implemented `Manager` with `SetStatus()`, `GetStatus()`, `IsDegraded()`, `AddWatcher()`, `AllStatus()`. No other package imports it.

**Fix:** Wire into client lifecycle:
- Auto-trigger degradation on breaker open / reconnect failure
- Integrate with health check
- Add `WithDegradationManager()` option

### 6.2 pkg/metrics/ — Partially Integrated

`RecordAPICall()`, `RecordReconnect()`, etc. are called from `internal/client/client.go`. But:
- `RecordRateLimited()` — never called (rate limiter doesn't report)
- `RecordRetry()` — never called (retry doesn't report)
- `RecordBreakerState()` — never called (breaker doesn't report)
- `StartAPITracking()` / `APICallTracker.End()` — never used
- `Init()` / `InitWithServer()` — never called (Prometheus endpoint never started)

**Fix:**
- Wire `RecordRateLimited()` into rate limiter rejection path
- Wire `RecordRetry()` into retry path
- Wire `RecordBreakerState()` into breaker state transitions
- Document `InitWithServer()` in README

### 6.3 pkg/tracing/otel/ — Never Imported

OpenTelemetry backend is fully implemented but never imported by any SDK code.

**Fix:** Document usage in `pkg/tracing/doc.go` and README. Add example.

### 6.4 Files Changed

- `internal/client/client.go` — wire degradation, metrics
- `pkg/breaker/breaker.go` — add metrics callback
- `pkg/retry/retry.go` — add metrics callback
- `pkg/ratelimit/ratelimit.go` — add metrics callback
- `docs/USAGE.md` — document InitWithServer, otel tracing

---

## Step 7: Middleware/Interceptor Pattern (P3)

### 7.1 Current Architecture

Breaker, retry, and rate limiter are hardcoded into `RequestContext()`:
```go
func (c *Client) requestContextInternal(...) {
    if c.breaker != nil && !c.breaker.Allow() { ... }
    if c.rateLimiter != nil { c.rateLimiter.Wait(protoID) }
    if c.retryConfig != nil && !isTradingProto(protoID) {
        retry.Do(func() error { ... })
    }
}
```

### 7.2 Proposed Interceptor Pattern

```go
type RequestInterceptor func(ctx context.Context, protoID int32, req proto.Message, handler RequestHandler) (proto.Message, error)
type RequestHandler func(ctx context.Context, protoID int32, req proto.Message) (proto.Message, error)

func BreakerInterceptor(b *breaker.Breaker) RequestInterceptor { ... }
func RetryInterceptor(cfg retry.Config, excludeProtos ...int32) RequestInterceptor { ... }
func RateLimitInterceptor(rl *ratelimit.ProtoLimiter) RequestInterceptor { ... }
func MetricsInterceptor(m *metrics.Recorder) RequestInterceptor { ... }
func TracingInterceptor(t tracing.Tracer) RequestInterceptor { ... }
func LoggingInterceptor(l *slog.Logger) RequestInterceptor { ... }
func DegradationInterceptor(dm *degradation.Manager) RequestInterceptor { ... }
```

### 7.3 Benefits

- Composable: users can add/remove interceptors
- Orderable: interceptors execute in chain order
- Testable: each interceptor is independently testable
- Extensible: users can write custom interceptors

### 7.4 Files Changed

- `internal/client/interceptor.go` — new file
- `internal/client/client.go` — refactor RequestContext to use interceptor chain
- `pkg/breaker/interceptor.go` — new
- `pkg/retry/interceptor.go` — new
- `pkg/ratelimit/interceptor.go` — new
- `pkg/metrics/interceptor.go` — new
- `pkg/tracing/interceptor.go` — new

---

## Step 8: Error Wrapping Consistency (P3)

### 8.1 Current State

Mixed error wrapping patterns:
- `pkg/qot/*.go` uses `wrapError(funcName, retType, retMsg)`
- `pkg/trd/*.go` uses `wrapError(funcName, retType, retMsg)`
- `internal/client/client.go` uses `fmt.Errorf("...")` for some errors
- Some functions return bare `err` without wrapping

### 8.2 Fix

Standardize on `wrapError()` for all proto API errors and `fmt.Errorf("FuncName: ...")` for validation/internal errors. Ensure all error messages include the function name for debugging.

### 8.3 Files Changed

- `internal/client/client.go` — standardize error messages
- Various `pkg/qot/*.go` and `pkg/trd/*.go` files

---

## Step 9: Test Gaps (P3)

### 9.1 Missing Test Coverage

| Area | Test Needed |
|------|-------------|
| Goroutine leak | Verify all goroutines exit on Close() |
| Reconnect cycle | Test full disconnect → reconnect flow |
| Trading idempotency | Verify PlaceOrder/ModifyOrder/CancelOrder never retry |
| CancelAllOrder ForAll | Test the ForAll=true path |
| Pool contention | Test concurrent Get/Put with slow dial |
| Middleware chain | Test interceptor ordering and short-circuit |
| AES CBC round-trip | Test encrypt → decrypt with PKCS#7 padding |
| Breaker halfOpenMax | Test that halfOpenMax is enforced |
| New APIs | Test GetFinancialsEarningsPriceMove/History |

### 9.2 Files Changed

- `internal/client/client_test.go`
- `internal/client/pool_test.go`
- `internal/client/aes_test.go`
- `pkg/breaker/breaker_test.go`
- `pkg/qot/financials_test.go` — new

---

## Step 10: Documentation & Release

### 10.1 Update CHANGELOG.md

Add all Phase 5 items under `[Unreleased]`:
- Fixed ProtoID mismatch (GetTradeDate was using wrong ProtoID 3225)
- Added GetFinancialsEarningsPriceMove (ProtoID 3225)
- Added GetFinancialsEarningsPriceHistory (ProtoID 3226)
- Fixed proto safety violations (replaced GetXxx() with util.ProtoXxx())
- Fixed concurrency bugs (logger race, SetTracer race, setter thread safety)
- Fixed nil dereference in Conn.LocalAddr/RemoteAddr
- Fixed AES CBC PKCS#7 unpadding
- Fixed breaker halfOpenMax enforcement
- Fixed pool contention (mutex no longer held during TCP dial)
- Added middleware/interceptor pattern
- Wired degradation manager, metrics callbacks, otel tracing

### 10.2 Update IMPLEMENTATION_PLAN.md

Mark Phase 5 items as DONE.

### 10.3 Update README.md

- Add GetFinancialsEarningsPriceMove/History to API reference table
- Document interceptor pattern
- Document metrics InitWithServer() usage
- Document otel tracing usage

### 10.4 Commit & Push

```bash
git add -A
git commit -m "phase5: bug fixes, missing APIs, architecture hardening"
git push origin main
git push gitee main
```

---

## Risk Assessment

| Step | Risk Level | Reason |
|------|-----------|--------|
| Step 1 (ProtoID fix) | **CRITICAL** | Changes wire protocol; must be correct |
| Step 2 (New APIs) | **MEDIUM** | New proto files; need regen |
| Step 3 (Proto safety) | **LOW** | Mechanical replacement; well-tested pattern |
| Step 4 (Concurrency) | **MEDIUM** | Lock ordering changes; potential deadlocks |
| Step 5 (AES/Breaker/Pool) | **HIGH** | AES padding affects all encrypted traffic; pool contention fix changes concurrency model |
| Step 6 (Dead code) | **LOW** | Additive only; no behavior change |
| Step 7 (Interceptors) | **MEDIUM** | Refactoring core request path; needs thorough testing |
| Step 8 (Error wrapping) | **LOW** | Cosmetic; no behavior change |
| Step 9 (Tests) | **LOW** | Additive only |

---

## Execution Order

1. Step 1 (ProtoID fix) — must be first, blocks Step 2
2. Step 2 (New APIs) — depends on Step 1
3. Step 3 (Proto safety) — independent, can parallel with Step 2
4. Step 4 (Concurrency) — independent
5. Step 5 (AES/Breaker/Pool) — independent
6. Step 6 (Dead code wiring) — depends on Step 7 for interceptor pattern
7. Step 7 (Interceptors) — depends on Steps 4-5 for thread-safe primitives
8. Step 8 (Error wrapping) — independent, low priority
9. Step 9 (Tests) — after all code changes
10. Step 10 (Docs & Release) — last

---

## Verification Checklist

- [ ] `go build ./...` passes
- [ ] `go vet ./...` passes
- [ ] `go test -race ./...` passes
- [ ] All ProtoIDs match official Futu API v10.6 documentation
- [ ] No `GetXxx()` calls on proto messages (except enum-typed getters)
- [ ] No data races detected by `-race` flag
- [ ] AES encrypt/decrypt round-trip test passes
- [ ] Breaker halfOpenMax enforced in test
- [ ] Pool Get/Put concurrent test passes without timeout
- [ ] CHANGELOG.md updated
- [ ] Committed and pushed to both remotes
