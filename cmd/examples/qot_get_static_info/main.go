// Example: GetStaticInfo - 獲取股票靜態信息
//
// This example demonstrates how to use the GetStaticInfo API to retrieve
// static information about stocks such as name, lot size, listing date, etc.
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

	fmt.Println("=== GetStaticInfo Example / 獲取股票靜態信息示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)

	// Get static info for all HK stocks / 獲取所有港股靜態信息
	req := &qot.GetStaticInfoRequest{
		Market:  hkMarket,
		SecType: int32(qotcommon.SecurityType_SecurityType_Eqty), // Common stocks / 普通股
	}

	resp, err := qot.GetStaticInfo(cli, req)
	if err != nil {
		log.Printf("GetStaticInfo failed: %v", err)
		return
	}

	fmt.Printf("📊 Static Info / 靜態信息\n")
	fmt.Printf("Market / 市場: HK / 港股\n")
	fmt.Printf("Type / 類型: Eqty / 普通股\n")
	fmt.Printf("\nFound %d securities / 找到%d只股票\n\n", len(resp.StaticInfoList), len(resp.StaticInfoList))

	// Display first 20 stocks / 顯示前20只股票
	displayCount := 20
	if len(resp.StaticInfoList) < displayCount {
		displayCount = len(resp.StaticInfoList)
	}

	fmt.Printf("  %-10s %-20s %-10s %-12s %-12s\n", "Code", "Name", "Lot Size", "List Date", "SecType")
	fmt.Println("  " + "----------------------------------------------------------------------------")

	for i := 0; i < displayCount; i++ {
		info := resp.StaticInfoList[i]
		basic := info.GetBasic()
		if basic == nil {
			continue
		}

		fmt.Printf("  %-10s %-20s %-10d %-12s %-12d\n",
			basic.GetSecurity().GetCode(),
			basic.GetName(),
			basic.GetLotSize(),
			basic.GetListTime(),
			basic.GetSecType())
	}

	if len(resp.StaticInfoList) > displayCount {
		fmt.Printf("  ... and %d more securities\n", len(resp.StaticInfoList)-displayCount)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
