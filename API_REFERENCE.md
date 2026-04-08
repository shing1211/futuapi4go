# FutuAPI4Go API Reference

A complete reference for the `futuapi4go` SDK. All functions accept a `*futuapi.Client` as the first parameter and return `(result, error)`.

## Table of Contents

- [Client](#client)
- [Market Data (qot)](#market-data-qot)
  - [Quote Data](#quote-data)
  - [K-Line Data](#k-line-data)
  - [Order Book & Trades](#order-book--trades)
  - [Market Information](#market-information)
  - [Stock Screening & Analysis](#stock-screening--analysis)
  - [Securities & Plates](#securities--plates)
  - [History Data](#history-data)
  - [Subscriptions & Push](#subscriptions--push)
- [Trading (trd)](#trading-trd)
  - [Account Management](#account-management)
  - [Market Data](#market-data)
  - [Order Management](#order-management)
  - [Order History](#order-history)
  - [Advanced Trading](#advanced-trading)
- [System (sys)](#system-sys)

---

## Client

### `futuapi.New(opts ...Option) *Client`

Creates a new client with optional configuration.

```go
cli := futuapi.New(
    futuapi.WithDialTimeout(10 * time.Second),
    futuapi.WithKeepAliveInterval(30 * time.Second),
    futuapi.WithMaxRetries(3),
    futuapi.WithLogLevel(0),
)
defer cli.Close()
```

### `cli.Connect(addr string) error`

Connects to Futu OpenD at the given address. Address format: `host:port` (e.g., `127.0.0.1:11111`).

```go
if err := cli.Connect("127.0.0.1:11111"); err != nil {
    log.Fatalf("Connect failed: %v", err)
}
```

### `cli.Close()`

Closes the TCP connection and stops all background goroutines.

### `cli.EnsureConnected() error`

Returns an error if the client is not connected. Use this at the start of every API call.

### `cli.IsConnected() bool`

Returns whether the client is currently connected.

### `cli.GetConnID() uint64`

Returns the connection ID assigned by OpenD after successful connection.

### `cli.GetServerVer() int32`

Returns the OpenD server version.

### `cli.NextSerialNo() uint32`

Returns the next serial number for packet tracking.

### `cli.WithContext(ctx context.Context) *Client`

Returns a new client that uses the given context for timeouts and cancellation.

### `cli.Context() context.Context`

Returns the client's current context.

### `cli.RegisterHandler(protoID uint32, h Handler)`

Registers a handler for incoming push notifications with the given protocol ID.

### `cli.SetPushHandler(h PacketHandler)`

Sets the global push notification handler.

---

## Market Data (qot)

Import: `"gitee.com/shing1211/futuapi4go/pkg/qot"`

All functions in this package require a connected client. They use the `qotcommon` package for security definitions and enum values.

### Quote Data

#### `GetBasicQot(c *Client, securityList []*qotcommon.Security) ([]*BasicQot, error)`

Retrieves real-time quotes for one or more securities.

```go
hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
securities := []*qotcommon.Security{
    {Market: &hkMarket, Code: ptrStr("00700")},
}
quotes, err := qot.GetBasicQot(cli, securities)
for _, q := range quotes {
    fmt.Printf("%s: %.2f (Vol: %d)\n", q.Name, q.CurPrice, q.Volume)
}
```

**Response fields:** `Security`, `Name`, `IsSuspended`, `UpdateTime`, `HighPrice`, `OpenPrice`, `LowPrice`, `CurPrice`, `LastClosePrice`, `Volume`, `Turnover`, `TurnoverRate`, `Amplitude`

---

### K-Line Data

#### `GetKL(c *Client, req *GetKLRequest) (*GetKLResponse, error)`

Retrieves K-line (candlestick) data for a security.

```go
req := &qot.GetKLRequest{
    Security:  security,
    RehabType: int32(qotcommon.RehabType_RehabType_None),
    KLType:    int32(qotcommon.KLType_KLType_Day),
    ReqNum:    100,
}
resp, err := qot.GetKL(cli, req)
// KLType values: KLType_1Min, KLType_5Min, KLType_15Min, KLType_30Min,
//                KLType_60Min, KLType_Day, KLType_Week, KLType_Month
// RehabType values: RehabType_None, RehabType_Forward, RehabType_Backward
```

**Request fields:** `Security`, `RehabType`, `KLType`, `ReqNum`

**Response fields:** `Security`, `Name`, `KLList` (each KLine has: `Time`, `IsBlank`, `HighPrice`, `OpenPrice`, `LowPrice`, `ClosePrice`, `LastClosePrice`, `Volume`, `Turnover`, `ChangeRate`, `Timestamp`)

#### `RequestHistoryKL(c *Client, req *RequestHistoryKLRequest) (*RequestHistoryKLResponse, error)`

Requests historical K-line data (server-side pagination).

```go
req := &qot.RequestHistoryKLRequest{
    Security:       security,
    RehabType:      int32(qotcommon.RehabType_RehabType_None),
    KlType:         int32(qotcommon.KLType_KLType_Day),
    BeginTime:       "2024-01-01",
    EndTime:         "2024-12-31",
    AutoComplete:    false,
}
```

**Request fields:** `Security`, `RehabType`, `KlType`, `BeginTime`, `EndTime`, `AutoComplete`, `MaxAckKLNum`, `NeedKLFieldsFlags`, `Sort`

**Response fields:** `Security`, `ReqNum`, `PropertyList`, `KLList`, `NextTime`

#### `RequestHistoryKLQuota(c *Client, req *RequestHistoryKLQuotaRequest) (*RequestHistoryKLQuotaResponse, error)`

Queries the quota for historical K-line requests.

---

### Order Book & Trades

#### `GetOrderBook(c *Client, req *GetOrderBookRequest) (*GetOrderBookResponse, error)`

Retrieves the current order book (买卖盘) for a security.

```go
req := &qot.GetOrderBookRequest{
    Security: security,
    Num:      10,
}
resp, err := qot.GetOrderBook(cli, req)
// resp.OrderBookAskList:卖盘 (ask), resp.OrderBookBidList:买盘 (bid)
```

**Request fields:** `Security`, `Num`

**Response fields:** `Security`, `Name`, `OrderBookAskList`, `OrderBookBidList`, `SvrRecvTimeBid`, `SvrRecvTimeAsk`

Each OrderBook entry has: `Price`, `Volume`, `OrderCount`, `DetailList`

#### `GetTicker(c *Client, req *GetTickerRequest) (*GetTickerResponse, error)`

Retrieves tick-by-tick transaction data (逐笔成交).

```go
req := &qot.GetTickerRequest{
    Security:  security,
    MaxRetNum: 100,
}
resp, err := qot.GetTicker(cli, req)
for _, t := range resp.TickerList {
    fmt.Printf("%s: %.2f x %d\n", t.Time, t.Price, t.Volume)
}
```

**Request fields:** `Security`, `MaxRetNum`

**Response fields:** `Security`, `Name`, `TickerList` (each Ticker has: `Time`, `Sequence`, `Dir`, `Price`, `Volume`, `Turnover`, `RecvTime`, `Type`, `TypeSign`, `Timestamp`)

#### `GetRT(c *Client, req *GetRTRequest) (*GetRTResponse, error)`

Retrieves分时数据 (time-share / intraday price chart data).

```go
req := &qot.GetRTRequest{Security: security}
resp, err := qot.GetRT(cli, req)
```

**Request fields:** `Security`

**Response fields:** `Security`, `Name`, `RTList` (each RT has: `Time`, `Price`, `LastClosePrice`, `AvgPrice`, `Volume`, `Turnover`)

#### `GetBroker(c *Client, req *GetBrokerRequest) (*GetBrokerResponse, error)`

Retrieves brokerage (经纪队列) data showing top brokers on each side of the order book.

```go
req := &qot.GetBrokerRequest{
    Security: security,
    Num:      10,
}
resp, err := qot.GetBroker(cli, req)
// resp.AskBrokerList:卖盘经纪, resp.BidBrokerList:买盘经纪
```

**Request fields:** `Security`, `Num`

**Response fields:** `Security`, `Name`, `AskBrokerList`, `BidBrokerList` (each Broker has: `ID`, `Name`, `Pos`, `Volume`)

---

### Market Information

#### `GetStaticInfo(c *Client, req *GetStaticInfoRequest) (*GetStaticInfoResponse, error)`

Retrieves static security information (listing date, lot size, etc.).

```go
req := &qot.GetStaticInfoRequest{
    SecurityList: []*qotcommon.Security{security},
}
resp, err := qot.GetStaticInfo(cli, req)
```

**Request fields:** `SecurityList`

#### `GetTradeDate(c *Client, req *GetTradeDateRequest) (*GetTradeDateResponse, error)`

Retrieves trading dates for a market.

```go
req := &qot.GetTradeDateRequest{
    Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
}
resp, err := qot.GetTradeDate(cli, req)
for _, td := range resp.TradeDateList {
    fmt.Printf("Trading date: %s\n", td)
}
```

**Request fields:** `Market`, `BeginTime`, `EndTime`

**Response fields:** `TradeDateList` (strings in `YYYY-MM-DD` format)

#### `GetMarketState(c *Client, req *GetMarketStateRequest) (*GetMarketStateResponse, error)`

Checks whether a market is open, closed, or in pre/post market state.

```go
req := &qot.GetMarketStateRequest{
    SecurityList: []*qotcommon.Security{security},
}
resp, err := qot.GetMarketState(cli, req)
```

#### `GetCapitalFlow(c *Client, req *GetCapitalFlowRequest) (*GetCapitalFlowResponse, error)`

Retrieves capital flow data showing money movement between different user categories.

```go
req := &qot.GetCapitalFlowRequest{
    Security: security,
}
resp, err := qot.GetCapitalFlow(cli, req)
```

#### `GetCapitalDistribution(c *Client, security *qotcommon.Security) (*GetCapitalDistributionResponse, error)`

Retrieves capital distribution by user category (large/medium/small orders).

```go
resp, err := qot.GetCapitalDistribution(cli, security)
```

---

### Stock Screening & Analysis

#### `StockFilter(c *Client, req *StockFilterRequest) (*StockFilterResponse, error)`

Screens securities based on technical and fundamental criteria.

```go
req := &qot.StockFilterRequest{
    Market:          int32(qotcommon.QotMarket_QotMarket_HK_Security),
    FilterDataList:  []qot.FilterData{},
    SortField:       int32(0),
    Ascend:          true,
    PageNum:         1,
}
resp, err := qot.StockFilter(cli, req)
```

#### `GetOptionChain(c *Client, req *GetOptionChainRequest) (*GetOptionChainResponse, error)`

Retrieves option chain data for an underlying security.

```go
req := &qot.GetOptionChainRequest{
    UnderlyingSecurity: security,
    Type:                int32(0), // 0=All, 1=Call, 2=Put
}
resp, err := qot.GetOptionChain(cli, req)
```

**Request fields:** `UnderlyingSecurity`, `Type`, `ExpiryDateBegin`, `ExpiryDateEnd`, `StrikePriceMin`, `StrikePriceMax`

#### `GetOptionExpirationDate(c *Client, req *GetOptionExpirationDateRequest) (*GetOptionExpirationDateResponse, error)`

Retrieves available expiration dates for options on a given underlying.

#### `GetWarrant(c *Client, req *GetWarrantRequest) (*GetWarrantResponse, error)`

Retrieves warrant (窝轮) data for a specified issuer.

#### `GetSuspend(c *Client, req *GetSuspendRequest) (*GetSuspendResponse, error)`

Retrieves securities currently suspended from trading.

#### `GetCodeChange(c *Client, req *GetCodeChangeRequest) (*GetCodeChangeResponse, error)`

Retrieves securities with code changes (e.g., name changes, delistings).

#### `GetFutureInfo(c *Client, req *GetFutureInfoRequest) (*GetFutureInfoResponse, error)`

Retrieves futures contract information.

#### `GetIpoList(c *Client, req *GetIpoListRequest) (*GetIpoListResponse, error)`

Retrieves the upcoming IPO calendar for a market.

#### `GetHoldingChangeList(c *Client, req *GetHoldingChangeListRequest) (*GetHoldingChangeListResponse, error)`

Retrieves shareholder holding changes (major shareholders).

---

### Securities & Plates

#### `GetPlateSet(c *Client, req *GetPlateSetRequest) (*GetPlateSetResponse, error)`

Retrieves available plate (板块) sets for a market.

```go
req := &qot.GetPlateSetRequest{
    Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
    PlateSetType: int32(0), // plate type enum
}
resp, err := qot.GetPlateSet(cli, req)
```

**Request fields:** `Market`, `PlateSetType`

**Response fields:** `PlateInfoList` (each has: `ID`, `Name`, `PlateType`)

#### `GetPlateSecurity(c *Client, req *GetPlateSecurityRequest) (*GetPlateSecurityResponse, error)`

Retrieves all securities belonging to a specific plate.

```go
req := &qot.GetPlateSecurityRequest{
    Plate: security, // use Security from PlateInfo as plate identifier
}
resp, err := qot.GetPlateSecurity(cli, req)
```

#### `GetOwnerPlate(c *Client, req *GetOwnerPlateRequest) (*GetOwnerPlateResponse, error)`

Retrieves all plates that a given security belongs to.

#### `GetReference(c *Client, req *GetReferenceRequest) (*GetReferenceResponse, error)`

Retrieves related securities (e.g., Warrants referencing an underlying, convertibles, etc.).

```go
req := &qot.GetReferenceRequest{
    Security:      security,
    ReferenceType: int32(0), // reference type enum
}
resp, err := qot.GetReference(cli, req)
```

#### `GetUserSecurity(c *Client, groupName string) (*GetUserSecurityResponse, error)`

Retrieves securities in a user's custom group.

```go
resp, err := qot.GetUserSecurity(cli, "My Watchlist")
```

#### `GetUserSecurityGroup(c *Client, req *GetUserSecurityGroupRequest) (*GetUserSecurityGroupResponse, error)`

Retrieves all user security groups.

#### `ModifyUserSecurity(c *Client, req *ModifyUserSecurityRequest) (*ModifyUserSecurityResponse, error)`

Adds or removes securities from a user security group.

```go
req := &qot.ModifyUserSecurityRequest{
    GroupName: "My Watchlist",
    SecurityList: []*qotcommon.Security{security},
    ModType: int32(1), // 1=Add, 2=Remove
}
```

---

### Price Alerts

#### `GetPriceReminder(c *Client, security *qotcommon.Security, market int32) (*GetPriceReminderResponse, error)`

Retrieves all price reminders set for a security.

#### `SetPriceReminder(c *Client, req *SetPriceReminderRequest) (*SetPriceReminderResponse, error)`

Creates, modifies, or deletes a price reminder.

```go
req := &qot.SetPriceReminderRequest{
    Security:  security,
    Key:       0, // 0=create, >0=modify/delete
    Type:      int32(1), // 1=Above, 2=Below, 3=Change%
    Value:     360.0,
    Note:      "Alert me at 360",
    ModType:   int32(1), // 1=Add, 2=Modify, 3=Delete
}
resp, err := qot.SetPriceReminder(cli, req)
```

---

### History Data

#### `RequestRehab(c *Client, req *RequestRehabRequest) (*RequestRehabResponse, error)`

Requests pre-split adjusted price data (复权数据) for historical K-lines.

---

### Subscriptions & Push

#### `Subscribe(c *Client, req *SubscribeRequest) (*SubscribeResponse, error)`

Subscribes to real-time data for one or more securities.

```go
req := &qot.SubscribeRequest{
    SecurityList:     []*qotcommon.Security{security},
    SubTypeList:     []qot.SubType{qot.SubType_Basic, qot.SubType_KL, qot.SubType_Ticker},
    IsSubOrUnSub:    true, // true=subscribe, false=unsubscribe
}
resp, err := qot.Subscribe(cli, req)
```

**Request fields:** `SecurityList`, `SubTypeList`, `IsSubOrUnSub`, `IsRegOrUnRegPush`

**SubType values:** `SubType_Basic` (实时行情), `SubType_KL` (K线), `SubType_Ticker` (逐笔), `SubType_OrderBook` (买卖盘), `SubType_Broker` (经纪队列), `SubType_RT` (分时)

#### `RegQotPush(c *Client, req *RegQotPushRequest) (*RegQotPushResponse, error)`

Registers or unregisters for push notifications on specific data types.

#### `GetSubInfo(c *Client) (*GetSubInfoResponse, error)`

Returns current subscription quota usage.

---

## Trading (trd)

Import: `"gitee.com/shing1211/futuapi4go/pkg/trd"`

All trading functions require an unlocked trading account. Pass the `AccID` and `TrdMarket` via a `Header` in each request.

### Account Management

#### `GetAccList(c *Client, trdCategory int32, needGeneralSecAccount bool) (*GetAccListResponse, error)`

Lists all trading accounts accessible to the current user.

```go
// Security account category
resp, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
// Futures account category: TrdCategory_Futures
// US options: TrdCategory_USOption
for _, acc := range resp.AccList {
    fmt.Printf("Account: %d, Type: %d\n", acc.AccID, acc.AccType)
}
```

**Response fields:** `AccList` (each Acc has: `TrdEnv`, `AccID`, `AccType`, `CardNum`, `AccStatus`)

#### `UnlockTrade(c *Client, req *UnlockTradeRequest) error`

Unlocks the trading account with password MD5 hash.

```go
req := &trd.UnlockTradeRequest{
    Unlock: true,
    PwdMD5: "md5hash_of_password",
}
if err := trd.UnlockTrade(cli, req); err != nil {
    log.Printf("Unlock failed: %v", err)
}
```

#### `GetFunds(c *Client, req *GetFundsRequest) (*GetFundsResponse, error)`

Retrieves account funds and buying power.

```go
req := &trd.GetFundsRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
}
resp, err := trd.GetFunds(cli, req)
fmt.Printf("Power: %.2f, Cash: %.2f\n", resp.Funds.Power, resp.Funds.Cash)
```

**Request fields:** `AccID`, `TrdMarket`

**Response fields:** `Funds` (has: `Power`, `TotalAssets`, `Cash`, `MarketVal`, `FrozenCash`, `DebtCash`, `Currency`, `AvailableFunds`)

#### `GetPositionList(c *Client, req *GetPositionListRequest) (*GetPositionListResponse, error)`

Retrieves current positions.

```go
req := &trd.GetPositionListRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
}
resp, err := trd.GetPositionList(cli, req)
for _, p := range resp.PositionList {
    fmt.Printf("%s: Qty=%.0f, Val=%.2f, PL=%.2f\n", p.Code, p.Qty, p.Val, p.PlVal)
}
```

**Request fields:** `AccID`, `TrdMarket`

**Response fields:** `PositionList` (each Position has: `Code`, `Name`, `Qty`, `CanSellQty`, `Price`, `CostPrice`, `Val`, `PlVal`, `PlRatio`)

---

### Order Management

#### `PlaceOrder(c *Client, req *PlaceOrderRequest) (*PlaceOrderResponse, error)`

Places a new buy or sell order.

```go
req := &trd.PlaceOrderRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
    Code:      "00700",
    TrdSide:   int32(trdcommon.TrdSide_TrdSide_Buy),
    OrderType: int32(trdcommon.OrderType_OrderType_Normal),
    Price:     350.00,
    Qty:       100.0,
}
resp, err := trd.PlaceOrder(cli, req)
fmt.Printf("OrderID: %d\n", resp.OrderID)

// TrdSide: TrdSide_Buy, TrdSide_Sell
// OrderType: OrderType_Normal, OrderType_Market, OrderType_Auction,
//            OrderType_AuctionLimit, OrderType_SpecialLimit
```

**Request fields:** `AccID`, `TrdMarket`, `Code`, `TrdSide`, `OrderType`, `Price`, `Qty`

**Response fields:** `OrderID`, `OrderIDEx`

#### `ModifyOrder(c *Client, req *ModifyOrderRequest) error`

Modifies or cancels an existing order.

```go
req := &trd.ModifyOrderRequest{
    AccID:         accID,
    TrdMarket:     int32(trdcommon.TrdMarket_TrdMarket_HK),
    OrderID:       orderID,
    ModifyOrderOp: int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Normal),
    Price:         360.00, // new price (0 to keep unchanged)
    Qty:           50.0,   // new quantity (0 to keep unchanged)
}
if err := trd.ModifyOrder(cli, req); err != nil {
    log.Printf("Modify failed: %v", err)
}

// ModifyOrderOp: ModifyOrderOp_Normal (修改), ModifyOrderOp_Cancel (撤销)
```

**Request fields:** `AccID`, `TrdMarket`, `OrderID`, `ModifyOrderOp`, `Price`, `Qty`

#### `GetOrderList(c *Client, req *GetOrderListRequest) (*GetOrderListResponse, error)`

Retrieves today's pending and recently filled orders.

```go
req := &trd.GetOrderListRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
}
resp, err := trd.GetOrderList(cli, req)
for _, o := range resp.OrderList {
    fmt.Printf("Order %d: %s %s @ %.2f, Qty=%.0f, Status=%d\n",
        o.OrderID, o.Code, o.TrdSide == 1 ? "BUY" : "SELL", o.Price, o.Qty, o.OrderStatus)
}
```

**Request fields:** `AccID`, `TrdMarket`

**Response fields:** `OrderList` (each Order has: `OrderID`, `Code`, `Name`, `TrdSide`, `OrderType`, `OrderStatus`, `Price`, `Qty`, `FillQty`, `CreateTime`, `UpdateTime`, `FillAvgPrice`)

#### `GetOrderFillList(c *Client, req *GetOrderFillListRequest) (*GetOrderFillListResponse, error)`

Retrieves today's order fills (成交记录).

```go
req := &trd.GetOrderFillListRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
}
resp, err := trd.GetOrderFillList(cli, req)
for _, f := range resp.OrderFillList {
    fmt.Printf("Fill %d: %s @ %.2f x %.0f\n", f.FillID, f.Code, f.Price, f.Qty)
}
```

**Request fields:** `AccID`, `TrdMarket`

**Response fields:** `OrderFillList` (each OrderFill has: `OrderID`, `FillID`, `Code`, `Name`, `TrdSide`, `Price`, `Qty`, `CreateTime`)

---

### Order History

#### `GetHistoryOrderList(c *Client, req *GetHistoryOrderListRequest) (*GetHistoryOrderListResponse, error)`

Retrieves historical orders within a date range.

```go
req := &trd.GetHistoryOrderListRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
    StartTime: "2024-01-01",
    EndTime:   "2024-12-31",
}
```

**Request fields:** `AccID`, `TrdMarket`, `StartTime`, `EndTime`, `FilterConditions`

**Response fields:** `OrderList`

#### `GetHistoryOrderFillList(c *Client, req *GetHistoryOrderFillListRequest) (*GetHistoryOrderFillListResponse, error)`

Retrieves historical order fills within a date range.

**Request fields:** `AccID`, `TrdMarket`, `StartTime`, `EndTime`, `FilterConditions`

---

### Advanced Trading

#### `GetOrderFee(c *Client, req *GetOrderFeeRequest) (*GetOrderFeeResponse, error)`

Calculates the fees for a potential order.

```go
req := &trd.GetOrderFeeRequest{
    AccID:         accID,
    TrdMarket:     int32(trdcommon.TrdMarket_TrdMarket_HK),
    OrderIDExList: []string{orderIDEx},
}
resp, err := trd.GetOrderFee(cli, req)
```

#### `GetMarginRatio(c *Client, req *GetMarginRatioRequest) (*GetMarginRatioResponse, error)`

Retrieves margin ratio information for short selling.

```go
req := &trd.GetMarginRatioRequest{
    AccID:        accID,
    TrdMarket:    int32(trdcommon.TrdMarket_TrdMarket_HK),
    SecurityList: []*qotcommon.Security{security},
}
resp, err := trd.GetMarginRatio(cli, req)
```

#### `GetMaxTrdQtys(c *Client, req *GetMaxTrdQtysRequest) (*GetMaxTrdQtysResponse, error)`

Calculates maximum tradable quantities for a given price.

```go
req := &trd.GetMaxTrdQtysRequest{
    AccID:     accID,
    TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
    SecurityList: []*qotcommon.Security{security},
    PriceList:    []float64{350.0},
}
resp, err := trd.GetMaxTrdQtys(cli, req)
```

#### `SubAccPush(c *Client, req *SubAccPushRequest) error`

Subscribes to account push notifications (order updates, fill updates).

#### `ReconfirmOrder(c *Client, req *ReconfirmOrderRequest) (*ReconfirmOrderResponse, error)`

Requests order re-confirmation for risky operations.

#### `GetFlowSummary(c *Client, req *GetFlowSummaryRequest) (*GetFlowSummaryResponse, error)`

Retrieves trading flow summary data.

---

## System (sys)

Import: `"gitee.com/shing1211/futuapi4go/pkg/sys"`

### `GetGlobalState(c *Client) (*GetGlobalStateResponse, error)`

Retrieves the global connection state from OpenD.

```go
state, err := sys.GetGlobalState(cli)
fmt.Printf("Server Ver: %d, Build: %d\n", state.ServerVer, state.ServerBuildNo)
fmt.Printf("Qot Logined: %v, Trd Logined: %v\n", state.QotLogined, state.TrdLogined)
```

**Response fields:** `ConnID`, `ServerVer`, `ServerBuildNo`, `Time`, `LocalTime`, `QotLogined`, `TrdLogined`, `QotSvrIpAddr`, `TrdSvrIpAddr`, `MarketHK`, `MarketUS`, `MarketSH`, `MarketSZ`

### `GetUserInfo(c *Client) (*GetUserInfoResponse, error)`

Retrieves user account information.

**Response fields:** `UserID`, `NickName`, `AvatarUrl`, `ApiLevel`, `IsNeedAgreeDisclaimer`

### `GetDelayStatistics(c *Client) (*GetDelayStatisticsResponse, error)`

Retrieves connection delay statistics.

### `Verification(c *Client, req *VerificationRequest) error`

Submits verification code for account operations.

---

## Push Notifications

Import: `"gitee.com/shing1211/futuapi4go/pkg/push"`

Register push handlers on the client to receive real-time updates:

```go
cli.RegisterHandler(push.ProtoID_Qot_UpdateBasicQot, func(header *push.Header, body []byte) {
    data, err := push.ParseUpdateBasicQot(body)
    if err != nil {
        log.Printf("Parse error: %v", err)
        return
    }
    fmt.Printf("Price update: %s -> %.2f\n", data.Security.GetCode(), data.CurPrice)
})

cli.RegisterHandler(push.ProtoID_Qot_UpdateOrderBook, func(header *push.Header, body []byte) {
    data, err := push.ParseUpdateOrderBook(body)
    // ...
})
```

### Qot Push Handlers

| ProtoID | Parse Function | Description |
|---------|---------------|-------------|
| 3101 | `ParseUpdateBasicQot` | Real-time price updates |
| 3102 | `ParseUpdateKL` | K-line updates |
| 3103 | `ParseUpdateOrderBook` | Order book updates |
| 3104 | `ParseUpdateTicker` | Tick-by-tick trades |
| 3105 | `ParseUpdateRT` | Intraday time-share updates |
| 3106 | `ParseUpdateBroker` | Broker queue updates |
| 3107 | `ParseUpdatePriceReminder` | Price alert triggers |

### Trd Push Handlers

| ProtoID | Parse Function | Description |
|---------|---------------|-------------|
| 7001 | `ParseUpdateOrder` | Order status updates |
| 7002 | `ParseUpdateOrderFill` | Order fill (成交) notifications |
| 7003 | `ParseTrdNotify` | Trading system notifications |

---

## Common Enums

### Markets (`qotcommon`)

```go
qotcommon.QotMarket_QotMarket_HK_Security  // 港股
qotcommon.QotMarket_QotMarket_US_Security  // 美股
qotcommon.QotMarket_QotMarket_CN_Security  // A股
qotcommon.QotMarket_QotMarket_CN_Future    // 期貨
qotcommon.QotMarket_QotMarket_HK_Future    // 港期
qotcommon.QotMarket_QotMarket_SG_Future    // 新加坡期貨
```

### Trading Markets (`trdcommon`)

```go
trdcommon.TrdMarket_TrdMarket_HK    // 香港
trdcommon.TrdMarket_TrdMarket_US    // 美國
trdcommon.TrdMarket_TrdMarket_CN   // A股
trdcommon.TrdMarket_TrdMarket_CN_Future // 期貨
trdcommon.TrdMarket_TrdMarket_HK_Future  // 港期
```

### Trading Sides (`trdcommon`)

```go
trdcommon.TrdSide_TrdSide_Buy  // 買入
trdcommon.TrdSide_TrdSide_Sell // 賣出
```

### Order Types (`trdcommon`)

```go
trdcommon.OrderType_OrderType_Normal      // 普通訂單
trdcommon.OrderType_OrderType_Market      // 市價單
trdcommon.OrderType_OrderType_Auction     // 競價單
trdcommon.OrderType_OrderType_AuctionLimit // 競價限價
trdcommon.OrderType_OrderType_SpecialLimit // 特殊限價
```

### K-Line Types (`qotcommon`)

```go
qotcommon.KLType_KLType_1Min   // 1分鐘
qotcommon.KLType_KLType_5Min   // 5分鐘
qotcommon.KLType_KLType_15Min  // 15分鐘
qotcommon.KLType_KLType_30Min  // 30分鐘
qotcommon.KLType_KLType_60Min  // 60分鐘
qotcommon.KLType_KLType_Day    // 日K
qotcommon.KLType_KLType_Week   // 週K
qotcommon.KLType_KLType_Month  // 月K
qotcommon.KLType_KLType_Year   // 年K
qotcommon.KLType_KLType_1Hour  // 1小時
qotcommon.KLType_KLType_3Hour  // 3小時
qotcommon.KLType_KLType_4Hour  // 4小時
qotcommon.KLType_KLType_Invalid // 無效
```

### Rehabilitation Types (`qotcommon`)

```go
qotcommon.RehabType_RehabType_None     // 不復權
qotcommon.RehabType_RehabType_Forward  // 前復權
qotcommon.RehabType_RehabType_Backward  // 後復權
```

### Return Codes

All API responses include a `RetType` field:

```go
common.RetType_RetType_Succeed  // 0, 成功
common.RetType_RetType_Failed   // -1, 失敗
common.RetType_RetType_Timeout   // -2, 處理超時
```

---

## Connection Pool

For high-throughput scenarios, use the connection pool:

```go
pool := futuapi.NewClientPool(futuapi.DefaultPoolConfig("127.0.0.1:11111"))
defer pool.Close()

cli, err := pool.Get(futuapi.PoolTypeMarketData)
if err != nil {
    log.Fatalf("Pool get failed: %v", err)
}
defer pool.Put(cli)

// Use client...
```

---

## Error Handling

All API functions return `(result, error)`. Always check the error:

```go
quotes, err := qot.GetBasicQot(cli, securities)
if err != nil {
    log.Printf("API failed: %v", err)
    return
}
```

Connection errors can be handled with automatic reconnection:

```go
cli := futuapi.New(
    futuapi.WithMaxRetries(3),
    futuapi.WithReconnectInterval(3*time.Second),
    futuapi.WithReconnectBackoff(1.5),
)
```

---

## Simulator

A local mock server is available at `cmd/simulator/`:

```go
srv := simulator.New("127.0.0.1:22222")
srv.RegisterDefaultHandlers()
srv.RegisterQotHandlers()
srv.RegisterTrdHandlers()
srv.RegisterPushHandlers()
srv.AddSecurity(int32(qotcommon.QotMarket_QotMarket_HK_Security), "00700")
if err := srv.Start(); err != nil {
    log.Fatalf("Server start failed: %v", err)
}
defer srv.Stop()

cli := futuapi.New()
if err := cli.Connect("127.0.0.1:22222"); err != nil {
    log.Fatalf("Connect failed: %v", err)
}
```

Run the simulator:
```bash
go run cmd/simulator/main.go
```
