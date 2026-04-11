// Example: Simple Moving Average Crossover Strategy
// 示例：簡單移動平均線交叉策略
//
// This example demonstrates a basic algorithmic trading strategy using
// the Simple Moving Average (SMA) crossover technique.
//
// Strategy:
// - When short-term SMA crosses above long-term SMA → BUY signal
// - When short-term SMA crosses below long-term SMA → SELL signal
//
// Usage:
//   go run main.go [account_id]
//
// Note: This is a DRY RUN by default. Set dryRun=false to execute real trades.

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

// Strategy parameters / 策略參數
const (
	ShortPeriod = 5   // Short-term SMA period / 短期SMA週期
	LongPeriod  = 20  // Long-term SMA period / 長期SMA週期
	TradeQty    = 100 // Number of shares to trade / 交易股數
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== SMA Crossover Strategy / SMA交叉策略 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Configuration / 配置
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)
	code := "HSImain" // HSI Futures Main Contract
	security := &qotcommon.Security{
		Market: &hkFutureMarket,
		Code:   &code,
	}

	// Get account ID / 獲取賬戶ID
	accID := getAccountID(cli)
	if accID == 0 {
		fmt.Println("No trading account found")
		return
	}

	fmt.Printf("📊 Strategy Configuration / 策略配置:\n")
	fmt.Printf("   Security / 股票:        %s\n", code)
	fmt.Printf("   Short SMA Period / 短期: %d days\n", ShortPeriod)
	fmt.Printf("   Long SMA Period / 長期:  %d days\n", LongPeriod)
	fmt.Printf("   Trade Qty / 交易數量:    %d\n", TradeQty)
	fmt.Println()

	// Fetch K-line data / 獲取K線數據
	fmt.Println("📈 Fetching K-line data / 正在獲取K線數據...")
	klData := fetchKLData(cli, security, LongPeriod+10)
	if len(klData) < LongPeriod {
		fmt.Printf("❌ Insufficient data: need %d, got %d\n", LongPeriod, len(klData))
		return
	}
	fmt.Printf("✅ Retrieved %d K-line records\n\n", len(klData))

	// Calculate SMAs / 計算SMA
	fmt.Println("📐 Calculating Moving Averages / 正在計算移動平均線...")
	shortSMA := calculateSMA(klData, ShortPeriod)
	longSMA := calculateSMA(klData, LongPeriod)

	fmt.Printf("   Short SMA (%d) / 短期SMA: %.2f\n", ShortPeriod, shortSMA[len(shortSMA)-1])
	fmt.Printf("   Long SMA (%d) / 長期SMA:  %.2f\n\n", LongPeriod, longSMA[len(longSMA)-1])

	// Generate signal / 生成信號
	signal := generateSignal(shortSMA, longSMA)

	fmt.Println("🎯 Trading Signal / 交易信號:")
	switch signal {
	case 1:
		fmt.Println("   🟢 BUY Signal / 買入信號")
		fmt.Println("   Reason: Short SMA crossed above Long SMA")
		fmt.Println("   原因：短期SMA向上穿越長期SMA")
	case -1:
		fmt.Println("   🔴 SELL Signal / 賣出信號")
		fmt.Println("   Reason: Short SMA crossed below Long SMA")
		fmt.Println("   原因：短期SMA向下穿越長期SMA")
	default:
		fmt.Println("   ⚪ HOLD / 持有")
		fmt.Println("   Reason: No crossover detected")
		fmt.Println("   原因：未检测到交叉")
	}
	fmt.Println()

	// Execute trade / 執行交易
	dryRun := true // Set to false to execute real trades / 設置為false以執行真實交易

	if dryRun {
		fmt.Println("⚠️  DRY RUN MODE / 模擬運行模式")
		fmt.Println("   No real orders will be placed")
		fmt.Println("   不會下達真實訂單")
		fmt.Println()
		executeTrade(cli, accID, security, signal, TradeQty, dryRun)
	} else {
		fmt.Println("🚨 LIVE TRADING MODE / 真實交易模式")
		fmt.Println("   Real orders will be placed!")
		fmt.Println("   將下達真實訂單！")
		executeTrade(cli, accID, security, signal, TradeQty, dryRun)
	}

	fmt.Println("\n=== Strategy Complete / 策略完成 ===")
}

// fetchKLData retrieves daily K-line data / 獲取日K線數據
func fetchKLData(cli *futuapi.Client, security *qotcommon.Security, count int32) []float64 {
	req := &qot.GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_Forward), // Forward adjusted / 前復權
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    count,
	}

	resp, err := qot.GetKL(cli, req)
	if err != nil {
		log.Printf("GetKL failed: %v", err)
		return nil
	}

	// Extract close prices / 提取收盤價
	prices := make([]float64, len(resp.KLList))
	for i, kl := range resp.KLList {
		prices[i] = kl.ClosePrice
	}

	return prices
}

// calculateSMA calculates Simple Moving Average / 計算簡單移動平均線
func calculateSMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	sma := make([]float64, len(prices)-period+1)
	for i := range sma {
		sum := 0.0
		for j := i; j < i+period; j++ {
			sum += prices[j]
		}
		sma[i] = sum / float64(period)
	}

	return sma
}

// generateSignal generates trading signal based on SMA crossover
// 根據SMA交叉生成交易信號
func generateSignal(shortSMA, longSMA []float64) int {
	if len(shortSMA) < 2 || len(longSMA) < 2 {
		return 0
	}

	// Align arrays / 對齊數組
	shortLen := len(shortSMA)
	longLen := len(longSMA)
	minLen := shortLen
	if longLen < minLen {
		minLen = longLen
	}

	if minLen < 2 {
		return 0
	}

	// Check crossover / 檢查交叉
	prevShort := shortSMA[minLen-2]
	prevLong := longSMA[minLen-2]
	currShort := shortSMA[minLen-1]
	currLong := longSMA[minLen-1]

	// Bullish crossover (golden cross) / 黃金交叉
	if prevShort <= prevLong && currShort > currLong {
		return 1 // BUY
	}

	// Bearish crossover (death cross) / 死亡交叉
	if prevShort >= prevLong && currShort < currLong {
		return -1 // SELL
	}

	return 0 // HOLD
}

// executeTrade executes the trading signal / 執行交易信號
func executeTrade(cli *futuapi.Client, accID uint64, security *qotcommon.Security, signal int, qty int, dryRun bool) {
	if signal == 0 {
		fmt.Println("   Action: HOLD / 操作：持有")
		fmt.Println("   No action required / 無需操作")
		return
	}

	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
	action := "BUY / 買入"
	if signal == -1 {
		trdSide = int32(trdcommon.TrdSide_TrdSide_Sell)
		action = "SELL / 賣出"
	}

	fmt.Printf("   Action: %s\n", action)
	fmt.Printf("   Quantity: %d\n", qty)
	fmt.Printf("   Account: %d\n", accID)

	if dryRun {
		fmt.Println("   Status: DRY RUN - No order placed / 狀態：模擬運行 - 未下單")
	} else {
		// Get current price for limit order / 獲取當前價格以設置限價單
		quotes, err := qot.GetBasicQot(cli, []*qotcommon.Security{security})
		if err != nil {
			fmt.Printf("   ❌ Failed to get quote: %v\n", err)
			return
		}

		price := quotes[0].CurPrice
		if signal == 1 {
			price *= 1.01 // Slightly above market for buy / 買入略高於市場價
		} else {
			price *= 0.99 // Slightly below market for sell / 賣出略低於市場價
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
			fmt.Printf("   ❌ Order failed: %v\n", err)
			return
		}

		fmt.Printf("   ✅ Order placed! OrderID: %d\n", resp.OrderID)
		fmt.Printf("   Price: %.2f\n", price)
	}
}

// getAccountID retrieves the trading account ID / 獲取交易賬戶ID
func getAccountID(cli *futuapi.Client) uint64 {
	// Check command line / 檢查命令行
	if len(os.Args) > 1 {
		id, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err == nil {
			return id
		}
	}

	// Get first account / 獲取第一個賬戶
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
