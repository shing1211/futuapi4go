// Example: Subscribe - 訂閱實時行情
//
// This example demonstrates how to use the Subscribe API to subscribe
// to real-time market data for securities.
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== Subscribe Example / 訂閱實時行情示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Define securities to subscribe / 定義要訂閱的股票 (using HK futures for HSI)
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)
	securities := []*qotcommon.Security{
		{Market: &hkFutureMarket, Code: ptrStr("HSImain")}, // HSI Futures Main Contract / 恆生指數期貨
	}

	// Subscribe to multiple data types / 訂閱多種數據類型
	subTypes := []qot.SubType{
		qot.SubType_Basic,     // Real-time quotes / 實時報價
		qot.SubType_OrderBook, // Order book / 買賣盤
		qot.SubType_Ticker,    // Tick-by-tick / 逐筆
		qot.SubType_KL,        // K-line / K線
		qot.SubType_RT,        // Real-time minute data / 實時分時
		qot.SubType_Broker,    // Broker queue / 經紀隊列
	}

	req := &qot.SubscribeRequest{
		SecurityList:     securities,
		SubTypeList:      subTypes,
		IsSubOrUnSub:     true, // true=subscribe, false=unsubscribe
		IsRegOrUnRegPush: true, // true=register for push, false=unregister
	}

	fmt.Printf("Subscribing to %d data types for %d securities...\n\n", len(subTypes), len(securities))

	resp, err := qot.Subscribe(cli, req)
	if err != nil {
		log.Printf("Subscribe failed: %v", err)
		return
	}

	fmt.Printf("✅ Subscribe Result / 訂閱結果:\n")
	fmt.Printf("   RetType: %d\n", resp.RetType)
	fmt.Printf("   RetMsg: %s\n\n", resp.RetMsg)

	// Check subscription info / 檢查訂閱信息
	subInfo, err := qot.GetSubInfo(cli)
	if err != nil {
		fmt.Printf("⚠️  GetSubInfo failed: %v (this is normal during testing)\n", err)
	} else {
		fmt.Printf("📊 Subscription Info / 訂閱信息:\n")
		fmt.Printf("   Total Used Quota / 已用額度: %d\n", subInfo.TotalUsedQuota)
		fmt.Printf("   Remaining Quota / 剩餘額度:   %d\n", subInfo.RemainQuota)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
