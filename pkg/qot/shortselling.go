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

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdailyshortvolume"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshortinterest"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettoptenbuysellbrokers"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetTopTenBuySellBrokers = 3247
	ProtoID_GetDailyShortVolume     = 3248
	ProtoID_GetShortInterest        = 3249
)

type TopTenBrokerItem struct {
	NetVol        int64
	BrokerName    string
	BuySellType   int32
	AvgPrice      float64
	TotalVol      float64
	TotalTurnover float64
}

type GetTopTenBuySellBrokersRequest struct {
	Security   *qotcommon.Security
	DaysBefore int32
}

type GetTopTenBuySellBrokersResponse struct {
	IsRealTime  bool
	DataTime    int64
	DataTimeStr string
	BrokerList  []*TopTenBrokerItem
}

func GetTopTenBuySellBrokers(ctx context.Context, c *futuapi.Client, req *GetTopTenBuySellBrokersRequest) (*GetTopTenBuySellBrokersResponse, error) {
	if req == nil {
		return nil, wrapError("GetTopTenBuySellBrokers", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetTopTenBuySellBrokers", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgettoptenbuysellbrokers.C2S{
		Security: req.Security,
	}
	if req.DaysBefore != 0 {
		c2s.DaysBefore = &req.DaysBefore
	}
	pkt := &qotgettoptenbuysellbrokers.Request{C2S: c2s}
	var rsp qotgettoptenbuysellbrokers.Response

	if err := c.RequestContext(ctx, ProtoID_GetTopTenBuySellBrokers, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTopTenBuySellBrokers", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetTopTenBuySellBrokers", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetTopTenBuySellBrokersResponse{
		IsRealTime:  util.ProtoBool(s2c.IsRealTime),
		DataTime:    util.ProtoInt64(s2c.DataTime),
		DataTimeStr: util.ProtoStr(s2c.DataTimeStr),
		BrokerList:  make([]*TopTenBrokerItem, 0, len(s2c.BrokerList)),
	}

	for _, b := range s2c.BrokerList {
		if b == nil {
			continue
		}
		result.BrokerList = append(result.BrokerList, &TopTenBrokerItem{
			NetVol:        util.ProtoInt64(b.NetVol),
			BrokerName:    util.ProtoStr(b.BrokerName),
			BuySellType:   int32(b.GetBuySellType()),
			AvgPrice:      util.ProtoFloat64(b.AvgPrice),
			TotalVol:      util.ProtoFloat64(b.TotalVol),
			TotalTurnover: util.ProtoFloat64(b.TotalTurnover),
		})
	}

	return result, nil
}

type GetDailyShortVolumeRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetDailyShortVolumeResponse struct {
	UsItemList           []*qotgetdailyshortvolume.UsDailyShortVolumeItem
	HkItemList           []*qotgetdailyshortvolume.HkDailyShortVolumeItem
	NextKey              string
	AggregatedShort      int64
	AggregatedShortRatio float64
	NewTimeStr           string
}

func GetDailyShortVolume(ctx context.Context, c *futuapi.Client, req *GetDailyShortVolumeRequest) (*GetDailyShortVolumeResponse, error) {
	if req == nil {
		return nil, wrapError("GetDailyShortVolume", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetDailyShortVolume", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetdailyshortvolume.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetdailyshortvolume.Request{C2S: c2s}
	var rsp qotgetdailyshortvolume.Response

	if err := c.RequestContext(ctx, ProtoID_GetDailyShortVolume, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDailyShortVolume", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetDailyShortVolume", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetDailyShortVolumeResponse{
		UsItemList:           s2c.UsItemList,
		HkItemList:           s2c.HkItemList,
		NextKey:              util.ProtoStr(s2c.NextKey),
		AggregatedShort:      util.ProtoInt64(s2c.AggregatedShort),
		AggregatedShortRatio: util.ProtoFloat64(s2c.AggregatedShortRatio),
		NewTimeStr:           util.ProtoStr(s2c.NewTimeStr),
	}, nil
}

type GetShortInterestRequest struct {
	Security *qotcommon.Security
	NextKey  string
	Num      int32
}

type GetShortInterestResponse struct {
	UsItemList []*qotgetshortinterest.UsShortInterestItem
	HkItemList []*qotgetshortinterest.HkShortInterestItem
	NextKey    string
}

func GetShortInterest(ctx context.Context, c *futuapi.Client, req *GetShortInterestRequest) (*GetShortInterestResponse, error) {
	if req == nil {
		return nil, wrapError("GetShortInterest", int32(common.RetType_RetType_Unknown), "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetShortInterest", int32(common.RetType_RetType_Unknown), "Security is nil")
	}
	c2s := &qotgetshortinterest.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetshortinterest.Request{C2S: c2s}
	var rsp qotgetshortinterest.Response

	if err := c.RequestContext(ctx, ProtoID_GetShortInterest, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShortInterest", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetShortInterest", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetShortInterestResponse{
		UsItemList: s2c.UsItemList,
		HkItemList: s2c.HkItemList,
		NextKey:    util.ProtoStr(s2c.NextKey),
	}, nil
}