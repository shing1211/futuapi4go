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

// Example: Market Making Strategy / 做市策略
//
// This example demonstrates a basic market making algorithm that
// simultaneously places bid and ask orders around the current price
// to profit from the bid-ask spread.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// Market making parameters / 做市參數
const (
	SpreadPercent = 0.005 // 0.5% spread / 0.5% 價差
	OrderQty      = 100   // Shares per side / 每邊股數
	Layers        = 3     // Number of layers per side / 每邊層數
	LayerSpread   = 0.002 // Additional spread per layer / 每層額外價差
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== Market Making Strategy / 做市策略 ===")
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
	fmt.Printf("   Spread / 價差:          %.2f%%\n", SpreadPercent*100)
	fmt.Printf("   Order Qty / 訂單數量:    %d shares\n", OrderQty)
	fmt.Printf("   Layers / 層數:          %d per side\n", Layers)
	fmt.Println()

	// Get current market data / 獲取當前市場數據
	fmt.Println("📈 Fetching market data / 正在獲取市場數據...")
	currentPrice, bidPrice, askPrice := getMarketData(cli, security)
	if currentPrice <= 0 {
		fmt.Println("❌ Failed to get market data")
		return
	}

	fmt.Printf("✅ Current Market / 當前市場:\n")
	fmt.Printf("   Mid Price / 中間價: HK$%.2f\n", currentPrice)
	fmt.Printf("   Bid / 買入價:      HK$%.2f\n", bidPrice)
	fmt.Printf("   Ask / 賣出價:      HK$%.2f\n\n", askPrice)

	// Calculate order prices / 計算訂單價格
	fmt.Println("📐 Calculating Order Prices / 正在計算訂單價格...")
	bidLevels, askLevels := calculateMarketMakingLevels(currentPrice)

	fmt.Printf("   Bid Levels (Buy Orders) / 買入級別:\n")
	for i, price := range bidLevels {
		fmt.Printf("     Layer %d: HK$%.2f (spread: %.2f%%)\n",
			i+1, price, (currentPrice-price)/currentPrice*100)
	}

	fmt.Printf("\n   Ask Levels (Sell Orders) / 賣出級別:\n")
	for i, price := range askLevels {
		fmt.Printf("     Layer %d: HK$%.2f (spread: %.2f%%)\n",
			i+1, price, (price-currentPrice)/currentPrice*100)
	}
	fmt.Println()

	// Calculate expected profit / 計算預期利潤
	expectedProfit := float64(Layers) * OrderQty * (askLevels[0] - bidLevels[0])
	fmt.Printf("💰 Expected Profit per Cycle / 每輪預期利潤: HK$%.2f\n\n", expectedProfit)

	// Place market making orders / 下達做市訂單
	dryRun := true // Set to false for live trading / 設置為false以進行真實交易

	fmt.Println("📝 Placing Market Making Orders / 正在下達做市訂單...")
	bidOrders := placeOrders(cli, accID, security, bidLevels, "Buy", OrderQty, dryRun)
	askOrders := placeOrders(cli, accID, security, askLevels, "Sell", OrderQty, dryRun)

	fmt.Printf("\n📊 Order Summary / 訂單摘要:\n")
	fmt.Printf("   Bid Orders / 買入訂單:   %d\n", bidOrders)
	fmt.Printf("   Ask Orders / 賣出訂單:   %d\n", askOrders)
	fmt.Printf("   Total Spread / 總價差:   HK$%.2f\n", askLevels[0]-bidLevels[0])

	if dryRun {
		fmt.Println("\n⚠️  DRY RUN MODE - No real orders placed")
		fmt.Println("   模擬運行模式 - 沒有下達真實訂單")
	}

	fmt.Println("\n=== Strategy Complete / 策略完成 ===")
}

// getMarketData retrieves current market prices / 獲取當前市場價格
func getMarketData(cli *futuapi.Client, security *qotcommon.Security) (mid, bid, ask float64) {
	// Get basic quote / 獲取基本報價
	quotes, err := qot.GetBasicQot(cli, []*qotcommon.Security{security})
	if err != nil {
		log.Printf("GetBasicQot failed: %v", err)
		return 0, 0, 0
	}

	if len(quotes) == 0 {
		return 0, 0, 0
	}

	mid = quotes[0].CurPrice

	// Get order book for better bid/ask / 獲取買賣盤以獲得更好的買入/賣出價
	obReq := &qot.GetOrderBookRequest{
		Security: security,
		Num:      1,
	}
	obResp, err := qot.GetOrderBook(cli, obReq)
	if err != nil {
		// Fallback to basic quote / 回退到基本報價
		bid = mid * 0.999
		ask = mid * 1.001
		return
	}

	if len(obResp.OrderBookBidList) > 0 {
		bid = obResp.OrderBookBidList[0].Price
	} else {
		bid = mid * 0.999
	}

	if len(obResp.OrderBookAskList) > 0 {
		ask = obResp.OrderBookAskList[0].Price
	} else {
		ask = mid * 1.001
	}

	return
}

// calculateMarketMakingLevels calculates bid and ask price levels
// 計算買入和賣出價格級別
func calculateMarketMakingLevels(midPrice float64) ([]float64, []float64) {
	bidLevels := make([]float64, 0, Layers)
	askLevels := make([]float64, 0, Layers)

	for i := 0; i < Layers; i++ {
		spread := SpreadPercent + float64(i)*LayerSpread
		bidPrice := midPrice * (1 - spread)
		askPrice := midPrice * (1 + spread)

		// Round to 2 decimal places / 四捨五入到小數點後2位
		bidLevels = append(bidLevels, roundTo2(bidPrice))
		askLevels = append(askLevels, roundTo2(askPrice))
	}

	return bidLevels, askLevels
}

// placeOrders places orders at specified price levels / 在指定價格級別下單
func placeOrders(cli *futuapi.Client, accID uint64, security *qotcommon.Security,
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
	}

	return orderCount
}

// roundTo2 rounds to 2 decimal places / 四捨五入到小數點後2位
func roundTo2(val float64) float64 {
	return float64(int(val*100+0.5)) / 100
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
