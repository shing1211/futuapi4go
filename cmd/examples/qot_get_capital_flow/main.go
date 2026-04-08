// Example: GetCapitalFlow - 獲取資金流向
//
// This example demonstrates how to use the GetCapitalFlow API to retrieve
// capital flow data showing money movement in/out of a security.
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

	fmt.Println("=== GetCapitalFlow Example / 獲取資金流向示例 ===")
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

	// Get capital flow (daily) / 獲取資金流向（日線）
	req := &qot.GetCapitalFlowRequest{
		Security:   security,
		PeriodType: 1, // Daily / 日線
	}

	resp, err := qot.GetCapitalFlow(cli, req)
	if err != nil {
		log.Printf("GetCapitalFlow failed: %v", err)
		return
	}

	fmt.Printf("📊 Capital Flow / 資金流向 for %s\n\n", security.GetCode())

	// Display capital flow data / 顯示資金流向數據
	displayCount := 10
	if len(resp.FlowItemList) < displayCount {
		displayCount = len(resp.FlowItemList)
	}

	fmt.Printf("  %-20s %-12s %-12s %-12s %-12s\n",
		"Time", "In Flow", "Out Flow", "Net Flow", "Main In")
	fmt.Println("  " + "------------------------------------------------------------")

	for i := 0; i < displayCount; i++ {
		flow := resp.FlowItemList[i]
		netFlow := flow.InFlow - flow.OutFlow
		fmt.Printf("  %-20s %-12.0f %-12.0f %-12.0f %-12.0f\n",
			flow.Time, flow.InFlow, flow.OutFlow, netFlow, flow.MainInFlow)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
