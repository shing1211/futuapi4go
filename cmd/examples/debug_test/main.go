package main

import (
	"fmt"
	"os"
	"time"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
	"gitee.com/shing1211/futuapi4go/pkg/sys"
)

func main() {
	fmt.Println("=== FutuAPI4Go Debug Test ===")

	cli := futuapi.New()
	defer cli.Close()

	addr := "127.0.0.1:11111"
	if a := os.Getenv("FUTU_ADDR"); a != "" {
		addr = a
	}

	fmt.Printf("📡 Connecting to %s...\n", addr)
	if err := cli.Connect(addr); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Connected!\n")
	fmt.Printf("   ConnID:    %d\n", cli.GetConnID())
	fmt.Printf("   ServerVer: %d\n", cli.GetServerVer())
	fmt.Println()

	// Test GetGlobalState with raw bytes
	fmt.Println("🌐 Testing GetGlobalState...")
	state, err := sys.GetGlobalState(cli)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  ✅ Got global state\n")
	fmt.Printf("    MarketHK: %d\n", state.MarketHK)
	fmt.Printf("    MarketUS: %d\n", state.MarketUS)
	fmt.Printf("    MarketSH: %d\n", state.MarketSH)
	fmt.Printf("    MarketSZ: %d\n", state.MarketSZ)
	fmt.Printf("    QotLogined: %v\n", state.QotLogined)
	fmt.Printf("    TrdLogined: %v\n", state.TrdLogined)
	fmt.Printf("    ServerVer: %d\n", state.ServerVer)
	fmt.Printf("    ConnID: %d\n", state.ConnID)

	// Check if markets are in a valid state
	if state.MarketHK == 0 {
		fmt.Println("\n⚠️  HK market state is 0 (unknown)")
	}
	fmt.Println()

	// Define market and security
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	sec := &qotcommon.Security{Market: &hkMarket, Code: &code}

	// Test 1: GetStaticInfo
	fmt.Println("📋 Test 1: GetStaticInfo")
	staticReq := &qot.GetStaticInfoRequest{
		Market:  hkMarket,
		SecType: int32(qotcommon.SecurityType_SecurityType_Eqty),
	}
	staticResp, err := qot.GetStaticInfo(cli, staticReq)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		fmt.Printf("  ✅ Got %d securities\n", len(staticResp.StaticInfoList))
	}
	fmt.Println()

	// Test 2: Subscribe
	fmt.Println("📡 Test 2: Subscribe")
	subReq := &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      []qot.SubType{qot.SubType_Basic},
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}
	subResp, err := qot.Subscribe(cli, subReq)
	if err != nil {
		fmt.Printf("  ⚠️  Warning: %v\n", err)
	} else {
		fmt.Printf("  ✅ Success (RetType=%d)\n", subResp.RetType)
	}
	fmt.Println()

	// Test 3: GetSubInfo
	fmt.Println("📊 Test 3: GetSubInfo")
	subInfoResp, err := qot.GetSubInfo(cli)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		fmt.Printf("  ✅ UsedQuota=%d, Remain=%d\n", subInfoResp.TotalUsedQuota, subInfoResp.RemainQuota)
	}
	fmt.Println()

	// Test 4: GetBasicQot
	fmt.Println("📊 Test 4: GetBasicQot")
	c2s := &qotgetbasicqot.C2S{SecurityList: []*qotcommon.Security{sec}}
	req := &qotgetbasicqot.Request{C2S: c2s}
	bodyBytes, _ := proto.Marshal(req)
	fmt.Printf("  Request: %d bytes\n", len(bodyBytes))

	quotes, err := qot.GetBasicQot(cli, []*qotcommon.Security{sec})
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		for _, q := range quotes {
			fmt.Printf("  ✅ %s (%s): Price=%.2f\n", q.Security.GetCode(), q.Name, q.CurPrice)
		}
	}
	fmt.Println()

	// Test 5: GetKL
	fmt.Println("📈 Test 5: GetKL (Daily)")
	klReq := &qot.GetKLRequest{
		Security:  sec,
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    3,
	}
	klResp, err := qot.GetKL(cli, klReq)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		fmt.Printf("  ✅ Got %d K-lines\n", len(klResp.KLList))
	}
	fmt.Println()

	// Test 6: GetOrderBook
	fmt.Println("📕 Test 6: GetOrderBook")
	obReq := &qot.GetOrderBookRequest{
		Security: sec,
		Num:      5,
	}
	obResp, err := qot.GetOrderBook(cli, obReq)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
	} else {
		fmt.Printf("  ✅ Got order book (asks=%d, bids=%d)\n",
			len(obResp.OrderBookAskList), len(obResp.OrderBookBidList))
	}

	fmt.Println("\n🎉 Test Complete!")
}

// Wait briefly for any push notifications
func wait() { time.Sleep(2 * time.Second) }
