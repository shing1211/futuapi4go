// Example: GetBasicQot - 獲取實時行情
//
// This example demonstrates how to use the GetBasicQot API to retrieve
// real-time market quotes for one or more securities.
//
// Usage:
//   go run main.go
//
// Note: Requires Futu OpenD running and logged in to quote server.

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
	// Create client
	cli := futuapi.New()
	defer cli.Close()

	// Connect to OpenD
	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetBasicQot Example / 獲取實時行情示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Define securities to query / 定義要查詢的股票 (using HK futures for HSI)
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)

	securities := []*qotcommon.Security{
		{Market: &hkFutureMarket, Code: ptrStr("HSImain")}, // HSI Futures Main Contract / 恆生指數期貨
	}

	fmt.Printf("Querying quotes for %d securities...\n", len(securities))

	// Call GetBasicQot / 調用 GetBasicQot
	quotes, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		log.Printf("GetBasicQot failed: %v", err)
		return
	}

	// Display results / 顯示結果
	fmt.Println("\n=== Real-time Quotes / 實時行情 ===")
	for _, q := range quotes {
		fmt.Printf("\n📈 %s (%s)\n", q.Security.GetCode(), q.Name)
		fmt.Printf("   Current Price / 現價:  %.2f\n", q.CurPrice)
		fmt.Printf("   Open Price / 開盤價:   %.2f\n", q.OpenPrice)
		fmt.Printf("   High Price / 最高價:   %.2f\n", q.HighPrice)
		fmt.Printf("   Low Price / 最低價:    %.2f\n", q.LowPrice)
		fmt.Printf("   Last Close / 昨收價:   %.2f\n", q.LastClosePrice)
		fmt.Printf("   Volume / 成交量:       %d\n", q.Volume)
		fmt.Printf("   Turnover / 成交額:     %.0f\n", q.Turnover)
		fmt.Printf("   Update Time / 更新時間: %s\n", q.UpdateTime)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

// Helper functions / 輔助函數
func ptrStr(s string) *string { return &s }
