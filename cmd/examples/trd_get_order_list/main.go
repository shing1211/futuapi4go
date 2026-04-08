// Example: GetOrderList - 獲取訂單列表
//
// This example demonstrates how to use the GetOrderList API to retrieve
// current orders in a trading account.
//
// Usage:
//   go run main.go [acc_id]

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"gitee.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetOrderList Example / 獲取訂單列表示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get account ID / 獲取賬戶ID
	var accID uint64
	if len(os.Args) > 1 {
		id, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalf("Invalid account ID: %v", err)
		}
		accID = id
	} else {
		accResp, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
		if err != nil {
			log.Printf("GetAccList failed: %v", err)
			return
		}
		if len(accResp.AccList) == 0 {
			fmt.Println("No trading accounts found")
			return
		}
		accID = accResp.AccList[0].AccID
	}

	fmt.Printf("Using Account / 使用賬戶: %d\n\n", accID)

	// Get order list / 獲取訂單列表
	req := &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: 0, // 0 = all markets / 0 = 所有市場
	}

	resp, err := trd.GetOrderList(cli, req)
	if err != nil {
		log.Printf("GetOrderList failed: %v", err)
		return
	}

	fmt.Printf("📊 Order List / 訂單列表 (Account: %d)\n\n", accID)
	fmt.Printf("Found %d orders\n\n", len(resp.OrderList))

	if len(resp.OrderList) == 0 {
		fmt.Println("No orders found / 沒有訂單")
		fmt.Println("\n=== Example Complete / 示例完成 ===")
		return
	}

	// Display orders / 顯示訂單
	fmt.Printf("  %-12s %-10s %-8s %-6s %-6s %-8s %-8s %-10s\n",
		"OrderID", "Code", "Side", "Type", "Status", "Qty", "Price", "Filled")
	fmt.Println("  " + "---------------------------------------------------------------------------")

	for _, o := range resp.OrderList {
		side := "Buy"
		if o.TrdSide == int32(trdcommon.TrdSide_TrdSide_Sell) {
			side = "Sell"
		}

		orderType := "Normal"
		if o.OrderType == int32(trdcommon.OrderType_OrderType_Market) {
			orderType = "Market"
		}

		status := "Unknown"
		switch o.OrderStatus {
		case int32(trdcommon.OrderState_OrderState_Submitting):
			status = "Submitting"
		case int32(trdcommon.OrderState_OrderState_Submitted):
			status = "Submitted"
		case int32(trdcommon.OrderState_OrderState_FilledAll):
			status = "Filled"
		case int32(trdcommon.OrderState_OrderState_PartialDone):
			status = "Partial"
		case int32(trdcommon.OrderState_OrderState_Canceled):
			status = "Canceled"
		}

		fmt.Printf("  %-12d %-10s %-8s %-6s %-6s %-8.0f %-8.2f %-10.0f\n",
			o.OrderID, o.Code, side, orderType, status, o.Qty, o.Price, o.DealtQty)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
