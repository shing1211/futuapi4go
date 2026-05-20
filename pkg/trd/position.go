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

package trd

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdflowsummary"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// AccCashInfo represents per-currency cash information (futures accounts).
type AccCashInfo struct {
	Currency         int32
	Cash             float64
	AvailableBalance float64
	NetCashPower     float64
}

// AccMarketInfo represents per-market asset information.
type AccMarketInfo struct {
	TrdMarket int32
	Assets    float64
}

// Funds represents the capital and asset information of a trading account.
type Funds struct {
	Power             float64
	TotalAssets       float64
	Cash              float64
	MarketVal         float64
	FrozenCash        float64
	DebtCash          float64
	AvlWithdrawalCash float64
	Currency          int32
	AvailableFunds    float64
	UnrealizedPL      float64
	RealizedPL        float64
	RiskLevel         int32
	InitialMargin     float64
	MaintenanceMargin float64
	MaxPowerShort     float64
	NetCashPower      float64
	LongMv            float64
	ShortMv           float64
	PendingAsset      float64
	MaxWithdrawal     float64
	RiskStatus        int32
	MarginCallMargin  float64
	// IsPDT indicates whether the account is a Pattern Day Trader (US margin accounts).
	IsPDT             bool
	// PDTSeq is the PDT sequence number.
	PDTSeq            string
	BeginningDTBP     float64
	RemainingDTBP     float64
	DtCallAmount      float64
	DtStatus          int32
	CashInfoList      []AccCashInfo
	MarketInfoList    []AccMarketInfo
	SecuritiesAssets  float64
	FundAssets        float64
	BondAssets        float64
}

// GetFundsRequest is the request to retrieve account funds.
type GetFundsRequest struct {
	AccID         uint64
	TrdMarket     constant.TrdMarket
	TrdEnv        constant.TrdEnv
	RefreshCache  bool
	Currency      int32
	AssetCategory int32
}

// GetFundsResponse is the response containing account funds information.
type GetFundsResponse struct {
	Funds *Funds
}

// GetFunds retrieves the account funds information including cash, assets, and available funds.
// Returns the funds data or an error if the request fails.
func GetFunds(ctx context.Context, c *futuapi.Client, req *GetFundsRequest) (*GetFundsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFunds: request is nil")
	}
	if req.AccID == 0 {
		return nil, constant.ErrInvalidAccID
	}
	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetfunds.C2S{
		Header: header,
	}
	if req.RefreshCache {
		c2s.RefreshCache = &req.RefreshCache
	}
	if req.Currency != 0 {
		c2s.Currency = &req.Currency
	}
	if req.AssetCategory != 0 {
		c2s.AssetCategory = &req.AssetCategory
	}

	pkt := &trdgetfunds.Request{C2S: c2s}
	var rsp trdgetfunds.Response

	if err := c.RequestContext(ctx, ProtoID_GetFunds, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFunds", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFunds", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	f := s2c.Funds
	if f == nil {
		return nil, fmt.Errorf("GetFunds: funds is nil")
	}
	return &GetFundsResponse{
		Funds: &Funds{
			Power:             util.ProtoFloat64(f.Power),
			TotalAssets:       util.ProtoFloat64(f.TotalAssets),
			Cash:              util.ProtoFloat64(f.Cash),
			MarketVal:         util.ProtoFloat64(f.MarketVal),
			FrozenCash:        util.ProtoFloat64(f.FrozenCash),
			DebtCash:          util.ProtoFloat64(f.DebtCash),
			AvlWithdrawalCash: util.ProtoFloat64(f.AvlWithdrawalCash),
			Currency:          util.ProtoInt32(f.Currency),
			AvailableFunds:    util.ProtoFloat64(f.AvailableFunds),
			UnrealizedPL:      util.ProtoFloat64(f.UnrealizedPL),
			RealizedPL:        util.ProtoFloat64(f.RealizedPL),
			RiskLevel:         util.ProtoInt32(f.RiskLevel),
			InitialMargin:     util.ProtoFloat64(f.InitialMargin),
			MaintenanceMargin: util.ProtoFloat64(f.MaintenanceMargin),
			MaxPowerShort:     util.ProtoFloat64(f.MaxPowerShort),
			NetCashPower:      util.ProtoFloat64(f.NetCashPower),
			LongMv:            util.ProtoFloat64(f.LongMv),
			ShortMv:           util.ProtoFloat64(f.ShortMv),
			PendingAsset:      util.ProtoFloat64(f.PendingAsset),
			MaxWithdrawal:     util.ProtoFloat64(f.MaxWithdrawal),
			RiskStatus:        util.ProtoInt32(f.RiskStatus),
			MarginCallMargin:  util.ProtoFloat64(f.MarginCallMargin),
			IsPDT:             util.ProtoBool(f.IsPdt),
			PDTSeq:            util.ProtoStr(f.PdtSeq),
			BeginningDTBP:     util.ProtoFloat64(f.BeginningDTBP),
			RemainingDTBP:     util.ProtoFloat64(f.RemainingDTBP),
			DtCallAmount:      util.ProtoFloat64(f.DtCallAmount),
			DtStatus:          util.ProtoInt32(f.DtStatus),
			CashInfoList:      accCashInfoListToGo(f.CashInfoList),
			MarketInfoList:    accMarketInfoListToGo(f.MarketInfoList),
			SecuritiesAssets:  util.ProtoFloat64(f.SecuritiesAssets),
			FundAssets:        util.ProtoFloat64(f.FundAssets),
			BondAssets:        util.ProtoFloat64(f.BondAssets),
		},
	}, nil
}

func accCashInfoListToGo(in []*trdcommon.AccCashInfo) []AccCashInfo {
	out := make([]AccCashInfo, 0, len(in))
	for _, c := range in {
		if c == nil {
			continue
		}
		out = append(out, AccCashInfo{
			Currency:         util.ProtoInt32(c.Currency),
			Cash:             util.ProtoFloat64(c.Cash),
			AvailableBalance: util.ProtoFloat64(c.AvailableBalance),
			NetCashPower:     util.ProtoFloat64(c.NetCashPower),
		})
	}
	return out
}

func accMarketInfoListToGo(in []*trdcommon.AccMarketInfo) []AccMarketInfo {
	out := make([]AccMarketInfo, 0, len(in))
	for _, m := range in {
		if m == nil {
			continue
		}
		out = append(out, AccMarketInfo{
			TrdMarket: util.ProtoInt32(m.TrdMarket),
			Assets:    util.ProtoFloat64(m.Assets),
		})
	}
	return out
}

// Position represents a stock position with quantity, price, cost, and profit/loss information.
type Position struct {
	PositionID       uint64
	PositionSide     int32
	Code             string
	Name             string
	Qty              float64
	CanSellQty       float64
	Price            float64
	CostPrice        float64
	Val              float64
	PlVal            float64
	PlRatio          float64
	SecMarket        int32
	TdPlVal          float64
	TdTrdVal         float64
	TdBuyVal         float64
	TdBuyQty         float64
	TdSellVal        float64
	TdSellQty        float64
	UnrealizedPL     float64
	RealizedPL       float64
	Currency         int32
	TrdMarket        int32
	DilutedCostPrice float64
	AverageCostPrice float64
	AveragePlRatio   float64
}

// GetPositionListRequest is the request to retrieve position list.
type GetPositionListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
	FilterPLRatioMin float64
	FilterPLRatioMax float64
	RefreshCache     bool
	AssetCategory    int32
}

// GetPositionListResponse is the response containing a list of positions.
type GetPositionListResponse struct {
	PositionList []*Position
}

// GetPositionList retrieves the current position list for the account.
// Returns the position list or an error if the request fails.
func GetPositionList(ctx context.Context, c *futuapi.Client, req *GetPositionListRequest) (*GetPositionListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetPositionList: request is nil")
	}
	if req.AccID == 0 && req.TrdEnv == 0 && req.TrdMarket == 0 {
		return nil, fmt.Errorf("GetPositionList: at least one of AccID, TrdEnv, or TrdMarket must be set")
	}
	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetpositionlist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
	}
	if req.FilterPLRatioMin != 0 {
		c2s.FilterPLRatioMin = &req.FilterPLRatioMin
	}
	if req.FilterPLRatioMax != 0 {
		c2s.FilterPLRatioMax = &req.FilterPLRatioMax
	}
	if req.RefreshCache {
		c2s.RefreshCache = &req.RefreshCache
	}
	if req.AssetCategory != 0 {
		c2s.AssetCategory = &req.AssetCategory
	}

	pkt := &trdgetpositionlist.Request{C2S: c2s}
	var rsp trdgetpositionlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetPositionList, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPositionList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetPositionList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetPositionListResponse{
		PositionList: make([]*Position, 0, len(s2c.PositionList)),
	}

	for _, p := range s2c.PositionList {
		if p == nil {
			continue
		}
		result.PositionList = append(result.PositionList, &Position{
			PositionID:       util.ProtoUint64(p.PositionID),
			PositionSide:     util.ProtoInt32(p.PositionSide),
			Code:             util.ProtoStr(p.Code),
			Name:             util.ProtoStr(p.Name),
			Qty:              util.ProtoFloat64(p.Qty),
			CanSellQty:       util.ProtoFloat64(p.CanSellQty),
			Price:            util.ProtoFloat64(p.Price),
			CostPrice:        util.ProtoFloat64(p.CostPrice),
			Val:              util.ProtoFloat64(p.Val),
			PlVal:            util.ProtoFloat64(p.PlVal),
			PlRatio:          util.ProtoFloat64(p.PlRatio),
			SecMarket:        util.ProtoInt32(p.SecMarket),
			TdPlVal:          util.ProtoFloat64(p.TdPlVal),
			TdTrdVal:         util.ProtoFloat64(p.TdTrdVal),
			TdBuyVal:         util.ProtoFloat64(p.TdBuyVal),
			TdBuyQty:         util.ProtoFloat64(p.TdBuyQty),
			TdSellVal:        util.ProtoFloat64(p.TdSellVal),
			TdSellQty:        util.ProtoFloat64(p.TdSellQty),
			UnrealizedPL:     util.ProtoFloat64(p.UnrealizedPL),
			RealizedPL:       util.ProtoFloat64(p.RealizedPL),
			Currency:         util.ProtoInt32(p.Currency),
			TrdMarket:        util.ProtoInt32(p.TrdMarket),
			DilutedCostPrice: util.ProtoFloat64(p.DilutedCostPrice),
			AverageCostPrice: util.ProtoFloat64(p.AverageCostPrice),
			AveragePlRatio:   util.ProtoFloat64(p.AveragePlRatio),
		})
	}

	return result, nil
}

// GetMarginRatioRequest is the request to retrieve margin ratio information.
type GetMarginRatioRequest struct {
	AccID        uint64
	TrdMarket    constant.TrdMarket
	TrdEnv       constant.TrdEnv
	SecurityList []*qotcommon.Security
}

// MarginRatioInfo represents margin ratio information for a security, including long/short permits and fee rates.
type MarginRatioInfo struct {
	Security        *qotcommon.Security
	IsLongPermit    bool
	IsShortPermit   bool
	ShortPoolRemain float64
	ShortFeeRate    float64
	AlertLongRatio  float64
	AlertShortRatio float64
	ImLongRatio     float64
	ImShortRatio    float64
	McmLongRatio    float64
	McmShortRatio   float64
	MmLongRatio     float64
	MmShortRatio    float64
}

// GetMarginRatioResponse is the response containing margin ratio information.
type GetMarginRatioResponse struct {
	MarginRatioInfoList []*MarginRatioInfo
}

// GetMarginRatio retrieves the margin ratio information for specified securities.
// Returns the margin ratio list or an error if the request fails.
func GetMarginRatio(ctx context.Context, c *futuapi.Client, req *GetMarginRatioRequest) (*GetMarginRatioResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMarginRatio: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetmarginratio.C2S{
		Header:       header,
		SecurityList: req.SecurityList,
	}

	pkt := &trdgetmarginratio.Request{C2S: c2s}
	var rsp trdgetmarginratio.Response

	if err := c.RequestContext(ctx, ProtoID_GetMarginRatio, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMarginRatio", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetMarginRatio", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetMarginRatioResponse{
		MarginRatioInfoList: make([]*MarginRatioInfo, 0, len(s2c.MarginRatioInfoList)),
	}

	for _, m := range s2c.MarginRatioInfoList {
		if m == nil {
			continue
		}
		result.MarginRatioInfoList = append(result.MarginRatioInfoList, &MarginRatioInfo{
			Security:        m.Security,
			IsLongPermit:    util.ProtoBool(m.IsLongPermit),
			IsShortPermit:   util.ProtoBool(m.IsShortPermit),
			ShortPoolRemain: util.ProtoFloat64(m.ShortPoolRemain),
			ShortFeeRate:    util.ProtoFloat64(m.ShortFeeRate),
			AlertLongRatio:  util.ProtoFloat64(m.AlertLongRatio),
			AlertShortRatio: util.ProtoFloat64(m.AlertShortRatio),
			ImLongRatio:     util.ProtoFloat64(m.ImLongRatio),
			ImShortRatio:    util.ProtoFloat64(m.ImShortRatio),
			McmLongRatio:    util.ProtoFloat64(m.McmLongRatio),
			McmShortRatio:   util.ProtoFloat64(m.McmShortRatio),
			MmLongRatio:     util.ProtoFloat64(m.MmLongRatio),
			MmShortRatio:    util.ProtoFloat64(m.MmShortRatio),
		})
	}

	return result, nil
}

// GetMaxTrdQtysRequest is the request to retrieve maximum tradable quantities.
type GetMaxTrdQtysRequest struct {
	AccID              uint64
	TrdMarket          constant.TrdMarket
	TrdEnv             constant.TrdEnv
	OrderType          constant.OrderType
	Code               string
	Price              float64
	OrderID            uint64
	AdjustPrice        bool
	AdjustSideAndLimit float64
	SecMarket          constant.TrdSecMarket
	OrderIDEx          string
	Session            int32
	PositionID         uint64
}

// MaxTrdQtysInfo represents the maximum tradable quantities for various trading scenarios.
type MaxTrdQtysInfo struct {
	MaxCashBuy          float64
	MaxCashAndMarginBuy float64
	MaxPositionSell     float64
	MaxSellShort        float64
	MaxBuyBack          float64
	LongRequiredIM      float64
	ShortRequiredIM     float64
	Session             int32
}

// GetMaxTrdQtysResponse is the response containing maximum tradable quantities.
type GetMaxTrdQtysResponse struct {
	MaxTrdQtys *MaxTrdQtysInfo
}

// GetMaxTrdQtys retrieves the maximum tradable quantities for a given order.
// Returns the maximum quantities or an error if the request fails.
func GetMaxTrdQtys(ctx context.Context, c *futuapi.Client, req *GetMaxTrdQtysRequest) (*GetMaxTrdQtysResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMaxTrdQtys: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.Code == "" {
		return nil, fmt.Errorf("security code is required")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)
	orderType := int32(req.OrderType)
	secMarket := int32(req.SecMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetmaxtrdqtys.C2S{
		Header:    header,
		OrderType: &orderType,
		Code:      &req.Code,
		Price:     &req.Price,
	}
	if req.OrderID != 0 {
		c2s.OrderID = &req.OrderID
	}
	if req.AdjustPrice {
		c2s.AdjustPrice = &req.AdjustPrice
	}
	if req.AdjustSideAndLimit != 0 {
		c2s.AdjustSideAndLimit = &req.AdjustSideAndLimit
	}
	if req.SecMarket != 0 {
		c2s.SecMarket = &secMarket
	}
	if req.OrderIDEx != "" {
		c2s.OrderIDEx = &req.OrderIDEx
	}
	if req.Session != 0 {
		c2s.Session = &req.Session
	}
	if req.PositionID != 0 {
		c2s.PositionID = &req.PositionID
	}

	pkt := &trdgetmaxtrdqtys.Request{C2S: c2s}
	var rsp trdgetmaxtrdqtys.Response

	if err := c.RequestContext(ctx, ProtoID_GetMaxTrdQtys, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMaxTrdQtys", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetMaxTrdQtys", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	m := s2c.MaxTrdQtys
	if m == nil {
		return nil, fmt.Errorf("GetMaxTrdQtys: maxTrdQtys is nil")
	}
	return &GetMaxTrdQtysResponse{
		MaxTrdQtys: &MaxTrdQtysInfo{
			MaxCashBuy:          util.ProtoFloat64(m.MaxCashBuy),
			MaxCashAndMarginBuy: util.ProtoFloat64(m.MaxCashAndMarginBuy),
			MaxPositionSell:     util.ProtoFloat64(m.MaxPositionSell),
			MaxSellShort:        util.ProtoFloat64(m.MaxSellShort),
			MaxBuyBack:          util.ProtoFloat64(m.MaxBuyBack),
			LongRequiredIM:      util.ProtoFloat64(m.LongRequiredIM),
			ShortRequiredIM:     util.ProtoFloat64(m.ShortRequiredIM),
			Session:             util.ProtoInt32(m.Session),
		},
	}, nil
}

// GetFlowSummaryRequest is the request to retrieve fund flow summary for a clearing date.
type GetFlowSummaryRequest struct {
	AccID             uint64
	TrdMarket         constant.TrdMarket
	TrdEnv            constant.TrdEnv
	ClearingDate      string
	CashFlowDirection int32
}

// FlowSummaryInfo represents a single cash flow entry.
type FlowSummaryInfo struct {
	CashFlowID        uint64
	ClearingDate      string
	SettlementDate    string
	Currency          int32
	CashFlowType      string
	CashFlowDirection int32
	CashFlowAmount    float64
	CashFlowRemark    string
}

// flowSummaryInfoFromProto converts a proto FlowSummaryInfo to a wrapped type.
func flowSummaryInfoFromProto(f *trdflowsummary.FlowSummaryInfo) *FlowSummaryInfo {
	if f == nil {
		return nil
	}
	return &FlowSummaryInfo{
		CashFlowID:        util.ProtoUint64(f.CashFlowID),
		ClearingDate:      util.ProtoStr(f.ClearingDate),
		SettlementDate:    util.ProtoStr(f.SettlementDate),
		Currency:          util.ProtoInt32(f.Currency),
		CashFlowType:      util.ProtoStr(f.CashFlowType),
		CashFlowDirection: util.ProtoInt32(f.CashFlowDirection),
		CashFlowAmount:    util.ProtoFloat64(f.CashFlowAmount),
		CashFlowRemark:    util.ProtoStr(f.CashFlowRemark),
	}
}

// GetFlowSummaryResponse is the response containing the fund flow summary.
type GetFlowSummaryResponse struct {
	FlowSummaryList []*FlowSummaryInfo
}

// GetFlowSummary retrieves the fund flow summary for a specified clearing date.
// Returns the flow summary list or an error if the request fails.
func GetFlowSummary(ctx context.Context, c *futuapi.Client, req *GetFlowSummaryRequest) (*GetFlowSummaryResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFlowSummary: request is nil")
	}
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.ClearingDate == "" && req.CashFlowDirection != 0 {
		return nil, fmt.Errorf("clearing date is required when cash flow direction is specified")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
		TrdEnv:    &trdEnv,
	}

	c2s := &trdflowsummary.C2S{
		Header:            header,
		ClearingDate:      &req.ClearingDate,
		CashFlowDirection: &req.CashFlowDirection,
	}

	pkt := &trdflowsummary.Request{C2S: c2s}
	var rsp trdflowsummary.Response

	if err := c.RequestContext(ctx, ProtoID_GetFlowSummary, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFlowSummary", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFlowSummary", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	flowSummaryList := s2c.FlowSummaryInfoList
	result := make([]*FlowSummaryInfo, 0, len(flowSummaryList))
	for _, item := range flowSummaryList {
		if item == nil {
			continue
		}
		result = append(result, flowSummaryInfoFromProto(item))
	}
	return &GetFlowSummaryResponse{
		FlowSummaryList: result,
	}, nil
}
