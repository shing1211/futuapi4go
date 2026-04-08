// Example: GetOrderBook - 獲取買賣盤數據
//
// This example demonstrates how to use the GetOrderBook API to retrieve
// order book (bid/ask) data for a security.
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"
	"os"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetOrderBook Example / 獲取買賣盤示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{
		Market: &hkMarket,
		Code:   ptrStr("00700"), // Tencent / 騰訊
	}

	// Get order book with 10 levels / 獲取10檔買賣盤
	req := &qot.GetOrderBookRequest{
		Security: security,
		Num:      10,
	}

	resp, err := qot.GetOrderBook(cli, req)
	if err != nil {
		log.Printf("GetOrderBook failed: %v", err)
		return
	}

	fmt.Printf("📊 Order Book for %s (%s)\n", security.GetCode(), resp.Name)
	fmt.Println()

	// Display ask orders (sell orders) / 顯示賣盤
	fmt.Println("🔴 Ask Orders (Sell) / 賣盤:")
	fmt.Printf("  %-6s %-12s %-12s %-10s\n", "Level", "Price", "Volume", "Orders")
	fmt.Println("  " + "--------------------------------------------")

	// Ask list is in reverse order (highest price first)
	for i := len(resp.OrderBookAskList) - 1; i >= 0; i-- {
		ob := resp.OrderBookAskList[i]
		fmt.Printf("  %-6d %-12.2f %-12d %-10d\n",
			len(resp.OrderBookAskList)-i, ob.Price, ob.Volume, ob.OrderCount)
	}

	fmt.Println()

	// Display bid orders (buy orders) / 顯示買盤
	fmt.Println("🟢 Bid Orders (Buy) / 買盤:")
	fmt.Printf("  %-6s %-12s %-12s %-10s\n", "Level", "Price", "Volume", "Orders")
	fmt.Println("  " + "--------------------------------------------")

	for i, ob := range resp.OrderBookBidList {
		fmt.Printf("  %-6d %-12.2f %-12d %-10d\n",
			i+1, ob.Price, ob.Volume, ob.OrderCount)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
