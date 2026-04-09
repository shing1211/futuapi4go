// Example: GetPositionList - 獲取持倉列表
//
// This example demonstrates how to use the GetPositionList API to retrieve
// current holdings/positions in a trading account.
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetPositionList Example / 獲取持倉列表示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get account list / 獲取賬戶列表
	accResp, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		log.Printf("GetAccList failed: %v", err)
		return
	}

	if len(accResp.AccList) == 0 {
		fmt.Println("No trading accounts found")
		return
	}

	acc := accResp.AccList[0]

	// Get position list / 獲取持倉列表
	req := &trd.GetPositionListRequest{
		AccID:     acc.AccID,
		TrdMarket: 0, // 0 = all markets / 0 = 所有市場
	}

	resp, err := trd.GetPositionList(cli, req)
	if err != nil {
		log.Printf("GetPositionList failed: %v", err)
		return
	}

	fmt.Printf("📊 Portfolio / 投資組合 (AccID: %d)\n\n", acc.AccID)
	fmt.Printf("Found %d positions\n\n", len(resp.PositionList))

	if len(resp.PositionList) == 0 {
		fmt.Println("No positions held / 目前沒有持倉")
		fmt.Println("\n=== Example Complete / 示例完成 ===")
		return
	}

	// Display header / 顯示表頭
	fmt.Printf("  %-10s %-12s %-8s %-8s %-10s %-10s %-10s %-12s\n",
		"Code", "Name", "Qty", "Avail", "Cost", "Price", "Value", "P/L")
	fmt.Println("  " + "------------------------------------------------------------------------------")

	// Display each position / 顯示每個持倉
	totalValue := 0.0
	totalPL := 0.0

	for _, pos := range resp.PositionList {
		value := pos.Val
		pl := pos.PlVal

		totalValue += value
		totalPL += pl

		fmt.Printf("  %-10s %-12s %-8.0f %-8.0f %-10.2f %-10.2f %-10.2f %-12.2f\n",
			pos.Code, pos.Name, pos.Qty, pos.CanSellQty,
			pos.CostPrice, pos.Price, value, pl)
	}

	fmt.Println("  " + "------------------------------------------------------------------------------")
	fmt.Printf("  %-42s %-10.2f %-12.2f\n", "Total / 合計:", totalValue, totalPL)

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

