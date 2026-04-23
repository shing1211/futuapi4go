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

## High-Level Client Wrappers

Import: `"github.com/shing1211/futuapi4go/client"`

The `client` package provides high-level wrapper functions that simplify common workflows. They accept a `*client.Client` and return user-friendly types.

### Market Data Wrappers

#### `client.GetQuote(c *Client, market int32, code string) (*Quote, error)`

Retrieves real-time quote for a single security. Returns a `Quote` struct with all price/volume data.

```go
quote, err := client.GetQuote(c, client.Market_HK_Security, "00700")
```

**Response:** `Quote` struct with `Market`, `Code`, `Name`, `CurPrice`, `OpenPrice`, `HighPrice`, `LowPrice`, `LastClose`, `Volume`, `Turnover`, `TurnoverRate`, `Amplitude`, `UpdateTime`, `IsSuspend`.

#### `client.GetKLines(c *Client, market int32, code string, klType int32, num int) ([]KLine, error)`

Retrieves K-line (candlestick) data.

```go
klines, err := client.GetKLines(c, client.Market_HK_Security, "00700", client.KLType_Day, 100)
```

**Response:** `[]KLine` with `Time`, `Open`, `High`, `Low`, `Close`, `Volume`, `Turnover`, `LastClose`, `ChangeRate`, `Timestamp`.

#### `client.GetOrderBook(c *Client, market int32, code string, num int) (*OrderBook, error)`

Retrieves order book (买卖盘). Returns `OrderBook` with `BidList` and `AskList` slices.

#### `client.GetTicker(c *Client, market int32, code string, num int) ([]Ticker, error)`

Retrieves tick-by-tick trade data.

#### `client.GetRT(c *Client, market int32, code string) ([]RT, error)`

Retrieves 分时 (time-share / intraday) price data.

#### `client.GetBroker(c *Client, market int32, code string, num int) ([]Broker, []Broker, error)`

Returns `(askBrokers, bidBrokers, error)` — top brokers on each side of the order book.

#### `client.GetStaticInfo(c *Client, market int32, code string) ([]StaticInfo, error)`

Retrieves static security info (listing date, lot size).

#### `client.GetTradeDate(c *Client, market int32, startDate, endDate string) ([]string, error)`

Retrieves trading dates in `YYYY-MM-DD` format.

#### `client.GetFutureInfo(c *Client, code string) ([]FutureInfo, error)`

Retrieves futures contract info.

#### `client.GetPlateSet(c *Client, market int32) ([]Plate, error)`

Retrieves available plate (板块) sets.

#### `client.GetIpoList(c *Client, market int32) ([]IpoData, error)`

Retrieves IPO calendar.

#### `client.GetMarketState(c *Client, market int32, code string) (int32, error)`

Returns market state (open/closed/pre-market).

#### `client.GetCapitalFlow(c *Client, market int32, code string) ([]CapitalFlow, error)`

Retrieves capital flow data.

#### `client.GetCapitalDistribution(c *Client, market int32, code string) (*CapitalDistribution, error)`

Retrieves capital distribution by order size.

#### `client.GetOwnerPlate(c *Client, market int32, code string) ([]string, error)`

Returns all plates that a security belongs to.

#### `client.GetReference(c *Client, market int32, code string, refType int32) ([]StaticInfo, error)`

Retrieves related securities (warrants, convertibles, etc.).

#### `client.GetPlateSecurity(c *Client, market int32, plateCode string) ([]StaticInfo, error)`

Retrieves all securities in a specific plate.

#### `client.GetOptionExpirationDate(c *Client, market int32, code string) ([]OptionExpiration, error)`

Returns available expiration dates for options.

#### `client.GetOptionChain(c *Client, market int32, code string, ...) ([]*OptChain, error)`

Retrieves full option chain data.

#### `client.GetWarrant(c *Client, market int32, code string, ...) ([]*WarrantData, error)`

Retrieves warrant (窝轮) data.

#### `client.GetSecuritySnapshot(c *Client, securities []*qotcommon.Security) ([]*Snapshot, error)`

Retrieves snapshot data for multiple securities in one call.

#### `client.GetCodeChange(c *Client, securities []*qotcommon.Security) ([]*CodeChangeInfo, error)`

Retrieves securities with code changes (name, delisting, etc.).

#### `client.GetSuspend(c *Client, securities []*qotcommon.Security, beginTime, endTime string) ([]*SuspendInfo, error)`

Retrieves suspended securities.

### Subscription Wrappers

#### `client.Subscribe(c *Client, market int32, code string, subTypes []int32) error`

Subscribes to real-time data. `subTypes` values: `client.SubType_Basic`, `client.SubType_KL`, `client.SubType_Ticker`, `client.SubType_OrderBook`, `client.SubType_Broker`, `client.SubType_RT`.

```go
err := client.Subscribe(c, client.Market_HK_Security, "00700", []int32{client.SubType_Basic, client.SubType_KL})
```

#### `client.Unsubscribe(c *Client, market int32, code string, subTypes []int32) error`

Unsubscribes from real-time data.

#### `client.UnsubscribeAll(c *Client) error`

Unsubscribes from all real-time data.

#### `client.QuerySubscription(c *Client) (*qot.GetSubInfoResponse, error)`

Returns current subscription quota usage and subscribed securities.

#### `client.RegQotPush(c *Client, market int32, code string, ...) error`

Registers for push notifications on specific data types.

### Trading Wrappers

#### `client.GetAccountList(c *Client) ([]Account, error)`

Lists all trading accounts. Returns `[]Account` with `AccID`, `AccType`, `TrdEnv`, etc.

```go
accs, err := client.GetAccountList(c)
accID := accs[0].AccID
```

#### `client.UnlockTrading(c *Client, pwdMD5 string) error`

Unlocks trading with MD5-hashed password.

```go
err := client.UnlockTrading(c, "md5hash")
```

#### `client.PlaceOrder(c *Client, accID uint64, market int32, code string, side, orderType int32, price float64, qty float64) (*PlaceOrderResult, error)`

Places a buy or sell order. Returns `*PlaceOrderResult` with `OrderID` and `OrderIDEx`.

```go
result, err := client.PlaceOrder(c, accID, client.Market_HK_Security, "00700",
    client.Side_Buy, client.OrderType_Normal, 350.00, 100)
```

#### `client.ModifyOrder(c *Client, accID uint64, market int32, orderID uint64, modifyOp int32, price float64, qty float64) (*trd.ModifyOrderResponse, error)`

Modifies or cancels an existing order. Returns `*trd.ModifyOrderResponse` with `Header`, `OrderID`, `OrderIDEx`.

```go
resp, err := client.ModifyOrder(c, accID, client.Market_HK_Security, orderID,
    client.ModifyOp_Normal, 360.00, 0)
// Use resp.OrderID, resp.OrderIDEx on success
```

#### `client.CancelAllOrder(c *Client, accID uint64, market int32, trdEnv int32) error`

Cancels all pending orders for the account.

#### `client.GetPositionList(c *Client, accID uint64) ([]Position, error)`

Retrieves current positions.

#### `client.GetAccountInfo(c *Client, accID uint64, market int32) (*Funds, error)`

Retrieves full account information including multi-currency cash (per-currency cash, available balance, net cash power) and per-market assets. Maps to Python's `accinfo_query`. This is the recommended way to get account data.

```go
funds, err := client.GetAccountInfo(c, accID, constant.TrdMarket_HK)
// Access per-currency cash:
for _, ci := range funds.CashInfoList {
    fmt.Printf("Currency=%d, Cash=%.2f, Available=%.2f\n",
        ci.Currency, ci.Cash, ci.AvailableBalance)
}
// Access per-market assets:
for _, mi := range funds.MarketInfoList {
    fmt.Printf("Market=%d, Assets=%.2f\n", mi.TrdMarket, mi.Assets)
}
```

**Response:** `Funds` struct with all base fields plus:
- `CashInfoList []AccCashInfo` — per-currency: `Currency`, `Cash`, `AvailableBalance`, `NetCashPower`
- `MarketInfoList []AccMarketInfo` — per-market: `TrdMarket`, `Assets`
- Plus all base fields: `Power`, `TotalAssets`, `Cash`, `MarketVal`, `FrozenCash`, `DebtCash`, `AvlWithdrawalCash`, `Currency`, `AvailableFunds`, `UnrealizedPL`, `RealizedPL`, `RiskLevel`, `InitialMargin`, `MaintenanceMargin`, `MaxPowerShort`, `NetCashPower`, `LongMv`, `ShortMv`, `PendingAsset`, `MaxWithdrawal`, `RiskStatus`, `MarginCallMargin`, `IsPDT`, `PDTSeq`, `BeginningDTBP`, `RemainingDTBP`, `DtCallAmount`, `DtStatus`

#### `client.GetFunds(c *Client, accID uint64) (*Funds, error)`

Retrieves account funds. Auto-selects the first available account and market. Internally calls `GetAccountInfo`.

#### `client.GetFlowSummary(c *Client, accID uint64, market int32, clearingDate string, direction int32) ([]*FlowSummaryInfo, error)`

Retrieves account cash flow entries (清算资金流水). Maps to Python's `get_acc_cash_flow`.
- `clearingDate`: clearing date in "YYYY-MM-DD" format, empty means today.
- `direction`: 0=none, 1=in, 2=out. Use `constant.CashFlowDirection_In` etc.

```go
flows, err := client.GetFlowSummary(c, accID, constant.TrdMarket_HK, "2026-04-23", 0)
for _, f := range flows {
    fmt.Printf("ID=%d Date=%s Type=%s Dir=%d Amount=%.2f\n",
        f.CashFlowID, f.ClearingDate, f.CashFlowType, f.CashFlowDirection, f.CashFlowAmount)
}
```

**Response:** `[]*FlowSummaryInfo` with `CashFlowID`, `ClearingDate`, `SettlementDate`, `Currency`, `CashFlowType`, `CashFlowDirection`, `CashFlowAmount`, `CashFlowRemark`

#### `client.GetAccTradingInfo(c *Client, accID uint64, market int32, code string, orderType int32, price float64) (*AccTradingInfo, error)`

Retrieves maximum tradable quantities and margin requirements for a security. Maps to Python's `acctradinginfo_query`.

```go
info, err := client.GetAccTradingInfo(c, accID, constant.Market_HK, "00700",
    constant.OrderType_Normal, 350.00)
fmt.Printf("Max Cash Buy: %.0f, Max Sell: %.0f\n",
    info.MaxCashBuy, info.MaxPositionSell)
```

**Response:** `AccTradingInfo` with `MaxCashBuy`, `MaxCashAndMarginBuy`, `MaxPositionSell`, `MaxSellShort`, `MaxBuyBack`, `LongRequiredIM`, `ShortRequiredIM`

#### `client.GetOrderList(c *Client, accID uint64) ([]Order, error)`

Retrieves today's orders.

#### `client.GetHistoryOrderList(c *Client, accID uint64, market int32, startDate, endDate string) ([]Order, error)`

Retrieves historical orders within a date range.

#### `client.GetOrderFillList(c *Client, accID uint64) ([]OrderFill, error)`

Retrieves today's order fills.

#### `client.GetHistoryOrderFillList(c *Client, accID uint64, market int32) ([]OrderFill, error)`

Retrieves historical order fills.

#### `client.GetMaxTrdQtys(c *Client, accID uint64, market int32, code string, orderType int32, price float64) (*MaxTrdQtysInfo, error)`

Calculates maximum tradable quantities.

#### `client.GetOrderFee(c *Client, accID uint64, market int32, orderIDExList []string) ([]*OrderFeeInfo, error)`

Retrieves fee information for orders.

#### `client.GetMarginRatio(c *Client, accID uint64, market int32, securities []*qotcommon.Security) ([]*MarginRatioInfo, error)`

Retrieves margin ratio information.

#### `client.SubAccPush(c *Client, accIDList []uint64) error`

Subscribes to account push notifications.

#### `client.ReconfirmOrder(c *Client, accID uint64, market int32, orderID uint64, reason int32) (*ReconfirmOrderResult, error)`

Requests order re-confirmation for risky operations.

### User Security Wrappers

#### `client.GetUserSecurityGroup(c *Client) ([]UserSecurityGroup, error)`

Returns all user security groups.

#### `client.GetUserSecurity(c *Client, groupName string) ([]StaticInfo, error)`

Returns securities in a specific group.

#### `client.ModifyUserSecurity(c *Client, groupName string, op int32, market int32, codes []string) error`

Adds or removes securities from a group.

### Price Alerts

#### `client.SetPriceReminder(c *Client, market int32, code string, op, reminderType, freq int32, value float64, note string) (int64, error)`

Creates, modifies, or deletes a price reminder. Returns the reminder key.

#### `client.GetPriceReminder(c *Client, market int32, code string) ([]*PriceReminderInfo, error)`

Returns all price reminders for a security.

### Historical Data Wrappers

#### `client.RequestHistoryKL(c *Client, market int32, code string, klType int32, startDate, endDate string) ([]KLine, error)`

Requests historical K-line data.

#### `client.RequestHistoryKLQuota(c *Client) (*HistoryKLQuotaInfo, error)`

Returns historical K-line quota usage.

#### `client.RequestTradeDate(c *Client, market int32, startDate, endDate string, code string) ([]string, error)`

Requests trading dates.

#### `client.RequestRehab(c *Client, market int32, code string) ([]*RehabInfo, error)`

Requests pre-split adjusted price data.

### Analysis Wrappers

#### `client.StockFilter(c *Client, market int32, begin, num int32) ([]*StockFilterResult, error)`

Screens securities based on criteria.

### System Wrappers

#### `client.GetGlobalState(c *Client) (*GlobalState, error)`

Returns OpenD connection state and market statuses.

#### `client.GetUserInfo(c *Client) (*UserInfo, error)`

Returns user account information.

#### `client.GetDelayStatistics(c *Client) (*DelayStatistics, error)`

Returns connection delay statistics.

#### `client.GetSubInfo(c *Client) (*SubInfo, error)`

Returns subscription info (simplified wrapper around QuerySubscription).

### Portfolio Wrappers

#### `client.GetHoldingChangeList(c *Client, market int32, code string, holderCategory int32, beginTime, endTime string) ([]*HoldingChangeInfo, error)`

Retrieves shareholder holding changes.

### Constants

All constants are package-level variables on `client`:

| Category | Constants |
|----------|-----------|
| Markets | `Market_HK_Security`, `Market_US_Security`, `Market_CNSH_Security`, `Market_CNSZ_Security`, `Market_HK_Future`, `Market_CN_Future`, `Market_SG_Future`, `Market_JP_Future` |
| Sides | `Side_Buy`, `Side_Sell` |
| Order Types | `OrderType_Normal`, `OrderType_Market`, `OrderType_Stop`, `OrderType_Auction`, `OrderType_AuctionLimit`, `OrderType_SpecialLimit` |
| K-Line Types | `KLType_1Min` through `KLType_Month` |
| Subscription Types | `SubType_Basic`, `SubType_KL`, `SubType_Ticker`, `SubType_OrderBook`, `SubType_Broker`, `SubType_RT` |
| Modify Order Ops | `ModifyOp_Normal`, `ModifyOp_Cancel` |

---

## Internal Client

The following functions and types are in the `internal/client` package (`futuapi`).

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

### `cli.GetLoginUserID() uint64`

Returns the Futu/NiuNiu user ID that logged into OpenD. Zero if not logged in.

### `cli.IsEncrypt() bool`

Returns true if the connection uses AES encryption after the handshake.

### `cli.CanSendProto(protoID uint32) bool`

Reports whether a request for the given proto ID can be sent based on current connection state. InitConnect can be sent when connected; all other protos require the connection to be fully ready (serverVer > 0).

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

Import: `"github.com/shing1211/futuapi4go/pkg/qot"`

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

Import: `"github.com/shing1211/futuapi4go/pkg/trd"`

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

#### `ModifyOrder(c *Client, req *ModifyOrderRequest) (*ModifyOrderResponse, error)`

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
resp, err := trd.ModifyOrder(cli, req)
if err != nil {
    log.Printf("Modify failed: %v", err)
} else {
    fmt.Printf("Order modified: %d / %s\n", resp.OrderID, resp.OrderIDEx)
}

// ModifyOrderOp: ModifyOrderOp_Normal (修改), ModifyOrderOp_Cancel (撤销)
```

**Request fields:** `AccID`, `TrdMarket`, `OrderID`, `ModifyOrderOp`, `Price`, `Qty`

**Response fields:** `Header`, `OrderID`, `OrderIDEx`

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

Import: `"github.com/shing1211/futuapi4go/pkg/sys"`

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

Import: `"github.com/shing1211/futuapi4go/pkg/push"`

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

| Constant | ProtoID | Parse Function | Description |
|----------|---------|---------------|-------------|
| `ProtoID_Qot_UpdateBasicQot` | 3005 | `ParseUpdateBasicQot` | Real-time price updates |
| `ProtoID_Qot_UpdateKL` | 3007 | `ParseUpdateKL` | K-line updates |
| `ProtoID_Qot_UpdateOrderBook` | 3013 | `ParseUpdateOrderBook` | Order book updates |
| `ProtoID_Qot_UpdateTicker` | 3011 | `ParseUpdateTicker` | Tick-by-tick trades |
| `ProtoID_Qot_UpdateRT` | 3009 | `ParseUpdateRT` | Intraday time-share updates |
| `ProtoID_Qot_UpdateBroker` | 3015 | `ParseUpdateBroker` | Broker queue updates |
| `ProtoID_Qot_PushPriceReminder` | 3107 | `ParseUpdatePriceReminder` | Price alert triggers |

### Trd Push Handlers

| Constant | ProtoID | Parse Function | Description |
|----------|---------|---------------|-------------|
| `ProtoID_Trd_UpdateOrder` | 2208 | `ParseUpdateOrder` | Order status updates |
| `ProtoID_Trd_UpdateOrderFill` | 2218 | `ParseUpdateOrderFill` | Order fill (成交) notifications |
| `ProtoID_Trd_Notify` | 2207 | `ParseTrdNotify` | Trading system notifications |

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


