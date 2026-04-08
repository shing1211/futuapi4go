// Market Data Examples - Basic Quote APIs
//
// This example demonstrates how to use basic market data APIs:
// - GetBasicQot: Real-time quotes
// - GetKL: K-line (candlestick) data
// - GetOrderBook: Order book (bid/ask)
// - GetTicker: Tick-by-tick trades
// - GetRT: Real-time minute data
// - GetBroker: Broker queue
//
// Usage:
//   With real OpenD: go run main.go
//   With simulator:  Start simulator first, then run

package main

import (
	"fmt"
	"log"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	// Create client
	cli := futuapi.New()
	defer cli.Close()

	// Connect to OpenD or Simulator
	addr := "127.0.0.1:11111"
	fmt.Printf("=== Connecting to %s ===\n", addr)
	
	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✓ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Define securities to query
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	usMarket := int32(qotcommon.QotMarket_QotMarket_US_Security)
	shMarket := int32(qotcommon.QotMarket_QotMarket_CNSH_Security)

	securities := []*qotcommon.Security{
		{Market: &hkMarket, Code: ptrStr("00700")},  // Tencent
		{Market: &hkMarket, Code: ptrStr("09988")},  // Alibaba
		{Market: &usMarket, Code: ptrStr("AAPL")},   // Apple
		{Market: &shMarket, Code: ptrStr("600519")}, // Kweichow Moutai
	}

	// 1. Get Real-time Quotes
	fmt.Println("=== 1. Real-time Quotes (GetBasicQot) ===")
	quotes, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		log.Printf("GetBasicQot failed: %v", err)
	} else {
		for _, q := range quotes {
			fmt.Printf("  %s (%s): Price=%.2f | Open=%.2f | High=%.2f | Low=%.2f | Vol=%d | Turnover=%.0f\n",
				q.Security.GetCode(), q.Name, q.CurPrice, q.OpenPrice,
				q.HighPrice, q.LowPrice, q.Volume, q.Turnover)
		}
	}
	fmt.Println()

	// 2. Get K-line Data (Daily)
	fmt.Println("=== 2. K-line Data - Daily (GetKL) ===")
	klReq := &qot.GetKLRequest{
		Security:  securities[0],
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    5,
	}
	
	klResp, err := qot.GetKL(cli, klReq)
	if err != nil {
		log.Printf("GetKL failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", klResp.Security.GetCode(), klResp.Name)
		fmt.Printf("  %-20s %-10s %-10s %-10s %-10s %-12s\n",
			"Time", "Open", "High", "Low", "Close", "Volume")
		for _, kl := range klResp.KLList {
			fmt.Printf("  %-20s %-10.2f %-10.2f %-10.2f %-10.2f %-12d\n",
				kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice,
				kl.ClosePrice, kl.Volume)
		}
	}
	fmt.Println()

	// 3. Get K-line Data (1-Minute)
	fmt.Println("=== 3. K-line Data - 1 Minute (GetKL) ===")
	klReq1Min := &qot.GetKLRequest{
		Security:  securities[0],
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_1Min),
		ReqNum:    5,
	}
	
	klResp1Min, err := qot.GetKL(cli, klReq1Min)
	if err != nil {
		log.Printf("GetKL 1-min failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", klResp1Min.Security.GetCode(), klResp1Min.Name)
		for _, kl := range klResp1Min.KLList {
			fmt.Printf("  %s: O=%.2f H=%.2f L=%.2f C=%.2f V=%d\n",
				kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice,
				kl.ClosePrice, kl.Volume)
		}
	}
	fmt.Println()

	// 4. Get Order Book
	fmt.Println("=== 4. Order Book (GetOrderBook) ===")
	obReq := &qot.GetOrderBookRequest{
		Security: securities[0],
		Num:      5, // Top 5 levels
	}
	
	obResp, err := qot.GetOrderBook(cli, obReq)
	if err != nil {
		log.Printf("GetOrderBook failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", obResp.Security.GetCode(), obResp.Name)
		fmt.Println("  Asks (Sell orders):")
		for i, ask := range obResp.OrderBookAskList {
			fmt.Printf("    %d: Price=%.2f | Volume=%d | Orders=%d\n",
				i+1, ask.Price, ask.Volume, ask.OrderCount)
		}
		fmt.Println("  Bids (Buy orders):")
		for i, bid := range obResp.OrderBookBidList {
			fmt.Printf("    %d: Price=%.2f | Volume=%d | Orders=%d\n",
				i+1, bid.Price, bid.Volume, bid.OrderCount)
		}
	}
	fmt.Println()

	// 5. Get Ticker (Recent Trades)
	fmt.Println("=== 5. Recent Trades (GetTicker) ===")
	tickerReq := &qot.GetTickerRequest{
		Security: securities[0],
		Num:      10,
	}
	
	tickerResp, err := qot.GetTicker(cli, tickerReq)
	if err != nil {
		log.Printf("GetTicker failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", tickerResp.Security.GetCode(), tickerResp.Name)
		fmt.Printf("  %-20s %-10s %-10s %-12s %-10s\n",
			"Time", "Price", "Volume", "Turnover", "Direction")
		for _, t := range tickerResp.TickerList {
			dir := "Neutral"
			if t.Dir == 1 {
				dir = "Buy"
			} else if t.Dir == 2 {
				dir = "Sell"
			}
			fmt.Printf("  %-20s %-10.2f %-10d %-12.2f %-10s\n",
				t.Time, t.Price, t.Volume, t.Turnover, dir)
		}
	}
	fmt.Println()

	// 6. Get Real-time Minute Data
	fmt.Println("=== 6. Real-time Minute Data (GetRT) ===")
	rtReq := &qot.GetRTRequest{
		Security: securities[0],
	}
	
	rtResp, err := qot.GetRT(cli, rtReq)
	if err != nil {
		log.Printf("GetRT failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", rtResp.Security.GetCode(), rtResp.Name)
		// Show first 5 entries
		count := 5
		if len(rtResp.RTList) < count {
			count = len(rtResp.RTList)
		}
		fmt.Printf("  %-20s %-10s %-12s %-12s\n",
			"Time", "Price", "Volume", "AvgPrice")
		for i := 0; i < count; i++ {
			rt := rtResp.RTList[i]
			fmt.Printf("  %-20s %-10.2f %-12d %-12.2f\n",
				rt.Time, rt.Price, rt.Volume, rt.AvgPrice)
		}
	}
	fmt.Println()

	// 7. Get Broker Queue
	fmt.Println("=== 7. Broker Queue (GetBroker) ===")
	brokerReq := &qot.GetBrokerRequest{
		Security: securities[0],
		Num:      10,
	}
	
	brokerResp, err := qot.GetBroker(cli, brokerReq)
	if err != nil {
		log.Printf("GetBroker failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s (%s)\n", brokerResp.Security.GetCode(), brokerResp.Name)
		fmt.Println("  Ask Brokers:")
		for _, b := range brokerResp.AskBrokerList {
			fmt.Printf("    Pos=%d | Broker=%s | Volume=%d\n",
				b.Pos, b.Name, b.Volume)
		}
		fmt.Println("  Bid Brokers:")
		for _, b := range brokerResp.BidBrokerList {
			fmt.Printf("    Pos=%d | Broker=%s | Volume=%d\n",
				b.Pos, b.Name, b.Volume)
		}
	}
	fmt.Println()

	// 8. Get Security Snapshot (Comprehensive data)
	fmt.Println("=== 8. Security Snapshot (GetSecuritySnapshot) ===")
	snapReq := &qot.GetSecuritySnapshotRequest{
		SecurityList: securities[:1], // Just first one
	}
	
	snapResp, err := qot.GetSecuritySnapshot(cli, snapReq)
	if err != nil {
		log.Printf("GetSecuritySnapshot failed: %v", err)
	} else {
		for _, snap := range snapResp.SnapshotList {
			basic := snap.GetBasic()
			if basic == nil {
				continue
			}
			sec := basic.GetSecurity()
			fmt.Printf("  %s (%s)\n", sec.GetCode(), basic.GetName())
			fmt.Printf("    CurPrice: %.2f | Volume: %d | Turnover: %.0f\n",
				basic.GetCurPrice(), basic.GetVolume(), basic.GetTurnover())
			fmt.Printf("    High: %.2f | Low: %.2f | Open: %.2f\n",
				basic.GetHighPrice(), basic.GetLowPrice(), basic.GetOpenPrice())
		}
	}
	fmt.Println()

	fmt.Println("=== Examples Complete ===")
	fmt.Println("Tip: This example works with both simulator AND real Futu OpenD!")
}

// Helper functions
func ptrStr(s string) *string { return &s }
func ptrInt32(v int32) *int32 { return &v }
func ptrInt64(v int64) *int64 { return &v }
func ptrFloat64(v float64) *float64 { return &v }
func ptrBool(v bool) *bool { return &v }
