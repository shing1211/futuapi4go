// Package qot provides market data APIs for the Futu OpenD SDK.
//
// This package covers real-time quotes, K-lines, order book, tick data,
// broker queue, capital flow, stock screening, options, warrants, and
// historical data requests. All functions require a connected client.
//
// For Python SDK migration, use the constant package for Python-style constants:
//
//	import (
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/qot"
//	    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
//	)
//
//	// Create security with Python-style constants
//	market := constant.Market_HK
//	code := "00700"
//	securities := []*qotcommon.Security{
//	    {Market: &market, Code: &code},
//	}
//
//	// Use constant package for K-line types, rehab types, etc.
//	// KLType: constant.KLType_K_Day, constant.KLType_K_1Min, etc.
//	// RehabType: constant.RehabType_Forward (QFQ), constant.RehabType_Backward (HFQ)
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

// quotes, err := qot.GetBasicQot(cli, securities)
package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategy"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyanalysis"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyspreads"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// wrapError standardizes error messages for proto response failures
func wrapError(funcName string, retType int32, retMsg string) error {
	code := constant.ErrorCode(retType)
	return constant.NewFutuError(code, funcName, retMsg)
}

const (
	ProtoID_GetBasicQot             = 3004
	ProtoID_GetKL                   = 3006
	ProtoID_GetOrderBook            = 3012
	ProtoID_GetTicker               = 3010
	ProtoID_GetRT                   = 3008
	ProtoID_GetSecuritySnapshot     = 3203
	ProtoID_GetBroker               = 3014
	ProtoID_GetStaticInfo           = 3202
	ProtoID_GetPlateSet             = 3204
	ProtoID_GetPlateSecurity        = 3205
	ProtoID_GetSuspend              = 3201
	ProtoID_GetCodeChange           = 3216
	ProtoID_GetFutureInfo           = 3218
	ProtoID_GetIpoList              = 3217
	ProtoID_GetHistoryKL           = 3101
	ProtoID_Qot_GetRehab              = 3102
	ProtoID_GetHoldingChangeList    = 3208
	ProtoID_RequestRehab            = 3105
	ProtoID_GetUserSecurityGroup    = 3222
	ProtoID_ModifyUserSecurity      = 3214
	ProtoID_SetPriceReminder        = 3220
	ProtoID_GetCapitalFlow          = 3211
	ProtoID_GetCapitalDistribution  = 3212
	ProtoID_StockFilter             = 3215
	ProtoID_GetOptionChain          = 3209
	ProtoID_GetOptionExpirationDate = 3224
	ProtoID_GetWarrant              = 3210
	ProtoID_GetUserSecurity         = 3213
	ProtoID_GetPriceReminder        = 3221
	ProtoID_RequestTradeDate        = 3219
	ProtoID_Subscribe               = 3001
	ProtoID_RegQotPush              = 3002
	ProtoID_RequestHistoryKL        = 3103
	ProtoID_RequestHistoryKLQuota   = 3104

	// Screen APIs (v10.6+)
	ProtoID_StockScreen                = 3252
	ProtoID_OptionScreen               = 3253
	ProtoID_WarrantScreen              = 3254
	ProtoID_GetOptionQuote             = 3255
	ProtoID_GetOptionStrategy          = 3256
	ProtoID_GetOptionStrategyAnalysis  = 3257
	ProtoID_GetOptionStrategySpread    = 3258
)

// BasicQot represents basic quote data for a security.
type BasicQot struct {
	Security       *qotcommon.Security
	Name           string
	IsSuspended    bool
	UpdateTime     string
	HighPrice      float64
	OpenPrice      float64
	LowPrice       float64
	CurPrice       float64
	LastClosePrice float64
	Volume         int64
	Turnover       float64
	TurnoverRate   float64
	Amplitude      float64
	ListTime         string
	PriceSpread      float64
	DarkStatus       int32
	ListTimestamp    float64
	UpdateTimestamp  float64
	SecStatus       int32
	OptionExData    *qotcommon.OptionBasicQotExData
	PreMarket       *qotcommon.PreAfterMarketData
	AfterMarket     *qotcommon.PreAfterMarketData
	FutureExData    *qotcommon.FutureBasicQotExData
	Overnight       *qotcommon.PreAfterMarketData
	WarrantExData   *qotcommon.WarrantBasicQotExData
}

// GetBasicQot returns basic quote data for the given securities.
func GetBasicQot(ctx context.Context, c *futuapi.Client, securityList []*qotcommon.Security) ([]*BasicQot, error) {
	if len(securityList) == 0 {
		return nil, fmt.Errorf("GetBasicQot: security list is empty")
	}
	c2s := &qotgetbasicqot.C2S{
		SecurityList: securityList,
	}
	req := &qotgetbasicqot.Request{C2S: c2s}
	var rsp qotgetbasicqot.Response

	if err := c.RequestContext(ctx, ProtoID_GetBasicQot, req, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetBasicQot", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetBasicQot", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := make([]*BasicQot, 0, len(s2c.BasicQotList))
	for _, bq := range s2c.BasicQotList {
		if bq == nil {
			continue
		}
		result = append(result, &BasicQot{
			Security:       bq.Security,
			Name:           util.ProtoStr(bq.Name),
			IsSuspended:    util.ProtoBool(bq.IsSuspended),
			UpdateTime:     util.ProtoStr(bq.UpdateTime),
			HighPrice:      util.ProtoFloat64(bq.HighPrice),
			OpenPrice:      util.ProtoFloat64(bq.OpenPrice),
			LowPrice:       util.ProtoFloat64(bq.LowPrice),
			CurPrice:       util.ProtoFloat64(bq.CurPrice),
			LastClosePrice: util.ProtoFloat64(bq.LastClosePrice),
			Volume:         util.ProtoInt64(bq.Volume),
			Turnover:       util.ProtoFloat64(bq.Turnover),
			TurnoverRate:   util.ProtoFloat64(bq.TurnoverRate),
			Amplitude:      util.ProtoFloat64(bq.Amplitude),
			ListTime:        util.ProtoStr(bq.ListTime),
			PriceSpread:     util.ProtoFloat64(bq.PriceSpread),
			DarkStatus:      util.ProtoInt32(bq.DarkStatus),
			ListTimestamp:   util.ProtoFloat64(bq.ListTimestamp),
			UpdateTimestamp: util.ProtoFloat64(bq.UpdateTimestamp),
			SecStatus:       util.ProtoInt32(bq.SecStatus),
			OptionExData:   bq.OptionExData,
			PreMarket:       bq.PreMarket,
			AfterMarket:     bq.AfterMarket,
			FutureExData:   bq.FutureExData,
			Overnight:       bq.Overnight,
			WarrantExData:  bq.WarrantExData,
		})
	}

	return result, nil
}

// KLine represents a single K-line (candlestick) data point.
type KLine struct {
	Time           string
	IsBlank        bool
	HighPrice      float64
	OpenPrice      float64
	LowPrice       float64
	ClosePrice     float64
	LastClosePrice float64
	Volume         int64
	Turnover       float64
	TurnoverRate   float64
	Pe             float64
	ChangeRate     float64
	Timestamp      float64
}

// GetOptionQuoteRequest defines parameters for GetOptionQuote.
type GetOptionQuoteRequest struct {
	MultiLegs []*qotcommon.ComboLeg
}

// GetOptionQuoteResponse is the response type for GetOptionQuote.
type GetOptionQuoteResponse struct {
	OptionQuoteList []*qotgetoptionquote.OptionQuote
}

// GetOptionQuote returns real-time quotes for option combo legs.
func GetOptionQuote(ctx context.Context, c *futuapi.Client, req *GetOptionQuoteRequest) (*GetOptionQuoteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionQuote: request is nil")
	}
	if len(req.MultiLegs) == 0 {
		return nil, fmt.Errorf("GetOptionQuote: MultiLegs is empty")
	}
	c2s := &qotgetoptionquote.C2S{
		MultiLegs: req.MultiLegs,
	}
	pkt := &qotgetoptionquote.Request{C2S: c2s}
	var rsp qotgetoptionquote.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionQuote, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionQuote", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionQuote", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetOptionQuoteResponse{
		OptionQuoteList: s2c.OptionQuoteList,
	}, nil
}

// GetOptionStrategyRequest defines parameters for GetOptionStrategy.
type GetOptionStrategyRequest struct {
	Owner           *qotcommon.Security
	OptionStrategy  int32
	ExpireTime      string
	FarExpireTime   string
	Spread          float64
	OptionType      int32
	StrikePrice     float64
	IndexOptionType int32
}

// GetOptionStrategyResponse is the response type for GetOptionStrategy.
type GetOptionStrategyResponse struct {
	StrategyList []*qotgetoptionstrategy.OptionStrategyItem
}

// GetOptionStrategy returns option strategy combo lists for the given underlying.
func GetOptionStrategy(ctx context.Context, c *futuapi.Client, req *GetOptionStrategyRequest) (*GetOptionStrategyResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionStrategy: request is nil")
	}
	if req.Owner == nil {
		return nil, fmt.Errorf("GetOptionStrategy: Owner is nil")
	}
	c2s := &qotgetoptionstrategy.C2S{
		Owner:          req.Owner,
		OptionStrategy: &req.OptionStrategy,
	}
	if req.ExpireTime != "" {
		c2s.ExpireTime = &req.ExpireTime
	}
	if req.FarExpireTime != "" {
		c2s.FarExpireTime = &req.FarExpireTime
	}
	if req.Spread != 0 {
		c2s.Spread = &req.Spread
	}
	if req.OptionType != 0 {
		c2s.OptionType = &req.OptionType
	}
	if req.StrikePrice != 0 {
		c2s.StrikePrice = &req.StrikePrice
	}
	if req.IndexOptionType != 0 {
		c2s.IndexOptionType = &req.IndexOptionType
	}
	pkt := &qotgetoptionstrategy.Request{C2S: c2s}
	var rsp qotgetoptionstrategy.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionStrategy, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionStrategy", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionStrategy", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetOptionStrategyResponse{
		StrategyList: s2c.StrategyList,
	}, nil
}

// GetOptionStrategyAnalysisRequest defines parameters for GetOptionStrategyAnalysis.
type GetOptionStrategyAnalysisRequest struct {
	MultiLegs []*qotcommon.ComboLeg
}

// GetOptionStrategyAnalysisResponse is the response type for GetOptionStrategyAnalysis.
type GetOptionStrategyAnalysisResponse struct {
	Code            string
	Name            string
	OptionStrategy  int32
	Bid1            float64
	Ask1            float64
	MaxProfit       float64
	MaxLoss         float64
	BreakevenPoints []float64
	ProbOfProfit    float64
	Delta           float64
	Theta           float64
}

// GetOptionStrategyAnalysis returns P&L analysis for an option strategy combination.
func GetOptionStrategyAnalysis(ctx context.Context, c *futuapi.Client, req *GetOptionStrategyAnalysisRequest) (*GetOptionStrategyAnalysisResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionStrategyAnalysis: request is nil")
	}
	if len(req.MultiLegs) == 0 {
		return nil, fmt.Errorf("GetOptionStrategyAnalysis: MultiLegs is empty")
	}
	c2s := &qotgetoptionstrategyanalysis.C2S{
		MultiLegs: req.MultiLegs,
	}
	pkt := &qotgetoptionstrategyanalysis.Request{C2S: c2s}
	var rsp qotgetoptionstrategyanalysis.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionStrategyAnalysis, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionStrategyAnalysis", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionStrategyAnalysis", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetOptionStrategyAnalysisResponse{
		Code:            util.ProtoStr(s2c.Code),
		Name:            util.ProtoStr(s2c.Name),
		OptionStrategy:  util.ProtoInt32(s2c.OptionStrategy),
		Bid1:            util.ProtoFloat64(s2c.Bid1),
		Ask1:            util.ProtoFloat64(s2c.Ask1),
		MaxProfit:       util.ProtoFloat64(s2c.MaxProfit),
		MaxLoss:         util.ProtoFloat64(s2c.MaxLoss),
		BreakevenPoints: s2c.BreakevenPoints,
		ProbOfProfit:    util.ProtoFloat64(s2c.ProbOfProfit),
		Delta:           util.ProtoFloat64(s2c.Delta),
		Theta:           util.ProtoFloat64(s2c.Theta),
	}, nil
}

// GetOptionStrategySpreadRequest defines parameters for GetOptionStrategySpread.
type GetOptionStrategySpreadRequest struct {
	Owner           *qotcommon.Security
	OptionStrategy  int32
	ExpireTime      string
	FarExpireTime   string
	IndexOptionType int32
}

// GetOptionStrategySpreadResponse is the response type for GetOptionStrategySpread.
type GetOptionStrategySpreadResponse struct {
	SpreadList []float64
}

// GetOptionStrategySpread returns available spread values for an option strategy.
func GetOptionStrategySpread(ctx context.Context, c *futuapi.Client, req *GetOptionStrategySpreadRequest) (*GetOptionStrategySpreadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionStrategySpread: request is nil")
	}
	if req.Owner == nil {
		return nil, fmt.Errorf("GetOptionStrategySpread: Owner is nil")
	}
	c2s := &qotgetoptionstrategyspreads.C2S{
		Owner:          req.Owner,
		OptionStrategy: &req.OptionStrategy,
	}
	if req.ExpireTime != "" {
		c2s.ExpireTime = &req.ExpireTime
	}
	if req.FarExpireTime != "" {
		c2s.FarExpireTime = &req.FarExpireTime
	}
	if req.IndexOptionType != 0 {
		c2s.IndexOptionType = &req.IndexOptionType
	}
	pkt := &qotgetoptionstrategyspreads.Request{C2S: c2s}
	var rsp qotgetoptionstrategyspreads.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionStrategySpread, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionStrategySpread", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionStrategySpread", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetOptionStrategySpreadResponse{
		SpreadList: s2c.SpreadList,
	}, nil
}


// GetKLRequest defines parameters for GetKL.
type GetKLRequest struct {
	Security  *qotcommon.Security
	RehabType int32
	KLType    int32
	ReqNum    int32
}

// GetKLResponse is the response type for GetKL.
type GetKLResponse struct {
	Security *qotcommon.Security
	Name     string
	KLList   []*KLine
}

// GetKL returns K-line (candlestick) data for the given security.
func GetKL(ctx context.Context, c *futuapi.Client, req *GetKLRequest) (*GetKLResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetKL: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetKL: Security is nil")
	}
	if req.ReqNum <= 0 {
		return nil, fmt.Errorf("GetKL: ReqNum must be positive")
	}
	c2s := &qotgetkl.C2S{
		Security:  req.Security,
		RehabType: &req.RehabType,
		KlType:    &req.KLType,
		ReqNum:    &req.ReqNum,
	}
	pkt := &qotgetkl.Request{C2S: c2s}
	var rsp qotgetkl.Response

	if err := c.RequestContext(ctx, ProtoID_GetKL, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetKL", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetKL", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetKLResponse{
		Security: s2c.Security,
		Name:     util.ProtoStr(s2c.Name),
		KLList:   make([]*KLine, 0, len(s2c.KlList)),
	}

	for _, kl := range s2c.KlList {
		if kl == nil {
			continue
		}
		result.KLList = append(result.KLList, &KLine{
			Time:           util.ProtoStr(kl.Time),
			IsBlank:        util.ProtoBool(kl.IsBlank),
			HighPrice:      util.ProtoFloat64(kl.HighPrice),
			OpenPrice:      util.ProtoFloat64(kl.OpenPrice),
			LowPrice:       util.ProtoFloat64(kl.LowPrice),
			ClosePrice:     util.ProtoFloat64(kl.ClosePrice),
			LastClosePrice: util.ProtoFloat64(kl.LastClosePrice),
			Volume:         util.ProtoInt64(kl.Volume),
			Turnover:       util.ProtoFloat64(kl.Turnover),
			TurnoverRate:   util.ProtoFloat64(kl.TurnoverRate),
			Pe:             util.ProtoFloat64(kl.Pe),
			ChangeRate:     util.ProtoFloat64(kl.ChangeRate),
			Timestamp:      util.ProtoFloat64(kl.Timestamp),
		})
	}

	return result, nil
}
