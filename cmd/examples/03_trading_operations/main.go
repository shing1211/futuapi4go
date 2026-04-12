// Trading Operations Examples
//
// This example demonstrates how to use trading APIs:
// - GetAccList: Get account list
// - UnlockTrade: Unlock trading
// - GetFunds: Get account funds
// - PlaceOrder: Place orders
// - GetOrderList: Query orders
// - ModifyOrder: Modify orders
// - GetPositionList: Get positions
// - GetOrderFee: Get order fees
// - GetMarginRatio: Get margin ratio
// - GetMaxTrdQtys: Get max trade quantities
// - GetHistoryOrderList: Get history orders
// - GetOrderOrderFillList: Get order fills
//
// Usage:
//   go run main.go
//
// Note: Trading APIs require UnlockTrade to be called first!

package main

import (
	"fmt"
	"log"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	// Create and connect client
	cli := futuapi.New()
	defer cli.Close()

	addr := "127.0.0.1:11111"
	fmt.Printf("=== Connecting to %s ===\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✓ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Define markets
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	trdCategory := int32(trdcommon.TrdCategory_TrdCategory_Security)

	// 1. Get Account List
	fmt.Println("=== 1. Account List (GetAccList) ===")
	accResp, err := trd.GetAccList(cli, trdCategory, false)
	if err != nil {
		log.Printf("GetAccList failed: %v", err)
	} else {
		fmt.Printf("  Found %d accounts\n", len(accResp.AccList))
		for i, acc := range accResp.AccList {
			trdEnv := "Simulation"
			if acc.TrdEnv == int32(trdcommon.TrdEnv_TrdEnv_Real) {
				trdEnv = "Real"
			}
			fmt.Printf("  [%d] AccID=%d | TrdEnv=%s | AccType=%d | Status=%d\n",
				i+1, acc.AccID, trdEnv, acc.AccType, acc.AccStatus)
		}
	}
	fmt.Println()

	// For the rest of examples, use first account if available
	var accID uint64 = 123456789 // Default mock value
	if accResp != nil && len(accResp.AccList) > 0 {
		accID = accResp.AccList[0].AccID
	}

	// 2. Unlock Trade (REQUIRED for trading)
	fmt.Println("=== 2. Unlock Trading (UnlockTrade) ===")
	fmt.Println("  ⚠️  In real environment, you need to provide trade password")
	fmt.Println("  For simulator, this will return success")

	unlockReq := &trd.UnlockTradeRequest{Unlock: true, PwdMD5: "your_trade_password"}
	unlockErr := trd.UnlockTrade(cli, unlockReq)
	if unlockErr != nil {
		log.Printf("UnlockTrade failed: %v (expected in simulator)", unlockErr)
	} else {
		fmt.Println("  ✓ Trade unlocked successfully")
	}
	fmt.Println()

	// 3. Get Funds
	fmt.Println("=== 3. Account Funds (GetFunds) ===")
	fundsReq := &trd.GetFundsRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
	}

	fundsResp, err := trd.GetFunds(cli, fundsReq)
	if err != nil {
		log.Printf("GetFunds failed: %v", err)
	} else {
		funds := fundsResp.Funds
		fmt.Printf("  Account: %d\n", accID)
		fmt.Printf("  Total Assets:        %.2f\n", funds.TotalAssets)
		fmt.Printf("  Cash:                %.2f\n", funds.Cash)
		fmt.Printf("  Market Value:        %.2f\n", funds.MarketVal)
		fmt.Printf("  Available Funds:     %.2f\n", funds.AvailableFunds)
		fmt.Printf("  Frozen Cash:         %.2f\n", funds.FrozenCash)
		fmt.Printf("  Power:               %.2f\n", funds.Power)
	}
	fmt.Println()

	// 4. Get Positions
	fmt.Println("=== 4. Position List (GetPositionList) ===")
	posReq := &trd.GetPositionListRequest{
		AccID:     accID,
		TrdMarket: 0, // 0 means all markets
	}

	posResp, err := trd.GetPositionList(cli, posReq)
	if err != nil {
		log.Printf("GetPositionList failed: %v", err)
	} else {
		fmt.Printf("  Account: %d\n", accID)
		fmt.Printf("  Found %d positions\n", len(posResp.PositionList))

		if len(posResp.PositionList) > 0 {
			fmt.Printf("  %-10s %-10s %-10s %-10s %-12s %-12s %-12s %-12s\n",
				"Code", "Name", "Qty", "AvailQty", "CostPrice", "Price", "Val", "PlVal")
			for _, pos := range posResp.PositionList {
				fmt.Printf("  %-10s %-10s %-10.0f %-10.0f %-10.4f %-12.2f %-12.2f %-12.2f\n",
					pos.Code, pos.Name, pos.Qty, pos.CanSellQty,
					pos.CostPrice, pos.Price, pos.Val, pos.PlVal)
			}
		} else {
			fmt.Println("  No positions found (this is normal for simulator)")
		}
	}
	fmt.Println()

	// 5. Place Order (Buy)
	fmt.Println("=== 5. Place Order - BUY (PlaceOrder) ===")

	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
	orderType := int32(trdcommon.OrderType_OrderType_Normal)
	qty := float64(100)
	price := 350.00

	placeReq := &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		Code:      "00700",
		TrdSide:   trdSide,
		OrderType: orderType,
		Price:     price,
		Qty:       qty,
	}

	placeResp, err := trd.PlaceOrder(cli, placeReq)
	var orderID uint64
	if err != nil {
		log.Printf("PlaceOrder failed: %v", err)
	} else {
		orderID = placeResp.OrderID
		fmt.Printf("  ✓ Order placed successfully!\n")
		fmt.Printf("  OrderID: %d\n", orderID)
		fmt.Printf("  Action: BUY %.0f shares of %s at %.2f\n",
			qty, "00700", price)
	}
	fmt.Println()

	// 6. Get Order List
	fmt.Println("=== 6. Order List (GetOrderList) ===")
	orderReq := &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
	}

	orderResp, err := trd.GetOrderList(cli, orderReq)
	if err != nil {
		log.Printf("GetOrderList failed: %v", err)
	} else {
		fmt.Printf("  Account: %d\n", accID)
		fmt.Printf("  Found %d orders\n", len(orderResp.OrderList))

		if len(orderResp.OrderList) > 0 {
			fmt.Printf("  %-12s %-10s %-10s %-12s %-10s %-10s %-12s\n",
				"OrderID", "Code", "Side", "Type", "Qty", "Price", "Status")
			for _, o := range orderResp.OrderList {
				side := "BUY"
				if o.TrdSide == int32(trdcommon.TrdSide_TrdSide_Sell) {
					side = "SELL"
				}

				status := "Unknown"
				switch o.OrderStatus {
				case int32(trdcommon.OrderStatus_OrderStatus_Submitting):
					status = "Submitting"
				case int32(trdcommon.OrderStatus_OrderStatus_Submitted):
					status = "Submitted"
				case int32(trdcommon.OrderStatus_OrderStatus_Filled_All):
					status = "Filled"
				case int32(trdcommon.OrderStatus_OrderStatus_Filled_Part):
					status = "Partial"
				case int32(trdcommon.OrderStatus_OrderStatus_Cancelled_All):
					status = "Cancelled"
				}

				fmt.Printf("  %-12d %-10s %-10s %-12d %-10.0f %-10.2f %-12s\n",
					o.OrderID, o.Code, side, o.OrderType, o.Qty, o.Price, status)
			}
		} else {
			fmt.Println("  No orders found")
		}
	}
	fmt.Println()

	// 7. Modify Order
	fmt.Println("=== 7. Modify Order (ModifyOrder) ===")
	if orderID > 0 {
		modifyReq := &trd.ModifyOrderRequest{
			AccID:         accID,
			TrdMarket:     hkMarket,
			OrderID:       orderID,
			ModifyOrderOp: int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Normal),
			Qty:           200,
			Price:         360.00,
		}

		modifyResp, modifyErr := trd.ModifyOrder(cli, modifyReq)
		if modifyErr != nil {
			log.Printf("ModifyOrder failed: %v", modifyErr)
		} else {
			fmt.Println("  ✓ Order modified successfully!")
			fmt.Printf("  OrderID: %d, OrderIDEx: %s\n", modifyResp.OrderID, modifyResp.OrderIDEx)
			fmt.Printf("  New Qty: 200, New Price: 360.00\n")
		}
	} else {
		fmt.Println("  ⚠️  No order to modify (place order first)")
	}
	fmt.Println()

	// 8. Get Order Fills/Executions
	fmt.Println("=== 8. Order Fills (GetOrderOrderFillList) ===")
	fillReq := &trd.GetOrderFillListRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
	}

	fillResp, err := trd.GetOrderFillList(cli, fillReq)
	if err != nil {
		log.Printf("GetOrderFillList failed: %v", err)
	} else {
		fmt.Printf("  Account: %d\n", accID)
		fmt.Printf("  Found %d fills\n", len(fillResp.OrderFillList))

		if len(fillResp.OrderFillList) > 0 {
			fmt.Printf("  %-12s %-12s %-10s %-12s %-10s\n",
				"FillID", "OrderID", "Code", "Qty", "Price")
			for _, f := range fillResp.OrderFillList {
				fmt.Printf("  %-12d %-12d %-10s %-12.0f %-10.2f\n",
					f.FillID, f.OrderID, f.Code, f.Qty, f.Price)
			}
		} else {
			fmt.Println("  No fills found")
		}
	}
	fmt.Println()

	// 9. Get Order Fee
	fmt.Println("=== 9. Order Fees (GetOrderFee) ===")
	feeReq := &trd.GetOrderFeeRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
	}

	feeResp, err := trd.GetOrderFee(cli, feeReq)
	if err != nil {
		log.Printf("GetOrderFee failed: %v", err)
	} else {
		fmt.Printf("  Account: %d\n", accID)
		fmt.Printf("  Found %d fee records\n", len(feeResp.OrderFeeList))

		if len(feeResp.OrderFeeList) > 0 {
			for _, fee := range feeResp.OrderFeeList {
				fmt.Printf("  OrderID: %s | Fee: %.2f\n", fee.OrderIDEx, fee.FeeAmount)
			}
		} else {
			fmt.Println("  No fees found")
		}
	}
	fmt.Println()

	// 10. Get Margin Ratio
	fmt.Println("=== 10. Margin Ratio (GetMarginRatio) ===")
	code007 := "00700"
	sec007 := &qotcommon.Security{Market: &hkMarket, Code: &code007}
	marginReq := &trd.GetMarginRatioRequest{
		AccID:        accID,
		TrdMarket:    hkMarket,
		SecurityList: []*qotcommon.Security{sec007},
	}

	marginResp, err := trd.GetMarginRatio(cli, marginReq)
	if err != nil {
		log.Printf("GetMarginRatio failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s\n", "00700")
		if len(marginResp.MarginRatioInfoList) > 0 {
			info := marginResp.MarginRatioInfoList[0]
			fmt.Printf("  IM Long Ratio:   %.4f\n", info.ImLongRatio)
			fmt.Printf("  IM Short Ratio:  %.4f\n", info.ImShortRatio)
		}
	}
	fmt.Println()

	// 11. Get Max Trade Quantities
	fmt.Println("=== 11. Max Trade Quantities (GetMaxTrdQtys) ===")
	maxQtyReq := &trd.GetMaxTrdQtysRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		Code:      "00700",
		Price:     350.00,
		OrderType: int32(trdcommon.OrderType_OrderType_Normal),
	}

	maxQtyResp, err := trd.GetMaxTrdQtys(cli, maxQtyReq)
	if err != nil {
		log.Printf("GetMaxTrdQtys failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s @ %.2f\n", "00700", 350.00)
		fmt.Printf("  Max Cash Buy:        %.0f\n", maxQtyResp.MaxTrdQtys.MaxCashBuy)
		fmt.Printf("  Max Cash+Margin Buy: %.0f\n", maxQtyResp.MaxTrdQtys.MaxCashAndMarginBuy)
		fmt.Printf("  Max Position Sell:   %.0f\n", maxQtyResp.MaxTrdQtys.MaxPositionSell)
		fmt.Printf("  Max Short Sell:      %.0f\n", maxQtyResp.MaxTrdQtys.MaxSellShort)
	}
	fmt.Println()

	// 12. Get History Orders
	fmt.Println("=== 12. History Orders (GetHistoryOrderList) ===")
	beginTimeStr := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endTimeStr := time.Now().Format("2006-01-02")

	histReq := &trd.GetHistoryOrderListRequest{
		AccID:     accID,
		TrdMarket: hkMarket,
		FilterConditions: &trdcommon.TrdFilterConditions{
			BeginTime: &beginTimeStr,
			EndTime:   &endTimeStr,
		},
	}

	histResp, err := trd.GetHistoryOrderList(cli, histReq)
	if err != nil {
		log.Printf("GetHistoryOrderList failed: %v", err)
	} else {
		fmt.Printf("  Period: %s to %s\n", beginTimeStr, endTimeStr)
		fmt.Printf("  Found %d historical orders\n", len(histResp.OrderList))

		if len(histResp.OrderList) > 0 {
			fmt.Printf("  Showing first 3:\n")
			count := 3
			if len(histResp.OrderList) < count {
				count = len(histResp.OrderList)
			}
			for i := 0; i < count; i++ {
				o := histResp.OrderList[i]
				side := "BUY"
				if o.TrdSide != nil && *o.TrdSide == int32(trdcommon.TrdSide_TrdSide_Sell) {
					side = "SELL"
				}
				codeStr := "N/A"
				if o.Code != nil {
					codeStr = *o.Code
				}
				nameStr := "N/A"
				if o.Name != nil {
					nameStr = *o.Name
				}
				fmt.Printf("    %s (%s) | %s | Qty=%.0f | Price=%.2f | Status=%d\n",
					codeStr, nameStr, side, o.GetQty(), o.GetPrice(), o.GetOrderStatus())
			}
		}
	}
	fmt.Println()

	fmt.Println("=== Trading Examples Complete ===")
	fmt.Println("⚠️  Note: Simulator returns mock data")
	fmt.Println("💡  Tip: In real environment, always:")
	fmt.Println("   1. Unlock trade with correct password")
	fmt.Println("   2. Check account balance before trading")
	fmt.Println("   3. Verify order details before placing")
	fmt.Println("   4. Monitor order status after placement")
}

// Helper functions
func ptrStr(s string) *string {
	return &s
}

func ptrInt32(v int32) *int32 {
	return &v
}

func ptrInt64(v int64) *int64 {
	return &v
}

func ptrFloat64(v float64) *float64 {
	return &v
}

// Ensure imports are used
var _ = time.Now
var _ = trdcommon.TrdCategory_TrdCategory_Security
