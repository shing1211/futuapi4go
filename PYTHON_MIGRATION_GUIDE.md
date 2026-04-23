# Python to Go Migration Guide

## Overview

This guide helps Python developers migrate from the `futu-api` Python SDK to the `futuapi4go` Go SDK. The Go SDK is designed to mirror the Python SDK's API structure and naming conventions for easy migration.

## Key Differences

| Aspect | Python SDK | Go SDK |
|--------|-----------|--------|
| Import | `import futu as ft` | `github.com/shing1211/futuapi4go/pkg/constant` |
| Context Manager | `with` statement | `defer ctx.Close()` |
| Async/Handlers | Callback classes | Channel-based or function handlers |
| Enums | Class attributes | Package-level constants |
| Error Handling | Return codes | Go error returns |

## Code Comparison Examples

### 1. Creating a Quote Context

**Python:**
```python
import futu as ft

quote_ctx = ft.OpenQuoteContext(host="127.0.0.1", port=11111)
# Use the context
quote_ctx.close()
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/client"
)

cli := client.New(client.WithAddress("127.0.0.1:11111"))
defer cli.Close()
```

### 2. Market Constants

**Python:**
```python
market = ft.Market.HK
sec_type = ft.SecurityType.STOCK
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/constant"
)

market := constant.Market_HK
secType := constant.SecurityType_Stock
```

### 3. Subscription Types

**Python:**
```python
quote_ctx.subscribe(code, [
    ft.SubType.QUOTE,
    ft.SubType.TICKER,
    ft.SubType.K_DAY,
    ft.SubType.ORDER_BOOK
])
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/qot"
    "github.com/shing1211/futuapi4go/pkg/constant"
)

_, err := qot.Subscribe(cli.Inner(), &qot.SubscribeRequest{
    SecurityList: securities,
    SubTypeList: []qot.SubType{
        qot.SubType(constant.SubType_Quote),
        qot.SubType(constant.SubType_Ticker),
        qot.SubType(constant.SubType_K_Day),
        qot.SubType(constant.SubType_OrderBook),
    },
    IsSubOrUnSub: true,
})
```

### 4. K-line Types

**Python:**
```python
kl_type = ft.KLType.K_DAY
rehab_type = ft.AuType.QFQ  # Forward adjustment
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/constant"
)

klType := constant.KLType_K_Day
rehabType := constant.RehabType_Forward  // QFQ in Python
```

### 5. Trading Environment

**Python:**
```python
trd_env = ft.TrdEnv.SIMULATE  # or ft.TrdEnv.REAL
trd_side = ft.TrdSide.BUY
order_type = ft.OrderType.NORMAL  # or ft.OrderType.MARKET
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/constant"
)

trdEnv := constant.TrdEnv_Simulate  // or constant.TrdEnv_Real
trdSide := constant.TrdSide_Buy
orderType := constant.OrderType_Normal  // or constant.OrderType_Market
```

### 6. Getting Market Snapshot

**Python:**
```python
ret, data = quote_ctx.get_market_snapshot(['HK.00700'])
if ret == ftRET_OK:
    print(data)
```

**Go:**
```go
import (
    "context"
    "github.com/shing1211/futuapi4go/pkg/qot"
    qotcommon "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

market := int32(constant.Market_HK)
code := "00700"
security := &qotcommon.Security{
    Market: &market,
    Code: &code,
}

resp, err := qot.GetSecuritySnapshot(ctx.Background(), cli.Inner(), &qot.GetSecuritySnapshotRequest{
    SecurityList: []*qotcommon.Security{security},
})
if err != nil {
    log.Fatal(err)
}
// Process resp.SnapshotList
```

### 7. Getting Historical K-lines

**Python:**
```python
ret, data = quote_ctx.request_history_kline(
    code='HK.00700',
    start='2024-01-01',
    end='2024-12-31',
    ktype=ft.KLType.K_DAY,
    autype=ft.AuType.QFQ
)
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/qot"
    "github.com/shing1211/futuapi4go/pkg/constant"
)

market := int32(constant.Market_HK)
code := "00700"
security := &qotcommon.Security{
    Market: &market,
    Code: &code,
}

resp, err := qot.RequestHistoryKL(cli.Inner(), &qot.RequestHistoryKLRequest{
    Security:    security,
    RehabType:   int32(constant.RehabType_Forward),
    KlType:      int32(constant.KLType_K_Day),
    BeginTime:   "2024-01-01",
    EndTime:     "2024-12-31",
    MaxAckKLNum: 1000,
})
```

### 8. Placing an Order

**Python:**
```python
ret, data = trade_ctx.place_order(
    price=350.0,
    qty=100,
    code='HK.00700',
    trd_side=ft.TrdSide.BUY,
    order_type=ft.OrderType.NORMAL,
    trd_env=ft.TrdEnv.SIMULATE
)
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/trd"
    "github.com/shing1211/futuapi4go/pkg/constant"
)

resp, err := trd.PlaceOrder(cli.Inner(), &trd.PlaceOrderRequest{
    AccID:     accID,
    TrdMarket: int32(constant.Market_HK),
    TrdEnv:    constant.TrdEnv_Simulate,
    Code:      "00700",
    TrdSide:   constant.TrdSide_Buy,
    OrderType: constant.OrderType_Normal,
    Price:     350.0,
    Qty:       100,
})
```

### 9. Querying Positions

**Python:**
```python
ret, data = trade_ctx.position_list_query(
    trd_env=ft.TrdEnv.SIMULATE,
    acc_id=acc_id
)
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/pkg/trd"
)

resp, err := trd.GetPositionList(cli.Inner(), &trd.GetPositionListRequest{
    AccID:     accID,
    TrdMarket: int32(constant.Market_HK),
    TrdEnv:    constant.TrdEnv_Simulate,
})
```

### 10. Push Handlers

**Python:**
```python
class MyTickerHandler(ft.TickerHandlerBase):
    def on_recv_connect(self, packet_handler):
        return

    def on_recv_data(self, rsp_pb):
        ticker_data = rsp_pb
        print(f"Ticker: {ticker_data}")
        return

quote_ctx.set_handler(MyTickerHandler())
quote_ctx.subscribe('HK.00700', [ft.SubType.TICKER])
```

**Go:**
```go
import (
    "github.com/shing1211/futuapi4go/client"
    "github.com/shing1211/futuapi4go/pkg/push"
)

cli.RegisterHandler(client.ProtoID_Qot_UpdateTicker, func(protoID uint32, body []byte) {
    ticker, err := push.ParseUpdateTicker(body)
    if err != nil || ticker == nil {
        return
    }
    fmt.Printf("Ticker: %v\n", ticker)
})

_, err := qot.Subscribe(cli.Inner(), &qot.SubscribeRequest{
    SecurityList: []*qotcommon.Security{sec},
    SubTypeList: []qot.SubType{qot.SubType(constant.SubType_Ticker)},
    IsSubOrUnSub:     true,
    IsRegOrUnRegPush: true,
})
```

## Complete Enum/Constant Reference

### Markets (行情市场)

| Python | Go | Value |
|--------|-----|-------|
| `ft.Market.HK` | `constant.Market_HK` | 1 |
| `ft.Market.US` | `constant.Market_US` | 11 |
| `ft.Market.SH` | `constant.Market_SH` | 21 |
| `ft.Market.SZ` | `constant.Market_SZ` | 22 |
| `ft.Market.SG` | `constant.Market_SG` | 31 |
| `ft.Market.JP` | `constant.Market_JP` | 41 |
| `ft.Market.AU` | `constant.Market_AU` | 51 |
| `ft.Market.MY` | `constant.Market_MY` | 61 |
| `ft.Market.CA` | `constant.Market_CA` | 71 |
| `ft.Market.FX` | `constant.Market_FX` | 81 |

### Security Types (证券类型)

| Python | Go | Value |
|--------|-----|-------|
| `ft.SecurityType.STOCK` | `constant.SecurityType_Stock` | 3 |
| `ft.SecurityType.ETF` | `constant.SecurityType_ETF` | 4 |
| `ft.SecurityType.WARRANT` | `constant.SecurityType_Warrant` | 5 |
| `ft.SecurityType.INDEX` | `constant.SecurityType_Index` | 6 |
| `ft.SecurityType.FUTURE` | `constant.SecurityType_Future` | 10 |

### K-Line Types (K线类型)

| Python | Go | Value |
|--------|-----|-------|
| `ft.KLType.K_1M` | `constant.KLType_K_1Min` | 1 |
| `ft.KLType.K_5M` | `constant.KLType_K_5Min` | 2 |
| `ft.KLType.K_15M` | `constant.KLType_K_15Min` | 3 |
| `ft.KLType.K_30M` | `constant.KLType_K_30Min` | 4 |
| `ft.KLType.K_60M` | `constant.KLType_K_60Min` | 5 |
| `ft.KLType.K_DAY` | `constant.KLType_K_Day` | 6 |
| `ft.KLType.K_WEEK` | `constant.KLType_K_Week` | 7 |
| `ft.KLType.K_MON` | `constant.KLType_K_Month` | 8 |

### Subscription Types (订阅类型)

| Python | Go | Value |
|--------|-----|-------|
| `ft.SubType.QUOTE` | `constant.SubType_Quote` | 1 |
| `ft.SubType.ORDER_BOOK` | `constant.SubType_OrderBook` | 2 |
| `ft.SubType.TICKER` | `constant.SubType_Ticker` | 4 |
| `ft.SubType.RT_DATA` | `constant.SubType_RT` | 5 |
| `ft.SubType.K_DAY` | `constant.SubType_K_Day` | 6 |
| `ft.SubType.K_1M` | `constant.SubType_K_1Min` | 11 |
| `ft.SubType.BROKER` | `constant.SubType_Broker` | 14 |

### Rehab Types (复权类型)

| Python | Go | Value |
|--------|-----|-------|
| `ft.AuType.NONE` | `constant.RehabType_None` | 0 |
| `ft.AuType.QFQ` | `constant.RehabType_Forward` | 1 |
| `ft.AuType.HFQ` | `constant.RehabType_Backward` | 2 |

### Trading Environment (交易环境)

| Python | Go | Value |
|--------|-----|-------|
| `ft.TrdEnv.REAL` | `constant.TrdEnv_Real` | 1 |
| `ft.TrdEnv.SIMULATE` | `constant.TrdEnv_Simulate` | 0 |

### Trade Sides (交易方向)

| Python | Go | Value |
|--------|-----|-------|
| `ft.TrdSide.BUY` | `constant.TrdSide_Buy` | 1 |
| `ft.TrdSide.SELL` | `constant.TrdSide_Sell` | 2 |

### Order Types (订单类型)

| Python | Go | Value |
|--------|-----|-------|
| `ft.OrderType.NORMAL` | `constant.OrderType_Normal` | 1 |
| `ft.OrderType.MARKET` | `constant.OrderType_Market` | 2 |

### Order Status (订单状态)

| Python | Go | Value |
|--------|-----|-------|
| `ft.OrderStatus.SUBMITTED` | `constant.OrderStatus_Submitted` | 6 |
| `ft.OrderStatus.FILLED_ALL` | `constant.OrderStatus_FilledAll` | 8 |
| `ft.OrderStatus.CANCELLED_ALL` | `constant.OrderStatus_CancelledAll` | 12 |

### Modify Order Operations (改单操作)

| Python | Go | Value |
|--------|-----|-------|
| `ft.ModifyOrderOp.NORMAL` | `constant.ModifyOrderOp_Normal` | 1 |
| `ft.ModifyOrderOp.CANCEL` | `constant.ModifyOrderOp_Cancel` | 2 |

## Common Patterns

### Pattern 1: Error Handling

**Python:**
```python
ret, data = quote_ctx.get_market_snapshot(['HK.00700'])
if ret != ftRET_OK:
    print(f"Error: {data}")
else:
    print(data)
```

**Go:**
```go
resp, err := qot.GetSecuritySnapshot(ctx.Background(), cli.Inner(), req)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return err
}
// Process resp
```

### Pattern 2: Timeouts

**Python:**
```python
# Python handles timeouts differently
quote_ctx.set_max_wait_time(30)  # Set timeout
```

**Go:**
```go
import "context"

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

resp, err := qot.GetSecuritySnapshot(ctx, cli.Inner(), req)
// Context will timeout automatically
```

### Pattern 3: Pagination

**Python:**
```python
ret, data = quote_ctx.request_history_kline(
    code='HK.00700',
    start='2024-01-01',
    end='2024-12-31',
    ktype=ft.KLType.K_DAY,
    autype=ft.AuType.NONE,
    max_count=1000
)
```

**Go:**
```go
resp, err := qot.RequestHistoryKL(cli.Inner(), &qot.RequestHistoryKLRequest{
    Security:    security,
    KlType:      constant.KLType_K_Day,
    RehabType:   constant.RehabType_None,
    BeginTime:   "2024-01-01",
    EndTime:     "2024-12-31",
    MaxAckKLNum: 1000,
})

// If NextReqKey is not empty, there are more pages
for len(resp.NextReqKey) > 0 {
    resp, err = qot.RequestHistoryKL(cli.Inner(), &qot.RequestHistoryKLRequest{
        Security:    security,
        KlType:      constant.KLType_K_Day,
        RehabType:   constant.RehabType_None,
        BeginTime:   "2024-01-01",
        EndTime:     "2024-12-31",
        MaxAckKLNum: 1000,
        NextReqKey:  resp.NextReqKey,  // Pass the pagination key
    })
}
```

## Performance Tips

1. **Use Context for Cancellation:**
   - Always pass context for request cancellation
   - Set appropriate timeouts

2. **Connection Pooling:**
   - The Go SDK supports connection pooling
   - Use `client.New()` for automatic connection management

3. **Concurrent Requests:**
   - Go's goroutines make concurrent requests easy
   - Use `sync.WaitGroup` for parallel operations

```go
var wg sync.WaitGroup
results := make(chan *qot.SecuritySnapshotResponse, len(securities))

for _, sec := range securities {
    wg.Add(1)
    go func(s *qotcommon.Security) {
        defer wg.Done()
        resp, err := qot.GetSecuritySnapshot(ctx, cli.Inner(), &qot.GetSecuritySnapshotRequest{
            SecurityList: []*qotcommon.Security{s},
        })
        if err == nil {
            results <- resp
        }
    }(sec)
}

wg.Wait()
close(results)
```

## What's Next?

- Check the `AGENTS.md` file for SDK debugging tips
- Look at the `examples/` directory for complete code examples
- Review the `pkg/constant/constant.go` file for all available constants
- See the demo project `futuapi4go-demo` for a comprehensive example

## Getting Help

If you encounter issues:

1. Check the Python SDK documentation: https://openapi.futunn.com/futu-api-doc/
2. Open an issue on GitHub: https://github.com/shing1211/futuapi4go/issues
3. Review the proto reference: `docs/FUTU_PROTO_REF.md`
