// Package qot provides market data query APIs for the Futu OpenD SDK.
// It covers real-time quotes, K-lines, order book, tick data, broker queue,
// capital flow, stock screening, options, warrants, and plate/industry data.
//
// Key functions:
//   - GetBasicQot: Real-time quote snapshot
//   - GetKL: K-line data
//   - GetOrderBook: Bid/ask depth
//   - GetTicker: Recent trades
//   - GetCapitalFlow: Money flow
//   - StockFilter: Screen stocks by conditions
//
// Each function follows the pattern:
//
//	qot.GetBasicQot(ctx, &qot.GetBasicQotInput{Code: "HK.00700"})
//
// 简体中文:
// qot 包提供富途 OpenD SDK 的行情数据查询 API，
// 涵盖实时报价、K 线、买卖盘、逐笔成交、经纪队列、资金流向、股票筛选、期权、权证和板块行业数据。
//
// 繁體中文:
// qot 包提供富途 OpenD SDK 的行情數據查詢 API，
// 涵蓋實時報價、K 線、買賣盤、逐筆成交、經紀隊列、資金流向、股票篩選、期權、權證和板塊行業數據。
package qot
