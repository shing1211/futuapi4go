// Package trd provides trading APIs for the Futu OpenD SDK,
// including order placement, modification, cancellation, position/portfolio
// queries, account management, and order building.
//
// Key functions:
//   - PlaceOrder: Place a new order
//   - ModifyOrder: Modify an existing order
//   - CancelOrder: Cancel an order
//   - GetOrderList: Query open orders
//   - GetPositionList: Query positions
//   - GetAccountList: Query trade accounts
//   - GetFunds: Query account funds
//
// Order building:
//
//	order := trd.NewOrderBuilder(trdcommon.OrderType_Normal, trdcommon.TrdSide_Buy).
//	    StockCode("HK.00700").
//	    Price(380.0).
//	    Qty(100).
//	    Build()
//
// 简体中文:
// trd 包提供富途 OpenD SDK 的交易 API，
// 包括下单、改单、撤单、持仓/组合查询、账户管理以及订单构建器。
//
// 繁體中文:
// trd 包提供富途 OpenD SDK 的交易 API，
// 包括下單、改單、撤單、持倉/組合查詢、賬戶管理以及訂單構建器。
package trd
