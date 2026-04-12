# Changelog

All notable changes to this project will be documented in this file.

## [0.6.0] - 2026-04-12

### Added

#### Push Notification Handler API
The public `client.Client` wrapper had zero push handler support. Added:

- `Client.SetPushHandler(protoID, handler)` — register per-protoID push handler
- `PushQuote`, `PushKLine`, `PushOrderBook`, `PushTicker` types
- `ParsePushQuote()`, `ParsePushKLine()`, `ParsePushOrderBook()`, `ParsePushTicker()`
- `OBItem` type for order book push data
- ProtoID constants re-exported for convenience (3005, 3007, 3011, 3013, 3009, 3015)

#### 100% Proto Field Coverage
Complete proto field audit across all 59 wrapper functions — zero data loss, no partial mappings.

- `GlobalState`: Added `MarketHKFuture`, `MarketUSFuture`, `MarketSGFuture`, `MarketJPFuture`, `ProgramStatus` fields
- `PlaceOrderResult`: Added `OrderIDEx` field
- `ModifyOrder`: Now returns `*ModifyOrderResponse` with `Header`, `OrderID`, `OrderIDEx` (previously returned `error` only)
- `ReconfirmOrder`: Now returns `*ReconfirmOrderResult` with `AccID`, `TrdEnv`, `TrdMarket`, `OrderID`
- `GetDelayStatistics`: Added `ReqReplyList` and `PlaceOrderList` (was only mapping QotPushStatisticsList)
- `GetPriceReminder`: Added `ReminderSessionList` to `PriceReminderItemInfo`, switched to own struct types instead of raw proto
- `RequestHistoryKLQuota`: Added `DetailList` with `HistoryKLQuotaDetail` struct
- `RequestHistoryKL`: Fixed all 11 KLine fields mapped; fixed manual pointer unwrapping
- `RequestHistoryKL`: Added automatic pagination via `NextReqKey` — fetches all available K-lines automatically
- `RequestHistoryKLWithLimit`: New function with configurable page size
- `DelayStatistics`: Fixed hardcoded zeros — now returns actual statistics
- `StockFilter`: Fixed hardcoded zeros — now parses BaseDataList (`CurPrice`, `ChangeRate`, `Volume`)
- `FutureInfo`: Removed non-existent `InstType` field
- `GetSubInfo`: Fixed SubTypes extraction from ConnSubInfo list

#### All Structs Now Fully Mapped (59/59)
- `Quote`: Added `Name`, `LastClose`, `Turnover`, `TurnoverRate`, `Amplitude`
- `KLine`: Added `LastClose`, `Turnover`, `ChangeRate`, `Timestamp`
- `Ticker`: Added `Sequence`, `Turnover`, `RecvTime`, `Type`, `TypeSign`, `Timestamp`
- `RT`: Added `LastClose`, `AvgPrice`, `Turnover`
- `OrderBook`: Added `SvrRecvTimeBid/Ask` timestamps
- `OrderBookItem`: Added `OrderCount`, `DetailList`, `OrderBookDetail` struct
- `Broker`: Added `Pos`, `Volume`
- `FutureInfo`: Added 12 fields (`Owner`, `Exchange`, `ContractType`, `ContractSize`, `MinVar`, `QuoteUnit`, etc.)
- `Account`: Added `TrdMarketAuthList`, `SecurityFirm`, `SimAccType`, `UniCardNum`, `AccRole`, `JpAccType`
- `CapitalFlow`: Added `Timestamp`
- `CapitalDistribution`: Added `UpdateTime`, `UpdateTimestamp`
- `StaticInfo`: Added `ListTime`, `LotSize`
- `IpoData`: Added `ListTimestamp`
- `UserSecurityGroup`: Added `GroupType`
- `UserInfo`: Added `AvatarUrl` mapping
- `Snapshot`: Added 25 fields (`ListTime`, `UpdateTime`, `TurnoverRate`, `AskPrice`, `BidPrice`, `EnableMargin`, `ShortSellRate`, `Amplitude`, `52W High/Low`, etc.)
- `Position`: Added 15 fields (`SecMarket`, `TdPlVal`, `TdTrdVal`, `TdBuyVal`, `TdBuyQty`, `TdSellVal`, `TdSellQty`, `UnrealizedPL`, `RealizedPL`, `Currency`, `TrdMarket`, `DilutedCostPrice`, `AverageCostPrice`, `AveragePnLRate`)
- `Funds`: Added 16 fields (`FrozenCash`, `DebtCash`, `AvlWithdrawalCash`, `RiskLevel`, `InitialMargin`, `MaintenanceMargin`, `MaxPowerShort`, `NetCashPower`, `LongMv`, `ShortMv`, `PendingAsset`, `MaxWithdrawal`, `RiskStatus`, `MarginCallMargin`, `IsPDT`, `PDTSeq`)
- `Order`: Added 15 fields (`OrderIDEx`, `FillQty`, `FillAvgPrice`, `CreateTime`, `UpdateTime`, `LastErrMsg`, `SecMarket`, `CreateTimestamp`, `UpdateTimestamp`, `Remark`, `TimeInForce`, `FillOutsideRTH`, `AuxPrice`, `TrailType`, `TrailValue`, `TrailSpread`, `Currency`, `TrdMarket`, `Session`)
- `OrderFill`: Added all 11 missing fields (`FillIDEx`, `OrderIDEx`, `CreateTime`, `CounterBrokerID`, `CounterBrokerName`, `SecMarket`, `CreateTimestamp`, `UpdateTimestamp`, `Status`, `TrdMarket`, `JpAccType`)
- `OrderFeeInfo`: Added `FeeList` with `OrderFeeItemInfo`
- `WarrantData`: All 32 proto fields now fully mapped

#### Internal Package Updates
- `pkg/sys/system.go`: `GetGlobalStateResponse` — added market future fields + `ProgramStatus`
- `pkg/qot/quote.go`: `RequestHistoryKLResponse` — use internal `KLine` type, map all fields; `PriceReminderItemInfo` — add `ReminderSessionList`
- `pkg/trd/trade.go`: `PlaceOrderResponse` — add `OrderIDEx`; `ModifyOrder` — add `ModifyOrderResponse` struct; `GetAccList` — add all 6 missing fields; `GetHistoryOrderFillList` — add `JpAccType`

#### Example Fixes
- `04_push_subscriptions`: Fixed `fmt` format specifiers for proto type changes
- `03_trading_operations`: Updated `ModifyOrder` call for new return signature
- `trd_modify_order`: Updated `ModifyOrder` call for new return signature

### Documentation
- `PROTO_VERIFICATION.md`: Comprehensive field-by-field report, 59/59 wrappers verified

## [0.4.1] - 2026-04-08

### Fixed

#### Protobuf Wrapper Layer Compatibility
- Fixed wrapper structs missing fields that caused example compilation failures
- `BasicQot` in `pkg/qot/quote.go` expanded with missing fields: `IsSuspended`, `UpdateTime`, `LastClosePrice`, `TurnoverRate`, `Amplitude`
- `ModifyOrderRequest` in `pkg/trd/trade.go` added `ModifyOrderOp` field

#### Example Code Fixes
- **qot_get_order_book**: Fixed `ob.OrederCount` → `ob.OrderCount` (proto has typo)
- **qot_get_trade_date**: Fixed `td.TradeDate` → `td.GetTime()` (field is `time`)
- **trd_unlock_trade**: Added `Unlock: true` field, changed `PWD` → `PwdMD5`
- **trd_place_order**: Removed non-existent `PriceType` field
- **trd_get_order_list**: Fixed `OrderState_*` → `OrderStatus_*`, `DealtQty` → `FillQty`
- **trd_modify_order**: Fixed `ModifyType` → `ModifyOrderOp`, `ModifyOrderType_*` → `ModifyOrderOp_*`
- **trd_get_order_list**: Fixed `OrderState_*` → `OrderStatus_*`, `DealtQty` → `FillQty`
- **02_market_data_advanced**: Fixed `basic.GetCode()` → `basic.GetSecurity().GetCode()`, fixed `GetFutureInfoRequest` fields, fixed capital flow field access, fixed `IpoPrice` field
- **03_trading_operations**: Fixed `UnlockTradeRequest`, `OrderStatus` enums, `ModifyOrderOp`, `OrderFillList`, `GetOrderFeeResponse`, `GetMarginRatioRequest`, `GetMaxTrdQtysResponse`, `GetHistoryOrderListRequest`, `TrdSide` pointer type
- **04_push_subscriptions**: Fixed `SubInfoList` → `ConnSubInfoList` with proper nesting, fixed `RegQotPushRequest` struct
- **05_comprehensive_demo**: Fixed capital flow fields, `GetMaxTrdQtysResponse`, `RegQotPush` request type
- **algo_breakout_trading**: Fixed variable scope issue for `stopLoss`/`takeProfit`
- **qot_stock_filter**: Fixed `BaseData.FieldName` pointer type, `AllCount` type conversion
- **qot_get_capital_flow**: Fixed capital flow field access

#### Test Fixes
- **test/integration**: Removed unused `fmt` import

#### Build/Vet Fixes
- Fixed `fmt.Println` → `fmt.Print` for strings containing printf directives in example code comments
- Fixed `fmt.Printf` format specifier mismatches (wrong argument types for `*string`, `*float64` fields)

### Verified

- **20/20 example compile tests pass** against live OpenD with simulated account
- All `go build ./...` succeeds
- All `go test ./...` passes (unit + integration + example compilation tests)

## [0.4.0] - 2026-04-08

### Added

#### Push Notification Support
- Push notification handler with serial matching
- Prevents push notifications from being consumed as request responses
- Support for all Qot and Trd push types

#### Metrics & Instrumentation
- Client metrics collection for API calls
- Latency tracking for request/response cycles
- Success/failure rate monitoring
- Connection pool metrics

#### Health Check
- Health check endpoint for client pool
- Periodic connectivity monitoring
- Auto-reconnection on health check failure

#### Version Information
- GetVersionInfo API implementation
- SDK version reporting
- OpenD version compatibility checking

#### Release Checklist
- Production readiness checklist
- Code quality gates
- Testing requirements
- Documentation requirements

#### Options Trading APIs
- GetOptionChain (2304) - Get option chain data
- GetOptionExpirationDate (2305) - Get option expiration dates

### Fixed
- Protocol header validation improvements
- Error handling for edge cases
- Connection state management fixes

### Testing
- 64 tests passing across 5 packages
- Unit tests for core client functionality
- Integration tests with OpenD simulator
- Concurrent access and race condition tests

## [0.3.0] - 2026-04-07

### Added

#### OpenD Simulator
- TCP server core (46-byte protocol header, LittleEndian)
- System API handlers (4): InitConnect, KeepAlive, GetGlobalState, GetUserInfo
- Qot market data handlers (42): Coverage for all Qot APIs
- Trd trading handlers (13): Coverage for all trading APIs
- Simulator example program (examples/simulator/main.go)

### Fixed

#### SDK Bug Fixes
- qot/quote.go: Subscribe - Added missing retType error check
- qot/quote.go: ModifyUserSecurity - Added missing retType error check
- qot/quote.go: RegQotPush - Added missing retType error check

### Documentation
- Updated IMPLEMENTATION.md with simulator stats
- Updated SIMULATOR.md with complete implementation status
- Updated README.md with project status

## [0.2.0] - 2026-04-07

### Added

#### Qot - Market Data API (29 APIs)
- GetBasicQot (2101) - Get real-time quotes
- GetKL (2102) - Get real-time K-lines
- GetOrderBook (2106) - Get order book
- GetTicker (2107) - Get tick-by-tick trades
- GetRT (2108) - Get real-time minute data
- GetSecuritySnapshot (2110) - Get security snapshot
- GetBroker (2111) - Get broker queue
- GetStaticInfo (2201) - Get security static info
- GetPlateSet (2202) - Get plate/sector set
- GetPlateSecurity (2203) - Get securities in plate
- GetOwnerPlate (2204) - Get owner plates
- GetReference (2205) - Get reference data
- GetTradeDate (2206) - Get trading dates
- RequestTradeDate (2207) - Request trading dates
- GetMarketState (2208) - Get market state
- GetSuspend (2209) - Get suspension info
- GetCodeChange (2210) - Get code change info
- GetFutureInfo (2211) - Get futures info
- GetIpoList (2212) - Get IPO list
- GetHoldingChangeList (2213) - Get holding change list
- RequestRehab (2214) - Request rehabilitation data
- GetCapitalFlow (2301) - Get capital flow
- GetCapitalDistribution (2302) - Get capital distribution
- StockFilter (2303) - Stock screening
- GetOptionChain (2304) - Get option chain
- GetOptionExpirationDate (2305) - Get option expiration dates
- GetWarrant (2306) - Get warrant info
- GetUserSecurity (2401) - Get user watchlist
- GetUserSecurityGroup (2402) - Get user watchlist groups
- ModifyUserSecurity (2403) - Modify user watchlist
- GetPriceReminder (2404) - Get price reminders
- SetPriceReminder (2405) - Set price reminders
- Subscribe (3001) - Subscribe to real-time data
- GetSubInfo (3002) - Get subscription info
- RegQotPush (3003) - Register for quote push
- RequestHistoryKLQuota (3104) - Get historical K-line quota usage
- RequestHistoryKL (2104) - Request historical K-lines (async)

#### Qot - Push Notifications (7 handlers)
- Qot_UpdateBasicQot (3101) - Real-time quote push
- Qot_UpdateKL (3102) - K-line push
- Qot_UpdateOrderBook (3103) - Order book push
- Qot_UpdateTicker (3104) - Tick-by-tick push
- Qot_UpdateRT (3105) - Minute data push
- Qot_UpdateBroker (3106) - Broker queue push
- Qot_UpdatePriceReminder (3107) - Price reminder push

#### Trd - Trading API (14 APIs)
- GetAccList (4001) - Get account list
- UnlockTrade (4002) - Unlock trading password
- GetFunds (4003) - Get account funds
- GetOrderFee (4004) - Get order fees
- GetMarginRatio (4005) - Get margin ratio
- GetMaxTrdQtys (4006) - Get max trade quantities
- PlaceOrder (5001) - Place order
- ModifyOrder (5002) - Modify order
- GetOrderList (5003) - Query order list
- GetHistoryOrderList (5004) - Query historical orders
- GetOrderFillList (5005) - Query fill list
- GetHistoryOrderFillList (5006) - Query historical fills
- GetPositionList (6001) - Get position list
- SubAccPush (7005) - Account push subscription
- ReconfirmOrder (7004) - Order confirmation

#### Trd - Push Notifications (3 handlers)
- Trd_UpdateOrder (7001) - Order status push
- Trd_UpdateOrderFill (7002) - Fill push
- Trd_Notify (7003) - Trade notification push

#### System - System API (4 APIs)
- GetGlobalState (1004) - Get global state
- GetUserInfo (1005) - Get user info
- GetDelayStatistics (1006) - Get delay statistics
- Verification (8001) - Verification interface

#### System - Push Notifications (1 handler)
- Notify (1003) - System notification push

### Updated
- Protobuf definitions updated to v10.2.6208 (74 proto files)
- README.md added detailed API implementation status table

## [0.1.0] - 2026-04-07

### Added
- Initial release
- Core client implementation (TCP connection, protocol encapsulation)
- InitConnect connection initialization
- Basic Protobuf message definitions
- README, license, and other base files
