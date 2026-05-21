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

package qot

import (
	"context"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinsiderholderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinsidertradelist"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetInsiderHolderList = 3241
	ProtoID_GetInsiderTradeList  = 3242
)

type GetInsiderHolderListRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetInsiderHolderListResponse struct {
	ItemList           []*qotgetinsiderholderlist.OwnerInsiderHolderItem
	AllCount           int32
	NextKey            string
	InsiderTotalCount  int32
	InsiderBoughtCount int32
	InsiderSoldCount   int32
}

func GetInsiderHolderList(ctx context.Context, c *futuapi.Client, req *GetInsiderHolderListRequest) (*GetInsiderHolderListResponse, error) {
	if req == nil {
		return nil, wrapError("GetInsiderHolderList", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetInsiderHolderList", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetinsiderholderlist.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num > 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetinsiderholderlist.Request{C2S: c2s}
	var rsp qotgetinsiderholderlist.Response

	if err := c.RequestContext(ctx, ProtoID_GetInsiderHolderList, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInsiderHolderList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetInsiderHolderList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetInsiderHolderListResponse{
		ItemList:           s2c.ItemList,
		AllCount:           util.ProtoInt32(s2c.AllCount),
		NextKey:            util.ProtoStr(s2c.NextKey),
		InsiderTotalCount:  util.ProtoInt32(s2c.InsiderTotalCount),
		InsiderBoughtCount: util.ProtoInt32(s2c.InsiderBoughtCount),
		InsiderSoldCount:   util.ProtoInt32(s2c.InsiderSoldCount),
	}

	return result, nil
}

type GetInsiderTradeListRequest struct {
	Security *qotcommon.Security
	HolderId int64
	NextKey  string
	Num      int32
}

type GetInsiderTradeListResponse struct {
	ItemList []*qotgetinsidertradelist.OwnerInsiderTradeItem
	AllCount int32
	NextKey  string
}

func GetInsiderTradeList(ctx context.Context, c *futuapi.Client, req *GetInsiderTradeListRequest) (*GetInsiderTradeListResponse, error) {
	if req == nil {
		return nil, wrapError("GetInsiderTradeList", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetInsiderTradeList", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetinsidertradelist.C2S{
		Security: req.Security,
	}
	if req.HolderId != 0 {
		c2s.HolderId = &req.HolderId
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num > 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetinsidertradelist.Request{C2S: c2s}
	var rsp qotgetinsidertradelist.Response

	if err := c.RequestContext(ctx, ProtoID_GetInsiderTradeList, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInsiderTradeList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetInsiderTradeList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetInsiderTradeListResponse{
		ItemList: s2c.ItemList,
		AllCount: util.ProtoInt32(s2c.AllCount),
		NextKey:  util.ProtoStr(s2c.NextKey),
	}

	return result, nil
}
