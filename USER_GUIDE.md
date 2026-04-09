# futuapi4go User Guide

This guide is for quantitative traders, covering how to use the futuapi4go SDK for market data queries and trading operations.

## Table of Contents

1. [Quick Start](#quick-start)
2. [Connection Management](#connection-management)
3. [Market Data Queries](#market-data-queries)
4. [Trading Operations](#trading-operations)
5. [Real-time Push](#real-time-push)
6. [FAQ](#faq)

---

## Quick Start

### Installation

```bash
go get github.com/shing1211/futuapi4go
```

### Basic Usage Flow

```go
package main

import (
    "fmt"
    "log"
    futuapi "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/qot"
    "github.com/shing1211/futuapi4go/pb/qotcommon"
)

func main() {
    // 1. Create client
    cli := futuapi.New()
    
    // 2. Connect to OpenD
    err := cli.Connect("127.0.0.1:11111")
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()
    
    // 3. Call API
    market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
    code := "00700"
    securities := []*qotcommon.Security{
        {Market: &market, Code: &code},
    }
    
    result, err := qot.GetBasicQot(cli, securities)
    if err != nil {
        log.Fatal(err)
    }
    
    // 4. Process results
    for _, bq := range result {
        fmt.Printf("%s %s: CurPrice=%.2f\n", 
            bq.Security.GetCode(), bq.Name, bq.CurPrice)
    }
}
```

---

## Connection Management

### Create Connection

```go
cli := futuapi.New()
err := cli.Connect("127.0.0.1:11111")
if err != nil {
    log.Fatal(err)
}
defer cli.Close()
```

### Initialize Connection (Get Connection ID)

```go
// Initialize connection, get user info
userInfo, err := sys.InitConnect(cli, "your_app_id", "your_hash")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("User: %s, ConnID: %d\n", userInfo.GetNickName(), userInfo.GetConnectionID())
```

### Heartbeat Keep-alive

SDK automatically sends heartbeat packets to maintain connection, no manual operation needed.

---

## Market Data Queries

### Get Real-time Quotes (GetBasicQot)

```go
market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
code := "00700"
securities := []*qotcommon.Security{
    {Market: &market, Code: &code},
}

result, err := qot.GetBasicQot(cli, securities)
if err != nil {
    log.Fatal(err)
}

for _, bq := range result {
    fmt.Printf("%s: CurPrice=%.2f, Change=%.2f%%\n",
        bq.Name, bq.CurPrice, bq.ChangeRate)
}
```

### Get K-Line (GetKL)

```go
req := &qot.GetKLRequest{
    Security:  &qotcommon.Security{Market: &market, Code: &code},
    RehabType: int32(qotcommon.RehabType_RehabType_None),
    KLType:    int32(qotcommon.KLType_KLType_Day),
    ReqNum:    100,
}

result, err := qot.GetKL(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, kl := range result.KLList {
    fmt.Printf("%s: Open=%.2f, High=%.2f, Low=%.2f, Close=%.2f\n",
        kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice, kl.ClosePrice)
}
```

### Get Order Book (GetOrderBook)

```go
req := &qot.GetOrderBookRequest{
    Security: &qotcommon.Security{Market: &market, Code: &code},
    Num:      10,
}

result, err := qot.GetOrderBook(cli, req)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Bid:")
for _, bid := range result.OrderBookBidList {
    fmt.Printf("  Price=%.2f, Volume=%d\n", bid.Price, bid.Volume)
}

fmt.Println("Ask:")
for _, ask := range result.OrderBookAskList {
    fmt.Printf("  Price=%.2f, Volume=%d\n", ask.Price, ask.Volume)
}
```

### Get Minute Data (GetRT)

```go
req := &qot.GetRTRequest{
    Security: &qotcommon.Security{Market: &market, Code: &code},
}

result, err := qot.GetRT(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, rt := range result.RTList {
    fmt.Printf("%s: Price=%.2f, Volume=%d\n",
        rt.Time, rt.Price, rt.Volume)
}
```

### Get Capital Flow (GetCapitalFlow)

```go
req := &qot.GetCapitalFlowRequest{
    Security:   &qotcommon.Security{Market: &market, Code: &code},
    PeriodType: 1, // Daily
}

result, err := qot.GetCapitalFlow(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, f := range result.FlowItemList {
    fmt.Printf("%s: MainInFlow=%.2f\n", f.Time, f.MainInFlow)
}
```

### Stock Filter (StockFilter)

```go
req := &qot.StockFilterRequest{
    Begin:  0,
    Num:    10,
    Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
    BaseFilterList: []*qotstockfilter.BaseFilter{
        {
            FieldName:  int32(qotstockfilter.StockField_StockField_CurPrice),
            FilterMin:  proto.Float64(10.0),
            FilterMax:  proto.Float64(100.0),
            IsNoFilter: proto.Bool(false),
        },
    },
}

result, err := qot.StockFilter(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, d := range result.DataList {
    fmt.Printf("%s: %s\n", d.Security.GetCode(), d.Name)
}
```

### Get Option Chain (GetOptionChain)

```go
req := &qot.GetOptionChainRequest{
    Owner:      &qotcommon.Security{Market: &market, Code: &code},
    BeginTime:  "2024-01-01",
    EndTime:    "2024-12-31",
    DataFilter: nil,
}

result, err := qot.GetOptionChain(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, chain := range result.OptionChain {
    fmt.Printf("StrikeDate: %s\n", chain.StrikeTime)
    for _, opt := range chain.Option {
        if opt.Call != nil {
            fmt.Printf("  Call: %s\n", opt.Call.GetCode())
        }
        if opt.Put != nil {
            fmt.Printf("  Put: %s\n", opt.Put.GetCode())
        }
    }
}
```

---

## Trading Operations

### Unlock Trading

```go
// Must unlock before trading
err = trd.UnlockTrade(cli, "your_trade_password")
if err != nil {
    log.Fatal(err)
}
```

### Query Account Funds

```go
// Get account list
accList, err := trd.GetAccList(cli)
if err != nil {
    log.Fatal(err)
}

// Use first account
acc := accList[0]

// Query funds
funds, err := trd.GetFunds(cli, acc.AccID, int32(trdcommon.TrdMarket_TrdMarket_HK))
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Cash: %.2f, Frozen: %.2f\n", 
    funds.GetCash(), funds.GetFrozenCash())
```

### Query Positions

```go
positions, err := trd.GetPositionList(cli, acc.AccID, 0, nil)
if err != nil {
    log.Fatal(err)
}

for _, pos := range positions.PositionList {
    fmt.Printf("%s: Qty=%d, Cost=%.2f, Current=%.2f\n",
        pos.Security.GetCode(),
        pos.GetQty(),
        pos.GetCostPrice(),
        pos.GetMarketVal())
}
```

### Place Order

```go
// Buy 100 shares of Tencent
orderID, err := trd.PlaceOrder(cli, &trd.PlaceOrderRequest{
    AccID:        acc.AccID,
    TrdSide:      int32(trdcommon.TrdSide_TrdSide_Buy),
    OrderType:    int32(trdcommon.OrderType_OrderType_Normal),
    Market:       int32(trdcommon.TrdMarket_TrdMarket_HK),
    Security:     &trdcommon.Security{Market: &market, Code: &code},
    Qty:          100,
    Price:        350.00,
    PriceType:    int32(trdcommon.PriceType_PriceType_Normal),
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("OrderID: %s\n", orderID)
```

### Modify Order

```go
err = trd.ModifyOrder(cli, &trd.ModifyOrderRequest{
    AccID:     acc.AccID,
    OrderID:   orderID,
    Market:    int32(trdcommon.TrdMarket_TrdMarket_HK),
    ModifyType: int32(trdcommon.ModifyOrderType_ModifyOrderType_Normal),
    Qty:       200, // Modify quantity
    Price:     360.00, // Modify price
})
if err != nil {
    log.Fatal(err)
}
```

### Query Order List

```go
orders, err := trd.GetOrderList(cli, acc.AccID, 0, nil)
if err != nil {
    log.Fatal(err)
}

for _, o := range orders.OrderList {
    fmt.Printf("Order %s: Status=%d, Qty=%d, Price=%.2f\n",
        o.GetOrderID(), o.GetState(), o.GetQty(), o.GetPrice())
}
```

---

## Real-time Push

### Subscribe to Quotes

```go
// Set push callback
cli.SetQotPushHandler(func(packet *conn.Packet) {
    switch packet.ProtoID {
    case qot.ProtoID_GetBasicQot:
        // Handle quote push
    case qot.ProtoID_GetKL:
        // Handle K-line push
    }
})

// Subscribe to real-time data
security := &qotcommon.Security{Market: &market, Code: &code}
_, err = qot.Subscribe(cli, &qot.SubscribeRequest{
    SecurityList:     []*qotcommon.Security{security},
    SubTypeList:      []qot.SubType{qot.SubType_Basic, qot.SubType_KL},
    IsSubOrUnSub:     true,
    IsRegOrUnRegPush: true,
})
```

### Order Status Push

```go
// Set trading push callback
cli.SetTrdPushHandler(func(packet *conn.Packet) {
    switch packet.ProtoID {
    case trd.ProtoID_UpdateOrder:
        // Handle order update
    case trd.ProtoID_UpdateOrderFill:
        // Handle fill update
    }
})
```

---

## FAQ

### Q: Connection failed, what to do?

1. Confirm Futu OpenD is started and running
2. Confirm port number is correct (default 11111)
3. Confirm network connection is normal

```go
err := cli.Connect("127.0.0.1:11111")
if err != nil {
    log.Fatal("Connection failed:", err)
}
```

### Q: How to handle errors?

All API calls may return errors, recommended to handle uniformly:

```go
result, err := qot.GetBasicQot(cli, securities)
if err != nil {
    // Distinguish error types
    if strings.Contains(err.Error(), "timeout") {
        // Handle timeout
    } else if strings.Contains(err.Error(), "not connected") {
        // Handle disconnect
    } else {
        log.Fatal(err)
    }
}
```

### Q: How to get quotes for multiple stocks?

```go
securities := []*qotcommon.Security{
    {Market: &market, Code: &code1},
    {Market: &market, Code: &code2},
    {Market: &market, Code: &code3},
}

result, err := qot.GetBasicQot(cli, securities)
```

### Q: How to set price alerts?

```go
// Get price alerts
result, err := qot.GetPriceReminder(cli, security, market)

// Setting alerts requires operation in Futu OpenD client
```

### Q: What preparation is needed before trading?

1. Unlock trading password: `trd.UnlockTrade()`
2. Get trading account: `trd.GetAccList()`
3. Ensure account has sufficient funds

---

## Market Constants Reference

### Stock Markets (QotMarket)

| Market | Value | Description |
|------|-----|------|
| HK_Security | 1 | Hong Kong |
| US_Security | 11 | US stocks |
| SH_Security | 31 | Shanghai |
| SZ_Security | 32 | Shenzhen |

### K-Line Types (KLType)

| Type | Value | Description |
|------|-----|------|
| KLType_Min1 | 1 | 1 minute |
| KLType_Min5 | 2 | 5 minutes |
| KLType_Min15 | 3 | 15 minutes |
| KLType_Min30 | 4 | 30 minutes |
| KLType_Min60 | 5 | 60 minutes |
| KLType_Day | 4 | Daily |
| KLType_Week | 5 | Weekly |
| KLType_Month | 6 | Monthly |

### Trade Direction (TrdSide)

| Direction | Value | Description |
|------|-----|------|
| Buy | 1 | Buy |
| Sell | 2 | Sell |

### Order Status (OrderState)

| Status | Value | Description |
|------|-----|------|
| Unknown | 0 | Unknown |
| Submitting | 1 | Submitting |
| Submitted | 2 | Submitted |
| Filled | 3 | Filled |
| PartiallyFilled | 4 | Partially Filled |
| Cancelled | 5 | Cancelled |
| Rejected | 6 | Rejected |
