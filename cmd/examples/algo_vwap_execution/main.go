// Example: VWAP Execution Strategy / 成交量加權平均價格執行策略
//
// This example demonstrates an intelligent order execution algorithm that
// uses Volume-Weighted Average Price (VWAP) to minimize market impact
// when executing large orders.
//
// Strategy:
// - Calculate VWAP from historical data
// - Split large order into smaller child orders
// - Execute when price is below VWAP (for buys) or above VWAP (for sells)
// - Track execution quality vs VWAP benchmark
//
// Usage:
//   go run main.go [account_id] [total_shares]
//
// Note: DRY RUN by default for safety.

package main

import (
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

// VWAP execution parameters / VWAP執行參數
const (
	HistoryDays    = 10  // Days to calculate VWAP / 計算VWAP的天數
	ChildOrderPct  = 10  // % of total per child order / 每子單百分比
	PriceThreshold = 0.3 // Execute if price is X% better than VWAP / 價格優於VWAP X%時執行
	MaxAttempts    = 5   // Maximum execution attempts / 最大執行次數
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== VWAP Execution Strategy / VWAP執行策略 ===")
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

	// Get total shares to execute from command line or default
	totalShares := 1000
	if len(os.Args) > 2 {
		shares, err := strconv.Atoi(os.Args[2])
		if err == nil && shares > 0 {
			totalShares = shares
		}
	}

	fmt.Printf("📊 Execution Configuration / 執行配置:\n")
	fmt.Printf("   Security / 股票:         %s\n", code)
	fmt.Printf("   Total Shares / 總股數:    %d\n", totalShares)
	fmt.Printf("   History Days / 歷史天數:  %d\n", HistoryDays)
	fmt.Printf("   Child Order Size / 子單大小: %d shares (%d%%)\n",
		totalShares*ChildOrderPct/100, ChildOrderPct)
	fmt.Printf("   Price Threshold / 價格閾值: %.1f%%\n", PriceThreshold)
	fmt.Println()

	// Fetch historical data / 獲取歷史數據
	fmt.Println("📈 Fetching Historical Data / 正在獲取歷史數據...")
	klData := fetchKLData(cli, security, HistoryDays)
	if len(klData) == 0 {
		fmt.Println("❌ Failed to fetch historical data")
		return
	}
	fmt.Printf("✅ Retrieved %d days of data\n\n", len(klData))

	// Calculate VWAP / 計算VWAP
	fmt.Println("📐 Calculating VWAP / 正在計算VWAP...")
	vwap := calculateVWAP(klData)
	fmt.Printf("   VWAP / 成交量加權平均價: HK$%.2f\n\n", vwap)

	// Get current price / 獲取當前價格
	fmt.Println("💰 Getting Current Price / 正在獲取當前價格...")
	currentPrice := getCurrentPrice(cli, security)
	if currentPrice <= 0 {
		fmt.Println("❌ Failed to get current price")
		return
	}
	fmt.Printf("   Current Price / 當前價格: HK$%.2f\n", currentPrice)

	// Compare with VWAP / 與VWAP比較
	priceDiff := (currentPrice - vwap) / vwap * 100
	fmt.Printf("   vs VWAP / 對比VWAP: %.2f%%\n", priceDiff)

	if priceDiff < 0 {
		fmt.Println("   ✅ Price is BELOW VWAP - Good time to BUY")
		fmt.Println("   ✅ 價格低於VWAP - 買入的好時機")
	} else {
		fmt.Println("   ⚠️  Price is ABOVE VWAP - Consider waiting for better price")
		fmt.Println("   ⚠️  價格高於VWAP - 考慮等待更好的價格")
	}
	fmt.Println()

	// Calculate child orders / 計算子單
	numChildOrders := int(math.Ceil(float64(totalShares) / float64(totalShares*ChildOrderPct/100)))
	childOrderSize := totalShares * ChildOrderPct / 100

	fmt.Printf("📋 Execution Plan / 執行計劃:\n")
	fmt.Printf("   Total Shares / 總股數:       %d\n", totalShares)
	fmt.Printf("   Child Orders / 子單數:       %d\n", numChildOrders)
	fmt.Printf("   Size per Order / 每單大小:    %d shares\n", childOrderSize)
	fmt.Println()

	// Execute orders / 執行訂單
	dryRun := true // Set to false for live execution / 設置為false以進行真實執行

	fmt.Println("🚀 Executing Orders / 正在執行訂單...")
	executeVWAPOrders(cli, accID, security, totalShares, childOrderSize, numChildOrders,
		vwap, currentPrice, dryRun)

	fmt.Println("\n=== Execution Complete / 執行完成 ===")
}

// KLData holds K-line data for VWAP calculation / VWAP計算的K線數據
type KLData struct {
	Close    float64
	Volume   int64
	Turnover float64
}

// fetchKLData retrieves historical K-line data / 獲取歷史K線數據
func fetchKLData(cli *futuapi.Client, security *qotcommon.Security, days int32) []KLData {
	req := &qot.GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_Forward),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    days,
	}

	resp, err := qot.GetKL(cli, req)
	if err != nil {
		log.Printf("GetKL failed: %v", err)
		return nil
	}

	data := make([]KLData, len(resp.KLList))
	for i, kl := range resp.KLList {
		data[i] = KLData{
			Close:    kl.ClosePrice,
			Volume:   kl.Volume,
			Turnover: kl.Turnover,
		}
	}

	return data
}

// calculateVWAP calculates Volume-Weighted Average Price / 計算成交量加權平均價格
func calculateVWAP(data []KLData) float64 {
	if len(data) == 0 {
		return 0
	}

	totalTurnover := 0.0
	totalVolume := int64(0)

	for _, d := range data {
		totalTurnover += d.Turnover
		totalVolume += d.Volume
	}

	if totalVolume == 0 {
		return 0
	}

	return totalTurnover / float64(totalVolume)
}

// getCurrentPrice retrieves current market price / 獲取當前市場價格
func getCurrentPrice(cli *futuapi.Client, security *qotcommon.Security) float64 {
	quotes, err := qot.GetBasicQot(cli, []*qotcommon.Security{security})
	if err != nil {
		log.Printf("GetBasicQot failed: %v", err)
		return 0
	}

	if len(quotes) == 0 {
		return 0
	}

	return quotes[0].CurPrice
}

// executeVWAPOrders executes VWAP-based orders / 執行基於VWAP的訂單
func executeVWAPOrders(cli *futuapi.Client, accID uint64, security *qotcommon.Security,
	totalShares, childOrderSize, numChildOrders int, vwap, currentPrice float64, dryRun bool) {

	filledShares := 0
	totalCost := 0.0
	attempts := 0

	for i := 0; i < numChildOrders && attempts < MaxAttempts; i++ {
		// Get current price for this attempt / 獲取此次嘗試的當前價格
		price := getCurrentPrice(cli, security)
		if price <= 0 {
			fmt.Printf("   ❌ Attempt %d: Failed to get price\n", i+1)
			attempts++
			continue
		}

		// Check if price is favorable / 檢查價格是否有利
		priceDiff := (price - vwap) / vwap * 100

		orderSize := childOrderSize
		if filledShares+orderSize > totalShares {
			orderSize = totalShares - filledShares
		}

		fmt.Printf("\n   📊 Child Order %d/%d / 子單 %d/%d:\n", i+1, numChildOrders, i+1, numChildOrders)
		fmt.Printf("      Current Price: HK$%.2f (vs VWAP: %.2f%%)\n", price, priceDiff)
		fmt.Printf("      Order Size: %d shares\n", orderSize)

		if dryRun {
			fmt.Printf("      [DRY RUN] Would BUY %d @ HK$%.2f\n", orderSize, price)
			filledShares += orderSize
			totalCost += price * float64(orderSize)
		} else {
			// Execute real order / 執行真實訂單
			req := &trd.PlaceOrderRequest{
				AccID:     accID,
				TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
				Code:      security.GetCode(),
				TrdSide:   int32(trdcommon.TrdSide_TrdSide_Buy),
				OrderType: int32(trdcommon.OrderType_OrderType_Normal),
				Price:     price,
				Qty:       float64(orderSize),
			}

			resp, err := trd.PlaceOrder(cli, req)
			if err != nil {
				fmt.Printf("      ❌ Order failed: %v\n", err)
				attempts++
				continue
			}

			fmt.Printf("      ✅ Order placed! OrderID: %d\n", resp.OrderID)
			filledShares += orderSize
			totalCost += price * float64(orderSize)
		}

		// Simulate waiting between orders / 模擬訂單之間的等待
		// In production, you'd add: time.Sleep(30 * time.Second)
	}

	// Execution summary / 執行摘要
	fmt.Println("\n   📋 Execution Summary / 執行摘要:")
	fmt.Printf("      Total Filled / 總成交:     %d/%d shares (%.1f%%)\n",
		filledShares, totalShares, float64(filledShares)/float64(totalShares)*100)

	if filledShares > 0 {
		avgPrice := totalCost / float64(filledShares)
		fmt.Printf("      Average Price / 平均價:  HK$%.2f\n", avgPrice)
		fmt.Printf("      VWAP / VWAP:            HK$%.2f\n", vwap)

		savings := (vwap - avgPrice) / vwap * 100
		if savings > 0 {
			fmt.Printf("      ✅ Execution Savings / 執行節省: HK$%.2f (%.2f%% better than VWAP)\n",
				(vwap-avgPrice)*float64(filledShares), savings)
		} else {
			fmt.Printf("      ⚠️  Execution Cost / 執行成本: HK$%.2f (%.2f%% worse than VWAP)\n",
				(avgPrice-vwap)*float64(filledShares), -savings)
		}
	}
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
