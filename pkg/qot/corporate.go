// Package qot provides market data APIs for the Futu OpenD SDK.
//
// This package covers real-time quotes, K-lines, order book, tick data,
// broker queue, capital flow, stock screening, options, warrants, and
// historical data requests. All functions require a connected client.
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

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsdividends"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsbuybacks"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsstocksplits"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetCorporateActionsDividends  = 3234
	ProtoID_GetCorporateActionsBuybacks   = 3235
	ProtoID_GetCorporateActionsStockSplits = 3236
)

type GetCorporateActionsDividendsRequest struct {
	Security *qotcommon.Security
}

type GetCorporateActionsDividendsResponse struct {
	DividendList []*qotgetcorporateactionsdividends.DividendItem
}

func GetCorporateActionsDividends(ctx context.Context, c *futuapi.Client, req *GetCorporateActionsDividendsRequest) (*GetCorporateActionsDividendsResponse, error) {
	if req == nil {
		return nil, wrapError("GetCorporateActionsDividends", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetCorporateActionsDividends", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetcorporateactionsdividends.C2S{
		Security: req.Security,
	}
	pkt := &qotgetcorporateactionsdividends.Request{C2S: c2s}
	var rsp qotgetcorporateactionsdividends.Response

	if err := c.RequestContext(ctx, ProtoID_GetCorporateActionsDividends, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCorporateActionsDividends", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCorporateActionsDividends", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetCorporateActionsDividendsResponse{
		DividendList: make([]*qotgetcorporateactionsdividends.DividendItem, 0, len(s2c.DividendList)),
	}

	for _, item := range s2c.DividendList {
		if item == nil {
			continue
		}
		result.DividendList = append(result.DividendList, item)
	}

	return result, nil
}

type GetCorporateActionsBuybacksRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetCorporateActionsBuybacksResponse struct {
	HkBuyBackList []*qotgetcorporateactionsbuybacks.HKBuyBackItem
	ABuyBackList  []*qotgetcorporateactionsbuybacks.ABuyBackItem
	NextKey       string
}

func GetCorporateActionsBuybacks(ctx context.Context, c *futuapi.Client, req *GetCorporateActionsBuybacksRequest) (*GetCorporateActionsBuybacksResponse, error) {
	if req == nil {
		return nil, wrapError("GetCorporateActionsBuybacks", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetCorporateActionsBuybacks", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetcorporateactionsbuybacks.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num > 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetcorporateactionsbuybacks.Request{C2S: c2s}
	var rsp qotgetcorporateactionsbuybacks.Response

	if err := c.RequestContext(ctx, ProtoID_GetCorporateActionsBuybacks, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCorporateActionsBuybacks", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCorporateActionsBuybacks", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetCorporateActionsBuybacksResponse{
		HkBuyBackList: make([]*qotgetcorporateactionsbuybacks.HKBuyBackItem, 0, len(s2c.HkBuyBackList)),
		ABuyBackList:  make([]*qotgetcorporateactionsbuybacks.ABuyBackItem, 0, len(s2c.ABuyBackList)),
		NextKey:       util.ProtoStr(s2c.NextKey),
	}

	for _, item := range s2c.HkBuyBackList {
		if item == nil {
			continue
		}
		result.HkBuyBackList = append(result.HkBuyBackList, item)
	}

	for _, item := range s2c.ABuyBackList {
		if item == nil {
			continue
		}
		result.ABuyBackList = append(result.ABuyBackList, item)
	}

	return result, nil
}

type GetCorporateActionsStockSplitsRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetCorporateActionsStockSplitsResponse struct {
	SplitItemList []*qotgetcorporateactionsstocksplits.StockSplitItem
	NextKey       string
}

func GetCorporateActionsStockSplits(ctx context.Context, c *futuapi.Client, req *GetCorporateActionsStockSplitsRequest) (*GetCorporateActionsStockSplitsResponse, error) {
	if req == nil {
		return nil, wrapError("GetCorporateActionsStockSplits", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetCorporateActionsStockSplits", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetcorporateactionsstocksplits.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num > 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetcorporateactionsstocksplits.Request{C2S: c2s}
	var rsp qotgetcorporateactionsstocksplits.Response

	if err := c.RequestContext(ctx, ProtoID_GetCorporateActionsStockSplits, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCorporateActionsStockSplits", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCorporateActionsStockSplits", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetCorporateActionsStockSplitsResponse{
		SplitItemList: make([]*qotgetcorporateactionsstocksplits.StockSplitItem, 0, len(s2c.SplitItemList)),
		NextKey:       util.ProtoStr(s2c.NextKey),
	}

	for _, item := range s2c.SplitItemList {
		if item == nil {
			continue
		}
		result.SplitItemList = append(result.SplitItemList, item)
	}

	return result, nil
}