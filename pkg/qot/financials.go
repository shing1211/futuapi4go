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

package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialrevenuebreakdown"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsearnpricehist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsearnpricemove"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsstatements"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetFinancialsStatements       = 3227
	ProtoID_GetFinancialsRevenueBreakdown = 3228
)

type GetFinancialsStatementsRequest struct {
	Security       *qotcommon.Security
	StatementType  int32
	FinancialType  int32
	CurrencyCode   string
	NextKey        string
	Num            int32
}

type GetFinancialsStatementsResponse struct {
	StructureList []*qotgetfinancialsstatements.FinancialFieldInfo
	ReportList    []*qotgetfinancialsstatements.FinancialReport
	NextKey       string
}

func GetFinancialsStatements(ctx context.Context, c *futuapi.Client, req *GetFinancialsStatementsRequest) (*GetFinancialsStatementsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsStatements: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsStatements: Security is nil")
	}
	c2s := &qotgetfinancialsstatements.C2S{
		Security: req.Security,
	}
	if req.StatementType != 0 {
		v := qotcommon.FinancialStatementsType(req.StatementType)
		c2s.StatementType = &v
	}
	if req.FinancialType != 0 {
		v := qotcommon.F10Type(req.FinancialType)
		c2s.FinancialType = &v
	}
	if req.CurrencyCode != "" {
		c2s.CurrencyCode = &req.CurrencyCode
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetfinancialsstatements.Request{C2S: c2s}
	var rsp qotgetfinancialsstatements.Response

	if err := c.RequestContext(ctx, ProtoID_GetFinancialsStatements, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsStatements", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsStatements", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetFinancialsStatementsResponse{
		StructureList: s2c.StructureList,
		ReportList:    s2c.ReportList,
		NextKey:      util.ProtoStr(s2c.NextKey),
	}, nil
}

type GetFinancialsRevenueBreakdownRequest struct {
	Security      *qotcommon.Security
	Date          uint32
	FinancialType int32
	CurrencyCode  string
}

type GetFinancialsRevenueBreakdownResponse struct {
	Period         string
	BreakdownList  []*qotgetfinancialrevenuebreakdown.RevenueBreakdownGroup
	CurrencyCode   string
	ScreenDateList []*qotgetfinancialrevenuebreakdown.ScreenDate
}

func GetFinancialsRevenueBreakdown(ctx context.Context, c *futuapi.Client, req *GetFinancialsRevenueBreakdownRequest) (*GetFinancialsRevenueBreakdownResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsRevenueBreakdown: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsRevenueBreakdown: Security is nil")
	}
	c2s := &qotgetfinancialrevenuebreakdown.C2S{
		Security: req.Security,
	}
	if req.Date != 0 {
		c2s.Date = &req.Date
	}
	if req.FinancialType != 0 {
		v := qotcommon.F10Type(req.FinancialType)
		c2s.FinancialType = &v
	}
	if req.CurrencyCode != "" {
		c2s.CurrencyCode = &req.CurrencyCode
	}
	pkt := &qotgetfinancialrevenuebreakdown.Request{C2S: c2s}
	var rsp qotgetfinancialrevenuebreakdown.Response

	if err := c.RequestContext(ctx, ProtoID_GetFinancialsRevenueBreakdown, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsRevenueBreakdown", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsRevenueBreakdown", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetFinancialsRevenueBreakdownResponse{
		Period:        util.ProtoStr(s2c.Period),
		BreakdownList: s2c.BreakdownList,
		CurrencyCode:  util.ProtoStr(s2c.CurrencyCode),
		ScreenDateList: s2c.ScreenDateList,
	}, nil
}

// PricePerformanceRow represents price data for a single trading day around earnings.
type PricePerformanceRow struct {
	TradingDay     int64
	TradingDayStr  string
	ClosePrice     float64
	OpenPrice      float64
	HighestPrice   float64
	LowestPrice    float64
	LastClosePrice float64
	OptionIV       float64
	OptionHV       float64
}

// ReportCycleQuoteData represents earnings price data for one reporting period.
type ReportCycleQuoteData struct {
	FiscalYear       int32
	FinancialType    int32
	PeriodText       string
	PubTradingDay    int64
	PubTradingDayStr string
	PubType          qotcommon.EarningsPubTimeType
	PriceInfoIndex   int32
	ItemList         []*PricePerformanceRow
}

// GetFinancialsEarningsPriceMoveResponse is the response type for GetFinancialsEarningsPriceMove.
type GetFinancialsEarningsPriceMoveResponse struct {
	DetailList []*ReportCycleQuoteData
}

// PricePerformanceRowFromProto converts a proto PricePerformanceRow to our domain type.
func PricePerformanceRowFromProto(p *qotgetfinancialsearnpricemove.PricePerformanceRow) *PricePerformanceRow {
	if p == nil {
		return nil
	}
	return &PricePerformanceRow{
		TradingDay:     util.ProtoInt64(p.TradingDay),
		TradingDayStr:  util.ProtoStr(p.TradingDayStr),
		ClosePrice:     util.ProtoFloat64(p.ClosePrice),
		OpenPrice:      util.ProtoFloat64(p.OpenPrice),
		HighestPrice:   util.ProtoFloat64(p.HighestPrice),
		LowestPrice:    util.ProtoFloat64(p.LowestPrice),
		LastClosePrice: util.ProtoFloat64(p.LastClosePrice),
		OptionIV:       util.ProtoFloat64(p.OptionIV),
		OptionHV:       util.ProtoFloat64(p.OptionHV),
	}
}

// ReportCycleQuoteFromProto converts a proto ReportCycleQuote to our domain type.
func ReportCycleQuoteFromProto(r *qotgetfinancialsearnpricemove.ReportCycleQuote) *ReportCycleQuoteData {
	if r == nil {
		return nil
	}
	data := &ReportCycleQuoteData{
		FiscalYear:       util.ProtoInt32(r.FiscalYear),
		FinancialType:    util.ProtoInt32(r.FinancialType),
		PeriodText:       util.ProtoStr(r.PeriodText),
		PubTradingDay:    util.ProtoInt64(r.PubTradingDay),
		PubTradingDayStr: util.ProtoStr(r.PubTradingDayStr),
		PubType:          r.GetPubType(),
		PriceInfoIndex:   util.ProtoInt32(r.PriceInfoIndex),
	}
	for _, item := range r.ItemList {
		if item == nil {
			continue
		}
		data.ItemList = append(data.ItemList, PricePerformanceRowFromProto(item))
	}
	return data
}

// GetFinancialsEarningsPriceMoveRequest defines parameters for GetFinancialsEarningsPriceMove.
type GetFinancialsEarningsPriceMoveRequest struct {
	Security    *qotcommon.Security
	PeriodCount int32
}

// GetFinancialsEarningsPriceMove retrieves earnings price move data.
func GetFinancialsEarningsPriceMove(ctx context.Context, c *futuapi.Client, req *GetFinancialsEarningsPriceMoveRequest) (*GetFinancialsEarningsPriceMoveResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceMove: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceMove: security is required")
	}

	c2s := &qotgetfinancialsearnpricemove.C2S{
		Security:    req.Security,
		PeriodCount: &req.PeriodCount,
	}
	pkt := &qotgetfinancialsearnpricemove.Request{C2S: c2s}
	var rsp qotgetfinancialsearnpricemove.Response

	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetFinancialsEarningsPriceMove, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsEarningsPriceMove", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsEarningsPriceMove", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetFinancialsEarningsPriceMoveResponse{}
	for _, d := range s2c.DetailList {
		if d == nil {
			continue
		}
		result.DetailList = append(result.DetailList, ReportCycleQuoteFromProto(d))
	}
	return result, nil
}

// PriceInfoData represents price data for the earnings announcement day.
type PriceInfoData struct {
	TradingDay     int64
	TradingDayStr  string
	ClosePrice     float64
	OpenPrice      float64
	HighestPrice   float64
	LowestPrice    float64
	LastClosePrice float64
	Volume         float64
}

// FinScheduleInfoData represents close price at a single trading day offset.
type FinScheduleInfoData struct {
	Delta      int32
	ClosePrice float64
}

// PriceHistoryOnEarningsDaysData represents historical data for one earnings period.
type PriceHistoryOnEarningsDaysData struct {
	FiscalYear              int32
	FinancialType           int32
	PeriodText              string
	IsCurrent               bool
	PubTradingDay           int64
	PubTradingDayStr        string
	PubTime                 int64
	PubTimeStr              string
	PubType                 qotcommon.EarningsPubTimeType
	PredictVolaRatioNewest  float64
	PredictVolaRatioHighest float64
	PredictVolaValNewest    float64
	PredictVolaValHighest   float64
	OptionIVCrush           float64
	OptionStrikeDateIVCrush float64
	PriceInfo               *PriceInfoData
	ScheduleInfoList        []*FinScheduleInfoData
}

// GetFinancialsEarningsPriceHistoryResponse is the response type for GetFinancialsEarningsPriceHistory.
type GetFinancialsEarningsPriceHistoryResponse struct {
	DetailList []*PriceHistoryOnEarningsDaysData
}

// PriceInfoFromProto converts a proto PriceInfo to our domain type.
func PriceInfoFromProto(p *qotgetfinancialsearnpricehist.PriceInfo) *PriceInfoData {
	if p == nil {
		return nil
	}
	return &PriceInfoData{
		TradingDay:     util.ProtoInt64(p.TradingDay),
		TradingDayStr:  util.ProtoStr(p.TradingDayStr),
		ClosePrice:     util.ProtoFloat64(p.ClosePrice),
		OpenPrice:      util.ProtoFloat64(p.OpenPrice),
		HighestPrice:   util.ProtoFloat64(p.HighestPrice),
		LowestPrice:    util.ProtoFloat64(p.LowestPrice),
		LastClosePrice: util.ProtoFloat64(p.LastClosePrice),
		Volume:         util.ProtoFloat64(p.Volume),
	}
}

// FinScheduleInfoFromProto converts a proto FinScheduleInfo to our domain type.
func FinScheduleInfoFromProto(f *qotgetfinancialsearnpricehist.FinScheduleInfo) *FinScheduleInfoData {
	if f == nil {
		return nil
	}
	return &FinScheduleInfoData{
		Delta:      util.ProtoInt32(f.Delta),
		ClosePrice: util.ProtoFloat64(f.ClosePrice),
	}
}

// PriceHistoryFromProto converts a proto PriceHistoryOnEarningsDays to our domain type.
func PriceHistoryFromProto(p *qotgetfinancialsearnpricehist.PriceHistoryOnEarningsDays) *PriceHistoryOnEarningsDaysData {
	if p == nil {
		return nil
	}
	data := &PriceHistoryOnEarningsDaysData{
		FiscalYear:              util.ProtoInt32(p.FiscalYear),
		FinancialType:           util.ProtoInt32(p.FinancialType),
		PeriodText:              util.ProtoStr(p.PeriodText),
		IsCurrent:               util.ProtoBool(p.IsCurrent),
		PubTradingDay:           util.ProtoInt64(p.PubTradingDay),
		PubTradingDayStr:        util.ProtoStr(p.PubTradingDayStr),
		PubTime:                 util.ProtoInt64(p.PubTime),
		PubTimeStr:              util.ProtoStr(p.PubTimeStr),
		PubType:                 p.GetPubType(),
		PredictVolaRatioNewest:  util.ProtoFloat64(p.PredictVolaRatioNewest),
		PredictVolaRatioHighest: util.ProtoFloat64(p.PredictVolaRatioHighest),
		PredictVolaValNewest:    util.ProtoFloat64(p.PredictVolaValNewest),
		PredictVolaValHighest:   util.ProtoFloat64(p.PredictVolaValHighest),
		OptionIVCrush:           util.ProtoFloat64(p.OptionIVCrush),
		OptionStrikeDateIVCrush: util.ProtoFloat64(p.OptionStrikeDateIVCrush),
		PriceInfo:               PriceInfoFromProto(p.PriceInfo),
	}
	for _, f := range p.ScheduleInfoList {
		if f == nil {
			continue
		}
		data.ScheduleInfoList = append(data.ScheduleInfoList, FinScheduleInfoFromProto(f))
	}
	return data
}

// GetFinancialsEarningsPriceHistory retrieves earnings price history data.
func GetFinancialsEarningsPriceHistory(ctx context.Context, c *futuapi.Client, req *GetFinancialsEarningsPriceHistoryRequest) (*GetFinancialsEarningsPriceHistoryResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceHistory: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceHistory: security is required")
	}

	c2s := &qotgetfinancialsearnpricehist.C2S{
		Security: req.Security,
	}
	pkt := &qotgetfinancialsearnpricehist.Request{C2S: c2s}
	var rsp qotgetfinancialsearnpricehist.Response

	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetFinancialsEarningsPriceHistory, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsEarningsPriceHistory", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsEarningsPriceHistory", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetFinancialsEarningsPriceHistoryResponse{}
	for _, d := range s2c.DetailList {
		if d == nil {
			continue
		}
		result.DetailList = append(result.DetailList, PriceHistoryFromProto(d))
	}
	return result, nil
}

// GetFinancialsEarningsPriceHistoryRequest defines parameters for GetFinancialsEarningsPriceHistory.
type GetFinancialsEarningsPriceHistoryRequest struct {
	Security *qotcommon.Security
}