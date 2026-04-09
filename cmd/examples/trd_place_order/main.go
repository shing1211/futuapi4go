// Example: PlaceOrder - 下單交易
//
// This example demonstrates how to use the PlaceOrder API to place
// a buy or sell order in the trading account.
//
// Usage:
//   go run main.go [acc_id]
//
// Note: This is a DRY RUN example - it shows how to construct the order
// but does not actually submit it to avoid accidental trades.
// Remove the dry run flag to place real orders.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	fmt.Println("=== PlaceOrder Example / 下單交易示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get account ID from argument or use first account / 從參數獲取賬戶ID或使用第一個賬戶
	var accID uint64
	if len(os.Args) > 1 {
		id, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalf("Invalid account ID: %v", err)
		}
		accID = id
	} else {
		// Get first account / 獲取第一個賬戶
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

	// Construct order / 構建訂單
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)          // Buy / 買入
	orderType := int32(trdcommon.OrderType_OrderType_Normal) // Normal order / 普通單

	// Example: Buy 100 shares of Tencent at 350.00 / 以350.00買入騰訊100股
	req := &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		Code:      "00700",
		TrdSide:   trdSide,
		OrderType: orderType,
		Price:     350.00,
		Qty:       100.0,
	}

	fmt.Println("📝 Order Details / 訂單詳情:")
	fmt.Printf("  Action / 操作:        BUY / 買入\n")
	fmt.Printf("  Code / 代碼:          %s\n", req.Code)
	fmt.Printf("  Quantity / 數量:      %.0f\n", req.Qty)
	fmt.Printf("  Price / 價格:         %.2f\n", req.Price)
	fmt.Printf("  Order Type / 訂單類型: Normal / 普通單\n")
	fmt.Printf("  Price Type / 價格類型: Limit / 限價單\n")
	fmt.Println()

	// DRY RUN - Comment out to actually place the order / 乾運行 - 取消註釋以實際下單
	dryRun := true

	if dryRun {
		fmt.Println("⚠️  DRY RUN MODE / 乾運行模式")
		fmt.Println("   Order was NOT actually submitted")
		fmt.Println("   Remove 'dryRun = true' to place real orders")
		fmt.Println()
		fmt.Println("   To place order, uncomment this code:")
		fmt.Println("   resp, err := trd.PlaceOrder(cli, req)")
		fmt.Println("   if err != nil {")
		fmt.Println("       log.Printf(\"PlaceOrder failed\", err)")
		fmt.Println("       return")
		fmt.Println("   }")
		fmt.Println("   fmt.Printf(\"Order placed, OrderID\", resp.OrderID)")
	} else {
		// Live order placement / 實際下單
		resp, err := trd.PlaceOrder(cli, req)
		if err != nil {
			log.Printf("PlaceOrder failed: %v", err)
			return
		}
		fmt.Printf("✅ Order placed successfully! / 下單成功！\n")
		fmt.Printf("   OrderID / 訂單ID: %d\n", resp.OrderID)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

