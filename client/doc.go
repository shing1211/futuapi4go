// Package client provides the main public API for connecting to Futu OpenD
// and performing market data queries, trading operations, and system commands.
//
// Usage:
//
//	cli := client.New(client.WithRSAPrivateKey(pem))
//	cli.Connect("127.0.0.1:11111")
//	defer cli.Close()
//
//	// Market data
//	snap, _ := cli.GetSecuritySnapshot(ctx, "HK.00700")
//
//	// Trading
//	orderID, _ := cli.PlaceOrder(ctx, &trd.PlaceOrderInput{...})
//
//	// Push notifications
//	cli.RegisterHandler(push.ProtoID_Qot_UpdateBasicQot, handler)
//
// 简体中文:
// client 包提供了连接富途 OpenD 并执行行情查询、交易操作和系统命令的公共 API。
//
// 繁體中文:
// client 包提供了連接富途 OpenD 並執行行情查詢、交易操作和系統命令的公共 API。
package client
