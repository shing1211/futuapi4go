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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFunds", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetFunds: s2c is nil")
	}

	f := s2c.GetFunds()
	if f == nil {
		return nil, fmt.Errorf("GetFunds: funds is nil")
	}
	return &GetFundsResponse{
		Funds: &Funds{
			Power:             f.GetPower(),
			TotalAssets:       f.GetTotalAssets(),
			Cash:              f.GetCash(),
			MarketVal:         f.GetMarketVal(),
			FrozenCash:        f.GetFrozenCash(),
			DebtCash:          f.GetDebtCash(),
			AvlWithdrawalCash: f.GetAvlWithdrawalCash(),
			Currency:          f.GetCurrency(),
			AvailableFunds:    f.GetAvailableFunds(),
			UnrealizedPL:      f.GetUnrealizedPL(),
			RealizedPL:        f.GetRealizedPL(),
			RiskLevel:         f.GetRiskLevel(),
			InitialMargin:     f.GetInitialMargin(),
			MaintenanceMargin: f.GetMaintenanceMargin(),
			MaxPowerShort:     f.GetMaxPowerShort(),
			NetCashPower:      f.GetNetCashPower(),
			LongMv:            f.GetLongMv(),
			ShortMv:           f.GetShortMv(),
			PendingAsset:      f.GetPendingAsset(),
			MaxWithdrawal:     f.GetMaxWithdrawal(),
			RiskStatus:        f.GetRiskStatus(),
			MarginCallMargin:  f.GetMarginCallMargin(),
			IsPDT:             f.GetIsPdt(),
			PDTSeq:            f.GetPdtSeq(),
			BeginningDTBP:     f.GetBeginningDTBP(),
			RemainingDTBP:     f.GetRemainingDTBP(),
			DtCallAmount:      f.GetDtCallAmount(),
			DtStatus:          f.GetDtStatus(),
			CashInfoList:      accCashInfoListToGo(f.GetCashInfoList()),
			MarketInfoList:    accMarketInfoListToGo(f.GetMarketInfoList()),
			SecuritiesAssets:  f.GetSecuritiesAssets(),
			FundAssets:        f.GetFundAssets(),
			BondAssets:        f.GetBondAssets(),
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
			Currency:         c.GetCurrency(),
			Cash:             c.GetCash(),
			AvailableBalance: c.GetAvailableBalance(),
			NetCashPower:     c.GetNetCashPower(),
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
			TrdMarket: m.GetTrdMarket(),
			Assets:    m.GetAssets(),
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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPositionList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPositionList: s2c is nil")
	}

	result := &GetPositionListResponse{
		PositionList: make([]*Position, 0, len(s2c.GetPositionList())),
	}

	for _, p := range s2c.GetPositionList() {
		if p == nil {
			continue
		}
		result.PositionList = append(result.PositionList, &Position{
			PositionID:       p.GetPositionID(),
			PositionSide:     p.GetPositionSide(),
			Code:             p.GetCode(),
			Name:             p.GetName(),
			Qty:              p.GetQty(),
			CanSellQty:       p.GetCanSellQty(),
			Price:            p.GetPrice(),
			CostPrice:        p.GetCostPrice(),
			Val:              p.GetVal(),
			PlVal:            p.GetPlVal(),
			PlRatio:          p.GetPlRatio(),
			SecMarket:        p.GetSecMarket(),
			TdPlVal:          p.GetTdPlVal(),
			TdTrdVal:         p.GetTdTrdVal(),
			TdBuyVal:         p.GetTdBuyVal(),
			TdBuyQty:         p.GetTdBuyQty(),
			TdSellVal:        p.GetTdSellVal(),
			TdSellQty:        p.GetTdSellQty(),
			UnrealizedPL:     p.GetUnrealizedPL(),
			RealizedPL:       p.GetRealizedPL(),
			Currency:         p.GetCurrency(),
			TrdMarket:        p.GetTrdMarket(),
			DilutedCostPrice: p.GetDilutedCostPrice(),
			AverageCostPrice: p.GetAverageCostPrice(),
			AveragePlRatio:   p.GetAveragePlRatio(),
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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMarginRatio", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetMarginRatio: s2c is nil")
	}

	result := &GetMarginRatioResponse{
		MarginRatioInfoList: make([]*MarginRatioInfo, 0, len(s2c.GetMarginRatioInfoList())),
	}

	for _, m := range s2c.GetMarginRatioInfoList() {
		if m == nil {
			continue
		}
		result.MarginRatioInfoList = append(result.MarginRatioInfoList, &MarginRatioInfo{
			Security:        m.GetSecurity(),
			IsLongPermit:    m.GetIsLongPermit(),
			IsShortPermit:   m.GetIsShortPermit(),
			ShortPoolRemain: m.GetShortPoolRemain(),
			ShortFeeRate:    m.GetShortFeeRate(),
			AlertLongRatio:  m.GetAlertLongRatio(),
			AlertShortRatio: m.GetAlertShortRatio(),
			ImLongRatio:     m.GetImLongRatio(),
			ImShortRatio:    m.GetImShortRatio(),
			McmLongRatio:    m.GetMcmLongRatio(),
			McmShortRatio:   m.GetMcmShortRatio(),
			MmLongRatio:     m.GetMmLongRatio(),
			MmShortRatio:    m.GetMmShortRatio(),
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

	pkt := &trdgetmaxtrdqtys.Request{C2S: c2s}
	var rsp trdgetmaxtrdqtys.Response

	if err := c.RequestContext(ctx, ProtoID_GetMaxTrdQtys, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMaxTrdQtys", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetMaxTrdQtys: s2c is nil")
	}

	m := s2c.GetMaxTrdQtys()
	if m == nil {
		return nil, fmt.Errorf("GetMaxTrdQtys: maxTrdQtys is nil")
	}
	return &GetMaxTrdQtysResponse{
		MaxTrdQtys: &MaxTrdQtysInfo{
			MaxCashBuy:          m.GetMaxCashBuy(),
			MaxCashAndMarginBuy: m.GetMaxCashAndMarginBuy(),
			MaxPositionSell:     m.GetMaxPositionSell(),
			MaxSellShort:        m.GetMaxSellShort(),
			MaxBuyBack:          m.GetMaxBuyBack(),
			LongRequiredIM:      m.GetLongRequiredIM(),
			ShortRequiredIM:     m.GetShortRequiredIM(),
			Session:             m.GetSession(),
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

// GetFlowSummaryResponse is the response containing the fund flow summary.
type GetFlowSummaryResponse struct {
	FlowSummaryList []*trdflowsummary.FlowSummaryInfo
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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFlowSummary", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetFlowSummary: s2c is nil")
	}

	return &GetFlowSummaryResponse{
		FlowSummaryList: s2c.GetFlowSummaryInfoList(),
	}, nil
}
