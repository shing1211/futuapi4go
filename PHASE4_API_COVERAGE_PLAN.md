# Phase 4: Missing High-Level API Coverage — Detailed Design & Implementation Plan

*Generated: 2026-05-21*
*SDK Version: v0.14.0 (Futu Protocol v10.8.6808)*

---

## 1. Overview

Phase 4 adds 25 new quote/market-data API implementations to `pkg/qot/`, plus corresponding `client/quote_api.go` wrappers and `client/fluent_api.go` methods. All trading APIs (pkg/trd/) are already fully covered — zero gaps there.

### Audit Summary

| Layer | Total APIs | Implemented | Missing |
|-------|-----------|-------------|---------|
| **pkg/qot/** (Quote) | 62 request/response | 37 | **25** (ProtoIDs 3227-3251) |
| **pkg/trd/** (Trade) | 16 | 16 | **0** (fully covered) |
| **pkg/sys/** (System) | 7 | 7 | **0** (fully covered) |
| **client/ wrappers** | — | — | **25 quote + 1 existing (RequestHistoryKLQuota) + 8 trade convenience** |
| **Screening APIs** | 5 | 0 | **5 deferred** (no ProtoID assigned yet) |

---

## 2. Scope

### 2.1 In Scope — 25 New Quote APIs (ProtoIDs 3227-3251)

All proto pb packages already exist in `pkg/pb/` with generated Go bindings. ProtoIDs already defined in `pkg/constant/constant.go`.

### 2.2 In Scope — Client Wrapper Gaps

| Missing Wrapper | pkg Function | Priority |
|---|---|---|
| `RequestHistoryKLQuota` | `qot.RequestHistoryKLQuota` | HIGH |
| `QuickBuy` | `trd.QuickBuy` | MEDIUM |
| `QuickSell` | `trd.QuickSell` | MEDIUM |
| `QuickMarketBuy` | `trd.QuickMarketBuy` | MEDIUM |
| `QuickMarketSell` | `trd.QuickMarketSell` | MEDIUM |
| `GetPositions` | `trd.GetPositions` | MEDIUM |
| `GetTodayFills` | `trd.GetTodayFills` | MEDIUM |
| `GetTodayOrders` | `trd.GetTodayOrders` | MEDIUM |
| `GetAccountFunds` | `trd.GetAccountFunds` | MEDIUM |

### 2.3 Deferred — 5 Screening APIs (No ProtoID)

| Proto File | PB Package | Reason |
|---|---|---|
| Qot_StockScreen.proto | qotstockscreen | No ProtoID in constant.go |
| Qot_WarrantScreen.proto | qotwarrantscreen | No ProtoID in constant.go |
| Qot_OptionScreen.proto | qotoptionscreen | No ProtoID in constant.go |
| Qot_GetFinancialsEarningsPriceMove.proto | qotgetfinancialsearnpricemove | No ProtoID in constant.go |
| Qot_GetFinancialsEarningsPriceHistory.proto | qotgetfinancialsearnpricehist | No ProtoID in constant.go |

### 2.4 Not In Scope (Intentional Exclusions)

- `NewHistoryKLineIterator` — utility constructor, not an API call
- `NewAuditLogger`, `ValidateOrder`, `HasErrors` — utility functions, not API calls
- `CancelAllOrders` — already covered by `client/trade_api.go:CancelAllOrder`

---

## 3. File Organization

New APIs organized into **9 new files** in `pkg/qot/`:

| File | APIs | Category |
|------|------|----------|
| `financials.go` | GetFinancialsStatements (3227), GetFinancialsRevenueBreakdown (3228) | Financial Data |
| `research.go` | GetResearchAnalystConsensus (3229), GetResearchRatingSummary (3230), GetResearchMorningstarReport (3231) | Research/Analyst |
| `valuation.go` | GetValuationDetail (3232), GetValuationPlateStockList (3233) | Valuation |
| `corporate.go` | GetCorporateActionsDividends (3234), GetCorporateActionsBuybacks (3235), GetCorporateActionsStockSplits (3236) | Corporate Actions |
| `shareholders.go` | GetShareholdersOverview (3237), GetShareholdersHoldingChanges (3238), GetShareholdersHolderDetail (3239), GetShareholdersInstitutional (3240) | Shareholders |
| `insider.go` | GetInsiderHolderList (3241), GetInsiderTradeList (3242) | Insider Trading |
| `company.go` | GetCompanyProfile (3243), GetCompanyExecutives (3244), GetCompanyExecutiveBackground (3245), GetCompanyOperationalEfficiency (3246) | Company Info |
| `shortselling.go` | GetTopTenBuySellBrokers (3247), GetDailyShortVolume (3248), GetShortInterest (3249) | Short Selling & Brokers |
| `option_extra.go` | GetOptionVolatility (3250), GetOptionExerciseProbability (3251) | Options (Extended) |

---

## 4. Implementation Pattern

Every new function follows this exact pattern (derived from existing codebase):

```go
func Xxx(ctx context.Context, c *futuapi.Client, req *XxxRequest) (*XxxResponse, error) {
    if req == nil {
        return nil, fmt.Errorf("Xxx: request is nil")
    }
    if req.Security == nil {
        return nil, fmt.Errorf("Xxx: security is required")
    }

    c2s := &somepb.C2S{
        Security: req.Security,
        // required scalars: &req.Field
        // optional scalars: conditional &req.Field
        // slices/messages: direct
    }
    if req.OptField != 0 {
        c2s.OptField = &req.OptField
    }
    if req.OptStr != "" {
        c2s.OptStr = &req.OptStr
    }

    pkt := &somepb.Request{C2S: c2s}
    var rsp somepb.Response

    if err := c.RequestContext(ctx, ProtoID_Xxx, pkt, &rsp); err != nil {
        return nil, err
    }
    if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
        return nil, wrapError("Xxx", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
    }
    s2c := rsp.S2C
    if s2c == nil {
        return nil, wrapError("Xxx", int32(common.RetType_RetType_Unknown), "s2c is nil")
    }

    result := &XxxResponse{...}
    for _, item := range s2c.ItemList {
        if item == nil {
            continue
        }
        result.Items = append(result.Items, &XxxItem{
            Name:  util.ProtoStr(item.Name),
            Value: util.ProtoFloat64(item.Value),
        })
    }
    return result, nil
}
```

### Key Rules

1. **`util.ProtoXxx()` for ALL scalar proto field reads** — never dereference proto pointer fields directly
2. **`wrapError()` for all error returns** — never construct `FutuError` directly
3. **Nil-guard hierarchy**: request nil → required fields → S2C nil → list item nil
4. **C2S field assignment**: slices/messages direct, required scalars `&field`, optional scalars conditional
5. **Pass-through proto types** for complex nested data (FinancialFieldInfo, FinancialReport, etc.)
6. **Custom request/response types** — never expose raw C2S/S2C to callers
7. **No comments** in code (per AGENTS.md code style)
8. **Context as FIRST parameter** to all public APIs

---

## 5. Detailed API Specifications

### 5.1 financials.go

#### GetFinancialsStatements (ProtoID 3227)

**PB Package**: `qotgetfinancialsstatements`

**C2S Fields**:
| Field | Type | Required | Notes |
|-------|------|----------|-------|
| Security | *qotcommon.Security | YES | Stock |
| StatementType | *qotcommon.FinancialStatementsType | NO | Enum: Income=1, BalanceSheet=2, CashFlow=3 |
| FinancialType | *qotcommon.F10Type | NO | Enum: 0-7, 9-11, default 10(QuarterlyAnnual) |
| CurrencyCode | *string | NO | ISO 4217 (CNY, USD, HKD, etc.) |
| NextKey | *string | NO | Pagination key |
| Num | *int32 | NO | Page size, default 10, range 1-50 |

**S2C Fields**:
| Field | Type | Notes |
|-------|------|-------|
| StructureList | []*FinancialFieldInfo | Field structure list |
| ReportList | []*FinancialReport | Financial report data |
| NextKey | *string | Pagination key, "-1" = no more |

**Custom Types**:
```go
type GetFinancialsStatementsRequest struct {
    Security       *qotcommon.Security
    StatementType  int32
    FinancialType  int32
    CurrencyCode   string
    NextKey        string
    Num            int32
}

type GetFinancialsStatementsResponse struct {
    StructureList []*qotgetfinancialsstatements.FinancialFieldInfo
    ReportList    []*qotgetfinancialsstatements.FinancialReport
    NextKey       string
}
```

#### GetFinancialsRevenueBreakdown (ProtoID 3228)

**PB Package**: `qotgetfinancialrevenuebreakdown`

**C2S Fields**:
| Field | Type | Required | Notes |
|-------|------|----------|-------|
| Security | *qotcommon.Security | YES | Stock |
| Date | *uint32 | NO | Unix timestamp, 0=latest |
| FinancialType | *qotcommon.F10Type | NO | Enum: 0-7, 9, default 0 |
| CurrencyCode | *string | NO | ISO 4217 |

**S2C Fields**:
| Field | Type | Notes |
|-------|------|-------|
| Period | *string | e.g. "2024/Q3", "2024/FY" |
| BreakdownList | []*RevenueBreakdownGroup | Revenue breakdown data |
| CurrencyCode | *string | ISO 4217 |
| ScreenDateList | []*ScreenDate | Available date list |

**Custom Types**:
```go
type GetFinancialsRevenueBreakdownRequest struct {
    Security      *qotcommon.Security
    Date          uint32
    FinancialType int32
    CurrencyCode  string
}

type GetFinancialsRevenueBreakdownResponse struct {
    Period         string
    BreakdownList  []*qotgetfinancialrevenuebreakdown.RevenueBreakdownGroup
    CurrencyCode   string
    ScreenDateList []*qotgetfinancialrevenuebreakdown.ScreenDate
}
```

---

### 5.2 research.go

#### GetResearchAnalystConsensus (ProtoID 3229)

**PB Package**: `qotgetresearchanalystconsensus`

**C2S Fields**:
| Field | Type | Required |
|-------|------|----------|
| Security | *qotcommon.Security | YES |

**S2C Fields**:
| Field | Type | Notes |
|-------|------|-------|
| Highest | *float64 | Highest target price |
| Average | *float64 | Average target price |
| Lowest | *float64 | Lowest target price |
| Rating | *qotcommon.ResearchRatingType | Enum: Sell/Hold/Buy etc. |
| Total | *int32 | Number of analysts |
| UpdateTime | *int64 | Timestamp (seconds) |
| UpdateTimeStr | *string | YYYY-MM-DD format |
| Buy | *float64 | Buy rating % |
| Hold | *float64 | Hold rating % |
| Sell | *float64 | Sell rating % |
| StrongBuy | *float64 | Strong Buy % (non-US only) |
| Underperform | *float64 | Underperform % (non-US only) |

**Custom Types**:
```go
type GetResearchAnalystConsensusRequest struct {
    Security *qotcommon.Security
}

type GetResearchAnalystConsensusResponse struct {
    Highest       float64
    Average       float64
    Lowest        float64
    Rating        int32
    Total         int32
    UpdateTime    int64
    UpdateTimeStr string
    Buy           float64
    Hold          float64
    Sell          float64
    StrongBuy     float64
    Underperform  float64
}
```

#### GetResearchRatingSummary (ProtoID 3230)

**PB Package**: `qotgetresearchratingsummary`

**C2S Fields**:
| Field | Type | Required | Notes |
|-------|------|----------|-------|
| Security | *qotcommon.Security | YES | Stock |
| RatingDimensionType | *qotcommon.ResearchRatingDimensionType | NO | Default: institution |
| Uid | *string | NO | Empty=summary list, non-empty=detail |
| NextKey | *string | NO | Pagination |
| Num | *int32 | NO | Page size, default 10, range 1-20 |

**S2C Fields**: Complex nested types (InstInfo, AnalystInfo, RatingDetailItem, etc.)

**Custom Types**:
```go
type GetResearchRatingSummaryRequest struct {
    Security            *qotcommon.Security
    RatingDimensionType int32
    Uid                 string
    NextKey             string
    Num                 int32
}

type GetResearchRatingSummaryResponse struct {
    RatingSummaryList []*qotgetresearchratingsummary.RatingSummaryItem
    NextKey           string
}
```

#### GetResearchMorningstarReport (ProtoID 3231)

**PB Package**: `qotgetresearchmorningstarrpt`

**C2S Fields**:
| Field | Type | Required |
|-------|------|----------|
| Security | *qotcommon.Security | YES |

**S2C Fields**: Very rich — 25 fields including star rating, fair value, economic moat, uncertainty, financial health, bull/bear says, analyst notes, PDF URL, etc. All pass-through as raw S2C.

**Custom Types**:
```go
type GetResearchMorningstarReportRequest struct {
    Security *qotcommon.Security
}

type GetResearchMorningstarReportResponse struct {
    S2C *qotgetresearchmorningstarrpt.S2C
}
```

Note: Morningstar report S2C has 25+ fields with deeply nested types. Rather than wrapping each field, we pass through the entire S2C struct. Users access fields with `util.ProtoXxx()` helpers.

---

### 5.3 valuation.go

#### GetValuationDetail (ProtoID 3232)

**PB Package**: `qotgetvaluationdetail`

**C2S Fields**:
| Field | Type | Required | Notes |
|-------|------|----------|-------|
| Security | *qotcommon.Security | YES | Stock |
| ValuationType | *qotcommon.ValuationType | NO | Default 0 (recommended) |
| IntervalType | *qotcommon.ValuationIntervalType | NO | Historical data period |

**S2C Fields**: Complex nested — ValuationType, LastUpdateTime, Trend, MarketDistribution, PlateDistribution, ProfitGrowthRate. All pass-through.

**Custom Types**:
```go
type GetValuationDetailRequest struct {
    Security      *qotcommon.Security
    ValuationType int32
    IntervalType  int32
}

type GetValuationDetailResponse struct {
    S2C *qotgetvaluationdetail.S2C
}
```

#### GetValuationPlateStockList (ProtoID 3233)

**PB Package**: `qotgetvaluationplatestocklist`

**C2S Fields**:
| Field | Type | Required | Notes |
|-------|------|----------|-------|
| Security | *qotcommon.Security | YES | Stock/plate |
| ValuationType | *qotcommon.ValuationType | NO | 1-3, default 1(PE) |
| NextKey | *string | NO | Pagination |
| Num | *int32 | NO | Page size, default 10, range 1-50 |
| SortType | *qotcommon.SortType | NO | Sort direction |
| SortId | *qotcommon.SortField | NO | Sort column |
| FilterSecurity | *qotcommon.Security | NO | Plate filter |

**S2C Fields**: Count, StockList, NextKey, PlateList. Pass-through.

**Custom Types**:
```go
type GetValuationPlateStockListRequest struct {
    Security       *qotcommon.Security
    ValuationType  int32
    NextKey        string
    Num            int32
    SortType       int32
    SortId         int32
    FilterSecurity *qotcommon.Security
}

type GetValuationPlateStockListResponse struct {
    Count     int32
    StockList []*qotgetvaluationplatestocklist.StockItem
    NextKey   string
    PlateList []*qotgetvaluationplatestocklist.PlateItem
}
```

---

### 5.4 corporate.go

#### GetCorporateActionsDividends (ProtoID 3234)

**PB Package**: `qotgetcorporateactionsdividends`

**C2S**: Security (required)
**S2C**: DividendList []*DividendItem

```go
type GetCorporateActionsDividendsRequest struct {
    Security *qotcommon.Security
}

type GetCorporateActionsDividendsResponse struct {
    DividendList []*qotgetcorporateactionsdividends.DividendItem
}
```

#### GetCorporateActionsBuybacks (ProtoID 3235)

**PB Package**: `qotgetcorporateactionsbuybacks`

**C2S**: Security (required), NextKey (optional), Num (optional)
**S2C**: HkBuyBackList, ABuyBackList, NextKey

```go
type GetCorporateActionsBuybacksRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetCorporateActionsBuybacksResponse struct {
    HkBuyBackList []*qotgetcorporateactionsbuybacks.HKBuyBackItem
    ABuyBackList  []*qotgetcorporateactionsbuybacks.ABuyBackItem
    NextKey       string
}
```

#### GetCorporateActionsStockSplits (ProtoID 3236)

**PB Package**: `qotgetcorporateactionsstocksplits`

**C2S**: Security (required), NextKey (optional), Num (optional)
**S2C**: SplitItemList, NextKey

```go
type GetCorporateActionsStockSplitsRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetCorporateActionsStockSplitsResponse struct {
    SplitItemList []*qotgetcorporateactionsstocksplits.StockSplitItem
    NextKey       string
}
```

---

### 5.5 shareholders.go

#### GetShareholdersOverview (ProtoID 3237)

**PB Package**: `qotgetshareholdersoverview`

**C2S**: Security (required), PeriodId (optional, 0=latest)
**S2C**: MainHolderInfoList, HolderTypeInfoList, HoldingPeriodList

```go
type GetShareholdersOverviewRequest struct {
    Security *qotcommon.Security
    PeriodId int32
}

type GetShareholdersOverviewResponse struct {
    MainHolderInfoList []*qotgetshareholdersoverview.OwnershipStaticInfo
    HolderTypeInfoList []*qotgetshareholdersoverview.OwnershipStaticInfo
    HoldingPeriodList  []*qotgetshareholdersoverview.HoldingPeriodItem
}
```

#### GetShareholdersHoldingChanges (ProtoID 3238)

**PB Package**: `qotgetshareholdersholdingchanges`

**C2S**: Security (required), NextKey, Num, SortType, SortColumn, FilterType (all optional)
**S2C**: ItemList, NextKey

```go
type GetShareholdersHoldingChangesRequest struct {
    Security    *qotcommon.Security
    NextKey     string
    Num         int32
    SortType    int32
    SortColumn  int32
    FilterType  int32
}

type GetShareholdersHoldingChangesResponse struct {
    ItemList []*qotgetshareholdersholdingchanges.OwnerListItem
    NextKey  string
}
```

#### GetShareholdersHolderDetail (ProtoID 3239)

**PB Package**: `qotgetshareholdersholderdetail`

**C2S**: Security (required), RequestType, NextKey, Num, SortColumn, SortType, PeriodId, HolderId (all optional)
**S2C**: UpdateTime, UpdateTimeStr, NextKey, ItemList

```go
type GetShareholdersHolderDetailRequest struct {
    Security    *qotcommon.Security
    RequestType int32
    NextKey     string
    Num         int32
    SortColumn  int32
    SortType    int32
    PeriodId    int32
    HolderId    int32
}

type GetShareholdersHolderDetailResponse struct {
    UpdateTime    uint64
    UpdateTimeStr string
    NextKey       string
    ItemList      []*qotgetshareholdersholderdetail.OwnershipDetailItem
}
```

#### GetShareholdersInstitutional (ProtoID 3240)

**PB Package**: `qotgetshareholdersinstitutional`

**C2S**: Security (required), NextKey, Num (optional)
**S2C**: UpdateTime, UpdateTimeStr, NextKey, ItemList

```go
type GetShareholdersInstitutionalRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetShareholdersInstitutionalResponse struct {
    UpdateTime    uint64
    UpdateTimeStr string
    NextKey       string
    ItemList      []*qotgetshareholdersinstitutional.InstitutionHolderItem
}
```

---

### 5.6 insider.go

#### GetInsiderHolderList (ProtoID 3241)

**PB Package**: `qotgetinsiderholderlist`

**C2S**: Security (required), NextKey, Num (optional, range 1-20)
**S2C**: ItemList, AllCount, NextKey, InsiderTotalCount, InsiderBoughtCount, InsiderSoldCount

```go
type GetInsiderHolderListRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetInsiderHolderListResponse struct {
    ItemList           []*qotgetinsiderholderlist.OwnerInsiderHolderItem
    AllCount           int32
    NextKey            string
    InsiderTotalCount  int32
    InsiderBoughtCount int32
    InsiderSoldCount   int32
}
```

#### GetInsiderTradeList (ProtoID 3242)

**PB Package**: `qotgetinsidertradelist`

**C2S**: Security (required), HolderId (optional), NextKey, Num (optional)
**S2C**: ItemList, AllCount, NextKey

```go
type GetInsiderTradeListRequest struct {
    Security *qotcommon.Security
    HolderId int64
    NextKey  string
    Num      int32
}

type GetInsiderTradeListResponse struct {
    ItemList []*qotgetinsidertradelist.OwnerInsiderTradeItem
    AllCount int32
    NextKey  string
}
```

---

### 5.7 company.go

#### GetCompanyProfile (ProtoID 3243)

**PB Package**: `qotgetcompanyprofile`

**C2S**: Security (required)
**S2C**: ItemList []*CompanyLabItem

```go
type GetCompanyProfileRequest struct {
    Security *qotcommon.Security
}

type GetCompanyProfileResponse struct {
    ItemList []*qotgetcompanyprofile.CompanyLabItem
}
```

#### GetCompanyExecutives (ProtoID 3244)

**PB Package**: `qotgetcompanyexecutives`

**C2S**: Security (required)
**S2C**: DirectorList []*DirectorInfo

```go
type GetCompanyExecutivesRequest struct {
    Security *qotcommon.Security
}

type GetCompanyExecutivesResponse struct {
    DirectorList []*qotgetcompanyexecutives.DirectorInfo
}
```

#### GetCompanyExecutiveBackground (ProtoID 3245)

**PB Package**: `qotgetcompanyexecutivebackground`

**C2S**: Security (required), LeaderName (optional)
**S2C**: BriefBackground *string

```go
type GetCompanyExecutiveBackgroundRequest struct {
    Security   *qotcommon.Security
    LeaderName string
}

type GetCompanyExecutiveBackgroundResponse struct {
    BriefBackground string
}
```

#### GetCompanyOperationalEfficiency (ProtoID 3246)

**PB Package**: `qotgetcompanyoperationalefficiency`

**C2S**: Security (required), NextKey, Num, CurrencyCode, FinancialType (all optional)
**S2C**: ItemList, NextKey, CurrencyCode

```go
type GetCompanyOperationalEfficiencyRequest struct {
    Security      *qotcommon.Security
    NextKey       string
    Num           int32
    CurrencyCode  string
    FinancialType int32
}

type GetCompanyOperationalEfficiencyResponse struct {
    ItemList     []*qotgetcompanyoperationalefficiency.OperationalEfficiencyItem
    NextKey      string
    CurrencyCode string
}
```

---

### 5.8 shortselling.go

#### GetTopTenBuySellBrokers (ProtoID 3247)

**PB Package**: `qotgettoptenbuysellbrokers`

**C2S**: Security (required), DaysBefore (optional, 0=realtime)
**S2C**: IsRealTime, DataTime, DataTimeStr, BrokerList []*BrokerItem

**BrokerItem**: NetVol (int64), BrokerName (string), BuySellType (enum), AvgPrice (float64), TotalVol (float64), TotalTurnover (float64)

```go
type GetTopTenBuySellBrokersRequest struct {
    Security   *qotcommon.Security
    DaysBefore int32
}

type TopTenBrokerItem struct {
    NetVol        int64
    BrokerName    string
    BuySellType   int32
    AvgPrice      float64
    TotalVol      float64
    TotalTurnover float64
}

type GetTopTenBuySellBrokersResponse struct {
    IsRealTime  bool
    DataTime    int64
    DataTimeStr string
    BrokerList  []*TopTenBrokerItem
}
```

Note: This API has relatively simple nested types, so we create custom `TopTenBrokerItem` instead of pass-through.

#### GetDailyShortVolume (ProtoID 3248)

**PB Package**: `qotgetdailyshortvolume`

**C2S**: Security (required), NextKey, Num (optional)
**S2C**: UsItemList, HkItemList, NextKey, AggregatedShort, AggregatedShortRatio, NewTimeStr

```go
type GetDailyShortVolumeRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetDailyShortVolumeResponse struct {
    UsItemList            []*qotgetdailyshortvolume.UsDailyShortVolumeItem
    HkItemList            []*qotgetdailyshortvolume.HkDailyShortVolumeItem
    NextKey               string
    AggregatedShort       int64
    AggregatedShortRatio  float64
    NewTimeStr            string
}
```

#### GetShortInterest (ProtoID 3249)

**PB Package**: `qotgetshortinterest`

**C2S**: Security (required), NextKey, Num (optional)
**S2C**: UsItemList, HkItemList, NextKey

```go
type GetShortInterestRequest struct {
    Security *qotcommon.Security
    NextKey  string
    Num      int32
}

type GetShortInterestResponse struct {
    UsItemList []*qotgetshortinterest.UsShortInterestItem
    HkItemList []*qotgetshortinterest.HkShortInterestItem
    NextKey    string
}
```

---

### 5.9 option_extra.go

#### GetOptionVolatility (ProtoID 3250)

**PB Package**: `qotgetoptionvolatility`

**C2S**: Security (required), QueryTimePeriod (optional enum), HvTimePeriod (optional int32)
**S2C**: ItemList, AverageImpvol, ImpvolStatus, Analysis

```go
type GetOptionVolatilityRequest struct {
    Security         *qotcommon.Security
    QueryTimePeriod  int32
    HvTimePeriod     int32
}

type GetOptionVolatilityResponse struct {
    ItemList      []*qotgetoptionvolatility.VolatilityItem
    AverageImpvol float64
    ImpvolStatus  int32
    Analysis      string
}
```

#### GetOptionExerciseProbability (ProtoID 3251)

**PB Package**: `qotgetoptionexerciseprobability`

**C2S**: Security (required)
**S2C**: ItemList []*StrikeProbabilityItem

```go
type GetOptionExerciseProbabilityRequest struct {
    Security *qotcommon.Security
}

type GetOptionExerciseProbabilityResponse struct {
    ItemList []*qotgetoptionexerciseprobability.StrikeProbabilityItem
}
```

---

## 6. Client Wrapper Specifications

### 6.1 client/quote_api.go — 26 New Functions

Each follows the existing pattern: `func Xxx(ctx context.Context, c *Client, ...) (*XxxResponse, error)` with input validation.

| Function | Delegates To | Key Validation |
|----------|-------------|----------------|
| `GetFinancialsStatements` | `qot.GetFinancialsStatements` | Security nil |
| `GetFinancialsRevenueBreakdown` | `qot.GetFinancialsRevenueBreakdown` | Security nil |
| `GetResearchAnalystConsensus` | `qot.GetResearchAnalystConsensus` | Security nil |
| `GetResearchRatingSummary` | `qot.GetResearchRatingSummary` | Security nil |
| `GetResearchMorningstarReport` | `qot.GetResearchMorningstarReport` | Security nil |
| `GetValuationDetail` | `qot.GetValuationDetail` | Security nil |
| `GetValuationPlateStockList` | `qot.GetValuationPlateStockList` | Security nil |
| `GetCorporateActionsDividends` | `qot.GetCorporateActionsDividends` | Security nil |
| `GetCorporateActionsBuybacks` | `qot.GetCorporateActionsBuybacks` | Security nil |
| `GetCorporateActionsStockSplits` | `qot.GetCorporateActionsStockSplits` | Security nil |
| `GetShareholdersOverview` | `qot.GetShareholdersOverview` | Security nil |
| `GetShareholdersHoldingChanges` | `qot.GetShareholdersHoldingChanges` | Security nil |
| `GetShareholdersHolderDetail` | `qot.GetShareholdersHolderDetail` | Security nil |
| `GetShareholdersInstitutional` | `qot.GetShareholdersInstitutional` | Security nil |
| `GetInsiderHolderList` | `qot.GetInsiderHolderList` | Security nil |
| `GetInsiderTradeList` | `qot.GetInsiderTradeList` | Security nil |
| `GetCompanyProfile` | `qot.GetCompanyProfile` | Security nil |
| `GetCompanyExecutives` | `qot.GetCompanyExecutives` | Security nil |
| `GetCompanyExecutiveBackground` | `qot.GetCompanyExecutiveBackground` | Security nil |
| `GetCompanyOperationalEfficiency` | `qot.GetCompanyOperationalEfficiency` | Security nil |
| `GetTopTenBuySellBrokers` | `qot.GetTopTenBuySellBrokers` | Security nil |
| `GetDailyShortVolume` | `qot.GetDailyShortVolume` | Security nil |
| `GetShortInterest` | `qot.GetShortInterest` | Security nil |
| `GetOptionVolatility` | `qot.GetOptionVolatility` | Security nil |
| `GetOptionExerciseProbability` | `qot.GetOptionExerciseProbability` | Security nil |
| `RequestHistoryKLQuota` | `qot.RequestHistoryKLQuota` | (existing, just missing wrapper) |

### 6.2 client/fluent_api.go — 25 New QuoteAPI Methods

One-line delegation methods on `QuoteAPI` struct, matching the 25 new APIs above.

### 6.3 client/trade_api.go — 8 Convenience Wrappers

| Function | Delegates To | Key Validation |
|----------|-------------|----------------|
| `QuickBuy` | `trd.QuickBuy` | accID!=0, code!="", qty>0, price>0 |
| `QuickSell` | `trd.QuickSell` | accID!=0, code!="", qty>0, price>0 |
| `QuickMarketBuy` | `trd.QuickMarketBuy` | accID!=0, code!="", qty>0 |
| `QuickMarketSell` | `trd.QuickMarketSell` | accID!=0, code!="", qty>0 |
| `GetPositions` | `trd.GetPositions` | accID!=0 |
| `GetTodayFills` | `trd.GetTodayFills` | accID!=0 |
| `GetTodayOrders` | `trd.GetTodayOrders` | accID!=0 |
| `GetAccountFunds` | `trd.GetAccountFunds` | accID!=0 |

---

## 7. Execution Steps

### Step 1: Create 9 new pkg/qot/ implementation files
Each file contains: custom request/response types + implementation function.

### Step 2: Add client/quote_api.go wrappers (25 new + 1 missing existing)
Simple delegation wrappers with input validation.

### Step 3: Add client/fluent_api.go QuoteAPI methods (25 new)
One-line delegation methods.

### Step 4: Add client/trade_api.go convenience wrappers (8 functions)
Delegation wrappers for pkg/trd/convenience.go.

### Step 5: Build & Test
```
go build ./...
go vet ./...
go test -race ./pkg/qot/... ./client/...
```

### Step 6: Update docs
- CHANGELOG.md: Add Phase 4 items under [Unreleased]
- IMPLEMENTATION_PLAN.md: Mark Phase 4 as DONE
- PHASE4_API_COVERAGE_PLAN.md: Update status

### Step 7: Commit & Push
```
git add -A && git commit -m "feat: Phase 4 — add 25 new quote APIs + client wrappers"
git push origin main
git push gitee main
```

---

## 8. Design Decisions

1. **Pass-through proto types for complex nested data**: FinancialFieldInfo, FinancialReport, CompanyLabItem, DirectorInfo, etc. are deeply nested proto types. Wrapping each sub-field into custom Go types would add hundreds of lines with no real benefit. Users access fields with `util.ProtoXxx()` helpers.

2. **Full S2C pass-through for Morningstar & Valuation**: GetResearchMorningstarReport (25+ fields) and GetValuationDetail (complex nested) use full S2C pass-through. The S2C struct IS the response.

3. **Custom wrapper for TopTenBuySellBrokers**: This API has relatively simple nested types (BrokerItem with 6 scalar fields), so we create a custom `TopTenBrokerItem` struct for a cleaner API.

4. **Request struct pattern for all APIs**: Even Security-only APIs use `*XxxRequest` struct for consistency and forward compatibility.

5. **Separate files by domain**: 9 new domain-specific files keep files manageable and logically organized.

6. **No unit tests for new APIs yet**: These APIs require a live OpenD connection. Build/vet checks catch structural errors.

7. **Screening APIs deferred**: StockScreen, WarrantScreen, OptionScreen, and 2 EarningsPrice APIs have no ProtoID assigned yet.

---

## 9. Risk Assessment

- **Impact**: LOW — All changes are additive (new functions). No existing code is modified.
- **Breaking changes**: NONE — New public functions don't affect existing APIs.
- **Proto compatibility**: All pb packages already generated and building successfully.
- **ProtoID coverage**: All 25 ProtoIDs already defined in `pkg/constant/constant.go`.

---

## 10. Verification Checklist

- [ ] All 9 new pkg/qot/ files created and building
- [ ] All 25 new pkg/qot/ functions follow the implementation pattern
- [ ] All 26 client/quote_api.go wrappers added with validation
- [ ] All 25 client/fluent_api.go QuoteAPI methods added
- [ ] All 8 client/trade_api.go convenience wrappers added
- [ ] `go build ./...` passes
- [ ] `go vet ./...` passes
- [ ] `go test -race ./pkg/qot/... ./client/...` passes
- [ ] CHANGELOG.md updated
- [ ] IMPLEMENTATION_PLAN.md updated
- [ ] Changes committed and pushed to origin/main + gitee/main
