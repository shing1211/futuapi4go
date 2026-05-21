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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetvaluationdetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetvaluationplatestocklist"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetValuationDetail          = 3232
	ProtoID_GetValuationPlateStockList  = 3233
)

type GetValuationDetailRequest struct {
	Security      *qotcommon.Security
	ValuationType int32
	IntervalType  int32
}

type GetValuationDetailResponse struct {
	S2C *qotgetvaluationdetail.S2C
}

func GetValuationDetail(ctx context.Context, c *futuapi.Client, req *GetValuationDetailRequest) (*GetValuationDetailResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetValuationDetail: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetValuationDetail: Security is nil")
	}
	c2s := &qotgetvaluationdetail.C2S{
		Security: req.Security,
	}
	if req.ValuationType != 0 {
		vt := qotcommon.ValuationType(req.ValuationType)
		c2s.ValuationType = &vt
	}
	if req.IntervalType != 0 {
		it := qotcommon.ValuationIntervalType(req.IntervalType)
		c2s.IntervalType = &it
	}
	pkt := &qotgetvaluationdetail.Request{C2S: c2s}
	var rsp qotgetvaluationdetail.Response

	if err := c.RequestContext(ctx, ProtoID_GetValuationDetail, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetValuationDetail", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetValuationDetail", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetValuationDetailResponse{S2C: s2c}, nil
}

type GetValuationPlateStockListRequest struct {
	Security       *qotcommon.Security
	ValuationType  int32
	NextKey        string
	Num            int32
	SortType       int32
	SortId         int32
	FilterSecurity *qotcommon.Security
}

type GetValuationPlateStockListResponse struct {
	Count     int32
	StockList []*qotgetvaluationplatestocklist.StockItem
	NextKey   string
	PlateList []*qotgetvaluationplatestocklist.PlateItem
}

func GetValuationPlateStockList(ctx context.Context, c *futuapi.Client, req *GetValuationPlateStockListRequest) (*GetValuationPlateStockListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetValuationPlateStockList: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetValuationPlateStockList: Security is nil")
	}
	c2s := &qotgetvaluationplatestocklist.C2S{
		Security: req.Security,
	}
	if req.ValuationType != 0 {
		vt := qotcommon.ValuationType(req.ValuationType)
		c2s.ValuationType = &vt
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	if req.SortType != 0 {
		st := qotcommon.SortType(req.SortType)
		c2s.SortType = &st
	}
	if req.SortId != 0 {
		si := qotcommon.SortField(req.SortId)
		c2s.SortId = &si
	}
	if req.FilterSecurity != nil {
		c2s.FilterSecurity = req.FilterSecurity
	}
	pkt := &qotgetvaluationplatestocklist.Request{C2S: c2s}
	var rsp qotgetvaluationplatestocklist.Response

	if err := c.RequestContext(ctx, ProtoID_GetValuationPlateStockList, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetValuationPlateStockList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetValuationPlateStockList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetValuationPlateStockListResponse{
		Count:     util.ProtoInt32(s2c.Count),
		StockList: s2c.StockList,
		NextKey:   util.ProtoStr(s2c.NextKey),
		PlateList: s2c.PlateList,
	}

	return result, nil
}
