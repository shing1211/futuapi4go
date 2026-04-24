# Implementation Plan: futuapi4go Context Migration & Enhancement

## Overview
Upgrade all API functions to support `context.Context` for cancellation, timeout, and deadline propagation. Add input validation, nil safety on list iterations, and standardized error handling.

## Constraints
- Context must be first parameter for all API functions
- Maintain backward compatibility where possible (not always possible due to signature changes)
- Use `RequestContext` pattern to eliminate 20-line boilerplate per function
- Add nil guards on all list iteration loops

## Progress

### Completed Functions

#### pkg/qot/quote.go (6/30)
- [x] `GetBasicQot` - context, nil guard, wrapError
- [x] `GetKL` - context, nil guard, wrapError
- [x] `GetOrderBook` - context, nil guard, wrapError
- [x] `GetTicker` - context, nil guard, wrapError
- [x] `GetRT` - context, nil guard, wrapError
- [x] `GetBroker` - context, nil guard, wrapError

#### pkg/trd/trade.go (6/18)
- [x] `GetAccList` - context, nil guard, wrapError
- [x] `GetFunds` - context, nil guard, wrapError
- [x] `GetPositionList` - context, nil guard, wrapError
- [x] `PlaceOrder` - context, nil guard, wrapError, validation
- [x] `GetOrderList` - context, nil guard, wrapError
- [x] `ModifyOrder` - context, nil guard, wrapError, validation

#### client/client.go (wrappers)
- [x] All corresponding wrapper functions updated with context

### Remaining Functions

#### pkg/qot/quote.go (24 remaining)
```
GetStaticInfo
GetPlateSet
GetPlateSecurity
RequestTradeDate
RequestHistoryKL
GetSecuritySnapshot
Subscribe
GetCapitalFlow
GetCapitalDistribution
GetUserSecurity
GetPriceReminder
GetOptionExpirationDate
GetOptionChain
StockFilter
GetWarrant
GetSuspend
GetFutureInfo
GetCodeChange
GetIpoList
GetHoldingChangeList
GetUserSecurityGroup
ModifyUserSecurity
SetPriceReminder
RegQotPush
RequestRehab
RequestHistoryKLQuota
GetRehab
```

#### pkg/trd/trade.go (12 remaining)
```
UnlockTrade
GetOrderDetail
GetTrdFeecalc
GetUserInfo
GetFuturesInfo
ChangePassword
SetEntrustWave
SetOrderPrice
SetOrderDate
GetCreditInterestInfo
GetMaxBuySize
GetOrderWid
```

#### pkg/sys/system.go
```
Notify
KeepAlive
GetGlobalState
```

#### pkg/push/ (push handlers)
```
All push handler functions
```

### Test Files (blocked until library complete)
```
test/trd_api/trd_test.go
test/qot_api/qot_test.go
test/integration/integration_hsi_test.go
test/benchmark/benchmark_test.go
```

## Pattern for Updates

### Before
```go
func GetBroker(c *futuapi.Client, req *GetBrokerRequest) (*GetBrokerResponse, error) {
    if err := c.EnsureConnected(); err != nil {
        return nil, err
    }
    c2s := &qotgetbroker.C2S{...}

    pkt := &qotgetbroker.Request{C2S: c2s}

    body, err := proto.Marshal(pkt)
    if err != nil {
        return nil, err
    }

    serialNo := c.NextSerialNo()
    if err := c.Conn().WritePacket(ProtoID_GetBroker, serialNo, body); err != nil {
        return nil, err
    }

    apiTimeout := c.Conn().APITimeout()
    if apiTimeout == 0 {
        apiTimeout = 30 * time.Second
    }
    pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
    if err != nil {
        return nil, err
    }

    var rsp qotgetbroker.Response
    if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
        return nil, err
    }

    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, fmt.Errorf("GetBroker failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
    }

    s2c := rsp.GetS2C()
    if s2c == nil {
        return nil, fmt.Errorf("GetBroker: s2c is nil")
    }

    result := &GetBrokerResponse{...}
    for _, b := range s2c.GetBrokerAskList() {
        result.AskBrokerList = append(result.AskBrokerList, &Broker{...})
    }
    for _, b := range s2c.GetBrokerBidList() {
        result.BidBrokerList = append(result.BidBrokerList, &Broker{...})
    }
    return result, nil
}
```

### After
```go
func GetBroker(ctx context.Context, c *futuapi.Client, req *GetBrokerRequest) (*GetBrokerResponse, error) {
    // Input validation
    if req.Security == nil {
        return nil, fmt.Errorf("security is required")
    }

    c2s := &qotgetbroker.C2S{...}

    pkt := &qotgetbroker.Request{C2S: c2s}
    var rsp qotgetbroker.Response

    if err := c.RequestContext(ctx, ProtoID_GetBroker, pkt, &rsp); err != nil {
        return nil, err
    }

    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, wrapError("GetBroker", rsp.GetRetType(), rsp.GetRetMsg())
    }

    s2c := rsp.GetS2C()
    if s2c == nil {
        return nil, fmt.Errorf("GetBroker: s2c is nil")
    }

    result := &GetBrokerResponse{...}
    for _, b := range s2c.GetBrokerAskList() {
        if b == nil {
            continue
        }
        result.AskBrokerList = append(result.AskBrokerList, &Broker{...})
    }
    for _, b := range s2c.GetBrokerBidList() {
        if b == nil {
            continue
        }
        result.BidBrokerList = append(result.BidBrokerList, &Broker{...})
    }
    return result, nil
}
```

## Helper Functions Needed

### wrapError (already added to qot and trd packages)
```go
func wrapError(apiName string, retType int32, retMsg string) error {
    return fmt.Errorf("%s failed: retType=%d, retMsg=%s", apiName, retType, retMsg)
}
```

## Verification Commands
```bash
go build ./pkg/... ./client/...  # Library must always pass
go test ./pkg/... ./client/...  # Unit tests
go vet ./...                    # Lint
```

## Estimated Scope
- ~40 functions in pkg/ packages
- ~20 wrapper functions in client/
- ~4 test files with multiple test functions each

## Timeline
- Library functions: ~10-15 per session
- Test file updates: after library complete