// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Example: Grid Trading Strategy / 網格交易策略
//
// This example demonstrates a grid trading algorithm that places
// buy and sell orders at predetermined price levels.
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// Grid configuration / 網格配置
const (
	GridCount    = 10  // Number of grids / 網格數量
	GridSize     = 1.0 // Grid size in HKD / 每格大小（港幣）
	OrderPerGrid = 100 // Shares per grid / 每格股數
	MaxOrders    = 20  // Maximum open orders / 最大掛單數
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== Grid Trading Strategy / 網格交易策略 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Configuration / 配置
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)
	code := "HSImain" // HSI Futures Main Contract / 恆生指數期貨
	security := &qotcommon.Security{
		Market: &hkFutureMarket,
		Code:   &code,
	}

	// Get account ID / 獲取賬戶ID
	accID := getAccountID(cli)
	if accID == 0 {
		fmt.Println("No trading account found / 沒有找到交易賬戶")
		return
	}

	fmt.Printf("📊 Strategy Configuration / 策略配置:\n")
	fmt.Printf("   Security / 股票:        %s\n", code)
	fmt.Printf("   Grids / 網格數:         %d\n", GridCount)
	fmt.Printf("   Grid Size / 每格大小:   HK$%.2f\n", GridSize)
	fmt.Printf("   Per Grid / 每格數量:    %d shares\n", OrderPerGrid)
	fmt.Println()

	// Get current price / 獲取當前價格
	fmt.Println("💰 Fetching current price / 正在獲取當前價格...")
	currentPrice := getCurrentPrice(cli, security)
	if currentPrice <= 0 {
		fmt.Println("❌ Failed to get current price")
		return
	}
	fmt.Printf("✅ Current Price / 當前價格: HK$%.2f\n\n", currentPrice)

	// Calculate grid levels / 計算網格級別
	fmt.Println("📐 Calculating Grid Levels / 正在計算網格級別...")
	buyLevels, sellLevels := calculateGridLevels(currentPrice)

	fmt.Printf("   Buy Grids / 買入網格: %d levels\n", len(buyLevels))
	for i, price := range buyLevels {
		fmt.Printf("     Buy %d: HK$%.2f\n", i+1, price)
	}

	fmt.Printf("\n   Sell Grids / 賣出網格: %d levels\n", len(sellLevels))
	for i, price := range sellLevels {
		fmt.Printf("     Sell %d: HK$%.2f\n", i+1, price)
	}
	fmt.Println()

	// Place orders / 下單
	dryRun := true // Set to false for live trading / 設置為false以進行真實交易

	fmt.Println("📝 Placing Orders / 正在下單...")

	buyOrders := placeGridOrders(cli, accID, security, buyLevels, "Buy", OrderPerGrid, dryRun)
	sellOrders := placeGridOrders(cli, accID, security, sellLevels, "Sell", OrderPerGrid, dryRun)

	fmt.Printf("\n📊 Order Summary / 訂單摘要:\n")
	fmt.Printf("   Buy Orders / 買入訂單:   %d\n", buyOrders)
	fmt.Printf("   Sell Orders / 賣出訂單:  %d\n", sellOrders)
	fmt.Printf("   Total Orders / 總訂單:   %d\n", buyOrders+sellOrders)

	fmt.Println("\n=== Strategy Complete / 策略完成 ===")
	if dryRun {
		fmt.Println("⚠️  This was a DRY RUN. No real orders were placed.")
		fmt.Println("   這是模擬運行。沒有下達真實訂單。")
		fmt.Println("   Set dryRun=false in the code to place real orders.")
		fmt.Println("   在代碼中設置 dryRun=false 以下達真實訂單。")
	}
}

// getCurrentPrice retrieves the current market price / 獲取當前市場價格
func getCurrentPrice(cli *futuapi.Client, security *qotcommon.Security) float64 {
	quotes, err := qot.GetBasicQot(context.Background(),cli, []*qotcommon.Security{security})
	if err != nil {
		log.Printf("GetBasicQot failed: %v", err)
		return 0
	}

	if len(quotes) == 0 {
		return 0
	}

	return quotes[0].CurPrice
}

// calculateGridLevels calculates buy and sell price levels / 計算買入和賣出價格級別
func calculateGridLevels(currentPrice float64) ([]float64, []float64) {
	buyLevels := make([]float64, 0, GridCount)
	sellLevels := make([]float64, 0, GridCount)

	// Buy grids below current price / 當前價格下方的買入網格
	for i := 1; i <= GridCount; i++ {
		price := currentPrice - (float64(i) * GridSize)
		if price > 0 {
			buyLevels = append(buyLevels, math.Round(price*100)/100)
		}
	}

	// Sell grids above current price / 當前價格上方的賣出網格
	for i := 1; i <= GridCount; i++ {
		price := currentPrice + (float64(i) * GridSize)
		sellLevels = append(sellLevels, math.Round(price*100)/100)
	}

	return buyLevels, sellLevels
}

// placeGridOrders places orders at specified price levels
// 在指定價格級別下單
func placeGridOrders(cli *futuapi.Client, accID uint64, security *qotcommon.Security,
	prices []float64, side string, qty int, dryRun bool) int {

	orderCount := 0
	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
	if side == "Sell" {
		trdSide = int32(trdcommon.TrdSide_TrdSide_Sell)
	}

	for _, price := range prices {
		if dryRun {
			fmt.Printf("   [DRY RUN] %s %d @ HK$%.2f\n", side, qty, price)
			orderCount++
			continue
		}

		// Place real order / 下達真實訂單
		req := &trd.PlaceOrderRequest{
			AccID:     accID,
			TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
			Code:      security.GetCode(),
			TrdSide:   trdSide,
			OrderType: int32(trdcommon.OrderType_OrderType_Normal),
			Price:     price,
			Qty:       float64(qty),
		}

		resp, err := trd.PlaceOrder(cli, req)
		if err != nil {
			fmt.Printf("   ❌ Failed to place %s order @ HK$%.2f: %v\n", side, price, err)
			continue
		}

		fmt.Printf("   ✅ %s %d @ HK$%.2f | OrderID: %d\n", side, qty, price, resp.OrderID)
		orderCount++

		// Avoid rate limiting / 避免觸發速率限制
		// time.Sleep(100 * time.Millisecond)
	}

	return orderCount
}

// getAccountID retrieves trading account ID / 獲取交易賬戶ID
func getAccountID(cli *futuapi.Client) uint64 {
	if len(os.Args) > 1 {
		id, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err == nil {
			return id
		}
	}

	accResp, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		log.Printf("GetAccList failed: %v", err)
		return 0
	}

	if len(accResp.AccList) == 0 {
		return 0
	}

	return accResp.AccList[0].AccID
}
