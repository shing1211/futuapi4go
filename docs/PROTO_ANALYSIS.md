# Protobuf File Usage Analysis

## Summary

**Date**: 2026-04-11  
**Total Proto Files**: 74 → **73** (after cleanup)  
**Used Packages**: 73 (100%)  
**Unused Packages**: 1 (removed)  
**Duplicate Packages**: 0

---

## Removed Files

### UsedQuota.proto ❌ REMOVED

**Reason**: Completely unused throughout the codebase

**What it was**:
- Standalone proto message returning `usedSubQuota` and `usedKLineQuota`
- Had its own protocol ID but was never wired up
- No `ProtoID_UsedQuota` constant defined
- No simulator handler implemented
- No API wrapper function

**Why it's redundant**:
- Quota functionality is already provided by:
  - `Qot_GetSubInfo` - Returns `totalUsedQuota` and `remainQuota` per connection
  - `Qot_Common.ConnSubInfo` - Has `usedQuota` field
  - `Notify.UsedQuota` - Push notification for quota updates
  - `Qot_RequestHistoryKLQuota` - Returns historical K-line quota usage

**Files Removed**:
- `api/proto/UsedQuota.proto`
- `pkg/pb/usedquota/` (entire directory)

---

## All 73 Used Proto Files

### System & Common (9 files)

| Proto File | Package | Used By | Purpose |
|------------|---------|---------|---------|
| Common.proto | common | Everywhere | Shared types (RetType, enums) |
| InitConnect.proto | initconnect | internal/client | Connection handshake |
| KeepAlive.proto | keepalive | internal/client | Heartbeat |
| Notify.proto | notify | pkg/push, simulator | System notifications |
| GetGlobalState.proto | getglobalstate | pkg/sys | Global state query |
| GetUserInfo.proto | getuserinfo | pkg/sys | User info query |
| GetDelayStatistics.proto | getdelaystatistics | pkg/sys | Latency stats |
| Verification.proto | verification | pkg/sys | Verification code |
| Qot_Common.proto | qotcommon | pkg/qot, pkg/push | Shared quote types |
| Trd_Common.proto | trdcommon | pkg/trd, pkg/push | Shared trade types |

### Quote Request/Response (27 files)

| Proto File | Package | API Function | Status |
|------------|---------|--------------|--------|
| Qot_GetBasicQot.proto | qotgetbasicqot | GetBasicQot | ✅ Used |
| Qot_GetBroker.proto | qotgetbroker | GetBroker | ✅ Used |
| Qot_GetCapitalDistribution.proto | qotgetcapitaldistribution | GetCapitalDistribution | ✅ Used |
| Qot_GetCapitalFlow.proto | qotgetcapitalflow | GetCapitalFlow | ✅ Used |
| Qot_GetCodeChange.proto | qotgetcodechange | GetCodeChange | ✅ Used |
| Qot_GetFutureInfo.proto | qotgetfutureinfo | GetFutureInfo | ✅ Used |
| Qot_GetHoldingChangeList.proto | qotgetholdingchangelist | GetHoldingChangeList | ✅ Used |
| Qot_GetIpoList.proto | qotgetipolist | GetIpoList | ✅ Used |
| Qot_GetKL.proto | qotgetkl | GetKL | ✅ Used |
| Qot_GetMarketState.proto | qotgetmarketstate | GetMarketState | ✅ Used |
| Qot_GetOptionChain.proto | qotgetoptionchain | GetOptionChain | ✅ Used |
| Qot_GetOptionExpirationDate.proto | qotgetoptionexpirationdate | GetOptionExpirationDate | ✅ Used |
| Qot_GetOrderBook.proto | qotgetorderbook | GetOrderBook | ✅ Used |
| Qot_GetOwnerPlate.proto | qotgetownerplate | GetOwnerPlate | ✅ Used |
| Qot_GetPlateSecurity.proto | qotgetplatesecurity | GetPlateSecurity | ✅ Used |
| Qot_GetPlateSet.proto | qotgetplateset | GetPlateSet | ✅ Used |
| Qot_GetPriceReminder.proto | qotgetpricereminder | GetPriceReminder | ✅ Used |
| Qot_GetReference.proto | qotgetreference | GetReference | ✅ Used |
| Qot_GetRT.proto | qotgetrt | GetRT | ✅ Used |
| Qot_GetSecuritySnapshot.proto | qotgetsecuritysnapshot | GetSecuritySnapshot | ✅ Used |
| Qot_GetStaticInfo.proto | qotgetstaticinfo | GetStaticInfo | ✅ Used |
| Qot_GetSubInfo.proto | qotgetsubinfo | GetSubInfo | ✅ Used |
| Qot_GetSuspend.proto | qotgetsuspend | GetSuspend | ✅ Used |
| Qot_GetTicker.proto | qotgetticker | GetTicker | ✅ Used |
| Qot_GetTradeDate.proto | qotgettradedate | GetTradeDate | ✅ Used |
| Qot_GetUserSecurity.proto | qotgetusersecurity | GetUserSecurity | ✅ Used |
| Qot_GetUserSecurityGroup.proto | qotgetusersecuritygroup | GetUserSecurityGroup | ✅ Used |
| Qot_GetWarrant.proto | qotgetwarrant | GetWarrant | ✅ Used |

### Quote Configuration (8 files)

| Proto File | Package | API Function | Status |
|------------|---------|--------------|--------|
| Qot_ModifyUserSecurity.proto | qotmodifyusersecurity | ModifyUserSecurity | ✅ Used |
| Qot_RegQotPush.proto | qotregqotpush | RegQotPush | ✅ Used |
| Qot_RequestHistoryKL.proto | qotrequesthistorykl | RequestHistoryKL | ✅ Used |
| Qot_RequestHistoryKLQuota.proto | qotrequesthistoryklquota | RequestHistoryKLQuota | ✅ Used |
| Qot_RequestRehab.proto | qotrequestrehab | RequestRehab | ✅ Used |
| Qot_RequestTradeDate.proto | qotrequesttradedate | RequestTradeDate | ✅ Used |
| Qot_SetPriceReminder.proto | qotsetpricereminder | SetPriceReminder | ✅ Used |
| Qot_StockFilter.proto | qotstockfilter | StockFilter | ✅ Used |
| Qot_Sub.proto | qotsub | Subscribe | ✅ Used |

### Quote Push/Update (7 files)

| Proto File | Package | Push Handler | Status |
|------------|---------|--------------|--------|
| Qot_UpdateBasicQot.proto | qotupdatebasicqot | ParseUpdateBasicQot | ✅ Used |
| Qot_UpdateBroker.proto | qotupdatebroker | ParseUpdateBroker | ✅ Used |
| Qot_UpdateKL.proto | qotupdatekl | ParseUpdateKL | ✅ Used |
| Qot_UpdateOrderBook.proto | qotupdateorderbook | ParseUpdateOrderBook | ✅ Used |
| Qot_UpdatePriceReminder.proto | qotupdatepricereminder | ParseUpdatePriceReminder | ✅ Used |
| Qot_UpdateRT.proto | qotupdatert | ParseUpdateRT | ✅ Used |
| Qot_UpdateTicker.proto | qotupdateticker | ParseUpdateTicker | ✅ Used |

### Trading Request/Response (13 files)

| Proto File | Package | API Function | Status |
|------------|---------|--------------|--------|
| Trd_GetAccList.proto | trdgetacclist | GetAccList | ✅ Used |
| Trd_GetFunds.proto | trdgetfunds | GetFunds | ✅ Used |
| Trd_GetHistoryOrderFillList.proto | trdgethistoryorderfilllist | GetHistoryOrderFillList | ✅ Used |
| Trd_GetHistoryOrderList.proto | trdgethistoryorderlist | GetHistoryOrderList | ✅ Used |
| Trd_GetMarginRatio.proto | trdgetmarginratio | GetMarginRatio | ✅ Used |
| Trd_GetMaxTrdQtys.proto | trdgetmaxtrdqtys | GetMaxTrdQtys | ✅ Used |
| Trd_GetOrderFee.proto | trdgetorderfee | GetOrderFee | ✅ Used |
| Trd_GetOrderFillList.proto | trdgetorderfilllist | GetOrderFillList | ✅ Used |
| Trd_GetOrderList.proto | trdgetorderlist | GetOrderList | ✅ Used |
| Trd_GetPositionList.proto | trdgetpositionlist | GetPositionList | ✅ Used |
| Trd_ModifyOrder.proto | trdmodifyorder | ModifyOrder | ✅ Used |
| Trd_PlaceOrder.proto | trdplaceorder | PlaceOrder | ✅ Used |
| Trd_ReconfirmOrder.proto | trdreconfirmorder | ReconfirmOrder | ✅ Used |
| Trd_UnlockTrade.proto | trdunlocktrade | UnlockTrade | ✅ Used |
| Trd_FlowSummary.proto | trdflowsummary | GetFlowSummary | ✅ Used |
| Trd_SubAccPush.proto | trdsubaccpush | SubAccPush | ✅ Used |

### Trading Push/Update (3 files)

| Proto File | Package | Push Handler | Status |
|------------|---------|--------------|--------|
| Trd_Notify.proto | trdnotify | ParseTrdNotify | ✅ Used |
| Trd_UpdateOrder.proto | trdupdateorder | ParseUpdateOrder | ✅ Used |
| Trd_UpdateOrderFill.proto | trdupdateorderfill | ParseUpdateOrderFill | ✅ Used |

---

## Analysis of Similar-Function Protos

### GetTradeDate vs RequestTradeDate - NOT DUPLICATES

| Feature | GetTradeDate | RequestTradeDate |
|---------|--------------|------------------|
| **Market field** | Required | Required |
| **Security field** | No | Yes (optional) |
| **Priority** | Market-level | Security overrides market |
| **Use case** | General trading calendar | Per-security calendar |
| **Both used?** | ✅ Yes | ✅ Yes |

**Conclusion**: Different signatures for different use cases. Both actively used.

### GetSubInfo vs UsedQuota - NOT DUPLICATES (but UsedQuota unused)

| Feature | GetSubInfo | UsedQuota |
|---------|------------|-----------|
| **Returns** | Per-connection quota info | Global quota usage |
| **Includes** | totalUsedQuota, remainQuota | usedSubQuota, usedKLineQuota |
| **Used in code?** | ✅ Yes | ❌ No |
| **Has ProtoID?** | ✅ Yes | ❌ No |
| **Has API wrapper?** | ✅ Yes (GetSubInfo) | ❌ No |
| **Has simulator handler?** | ✅ Yes | ❌ No |

**Conclusion**: GetSubInfo is the active quota check. UsedQuota was defined but never integrated.

---

## Proto File Organization

### Directory Structure

```
api/proto/ (73 files)
├── Common.proto                    # Shared types
├── *_Connect_*.proto              # Connection management (2)
├── Qot_*.proto                    # Quote protocols (44)
├── Trd_*.proto                    # Trade protocols (18)
├── *_Statistics.proto             # Statistics (1)
├── *_User*.proto                  # User management (2)
└── *_Verification.proto           # Verification (1)

pkg/pb/ (73 packages)
├── common/                         # Shared types
├── qotcommon/                     # Quote shared types
├── trdcommon/                     # Trade shared types
├── qot*/                          # Quote request/response (27)
├── qotupdate*/                    # Quote push (7)
├── trd*/                          # Trade request/response (16)
└── trdupdate*/                    # Trade push (3)
```

### Naming Convention

**Proto files**: `Qot_<Action><Resource>.proto`
- Examples: `Qot_GetBasicQot.proto`, `Qot_UpdateKL.proto`

**Go packages**: Lowercase, abbreviated
- Examples: `qotgetbasicqot`, `qotupdatekl`

---

## Verification

### Build Test

```bash
go build ./...
# Result: ✓ Successful (zero errors)
```

### Test Compilation

```bash
go test -c ./test/qot_api
go test -c ./test/trd_api
go test -c ./test/integration
# Result: ✓ All compile successfully
```

### Import Check

```bash
# Search for any remaining references to usedquota
grep -r "usedquota" . --include="*.go"
# Result: No matches found (except in this documentation)
```

---

## Future Maintenance

### Adding New Proto Files

1. Create `.proto` file in `api/proto/`
2. Generate Go code: `protoc --go_out=pkg/pb api/proto/*.proto`
3. Add API wrapper in `pkg/qot/` or `pkg/trd/`
4. Add simulator handler in `cmd/simulator/`
5. Add tests in `test/`
6. Update this document

### Identifying Unused Protos

Run this check periodically:

```bash
# List all pb packages
ls pkg/pb/ | while read pkg; do
    # Check if imported anywhere
    count=$(grep -r "github.com/shing1211/futuapi4go/pkg/pb/$pkg" --include="*.go" | grep -v "pkg/pb/$pkg/" | wc -l)
    if [ $count -eq 0 ]; then
        echo "UNUSED: $pkg"
    fi
done
```

---

## Summary Statistics

| Metric | Count |
|--------|-------|
| **Total Proto Files (before)** | 74 |
| **Total Proto Files (after)** | 73 |
| **Removed (unused)** | 1 |
| **Actively used** | 73 (100%) |
| **Quote protocols** | 44 |
| **Trading protocols** | 18 |
| **System protocols** | 9 |
| **Push/update protocols** | 10 |
| **Duplicate protocols** | 0 |

---

**Status**: ✅ All proto files are now actively used. No duplicates found. Clean codebase.
