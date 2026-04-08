// FutuAPI4Go Comprehensive Demo
//
// This demo showcases the full capabilities of the FutuAPI4Go SDK:
// - Connection management
// - Market data queries (basic and advanced)
// - Trading operations
// - Real-time subscriptions
// - Error handling
//
// Usage:
//   With simulator:  Start simulator, then run this demo
//   With real OpenD:  Ensure OpenD is running, then run this demo
//
// go run main.go

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
	"gitee.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	printHeader()

	// Create client
	cli := client.New()
	defer cli.Close()

	// Connect
	addr := getEnvOrDefault("FUTU_ADDR", "127.0.0.1:11111")
	if err := connect(cli, addr); err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}

	// Run demos
	runMarketDataDemo(cli)
	runAdvancedMarketDataDemo(cli)
	runTradingDemo(cli)
	runSubscriptionDemo(cli)

	printFooter()
}

func printHeader() {
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                           ║")
	fmt.Println("║          FutuAPI4Go - Comprehensive Demo                 ║")
	fmt.Println("║                                                           ║")
	fmt.Println("║  World-class Golang SDK for Futu OpenD API               ║")
	fmt.Println("║  Version: 0.3.0                                          ║")
	fmt.Println("║                                                           ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printFooter() {
	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                           ║")
	fmt.Println("║  ✓ Demo Complete!                                        ║")
	fmt.Println("║                                                           ║")
	fmt.Println("║  Next Steps:                                             ║")
	fmt.Println("║  • Explore examples/01_* for detailed API usage         ║")
	fmt.Println("║  • Read USER_GUIDE.md for complete documentation        ║")
	fmt.Println("║  • Use simulator for local testing without OpenD        ║")
	fmt.Println("║                                                           ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
}

func connect(cli *client.Client, addr string) error {
	fmt.Printf("📡 Connecting to %s...\n", addr)
	
	if err := cli.Connect(addr); err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	
	fmt.Printf("✅ Connected successfully!\n")
	fmt.Printf("   Connection ID: %d\n", cli.GetConnID())
	fmt.Printf("   Server Version: %d\n", cli.GetServerVer())
	fmt.Println()
	return nil
}

func runMarketDataDemo(cli *client.Client) {
	printSection("1. Market Data - Basic")

	// Define securities
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	usMarket := int32(qotcommon.QotMarket_QotMarket_US_Security)
	shMarket := int32(qotcommon.QotMarket_QotMarket_CNSH_Security)

	securities := []*qotcommon.Security{
		{Market: &hkMarket, Code: ptrStr("00700")},  // Tencent
		{Market: &hkMarket, Code: ptrStr("09988")},  // Alibaba
		{Market: &usMarket, Code: ptrStr("AAPL")},   // Apple
		{Market: &shMarket, Code: ptrStr("600519")}, // Kweichow Moutai
	}

	// 1.1 Real-time Quotes
	fmt.Println("📊 Real-time Quotes:")
	quotes, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		fmt.Printf("   ❌ GetBasicQot failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Retrieved %d quotes\n", len(quotes))
		for _, q := range quotes {
			fmt.Printf("   • %s (%s): ¥%.2f | Vol: %d\n",
				q.Security.GetCode(), q.Name, q.CurPrice, q.Volume)
		}
	}
	fmt.Println()

	// 1.2 K-line Data
	fmt.Println("📈 K-line Data (Daily, 5 bars):")
	klReq := &qot.GetKLRequest{
		Security:  securities[0],
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    5,
	}
	
	klResp, err := qot.GetKL(cli, klReq)
	if err != nil {
		fmt.Printf("   ❌ GetKL failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Retrieved %d K-lines for %s\n", 
			len(klResp.KLList), klResp.Security.GetCode())
		for i, kl := range klResp.KLList {
			fmt.Printf("   [%d] %s | O: %.2f H: %.2f L: %.2f C: %.2f | V: %d\n",
				i+1, kl.Time, kl.OpenPrice, kl.HighPrice, 
				kl.LowPrice, kl.ClosePrice, kl.Volume)
		}
	}
	fmt.Println()

	// 1.3 Order Book
	fmt.Println("📕 Order Book (Top 5):")
	obReq := &qot.GetOrderBookRequest{
		Security: securities[0],
		Num:      5,
	}
	
	obResp, err := qot.GetOrderBook(cli, obReq)
	if err != nil {
		fmt.Printf("   ❌ GetOrderBook failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Order Book for %s\n", obResp.Security.GetCode())
		fmt.Printf("   Asks (Sell):\n")
		for i, ask := range obResp.OrderBookAskList {
			fmt.Printf("     [%d] %.2f | %d\n", i+1, ask.Price, ask.Volume)
		}
		fmt.Printf("   Bids (Buy):\n")
		for i, bid := range obResp.OrderBookBidList {
			fmt.Printf("     [%d] %.2f | %d\n", i+1, bid.Price, bid.Volume)
		}
	}
	fmt.Println()
}

func runAdvancedMarketDataDemo(cli *client.Client) {
	printSection("2. Market Data - Advanced")

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{
		Market: &hkMarket,
		Code:   ptrStr("00700"),
	}

	// 2.1 Capital Flow
	fmt.Println("💰 Capital Flow:")
	capFlowReq := &qot.GetCapitalFlowRequest{
		Security:   security,
		PeriodType: 1,
	}
	
	capFlowResp, err := qot.GetCapitalFlow(cli, capFlowReq)
	if err != nil {
		fmt.Printf("   ❌ GetCapitalFlow failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Retrieved %d capital flow records\n", 
			len(capFlowResp.FlowItemList))
		if len(capFlowResp.FlowItemList) > 0 {
			flow := capFlowResp.FlowItemList[0]
			fmt.Printf("   Latest: %s | Main In: %.2f | Main Out: %.2f\n",
				flow.Time, flow.MainInFlow, flow.MainOutFlow)
		}
	}
	fmt.Println()

	// 2.2 Capital Distribution
	fmt.Println("📊 Capital Distribution:")
	capDistResp, err := qot.GetCapitalDistribution(cli, security)
	if err != nil {
		fmt.Printf("   ❌ GetCapitalDistribution failed: %v\n", err)
	} else {
		cd := capDistResp.CapitalDistribution
		fmt.Printf("   ✓ Super In: %.2f | Super Out: %.2f\n",
			cd.CapitalInSuper, cd.CapitalOutSuper)
		fmt.Printf("     Big In: %.2f | Big Out: %.2f\n",
			cd.CapitalInBig, cd.CapitalOutBig)
	}
	fmt.Println()

	// 2.3 Stock Filter
	fmt.Println("🔍 Stock Filter:")
	filterReq := &qot.StockFilterRequest{
		Begin:  0,
		Num:    10,
		Market: hkMarket,
	}
	
	filterResp, err := qot.StockFilter(cli, filterReq)
	if err != nil {
		fmt.Printf("   ❌ StockFilter failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d stocks (showing %d)\n",
			filterResp.AllCount, len(filterResp.DataList))
	}
	fmt.Println()

	// 2.4 Trading Dates
	fmt.Println("📅 Trading Dates (April 2026):")
	tradeDateReq := &qot.GetTradeDateRequest{
		Market:    hkMarket,
		BeginTime: "2026-04-01",
		EndTime:   "2026-04-30",
	}
	
	tradeDateResp, err := qot.GetTradeDate(cli, tradeDateReq)
	if err != nil {
		fmt.Printf("   ❌ GetTradeDate failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d trading days\n", 
			len(tradeDateResp.TradeDateList))
	}
	fmt.Println()
}

func runTradingDemo(cli *client.Client) {
	printSection("3. Trading Operations")

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	trdCategory := int32(trdcommon.TrdCategory_TrdCategory_Security)

	// 3.1 Get Account List
	fmt.Println("👤 Account List:")
	accResp, err := trd.GetAccList(cli, trdCategory, false)
	if err != nil {
		fmt.Printf("   ❌ GetAccList failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d accounts\n", len(accResp.AccList))
		for i, acc := range accResp.AccList {
			fmt.Printf("   [%d] AccID: %d | Type: %d | Status: %d\n",
				i+1, acc.AccID, acc.AccType, acc.AccStatus)
		}
	}
	
	var accID uint64 = 123456789
	if accResp != nil && len(accResp.AccList) > 0 {
		accID = accResp.AccList[0].AccID
	}
	fmt.Println()

	// 3.2 Get Funds
	fmt.Println("💵 Account Funds:")
	fundsReq := &trd.GetFundsRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
	}
	
	fundsResp, err := trd.GetFunds(cli, fundsReq)
	if err != nil {
		fmt.Printf("   ❌ GetFunds failed: %v\n", err)
	} else {
		f := fundsResp.Funds
		fmt.Printf("   ✓ Total Assets: %.2f\n", f.TotalAssets)
		fmt.Printf("     Cash: %.2f | Market Value: %.2f\n", f.Cash, f.MarketVal)
		fmt.Printf("     Available: %.2f | Frozen: %.2f\n", 
			f.AvailableFunds, f.FrozenCash)
	}
	fmt.Println()

	// 3.3 Get Positions
	fmt.Println("📦 Positions:")
	posReq := &trd.GetPositionListRequest{
		AccID:     accID,
		TrdMarket: 0,
	}
	
	posResp, err := trd.GetPositionList(cli, posReq)
	if err != nil {
		fmt.Printf("   ❌ GetPositionList failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d positions\n", len(posResp.PositionList))
		for _, pos := range posResp.PositionList {
			fmt.Printf("   • %s (%s): Qty=%.0f | Cost=%.2f | Val=%.2f\n",
				pos.Code, pos.Name, pos.Qty, pos.CostPrice, pos.Val)
		}
	}
	fmt.Println()

	// 3.4 Place Order
	fmt.Println("📝 Place Order (Buy 100 shares of 00700 @ 350.00):")
	placeReq := &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		Code:      "00700",
		TrdSide:   int32(trdcommon.TrdSide_TrdSide_Buy),
		OrderType: int32(trdcommon.OrderType_OrderType_Normal),
		Price:     350.00,
		Qty:       100.0,
	}
	
	placeResp, err := trd.PlaceOrder(cli, placeReq)
	if err != nil {
		fmt.Printf("   ❌ PlaceOrder failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Order placed! OrderID: %d\n", placeResp.OrderID)
		
		// 3.5 Get Order List
		fmt.Println("📋 Order List:")
		orderReq := &trd.GetOrderListRequest{
			AccID:     accID,
			TrdMarket: hkMarket,
		}
		
		orderResp, err := trd.GetOrderList(cli, orderReq)
		if err != nil {
			fmt.Printf("   ❌ GetOrderList failed: %v\n", err)
		} else {
			fmt.Printf("   ✓ Found %d orders\n", len(orderResp.OrderList))
		}
	}
	fmt.Println()

	// 3.6 Max Trade Quantities
	fmt.Println("📊 Max Trade Quantities:")
	maxQtyReq := &trd.GetMaxTrdQtysRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		Code:      "00700",
		Price:     350.00,
		OrderType: int32(trdcommon.OrderType_OrderType_Normal),
	}
	
	maxQtyResp, err := trd.GetMaxTrdQtys(cli, maxQtyReq)
	if err != nil {
		fmt.Printf("   ❌ GetMaxTrdQtys failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Max Cash Buy: %.0f\n", maxQtyResp.MaxCashBuy)
		fmt.Printf("     Max Margin Buy: %.0f\n", maxQtyResp.MaxMarginBuy)
		fmt.Printf("     Max Sell: %.0f\n", maxQtyResp.MaxSell)
	}
	fmt.Println()
}

func runSubscriptionDemo(cli *client.Client) {
	printSection("4. Real-time Subscriptions")

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &hkMarket, Code: ptrStr("00700")},
	}

	// 4.1 Subscribe
	fmt.Println("📡 Subscribe to Real-time Data:")
	subReq := &qot.SubscribeRequest{
		SecurityList:         securities,
		SubTypeList:          []qot.SubType{qot.SubType_Basic, qot.SubType_KL},
		IsSubOrUnSub:         true,
		IsRegOrUnRegPush:     true,
	}
	
	subResp, err := qot.Subscribe(cli, subReq)
	if err != nil {
		fmt.Printf("   ❌ Subscribe failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Subscribed successfully (RetType: %d)\n", subResp.RetType)
	}
	fmt.Println()

	// 4.2 Get Sub Info
	fmt.Println("ℹ️  Subscription Info:")
	subInfoResp, err := qot.GetSubInfo(cli)
	if err != nil {
		fmt.Printf("   ❌ GetSubInfo failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Used Quota: %d | Remaining: %d\n",
			subInfoResp.TotalUsedQuota, subInfoResp.RemainQuota)
	}
	fmt.Println()

	// 4.3 Register Push
	fmt.Println("🔔 Register Push Notifications:")
	regReq := &qot.SubscribeRequest{
		SecurityList:         securities,
		SubTypeList:          []qot.SubType{qot.SubType_Basic},
		IsSubOrUnSub:         true,
		IsRegOrUnRegPush:     true,
	}
	
	regResp, err := qot.RegQotPush(cli, regReq)
	if err != nil {
		fmt.Printf("   ❌ RegQotPush failed: %v\n", err)
	} else {
		fmt.Printf("   ✓ Push registered (RetType: %d)\n", regResp.RetType)
	}
	fmt.Println()

	fmt.Println("💡 Push notifications are now active!")
	fmt.Println("   In a real application, you would set up handlers to process them.")
	fmt.Println()
}

func printSection(title string) {
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Printf("║  %-60s ║\n", title)
	fmt.Println("╠═══════════════════════════════════════════════════════════╣")
}

func getEnvOrDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

// Helper functions
func ptrStr(s string) *string {
	return &s
}

func ptrInt32(v int32) *int32 {
	return &v
}

func ptrFloat64(v float64) *float64 {
	return &v
}

// Ensure time import is used
var _ = time.Now
