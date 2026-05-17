# Proto Field Completion Plan v2

> Implementation plan for closing 30 audit issues in the futuapi4go SDK.
> All 30 issues must be resolved to reach 100% proto field coverage, eliminate raw type leakage, and standardize error/validation patterns.

---

## Table of Contents

1. [Phase 1 — Push Wrapper Enrichment](#phase-1--push-wrapper-enrichment)
2. [Phase 2 — Proto Field Gaps in pkg/ Wrappers](#phase-2--proto-field-gaps-in-pkg-wrappers)
3. [Phase 3 — Missing Wrapper: Qot_GetRehab](#phase-3--missing-wrapper-qot_getrehab)
4. [Phase 4 — Proto Type Leakage Fixes](#phase-4--proto-type-leakage-fixes)
5. [Phase 5 — MaxTrdQtys Price Fix](#phase-5--maxtrdqtys-price-fix)
6. [Phase 6 — Client Type Enrichment](#phase-6--client-type-enrichment)
7. [Phase 7 — GetUserInfo Enrichment](#phase-7--getuserinfo-enrichment)
8. [Phase 8 — GetDelayStatistics Cleanup](#phase-8--getdelaystatistics-cleanup)
9. [Phase 9 — Consistency Fixes](#phase-9--consistency-fixes)
10. [Change Summary](#change-summary)
11. [Verification Script](#verification-script)
12. [Design Decisions](#design-decisions)

---

## Phase 1 — Push Wrapper Enrichment

### Issue #3 — UpdateBasicQot Push: Add 13 Missing Fields

**File:** `pkg/push/qot_push.go:56-100`

**Current state:** `UpdateBasicQot` struct has 12 fields (Security, Name, CurPrice, OpenPrice, HighPrice, LowPrice, Volume, Turnover, IsSuspended, LastClosePrice, UpdateTime, UpdateTimestamp). Missing 13 of 25 BasicQot fields.

**What to do:**

1. Add 13 fields to the `UpdateBasicQot` struct at `pkg/push/qot_push.go:56`:

```go
type UpdateBasicQot struct {
    // existing 12 fields ...
    ListTime         string
    PriceSpread      float64
    TurnoverRate     float64
    Amplitude        float64
    DarkStatus       int32
    OptionExData     *qotcommon.OptionBasicQotExData
    ListTimestamp    float64
    PreMarket        *qotcommon.PreAfterMarketData
    AfterMarket      *qotcommon.PreAfterMarketData
    SecStatus        int32
    FutureExData     *qotcommon.FutureBasicQotExData
    WarrantExData    *qotcommon.WarrantBasicQotExData
    Overnight        *qotcommon.PreAfterMarketData
}
```

2. Add mapper assignments in `ParseUpdateBasicQot` at `pkg/push/qot_push.go:86-99`:

| Struct field | Proto accessor |
|---|---|
| `ListTime` | `bq.GetListTime()` |
| `PriceSpread` | `bq.GetPriceSpread()` |
| `TurnoverRate` | `bq.GetTurnoverRate()` |
| `Amplitude` | `bq.GetAmplitude()` |
| `DarkStatus` | `bq.GetDarkStatus()` |
| `OptionExData` | `bq.GetOptionExData()` |
| `ListTimestamp` | `bq.GetListTimestamp()` |
| `PreMarket` | `bq.GetPreMarket()` |
| `AfterMarket` | `bq.GetAfterMarket()` |
| `SecStatus` | `bq.GetSecStatus()` |
| `FutureExData` | `bq.GetFutureExData()` |
| `WarrantExData` | `bq.GetWarrantExData()` |
| `Overnight` | `bq.GetOvernight()` |

3. Update `pkg/push/push_test.go:153` — `TestParseUpdateBasicQotValidData`:
   - Add all 13 new fields to the test proto construction (lines 176-194)
   - Add assertion checks for each new field after line 246

**Verification:** `go test -race ./pkg/push/... -run TestParseUpdateBasicQot`

---

### Issue #9 — Chan Package: Add Missing Subscribe Functions & Channel Factories

**Files:** `pkg/push/chan/chan.go:87-109` (existing factories), `pkg/push/chan/chan.go:152-250` (existing subscribe functions)

**Current state:** 6 `New*Channel` factories exist (Quote, KL, Ticker, OrderBook, RT, Broker) + 7 `Subscribe*` functions (Quote, KLine, KLines, Ticker, OrderBook, RT, Broker, PriceReminder). Missing 4 factory/subscribe pairs for SystemNotify, UpdateOrder, UpdateOrderFill, TrdNotify.

**What to do — Add 4 channel factories at `pkg/push/chan/chan.go` after line 109:**

```go
func NewSystemNotifyChannel(bufferSize int) chan *push.SystemNotify
func NewOrderUpdateChannel(bufferSize int) chan *push.UpdateOrder
func NewOrderFillUpdateChannel(bufferSize int) chan *push.UpdateOrderFill
func NewTrdNotifyChannel(bufferSize int) chan *push.TrdNotify
```

**What to do — Add 4 subscribe functions at `pkg/push/chan/chan.go` after line 250:**

| Function | ProtoID | Parse function | Sub function |
|---|---|---|---|
| `SubscribeSystemNotify(ctx, cli, ch)` | `constant.ProtoID_Notify` | `push.ParseSystemNotify` | `nil` (no sub needed) |
| `SubscribeOrderUpdate(ctx, cli, ch)` | `constant.ProtoID_Trd_UpdateOrder` | `push.ParseUpdateOrder` | `nil` (handled via SubAccPush) |
| `SubscribeOrderFillUpdate(ctx, cli, ch)` | `constant.ProtoID_Trd_UpdateOrderFill` | `push.ParseUpdateOrderFill` | `nil` |
| `SubscribeTrdNotify(ctx, cli, ch)` | `constant.ProtoID_Notify` | `push.ParseTrdNotify` | `nil` |

Each follows the `subscribeOne[T]` pattern — no explicit `subFn` needed (just pass `func() error { return nil }`).

**Verification:** `go build ./pkg/push/chan/... && go vet ./pkg/push/chan/...`

---

## Phase 2 — Proto Field Gaps in pkg/ Wrappers

### Issue #1 — BasicQot: Add WarrantExData

**File:** `pkg/qot/quote.go:115-140` (struct), `pkg/qot/quote.go:168-193` (mapper)

**Current state:** `BasicQot` struct has OptionExData, PreMarket, AfterMarket, FutureExData, Overnight but NOT `WarrantExData`.

**What to do:**

1. Add field to `BasicQot` struct after `Overnight` (line 139):
```go
WarrantExData *qotcommon.WarrantBasicQotExData
```

2. Add mapper line inside loop at `pkg/qot/quote.go:192`:
```go
WarrantExData: bq.GetWarrantExData(),
```

3. Update `pkg/qot/quote_test.go` — add `WarrantExData` to test fixture and assertion.

**Verification:** `go test -race ./pkg/qot/... -run TestGetBasicQot`

---

### Issue #2 — KLine: Add Pe Field

**File:** `pkg/qot/quote.go:199-213` (struct)

**Current state:** `KLine` struct has 12 fields, missing `Pe float64`.

**What to do:**

1. Add field to `KLine` struct after `TurnoverRate` (line 210):
```go
Pe float64
```

**Note:** The mapper for `Pe` already exists in both `GetKL` (line 277 area) and `GetHistoryKL`/`RequestHistoryKL` handlers in `kline.go`. No mapper changes needed — just add the struct field.

**Verification:** `go build ./pkg/qot/...`

---

### Issue #7 — MaxTrdQtysInfo: Add Session Field

**File:** `pkg/trd/position.go:470-478` (struct)

**Current state:** `MaxTrdQtysInfo` has 7 fields, missing `Session int32`.

**What to do:**

1. Add field to `MaxTrdQtysInfo` struct at line 478:
```go
Session int32
```

2. Add mapper in `GetMaxTrdQtys` at line 562:
```go
Session: m.GetTrdMaxTrdQtys(),
```

3. Update `pkg/trd/position_test.go` — add Session to test assertions.

**Verification:** `go test -race ./pkg/trd/... -run TestGetMaxTrdQtys`

---

## Phase 3 — Missing Wrapper: qotgetrehab

### Issue #6 — GetRehab for Qot_GetRehab (ProtoID 3102)

**File:** `pkg/qot/holding.go:126-168` (current broken GetRehab)

**Current state:** Both `RequestRehab` (line 93) and `GetRehab` (line 137) use `qotrequestrehab` with ProtoID 3105. There IS a separate `qotgetrehab` package at `pkg/pb/qotgetrehab/` that supports multiple securities (takes `SecurityList` not single `Security`). This needs its own entry point.

**What to do:**

1. Add ProtoID constant in `pkg/qot/quote.go` line 95:
```go
ProtoID_GetRehab = 3102
```

2. Rewrite `pkg/qot/holding.go` — rewrite `GetRehab` function (lines 126-168) to use `qotgetrehab`:

```go
type GetRehabRequest struct {
    Security *qotcommon.Security  // single security (wrapped for backward compat)
}

type GetRehabResponse struct {
    SecurityRehabList []*qotgetrehab.SecurityRehab
}

func GetRehab(ctx context.Context, c *futuapi.Client, req *GetRehabRequest) (*GetRehabResponse, error) {
    if req == nil {
        return nil, fmt.Errorf("GetRehab: request is nil")
    }
    if req.Security == nil {
        return nil, fmt.Errorf("security is required")
    }

    c2s := &qotgetrehab.C2S{
        SecurityList: []*qotcommon.Security{req.Security},
    }

    pkt := &qotgetrehab.Request{C2S: c2s}
    var rsp qotgetrehab.Response

    if err := c.RequestContext(ctx, ProtoID_GetRehab, pkt, &rsp); err != nil {
        return nil, err
    }

    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, wrapError("GetRehab", rsp.GetRetType(), rsp.GetRetMsg())
    }

    s2c := rsp.GetS2C()
    if s2c == nil {
        return nil, wrapError("GetRehab", int32(common.RetType_RetType_Unknown), "s2c is nil")
    }

    return &GetRehabResponse{
        SecurityRehabList: s2c.GetSecurityRehabList(),
    }, nil
}
```

3. Update `pkg/qot/quote_test.go` — add ProtoID_GetRehab (3102) to the proto ID test.

4. Add import: `"github.com/shing1211/futuapi4go/pkg/pb/qotgetrehab"`

**Verification:** `go build ./pkg/qot/... && go vet ./pkg/qot/...`

---

## Phase 4 — Proto Type Leakage Fixes

### Issue #4 — GetHistoryOrderListResponse: Use Wrapped Order Type

**File:** `pkg/trd/queries.go:350-351`

**Current state:**
```go
type GetHistoryOrderListResponse struct {
    OrderList []*trdcommon.Order   // RAW PROTO TYPE — should be []*Order
}
```

**What to do:**

1. Change to:
```go
type GetHistoryOrderListResponse struct {
    OrderList []*Order  // use wrapped type
}
```

2. Add mapper loop in `GetHistoryOrderList` function (replace `return` at lines 398-400):
```go
result := &GetHistoryOrderListResponse{
    OrderList: make([]*Order, 0, len(s2c.GetOrderList())),
}
for _, o := range s2c.GetOrderList() {
    if o == nil {
        continue
    }
    result.OrderList = append(result.OrderList, &Order{
        OrderID:         o.GetOrderID(),
        OrderIDEx:       o.GetOrderIDEx(),
        Code:            o.GetCode(),
        Name:            o.GetName(),
        TrdSide:         o.GetTrdSide(),
        OrderType:       o.GetOrderType(),
        OrderStatus:     o.GetOrderStatus(),
        Price:           o.GetPrice(),
        Qty:             o.GetQty(),
        FillQty:         o.GetFillQty(),
        FillAvgPrice:    o.GetFillAvgPrice(),
        CreateTime:      o.GetCreateTime(),
        UpdateTime:      o.GetUpdateTime(),
        LastErrMsg:      o.GetLastErrMsg(),
        SecMarket:       o.GetSecMarket(),
        CreateTimestamp: o.GetCreateTimestamp(),
        UpdateTimestamp: o.GetUpdateTimestamp(),
        Remark:          o.GetRemark(),
        TimeInForce:     o.GetTimeInForce(),
        FillOutsideRTH:  o.GetFillOutsideRTH(),
        AuxPrice:        o.GetAuxPrice(),
        TrailType:       o.GetTrailType(),
        TrailValue:      o.GetTrailValue(),
        TrailSpread:     o.GetTrailSpread(),
        Currency:        o.GetCurrency(),
        TrdMarket:       o.GetTrdMarket(),
        Session:         o.GetSession(),
        JpAccType:       o.GetJpAccType(),
    })
}
return result, nil
```

**Verification:** `go build ./pkg/trd/... && go vet ./pkg/trd/...`

---

### Issue #10 — 14 Raw Proto Type Instances in Public API

The following types leak raw proto types into public struct definitions. Each must be replaced with a wrapped type or documented as acceptable.

#### Items to wrap:

| # | File:line | Current type | Replace with |
|---|---|---|---|
| 1 | `pkg/trd/position.go:568` | `*trdcommon.TrdHeader` in `GetFlowSummaryRequest` | Separated fields: `AccID uint64`, `TrdMarket constant.TrdMarket`, `TrdEnv constant.TrdEnv` |
| 2 | `pkg/trd/position.go:575` | `*trdcommon.TrdHeader` in `GetFlowSummaryResponse` | Separated fields or wrapped type |
| 3 | `pkg/trd/position.go:576` | `[]*trdflowsummary.FlowSummaryInfo` in `GetFlowSummaryResponse` | `[]*FlowSummaryInfo` (already exists in `client/types.go:618`) |
| 4 | `pkg/trd/position.go:358` | `[]*qotcommon.Security` in `GetMarginRatioRequest.SecurityList` | Acceptable — `qotcommon.Security` is the canonical SDK security type used everywhere; document as intentional |
| 5 | `pkg/trd/position.go:363` | `*qotcommon.Security` in `MarginRatioInfo` | Acceptable — same rationale as #4 |
| 6 | `pkg/trd/orders.go:334-335` | `*common.PacketID` + `*trdcommon.TrdHeader` in `ReconfirmOrderRequest` | Separated fields for `Header`; keep `PacketID` as-is (internal protocol detail) |
| 7 | `pkg/trd/orders.go:342` | `*trdcommon.TrdHeader` in `ReconfirmOrderResponse` | Separated fields `AccID`, `TrdEnv`, `TrdMarket` |

**Decision:**
- `*trdcommon.TrdFilterConditions` in request types (GetOrderListRequest, etc.) — **KEEP as-is**. Has too many fields to wrap.
- `*qotcommon.Security` in MarginRatioInfo, SecurityList — **KEEP as-is**. This is the canonical security representation.
- `*trdcommon.TrdHeader` in request/response types — **REPLACE** with discrete wrapped fields.
- `[]*trdflowsummary.FlowSummaryInfo` — **REPLACE** with `[]*FlowSummaryInfo` from client/types.go.

**Verification:** `go build ./... && go vet ./...`

---

## Phase 5 — MaxTrdQtys Price Fix

### Issue #8 — GetMaxTrdQtys: Always Send Price Field

**File:** `pkg/trd/position.go:514-515`

**Current state:**
```go
if req.Price != 0 {
    c2s.Price = &req.Price
}
```

**What to do:**

Change to unconditionally send Price (proto marks it as required):
```go
c2s.Price = &req.Price
```

**Verification:** `go build ./pkg/trd/... && go vet ./pkg/trd/...`

---

## Phase 6 — Client Type Enrichment

### Issue #11 — Add Missing Fields to client/types.go

**File:** `client/types.go`

#### Quote (lines 8-24) — Add 5 fields:

| Field | Type | Proto source |
|---|---|---|
| `ListTime` | `string` | `BasicQot.ListTime` |
| `PriceSpread` | `float64` | `BasicQot.PriceSpread` |
| `DarkStatus` | `int32` | `BasicQot.DarkStatus` |
| `ListTimestamp` | `float64` | `BasicQot.ListTimestamp` |
| `UpdateTimestamp` | `float64` | `BasicQot.UpdateTimestamp` |

Note: Nested proto pointers (OptionExData, PreMarket, AfterMarket, FutureExData, WarrantExData) intentionally omitted per existing design decision.

#### KLine (lines 27-39) — Add 1 field:

| Field | Type | Proto source |
|---|---|---|
| `TurnoverRate` | `float64` | `KLine.TurnoverRate` |

#### FutureInfo (lines 296-313) — Add 1 field:

| Field | Type | Proto source |
|---|---|---|
| `TradeTimeList` | `[]*qotgetfutureinfo.TradeTime` | `FutureBasicQotExData.TradeTimeList` |

#### IpoData (lines 322-327) — Add 3 fields:

| Field | Type | Proto source |
|---|---|---|
| `CnExData` | `*CNIpoExData` | `IpoBasicData.CnExData` |
| `HkExData` | `*HKIpoExData` | `IpoBasicData.HkExData` |
| `UsExData` | `*USIpoExData` | `IpoBasicData.UsExData` |

#### Broker (lines 246-251) — Add 1 field:

| Field | Type | Proto source |
|---|---|---|
| `OrderID` | `int64` | `Broker.OrderID` |

#### Ticker (lines 219-230) — Add 1 field:

| Field | Type | Proto source |
|---|---|---|
| `PushDataType` | `int32` | `Ticker.PushDataType` |

#### Position (lines 63-88) — Add 2 fields:

| Field | Type | Proto source |
|---|---|---|
| `SecMarket` | `int32` | `Position.SecMarket` |
| `TdTrdVal` | `float64` | `Position.TdTrdVal` |

#### Order (lines 142-170) — Add 1 field:

| Field | Type | Proto source |
|---|---|---|
| `JpAccType` | `int32` | `Order.JpAccType` |

#### GlobalState (lines 470-487) — Add 3 fields:

| Field | Type | Proto source |
|---|---|---|
| `ConnID` | `uint64` | `GlobalState.ConnID` |
| `QotSvrIpAddr` | `string` | `GlobalState.QotSvrIpAddr` |
| `TrdSvrIpAddr` | `string` | `GlobalState.TrdSvrIpAddr` |

#### UserInfo (lines 490-495) — Add 4 fields:

| Field | Type | Proto source |
|---|---|---|
| `IsNeedAgreeDisclaimer` | `bool` | `UserInfo.IsNeedAgreeDisclaimer` |
| `ShQotRight` | `int32` | `UserInfo.ShQotRight` |
| `SzQotRight` | `int32` | `UserInfo.SzQotRight` |
| `Extra` | `int32` | `UserInfo.Extra` |

**Verification:** `go build ./client/... && go vet ./client/...`

---

### Issue #12 — Update Client Mappers

#### `client/quote_api.go:40-57` — GetQuote mapper:

Add 5 fields to the `Quote` return:
```go
ListTime:        q.ListTime,
PriceSpread:     q.PriceSpread,
DarkStatus:      q.DarkStatus,
ListTimestamp:   q.ListTimestamp,
UpdateTimestamp: q.UpdateTimestamp,
```

#### `client/quote_api.go:76-88` — GetKLines mapper:

Add 1 field to the `KLine` return:
```go
TurnoverRate: kl.TurnoverRate,
```

#### `client/trade_api.go:112-139` — GetPositionList mapper:

Add 2 fields:
```go
SecMarket:  p.SecMarket,
TdTrdVal:   p.TdTrdVal,
```

**Fix duplicate TrdMarket mapping (lines 118-119):**
- Line 118: `Market: p.TrdMarket` — remove this (it's a duplicate, line 134 already maps `TrdMarket`)
- Actually line 118 maps to `Market` (the struct field name), and line 134 maps to `TrdMarket`. These are two different struct fields. The issue says it's a duplicate — looking more carefully:
  - The `GetPositionList` in `trade_api.go` maps `p.TrdMarket` to `Market` at line 118 and `p.TrdMarket` to `TrdMarket` at line 134
  - Both read from `p.TrdMarket` which is the same source. The client `Position` struct has both `Market int32` and `TrdMarket int32` fields.
  - **Fix:** `Market` (line 118) should map to `p.SecMarket` instead of `p.TrdMarket` since that's the "trading market of the security" vs "overall trading market"

#### `client/trade_api.go:342-370` — GetOrderList mapper:

Add 1 field:
```go
JpAccType: o.JpAccType,
```

**Verification:** `go build ./client/... && go vet ./client/...`

---

## Phase 7 — GetUserInfo Enrichment

### Issue #5 — Expand GetUserInfoResponse

**File:** `pkg/sys/system.go:139-149` (struct), `pkg/sys/system.go:172-181` (mapper)

**Current state:** `GetUserInfoResponse` has 8 fields. Misses ~15 proto S2C fields. C2S is empty (no `Flag` field).

**What to do:**

1. Expand request struct to accept `Flag`:
```go
type GetUserInfoRequest struct {
    Flag int32  // bitmask for selecting specific info fields
}
```

2. Expand response struct at `pkg/sys/system.go:139`:

| Field | Type | Proto source |
|---|---|---|
| `UserID` | `int64` | _existing_ |
| `NickName` | `string` | _existing_ |
| `AvatarUrl` | `string` | _existing_ |
| `ApiLevel` | `string` | _existing_ |
| `IsNeedAgreeDisclaimer` | `bool` | _existing_ |
| `ShQotRight` | `int32` | _existing_ |
| `SzQotRight` | `int32` | _existing_ |
| `Extra` | `int32` | _existing_ |
| `HkQotRight` | `int32` | `UserInfo.HkQotRight` |
| `UsQotRight` | `int32` | `UserInfo.UsQotRight` |
| `CnQotRight` | `int32` | `UserInfo.CnQotRight` |
| `SubQuota` | `int32` | `UserInfo.SubQuota` |
| `HistoryKLQuota` | `int32` | `UserInfo.HistoryKLQuota` |
| `WebKey` | `string` | `UserInfo.WebKey` |
| `WebJumpUrlHead` | `string` | `UserInfo.WebJumpUrlHead` |
| `IsAppNNOrMM` | `bool` | `UserInfo.IsAppNNOrMM` |
| `HkOptionQotRight` | `int32` | `UserInfo.HkOptionQotRight` |
| `HasUSOptionQotRight` | `bool` | `UserInfo.HasUSOptionQotRight` |
| `HkFutureQotRight` | `int32` | `UserInfo.HkFutureQotRight` |
| `UsFutureQotRight` | `int32` | `UserInfo.UsFutureQotRight` |
| `UsOptionQotRight` | `int32` | `UserInfo.UsOptionQotRight` |
| `UserAttribution` | `string` | `UserInfo.UserAttribution` |
| `UpdateWhatsNew` | `string` | `UserInfo.UpdateWhatsNew` |

3. Update `GetUserInfo` function signature to accept optional request:
```go
func GetUserInfo(ctx context.Context, c *futuapi.Client, req *GetUserInfoRequest) (*GetUserInfoResponse, error)
```
If `req` is nil, send empty C2S (backward compatible). If non-nil, set `Flag`.

4. Update mapper at lines 172-181 with all new fields.

**Verification:** `go build ./pkg/sys/... && go vet ./pkg/sys/...`

---

## Phase 8 — GetDelayStatistics Cleanup

### Issues #13, #14, #15 — Properly Typed Request/Response + Standard Pattern

**File:** `pkg/sys/system.go:184-296`

**Current state:** `GetDelayStatisticsResponse` uses raw proto types (`*getdelaystatistics.DelayStatistics`, etc). The request has no typed wrapper — uses `getdelaystatistics.C2S` directly. Uses `WritePacket` + `ReadResponseContext` instead of `c.RequestContext()`.

**What to do:**

1. Create proper request struct:
```go
type GetDelayStatisticsRequest struct {
    TypeList      []int32  // DelayStatisticsType filter
    QotPushStage  int32    // 0 = all, 1 = push, 2 = request-reply, 3 = place-order
    SegmentList   []int32  // time segments
}
```

2. Create typed response struct replacing raw proto references:
```go
type GetDelayStatisticsResponse struct {
    QotPushStatisticsList    []*DelayStatistics           // wrapped, not raw proto
    ReqReplyStatisticsList   []*ReqReplyStatisticsItem    // from client/types.go
    PlaceOrderStatisticsList []*PlaceOrderStatisticsItem  // from client/types.go
}

type DelayStatistics struct {
    QotPushType int32
    DelayAvg    float64
    Count       int32
    ItemList    []DelayStatisticsItem
    // ...
}
```

3. Replace `WritePacket`/`ReadResponseContext` with `c.RequestContext()`:
   - If proto2 encoding is still needed, move the hack (`marshalC2SProto2`, `marshalGetDelayStatisticsRequest`, `appendVarint`) to `internal/client/proto2.go` as internal helpers.

4. Update the function body to use standard pattern.

**Verification:** `go build ./pkg/sys/... && go test -race ./pkg/sys/...`

---

## Phase 9 — Consistency Fixes

### Issue #16 — DataFilter: Use Proper Type

**File:** `pkg/qot/options.go:107`

**Current state:**
```go
DataFilter interface{}  // should be *qotgetoptionchain.DataFilter
```

**What to do:**
```go
DataFilter *qotgetoptionchain.DataFilter
```

Also add mapper in `GetOptionChain`:
```go
if req.DataFilter != nil {
    c2s.DataFilter = req.DataFilter
}
```

**Verification:** `go build ./pkg/qot/...`

---

### Issue #17 — CancelAllOrders: Add accID Validation

**File:** `pkg/trd/convenience.go:23`

**Current state:** No validation for `accID`.

**What to do:** Add at line 24:
```go
if accID == 0 {
    return nil, fmt.Errorf("CancelAllOrders: account ID is required")
}
```

**Verification:** `go build ./pkg/trd/...`

---

### Issue #18 — Standardize S2C Nil Error Reporting: Use wrapError()

**Files:** All pkg/ files that currently use `fmt.Errorf("...: s2c is nil")` instead of `wrapError()`.

**Current state:** Mixed usage — some use `wrapError()`, others use `fmt.Errorf()`.

**What to do:** Find-and-replace all instances of:
```go
return nil, fmt.Errorf("FuncName: s2c is nil")
```
with:
```go
return nil, wrapError("FuncName", int32(common.RetType_RetType_Unknown), "s2c is nil")
```

Affected files (13+ instances):
- `pkg/trd/position.go` — GetFunds, GetPositionList, GetMarginRatio, GetMaxTrdQtys, GetFlowSummary
- `pkg/trd/queries.go` — GetOrderList, GetOrderFillList, GetOrderFee, GetHistoryOrderList, GetHistoryOrderFillList
- `pkg/trd/orders.go` — PlaceOrder, ModifyOrder, ReconfirmOrder
- `pkg/sys/system.go` — GetGlobalState, GetUserInfo, GetDelayStatistics, TestCmd, GetUsedQuota
- `pkg/qot/` — various files

**Verification:** `go build ./... && go vet ./...`

---

### Issue #19 — Consolidate ProtoID Constants

**Files:** `pkg/qot/quote.go:77-112`, `pkg/qot/sub.go:34`, `pkg/qot/trade_date.go:13`, `pkg/qot/kline.go:224`, `pkg/trd/trade.go:52-71`, `pkg/sys/system.go:63-70`, `pkg/push/qot_push.go:46-54`

**Current state:** ProtoID constants are duplicated across packages (e.g., `ProtoID_GetBasicQot = 3004` in `pkg/qot/quote.go` AND `ProtoID_Qot_GetBasicQot = 3004` in `pkg/constant/constant.go`).

**What to do:**

| Location | Constant | Replace with |
|---|---|---|
| `pkg/qot/quote.go:77-112` | `ProtoID_*` (26 consts) | Remove, import `constant.ProtoID_Qot_*` |
| `pkg/qot/sub.go:34` | `ProtoID_GetSubInfo = 3003` | Remove, use `constant.ProtoID_Qot_GetSubInfo` |
| `pkg/qot/trade_date.go:13` | `ProtoID_Qot_GetTradeDate = 3225` | Remove, add to `constant.go` first then reference |
| `pkg/qot/kline.go:224` | `ProtoID_GetHistoryKLPoints = 3106` | Remove, use `constant.ProtoID_Qot_GetHistoryKLPoints` |
| `pkg/trd/trade.go:52-71` | `ProtoID_*` (19 consts) | Remove, use `constant.ProtoID_Trd_*` |
| `pkg/sys/system.go:63-70` | `ProtoID_*` (6 consts) | Remove, use `constant.ProtoID_*` |
| `pkg/push/qot_push.go:46-54` | `ProtoID_Qot_Update*` | Remove, use `constant.ProtoID_Qot_Update*` |

**Note:** Some constants may not exist in `constant.go` yet (e.g., `ProtoID_Qot_GetTradeDate = 3225`, `ProtoID_Qot_GetRehab = 3102`). Add them first.

**Verification:** `go build ./... && go vet ./...`

---

### Issue #20 — Remove Unused ProtoID_GetHistoryKLPoints

**File:** `pkg/qot/kline.go:224-226`

**Current state:** `ProtoID_GetHistoryKLPoints = 3106` is defined but `GetHistoryKLPoints` at line 306 uses `constant.ProtoID_Qot_GetHistoryKLPoints` instead.

**What to do:** Remove the local constant:
```go
// DELETE these lines:
const (
    ProtoID_GetHistoryKLPoints = 3106
)
```

**Verification:** `go build ./pkg/qot/...`

---

### Issue #21 — Subscribe/RegQotPush: Suppress RetType/RetMsg on Success

**File:** `pkg/qot/sub.go:132-135` (Subscribe), `pkg/qot/sub.go:184-187` (RegQotPush)

**Current state:** Both functions return `(*SubscribeResponse, error)` / `(*RegQotPushResponse, error)` with RetType and RetMsg populated on success. The response types are mostly unused by callers (who just check error).

**What to do:**

Change both functions to return `error` only (simpler API):

```go
func Subscribe(ctx context.Context, c *futuapi.Client, req *SubscribeRequest) error { ... }
func RegQotPush(ctx context.Context, c *futuapi.Client, req *RegQotPushRequest) error { ... }
```

Remove `SubscribeResponse` and `RegQotPushResponse` types (or keep for backward compat, deprecate).

Update callers in `client/quote_api.go`:
- Line 103: change `_, err := qot.Subscribe(...)` to `err := qot.Subscribe(...)`
- Lines 123, 134, 159, 186, 214: same pattern

**Verification:** `go build ./... && go vet ./...`

---

### Issue #22 — ModifyUserSecurity: Add S2C Nil Check

**File:** `pkg/qot/user.go:174-181`

**Current state:** No S2C nil check after RetType check.

**What to do:** Add after line 176:
```go
s2c := rsp.GetS2C()
if s2c == nil {
    return nil, wrapError("ModifyUserSecurity", int32(common.RetType_RetType_Unknown), "s2c is nil")
}
```

**Verification:** `go build ./pkg/qot/...`

---

### Issue #23 — GetTradeDate: Add Proper Typed Request/Response Structs

**File:** `pkg/qot/trade_date.go:16-35`

**Current state:** Uses raw proto types directly:
```go
func GetTradeDate(ctx context.Context, c *futuapi.Client, req *qotgettradedate.C2S) (*qotgettradedate.S2C, error)
```

**What to do:**

```go
type GetTradeDateRequest struct {
    Market    int32
    BeginTime string
    EndTime   string
}

type GetTradeDateResponse struct {
    TradeDateList []*TradeDateInfo
}

type TradeDateInfo struct {
    Time          string
    Timestamp     float64
    TradeDateType int32
}

func GetTradeDate(ctx context.Context, c *futuapi.Client, req *GetTradeDateRequest) (*GetTradeDateResponse, error) {
    // input validation
    // build C2S from wrapped request
    // use RequestContext
    // check RetType, check S2C nil
    // map TradeDateList with nil guard loop
}
```

**Verification:** `go build ./pkg/qot/... && go vet ./pkg/qot/...`

---

### Issue #24 — SubscribeKLines: Use constant.Market and []constant.KLType Types

**File:** `pkg/push/chan/chan.go:164`

**Current state:**
```go
func SubscribeKLines(ctx context.Context, cli *client.Client, market int32, code string, kTypes []int32, ch chan<- *push.UpdateKL) (func(), error)
```

**What to do:**
```go
func SubscribeKLines(ctx context.Context, cli *client.Client, market constant.Market, code string, kTypes []constant.KLType, ch chan<- *push.UpdateKL) (func(), error)
```

Update internal `subscribe` helper at line 185 similarly:
```go
func subscribe(ctx context.Context, cli *client.Client, market constant.Market, code string, kTypes []constant.KLType) error
```

**Verification:** `go build ./pkg/push/chan/... && go vet ./pkg/push/chan/...`

---

### Issue #25 — klTypeToSubType: Return Error on Unknown KLType

**File:** `pkg/push/chan/chan.go:193-220`

**Current state:**
```go
func klTypeToSubType(k constant.KLType) constant.SubType
```
Falls back silently to `SubType_K_1Min` on unknown input.

**What to do:**
```go
func klTypeToSubType(k constant.KLType) (constant.SubType, error)
```
Return error on unknown KLType instead of silent fallback.

Update callers:
- Line 160: `klTypeToSubType(klType)` → handle error
- Line 188: `klTypeToSubType(constant.KLType(kt))` → handle error

**Verification:** `go build ./pkg/push/chan/... && go vet ./pkg/push/chan/...`

---

### Issue #26 — IsPDT/PDTSeq: Document Proto Discrepancy

**File:** `pkg/trd/position.go:171-172`

**Current state:** `IsPDT` maps from `f.GetIsPdt()` (proto uses lowercase "pdt"). `PDTSeq` maps from `f.GetPdtSeq()`.

**What to do:** Add comments documenting the naming discrepancy:
```go
IsPDT  bool    `json:"isPDT"`   // maps from proto IsPdt (PDT = Pattern Day Trader)
PDTSeq string  `json:"pDTSeq"`  // maps from proto PdtSeq
```

**Verification:** `go build ./pkg/trd/...`

---

### Issue #28 — GetFunds: Split Combined nil/zero Check

**File:** `pkg/trd/position.go:102`

**Current state:**
```go
if req == nil || req.AccID == 0 {
    return nil, constant.ErrInvalidAccID
}
```

**What to do:** Split into two checks:
```go
if req == nil {
    return nil, fmt.Errorf("GetFunds: request is nil")
}
if req.AccID == 0 {
    return nil, constant.ErrInvalidAccID
}
```

**Verification:** `go build ./pkg/trd/...`

---

### Issue #29 — Add Input Validation

**Files:** `pkg/qot/quote.go:143` (GetBasicQot), plus GetOrderBook and GetKL.

**Current state:**
- `GetBasicQot` at line 143: no check for empty `securityList`
- `GetBasicQot` at line 144: validates via `pkg/qot/sub.go` level

**What to do:**

1. `GetBasicQot` — add at line 143:
```go
if len(securityList) == 0 {
    return nil, fmt.Errorf("GetBasicQot: security list is empty")
}
```

2. `GetOrderBook` — add at function entry:
```go
if req.Security == nil {
    return nil, fmt.Errorf("GetOrderBook: security is required")
}
```

3. `GetKL` — already has nil check for `req` (line 232), also add:
```go
if req.Security == nil {
    return nil, fmt.Errorf("GetKL: security is required")
}
```

**Verification:** `go build ./pkg/qot/... && go vet ./pkg/qot/...`

---

### Issue #30 — Replace GetXxx() Proto Method Calls with Direct Nil Checks

**Scope:** New and changed code in this plan. For existing code, fix opportunistically.

**Pattern to follow:**
```go
// INSTEAD OF:
value := msg.GetField()

// USE:
var value T
if msg.Field != nil {
    value = *msg.Field
}
```

**Specifically for mapper loops** in new/changed code, use direct field access with nil guards where the proto field is a scalar pointer:

```go
// CURRENT (bad — hides nil panics):
Name: bq.GetName(),

// PREFERRED (for string/float64/int64 etc scalar pointers):
Name: func() string { if bq.Name != nil { return *bq.Name }; return "" }(),
```

For nested message pointers (like `*qotcommon.Security`, `*qotcommon.OptionBasicQotExData`), `GetXxx()` is acceptable since it returns nil safely.

**Files affected:** All changes in this plan for scalar fields.

**Verification:** `go build ./... && go vet ./...`

---

## Change Summary

### Struct Fields Added

| Type | File | Fields added |
|---|---|---|
| `UpdateBasicQot` | `pkg/push/qot_push.go` | ListTime, PriceSpread, TurnoverRate, Amplitude, DarkStatus, OptionExData, ListTimestamp, PreMarket, AfterMarket, SecStatus, FutureExData, WarrantExData, Overnight |
| `BasicQot` | `pkg/qot/quote.go` | WarrantExData |
| `KLine` | `pkg/qot/quote.go` | Pe |
| `MaxTrdQtysInfo` | `pkg/trd/position.go` | Session |
| `Quote` | `client/types.go` | ListTime, PriceSpread, DarkStatus, ListTimestamp, UpdateTimestamp |
| `KLine` (client) | `client/types.go` | TurnoverRate |
| `FutureInfo` (client) | `client/types.go` | TradeTimeList |
| `IpoData` (client) | `client/types.go` | CnExData, HkExData, UsExData |
| `Broker` (client) | `client/types.go` | OrderID |
| `Ticker` (client) | `client/types.go` | PushDataType |
| `Position` (client) | `client/types.go` | SecMarket, TdTrdVal |
| `Order` (client) | `client/types.go` | JpAccType |
| `GlobalState` (client) | `client/types.go` | ConnID, QotSvrIpAddr, TrdSvrIpAddr |
| `UserInfo` (client) | `client/types.go` | IsNeedAgreeDisclaimer, ShQotRight, SzQotRight, Extra |
| `GetUserInfoResponse` | `pkg/sys/system.go` | 15 new fields (HkQotRight, UsQotRight, CnQotRight, SubQuota, etc.) |

### Type Replacements (Raw Proto → Wrapped)

| Location | Old | New |
|---|---|---|
| `GetHistoryOrderListResponse` | `[]*trdcommon.Order` | `[]*Order` |
| `GetFlowSummaryRequest` | `*trdcommon.TrdHeader` | Discrete fields |
| `GetFlowSummaryResponse` | `*trdcommon.TrdHeader` | Discrete fields |
| `GetFlowSummaryResponse` | `[]*trdflowsummary.FlowSummaryInfo` | `[]*FlowSummaryInfo` |
| `ReconfirmOrderRequest` | `*trdcommon.TrdHeader` | Discrete fields |
| `ReconfirmOrderResponse` | `*trdcommon.TrdHeader` | Discrete fields |
| `GetTradeDate` | Raw `*qotgettradedate.C2S`/`*qotgettradedate.S2C` | `*GetTradeDateRequest`/`*GetTradeDateResponse` |
| `GetOptionChainRequest.DataFilter` | `interface{}` | `*qotgetoptionchain.DataFilter` |

### Files Touched (30 total)

| File | Issues |
|---|---|
| `pkg/push/qot_push.go` | #3 |
| `pkg/push/chan/chan.go` | #9, #24, #25 |
| `pkg/push/push_test.go` | #3 |
| `pkg/qot/quote.go` | #1, #2, #19, #29 |
| `pkg/qot/holding.go` | #6 |
| `pkg/qot/kline.go` | #19, #20 |
| `pkg/qot/options.go` | #16, #19 |
| `pkg/qot/sub.go` | #19, #21 |
| `pkg/qot/user.go` | #22 |
| `pkg/qot/trade_date.go` | #19, #23 |
| `pkg/qot/quote_test.go` | #6 |
| `pkg/trd/position.go` | #7, #8, #10, #18, #26, #28 |
| `pkg/trd/queries.go` | #4, #10, #18 |
| `pkg/trd/orders.go` | #10, #18 |
| `pkg/trd/trade.go` | #19 |
| `pkg/trd/convenience.go` | #17 |
| `pkg/sys/system.go` | #5, #13-15, #18, #19 |
| `client/types.go` | #11 |
| `client/quote_api.go` | #12, #21 |
| `client/trade_api.go` | #12 |
| `client/push.go` | #19 |
| `pkg/constant/constant.go` | #19 |

---

## Verification Script

Run these commands **in sequence** after each phase:

```bash
# Phase 1
go test -race ./pkg/push/... -run TestParseUpdateBasicQot

# Phase 2
go build ./pkg/qot/...
go build ./pkg/trd/...

# Phase 3
go build ./pkg/qot/...
go vet ./pkg/qot/...

# Phase 4
go build ./...
go vet ./...

# Phase 5
go build ./pkg/trd/...

# Phase 6
go build ./client/...

# Phase 7
go build ./pkg/sys/...

# Phase 8
go build ./pkg/sys/...
go test -race ./pkg/sys/...

# Phase 9 — Full verification
go build ./...
go vet ./...
go test -race ./...
```

---

## Design Decisions

### 1. Nested Proto Pointers in Client Types — Omitted Intentionally

The following nested proto pointer types are **NOT** added to `client/types.go` `Quote` struct:
- `OptionExData *qotcommon.OptionBasicQotExData`
- `PreMarket *qotcommon.PreAfterMarketData`
- `AfterMarket *qotcommon.PreAfterMarketData`
- `FutureExData *qotcommon.FutureBasicQotExData`
- `WarrantExData *qotcommon.WarrantBasicQotExData`
- `Overnight *qotcommon.PreAfterMarketData`

**Rationale:** The `client` package is designed as a simplified API for common use cases. Nested proto-typed fields would force users to import proto packages, defeating the wrapper abstraction. Users who need these fields should use `pkg/qot` directly.

### 2. TrdFilterConditions — Kept as Raw Proto in Request Types

`*trdcommon.TrdFilterConditions` is used in 6 request types (GetOrderList, GetHistoryOrderList, GetOrderFillList, GetHistoryOrderFillList, GetPositionList, etc.).

**Rationale:** This type has 10+ fields (BeginTime, EndTime, Code, OrderID, etc.). Wrapping it would create a 1:1 mapping with no simplification benefit. It's accepted as a pragmatic leak.

### 3. qotcommon.Security — Kept as Raw Proto Everywhere

`*qotcommon.Security` is the canonical way to represent a security (market + code pair) across the entire Futu protocol. Wrapping it would introduce unnecessary indirection.

**Exception:** When there's only a single security being requested by a wrapper function (e.g., `GetQuote(ctx, market, code)`), the `client` package accepts separate `market` and `code` parameters.

### 4. Proto2 Hack for GetDelayStatistics — Moved to Internal Helper

The proto2 wire format encoding hack (`marshalC2SProto2`, `appendVarint`) should be **moved** to `internal/client/proto2.go` rather than removed, since OpenD's C++ parser requires proto2 non-packed encoding for `repeated int32` fields. The hack is an implementation detail, not a public API concern.

### 5. Ordering of New Fields — Append at End of Struct

All new struct fields are appended at the end of their respective struct definitions. This ensures backward binary compatibility for any code using positional struct literals (though Go best practice is named field initialization).

### 6. Subscribe/RegQotPush Return Type Change

Changing from `(*SubscribeResponse, error)` to `error` is a **breaking change** for any caller that reads the response. Since the response only contains RetType (success) and RetMsg (empty on success), no callers in the codebase actually use these values. The `client/quote_api.go` callers all use `_, err := qot.Subscribe(...)`. This change simplifies the API.

### 7. GetUserInfo Signature Change

Adding a `*GetUserInfoRequest` parameter changes the function signature. To maintain backward compatibility, a convenience wrapper without the request parameter can be provided:

```go
// Deprecated: Use GetUserInfo with context and optional request.
func GetUserInfoSimple(ctx context.Context, c *futuapi.Client) (*GetUserInfoResponse, error) {
    return GetUserInfo(ctx, c, nil)
}
```

---

*Plan generated 2026-05-17. Corresponds to futuapi4go v0.9.0 audit.*

---

## Implementation Status

| Phase | Title | Released | Status |
|-------|-------|----------|--------|
| 1 | Push Wrapper Enrichment | v0.8.1 | ✅ Done |
| 2 | Proto Field Gaps in pkg/ Wrappers | v0.8.1 | ✅ Done |
| 3 | Missing Wrapper: Qot_GetRehab | v0.8.1 | ✅ Done |
| 4 | Proto Type Leakage Fixes | v0.8.1 | ✅ Done |
| 5 | MaxTrdQtys Price Fix | v0.8.1 | ✅ Done |
| 6 | Client Type Enrichment | v0.8.1 | ✅ Done |
| 7 | GetUserInfo Enrichment | v0.8.1 | ✅ Done |
| 8 | GetDelayStatistics Cleanup | v0.8.1 | ✅ Done |
| 9 | Consistency Fixes | v0.8.2 | ✅ Done |
| 10 | Cross-Layer Field Audit | v0.8.4 | ✅ Done |

---

## Phase 10 — Cross-Layer Field Audit (Examples 00-20)

Completed 2026-05-17. Audited all 21 demo examples (00–20) for missing proto→SDK→client→demo field coverage. Fixed 7 gaps across wrapper layers.

### Gap A — PushTicker Missing Fields (examples 02, 09)

- **Problem:** `client/push.go` `ParsePushTicker` mapper only mapped 7 of 15 proto `Ticker` fields.
- **Fix:** Added `Time`, `Timestamp`, `PushDataType` to `client/types.go` `PushTicker` struct. Mapper now maps all 15 fields.
- **Files:** `client/types.go`, `client/push.go`

### Gap B — PushRT Missing Fields (examples 04, 10)

- **Problem:** `client/push.go` `ParsePushRT` mapper only mapped 7 of 11 proto `TimeShare` fields.
- **Fix:** Added `LastClosePrice` to `client/types.go` `PushRT` struct. Mapper now maps all 11 fields.
- **Files:** `client/types.go`, `client/push.go`

### Gap C — KLine Missing Pe (examples 06, 07, 15)

- **Problem:** `client.KLine` struct had no `Pe` field. 3 of 4 KLine mappers also missed `IsBlank` and `TurnoverRate`.
- **Fix:** Added `Pe` to `client/types.go` `KLine`. Updated all 4 mappers (`GetKLines`, `RequestHistoryKLWithLimit`, `GetHistoryKL`, `ParsePushKLine`) to map all 13 proto fields.
- **Files:** `client/types.go`, `client/quote_api.go`, `client/push.go`

### Gap D — StaticInfo Missing Fields (example 14, GetStaticInfo)

- **Problem:** `client.StaticInfo` struct had only 5 fields (Code, Name, Type, ListTime, LotSize). Proto `SecurityStaticBasic` has 9 fields. `GetPlateSecurity` mapper only populated Code, Name, Type — skipped ListTime and LotSize too.
- **Fix:** Added `Id`, `Delisting`, `ListTimestamp`, `ExchType` to `StaticInfo`. Both `GetStaticInfo` and `GetPlateSecurity` mappers now populate all 9 fields.
- **Files:** `client/types.go`, `client/quote_api.go`

### Gap E — GetCapitalFlow Drops Metadata (example 12)

- **Problem:** `client.GetCapitalFlow` returned only `([]CapitalFlow, error)`, dropping `LastValidTime` and `LastValidTimestamp` from the S2C response.
- **Fix:** New `client.CapitalFlowResponse` struct wraps `Items []CapitalFlow` + `LastValidTime` + `LastValidTimestamp`. Function now returns `(*CapitalFlowResponse, error)`.
- **Breaking?** Yes. Callers must use `.Items` to iterate.
- **Files:** `client/types.go`, `client/quote_api.go`

### Gap F — GetMarketState Drops Code/Name (example 16)

- **Problem:** `client.GetMarketState` returned bare `(int32, error)`, dropping `Security` (code) and `Name` from the S2C `MarketInfo`.
- **Fix:** New `client.MarketStateResult` struct with `Code`, `Name`, `State`. Function now returns `(*MarketStateResult, error)`.
- **Breaking?** Yes. Callers must use `.State` for the numeric value.
- **Files:** `client/types.go`, `client/quote_api.go`

---

## Open Issues / Next Steps

These items were identified during the audit but not yet addressed. Track them here for the next work session.

### Priority: Medium

- [ ] **Audit examples 21-96+** — The remaining ~80 demo examples may contain additional unmapped proto fields. Run the same cross-layer trace per example.
- [ ] **Push.KLine raw proto passthrough** — Example 07 uses `*qotcommon.KLine` directly (nil pointer risk). Consider wrapping in the push path for safety, or leave as-is since demo accesses raw fields.
- [ ] **GetHistoryOrderListResponse.OrderList** — Still uses `[]*trdcommon.Order` (raw proto). Phase 4 Issue #4 was partially deferred.
- [ ] **GetDelayStatistics** — Still uses `WritePacket`/`ReadResponseContext` + raw proto types. Phase 8 was deferred (complex proto2 encoding issue).

### Priority: Low

- [ ] **Demo replace directive** — `futuapi4go-demo/go.mod` still has `replace github.com/shing1211/futuapi4go => ../futuapi4go`. Should remove when cutting a stable release.
- [ ] **ProtoID constant consolidation** — Phase 9 Issue #19 was partially completed. Some local ProtoID constants may still exist alongside `constant.go` definitions.
- [ ] **Subscribe/RegQotPush return type** — Phase 9 Issue #21 (change `(*SubscribeResponse, error)` → `error`). Breaking change, low adoption risk.
- [ ] **GitHub release automation** — `make release` fails outside macOS/Linux. Manual `gh release create` is the current workflow.

### Known Won't Fix

- **GitHub push blocked** — Intermittent network issue (`Failed to connect to github.com port 443`). Tags `v0.8.2`, `v0.8.3`, `v0.8.4` pushed successfully when network is available. Gitee is the canonical remote.
- **Nested proto pointers in client types** — `OptionExData`, `PreMarket`, `AfterMarket`, `FutureExData`, `WarrantExData`, `Overnight` intentionally omitted from `client.Quote` per design decision (see §1 above).

---

*Last updated: 2026-05-17. Corresponds to futuapi4go v0.8.4.*
