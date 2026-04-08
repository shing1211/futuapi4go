// Real-time Push Subscription Examples
//
// This example demonstrates how to use push notification APIs:
// - Subscribe: Subscribe to real-time data streams
// - GetSubInfo: Get subscription information
// - RegQotPush: Register for push notifications
// - Push Handlers: Handle real-time market data pushes
// - Trading Push Handlers: Handle order/fill notifications
//
// Usage:
//   go run main.go
//
// Note: Push notifications work asynchronously - the SDK receives
//       updates automatically after subscription

package main

import (
	"fmt"
	"log"
	"time"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	// Create and connect client
	cli := futuapi.New()
	defer cli.Close()

	addr := "127.0.0.1:11111"
	fmt.Printf("=== Connecting to %s ===\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✓ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Define securities to subscribe
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	usMarket := int32(qotcommon.QotMarket_QotMarket_US_Security)

	securities := []*qotcommon.Security{
		{Market: &hkMarket, Code: ptrStr("00700")}, // Tencent
		{Market: &hkMarket, Code: ptrStr("09988")}, // Alibaba
		{Market: &usMarket, Code: ptrStr("AAPL")},  // Apple
	}

	// 1. Subscribe to Real-time Data
	fmt.Println("=== 1. Subscribe to Real-time Data (Subscribe) ===")

	// Subscribe to multiple data types
	subTypes := []qot.SubType{
		qot.SubType_Basic,     // Real-time quotes
		qot.SubType_OrderBook, // Order book
		qot.SubType_Ticker,    // Tick-by-tick trades
		qot.SubType_KL,        // K-line data
		qot.SubType_RT,        // Real-time minute data
		qot.SubType_Broker,    // Broker queue
	}

	subReq := &qot.SubscribeRequest{
		SecurityList:     securities,
		SubTypeList:      subTypes,
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}

	subResp, err := qot.Subscribe(cli, subReq)
	if err != nil {
		log.Printf("Subscribe failed: %v", err)
	} else {
		fmt.Printf("  ✓ Successfully subscribed to %d data types for %d securities\n",
			len(subTypes), len(securities))
		fmt.Printf("  Return Type: %d\n", subResp.RetType)
		fmt.Printf("  Return Message: %s\n", subResp.RetMsg)
	}
	fmt.Println()

	// 2. Get Subscription Info
	fmt.Println("=== 2. Subscription Info (GetSubInfo) ===")
	subInfoResp, err := qot.GetSubInfo(cli)
	if err != nil {
		log.Printf("GetSubInfo failed: %v", err)
	} else {
		fmt.Printf("  Total Used Quota: %d\n", subInfoResp.TotalUsedQuota)
		fmt.Printf("  Remaining Quota: %d\n", subInfoResp.RemainQuota)

		if len(subInfoResp.ConnSubInfoList) > 0 {
			fmt.Printf("  Active Subscriptions:\n")
			for _, connInfo := range subInfoResp.ConnSubInfoList {
				for _, info := range connInfo.GetSubInfoList() {
					for _, sec := range info.GetSecurityList() {
						fmt.Printf("    %s | SubType=%d\n",
							sec.GetCode(), info.GetSubType())
					}
				}
			}
		}
	}
	fmt.Println()

	// 3. Register for Push Notifications
	fmt.Println("=== 3. Register for Push Notifications (RegQotPush) ===")

	regReq := &qot.RegQotPushRequest{
		SecurityList: securities,
		SubTypeList:  []int32{int32(qotcommon.SubType_SubType_Basic)},
		IsRegOrUnReg: true,
		IsFirstPush:  true,
	}

	regResp, err := qot.RegQotPush(cli, regReq)
	if err != nil {
		log.Printf("RegQotPush failed: %v", err)
	} else {
		fmt.Printf("  ✓ Push notifications registered\n")
		fmt.Printf("  Return Type: %d\n", regResp.RetType)
		fmt.Printf("  Return Message: %s\n", regResp.RetMsg)
	}
	fmt.Println()

	// 4. Set up Push Notification Handlers
	fmt.Println("=== 4. Push Notification Setup ===")
	fmt.Println("  In real applications, you would set up handlers like:")
	fmt.Println()
	fmt.Println("  cli.SetQotPushHandler(func(pkt *	futuapi.Packet) {")
	fmt.Println("    switch pkt.ProtoID {")
	fmt.Println("    case 3101: // Basic quote update")
	fmt.Println("      // Handle basic quote")
	fmt.Println("    case 3102: // K-line update")
	fmt.Println("      // Handle K-line")
	fmt.Println("    case 3103: // Order book update")
	fmt.Println("      // Handle order book")
	fmt.Println("    case 3104: // Ticker update")
	fmt.Println("      // Handle ticker")
	fmt.Println("    case 3105: // RT update")
	fmt.Println("      // Handle RT")
	fmt.Println("    case 3106: // Broker update")
	fmt.Println("      // Handle broker")
	fmt.Println("    }")
	fmt.Println("  })")
	fmt.Println()

	// 5. Example: What you'll receive in push notifications
	fmt.Println("=== 5. Push Notification Examples ===")
	fmt.Println()

	fmt.Println("  Basic Quote Push (3101):")
	fmt.Println("  ┌─────────────┬────────┬────────┬────────┬──────────┐")
	fmt.Println("  │ Code        │ Price  │ Open   │ High   │ Volume   │")
	fmt.Println("  ├─────────────┼────────┼────────┼────────┼──────────┤")
	fmt.Println("  │ 00700       │ 350.00 │ 348.00 │ 352.00 │ 10000000 │")
	fmt.Println("  │ 09988       │ 180.00 │ 178.00 │ 182.00 │ 5000000  │")
	fmt.Println("  └─────────────┴────────┴────────┴────────┴──────────┘")
	fmt.Println()

	fmt.Println("  Order Book Push (3103):")
	fmt.Println("  ┌──────┬────────┬────────┬──────┬────────┬────────┐")
	fmt.Println("  │ Side │ Price  │ Volume │ Side │ Price  │ Volume │")
	fmt.Println("  ├──────┼────────┼────────┼──────┼────────┼────────┤")
	fmt.Println("  │ Ask5 │ 350.05 │ 5000   │ Bid5 │ 349.95 │ 6000   │")
	fmt.Println("  │ Ask4 │ 350.04 │ 4000   │ Bid4 │ 349.96 │ 5500   │")
	fmt.Println("  │ Ask3 │ 350.03 │ 3000   │ Bid3 │ 349.97 │ 5000   │")
	fmt.Println("  │ Ask2 │ 350.02 │ 2000   │ Bid2 │ 349.98 │ 4500   │")
	fmt.Println("  │ Ask1 │ 350.01 │ 1000   │ Bid1 │ 349.99 │ 4000   │")
	fmt.Println("  └──────┴────────┴────────┴──────┴────────┴────────┘")
	fmt.Println()

	fmt.Println("  Ticker Push (3104):")
	fmt.Println("  ┌────────────────────┬────────┬────────┬──────────┬────────┐")
	fmt.Println("  │ Time               │ Price  │ Volume │ Turnover │ Dir    │")
	fmt.Println("  ├────────────────────┼────────┼────────┼──────────┼────────┤")
	fmt.Println("  │ 2026-04-08 10:30:00│ 350.00 │ 100    │ 35000.00 │ Buy    │")
	fmt.Println("  │ 2026-04-08 10:30:01│ 350.10 │ 200    │ 70020.00 │ Sell   │")
	fmt.Println("  │ 2026-04-08 10:30:02│ 349.90 │ 150    │ 52485.00 │ Neutral│")
	fmt.Println("  └────────────────────┴────────┴────────┴──────────┴────────┘")
	fmt.Println()

	fmt.Println("  K-line Push (3102):")
	fmt.Println("  ┌────────────────────┬────────┬────────┬────────┬────────┬──────────┐")
	fmt.Println("  │ Time               │ Open   │ High   │ Low    │ Close  │ Volume   │")
	fmt.Println("  ├────────────────────┼────────┼────────┼────────┼────────┼──────────┤")
	fmt.Println("  │ 2026-04-08 10:30:00│ 349.50 │ 350.50 │ 349.00 │ 350.00 │ 50000    │")
	fmt.Println("  │ 2026-04-08 10:31:00│ 350.00 │ 351.00 │ 349.80 │ 350.50 │ 45000    │")
	fmt.Println("  └────────────────────┴────────┴────────┴────────┴────────┴──────────┘")
	fmt.Println()

	// 6. Trading Push Notifications
	fmt.Println("=== 6. Trading Push Notifications ===")
	fmt.Println("  Trading push notifications include:")
	fmt.Println()
	fmt.Println("  • Order Status Update (7001)")
	fmt.Println("    - Triggered when order status changes")
	fmt.Println("    - Contains: OrderID, Status, Price, Qty, Filled Qty")
	fmt.Println()
	fmt.Println("  • Order Fill Update (7002)")
	fmt.Println("    - Triggered when order is partially or fully filled")
	fmt.Println("    - Contains: FillID, OrderID, Price, Qty, Turnover")
	fmt.Println()
	fmt.Println("  • Trade Notification (7003)")
	fmt.Println("    - General trade notifications")
	fmt.Println("    - Contains: Trade event type and details")
	fmt.Println()
	fmt.Println("  Setup example:")
	fmt.Println("  cli.SetTrdPushHandler(func(pkt *	futuapi.Packet) {")
	fmt.Println("    switch pkt.ProtoID {")
	fmt.Println("    case 7001: // Order status update")
	fmt.Println("      // Handle order update")
	fmt.Println("    case 7002: // Order fill update")
	fmt.Println("      // Handle fill update")
	fmt.Println("    case 7003: // Trade notification")
	fmt.Println("      // Handle trade notification")
	fmt.Println("    }")
	fmt.Println("  })")
	fmt.Println()

	// 7. Unsubscribe from Data
	fmt.Println("=== 7. Unsubscribe from Data ===")
	fmt.Println("  To unsubscribe:")
	fmt.Println()
	fmt.Println("  unsubReq := &qot.SubscribeRequest{")
	fmt.Println("    SecurityList:     securities,")
	fmt.Println("    SubTypeList:      subTypes,")
	fmt.Println("    IsSubOrUnSub:     false, // Set to false")
	fmt.Println("    IsRegOrUnRegPush: false, // Set to false")
	fmt.Println("  }")
	fmt.Println("  qot.Subscribe(cli, unsubReq)")
	fmt.Println()

	// 8. Best Practices
	fmt.Println("=== 8. Best Practices ===")
	fmt.Println("  1. Subscribe only to data types you actually need")
	fmt.Println("  2. Monitor your subscription quota with GetSubInfo")
	fmt.Println("  3. Unsubscribe when no longer needed")
	fmt.Println("  4. Handle push notifications asynchronously")
	fmt.Println("  5. Don't block in push handlers - process data in goroutines")
	fmt.Println("  6. Use thread-safe data structures for push data")
	fmt.Println()

	// Wait to receive some pushes (in real app, this would be event-driven)
	fmt.Println("=== Waiting for Push Notifications ===")
	fmt.Println("  In a real application, push notifications arrive asynchronously.")
	fmt.Println("  Waiting 5 seconds to demonstrate...")
	time.Sleep(5 * time.Second)
	fmt.Println("  Done!")
	fmt.Println()

	fmt.Println("=== Push Subscription Examples Complete ===")
	fmt.Println("💡  Tips:")
	fmt.Println("  • Use simulator to test push without real market data")
	fmt.Println("  • Push notifications are event-driven - set up handlers early")
	fmt.Println("  • Always check subscription quota to avoid hitting limits")
}

// Helper functions
func ptrStr(s string) *string {
	return &s
}

func ptrInt32(v int32) *int32 {
	return &v
}
