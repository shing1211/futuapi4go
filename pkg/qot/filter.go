// Package qot provides market data APIs for the Futu OpenD SDK.
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

package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetreference"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetwarrant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetReference = 3206
)

// StockFilterRequest defines parameters for StockFilter.
type StockFilterRequest struct {
	Begin                     int32
	Num                       int32
	Market                    int32
	Plate                     *qotcommon.Security
	BaseFilterList            []*qotstockfilter.BaseFilter
	AccumulateFilterList      []*qotstockfilter.AccumulateFilter
	FinancialFilterList       []*qotstockfilter.FinancialFilter
	PatternFilterList         []*qotstockfilter.PatternFilter
	CustomIndicatorFilterList []*qotstockfilter.CustomIndicatorFilter
}

// StockFilterData represents a single stock filter result.
type StockFilterData struct {
	Security                *qotcommon.Security
	Name                    string
	BaseDataList            []*qotstockfilter.BaseData
	AccumulateDataList      []*qotstockfilter.AccumulateData
	FinancialDataList       []*qotstockfilter.FinancialData
	CustomIndicatorDataList []*qotstockfilter.CustomIndicatorData
}

// StockFilterResponse is the response type for StockFilter.
type StockFilterResponse struct {
	LastPage bool
	AllCount int32
	DataList []*StockFilterData
}

// StockFilter filters stocks based on various criteria (选股).
func StockFilter(ctx context.Context, c *futuapi.Client, req *StockFilterRequest) (*StockFilterResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("StockFilter: request is nil")
	}
	if req.Market == 0 {
		return nil, fmt.Errorf("invalid market: must be non-zero")
	}
	if req.Num <= 0 || req.Num > 100 {
		return nil, fmt.Errorf("invalid num: must be between 1 and 100")
	}

	c2s := &qotstockfilter.C2S{
		Begin:                     &req.Begin,
		Num:                       &req.Num,
		Market:                    &req.Market,
		Plate:                     req.Plate,
		BaseFilterList:            req.BaseFilterList,
		AccumulateFilterList:      req.AccumulateFilterList,
		FinancialFilterList:       req.FinancialFilterList,
		PatternFilterList:         req.PatternFilterList,
		CustomIndicatorFilterList: req.CustomIndicatorFilterList,
	}

	pkt := &qotstockfilter.Request{C2S: c2s}
	var rsp qotstockfilter.Response

	if err := c.RequestContext(ctx, ProtoID_StockFilter, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("StockFilter", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("StockFilter", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &StockFilterResponse{
		LastPage: util.ProtoBool(s2c.LastPage),
		AllCount: util.ProtoInt32(s2c.AllCount),
		DataList: make([]*StockFilterData, 0, len(s2c.DataList)),
	}

	for _, d := range s2c.DataList {
		if d == nil {
			continue
		}
		result.DataList = append(result.DataList, &StockFilterData{
			Security:                d.Security,
			Name:                    util.ProtoStr(d.Name),
			BaseDataList:            d.BaseDataList,
			AccumulateDataList:      d.AccumulateDataList,
			FinancialDataList:       d.FinancialDataList,
			CustomIndicatorDataList: d.CustomIndicatorDataList,
		})
	}

	return result, nil
}

// GetWarrantRequest defines parameters for GetWarrant.
type GetWarrantRequest struct {
	Begin                 int32
	Num                   int32
	SortField             int32
	Ascend                bool
	Owner                 *qotcommon.Security
	TypeList              []int32
	IssuerList            []int32
	MaturityTimeMin       string
	MaturityTimeMax       string
	IpoPeriod             int32
	PriceType             int32
	Status                int32
	CurPriceMin           float64
	CurPriceMax           float64
	StrikePriceMin        float64
	StrikePriceMax        float64
	StreetMin             float64
	StreetMax             float64
	ConversionMin         float64
	ConversionMax         float64
	VolMin                uint64
	VolMax                uint64
	PremiumMin            float64
	PremiumMax            float64
	LeverageRatioMin      float64
	LeverageRatioMax      float64
	DeltaMin              float64
	DeltaMax              float64
	ImpliedMin            float64
	ImpliedMax            float64
	RecoveryPriceMin      float64
	RecoveryPriceMax      float64
	PriceRecoveryRatioMin float64
	PriceRecoveryRatioMax float64
}

// WarrantData represents detailed data for a single warrant.
type WarrantData struct {
	Stock              *qotcommon.Security
	Owner              *qotcommon.Security
	Type               int32
	Issuer             int32
	MaturityTime       string
	MaturityTimestamp  float64
	ListTime           string
	ListTimestamp      float64
	LastTradeTime      string
	LastTradeTimestamp float64
	RecoveryPrice      float64
	ConversionRatio    float64
	LotSize            int32
	StrikePrice        float64
	LastClosePrice     float64
	Name               string
	CurPrice           float64
	PriceChangeVal     float64
	ChangeRate         float64
	Status             int32
	BidPrice           float64
	AskPrice           float64
	BidVol             int64
	AskVol             int64
	Volume             int64
	Turnover           float64
	Score              float64
	Premium            float64
	BreakEvenPoint     float64
	Leverage           float64
	Ipop               float64
	PriceRecoveryRatio float64
	ConversionPrice    float64
	StreetRate         float64
	StreetVol          int64
	Amplitude          float64
	IssueSize          int64
	HighPrice          float64
	LowPrice           float64
	ImpliedVolatility  float64
	Delta              float64
	EffectiveLeverage  float64
	UpperStrikePrice   float64
	LowerStrikePrice   float64
	InLinePriceStatus  int32
}

// GetWarrantResponse is the response type for GetWarrant.
type GetWarrantResponse struct {
	LastPage        bool
	AllCount        int32
	WarrantDataList []*WarrantData
}

// GetWarrant returns warrant data filtered by the given criteria.
func GetWarrant(ctx context.Context, c *futuapi.Client, req *GetWarrantRequest) (*GetWarrantResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetWarrant: request is nil")
	}
	if req.Num <= 0 || req.Num > 100 {
		return nil, fmt.Errorf("invalid num: must be between 1 and 100")
	}

	c2s := &qotgetwarrant.C2S{
		Begin:     &req.Begin,
		Num:       &req.Num,
		SortField: &req.SortField,
		Ascend:    &req.Ascend,
		Owner:     req.Owner,
		TypeList:  req.TypeList,
	}
	if req.MaturityTimeMin != "" {
		c2s.MaturityTimeMin = &req.MaturityTimeMin
	}
	if req.MaturityTimeMax != "" {
		c2s.MaturityTimeMax = &req.MaturityTimeMax
	}
	if req.IpoPeriod != 0 {
		c2s.IpoPeriod = &req.IpoPeriod
	}
	if req.PriceType != 0 {
		c2s.PriceType = &req.PriceType
	}
	if req.Status != 0 {
		c2s.Status = &req.Status
	}
	if req.CurPriceMin != 0 {
		c2s.CurPriceMin = &req.CurPriceMin
	}
	if req.CurPriceMax != 0 {
		c2s.CurPriceMax = &req.CurPriceMax
	}
	if req.StrikePriceMin != 0 {
		c2s.StrikePriceMin = &req.StrikePriceMin
	}
	if req.StrikePriceMax != 0 {
		c2s.StrikePriceMax = &req.StrikePriceMax
	}
	if req.StreetMin != 0 {
		c2s.StreetMin = &req.StreetMin
	}
	if req.StreetMax != 0 {
		c2s.StreetMax = &req.StreetMax
	}
	if req.ConversionMin != 0 {
		c2s.ConversionMin = &req.ConversionMin
	}
	if req.ConversionMax != 0 {
		c2s.ConversionMax = &req.ConversionMax
	}
	if req.VolMin != 0 {
		c2s.VolMin = &req.VolMin
	}
	if req.VolMax != 0 {
		c2s.VolMax = &req.VolMax
	}
	if req.PremiumMin != 0 {
		c2s.PremiumMin = &req.PremiumMin
	}
	if req.PremiumMax != 0 {
		c2s.PremiumMax = &req.PremiumMax
	}
	if req.LeverageRatioMin != 0 {
		c2s.LeverageRatioMin = &req.LeverageRatioMin
	}
	if req.LeverageRatioMax != 0 {
		c2s.LeverageRatioMax = &req.LeverageRatioMax
	}
	if req.DeltaMin != 0 {
		c2s.DeltaMin = &req.DeltaMin
	}
	if req.DeltaMax != 0 {
		c2s.DeltaMax = &req.DeltaMax
	}
	if req.ImpliedMin != 0 {
		c2s.ImpliedMin = &req.ImpliedMin
	}
	if req.ImpliedMax != 0 {
		c2s.ImpliedMax = &req.ImpliedMax
	}
	if req.RecoveryPriceMin != 0 {
		c2s.RecoveryPriceMin = &req.RecoveryPriceMin
	}
	if req.RecoveryPriceMax != 0 {
		c2s.RecoveryPriceMax = &req.RecoveryPriceMax
	}
	if req.PriceRecoveryRatioMin != 0 {
		c2s.PriceRecoveryRatioMin = &req.PriceRecoveryRatioMin
	}
	if req.PriceRecoveryRatioMax != 0 {
		c2s.PriceRecoveryRatioMax = &req.PriceRecoveryRatioMax
	}

	pkt := &qotgetwarrant.Request{C2S: c2s}
	var rsp qotgetwarrant.Response

	if err := c.RequestContext(ctx, ProtoID_GetWarrant, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetWarrant", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetWarrant", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetWarrantResponse{
		LastPage:        util.ProtoBool(s2c.LastPage),
		AllCount:        util.ProtoInt32(s2c.AllCount),
		WarrantDataList: make([]*WarrantData, 0, len(s2c.WarrantDataList)),
	}

	for _, w := range s2c.WarrantDataList {
		if w == nil {
			continue
		}
		result.WarrantDataList = append(result.WarrantDataList, &WarrantData{
			Stock:              w.Stock,
			Owner:              w.Owner,
			Type:               util.ProtoInt32(w.Type),
			Issuer:             util.ProtoInt32(w.Issuer),
			MaturityTime:       util.ProtoStr(w.MaturityTime),
			MaturityTimestamp:  util.ProtoFloat64(w.MaturityTimestamp),
			ListTime:           util.ProtoStr(w.ListTime),
			ListTimestamp:      util.ProtoFloat64(w.ListTimestamp),
			LastTradeTime:      util.ProtoStr(w.LastTradeTime),
			LastTradeTimestamp: util.ProtoFloat64(w.LastTradeTimestamp),
			RecoveryPrice:      util.ProtoFloat64(w.RecoveryPrice),
			ConversionRatio:    util.ProtoFloat64(w.ConversionRatio),
			LotSize:            util.ProtoInt32(w.LotSize),
			StrikePrice:        util.ProtoFloat64(w.StrikePrice),
			LastClosePrice:     util.ProtoFloat64(w.LastClosePrice),
			Name:               util.ProtoStr(w.Name),
			CurPrice:           util.ProtoFloat64(w.CurPrice),
			PriceChangeVal:     util.ProtoFloat64(w.PriceChangeVal),
			ChangeRate:         util.ProtoFloat64(w.ChangeRate),
			Status:             util.ProtoInt32(w.Status),
			BidPrice:           util.ProtoFloat64(w.BidPrice),
			AskPrice:           util.ProtoFloat64(w.AskPrice),
			BidVol:             util.ProtoInt64(w.BidVol),
			AskVol:             util.ProtoInt64(w.AskVol),
			Volume:             util.ProtoInt64(w.Volume),
			Turnover:           util.ProtoFloat64(w.Turnover),
			Score:              util.ProtoFloat64(w.Score),
			Premium:            util.ProtoFloat64(w.Premium),
			BreakEvenPoint:     util.ProtoFloat64(w.BreakEvenPoint),
			Leverage:           util.ProtoFloat64(w.Leverage),
			Ipop:               util.ProtoFloat64(w.Ipop),
			PriceRecoveryRatio: util.ProtoFloat64(w.PriceRecoveryRatio),
			ConversionPrice:    util.ProtoFloat64(w.ConversionPrice),
			StreetRate:         util.ProtoFloat64(w.StreetRate),
			StreetVol:          util.ProtoInt64(w.StreetVol),
			Amplitude:          util.ProtoFloat64(w.Amplitude),
			IssueSize:          util.ProtoInt64(w.IssueSize),
			HighPrice:          util.ProtoFloat64(w.HighPrice),
			LowPrice:           util.ProtoFloat64(w.LowPrice),
			ImpliedVolatility:  util.ProtoFloat64(w.ImpliedVolatility),
			Delta:              util.ProtoFloat64(w.Delta),
			EffectiveLeverage:  util.ProtoFloat64(w.EffectiveLeverage),
			UpperStrikePrice:   util.ProtoFloat64(w.UpperStrikePrice),
			LowerStrikePrice:   util.ProtoFloat64(w.LowerStrikePrice),
			InLinePriceStatus:  util.ProtoInt32(w.InLinePriceStatus),
		})
	}

	return result, nil
}

// GetReferenceRequest defines parameters for GetReference.
type GetReferenceRequest struct {
	Security      *qotcommon.Security
	ReferenceType int32
}

// GetReferenceResponse is the response type for GetReference.
type GetReferenceResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetReference returns related securities (e.g., futures underlying) for the given security.
func GetReference(ctx context.Context, c *futuapi.Client, req *GetReferenceRequest) (*GetReferenceResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetReference: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetreference.C2S{
		Security:      req.Security,
		ReferenceType: &req.ReferenceType,
	}

	pkt := &qotgetreference.Request{C2S: c2s}
	var rsp qotgetreference.Response

	if err := c.RequestContext(ctx, ProtoID_GetReference, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetReference", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetReference", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetReferenceResponse{
		StaticInfoList: s2c.StaticInfoList,
	}, nil
}
