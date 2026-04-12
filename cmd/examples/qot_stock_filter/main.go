// Example: StockFilter - 股票篩選
//
// This example demonstrates how to use the StockFilter API to screen
// stocks based on various criteria such as price, market cap, etc.
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
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== StockFilter Example / 股票篩選示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)

	// Build filter request / 構建篩選請求
	req := &qot.StockFilterRequest{
		Begin:  0,
		Num:    20, // Get top 20 results / 獲取前20個結果
		Market: hkMarket,
	}

	// Example: Filter by price range 100-500 / 示例：篩選價格100-500的股票
	// Uncomment to use filters / 取消註釋以使用篩選器
	/*
		req.BaseFilterList = []*qotstockfilter.BaseFilter{
			{
				FieldName:  int32(qotstockfilter.StockField_StockField_CurPrice),
				FilterMin:  ptrFloat64(100.0),
				FilterMax:  ptrFloat64(500.0),
				IsNoFilter: ptrBool(false),
			},
		}
	*/

	fmt.Println("📊 Stock Filter / 股票篩選")
	fmt.Println("Market / 市場: HK / 港股")
	fmt.Println("Filter: Price 100-500 (commented out for demo)")
	fmt.Println()

	resp, err := qot.StockFilter(cli, req)
	if err != nil {
		log.Printf("StockFilter failed: %v", err)
		return
	}

	fmt.Printf("Found %d stocks matching criteria / 找到%d只符合條件的股票\n\n",
		resp.AllCount, resp.AllCount)

	// Display results / 顯示結果
	displayCount := 10
	if len(resp.DataList) < displayCount {
		displayCount = len(resp.DataList)
	}

	fmt.Printf("  %-10s %-20s %-12s\n", "Code", "Name", "Price")
	fmt.Println("  " + "----------------------------------------")

	for i := 0; i < displayCount; i++ {
		stock := resp.DataList[i]

		// Extract price from base data / 從基礎數據提取價格
		price := 0.0
		for _, bd := range stock.BaseDataList {
			if bd.FieldName != nil && *bd.FieldName == int32(qotstockfilter.StockField_StockField_CurPrice) {
				price = bd.GetValue()
				break
			}
		}

		fmt.Printf("  %-10s %-20s %-12.2f\n",
			stock.Security.GetCode(), stock.Name, price)
	}

	if resp.AllCount > int32(displayCount) {
		fmt.Printf("  ... and %d more stocks\n", resp.AllCount-int32(displayCount))
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrFloat64(v float64) *float64 { return &v }
func ptrBool(v bool) *bool          { return &v }
