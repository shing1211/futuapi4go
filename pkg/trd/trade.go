// Package trd provides trading APIs for the Futu OpenD SDK.
//
// This package covers account management, order placement and modification,
// position and funds queries, order history, and trading flow analysis.
// All trading functions require an unlocked trading account.
//
// For Python SDK migration, use the constant package for Python-style constants:
//
//	import (
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/trd"
//	)
//
//	// Trading environment: constant.TrdEnv_Real or constant.TrdEnv_Simulate
//	// Trade side: constant.TrdSide_Buy, constant.TrdSide_Sell
//	// Order type: constant.OrderType_Normal, constant.OrderType_Market
//	// TrdMarket: constant.TrdMarket_HK, constant.TrdMarket_US, etc.
//
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

// accs, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
// req := &trd.PlaceOrderRequest{
package trd

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdflowsummary"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdreconfirmorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdsubaccpush"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"

	"time"
)

const (
	ProtoID_GetAccList              = 2001
	ProtoID_UnlockTrade             = 2005
	ProtoID_GetFunds                = 2101
	ProtoID_GetOrderFee             = 2225
	ProtoID_GetMarginRatio          = 2223
	ProtoID_GetMaxTrdQtys           = 2111
	ProtoID_GetPositionList         = 2102
	ProtoID_GetOrderList            = 2201
	ProtoID_GetOrderFillList        = 2211
	ProtoID_GetHistoryOrderList     = 2221
	ProtoID_GetHistoryOrderFillList = 2222
	ProtoID_PlaceOrder              = 2202
	ProtoID_ModifyOrder             = 2205
	ProtoID_UpdateOrder             = 2208
	ProtoID_UpdateOrderFill         = 2218
	ProtoID_SubAccPush              = 2008
	ProtoID_ReconfirmOrder          = 2209
	ProtoID_GetFlowSummary          = 2226
)

// Acc represents a trading account with its environment, ID, type, and status.
type Acc struct {
	TrdEnv            int32
	AccID             uint64
	AccType           int32
	CardNum           string
	AccStatus         int32
	TrdMarketAuthList []int32
	SecurityFirm      int32
	SimAccType        int32
	UniCardNum        string
	AccRole           int32
	JpAccType         []int32
}

// GetAccListResponse is the response containing a list of trading accounts.
type GetAccListResponse struct {
	AccList []*Acc
}

// wrapError standardizes error messages for proto response failures
func wrapError(funcName string, retType int32, retMsg string) error {
	return fmt.Errorf("%s failed: retType=%d, retMsg=%s", funcName, retType, retMsg)
}

// GetAccList retrieves the list of trading accounts, optionally including general security account info.
// Returns the account list or an error if the request fails.
func GetAccList(ctx context.Context, c *futuapi.Client, trdCategory constant.TrdCategory, needGeneralSecAccount bool) (*GetAccListResponse, error) {
	trdCategoryInt := int32(trdCategory)
	c2s := &trdgetacclist.C2S{
		UserID:                proto.Uint64(0), // Deprecated but required by protocol, set to 0
		TrdCategory:           &trdCategoryInt,
		NeedGeneralSecAccount: &needGeneralSecAccount,
	}

	pkt := &trdgetacclist.Request{C2S: c2s}
	var rsp trdgetacclist.Response

	if err := c.RequestContext(ctx, ProtoID_GetAccList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetAccList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetAccList: s2c is nil")
	}

	result := &GetAccListResponse{
		AccList: make([]*Acc, 0, len(s2c.GetAccList())),
	}

	for _, acc := range s2c.GetAccList() {
		if acc == nil {
			continue
		}
		result.AccList = append(result.AccList, &Acc{
			TrdEnv:            acc.GetTrdEnv(),
			AccID:             acc.GetAccID(),
			AccType:           acc.GetAccType(),
			CardNum:           acc.GetCardNum(),
			AccStatus:         acc.GetAccStatus(),
			TrdMarketAuthList: acc.GetTrdMarketAuthList(),
			SecurityFirm:      acc.GetSecurityFirm(),
			SimAccType:        acc.GetSimAccType(),
			UniCardNum:        acc.GetUniCardNum(),
			AccRole:           acc.GetAccRole(),
			JpAccType:         acc.GetJpAccType(),
		})
	}

	return result, nil
}

// AccCashInfo represents per-currency cash information (futures accounts).
type AccCashInfo struct {
	Currency        int32
	Cash            float64
	AvailableBalance float64
	NetCashPower    float64
}

// AccMarketInfo represents per-market asset information.
type AccMarketInfo struct {
	TrdMarket int32
	Assets    float64
}

// Funds represents the capital and asset information of a trading account.
// Maps to Python's accinfo_query return columns.
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
	IsPDT             bool
	PDTSeq            string
	BeginningDTBP     float64
	RemainingDTBP     float64
	DtCallAmount      float64
	DtStatus          int32
	CashInfoList      []AccCashInfo
	MarketInfoList    []AccMarketInfo
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
			Currency:        c.GetCurrency(),
			Cash:            c.GetCash(),
			AvailableBalance: c.GetAvailableBalance(),
			NetCashPower:    c.GetNetCashPower(),
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

// Order represents an order with its ID, code, side, type, status, price, quantity, and fill information.
type Order struct {
	OrderID         uint64
	OrderIDEx       string
	Code            string
	Name            string
	TrdSide         int32
	OrderType       int32
	OrderStatus     int32
	Price           float64
	Qty             float64
	FillQty         float64
	FillAvgPrice    float64
	CreateTime      string
	UpdateTime      string
	LastErrMsg      string
	SecMarket       int32
	CreateTimestamp float64
	UpdateTimestamp float64
	Remark          string
	TimeInForce     int32
	FillOutsideRTH  bool
	AuxPrice        float64
	TrailType       int32
	TrailValue      float64
	TrailSpread     float64
	Currency        int32
	TrdMarket       int32
	Session         int32
	JpAccType       int32
}

// GetOrderListRequest is the request to retrieve order list.
type GetOrderListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
	FilterStatusList []int32
	RefreshCache     bool
}

// GetOrderListResponse is the response containing a list of orders.
type GetOrderListResponse struct {
	OrderList []*Order
}

// GetOrderList retrieves the current order list for the account.
// Returns the order list or an error if the request fails.
func GetOrderList(ctx context.Context, c *futuapi.Client, req *GetOrderListRequest) (*GetOrderListResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderlist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
		FilterStatusList: req.FilterStatusList,
	}
	if req.RefreshCache {
		c2s.RefreshCache = &req.RefreshCache
	}

	pkt := &trdgetorderlist.Request{C2S: c2s}
	var rsp trdgetorderlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderList: s2c is nil")
	}

	result := &GetOrderListResponse{
		OrderList: make([]*Order, 0, len(s2c.GetOrderList())),
	}

	for _, o := range s2c.GetOrderList() {
		if o == nil {
			continue
		}
		result.OrderList = append(result.OrderList, &Order{
			OrderID:         o.GetOrderID(),
			OrderIDEx:       o.GetOrderIDEx(),
			Code:            o.GetCode(),
			Name:            o.GetName(),
			TrdSide:         o.GetTrdSide(),
			OrderType:       o.GetOrderType(),
			OrderStatus:     o.GetOrderStatus(),
			Price:           o.GetPrice(),
			Qty:             o.GetQty(),
			FillQty:         o.GetFillQty(),
			FillAvgPrice:    o.GetFillAvgPrice(),
			CreateTime:      o.GetCreateTime(),
			UpdateTime:      o.GetUpdateTime(),
			LastErrMsg:      o.GetLastErrMsg(),
			SecMarket:       o.GetSecMarket(),
			CreateTimestamp: o.GetCreateTimestamp(),
			UpdateTimestamp: o.GetUpdateTimestamp(),
			Remark:          o.GetRemark(),
			TimeInForce:     o.GetTimeInForce(),
			FillOutsideRTH:  o.GetFillOutsideRTH(),
			AuxPrice:        o.GetAuxPrice(),
			TrailType:       o.GetTrailType(),
			TrailValue:      o.GetTrailValue(),
			TrailSpread:     o.GetTrailSpread(),
			Currency:        o.GetCurrency(),
			TrdMarket:       o.GetTrdMarket(),
			Session:         o.GetSession(),
			JpAccType:       o.GetJpAccType(),
		})
	}

	return result, nil
}

// OrderFill represents a filled (executed) order with its order ID, fill ID, code, side, price, and quantity.
type OrderFill struct {
	FillID            uint64
	FillIDEx          string
	OrderID           uint64
	OrderIDEx         string
	Code              string
	Name              string
	TrdSide           int32
	Price             float64
	Qty               float64
	CreateTime        string
	CounterBrokerID   int32
	CounterBrokerName string
	SecMarket         int32
	CreateTimestamp   float64
	UpdateTimestamp   float64
	Status            int32
	TrdMarket         int32
	JpAccType         int32
}

// GetOrderFillListRequest is the request to retrieve order fill list.
type GetOrderFillListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
}

// GetOrderFillListResponse is the response containing a list of order fills.
type GetOrderFillListResponse struct {
	OrderFillList []*OrderFill
}

// GetOrderFillList retrieves the current order fill (execution) list for the account.
// Returns the order fill list or an error if the request fails.
func GetOrderFillList(ctx context.Context, c *futuapi.Client, req *GetOrderFillListRequest) (*GetOrderFillListResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderfilllist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
	}

	pkt := &trdgetorderfilllist.Request{C2S: c2s}
	var rsp trdgetorderfilllist.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderFillList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderFillList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderFillList: s2c is nil")
	}

	result := &GetOrderFillListResponse{
		OrderFillList: make([]*OrderFill, 0, len(s2c.GetOrderFillList())),
	}

	for _, f := range s2c.GetOrderFillList() {
		if f == nil {
			continue
		}
		result.OrderFillList = append(result.OrderFillList, &OrderFill{
			FillID:            f.GetFillID(),
			FillIDEx:          f.GetFillIDEx(),
			OrderID:           f.GetOrderID(),
			OrderIDEx:         f.GetOrderIDEx(),
			Code:              f.GetCode(),
			Name:              f.GetName(),
			TrdSide:           f.GetTrdSide(),
			Price:             f.GetPrice(),
			Qty:               f.GetQty(),
			CreateTime:        f.GetCreateTime(),
			CounterBrokerID:   f.GetCounterBrokerID(),
			CounterBrokerName: f.GetCounterBrokerName(),
			SecMarket:         f.GetSecMarket(),
			CreateTimestamp:   f.GetCreateTimestamp(),
			UpdateTimestamp:   f.GetUpdateTimestamp(),
			Status:            f.GetStatus(),
			TrdMarket:         f.GetTrdMarket(),
			JpAccType:         f.GetJpAccType(),
		})
	}

	return result, nil
}

// PlaceOrderRequest is the request to place a new order.
type PlaceOrderRequest struct {
	AccID              uint64
	TrdMarket          constant.TrdMarket
	TrdEnv             constant.TrdEnv
	Code               string
	TrdSide            constant.TrdSide
	OrderType          constant.OrderType
	Price              float64
	Qty                float64
	AdjustPrice        bool
	AdjustSideAndLimit float64
	SecMarket          constant.TrdSecMarket
	Remark             string
	TimeInForce        int32
	FillOutsideRTH     bool
	AuxPrice           float64
	TrailType          constant.TrailType
	TrailValue         float64
	TrailSpread        float64
	Session            int32
	PositionID         uint64
}

// PlaceOrderResponse is the response containing the newly placed order ID.
type PlaceOrderResponse struct {
	OrderID   uint64
	OrderIDEx string
}

// PlaceOrder places a new order and returns the order ID.
// Returns the order ID or an error if the placement fails.
func PlaceOrder(ctx context.Context, c *futuapi.Client, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.Code == "" {
		return nil, fmt.Errorf("stock code is required")
	}
	if req.Qty <= 0 {
		return nil, fmt.Errorf("invalid quantity: must be positive")
	}
	if req.OrderType <= 0 {
		return nil, fmt.Errorf("invalid order type: must be valid order type constant")
	}
	if req.TrdSide <= 0 {
		return nil, fmt.Errorf("invalid trade side: must be buy/sell/other valid type")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)
	trdSide := int32(req.TrdSide)
	orderType := int32(req.OrderType)
	secMarket := int32(req.SecMarket)
	trailType := int32(req.TrailType)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdplaceorder.C2S{
		Header:    header,
		TrdSide:   &trdSide,
		OrderType: &orderType,
		Code:      &req.Code,
		Qty:       &req.Qty,
		PacketID: &common.PacketID{
			ConnID: proto.Uint64(c.GetConnID()),
		},
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
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
	if req.Remark != "" {
		c2s.Remark = &req.Remark
	}
	if req.TimeInForce != 0 {
		c2s.TimeInForce = &req.TimeInForce
	}
	if req.FillOutsideRTH {
		c2s.FillOutsideRTH = &req.FillOutsideRTH
	}
	if req.AuxPrice != 0 {
		c2s.AuxPrice = &req.AuxPrice
	}
	if req.TrailType != 0 {
		c2s.TrailType = &trailType
	}
	if req.TrailValue != 0 {
		c2s.TrailValue = &req.TrailValue
	}
	if req.TrailSpread != 0 {
		c2s.TrailSpread = &req.TrailSpread
	}
	if req.Session != 0 {
		c2s.Session = &req.Session
	}
	if req.PositionID != 0 {
		c2s.PositionID = &req.PositionID
	}

	serialNo := c.NextSerialNo()
	c2s.PacketID.SerialNo = &serialNo

	pkt := &trdplaceorder.Request{C2S: c2s}
	var rsp trdplaceorder.Response

	if err := c.RequestContext(ctx, ProtoID_PlaceOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("PlaceOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("PlaceOrder: s2c is nil")
	}

	return &PlaceOrderResponse{
		OrderID:   s2c.GetOrderID(),
		OrderIDEx: s2c.GetOrderIDEx(),
	}, nil
}

// ModifyOrderRequest is the request to modify an existing order (cancel, update price, or update quantity).
type ModifyOrderRequest struct {
	AccID              uint64
	TrdMarket          constant.TrdMarket
	TrdEnv             constant.TrdEnv
	OrderID            uint64
	ModifyOrderOp      constant.ModifyOrderOp
	Price              float64
	Qty                float64
	ForAll             bool
	TrdMarket2         constant.TrdMarket
	AdjustPrice        bool
	AdjustSideAndLimit float64
	AuxPrice           float64
	TrailType          constant.TrailType
	TrailValue         float64
	TrailSpread        float64
	OrderIDEx          string
}

// ModifyOrderResponse is the response returned after modifying an order.
type ModifyOrderResponse struct {
	AccID     uint64
	TrdEnv    int32
	TrdMarket int32
	OrderID   uint64
	OrderIDEx string
}

// ModifyOrder modifies or cancels an existing order.
// Returns the modification response or an error if it fails.
func ModifyOrder(ctx context.Context, c *futuapi.Client, req *ModifyOrderRequest) (*ModifyOrderResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.OrderID == 0 && req.OrderIDEx == "" {
		return nil, fmt.Errorf("order ID or OrderIDEx must be provided")
	}
	if req.ModifyOrderOp <= 0 {
		return nil, fmt.Errorf("invalid modify operation: must be valid order operation type")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)
	trdMarket2 := int32(req.TrdMarket2)
	modifyOrderOp := int32(req.ModifyOrderOp)
	trailType := int32(req.TrailType)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	orderID := req.OrderID
	c2s := &trdmodifyorder.C2S{
		Header:        header,
		OrderID:       &orderID,
		ModifyOrderOp: &modifyOrderOp,
		PacketID: &common.PacketID{
			ConnID: proto.Uint64(c.GetConnID()),
		},
	}
	// Optional fields - only set when provided
	if req.Qty != 0 {
		c2s.Qty = &req.Qty
	}
	if req.Price != 0 {
		c2s.Price = &req.Price
	}
	if req.ForAll {
		c2s.ForAll = &req.ForAll
	}
	if req.TrdMarket2 != 0 {
		c2s.TrdMarket = &trdMarket2
	}
	if req.AdjustPrice {
		c2s.AdjustPrice = &req.AdjustPrice
	}
	if req.AdjustSideAndLimit != 0 {
		c2s.AdjustSideAndLimit = &req.AdjustSideAndLimit
	}
	if req.AuxPrice != 0 {
		c2s.AuxPrice = &req.AuxPrice
	}
	if req.TrailType != 0 {
		c2s.TrailType = &trailType
	}
	if req.TrailValue != 0 {
		c2s.TrailValue = &req.TrailValue
	}
	if req.TrailSpread != 0 {
		c2s.TrailSpread = &req.TrailSpread
	}
	if req.OrderIDEx != "" {
		c2s.OrderIDEx = &req.OrderIDEx
	}

	serialNo := c.NextSerialNo()
	c2s.PacketID.SerialNo = &serialNo

	pkt := &trdmodifyorder.Request{C2S: c2s}
	var rsp trdmodifyorder.Response

	if err := c.RequestContext(ctx, ProtoID_ModifyOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("ModifyOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("ModifyOrder: s2c is nil")
	}

	return &ModifyOrderResponse{
		AccID:     s2c.GetHeader().GetAccID(),
		TrdEnv:    s2c.GetHeader().GetTrdEnv(),
		TrdMarket: s2c.GetHeader().GetTrdMarket(),
		OrderID:   s2c.GetOrderID(),
		OrderIDEx: s2c.GetOrderIDEx(),
	}, nil
}

// UnlockTradeRequest is the request to unlock or lock trading with a password.
type UnlockTradeRequest struct {
	Unlock       bool
	PwdMD5       constant.SensitiveString
	SecurityFirm int32
}

// UnlockTrade unlocks or locks trading functionality using the provided password.
// Returns an error if the unlock operation fails.
// UnlockTrade unlocks or locks trading with the given password.
func UnlockTrade(ctx context.Context, c *futuapi.Client, req *UnlockTradeRequest) error {
	// Input validation
	if req.PwdMD5.IsEmpty() {
		return fmt.Errorf("password MD5 is required")
	}

	pwdRaw := req.PwdMD5.Raw()
	c2s := &trdunlocktrade.C2S{
		Unlock:       &req.Unlock,
		PwdMD5:       &pwdRaw,
		SecurityFirm: &req.SecurityFirm,
	}

	pkt := &trdunlocktrade.Request{C2S: c2s}
	var rsp trdunlocktrade.Response

	if err := c.RequestContext(ctx, ProtoID_UnlockTrade, pkt, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return wrapError("UnlockTrade", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}

// GetOrderFeeRequest is the request to retrieve order fee information.
type GetOrderFeeRequest struct {
	AccID         uint64
	TrdMarket     constant.TrdMarket
	TrdEnv        constant.TrdEnv
	OrderIDExList []string
}

// OrderFeeInfo represents the fee information for a single order.
type OrderFeeInfo struct {
	OrderIDEx string
	FeeAmount float64
	FeeList   []*OrderFeeItemInfo
}

// OrderFeeItemInfo represents a single fee item with its title and value.
type OrderFeeItemInfo struct {
	Title string
	Value float64
}

// GetOrderFeeResponse is the response containing order fee information.
type GetOrderFeeResponse struct {
	OrderFeeList []*OrderFeeInfo
}

// GetOrderFee retrieves the fee details for specified orders.
// Returns the order fee list or an error if the request fails.
func GetOrderFee(ctx context.Context, c *futuapi.Client, req *GetOrderFeeRequest) (*GetOrderFeeResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}
	if len(req.OrderIDExList) == 0 {
		return nil, fmt.Errorf("order ID list is empty")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgetorderfee.C2S{
		Header:        header,
		OrderIdExList: req.OrderIDExList,
	}

	pkt := &trdgetorderfee.Request{C2S: c2s}
	var rsp trdgetorderfee.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderFee, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderFee", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderFee: s2c is nil")
	}

	result := &GetOrderFeeResponse{
		OrderFeeList: make([]*OrderFeeInfo, 0, len(s2c.GetOrderFeeList())),
	}

	for _, f := range s2c.GetOrderFeeList() {
		if f == nil {
			continue
		}
		feeInfo := &OrderFeeInfo{
			OrderIDEx: f.GetOrderIDEx(),
			FeeAmount: f.GetFeeAmount(),
			FeeList:   make([]*OrderFeeItemInfo, 0, len(f.GetFeeList())),
		}
		for _, item := range f.GetFeeList() {
			if item == nil {
				continue
			}
			feeInfo.FeeList = append(feeInfo.FeeList, &OrderFeeItemInfo{
				Title: item.GetTitle(),
				Value: item.GetValue(),
			})
		}
		result.OrderFeeList = append(result.OrderFeeList, feeInfo)
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
	// Input validation
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
}

// GetMaxTrdQtysResponse is the response containing maximum tradable quantities.
type GetMaxTrdQtysResponse struct {
	MaxTrdQtys *MaxTrdQtysInfo
}

// GetMaxTrdQtys retrieves the maximum tradable quantities for a given order.
// Returns the maximum quantities or an error if the request fails.
func GetMaxTrdQtys(ctx context.Context, c *futuapi.Client, req *GetMaxTrdQtysRequest) (*GetMaxTrdQtysResponse, error) {
	// Input validation
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
	}
	// Optional fields - only set when provided
	if req.Price != 0 {
		c2s.Price = &req.Price
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
		},
	}, nil
}

// GetHistoryOrderListRequest is the request to retrieve historical order list.
type GetHistoryOrderListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
	FilterStatusList []int32
}

// GetHistoryOrderListResponse is the response containing historical orders.
type GetHistoryOrderListResponse struct {
	OrderList []*trdcommon.Order
}

// GetHistoryOrderList retrieves the historical order list based on filter conditions.
// Returns the historical order list or an error if the request fails.
func GetHistoryOrderList(ctx context.Context, c *futuapi.Client, req *GetHistoryOrderListRequest) (*GetHistoryOrderListResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgethistoryorderlist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
		FilterStatusList: req.FilterStatusList,
	}
	if c2s.FilterConditions == nil {
		c2s.FilterConditions = &trdcommon.TrdFilterConditions{}
	}

	pkt := &trdgethistoryorderlist.Request{C2S: c2s}
	var rsp trdgethistoryorderlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetHistoryOrderList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryOrderList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHistoryOrderList: s2c is nil")
	}

	return &GetHistoryOrderListResponse{
		OrderList: s2c.GetOrderList(),
	}, nil
}

// GetHistoryOrderFillListRequest is the request to retrieve historical order fill list.
type GetHistoryOrderFillListRequest struct {
	AccID            uint64
	TrdMarket        constant.TrdMarket
	TrdEnv           constant.TrdEnv
	FilterConditions *trdcommon.TrdFilterConditions
}

// GetHistoryOrderFillListResponse is the response containing historical order fills.
type GetHistoryOrderFillListResponse struct {
	OrderFillList []*OrderFill
}

// GetHistoryOrderFillList retrieves the historical order fill (execution) list based on filter conditions.
// Returns the historical order fill list or an error if the request fails.
func GetHistoryOrderFillList(ctx context.Context, c *futuapi.Client, req *GetHistoryOrderFillListRequest) (*GetHistoryOrderFillListResponse, error) {
	// Input validation
	if req.AccID == 0 {
		return nil, fmt.Errorf("invalid account ID: must be non-zero")
	}

	trdEnv := int32(req.TrdEnv)
	trdMarket := int32(req.TrdMarket)

	header := &trdcommon.TrdHeader{
		TrdEnv:    &trdEnv,
		AccID:     &req.AccID,
		TrdMarket: &trdMarket,
	}

	c2s := &trdgethistoryorderfilllist.C2S{
		Header:           header,
		FilterConditions: req.FilterConditions,
	}
	if c2s.FilterConditions == nil {
		c2s.FilterConditions = &trdcommon.TrdFilterConditions{}
	}
	if c2s.FilterConditions.GetBeginTime() == "" {
		begin := time.Now().AddDate(0, 0, -30).Format("2006-01-02 15:04:05")
		c2s.FilterConditions.BeginTime = &begin
	}
	if c2s.FilterConditions.GetEndTime() == "" {
		end := time.Now().Format("2006-01-02 15:04:05")
		c2s.FilterConditions.EndTime = &end
	}

	pkt := &trdgethistoryorderfilllist.Request{C2S: c2s}
	var rsp trdgethistoryorderfilllist.Response

	if err := c.RequestContext(ctx, ProtoID_GetHistoryOrderFillList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryOrderFillList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHistoryOrderFillList: s2c is nil")
	}

	list := make([]*OrderFill, 0, len(s2c.GetOrderFillList()))
	for _, f := range s2c.GetOrderFillList() {
		if f == nil {
			continue
		}
		list = append(list, &OrderFill{
			FillID:            f.GetFillID(),
			FillIDEx:          f.GetFillIDEx(),
			OrderID:           f.GetOrderID(),
			OrderIDEx:         f.GetOrderIDEx(),
			Code:              f.GetCode(),
			Name:              f.GetName(),
			TrdSide:           f.GetTrdSide(),
			Price:             f.GetPrice(),
			Qty:               f.GetQty(),
			CreateTime:        f.GetCreateTime(),
			CounterBrokerID:   f.GetCounterBrokerID(),
			CounterBrokerName: f.GetCounterBrokerName(),
			SecMarket:         f.GetSecMarket(),
			CreateTimestamp:   f.GetCreateTimestamp(),
			UpdateTimestamp:   f.GetUpdateTimestamp(),
			Status:            f.GetStatus(),
			TrdMarket:         f.GetTrdMarket(),
			JpAccType:         f.GetJpAccType(),
		})
	}

	return &GetHistoryOrderFillListResponse{
		OrderFillList: list,
	}, nil
}

// SubAccPushRequest is the request to subscribe to account push notifications.
type SubAccPushRequest struct {
	AccIDList []uint64
}

// SubAccPush subscribes to account push notifications for the specified account IDs.
// Returns an error if the subscription fails.
func SubAccPush(ctx context.Context, c *futuapi.Client, req *SubAccPushRequest) error {
	// Input validation
	if len(req.AccIDList) == 0 {
		return fmt.Errorf("account ID list is empty")
	}

	c2s := &trdsubaccpush.C2S{
		AccIDList: req.AccIDList,
	}

	pkt := &trdsubaccpush.Request{C2S: c2s}
	var rsp trdsubaccpush.Response

	if err := c.RequestContext(ctx, ProtoID_SubAccPush, pkt, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return wrapError("SubAccPush", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}

// ReconfirmOrderRequest is the request to reconfirm an order with a specified reason.
type ReconfirmOrderRequest struct {
	PacketID        *common.PacketID
	Header          *trdcommon.TrdHeader
	OrderID         uint64
	ReconfirmReason int32
}

// ReconfirmOrderResponse is the response containing the reconfirmed order header and ID.
type ReconfirmOrderResponse struct {
	Header  *trdcommon.TrdHeader
	OrderID uint64
}

// ReconfirmOrder reconfirms an order that requires additional verification.
// Returns the reconfirmed order details or an error if the request fails.
func ReconfirmOrder(ctx context.Context, c *futuapi.Client, req *ReconfirmOrderRequest) (*ReconfirmOrderResponse, error) {
	// Input validation
	if req.OrderID == 0 {
		return nil, fmt.Errorf("invalid order ID: must be non-zero")
	}
	if req.Header == nil {
		return nil, fmt.Errorf("header is required")
	}

	c2s := &trdreconfirmorder.C2S{
		PacketID:        req.PacketID,
		Header:          req.Header,
		OrderID:         &req.OrderID,
		ReconfirmReason: &req.ReconfirmReason,
	}

	pkt := &trdreconfirmorder.Request{C2S: c2s}
	var rsp trdreconfirmorder.Response

	if err := c.RequestContext(ctx, ProtoID_ReconfirmOrder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("ReconfirmOrder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("ReconfirmOrder: s2c is nil")
	}

	return &ReconfirmOrderResponse{
		Header:  s2c.GetHeader(),
		OrderID: s2c.GetOrderID(),
	}, nil
}

// GetFlowSummaryRequest is the request to retrieve fund flow summary for a clearing date.
type GetFlowSummaryRequest struct {
	Header            *trdcommon.TrdHeader
	ClearingDate      string
	CashFlowDirection int32
}

// GetFlowSummaryResponse is the response containing the fund flow summary.
type GetFlowSummaryResponse struct {
	Header          *trdcommon.TrdHeader
	FlowSummaryList []*trdflowsummary.FlowSummaryInfo
}

// GetFlowSummary retrieves the fund flow summary for a specified clearing date.
// Returns the flow summary list or an error if the request fails.
func GetFlowSummary(ctx context.Context, c *futuapi.Client, req *GetFlowSummaryRequest) (*GetFlowSummaryResponse, error) {
	// Input validation
	if req.Header == nil {
		return nil, fmt.Errorf("header is required")
	}

	c2s := &trdflowsummary.C2S{
		Header:            req.Header,
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
		Header:          s2c.GetHeader(),
		FlowSummaryList: s2c.GetFlowSummaryInfoList(),
	}, nil
}
