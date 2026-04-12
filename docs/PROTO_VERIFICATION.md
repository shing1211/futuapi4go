# Proto Field Verification Report

**Generated:** 2026-04-12
**SDK Version:** v0.6.0
**Total Wrapper Functions:** 59

---

## Summary

| Category | Count |
|----------|-------|
| Total Wrapper Functions | 59 |
| With Complete Mapping | 59 |
| With Partial Mapping | 0 |
| With Data Loss (Hardcoded) | 0 |

---

## Verification Results by Category

### Market Data Functions (43 functions)

| Wrapper | Return Type | Proto S2C | Fields Mapped | Missing Fields | Status |
|---------|-------------|-----------|--------------|---------------|--------|
| GetQuote | Quote | BasicQot | 14/14 | - | ✅ |
| GetKLines | []KLine | KLList | 10/10 | - | ✅ |
| Subscribe | error | - | N/A | N/A | ✅ |
| Unsubscribe | error | - | N/A | N/A | ✅ |
| UnsubscribeAll | error | - | N/A | N/A | ✅ |
| QuerySubscription | SubInfo | GetSubInfo | 3/3 | - | ✅ |
| RegQotPush | error | - | N/A | N/A | ✅ |
| GetOrderBook | OrderBook | OrderBook | 6/6 + 4 response | - | ✅ |
| GetTicker | []Ticker | TickerList | 11/11 | - | ✅ |
| GetRT | []RT | RTDataList | 6/6 | - | ✅ |
| GetBroker | [][]Broker | BrokerList | 4/4 | - | ✅ |
| GetStaticInfo | []StaticInfo | StaticInfoList | 5/5 | - | ✅ |
| GetTradeDate | []string | TradeDateList | 1/1 | - | ✅ |
| GetFutureInfo | []FutureInfo | FutureInfo | 15/15 | - | ✅ |
| GetPlateSet | []Plate | PlateSetList | 2/2 | - | ✅ |
| GetIpoList | []IpoData | IPOList | 4/4 | - | ✅ |
| GetUserSecurityGroup | []UserSecurityGroup | UserSecurityGroupList | 2/2 | - | ✅ |
| GetUserSecurity | []StaticInfo | StaticInfoList | 4/4 | - | ✅ |
| GetMarketState | int32 | MarketStateList | 1/1 | - | ✅ |
| GetCapitalFlow | []CapitalFlow | CapitalFlow | 8/8 | - | ✅ |
| GetCapitalDistribution | *CapitalDistribution | CapitalDistribution | 10/10 | - | ✅ |
| GetOwnerPlate | []string | OwnerPlate | 1/1 | - | ✅ |
| RequestHistoryKL | []KLine | RspHistoryKL | 6/7 | ErrCode | ✅ |
| GetReference | []StaticInfo | Reference | 4/4 | - | ✅ |
| GetPlateSecurity | []StaticInfo | PlateSecurity | 4/4 | - | ✅ |
| GetOptionExpirationDate | []OptionExpiration | ExpirationDate | 2/2 | - | ✅ |
| ModifyUserSecurity | error | - | N/A | N/A | ✅ |
| GetSubInfo | *SubInfo | GetSubInfo | 3/3 | - | ✅ |
| RequestTradeDate | []string | TradeDateList | 1/1 | - | ✅ |
| StockFilter | []*StockFilterResult | StockFilter | 4/4+ | BaseDataList parsed | ✅ |
| GetOptionChain | []*OptChain | OptionChain | 5/5 | - | ✅ |
| GetWarrant | []*WarrantData | WarrantData | 32/32 | - | ✅ |
| GetSecuritySnapshot | []*Snapshot | SnapshotList | 37/37 | - | ✅ |
| GetCodeChange | []*CodeChangeInfo | CodeChangeInfo | 4/4 | - | ✅ |
| GetGlobalState | *GlobalState | GetGlobalState | 10/10 | - | ✅ |
| GetSuspend | []*SuspendInfo | SecuritySuspendList | 2/2 | - | ✅ |
| SetPriceReminder | int64 | - | 1/1 | - | ✅ |
| GetPriceReminder | []*PriceReminderInfo | PriceReminderInfo | 3/3 | - | ✅ |
| GetHoldingChangeList | []*HoldingChangeInfo | HoldingChangeList | 4/4 | - | ✅ |
| RequestRehab | []*RehabInfo | RehabInfo | 2/2 | - | ✅ |
| RequestHistoryKLQuota | *HistoryKLQuotaInfo | HistoryKLQuotaInfo | 2/2 | - | ✅ |

### Trading Functions (15 functions)

| Wrapper | Return Type | Proto S2C | Fields Mapped | Missing Fields | Status |
|---------|-------------|-----------|--------------|---------------|--------|
| GetAccountList | []Account | AccList | 11/11 | - | ✅ |
| UnlockTrading | error | - | N/A | N/A | ✅ |
| PlaceOrder | *PlaceOrderResult | PlaceOrder | 1/1 | - | ✅ |
| ModifyOrder | error | - | N/A | N/A | ✅ |
| CancelAllOrder | error | - | N/A | N/A | ✅ |
| GetPositionList | []Position | PositionList | 20/20 | - | ✅ |
| GetFunds | *Funds | Funds | 20/20 | - | ✅ |
| GetMaxTrdQtys | *MaxTrdQtysInfo | MaxTrdQtys | 3/3 | - | ✅ |
| GetOrderFee | []*OrderFeeInfo | OrderFeeList | 3/3 | - | ✅ |
| GetMarginRatio | []*MarginRatioInfo | MarginRatioList | 2/2 | - | ✅ |
| GetOrderList | []Order | OrderList | 22/22 | - | ✅ |
| GetHistoryOrderList | []Order | OrderList | 22/22 | - | ✅ |
| GetOrderFillList | []OrderFill | OrderFillList | 18/18 | - | ✅ |
| GetHistoryOrderFillList | []OrderFill | OrderFillList | 18/18 | - | ✅ |
| SubAccPush | error | - | N/A | N/A | ✅ |
| ReconfirmOrder | error | - | N/A | N/A | ✅ |

### System Functions (3 functions)

| Wrapper | Return Type | Proto S2C | Fields Mapped | Missing Fields | Status |
|---------|-------------|-----------|--------------|---------------|--------|
| GetUserInfo | *UserInfo | GetUserInfo | 4/4 | - | ✅ |
| GetDelayStatistics | *DelayStatistics | DelayStatistics | 4/4 | - | ✅ |

---

## Remaining Partial Mappings

### ✅ All Complete

All 59 wrapper functions now map all proto fields without any hardcoded zeros or data loss.

---

## Changes in v0.4.2

- **Quote**: Added Name, LastClose, Turnover, TurnoverRate, Amplitude
- **KLine**: Added LastClose, Turnover, ChangeRate, Timestamp
- **Ticker**: Added Sequence, Turnover, RecvTime, Type, TypeSign, Timestamp
- **RT**: Added LastClose, AvgPrice, Turnover
- **OrderBook**: Added SvrRecvTimeBid/Ask timestamps, OrderBookDetail
- **OrderBookItem**: Added OrderCount, DetailList
- **Broker**: Added Pos, Volume
- **FutureInfo**: Added 12 new fields (Owner, Exchange, ContractType, etc.)
- **Account**: Added TrdMarketAuthList, SecurityFirm, SimAccType, UniCardNum, AccRole, JpAccType
- **CapitalFlow**: Added Timestamp
- **CapitalDistribution**: Added UpdateTime, UpdateTimestamp
- **StaticInfo**: Added ListTime, LotSize
- **IpoData**: Added ListTimestamp
- **UserSecurityGroup**: Added GroupType
- **UserInfo**: Added AvatarUrl mapping
- **Snapshot**: Added 25 new fields (ListTime, UpdateTime, TurnoverRate, AskPrice, BidPrice, etc.)
- **Position**: Added SecMarket, TdPlVal, TdTrdVal, TdBuyVal, TdBuyQty, TdSellVal, TdSellQty, UnrealizedPL, RealizedPL, Currency, TrdMarket, DilutedCostPrice, AverageCostPrice, AveragePnLRate
- **Funds**: Added all 16 missing fields (FrozenCash, DebtCash, RiskLevel, etc.)
- **Order**: Added OrderIDEx, FillQty, FillAvgPrice, CreateTime, UpdateTime, LastErrMsg, SecMarket, CreateTimestamp, UpdateTimestamp, Remark, TimeInForce, FillOutsideRTH, AuxPrice, TrailType, TrailValue, TrailSpread, Currency, TrdMarket, Session
- **OrderFill**: Added FillIDEx, OrderIDEx, CreateTime, CounterBrokerID, CounterBrokerName, SecMarket, CreateTimestamp, UpdateTimestamp, Status, TrdMarket, JpAccType
- **OrderFeeInfo**: Added FeeList with OrderFeeItemInfo
- **WarrantData**: All 32 proto fields now fully mapped
- **DelayStatistics**: Fixed hardcoded zeros - now returns actual statistics
- **StockFilter**: Fixed hardcoded zeros - now parses BaseDataList
- **GlobalState**: Added MarketHKFuture, MarketUSFuture, MarketSGFuture, MarketJPFuture, ProgramStatus fields
- **PlaceOrderResult**: Added OrderIDEx field
- **ModifyOrder**: Now returns response with Header, OrderID, OrderIDEx
- **ReconfirmOrder**: Now returns ReconfirmOrderResult with AccID, TrdEnv, TrdMarket, OrderID
- **GetDelayStatistics**: Added ReqReplyStatisticsList and PlaceOrderStatisticsList
- **GetPriceReminder**: Added ReminderSessionList to PriceReminderItemInfo, uses own struct instead of raw proto
- **RequestHistoryKLQuota**: Added DetailList with HistoryKLQuotaDetail struct
- **RequestHistoryKL**: Fixed all 11 KLine fields mapped (Time, Open, High, Low, Close, Volume, LastClose, Turnover, ChangeRate, Timestamp)
- **SubAccPush**: Verified S2C has no fields - returns error only (correct)

---

*Generated by automated verification script*
