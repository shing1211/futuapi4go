# Migration Guide

This guide helps developers migrate from other SDKs (primarily the Python `py-futu-api`) to futuapi4go.

## Table of Contents

1. [Overview](#overview)
2. [Key Differences](#key-differences)
3. [Basic Migration](#basic-migration)
4. [API Mapping](#api-mapping)
5. [Trading Migration](#trading-migration)
6. [Push Notifications](#push-notifications)
7. [Common Patterns](#common-patterns)
8. [v0.6.0 Breaking Changes](#v060-breaking-changes)

---

## Overview

| Aspect | Python (py-futu-api) | Go (futuapi4go) |
|--------|----------------------|----------------|
| Package | `futu` | `github.com/shing1211/futuapi4go` |
| Client | `FutUOpenAPI` | `client.Client` or `futuapi.Client` |
| Connection | auto-initialized | manual `Connect()` |
| Async | asyncio | goroutines + callbacks |
| Types | proto-generated classes | struct types |

---

## Key Differences

### 1. Client Creation

**Python:**
```python
from futu import FutUOpenAPI
ctx = FutUOpenAPI(host="127.0.0.1", port=11111)
ctx.start()
```

**Go (High-level):**
```go
import "github.com/shing1211/futuapi4go/client"

cli := client.New()
defer cli.Close()
if err := cli.Connect("127.0.0.1:11111"); err != nil {
    log.Fatal(err)
}
```

**Go (Low-level):**
```go
import futuapi "github.com/shing1211/futuapi4go/internal/client"

cli := futuapi.New()
defer cli.Close()
if err := cli.Connect("127.0.0.1:11111"); err != nil {
    log.Fatal(err)
}
```

### 2. Security Definition

Python uses tuples or the `Security` class:

```python
from futu import Security
security = Security("HK", "00700")
# or tuple: ("HK", "00700")
```

Go requires pointer to struct:

```go
import "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
import "github.com/shing1211/futuapi4go/pkg/constant"

hk := int32(constant.Market_HK)
code := "00700"
security := &qotcommon.Security{
    Market: &hk,
    Code:   &code,
}
```

Helper function for cleaner code:

```go
func security(market constant.Market, code string) *qotcommon.Security {
    m := int32(market)
    return &qotcommon.Security{
        Market: &m,
        Code:   &code,
    }
}

// Usage - no cast needed!
sec := security(constant.Market_HK, "00700")
```

---

## Basic Migration

### Get Quote

**Python:**
```python
from futu import StockQuote
ret, data = ctx.query(StockQuote("00700"))
for q in data:
    print(q['cur_price'])
```

**Go:**
```go
quote, err := client.GetQuote(cli, client.Market_HK_Security, "00700")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Price: %.2f\n", quote.CurPrice)
```

### Get K-Line

**Python:**
```python
from futu import KLType, AuType
ret, data = ctx.query_history_kline("00700", KLType.KL_DAY, AuType.qfq, 100)
for k in data:
    print(k['close'])
```

**Go:**
```go
klines, err := client.GetKLines(cli, client.Market_HK_Security, "00700", client.KLType_Day, 100)
if err != nil {
    log.Fatal(err)
}
for _, kl := range klines {
    fmt.Printf("Close: %.2f\n", kl.Close)
}
```

### Get Order Book

**Python:**
```python
ret, data = ctx.query_order_book("00700", 10)
print(data['bid'])
print(data['ask'])
```

**Go:**
```go
ob, err := client.GetOrderBook(cli, client.Market_HK_Security, "00700", 10)
if err != nil {
    log.Fatal(err)
}
for _, bid := range ob.BidList {
    fmt.Printf("Bid: %.2f x %.0f\n", bid.Price, bid.Volume)
}
for _, ask := range ob.AskList {
    fmt.Printf("Ask: %.2f x %.0f\n", ask.Price, ask.Volume)
}
```

### Get Multiple Quotes

**Python:**
```python
ret, data = ctx.subscribe(["00700", "09988"], sub_types=[SubType QUOTE])
```

**Go:**
```go
// Set up subscription
securities := []*qotcommon.Security{
    security(client.Market_HK_Security, "00700"),
    security(client.Market_HK_Security, "09988"),
}
// Subscribe first
subReq := &qot.SubscribeRequest{
    SecurityList: securities,
    SubTypeList: []qot.SubType{qot.SubType_Basic},
    IsSubOrUnSub: true,
}
_, err := qot.Subscribe(cli, subReq)
// Then get snapshot
snapshot, err := client.GetSecuritySnapshot(cli, securities)
```

---

## API Mapping

### Market Data APIs

| Python | Go (High-level) | Go (Low-level) |
|--------|---------------|---------------|
| `query(StockQuote)` | `client.GetQuote()` | `qot.GetBasicQot()` |
| `query_history_kline()` | `client.GetKLines()` | `qot.GetKL()` |
| `query_order_book()` | `client.GetOrderBook()` | `qot.GetOrderBook()` |
| `subscribe()` | `client.Subscribe()` | `qot.Subscribe()` |
| `get_stock_basicinfo()` | `client.GetStaticInfo()` | `qot.GetStaticInfo()` |
| `get_plate_stock()` | `client.GetPlateSecurity()` | `qot.GetPlateSecurity()` |
| `get_plate_list()` | `client.GetPlateSet()` | `qot.GetPlateSet()` |
| `get_owner_plate()` | `client.GetOwnerPlate()` | `qot.GetOwnerPlate()` |
| `get_market_state()` | `client.GetMarketState()` | `qot.GetMarketState()` |
| `get_trade_date()` | `client.GetTradeDate()` | `qot.GetTradeDate()` |
| `get_rt()` | `client.GetRT()` | `qot.GetRT()` |
| `get_ticker()` | `client.GetTicker()` | `qot.GetTicker()` |
| `get_broker()` | `client.GetBroker()` | `qot.GetBroker()` |
| `get_capital_flow()` | `client.GetCapitalFlow()` | `qot.GetCapitalFlow()` |
| `get_capital_distribution()` | `client.GetCapitalDistribution()` | `qot.GetCapitalDistribution()` |
| `get_option_chain()` | `client.GetOptionChain()` | `qot.GetOptionChain()` |
| `get_warrant()` | `client.GetWarrant()` | `qot.GetWarrant()` |
| `stock_filter()` | `client.StockFilter()` | `qot.StockFilter()` |

### Trading APIs

| Python | Go (High-level) | Go (Low-level) |
|--------|---------------|---------------|
| `get_acc_list()` | `client.GetAccountList()` | `trd.GetAccList()` |
| `unlock_trade()` | `client.UnlockTrading()` | `trd.UnlockTrade()` |
| `place_order()` | `client.PlaceOrder()` | `trd.PlaceOrder()` |
| `modify_order()` | `client.ModifyOrder()` | `trd.ModifyOrder()` |
| `get_order_list()` | `client.GetOrderList()` | `trd.GetOrderList()` |
| `get_order_fill_list()` | `client.GetOrderFillList()` | `trd.GetOrderFillList()` |
| `get_position_list()` | `client.GetPositionList()` | `trd.GetPositionList()` |
| `get_history_order_list()` | `client.GetHistoryOrderList()` | `trd.GetHistoryOrderList()` |
| `get_funds()` | `client.GetFunds()` | `trd.GetFunds()` |
| `get_order_fee()` | `client.GetOrderFee()` | `trd.GetOrderFee()` |
| `get_max_trd_qtys()` | `client.GetMaxTrdQtys()` | `trd.GetMaxTrdQtys()` |

### System APIs

| Python | Go (High-level) | Go (Low-level) |
|--------|---------------|---------------|
| `get_global_state()` | `client.GetGlobalState()` | `sys.GetGlobalState()` |
| `get_user_info()` | `client.GetUserInfo()` | `sys.GetUserInfo()` |
| `keep_alive()` | (automatic) | (automatic) |

### Enums

| Python | Go |
|--------|-----|
| `KLType.KL_1MIN` | `qotcommon.KLType_KLType_1Min` |
| `KLType.KL_DAY` | `qotcommon.KLType_KLType_Day` |
| `SubType.QUOTE` | `qot.SubType_Basic` |
| `SubType.TICKER` | `qot.SubType_Ticker` |
| `TrdSide.BUY` | `trdcommon.TrdSide_TrdSide_Buy` |
| `TrdSide.SELL` | `trdcommon.TrdSide_TrdSide_Sell` |
| `OrderType.NORMAL` | `trdcommon.OrderType_OrderType_Normal` |
| `OrderType.MARKET` | `trdcommon.OrderType_OrderType_Market` |

---

## Trading Migration

### Setup and Unlock Trading

**Python:**
```python
# Get account list
ret, data = ctx.get_acc_list()
print(data)

# Unlock trading
ret = ctx.unlock_trade("your_password_md5")
print(ret)
```

**Go:**
```go
import "github.com/shing1211/futuapi4go/client"

accs, err := client.GetAccountList(cli)
if err != nil {
    log.Fatal(err)
}
accID := accs[0].AccID

err := client.UnlockTrading(cli, "your_password_md5")
if err != nil {
    log.Fatal(err)
}
```

### Place Order

**Python:**
```python
ret, data = ctx.place_order(
    trd_side=富途.TrdSide.BUY,
    code="00700",
    price=350.00,
    quantity=100,
    order_type=富途.OrderType.NORMAL
)
print(data)
```

**Go:**
```go
result, err := client.PlaceOrder(cli,
    accID,
    client.Market_HK_Security,
    "00700",
    client.Side_Buy,
    client.OrderType_Normal,
    350.00,
    100)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("OrderID: %d\n", result.OrderID)
```

### Query Positions

**Python:**
```python
ret, data = ctx.get_position_list()
for pos in data:
    print(f"{pos['code']}: {pos['qty']}")
```

**Go:**
```go
positions, err := client.GetPositionList(cli, accID)
if err != nil {
    log.Fatal(err)
}
for _, p := range positions {
    fmt.Printf("%s: %.0f\n", p.Code, p.Qty)
}
```

---

## Push Notifications

### Subscribe to Push

**Python:**
```python
from futu import SubType

# Set handler
def on_push(pack):
    print(pack)

ctx.subscribe_handler(pack, on_push)

# Subscribe
ctx.subscribe("00700", [SubType.QUOTE])
```

**Go:**
```go
import "github.com/shing1211/futuapi4go/pkg/push"

// Set up push handler
cli.SetPushHandler(func(header *push.Header, body []byte) {
    switch header.ProtoID {
    case push.ProtoID_Qot_UpdateBasicQot:
        data, _ := push.ParseUpdateBasicQot(body)
        fmt.Printf("Quote: %s -> %.2f\n", 
            data.Security.GetCode(), data.CurPrice)
    case push.ProtoID_Qot_UpdateKL:
        data, _ := push.ParseUpdateKL(body)
        fmt.Printf("KL: %s -> %.2f\n",
            data.Security.GetCode(), data.KL.ClosePrice)
    }
})

// Subscribe
err := client.Subscribe(cli, client.Market_HK_Security, "00700",
    []int32{client.SubType_Basic, client.SubType_KL})
```

### Push Types

| Push Type | Python | Go |
|----------|--------|-----|
| Quote update | `SubType.QUOTE` | `push.ProtoID_Qot_UpdateBasicQot` |
| K-line update | `SubType.K_LINE` | `push.ProtoID_Qot_UpdateKL` |
| Ticker update | `SubType.TICKER` | `push.ProtoID_Qot_UpdateTicker` |
| Order book | `SubType.BROKER` | `push.ProtoID_Qot_UpdateOrderBook` |
| Order update | `TrdUpdateOrder` | `push.ProtoID_Trd_UpdateOrder` |
| Order fill | `TrdUpdateOrderFill` | `push.ProtoID_Trd_UpdateOrderFill` |

---

## Common Patterns

### Error Handling

**Python:**
```python
ret, data = ctx.query_stock_quote("00700")
if ret == RETTYPE.OK:
    print(data)
else:
    print(f"Error: {ret}")
```

**Go:**
```go
quote, err := client.GetQuote(cli, client.Market_HK_Security, "00700")
if err != nil {
    log.Printf("Error: %v", err)
    return
}
fmt.Printf("Quote: %.2f\n", quote.CurPrice)
```

### Market Constants

```go
import "github.com/shing1211/futuapi4go/client"

// Markets
client.Market_HK_Security
client.Market_US_Security
client.Market_CNSH_Security  // Shanghai
client.Market_CNSZ_Security  // Shenzhen
client.Market_HK_Future

// K-Line Types
client.KLType_1Min
client.KLType_5Min
client.KLType_15Min
client.KLType_30Min
client.KLType_60Min
client.KLType_Day
client.KLType_Week
client.KLType_Month

// Subscription Types
client.SubType_Basic    // Quote
client.SubType_KL       // K-line
client.SubType_Ticker  // Tick
client.SubType_OrderBook
client.SubType_Broker
client.SubType_RT

// Trade Types
client.Side_Buy
client.Side_Sell
client.OrderType_Normal
client.OrderType_Market
```

### Using Context for Timeout

**Python:**
```python
ctx = FutUOpenAPI(...)
ctx.start()
# Use timeout parameter
ret, data = ctx.query(kline, timeout=5)
```

**Go:**
```go
cli := futuapi.New()
cli.SetReadTimeout(5 * time.Second)
cli.SetWriteTimeout(5 * time.Second)
```

---

## v0.6.0 Breaking Changes

### WithTradeEnv / GetTradeEnv Signature Change

**Before:**
```go
cli := client.New()
cli = cli.WithTradeEnv(1) // raw int32
env := cli.GetTradeEnv()  // returns int32
```

**After:**
```go
cli := client.New()
cli = cli.WithTradeEnv(constant.TrdEnv_Real)  // typed enum
env := cli.GetTradeEnv()                       // returns constant.TrdEnv
```

### WithTradeMarket / GetTradeMarket Signature Change

**Before:**
```go
cli = cli.WithTradeMarket(1) // raw int32
```

**After:**
```go
cli = cli.WithTradeMarket(constant.TrdMarket_HK) // typed enum
mkt := cli.GetTradeMarket()                       // returns constant.TrdMarket
```

### OrderBuilder.Build() Now Returns Error

**Before:**
```go
req := trd.NewOrder(accID, market, env).Buy("00700", 100).At(350.5).Build()
```

**After:**
```go
req, err := trd.NewOrder(accID, market, env).Buy("00700", 100).At(350.5).Build()
if err != nil {
    return err
}
```

### FutuError Additional Fields

FutuError now has `Category` and `Recovery` fields. Existing code accessing `FutuError{Code, Message, Func}` continues to work. New fields are auto-populated by `NewFutuError()` / `NewFutuErrorWithWrap()`.

### New Error Predicates

```go
// Classify errors programmatically
if constant.IsConnectionError(err) { /* reconnect */ }
if constant.IsTradingError(err)    { /* handle trading error */ }
cat := constant.CategoryOf(err)     // ErrorCategory: "connection", "trading", etc.
hint := constant.RecoveryHint(err)  // human-readable recovery suggestion
```

### Circuit Breaker

```go
cb := breaker.New(breaker.WithFailureThreshold(5), breaker.WithCooldown(10*time.Second))
cli := futuapi.New(futuapi.WithBreaker(cb))
// Business API calls are now protected; InitConnect/KeepAlive bypass breaker
```

### OrderBuilder AutoDetectMarket

```go
req, err := trd.NewOrder(accID, 0, env).
    Buy("00700.HK", 100).
    At(350.5).
    AutoDetectMarket(). // sets TrdMarket + SecMarket from code suffix
    Build()
```

---

## Complete Example

**Python:**
```python
from futu import FutUOpenAPI, SubType

ctx = FutUOpenAPI(host="127.0.0.1", port=11111)
ctx.start()

# Get quote
ret, quote = ctx.query(StockQuote("00700"))
print(quote[0]['cur_price'])

# Subscribe
ctx.subscribe_handler(SubType.QUOTE, on_push)
ctx.subscribe("00700", [SubType.QUOTE])

ctx.stop()
```

**Go:**
```go
import (
    "fmt"
    "log"

    "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/pkg/push"
)

func main() {
    cli := client.New()
    defer cli.Close()

    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }

    // Push handler
    cli.SetPushHandler(func(h *push.Header, body []byte) {
        data, err := push.ParseUpdateBasicQot(body)
        if err == nil {
            fmt.Printf("%s: %.2f\n", data.Security.GetCode(), data.CurPrice)
        }
    })

    // Get quote
    quote, err := cli.GetQuote(cli, client.Market_HK_Security, "00700")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Price: %.2f\n", quote.CurPrice)

    // Subscribe
    if err := cli.Subscribe(cli, client.Market_HK_Security, "00700",
        []int32{client.SubType_Basic}); err != nil {
        log.Printf("Subscribe error: %v", err)
    }

    // Keep running
    select {}
}
```