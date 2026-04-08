package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
	"gitee.com/shing1211/futuapi4go/pkg/sys"
)

func main() {
	fmt.Println("=== FutuAPI4Go Connection Test ===\n")
	
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
	
	// Test GetGlobalState
	fmt.Println("🌐 Testing GetGlobalState...")
	state, err := sys.GetGlobalState(cli)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  ✅ Got global state\n")
	fmt.Printf("    Market HK: %d | US: %d | SH: %d | SZ: %d\n", 
		state.MarketHK, state.MarketUS, state.MarketSH, state.MarketSZ)
	fmt.Printf("    QotLogined: %v | TrdLogined: %v\n", state.QotLogined, state.TrdLogined)
	fmt.Println()
	
	// Test GetBasicQot with detailed debug
	fmt.Println("📊 Testing GetBasicQot...")
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	sec := &qotcommon.Security{
		Market: &hkMarket,
		Code:   &code,
	}
	
	fmt.Printf("  Security: Market=%d, Code=%s\n", hkMarket, code)
	
	// Try with slice
	securities := []*qotcommon.Security{sec}
	
	fmt.Println("  Calling qot.GetBasicQot...")
	fmt.Printf("  Security: Market=%d, Code=%s\n", sec.GetMarket(), sec.GetCode())
	
	// Show what protobuf bytes we're sending
	c2s := &qotgetbasicqot.C2S{SecurityList: securities}
	req := &qotgetbasicqot.Request{C2S: c2s}
	bodyBytes, _ := proto.Marshal(req)
	fmt.Printf("  Request body: %d bytes\n", len(bodyBytes))
	fmt.Printf("  First 50 bytes: % x\n", bodyBytes[:min(len(bodyBytes), 50)])
	
	quotes, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("  ✅ Got %d quotes\n", len(quotes))
	if len(quotes) > 0 {
		q := quotes[0]
		fmt.Printf("  %s (%s)\n", q.Security.GetCode(), q.Name)
		fmt.Printf("    Price: %.2f\n", q.CurPrice)
		fmt.Printf("    Volume: %d\n", q.Volume)
	}
	
	fmt.Println("\n🎉 SDK is working!")
}

func ptrStr(s string) *string { return &s }
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
