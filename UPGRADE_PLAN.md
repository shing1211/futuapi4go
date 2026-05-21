# Upgrade Plan: futuapi4go v10.5.6508 → v10.6.6608

**Date:** 2026-05-21
**Goal:** Upgrade SDK to support Futu OpenD v10.6.6608

---

## Status Quo

| Item | Current (v10.5) | Target (v10.6) |
|------|-----------------|-----------------|
| Proto files | 79 in `api/proto/` | 104 in `api/proto/` (replace all) |
| `.pb.go` dirs | 79 in `pkg/pb/` | Regenerate all |
| `clientVer` | `1005` (`internal/client/client.go` lines 555, 698) | `1066` |
| Proto IDs defined | Up to 3251 (Qot_GetTradeDate) | Up to 3251 (same) |
| SDK version | v0.10.0 | v0.11.0 |

---

## What Changed in v10.6.6608

### 1. Enhanced Qot_StockFilter (Protocol ID 3215)
- **New filter types**: `PatternFilter` (K-line shape matching) and `CustomIndicatorFilter` (MA, EMA, RSI, MACD, BOLL, KDJ with relative position)
- **New enums in Qot_Common**: `CustomIndicatorField`, `PatternField`, `SortDir`, `RelativePosition`, `FinancialQuarter`, `FinancialStatementsType`, `F10Type`
- **Enhanced response**: `StockData` gains `repeated CustomIndicatorData`

### 2. New Fundamental Data APIs (25 new market data APIs, Protocol IDs 3227-3251)

| API | Proto ID | Description |
|-----|----------|-------------|
| GetFinancialsStatements | 3227 | Financial statements (income/balance/cash flow) with pagination |
| GetFinancialsRevenueBreakdown | 3228 | Revenue breakdown by segment |
| GetResearchAnalystConsensus | 3229 | Analyst consensus ratings |
| GetResearchRatingSummary | 3230 | Buy/hold/sell breakdown |
| GetResearchMorningstarReport | 3231 | Morningstar report |
| GetValuationDetail | 3232 | Valuation detail |
| GetValuationPlateStockList | 3233 | Valuation plate stock list |
| GetCorporateActionsDividends | 3234 | Dividend history |
| GetCorporateActionsBuybacks | 3235 | Share buyback history |
| GetCorporateActionsStockSplits | 3236 | Stock split history |
| GetShareholdersOverview | 3237 | Top shareholders overview |
| GetShareholdersHoldingChanges | 3238 | Shareholder holding changes |
| GetShareholdersHolderDetail | 3239 | Detailed holder info |
| GetShareholdersInstitutional | 3240 | Institutional holdings |
| GetInsiderHolderList | 3241 | Insider holder list |
| GetInsiderTradeList | 3242 | Insider trade list |
| GetCompanyProfile | 3243 | Company profile |
| GetCompanyExecutives | 3244 | Company executives |
| GetCompanyExecutiveBackground | 3245 | Executive background |
| GetCompanyOperationalEfficiency | 3246 | Operational efficiency |
| GetTopTenBuySellBrokers | 3247 | Top 10 brokers |
| GetDailyShortVolume | 3248 | Daily short volume |
| GetShortInterest | 3249 | Short interest |
| GetOptionVolatility | 3250 | Option volatility |
| GetOptionExerciseProbability | 3251 | Option exercise probability |

### 3. Trading Additions
- `TrdMarket.AU` (Australia) paper trading support
- `TrdSubAccType` may have new subtypes

---

## Execution Phases

### Phase 1 — Infrastructure: Create Proto Regeneration Script
**File to create:** `scripts/regen-all-protos.sh`

The old `scripts/` directory was deleted (commit f4d1ff3) because it contained Windows-only PowerShell scripts. Create a Linux shell replacement.

**Status**: ✓ Complete

---

### Phase 2 — Download and Replace Proto Files (MANUAL STEP)
**Action required**: Download v10.6.6608 Protobuf from https://www.futunn.com/hk/download/OpenAPI and replace all files in `api/proto/`.

1. Go to https://www.futunn.com/hk/download/OpenAPI
2. Download **Protobuf** package under **Ver.10.6.6608** section
3. Extract and copy all `.proto` files into `api/proto/` (overwrite existing)
4. Verify: `ls api/proto/*.proto | wc -l` should return 79+

**Status**: ✓ Complete (manual)

---

### Phase 3 — Regenerate Go PB Files
**Command**: `./scripts/regen-all-protos.sh`
**Prerequisites**: `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc` installed

```bash
# Install tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
apt install protobuf-compiler  # or brew install protobuf
```

**Status**: ✓ Complete

---

### Phase 4 — Update SDK Constants, Version, and Types
**Files to modify**:
- `internal/client/client.go` — clientVer 1005 → 1066 (2 locations)
- `pkg/constant/constant.go` — add ~25 new ProtoID constants (3227-3251)
- `client/types.go` — add new SDK types for new API responses

**Status**: ✓ Complete

---

### Phase 5 — Verify Build
```bash
go build ./...
go vet ./...
go test -race ./...
```

**Status**: ✓ Complete

---

### Phase 6 — Update Documentation
**Files to modify**:
- `README.md` — badge v10.5.6508 → v10.6.6608
- `AGENTS.md` — upgrade procedure section updated
- `CHANGELOG.md` — unreleased entry added

**Status**: ✓ Complete

---

## Files Modified Per Phase

| Phase | Files |
|-------|-------|
| 1 | `scripts/regen-all-protos.sh` (new) |
| 2 | `api/proto/*.proto` (replace all) |
| 3 | `pkg/pb/*/*.pb.go` (regenerate all) |
| 4 | `internal/client/client.go`, `pkg/constant/constant.go`, `client/types.go` |
| 5 | (none — verification only) |
| 6 | `README.md`, `AGENTS.md`, `CHANGELOG.md` |

---

## Release

- **SDK version**: v0.11.0
- **Tag**: `v0.11.0`
- **Push to**: origin + gitee (dual push)
