# FutuAPI4Go Examples

This directory contains comprehensive examples demonstrating how to use the FutuAPI4Go SDK for market data queries, trading operations, and real-time subscriptions.

## Example Programs

### 1. **01_market_data_basic** - Basic Market Data APIs
**Purpose**: Demonstrates fundamental market data APIs  
**APIs Covered**:
- `GetBasicQot` - Real-time quotes
- `GetKL` - K-line (candlestick) data
- `GetOrderBook` - Order book (bid/ask)
- `GetTicker` - Tick-by-tick trades
- `GetRT` - Real-time minute data
- `GetBroker` - Broker queue
- `GetSecuritySnapshot` - Comprehensive stock snapshot

**Run**:
```bash
cd 01_market_data_basic
go run main.go
```

**Best For**: Beginners learning the SDK, basic market data retrieval

---

### 2. **02_market_data_advanced** - Advanced Market Data APIs
**Purpose**: Shows advanced market analysis capabilities  
**APIs Covered**:
- `GetStaticInfo` - Stock static information
- `GetPlateSet` - Plate/sector information
- `GetCapitalFlow` - Capital flow analysis
- `GetCapitalDistribution` - Capital distribution by size
- `StockFilter` - Screen stocks by criteria
- `GetOptionExpirationDate` - Option expiration dates
- `GetWarrant` - Warrant information
- `GetTradeDate` - Trading calendar
- `GetFutureInfo` - Futures data
- `GetIpoList` - IPO listings

**Run**:
```bash
cd 02_market_data_advanced
go run main.go
```

**Best For**: Advanced market analysis, stock screening, options research

---

### 3. **03_trading_operations** - Trading APIs
**Purpose**: Complete trading workflow examples  
**APIs Covered**:
- `GetAccList` - Get account list
- `UnlockTrade` - Unlock trading password
- `GetFunds` - Get account funds
- `GetPositionList` - Get positions
- `PlaceOrder` - Place orders
- `GetOrderList` - Query orders
- `ModifyOrder` - Modify orders
- `GetOrderFillList` - Get order fills
- `GetOrderFee` - Get order fees
- `GetMarginRatio` - Get margin ratio
- `GetMaxTrdQtys` - Get max trade quantities
- `GetHistoryOrderList` - Get historical orders

**Run**:
```bash
cd 03_trading_operations
go run main.go
```

**Best For**: Learning trading operations, order management, position tracking

**Important**: Trading APIs require `UnlockTrade` to be called first in real environments!

---

### 4. **04_push_subscriptions** - Real-time Push Notifications
**Purpose**: Asynchronous real-time data subscriptions  
**APIs Covered**:
- `Subscribe` - Subscribe to real-time data streams
- `GetSubInfo` - Get subscription information
- `RegQotPush` - Register for push notifications
- Push handlers for market data
- Push handlers for trading updates

**Run**:
```bash
cd 04_push_subscriptions
go run main.go
```

**Best For**: Real-time monitoring, algorithmic trading, live dashboards

---

### 5. **05_comprehensive_demo** - All-in-One Showcase
**Purpose**: Demonstrates full SDK capabilities in one program  
**Features**:
- Connection management
- Market data queries (basic and advanced)
- Trading operations
- Real-time subscriptions
- Error handling
- Beautiful formatted output

**Run**:
```bash
cd 05_comprehensive_demo
go run main.go
```

**Best For**: Quick overview of SDK capabilities, testing installation

---

## Quick Start

### Using with Simulator (Recommended for Development)

1. **Start the simulator**:
```bash
go run ./cmd/simulator
```

2. **Run any example** (in another terminal):
```bash
cd cmd/examples/01_market_data_basic
go run main.go
```

### Using with Real Futu OpenD

1. **Ensure Futu OpenD is running** on `127.0.0.1:11111`

2. **Set environment variable** (optional):
```bash
export FUTU_ADDR="127.0.0.1:11111"
```

3. **Run examples**:
```bash
cd 01_market_data_basic
go run main.go
```

---

## Example Structure

Each example follows a consistent structure:

```go
package main

import (
    "fmt"
    "log"

    futuapi "github.com/shing1211/futuapi4go/internal/client"
    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
    "github.com/shing1211/futuapi4go/pkg/qot"
    "github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
    // 1. Create client
    cli := futuapi.New()
    defer cli.Close()
    
    // 2. Connect
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    
    // 3. Call APIs
    // ... API-specific code ...
}
```

---

## Best Practices

### 1. **Connection Management**
```go
cli := futuapi.New()
defer cli.Close() // Always close connection when done

if err := cli.Connect(addr); err != nil {
    log.Fatal(err)
}
```

### 2. **Error Handling**
```go
result, err := qot.GetBasicQot(cli, securities)
if err != nil {
    log.Printf("GetBasicQot failed: %v", err)
    // Handle error appropriately
}
```

### 3. **Resource Management**
- Always `defer cli.Close()` after successful connection
- Unsubscribe from push notifications when no longer needed
- Monitor subscription quota with `GetSubInfo`

### 4. **Push Notifications**
```go
// Set up handler BEFORE subscribing
cli.SetQotPushHandler(func(pkt *client.Packet) {
    switch pkt.ProtoID {
    case qot.ProtoID_Qot_UpdateBasicQot:
        // Handle basic quote update
    case qot.ProtoID_Qot_UpdateKL:
        // Handle K-line update
    }
})

// Then subscribe
qot.Subscribe(cli, subReq)
```

---

## Common Use Cases

### Market Data Retrieval
```go
// Get real-time quotes for multiple stocks
securities := []*qotcommon.Security{
    {Market: &hkMarket, Code: ptrStr("00700")},
    {Market: &hkMarket, Code: ptrStr("09988")},
}
quotes, _ := qot.GetBasicQot(cli, securities)
```

### Placing Orders
```go
orderReq := &trd.PlaceOrderRequest{
    AccID:     accID,
    TrdMarket: hkMarket,
    Code:      "00700",
    TrdSide:   int32(trdcommon.TrdSide_TrdSide_Buy),
    OrderType: int32(trdcommon.OrderType_OrderType_Normal),
    Price:     350.00,
    Qty:       100.0,
}
resp, _ := trd.PlaceOrder(cli, orderReq)
```

### Real-time Monitoring
```go
// Subscribe to multiple data types
subReq := &qot.SubscribeRequest{
    SecurityList:     securities,
    SubTypeList:      []qot.SubType{
        qot.SubType_Basic,
        qot.SubType_KL,
        qot.SubType_OrderBook,
    },
    IsSubOrUnSub:     true,
    IsRegOrUnRegPush: true,
}
qot.Subscribe(cli, subReq)
```

---

## Helper Functions

All examples include helper functions for creating pointers:

```go
func ptrStr(s string) *string { return &s }
func ptrInt32(v int32) *int32 { return &v }
func ptrFloat64(v float64) *float64 { return &v }
func ptrBool(v bool) *bool { return &v }
```

These are necessary because protobuf fields require pointer types.

---

## Notes

1. **Simulator vs Real OpenD**: 
   - Simulator returns mock data for testing
   - Real OpenD provides live market data
   - Examples work with both!

2. **API Coverage**:
   - All documented APIs are implemented
   - Some advanced APIs may return empty results in simulator (stubs)

3. **Trading Safety**:
   - Always test with simulator first
   - Use small quantities in real trading
   - Double-check order parameters before submission

---

## Additional Resources

- [docs/DEVELOPER.md](../docs/DEVELOPER.md) - SDK developer guide
- [docs/API_REFERENCE.md](../docs/API_REFERENCE.md) - API reference
- [docs/CHANGELOG.md](../docs/CHANGELOG.md) - Version history
- [ROADMAP.md](../ROADMAP.md) - Project roadmap

---

## Troubleshooting

### Connection Failed
```
Connection failed: dial: connection refused
```
**Solution**: Ensure Futu OpenD or simulator is running on the specified address

### API Returns Empty Results
```
Found 0 positions
```
**Solution**: Normal for simulator; try with real OpenD for actual data

### Subscription Quota Exceeded
```
Subscribe failed: retType=-1
```
**Solution**: Check quota with `GetSubInfo` and unsubscribe from unused subscriptions

---

## Support

For issues or questions:
- Open an issue on [github](https://github.com/shing1211/futuapi4go)
- Check existing documentation
- Review example code for usage patterns

---

**Happy Coding!**
