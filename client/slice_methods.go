package client

import (
	"fmt"
	"strings"

	"github.com/shing1211/futuapi4go/pkg/util"
)

// floatStr converts a float64 to its default string representation.
func floatStr(v float64) string { return fmt.Sprint(v) }

// intStr converts an int64/int32 to string.
func intStr[T int64 | int32 | int](v T) string { return fmt.Sprint(v) }

// boolStr converts a bool to string.
func boolStr(v bool) string { return fmt.Sprint(v) }

// uintStr converts uint64 to string.
func uintStr(v uint64) string { return fmt.Sprint(v) }

// csvLine writes a quoted CSV line.
func csvLine(v ...string) string {
	var b strings.Builder
	for i, s := range v {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteByte('"')
		b.WriteString(strings.ReplaceAll(s, `"`, `""`))
		b.WriteByte('"')
	}
	b.WriteString("\r\n")
	return b.String()
}

// --------------- Quote ---------------

type QuoteSlice []Quote

func (s QuoteSlice) ToJSON() string               { return util.ToJSON(s) }
func (s QuoteSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s QuoteSlice) Filter(fn func(Quote) bool) QuoteSlice {
	var r QuoteSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s QuoteSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Symbol", "Market", "Price", "Open", "High", "Low", "Volume", "Timestamp", "Name", "LastClose", "Turnover", "TurnoverRate", "Amplitude"))
	for _, v := range s {
		b.WriteString(csvLine(v.Symbol, intStr(v.Market), floatStr(v.Price), floatStr(v.Open), floatStr(v.High), floatStr(v.Low), intStr(v.Volume), v.Timestamp, v.Name, floatStr(v.LastClose), floatStr(v.Turnover), floatStr(v.TurnoverRate), floatStr(v.Amplitude)))
	}
	return b.String()
}

// --------------- KLine ---------------

type KLineSlice []KLine

func (s KLineSlice) ToJSON() string             { return util.ToJSON(s) }
func (s KLineSlice) ToJSONPretty() string       { return util.ToJSONPretty(s) }
func (s KLineSlice) Filter(fn func(KLine) bool) KLineSlice {
	var r KLineSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s KLineSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "Open", "High", "Low", "Close", "Volume", "LastClose", "Turnover", "ChangeRate", "Timestamp"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, floatStr(v.Open), floatStr(v.High), floatStr(v.Low), floatStr(v.Close), intStr(v.Volume), floatStr(v.LastClose), floatStr(v.Turnover), floatStr(v.ChangeRate), floatStr(v.Timestamp)))
	}
	return b.String()
}

// --------------- Position ---------------

type PositionSlice []Position

func (s PositionSlice) ToJSON() string               { return util.ToJSON(s) }
func (s PositionSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s PositionSlice) Filter(fn func(Position) bool) PositionSlice {
	var r PositionSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PositionSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Code", "Name", "Market", "Quantity", "CanSellQty", "CostPrice", "CurPrice", "MarketVal", "PnL", "PnLRate", "TodayBuyQty", "TodaySellQty", "TodayPnL", "UnrealizedPL", "RealizedPL", "Currency", "TrdMarket"))
	for _, v := range s {
		b.WriteString(csvLine(v.Code, v.Name, intStr(v.Market), floatStr(v.Quantity), floatStr(v.CanSellQty), floatStr(v.CostPrice), floatStr(v.CurPrice), floatStr(v.MarketVal), floatStr(v.PnL), floatStr(v.PnLRate), floatStr(v.TodayBuyQty), floatStr(v.TodaySellQty), floatStr(v.TodayPnL), floatStr(v.UnrealizedPL), floatStr(v.RealizedPL), intStr(v.Currency), intStr(v.TrdMarket)))
	}
	return b.String()
}

// --------------- Order ---------------

type OrderSlice []Order

func (s OrderSlice) ToJSON() string             { return util.ToJSON(s) }
func (s OrderSlice) ToJSONPretty() string       { return util.ToJSONPretty(s) }
func (s OrderSlice) Filter(fn func(Order) bool) OrderSlice {
	var r OrderSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("OrderID", "Code", "Name", "TrdSide", "OrderType", "OrderStatus", "Price", "Qty", "FillQty", "FillAvgPrice", "CreateTime", "UpdateTime", "Remark", "Currency", "TrdMarket"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.OrderID), v.Code, v.Name, intStr(v.TrdSide), intStr(v.OrderType), intStr(v.OrderStatus), floatStr(v.Price), floatStr(v.Qty), floatStr(v.FillQty), floatStr(v.FillAvgPrice), v.CreateTime, v.UpdateTime, v.Remark, intStr(v.Currency), intStr(v.TrdMarket)))
	}
	return b.String()
}

// --------------- OrderFill ---------------

type OrderFillSlice []OrderFill

func (s OrderFillSlice) ToJSON() string                { return util.ToJSON(s) }
func (s OrderFillSlice) ToJSONPretty() string          { return util.ToJSONPretty(s) }
func (s OrderFillSlice) Filter(fn func(OrderFill) bool) OrderFillSlice {
	var r OrderFillSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderFillSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("FillID", "OrderID", "Code", "Name", "TrdSide", "Price", "Qty", "CreateTime", "Status", "TrdMarket"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.FillID), uintStr(v.OrderID), v.Code, v.Name, intStr(v.TrdSide), floatStr(v.Price), floatStr(v.Qty), v.CreateTime, intStr(v.Status), intStr(v.TrdMarket)))
	}
	return b.String()
}

// --------------- Ticker ---------------

type TickerSlice []Ticker

func (s TickerSlice) ToJSON() string              { return util.ToJSON(s) }
func (s TickerSlice) ToJSONPretty() string        { return util.ToJSONPretty(s) }
func (s TickerSlice) Filter(fn func(Ticker) bool) TickerSlice {
	var r TickerSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s TickerSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "Sequence", "Price", "Volume", "Direction", "Turnover", "Type"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, intStr(v.Sequence), floatStr(v.Price), intStr(v.Volume), v.Direction, floatStr(v.Turnover), intStr(v.Type)))
	}
	return b.String()
}

// --------------- Broker ---------------

type BrokerSlice []Broker

func (s BrokerSlice) ToJSON() string              { return util.ToJSON(s) }
func (s BrokerSlice) ToJSONPretty() string        { return util.ToJSONPretty(s) }
func (s BrokerSlice) Filter(fn func(Broker) bool) BrokerSlice {
	var r BrokerSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s BrokerSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("ID", "Name", "Pos", "Volume"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.ID), v.Name, intStr(v.Pos), intStr(v.Volume)))
	}
	return b.String()
}

// --------------- OrderBookItem ---------------

type OrderBookItemSlice []OrderBookItem

func (s OrderBookItemSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s OrderBookItemSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s OrderBookItemSlice) Filter(fn func(OrderBookItem) bool) OrderBookItemSlice {
	var r OrderBookItemSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderBookItemSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Price", "Volume", "OrderCount"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.Price), intStr(v.Volume), intStr(v.OrderCount)))
	}
	return b.String()
}

// --------------- CapitalFlow ---------------

type CapitalFlowSlice []CapitalFlow

func (s CapitalFlowSlice) ToJSON() string                  { return util.ToJSON(s) }
func (s CapitalFlowSlice) ToJSONPretty() string            { return util.ToJSONPretty(s) }
func (s CapitalFlowSlice) Filter(fn func(CapitalFlow) bool) CapitalFlowSlice {
	var r CapitalFlowSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s CapitalFlowSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "InFlow", "MainInFlow", "SuperInFlow", "BigInFlow", "MidInFlow", "SmlInFlow", "Timestamp"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, floatStr(v.InFlow), floatStr(v.MainInFlow), floatStr(v.SuperInFlow), floatStr(v.BigInFlow), floatStr(v.MidInFlow), floatStr(v.SmlInFlow), floatStr(v.Timestamp)))
	}
	return b.String()
}

// --------------- CapitalDistribution ---------------

type CapitalDistributionSlice []CapitalDistribution

func (s CapitalDistributionSlice) ToJSON() string                         { return util.ToJSON(s) }
func (s CapitalDistributionSlice) ToJSONPretty() string                   { return util.ToJSONPretty(s) }
func (s CapitalDistributionSlice) Filter(fn func(CapitalDistribution) bool) CapitalDistributionSlice {
	var r CapitalDistributionSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s CapitalDistributionSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("MainInflow", "MainOutflow", "MidInflow", "MidOutflow", "SmallInflow", "SmallOutflow", "BigInflow", "BigOutflow", "UpdateTime"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.MainInflow), floatStr(v.MainOutflow), floatStr(v.MidInflow), floatStr(v.MidOutflow), floatStr(v.SmallInflow), floatStr(v.SmallOutflow), floatStr(v.BigInflow), floatStr(v.BigOutflow), v.UpdateTime))
	}
	return b.String()
}

// --------------- Funds ---------------

type FundsSlice []Funds

func (s FundsSlice) ToJSON() string             { return util.ToJSON(s) }
func (s FundsSlice) ToJSONPretty() string       { return util.ToJSONPretty(s) }
func (s FundsSlice) Filter(fn func(Funds) bool) FundsSlice {
	var r FundsSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s FundsSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Power", "TotalAssets", "Cash", "MarketVal", "FrozenCash", "DebtCash", "AvlWithdrawalCash", "Currency", "AvailableFunds", "UnrealizedPL", "RealizedPL", "RiskLevel"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.Power), floatStr(v.TotalAssets), floatStr(v.Cash), floatStr(v.MarketVal), floatStr(v.FrozenCash), floatStr(v.DebtCash), floatStr(v.AvlWithdrawalCash), intStr(v.Currency), floatStr(v.AvailableFunds), floatStr(v.UnrealizedPL), floatStr(v.RealizedPL), intStr(v.RiskLevel)))
	}
	return b.String()
}

// --------------- Account ---------------

type AccountSlice []Account

func (s AccountSlice) ToJSON() string               { return util.ToJSON(s) }
func (s AccountSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s AccountSlice) Filter(fn func(Account) bool) AccountSlice {
	var r AccountSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s AccountSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("AccID", "AccType", "TrdEnv", "CardNum", "AccStatus", "SecurityFirm", "SimAccType"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.AccID), intStr(v.AccType), intStr(v.TrdEnv), v.CardNum, intStr(v.AccStatus), intStr(v.SecurityFirm), intStr(v.SimAccType)))
	}
	return b.String()
}

// --------------- AccCashInfo ---------------

type AccCashInfoSlice []AccCashInfo

func (s AccCashInfoSlice) ToJSON() string                  { return util.ToJSON(s) }
func (s AccCashInfoSlice) ToJSONPretty() string            { return util.ToJSONPretty(s) }
func (s AccCashInfoSlice) Filter(fn func(AccCashInfo) bool) AccCashInfoSlice {
	var r AccCashInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s AccCashInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Currency", "Cash", "AvailableBalance", "NetCashPower"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.Currency), floatStr(v.Cash), floatStr(v.AvailableBalance), floatStr(v.NetCashPower)))
	}
	return b.String()
}

// --------------- AccMarketInfo ---------------

type AccMarketInfoSlice []AccMarketInfo

func (s AccMarketInfoSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s AccMarketInfoSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s AccMarketInfoSlice) Filter(fn func(AccMarketInfo) bool) AccMarketInfoSlice {
	var r AccMarketInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s AccMarketInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("TrdMarket", "Assets"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.TrdMarket), floatStr(v.Assets)))
	}
	return b.String()
}

// --------------- StockFilterResult ---------------

type StockFilterResultSlice []StockFilterResult

func (s StockFilterResultSlice) ToJSON() string                      { return util.ToJSON(s) }
func (s StockFilterResultSlice) ToJSONPretty() string                { return util.ToJSONPretty(s) }
func (s StockFilterResultSlice) Filter(fn func(StockFilterResult) bool) StockFilterResultSlice {
	var r StockFilterResultSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s StockFilterResultSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Code", "Name", "CurPrice", "ChangeRate", "Volume", "Turnover", "HighPrice", "LowPrice"))
	for _, v := range s {
		code := util.SecurityToString(v.Security)
		b.WriteString(csvLine(code, v.Name, floatStr(v.CurPrice), floatStr(v.ChangeRate), intStr(v.Volume), floatStr(v.Turnover), floatStr(v.HighPrice), floatStr(v.LowPrice)))
	}
	return b.String()
}

// --------------- HoldingChangeInfo ---------------

type HoldingChangeInfoSlice []HoldingChangeInfo

func (s HoldingChangeInfoSlice) ToJSON() string                       { return util.ToJSON(s) }
func (s HoldingChangeInfoSlice) ToJSONPretty() string                 { return util.ToJSONPretty(s) }
func (s HoldingChangeInfoSlice) Filter(fn func(HoldingChangeInfo) bool) HoldingChangeInfoSlice {
	var r HoldingChangeInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s HoldingChangeInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("HolderName", "HoldingQty", "HoldingRatio", "ChangeQty", "ChangeRatio", "Time"))
	for _, v := range s {
		b.WriteString(csvLine(v.HolderName, floatStr(v.HoldingQty), floatStr(v.HoldingRatio), floatStr(v.ChangeQty), floatStr(v.ChangeRatio), v.Time))
	}
	return b.String()
}

// --------------- RT ---------------

type RTSlice []RT

func (s RTSlice) ToJSON() string         { return util.ToJSON(s) }
func (s RTSlice) ToJSONPretty() string   { return util.ToJSONPretty(s) }
func (s RTSlice) Filter(fn func(RT) bool) RTSlice {
	var r RTSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s RTSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "Price", "Volume", "LastClose", "AvgPrice", "Turnover"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, floatStr(v.Price), intStr(v.Volume), floatStr(v.LastClose), floatStr(v.AvgPrice), floatStr(v.Turnover)))
	}
	return b.String()
}

// --------------- StaticInfo ---------------

type StaticInfoSlice []StaticInfo

func (s StaticInfoSlice) ToJSON() string                { return util.ToJSON(s) }
func (s StaticInfoSlice) ToJSONPretty() string          { return util.ToJSONPretty(s) }
func (s StaticInfoSlice) Filter(fn func(StaticInfo) bool) StaticInfoSlice {
	var r StaticInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s StaticInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Code", "Name", "Type", "ListTime", "LotSize"))
	for _, v := range s {
		b.WriteString(csvLine(v.Code, v.Name, intStr(v.Type), v.ListTime, intStr(v.LotSize)))
	}
	return b.String()
}

// --------------- SuspendInfo ---------------

type SuspendInfoSlice []SuspendInfo

func (s SuspendInfoSlice) ToJSON() string               { return util.ToJSON(s) }
func (s SuspendInfoSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s SuspendInfoSlice) Filter(fn func(SuspendInfo) bool) SuspendInfoSlice {
	var r SuspendInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s SuspendInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "Timestamp"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, floatStr(v.Timestamp)))
	}
	return b.String()
}

// --------------- BrokerItem ---------------

type BrokerItemSlice []BrokerItem

func (s BrokerItemSlice) ToJSON() string               { return util.ToJSON(s) }
func (s BrokerItemSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s BrokerItemSlice) Filter(fn func(BrokerItem) bool) BrokerItemSlice {
	var r BrokerItemSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s BrokerItemSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Price", "Volume", "BrokerID"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.Price), intStr(v.Volume), intStr(v.BrokerID)))
	}
	return b.String()
}

// --------------- OBItem ---------------

type OBItemSlice []OBItem

func (s OBItemSlice) ToJSON() string         { return util.ToJSON(s) }
func (s OBItemSlice) ToJSONPretty() string   { return util.ToJSONPretty(s) }
func (s OBItemSlice) Filter(fn func(OBItem) bool) OBItemSlice {
	var r OBItemSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OBItemSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Price", "Volume", "OrderCount"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.Price), intStr(v.Volume), intStr(v.OrderCount)))
	}
	return b.String()
}

// --------------- RehabInfo ---------------

type RehabInfoSlice []RehabInfo

func (s RehabInfoSlice) ToJSON() string              { return util.ToJSON(s) }
func (s RehabInfoSlice) ToJSONPretty() string        { return util.ToJSONPretty(s) }
func (s RehabInfoSlice) Filter(fn func(RehabInfo) bool) RehabInfoSlice {
	var r RehabInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s RehabInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Time", "FwdFactorA", "FwdFactorB", "BwdFactorA", "BwdFactorB", "SplitBase", "SplitErt", "AllotPrice"))
	for _, v := range s {
		b.WriteString(csvLine(v.Time, floatStr(v.FwdFactorA), floatStr(v.FwdFactorB), floatStr(v.BwdFactorA), floatStr(v.BwdFactorB), intStr(v.SplitBase), intStr(v.SplitErt), floatStr(v.AllotPrice)))
	}
	return b.String()
}

// --------------- OrderFeeItemInfo ---------------

type OrderFeeItemInfoSlice []OrderFeeItemInfo

func (s OrderFeeItemInfoSlice) ToJSON() string                    { return util.ToJSON(s) }
func (s OrderFeeItemInfoSlice) ToJSONPretty() string              { return util.ToJSONPretty(s) }
func (s OrderFeeItemInfoSlice) Filter(fn func(OrderFeeItemInfo) bool) OrderFeeItemInfoSlice {
	var r OrderFeeItemInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderFeeItemInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Title", "Value"))
	for _, v := range s {
		b.WriteString(csvLine(v.Title, floatStr(v.Value)))
	}
	return b.String()
}

// --------------- PushQuote ---------------

type PushQuoteSlice []PushQuote

func (s PushQuoteSlice) ToJSON() string              { return util.ToJSON(s) }
func (s PushQuoteSlice) ToJSONPretty() string        { return util.ToJSONPretty(s) }
func (s PushQuoteSlice) Filter(fn func(PushQuote) bool) PushQuoteSlice {
	var r PushQuoteSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PushQuoteSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Market", "Code", "Name", "CurPrice", "OpenPrice", "HighPrice", "LowPrice", "Volume", "Turnover"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.Market), v.Code, v.Name, floatStr(v.CurPrice), floatStr(v.OpenPrice), floatStr(v.HighPrice), floatStr(v.LowPrice), intStr(v.Volume), floatStr(v.Turnover)))
	}
	return b.String()
}

// --------------- PushKLine ---------------

type PushKLineSlice []PushKLine

func (s PushKLineSlice) ToJSON() string               { return util.ToJSON(s) }
func (s PushKLineSlice) ToJSONPretty() string         { return util.ToJSONPretty(s) }
func (s PushKLineSlice) Filter(fn func(PushKLine) bool) PushKLineSlice {
	var r PushKLineSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PushKLineSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Market", "Code", "Name", "KLType", "Time", "Open", "High", "Low", "Close", "Volume", "Turnover"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.Market), v.Code, v.Name, intStr(v.KLType), v.Time, floatStr(v.Open), floatStr(v.High), floatStr(v.Low), floatStr(v.Close), intStr(v.Volume), floatStr(v.Turnover)))
	}
	return b.String()
}

// --------------- PushOrderUpdate ---------------

type PushOrderUpdateSlice []PushOrderUpdate

func (s PushOrderUpdateSlice) ToJSON() string                     { return util.ToJSON(s) }
func (s PushOrderUpdateSlice) ToJSONPretty() string               { return util.ToJSONPretty(s) }
func (s PushOrderUpdateSlice) Filter(fn func(PushOrderUpdate) bool) PushOrderUpdateSlice {
	var r PushOrderUpdateSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PushOrderUpdateSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("OrderID", "Code", "SecMarket", "TrdSide", "Qty", "Price", "OrderStatus"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.OrderID), v.Code, intStr(v.SecMarket), intStr(v.TrdSide), floatStr(v.Qty), floatStr(v.Price), intStr(v.OrderStatus)))
	}
	return b.String()
}

// --------------- PushOrderFill ---------------

type PushOrderFillSlice []PushOrderFill

func (s PushOrderFillSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s PushOrderFillSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s PushOrderFillSlice) Filter(fn func(PushOrderFill) bool) PushOrderFillSlice {
	var r PushOrderFillSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PushOrderFillSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("OrderID", "Code", "SecMarket", "TrdSide", "Qty", "Price", "FillID", "CreateTime"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.OrderID), v.Code, intStr(v.SecMarket), intStr(v.TrdSide), floatStr(v.Qty), floatStr(v.Price), uintStr(v.FillID), v.FillCreateTime))
	}
	return b.String()
}

// --------------- PriceReminderItemInfo ---------------

type PriceReminderItemInfoSlice []PriceReminderItemInfo

func (s PriceReminderItemInfoSlice) ToJSON() string                         { return util.ToJSON(s) }
func (s PriceReminderItemInfoSlice) ToJSONPretty() string                   { return util.ToJSONPretty(s) }
func (s PriceReminderItemInfoSlice) Filter(fn func(PriceReminderItemInfo) bool) PriceReminderItemInfoSlice {
	var r PriceReminderItemInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s PriceReminderItemInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Key", "Type", "Freq", "Value", "Note", "IsEnable"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.Key), intStr(v.Type), intStr(v.Freq), floatStr(v.Value), v.Note, boolStr(v.IsEnable)))
	}
	return b.String()
}

// --------------- FlowSummaryInfo ---------------

type FlowSummaryInfoSlice []FlowSummaryInfo

func (s FlowSummaryInfoSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s FlowSummaryInfoSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s FlowSummaryInfoSlice) Filter(fn func(FlowSummaryInfo) bool) FlowSummaryInfoSlice {
	var r FlowSummaryInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s FlowSummaryInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("CashFlowID", "ClearingDate", "SettlementDate", "Currency", "CashFlowType", "CashFlowDirection", "CashFlowAmount", "CashFlowRemark"))
	for _, v := range s {
		b.WriteString(csvLine(uintStr(v.CashFlowID), v.ClearingDate, v.SettlementDate, intStr(v.Currency), v.CashFlowType, intStr(v.CashFlowDirection), floatStr(v.CashFlowAmount), v.CashFlowRemark))
	}
	return b.String()
}

// --------------- AccTradingInfo ---------------

type AccTradingInfoSlice []AccTradingInfo

func (s AccTradingInfoSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s AccTradingInfoSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s AccTradingInfoSlice) Filter(fn func(AccTradingInfo) bool) AccTradingInfoSlice {
	var r AccTradingInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s AccTradingInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("MaxCashBuy", "MaxCashAndMarginBuy", "MaxPositionSell", "MaxSellShort", "MaxBuyBack", "LongRequiredIM", "ShortRequiredIM"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.MaxCashBuy), floatStr(v.MaxCashAndMarginBuy), floatStr(v.MaxPositionSell), floatStr(v.MaxSellShort), floatStr(v.MaxBuyBack), floatStr(v.LongRequiredIM), floatStr(v.ShortRequiredIM)))
	}
	return b.String()
}

// --------------- OrderBookDetail ---------------

type OrderBookDetailSlice []OrderBookDetail

func (s OrderBookDetailSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s OrderBookDetailSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s OrderBookDetailSlice) Filter(fn func(OrderBookDetail) bool) OrderBookDetailSlice {
	var r OrderBookDetailSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderBookDetailSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("OrderID", "Volume"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.OrderID), intStr(v.Volume)))
	}
	return b.String()
}

// --------------- MaxTrdQtysInfo ---------------

type MaxTrdQtysInfoSlice []MaxTrdQtysInfo

func (s MaxTrdQtysInfoSlice) ToJSON() string                   { return util.ToJSON(s) }
func (s MaxTrdQtysInfoSlice) ToJSONPretty() string             { return util.ToJSONPretty(s) }
func (s MaxTrdQtysInfoSlice) Filter(fn func(MaxTrdQtysInfo) bool) MaxTrdQtysInfoSlice {
	var r MaxTrdQtysInfoSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s MaxTrdQtysInfoSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("MaxCashBuy", "MaxCashAndMarginBuy", "MaxPositionSell", "MaxSellShort", "MaxBuyBack"))
	for _, v := range s {
		b.WriteString(csvLine(floatStr(v.MaxCashBuy), floatStr(v.MaxCashAndMarginBuy), floatStr(v.MaxPositionSell), floatStr(v.MaxSellShort), floatStr(v.MaxBuyBack)))
	}
	return b.String()
}

// --------------- DelayStatisticsItem ---------------

type DelayStatisticsItemSlice []DelayStatisticsItem

func (s DelayStatisticsItemSlice) ToJSON() string                       { return util.ToJSON(s) }
func (s DelayStatisticsItemSlice) ToJSONPretty() string                 { return util.ToJSONPretty(s) }
func (s DelayStatisticsItemSlice) Filter(fn func(DelayStatisticsItem) bool) DelayStatisticsItemSlice {
	var r DelayStatisticsItemSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s DelayStatisticsItemSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("Begin", "End", "Count", "Proportion", "CumulativeRatio"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.Begin), intStr(v.End), intStr(v.Count), floatStr(v.Proportion), floatStr(v.CumulativeRatio)))
	}
	return b.String()
}

// --------------- ReqReplyStatisticsItem ---------------

type ReqReplyStatisticsItemSlice []ReqReplyStatisticsItem

func (s ReqReplyStatisticsItemSlice) ToJSON() string                       { return util.ToJSON(s) }
func (s ReqReplyStatisticsItemSlice) ToJSONPretty() string                 { return util.ToJSONPretty(s) }
func (s ReqReplyStatisticsItemSlice) Filter(fn func(ReqReplyStatisticsItem) bool) ReqReplyStatisticsItemSlice {
	var r ReqReplyStatisticsItemSlice
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
func (s ReqReplyStatisticsItemSlice) ToCSV() string {
	var b strings.Builder
	b.WriteString(csvLine("ProtoID", "Count", "TotalCostAvg", "OpenDCostAvg", "NetDelayAvg", "IsLocalReply"))
	for _, v := range s {
		b.WriteString(csvLine(intStr(v.ProtoID), intStr(v.Count), floatStr(v.TotalCostAvg), floatStr(v.OpenDCostAvg), floatStr(v.NetDelayAvg), boolStr(v.IsLocalReply)))
	}
	return b.String()
}
