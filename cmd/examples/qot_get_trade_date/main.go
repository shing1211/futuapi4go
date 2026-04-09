// Example: GetTradeDate - 獲取交易日期
//
// This example demonstrates how to use the GetTradeDate API to retrieve
// trading calendar dates for a specific market and date range.
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	fmt.Println("=== GetTradeDate Example / 獲取交易日期示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)

	// Calculate date range (current month) / 計算日期範圍（當月）
	now := time.Now()
	beginTime := now.Format("2006-01-01")
	endTime := now.AddDate(0, 1, 0).Format("2006-01-02")

	// Get trade dates / 獲取交易日期
	req := &qot.GetTradeDateRequest{
		Market:    hkMarket,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	resp, err := qot.GetTradeDate(cli, req)
	if err != nil {
		log.Printf("GetTradeDate failed: %v", err)
		return
	}

	fmt.Printf("📅 Trading Calendar / 交易日曆\n")
	fmt.Printf("Market / 市場: HK / 港股\n")
	fmt.Printf("Period / 期間:  %s to %s\n\n", beginTime, endTime)

	fmt.Printf("Found %d trading days / 找到%d個交易日\n\n", len(resp.TradeDateList), len(resp.TradeDateList))

	// Display trading dates / 顯示交易日期
	for i, td := range resp.TradeDateList {
		// Highlight today / 突出顯示今天
		marker := "  "
		if td.GetTime() == now.Format("2006-01-02") {
			marker = "🔵"
		}

		fmt.Printf("%s %d. %s\n", marker, i+1, td.GetTime())
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

