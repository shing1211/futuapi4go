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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersholderdetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersholdingchanges"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersinstitutional"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersoverview"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetShareholdersOverview    = 3237
	ProtoID_GetShareholdersHoldingChanges = 3238
	ProtoID_GetShareholdersHolderDetail = 3239
	ProtoID_GetShareholdersInstitutional = 3240
)

type GetShareholdersOverviewRequest struct {
	Security *qotcommon.Security
	PeriodId int32
}

type GetShareholdersOverviewResponse struct {
	MainHolderInfoList []*qotgetshareholdersoverview.OwnershipStaticInfo
	HolderTypeInfoList []*qotgetshareholdersoverview.OwnershipStaticInfo
	HoldingPeriodList  []*qotgetshareholdersoverview.HoldingPeriodItem
}

func GetShareholdersOverview(ctx context.Context, c *futuapi.Client, req *GetShareholdersOverviewRequest) (*GetShareholdersOverviewResponse, error) {
	if req == nil {
		return nil, wrapError("GetShareholdersOverview", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetShareholdersOverview", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetshareholdersoverview.C2S{
		Security: req.Security,
	}
	if req.PeriodId != 0 {
		c2s.PeriodId = &req.PeriodId
	}
	pkt := &qotgetshareholdersoverview.Request{C2S: c2s}
	var rsp qotgetshareholdersoverview.Response

	if err := c.RequestContext(ctx, ProtoID_GetShareholdersOverview, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShareholdersOverview", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetShareholdersOverview", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetShareholdersOverviewResponse{
		MainHolderInfoList: s2c.MainHolderInfoList,
		HolderTypeInfoList: s2c.HolderTypeInfoList,
		HoldingPeriodList:  s2c.HoldingPeriodList,
	}, nil
}

type GetShareholdersHoldingChangesRequest struct {
	Security   *qotcommon.Security
	NextKey    string
	Num        int32
	SortType   int32
	SortColumn int32
	FilterType int32
}

type GetShareholdersHoldingChangesResponse struct {
	ItemList []*qotgetshareholdersholdingchanges.OwnerListItem
	NextKey  string
}

func GetShareholdersHoldingChanges(ctx context.Context, c *futuapi.Client, req *GetShareholdersHoldingChangesRequest) (*GetShareholdersHoldingChangesResponse, error) {
	if req == nil {
		return nil, wrapError("GetShareholdersHoldingChanges", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetShareholdersHoldingChanges", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetshareholdersholdingchanges.C2S{
		Security: req.Security,
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
	if req.SortColumn != 0 {
		sc := qotcommon.SortField(req.SortColumn)
		c2s.SortColumn = &sc
	}
	if req.FilterType != 0 {
		ft := qotcommon.HoldingChangesFilterType(req.FilterType)
		c2s.FilterType = &ft
	}
	pkt := &qotgetshareholdersholdingchanges.Request{C2S: c2s}
	var rsp qotgetshareholdersholdingchanges.Response

	if err := c.RequestContext(ctx, ProtoID_GetShareholdersHoldingChanges, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShareholdersHoldingChanges", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetShareholdersHoldingChanges", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetShareholdersHoldingChangesResponse{
		ItemList: s2c.ItemList,
		NextKey:  util.ProtoStr(s2c.NextKey),
	}, nil
}

type GetShareholdersHolderDetailRequest struct {
	Security    *qotcommon.Security
	RequestType int32
	NextKey     string
	Num         int32
	SortColumn  int32
	SortType    int32
	PeriodId    int32
	HolderId    int32
}

type GetShareholdersHolderDetailResponse struct {
	UpdateTime    uint64
	UpdateTimeStr string
	NextKey       string
	ItemList      []*qotgetshareholdersholderdetail.OwnershipDetailItem
}

func GetShareholdersHolderDetail(ctx context.Context, c *futuapi.Client, req *GetShareholdersHolderDetailRequest) (*GetShareholdersHolderDetailResponse, error) {
	if req == nil {
		return nil, wrapError("GetShareholdersHolderDetail", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetShareholdersHolderDetail", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetshareholdersholderdetail.C2S{
		Security: req.Security,
	}
	if req.RequestType != 0 {
		rt := qotcommon.HolderDetailType(req.RequestType)
		c2s.RequestType = &rt
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	if req.SortColumn != 0 {
		sc := qotcommon.SortField(req.SortColumn)
		c2s.SortColumn = &sc
	}
	if req.SortType != 0 {
		st := qotcommon.SortType(req.SortType)
		c2s.SortType = &st
	}
	if req.PeriodId != 0 {
		c2s.PeriodId = &req.PeriodId
	}
	if req.HolderId != 0 {
		c2s.HolderId = &req.HolderId
	}
	pkt := &qotgetshareholdersholderdetail.Request{C2S: c2s}
	var rsp qotgetshareholdersholderdetail.Response

	if err := c.RequestContext(ctx, ProtoID_GetShareholdersHolderDetail, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShareholdersHolderDetail", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetShareholdersHolderDetail", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetShareholdersHolderDetailResponse{
		UpdateTime:    util.ProtoUint64(s2c.UpdateTime),
		UpdateTimeStr: util.ProtoStr(s2c.UpdateTimeStr),
		NextKey:       util.ProtoStr(s2c.NextKey),
		ItemList:      s2c.ItemList,
	}, nil
}

type GetShareholdersInstitutionalRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetShareholdersInstitutionalResponse struct {
	UpdateTime    uint64
	UpdateTimeStr string
	NextKey       string
	ItemList      []*qotgetshareholdersinstitutional.InstitutionHolderItem
}

func GetShareholdersInstitutional(ctx context.Context, c *futuapi.Client, req *GetShareholdersInstitutionalRequest) (*GetShareholdersInstitutionalResponse, error) {
	if req == nil {
		return nil, wrapError("GetShareholdersInstitutional", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetShareholdersInstitutional", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetshareholdersinstitutional.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetshareholdersinstitutional.Request{C2S: c2s}
	var rsp qotgetshareholdersinstitutional.Response

	if err := c.RequestContext(ctx, ProtoID_GetShareholdersInstitutional, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShareholdersInstitutional", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetShareholdersInstitutional", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetShareholdersInstitutionalResponse{
		UpdateTime:    util.ProtoUint64(s2c.UpdateTime),
		UpdateTimeStr: util.ProtoStr(s2c.UpdateTimeStr),
		NextKey:       util.ProtoStr(s2c.NextKey),
		ItemList:      s2c.ItemList,
	}, nil
}