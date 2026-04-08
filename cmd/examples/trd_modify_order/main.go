// Example: ModifyOrder - 修改訂單
//
// This example demonstrates how to use the ModifyOrder API to modify
// an existing order (price, quantity, or cancel).
//
// Usage:
//   go run main.go [acc_id] [order_id]

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

	fmt.Println("=== ModifyOrder Example / 修改訂單示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get parameters / 獲取參數
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [acc_id] [order_id]")
		fmt.Println("\nThis example shows how to modify an order:")
		fmt.Println("  - Change price from 350.00 to 360.00")
		fmt.Println("  - Change quantity from 100 to 200")
		fmt.Println("\nNote: DRY RUN mode - remove to actually modify")
		return
	}

	accID, _ := strconv.ParseUint(os.Args[1], 10, 64)
	orderID, _ := strconv.ParseUint(os.Args[2], 10, 64)

	fmt.Printf("Account ID / 賬戶ID: %d\n", accID)
	fmt.Printf("Order ID / 訂單ID:   %d\n\n", orderID)

	// Construct modify request / 構建修改請求
	req := &trd.ModifyOrderRequest{
		AccID:         accID,
		TrdMarket:     int32(trdcommon.TrdMarket_TrdMarket_HK),
		OrderID:       orderID,
		ModifyOrderOp: int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Normal), // Normal modify / 普通修改
		Qty:           200.0,                                               // New quantity / 新數量
		Price:         360.00,                                              // New price / 新價格
	}

	fmt.Println("📝 Modify Order Details / 修改訂單詳情:")
	fmt.Printf("  New Quantity / 新數量: %.0f\n", req.Qty)
	fmt.Printf("  New Price / 新價格:    %.2f\n", req.Price)
	fmt.Printf("  Modify Type / 修改類型: Normal / 普通修改\n")
	fmt.Println()

	// DRY RUN - Remove to actually modify / 乾運行 - 移除以實際修改
	dryRun := true

	if dryRun {
		fmt.Println("⚠️  DRY RUN MODE / 乾運行模式")
		fmt.Println("   Order was NOT actually modified")
		fmt.Println("   Remove 'dryRun = true' to modify real orders")
		fmt.Println()
		fmt.Println("   To modify order, uncomment this code:")
		fmt.Println("   err := trd.ModifyOrder(cli, req)")
		fmt.Println("   if err != nil {")
		fmt.Println("       log.Printf(\"ModifyOrder failed\", err)")
		fmt.Println("       return")
		fmt.Println("   }")
		fmt.Println("   fmt.Println(\"Order modified successfully\")")
	} else {
		// Live modification / 實際修改
		err := trd.ModifyOrder(cli, req)
		if err != nil {
			log.Printf("ModifyOrder failed: %v", err)
			return
		}
		fmt.Println("✅ Order modified successfully! / 訂單修改成功！")
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
