// Example: Breakout Trading Strategy / 突破交易策略
//
// This example demonstrates a breakout trading algorithm that detects
// price breakouts from consolidation ranges and enters positions
// in the breakout direction.
//
// Strategy:
// - Calculate price range over N periods (high/low)
// - When price breaks above resistance → BUY
// - When price breaks below support → SELL
// - Use stop-loss and take-profit for risk management
//
// Usage:
//   go run main.go [account_id]
//
// Note: DRY RUN by default for safety.

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
	LookbackPeriod = 20  // Periods to calculate range / 計算範圍的週期數
	BreakoutBuffer = 0.5 // % buffer to confirm breakout / 確認突破的緩衝百分比
	StopLossPct    = 2.0 // Stop loss percentage / 止損百分比
	TakeProfitPct  = 4.0 // Take profit percentage / 止盈百分比
	TradeQty       = 100 // Shares to trade / 交易股數
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== Breakout Trading Strategy / 突破交易策略 ===")
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
	fmt.Printf("   Security / 股票:            %s\n", code)
	fmt.Printf("   Lookback Period / 回顧週期:   %d days\n", LookbackPeriod)
	fmt.Printf("   Breakout Buffer / 突破緩衝:   %.1f%%\n", BreakoutBuffer)
	fmt.Printf("   Stop Loss / 止損:            %.1f%%\n", StopLossPct)
	fmt.Printf("   Take Profit / 止盈:          %.1f%%\n", TakeProfitPct)
	fmt.Printf("   Trade Qty / 交易數量:         %d\n", TradeQty)
	fmt.Println()

	// Fetch K-line data / 獲取K線數據
	fmt.Println("📈 Fetching K-line data / 正在獲取K線數據...")
	klData := fetchKLData(cli, security, LookbackPeriod+5)
	if len(klData) < LookbackPeriod {
		fmt.Printf("❌ Insufficient data: need %d, got %d\n", LookbackPeriod, len(klData))
		return
	}
	fmt.Printf("✅ Retrieved %d K-line records\n\n", len(klData))

	// Calculate support and resistance / 計算支撐位和阻力位
	fmt.Println("📐 Calculating Support & Resistance / 正在計算支撐位和阻力位...")
	high, low := calculateRange(klData[:LookbackPeriod])
	currentPrice := klData[len(klData)-1].Close

	fmt.Printf("   Resistance / 阻力位: HK$%.2f\n", high)
	fmt.Printf("   Support / 支撐位:    HK$%.2f\n", low)
	fmt.Printf("   Current Price / 當前價: HK$%.2f\n\n", currentPrice)

	// Detect breakout / 檢測突破
	breakoutType, breakoutPrice := detectBreakout(currentPrice, high, low)

	fmt.Println("🎯 Breakout Detection / 突破檢測:")
	switch breakoutType {
	case "UP":
		fmt.Printf("   🟢 UPWARD BREAKOUT / 向上突破\n")
		fmt.Printf("   Price broke above resistance at HK$%.2f\n", breakoutPrice)
		fmt.Printf("   價格突破阻力位 HK$%.2f\n", breakoutPrice)
	case "DOWN":
		fmt.Printf("   🔴 DOWNWARD BREAKOUT / 向下突破\n")
		fmt.Printf("   Price broke below support at HK$%.2f\n", breakoutPrice)
		fmt.Printf("   價格跌破支撐位 HK$%.2f\n", breakoutPrice)
	default:
		fmt.Println("   ⚪ NO BREAKOUT / 無突破")
		fmt.Println("   Price is within range / 價格在區間內")
	}
	fmt.Println()

	// Calculate risk management levels / 計算風險管理級別
	var stopLoss, takeProfit float64
	if breakoutType != "" {
		stopLoss, takeProfit = calculateRiskLevels(breakoutPrice, breakoutType)
		fmt.Println("🛡️  Risk Management / 風險管理:")
		fmt.Printf("   Entry / 入場:      HK$%.2f\n", breakoutPrice)
		fmt.Printf("   Stop Loss / 止損:   HK$%.2f (-%.1f%%)\n", stopLoss, StopLossPct)
		fmt.Printf("   Take Profit / 止盈: HK$%.2f (+%.1f%%)\n", takeProfit, TakeProfitPct)
		fmt.Println()
	}

	// Execute trade / 執行交易
	dryRun := true // Set to false for live trading / 設置為false以進行真實交易

	if dryRun {
		fmt.Println("⚠️  DRY RUN MODE / 模擬運行模式")
		fmt.Println("   No real orders will be placed")
		fmt.Println("   不會下達真實訂單")
	} else {
		fmt.Println("🚨 LIVE TRADING MODE / 真實交易模式")
	}
	fmt.Println()

	executeBreakoutTrade(cli, accID, security, breakoutType, breakoutPrice, TradeQty,
		stopLoss, takeProfit, dryRun)

	fmt.Println("\n=== Strategy Complete / 策略完成 ===")
}

// KLineData holds simplified K-line data / K線數據結構
type KLineData struct {
	Close  float64
	High   float64
	Low    float64
	Volume int64
}

// fetchKLData retrieves K-line data / 獲取K線數據
func fetchKLData(cli *futuapi.Client, security *qotcommon.Security, count int32) []KLineData {
	req := &qot.GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_Forward),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    count,
	}

	resp, err := qot.GetKL(cli, req)
	if err != nil {
		log.Printf("GetKL failed: %v", err)
		return nil
	}

	data := make([]KLineData, len(resp.KLList))
	for i, kl := range resp.KLList {
		data[i] = KLineData{
			Close:  kl.ClosePrice,
			High:   kl.HighPrice,
			Low:    kl.LowPrice,
			Volume: kl.Volume,
		}
	}

	return data
}

// calculateRange calculates the high/low range / 計算高低點範圍
func calculateRange(data []KLineData) (high, low float64) {
	if len(data) == 0 {
		return 0, 0
	}

	high = data[0].High
	low = data[0].Low

	for _, d := range data[1:] {
		if d.High > high {
			high = d.High
		}
		if d.Low < low {
			low = d.Low
		}
	}

	return
}

// detectBreakout detects if price has broken out of range
// 檢測價格是否突破區間
func detectBreakout(currentPrice, high, low float64) (string, float64) {
	buffer := high * BreakoutBuffer / 100

	// Upward breakout / 向上突破
	if currentPrice > high+buffer {
		return "UP", currentPrice
	}

	// Downward breakout / 向下突破
	if currentPrice < low-buffer {
		return "DOWN", currentPrice
	}

	return "", 0
}

// calculateRiskLevels calculates stop-loss and take-profit levels
// 計算止損和止盈位
func calculateRiskLevels(entryPrice float64, breakoutType string) (stopLoss, takeProfit float64) {
	if breakoutType == "UP" {
		stopLoss = entryPrice * (1 - StopLossPct/100)
		takeProfit = entryPrice * (1 + TakeProfitPct/100)
	} else {
		stopLoss = entryPrice * (1 + StopLossPct/100)
		takeProfit = entryPrice * (1 - TakeProfitPct/100)
	}

	return
}

// executeBreakoutTrade executes the breakout trade / 執行突破交易
func executeBreakoutTrade(cli *futuapi.Client, accID uint64, security *qotcommon.Security,
	breakoutType string, entryPrice float64, qty int, stopLoss, takeProfit float64, dryRun bool) {

	if breakoutType == "" {
		fmt.Println("   Action: WAIT / 操作：等待")
		fmt.Println("   No breakout detected - standing by")
		fmt.Println("   未檢測到突破 - 正在等待")
		return
	}

	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
	action := "BUY / 買入"
	if breakoutType == "DOWN" {
		trdSide = int32(trdcommon.TrdSide_TrdSide_Sell)
		action = "SELL SHORT / 賣空"
	}

	fmt.Printf("   Signal: %s\n", action)
	fmt.Printf("   Entry: HK$%.2f\n", entryPrice)
	fmt.Printf("   Qty: %d shares\n", qty)
	fmt.Printf("   Expected Value: HK$%.2f\n", entryPrice*float64(qty))

	if dryRun {
		fmt.Println("\n   [DRY RUN] Order would be:")
		fmt.Printf("   - %s %d @ HK$%.2f\n", action, qty, entryPrice)
		fmt.Printf("   - Stop Loss @ HK$%.2f\n", stopLoss)
		fmt.Printf("   - Take Profit @ HK$%.2f\n", takeProfit)
		fmt.Println("   Status: SIMULATED / 狀態：已模擬")
	} else {
		req := &trd.PlaceOrderRequest{
			AccID:     accID,
			TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
			Code:      security.GetCode(),
			TrdSide:   trdSide,
			OrderType: int32(trdcommon.OrderType_OrderType_Normal),
			Price:     entryPrice,
			Qty:       float64(qty),
		}

		resp, err := trd.PlaceOrder(cli, req)
		if err != nil {
			fmt.Printf("   ❌ Order failed: %v\n", err)
			return
		}

		fmt.Printf("\n   ✅ Order placed! OrderID: %d\n", resp.OrderID)
		fmt.Printf("   🛡️  Set stop-loss alert @ HK$%.2f\n", stopLoss)
		fmt.Printf("   🎯 Set take-profit alert @ HK$%.2f\n", takeProfit)
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
