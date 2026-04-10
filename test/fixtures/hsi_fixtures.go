package testutil

import (
	"fmt"
	"math"
	"time"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
)

// ============================================================================
// HSI (Hang Seng Index) Test Constants
// ============================================================================

// HSI Main Index constants
const (
	HSIMarket       = int32(qotcommon.QotMarket_QotMarket_HK_Security)
	HSICode         = "800100"
	HSIName         = "HANG SENG INDEX"
	HSICodeMain     = "HSI"
	HSISecurityKey  = "1.800100" // market.code format
)

// HSI Futures constants
const (
	HSIFuturesMarket = int32(qotcommon.QotMarket_QotMarket_HK_Future)
	HSIFuturesCode   = "HSImain"
	HSIFuturesName   = "HANG SENG INDEX FUTURES"
)

// Test trading account
const (
	TestAccID  = uint64(1234567890)
	TestTrdEnv = int32(trdcommon.TrdEnv_TrdEnv_Simulate)
	TestTrdMkt = int32(trdcommon.TrdMarket_TrdMarket_HK)
)

// ============================================================================
// Realistic HSI Market Data Fixtures
// ============================================================================

// HSIQuote returns a realistic real-time quote for HSI
func HSIQuote() *qotcommon.BasicQot {
	curPrice := float64(18523.45)
	openPrice := float64(18480.00)
	highPrice := float64(18590.12)
	lowPrice := float64(18420.50)
	lastClose := float64(18498.23)
	volume := uint64(2345678900)
	turnover := float64(98765432100.50)
	turnoverRate := float64(0.0234)
	amplitude := float64(0.92)
	updateTime := time.Now().Format("2006-01-02 15:04:05")

	market := HSIMarket
	code := HSICode

	return &qotcommon.BasicQot{
		Security: &qotcommon.Security{
			Market: &market,
			Code:   &code,
		},
		Name:         ptrStr(HSIName),
		IsSuspended:  ptrBool(false),
		UpdateTime:   ptrStr(updateTime),
		HighPrice:    ptrFloat64(highPrice),
		OpenPrice:    ptrFloat64(openPrice),
		LowPrice:     ptrFloat64(lowPrice),
		CurPrice:     ptrFloat64(curPrice),
		LastClosePrice: ptrFloat64(lastClose),
		Volume:       ptrInt64(int64(volume)),
		Turnover:     ptrFloat64(turnover),
		TurnoverRate: ptrFloat64(turnoverRate),
		Amplitude:    ptrFloat64(amplitude),
	}
}

// HSIOrderBookLevels returns realistic HSI order book with 10 levels
func HSIOrderBookLevels(numLevels int) (asks, bids []*qotcommon.OrderBook) {
	basePrice := 18523.45
	spread := 5.0 // HSI typical spread

	for i := 0; i < numLevels && i < 10; i++ {
		price := basePrice + float64(i+1)*spread
		vol := int64(100 + i*50)
		orderCount := int32(5 + i*2)

		asks = append(asks, &qotcommon.OrderBook{
			Price:      ptrFloat64(price),
			Volume:     ptrInt64(vol),
			OrderCount: ptrInt32(orderCount),
		})

		bidPrice := basePrice - float64(i+1)*spread
		bids = append(bids, &qotcommon.OrderBook{
			Price:      ptrFloat64(bidPrice),
			Volume:     ptrInt64(vol + 50),
			OrderCount: ptrInt32(orderCount + 3),
		})
	}
	return asks, bids
}

// HSITickerData returns realistic HSI tick-by-tick trades
func HSITickerData(count int) []*qotcommon.Ticker {
	tickers := make([]*qotcommon.Ticker, 0, count)
	baseTime := time.Now().Truncate(time.Hour)
	basePrice := 18523.45

	for i := 0; i < count; i++ {
		seq := int64(1000000 + i)
		price := basePrice + float64(i%10)*0.5 - 2.5
		volume := int64(10 + i%5*5)
		turnover := price * float64(volume)
		timeStr := baseTime.Add(time.Duration(i) * time.Second).Format("2006-01-02 15:04:05")
		dir := int32(i % 3) // 0=Buy, 1=Sell, 2=Neutral

		tickers = append(tickers, &qotcommon.Ticker{
			Time:     ptrStr(timeStr),
			Sequence: ptrInt64(seq),
			Dir:      ptrInt32(dir),
			Price:    ptrFloat64(price),
			Volume:   ptrInt64(volume),
			Turnover: ptrFloat64(turnover),
		})
	}
	return tickers
}

// HSIKLineData returns realistic HSI K-line data
func HSIKLineData(count int, klType int32) []*qotcommon.KLine {
	klines := make([]*qotcommon.KLine, 0, count)
	baseDate := time.Now().AddDate(0, 0, -count)
	basePrice := 18400.00

	for i := 0; i < count; i++ {
		var t time.Time
		switch klType {
		case int32(qotcommon.KLType_KLType_1Min):
			t = baseDate.Add(time.Duration(i) * time.Minute)
		case int32(qotcommon.KLType_KLType_5Min):
			t = baseDate.Add(time.Duration(i) * 5 * time.Minute)
		case int32(qotcommon.KLType_KLType_Day):
			t = baseDate.AddDate(0, 0, i)
		case int32(qotcommon.KLType_KLType_Week):
			t = baseDate.AddDate(0, 0, i*7)
		default:
			t = baseDate.AddDate(0, 0, i)
		}

		open := basePrice + float64(i)*2.5
		close := open + float64(i%5)*5.0 - 10.0
		high := math.Max(open, close) + float64(i%3)*3.0
		low := math.Min(open, close) - float64(i%3)*3.0
		volume := int64(1000000 + i*10000)
		turnover := float64(volume) * (open + close) / 2.0
		changeRate := (close - open) / open * 100.0

		timeStr := t.Format("2006-01-02 15:04:05")

		klines = append(klines, &qotcommon.KLine{
			Time:           ptrStr(timeStr),
			OpenPrice:      ptrFloat64(open),
			ClosePrice:     ptrFloat64(close),
			HighPrice:      ptrFloat64(high),
			LowPrice:       ptrFloat64(low),
			LastClosePrice: ptrFloat64(open),
			Volume:         ptrInt64(volume),
			Turnover:       ptrFloat64(turnover),
			ChangeRate:     ptrFloat64(changeRate),
		})
	}
	return klines
}

// HSIRTDData returns realistic HSI intraday time-share data
func HSIRTDData(count int) []*qotcommon.TimeShare {
	rtd := make([]*qotcommon.TimeShare, 0, count)
	baseTime := time.Now().Truncate(24 * time.Hour).Add(9 * time.Hour + 30*time.Minute) // HK market open
	basePrice := 18480.00

	for i := 0; i < count; i++ {
		t := baseTime.Add(time.Duration(i) * time.Minute)
		price := basePrice + float64(i)*1.2 - float64(i%10)*2.0
		volume := int64(50000 + i*1000)
		turnover := price * float64(volume)
		timeStr := t.Format("15:04")

		rtd = append(rtd, &qotcommon.TimeShare{
			Time:     ptrStr(timeStr),
			Price:    ptrFloat64(price),
			Volume:   ptrInt64(volume),
			Turnover: ptrFloat64(turnover),
		})
	}
	return rtd
}

// ============================================================================
// Trading Fixtures
// ============================================================================

// HSIOrder returns a realistic HSI order
func HSIOrder(orderID uint64) *trdcommon.Order {
	orderIDPtr := orderID
	price := 18520.00
	qty := float64(1)
	fillQty := float64(0)
	createTime := time.Now().Format("2006-01-02 15:04:05")
	code := HSIFuturesCode
	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
	orderType := int32(trdcommon.OrderType_OrderType_Normal)
	orderStatus := int32(trdcommon.OrderStatus_OrderStatus_Submitted)

	return &trdcommon.Order{
		OrderID:     &orderIDPtr,
		Code:        ptrStr(code),
		Name:        ptrStr(HSIFuturesName),
		TrdSide:     ptrInt32(trdSide),
		OrderType:   ptrInt32(orderType),
		OrderStatus: ptrInt32(orderStatus),
		Price:       ptrFloat64(price),
		Qty:         ptrFloat64(qty),
		FillQty:     ptrFloat64(fillQty),
		CreateTime:  ptrStr(createTime),
	}
}

// HSIOrderFill returns a realistic HSI order fill
func HSIOrderFill(fillID uint64, orderID uint64) *trdcommon.OrderFill {
	price := 18523.45
	qty := float64(1)
	createTime := time.Now().Format("2006-01-02 15:04:05")
	code := HSIFuturesCode
	trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)

	return &trdcommon.OrderFill{
		FillID:     ptrUint64(fillID),
		OrderID:    ptrUint64(orderID),
		Code:       ptrStr(code),
		Name:       ptrStr(HSIFuturesName),
		TrdSide:    ptrInt32(trdSide),
		Price:      ptrFloat64(price),
		Qty:        ptrFloat64(qty),
		CreateTime: ptrStr(createTime),
	}
}

// HSIPosition returns a realistic HSI position
func HSIPosition() *trdcommon.Position {
	code := HSIFuturesCode
	qty := float64(2)
	canSellQty := float64(2)
	price := 18523.45
	costPrice := 18480.00
	plVal := float64(86.90) // (18523.45 - 18480.00) * 2
	plRatio := 0.235        // 86.90 / (18480.00 * 2) * 100

	return &trdcommon.Position{
		Code:       ptrStr(code),
		Name:       ptrStr(HSIFuturesName),
		Qty:        ptrFloat64(qty),
		CanSellQty: ptrFloat64(canSellQty),
		Price:      ptrFloat64(price),
		CostPrice:  ptrFloat64(costPrice),
		PlVal:      ptrFloat64(plVal),
		PlRatio:    ptrFloat64(plRatio),
	}
}

// HSIFunds returns realistic account funds for HSI trading
func HSIFunds() *trdcommon.Funds {
	totalAssets := float64(500000.00)
	cash := float64(250000.00)
	marketVal := float64(74093.80) // 2 contracts * 18523.45 * 2 (multiplier)
	frozenCash := float64(37046.90)
	power := float64(425906.10)

	return &trdcommon.Funds{
		TotalAssets: ptrFloat64(totalAssets),
		Cash:        ptrFloat64(cash),
		MarketVal:   ptrFloat64(marketVal),
		FrozenCash:  ptrFloat64(frozenCash),
		Power:       ptrFloat64(power),
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

func ptrStr(s string) *string       { return &s }
func ptrFloat64(f float64) *float64 { return &f }
func ptrUint64(u uint64) *uint64    { return &u }
func ptrInt64(i int64) *int64       { return &i }
func ptrInt32(i int32) *int32       { return &i }
func ptrBool(b bool) *bool          { return &b }

// SecurityPtr returns a Security pointer for HSI
func SecurityPtr(market int32, code string) *qotcommon.Security {
	return &qotcommon.Security{
		Market: ptrInt32(market),
		Code:   ptrStr(code),
	}
}

// HSISecurity returns HSI main index security
func HSISecurity() *qotcommon.Security {
	return SecurityPtr(HSIMarket, HSICode)
}

// HSIFuturesSecurity returns HSI futures security
func HSIFuturesSecurity() *qotcommon.Security {
	return SecurityPtr(HSIFuturesMarket, HSIFuturesCode)
}

// FormatTestError formats error messages for tests
func FormatTestError(testName string, expected, actual interface{}) string {
	return fmt.Sprintf("[%s] expected %v, got %v", testName, expected, actual)
}

// AssertEqual checks equality and returns error message if not equal
func AssertEqual(testName string, expected, actual interface{}) string {
	if expected != actual {
		return FormatTestError(testName, expected, actual)
	}
	return ""
}

// AssertNotNil checks value is not nil
func AssertNotNil(testName string, v interface{}) string {
	if v == nil {
		return fmt.Sprintf("[%s] expected non-nil value", testName)
	}
	return ""
}
