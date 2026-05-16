// Package push provides handlers for parsing push notification payloads
// from Futu OpenD. Use RegisterHandler on the client to receive real-time
// market data and order updates.
//
// Usage:
//
//	cli.RegisterHandler(push.ProtoID_Qot_UpdateBasicQot, func(protoID uint32, body []byte) {
//	    data, err := push.ParseUpdateBasicQot(body)
//	    // ...
//	})
//
// Supported push types:
//   - Basic quote updates (real-time price changes)
//   - Order book (bid/ask) updates
//   - K-line updates
//   - Ticker (trade) updates
//   - Broker queue updates
//   - Order and order fill status updates
//   - Price reminder notifications
//
// 简体中文:
// push 包提供富途 OpenD 推送通知的解析处理器。
// 使用客户端的 RegisterHandler 方法接收实时行情和订单更新。
//
// 繁體中文:
// push 包提供富途 OpenD 推送通知的解析處理器。
// 使用客戶端的 RegisterHandler 方法接收實時行情和訂單更新。
package push
