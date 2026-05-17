# SDK Proto Field Completion — Implementation Plan [COMPLETED — Phase I Scope]

> **Status:** ✅ Phase I scope completed as of v0.8.1–v0.8.3.
> BasicQot, KLine, Ticker, RT, Broker, PushTicker, PriceReminder, CapitalFlow, MarketState, and all other core types now expose 100% of proto fields.
>
> ⚠️ This was the first phase of a larger effort. **Phase II is tracked in `PROTO_FIELD_COMPLETION_PLAN_v2.md`** where several items remain pending (see that file for current status).
>
> **Goal:** Fill all gaps where wrapper structs or mapper functions omit proto fields.
> Direct continuation of the RT fix (v0.8.1), applied systematically across all types.
>
> **Approach:** For each wrapper type: add missing struct fields, add missing mapper assignments.
> No new types, no proto regeneration, no breaking API changes (additive only).

---

## Phase 0 — RT Verification (Already Fixed)

The RT fix is complete. Confirm:
- `pkg/qot/market_data.go` — `RT` struct: `Minute`, `IsBlank`, `Timestamp` ✅
- `client/types.go` — `RT` type: `Minute`, `IsBlank`, `Timestamp` ✅
- `client/quote_api.go` — `GetRT` mapper: all 9 fields ✅

---

## Phase 1 — `pkg/qot/quote.go`: `BasicQot` Enrichment

**File:** `pkg/qot/quote.go`
**Proto:** `qotcommon.BasicQot` (25 fields total)

### 1.1 `BasicQot` struct — add missing fields

**Current (14 fields):**
```go
type BasicQot struct {
    Security       *qotcommon.Security
    Name           string
    IsSuspended    bool
    UpdateTime     string
    HighPrice      float64
    OpenPrice      float64
    LowPrice       float64
    CurPrice       float64
    LastClosePrice float64
    Volume         int64
    Turnover       float64
    TurnoverRate   float64
    Amplitude      float64
}
```

**Add (11 new fields, after `Amplitude`):**
```go
    ListTime         string                            // proto: listTime — 上市日期 (required)
    PriceSpread      float64                           // proto: priceSpread — 价差 (required)
    DarkStatus       int32                             // proto: darkStatus — 暗盘交易状态
    ListTimestamp    float64                           // proto: listTimestamp — 上市日期时间戳
    UpdateTimestamp  float64                           // proto: updateTimestamp — 更新时间戳
    SecStatus        int32                             // proto: secStatus — 股票状态
    OptionExData     *qotcommon.OptionBasicQotExData   // proto: optionExData — 期权特有
    PreMarket        *qotcommon.PreAfterMarketData      // proto: preMarket — 盘前
    AfterMarket      *qotcommon.PreAfterMarketData      // proto: afterMarket — 盘后
    FutureExData     *qotcommon.FutureBasicQotExData    // proto: futureExData — 期货特有
    Overnight        *qotcommon.PreAfterMarketData      // proto: overnight — 夜盘
```

> **Decision:** Nested optional messages (`OptionExData`, `PreMarket`, `AfterMarket`, `FutureExData`, `Overnight`) are exposed as raw proto pointers. Callers who need nested fields call `.GetXxx()` on the proto directly. Avoids proliferation of new sub-wrapper types while making all data accessible.

### 1.2 `GetBasicQot` mapper — add missing assignments

**Current (14 assignments).** **Add (11 new):**
```go
    ListTime:        bq.GetListTime(),
    PriceSpread:     bq.GetPriceSpread(),
    DarkStatus:      bq.GetDarkStatus(),
    ListTimestamp:   bq.GetListTimestamp(),
    UpdateTimestamp: bq.GetUpdateTimestamp(),
    SecStatus:       bq.GetSecStatus(),
    OptionExData:    bq.GetOptionExData(),
    PreMarket:       bq.GetPreMarket(),
    AfterMarket:     bq.GetAfterMarket(),
    FutureExData:    bq.GetFutureExData(),
    Overnight:       bq.GetOvernight(),
```

### 1.3 `pkg/qot/quote_test.go` — update tests

**Test:** `TestBasicQotStructFields` (line ~85)
Rename to `TestBasicQotStructFieldsComplete`, add struct init and assertions for all 25 fields.

---

## Phase 2 — `pkg/qot/market_data.go`: `Ticker`, `Broker` Enrichment

**File:** `pkg/qot/market_data.go`

### 2.1 `Ticker` struct — add `PushDataType`

**Proto:** `qotcommon.Ticker` (12 fields). Current: 10 fields, missing `PushDataType int32`.

**Add (1 new field):**
```go
    PushDataType int32  // proto: pushDataType
```

**Mapper:** After adding field, add `PushDataType: t.GetPushDataType()` to `GetTicker`.

### 2.2 `Broker` struct — add `OrderID`

**Proto:** `qotcommon.Broker` (5 fields). Current: 4 fields, missing `OrderID int64` (SF market).

**Add (1 new field):**
```go
    OrderID int64  // proto: orderID — SF market only
```

**Mapper:** After adding field, add `OrderID: b.GetOrderID()` to `GetBroker`.

### 2.3 Tests

- `TestTickerFields` (line ~274): add `PushDataType` to struct init and assertions.
- `TestBrokerFields` (line ~334): add `OrderID` to struct init and assertions.

---

## Phase 3 — `pkg/trd/position.go`: `Position`, `Funds` Enrichment

**File:** `pkg/trd/position.go`

### 3.1 `Position` struct — add `PositionSide`

**Proto:** `trdcommon.Position` (28 fields). Current: 24 fields, missing `PositionSide int32` (required in proto).

**Add (1 new field, after `PositionID`):**
```go
    PositionSide int32  // proto: positionSide — 持仓方向 (required)
```

**Mapper:** Add `PositionSide: p.GetPositionSide()` to `GetPositionList`.

### 3.2 `Funds` struct — add asset breakdown fields

**Proto:** `trdcommon.Funds` (36 fields). Current: 29 fields, missing 3.

**Add (3 new fields, after `DtStatus`):**
```go
    SecuritiesAssets float64  // proto: securitiesAssets — 证券资产净值
    FundAssets      float64  // proto: fundAssets — 基金资产净值
    BondAssets      float64  // proto: bondAssets — 债券资产净值
```

**Mapper:** Add 3 assignments to `GetFunds`.

### 3.3 Tests

- `TestPositionStructFields` (line ~105): add `PositionSide` assertions.
- `TestFundsStructFields` (line ~86): add `SecuritiesAssets`, `FundAssets`, `BondAssets` assertions.

---

## Phase 4 — `pkg/trd/max_qtys.go`: No Changes Needed

`MaxTrdQtysInfo` already has all 7 proto fields. The gap is in the **client layer** (Phase 6.4).

---

## Phase 5 — `pkg/push/qot_push.go`: `UpdateBasicQot` Enrichment

**File:** `pkg/push/qot_push.go`

### 5.1 `UpdateBasicQot` struct — add missing fields

**Proto:** `qotcommon.BasicQot`. Current: 9 fields.

**Add (4 new fields):**
```go
    IsSuspended     bool     // proto: isSuspended
    LastClosePrice  float64  // proto: lastClosePrice
    UpdateTime      string   // proto: updateTime
    UpdateTimestamp float64  // proto: updateTimestamp
```

### 5.2 `ParseUpdateBasicQot` mapper — add 4 assignments

### 5.3 `pkg/push/push_test.go` — `TestParseUpdateBasicQotValidData`

Add assertions for 4 new fields.

---

## Phase 6 — `client/types.go`: Multiple Types Enrichment

**File:** `client/types.go`

### 6.1 `Quote` struct — add 2 fields

**Add (2 new fields):**
```go
    IsSuspended bool  `json:"isSuspended"`  // from BasicQot.IsSuspended
    SecStatus  int32 `json:"secStatus"`   // from BasicQot.SecStatus
```

> **Decision:** Advanced fields (`ListTime`, `PriceSpread`, nested messages) intentionally NOT added to `client.Quote` to keep the JSON API surface lean. Power users use `pkg/qot.BasicQot` directly.

### 6.2 `Position` struct — add 1 field

**Add:**
```go
    PositionSide int32 `json:"positionSide"`  // from pkg/trd.Position.PositionSide
```

### 6.3 `Funds` struct — add 3 fields

**Add:**
```go
    SecuritiesAssets float64 `json:"securitiesAssets"`  // 证券资产净值
    FundAssets      float64 `json:"fundAssets"`        // 基金资产净值
    BondAssets      float64 `json:"bondAssets"`         // 债券资产净值
```

### 6.4 `MaxTrdQtysInfo` struct — add 2 fields

**Proto has 7 fields. Current client type has 6. Add:**
```go
    LongRequiredIM   float64 `json:"longRequiredIM"`    // 多头所需初始保证金
    ShortRequiredIM  float64 `json:"shortRequiredIM"`   // 空头所需初始保证金
```

### 6.5 `MarginRatioInfo` struct — add 7 fields

**Proto has 12 fields. Current client type has 6. Add:**
```go
    ShortPoolRemain float64 `json:"shortPoolRemain"`  // 融券池剩余
    AlertLongRatio  float64 `json:"alertLongRatio"`   // 融资预警比率
    AlertShortRatio float64 `json:"alertShortRatio"`  // 融券预警比率
    McmLongRatio    float64 `json:"mcmLongRatio"`     // 融资margin call保证金率
    McmShortRatio   float64 `json:"mcmShortRatio"`   // 融券margin call保证金率
    MmLongRatio     float64 `json:"mmLongRatio"`     // 融资维持保证金率
    MmShortRatio    float64 `json:"mmShortRatio"`    // 融券维持保证金率
```

### 6.6 `PushTicker` struct — add 5 fields

**Proto:** `qotcommon.Ticker`. Current: 9 fields.

**Add:**
```go
    Sequence  int64   `json:"sequence"`   // from Ticker.Sequence
    Dir       int32   `json:"dir"`        // from Ticker.Dir
    RecvTime  float64 `json:"recvTime"`   // from Ticker.RecvTime
    Type      int32   `json:"type"`       // from Ticker.Type
    TypeSign  int32   `json:"typeSign"`   // from Ticker.TypeSign
```

### 6.7 `PushRT` struct — add 3 fields

**Proto:** `qotcommon.TimeShare`. Current: 8 fields.

**Add:**
```go
    Minute    int32   `json:"minute"`    // from TimeShare.Minute
    IsBlank   bool    `json:"isBlank"`   // from TimeShare.IsBlank
    Timestamp float64 `json:"timestamp"`  // from TimeShare.Timestamp
```

### 6.8 `BrokerItem` struct — add 1 field

**Add:**
```go
    Name string `json:"name"`  // from Broker.name — 显示经纪商名称
```

### 6.9 BONUS: `KLine` struct — add 2 fields (also discovered)

**Proto:** `qotcommon.KLine` has `turnoverRate` and `isBlank`. `client/types.go` `KLine` lacks `IsBlank`. `pkg/qot/KLine` lacks `TurnoverRate`.

**Add to `client/types.go` `KLine`:**
```go
    IsBlank bool `json:"isBlank"`  // from KLine.IsBlank
```

**Add to `pkg/qot/kline.go` `KLine` struct and mapper:**
```go
    TurnoverRate float64  // proto: turnoverRate
    // mapper: TurnoverRate: kl.GetTurnoverRate()
```

---

## Phase 7 — Mappers: `client/quote_api.go` & `client/trade_api.go`

| File | Function | New Assignments |
|------|----------|----------------|
| `client/quote_api.go` | `GetQuote` | `IsSuspended`, `SecStatus` |
| `client/quote_api.go` | `GetKLines` | `IsBlank` |
| `client/trade_api.go` | `GetPositionList` | `PositionSide` |
| `client/trade_api.go` | `GetFunds` | `SecuritiesAssets`, `FundAssets`, `BondAssets` |
| `client/trade_api.go` | `GetMaxTrdQtys` | `LongRequiredIM`, `ShortRequiredIM` |
| `client/trade_api.go` | `GetMarginRatio` | `ShortPoolRemain`, `AlertLongRatio`, `AlertShortRatio`, `McmLongRatio`, `McmShortRatio`, `MmLongRatio`, `MmShortRatio` |

---

## Phase 8 — Demo Examples: Update Output

| Example | Current Output | Update |
|---------|---------------|--------|
| `01_quote` | shows price/open/high/low/vol | Add `isSuspended`, `secStatus` |
| `02_ticker` | shows price/volume | Add `sequence`, `dir`, `recvTime`, `type`, `typeSign` |
| `05_broker` | shows ID/Name/Pos/Volume | Add `orderID` |
| `04_rt` | already enriched ✅ | — |

---

## Complete Change Summary

### Struct changes (adding fields)

| File | Struct | New Fields | Count |
|------|--------|-----------|-------|
| `pkg/qot/quote.go` | `BasicQot` | `ListTime`, `PriceSpread`, `DarkStatus`, `ListTimestamp`, `UpdateTimestamp`, `SecStatus`, `OptionExData*`, `PreMarket*`, `AfterMarket*`, `FutureExData*`, `Overnight*` | 11 |
| `pkg/qot/market_data.go` | `Ticker` | `PushDataType` | 1 |
| `pkg/qot/market_data.go` | `Broker` | `OrderID` | 1 |
| `pkg/qot/kline.go` | `KLine` | `TurnoverRate` | 1 |
| `pkg/trd/position.go` | `Position` | `PositionSide` | 1 |
| `pkg/trd/position.go` | `Funds` | `SecuritiesAssets`, `FundAssets`, `BondAssets` | 3 |
| `pkg/push/qot_push.go` | `UpdateBasicQot` | `IsSuspended`, `LastClosePrice`, `UpdateTime`, `UpdateTimestamp` | 4 |
| `client/types.go` | `Quote` | `IsSuspended`, `SecStatus` | 2 |
| `client/types.go` | `Position` | `PositionSide` | 1 |
| `client/types.go` | `Funds` | `SecuritiesAssets`, `FundAssets`, `BondAssets` | 3 |
| `client/types.go` | `MaxTrdQtysInfo` | `LongRequiredIM`, `ShortRequiredIM` | 2 |
| `client/types.go` | `MarginRatioInfo` | `ShortPoolRemain`, `AlertLongRatio`, `AlertShortRatio`, `McmLongRatio`, `McmShortRatio`, `MmLongRatio`, `MmShortRatio` | 7 |
| `client/types.go` | `PushTicker` | `Sequence`, `Dir`, `RecvTime`, `Type`, `TypeSign` | 5 |
| `client/types.go` | `PushRT` | `Minute`, `IsBlank`, `Timestamp` | 3 |
| `client/types.go` | `BrokerItem` | `Name` | 1 |
| `client/types.go` | `KLine` | `IsBlank` | 1 |
| **Total** | | | **47** |

### Mapper changes (adding assignments)

| File | Function | New |
|------|----------|-----|
| `pkg/qot/quote.go` | `GetBasicQot` | 11 |
| `pkg/qot/market_data.go` | `GetTicker` | 1 |
| `pkg/qot/market_data.go` | `GetBroker` | 1 |
| `pkg/qot/kline.go` | `GetKLResponse` mapper | 1 |
| `pkg/trd/position.go` | `GetPositionList` | 1 |
| `pkg/trd/position.go` | `GetFunds` | 3 |
| `pkg/push/qot_push.go` | `ParseUpdateBasicQot` | 4 |
| `client/quote_api.go` | `GetQuote` | 2 |
| `client/quote_api.go` | `GetKLines` | 1 |
| `client/trade_api.go` | `GetPositionList` | 1 |
| `client/trade_api.go` | `GetFunds` | 3 |
| `client/trade_api.go` | `GetMaxTrdQtys` | 2 |
| `client/trade_api.go` | `GetMarginRatio` | 7 |
| **Total** | | **38** |

### Test changes

| File | Tests |
|------|-------|
| `pkg/qot/quote_test.go` | `TestBasicQotStructFields` → complete, `TestTickerFields`, `TestBrokerFields` |
| `pkg/trd/trade_test.go` | `TestPositionStructFields`, `TestFundsStructFields` |
| `pkg/push/push_test.go` | `TestParseUpdateBasicQotValidData` |

### Demo changes

| File | Change |
|------|--------|
| `examples/01_quote/main.go` | Show `isSuspended`, `secStatus` |
| `examples/02_ticker/main.go` | Show `sequence`, `dir`, `recvTime`, `type`, `typeSign` |
| `examples/05_broker/main.go` | Show `orderID` |

---

## Decisions Made

1. **Nested optional messages in `BasicQot`** — Exposed as raw proto pointers (`*qotcommon.Xxx`). Callers use proto getters directly. Avoids proliferation of sub-types.

2. **`Quote` in `client/types.go`** — Only `IsSuspended` and `SecStatus` added. Advanced fields omitted from public JSON API. Power users use `pkg/qot.BasicQot` directly.

3. **`OBItem.OrderCount` type mismatch** — `client/types.go` uses `int64`, proto has `int32`. Out of scope.

4. **`BrokerItem.BrokerID`** — Maps to proto `Broker.id`. Left as-is.

---

## Verification

```bash
go build ./...
go vet ./...
go test ./pkg/qot/... -run "BasicQot|TickerFields|BrokerFields" -v
go test ./pkg/trd/... -run "PositionStructFields|FundsStructFields" -v
go test ./pkg/push/... -run "UpdateBasicQot" -v
go test ./client/... -count=1
go test ./...
```

---

## Execution Order

1. Phase 1 — `pkg/qot/quote.go`: BasicQot (+11 fields)
2. Phase 2 — `pkg/qot/market_data.go`: Ticker + Broker (+2 fields)
3. Phase 3 — `pkg/trd/position.go`: Position + Funds (+4 fields)
4. Phase 4 — `pkg/trd/max_qtys.go`: No changes (gap is in client layer)
5. Phase 5 — `pkg/push/qot_push.go`: UpdateBasicQot (+4 fields)
6. Phase 6 — `client/types.go`: All 13 client types (+28 fields)
7. Phase 7 — `client/*_api.go`: All mappers (+15 assignments)
8. Phase 8 — Demo examples: Update output formatting
9. Tests — Update 4 test functions across 3 files
10. `docs/CHANGELOG.md` — Update `[Unreleased]` section
11. Verify — `go build ./... && go vet ./... && go test ./...`
12. Commit & tag v0.8.1
