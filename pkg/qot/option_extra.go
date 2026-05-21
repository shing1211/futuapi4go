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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionvolatility"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionexerciseprobability"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetOptionVolatility         = 3250
	ProtoID_GetOptionExerciseProbability = 3251
)

type GetOptionVolatilityRequest struct {
	Security        *qotcommon.Security
	QueryTimePeriod qotcommon.OptionVolatilityTimePeriodType
	HvTimePeriod    int32
}

type GetOptionVolatilityResponse struct {
	ItemList      []*qotgetoptionvolatility.VolatilityItem
	AverageImpvol float64
	ImpvolStatus  int32
	Analysis      string
}

func GetOptionVolatility(ctx context.Context, c *futuapi.Client, req *GetOptionVolatilityRequest) (*GetOptionVolatilityResponse, error) {
	if req == nil {
		return nil, wrapError("GetOptionVolatility", 0, "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetOptionVolatility", 0, "Security is nil")
	}
	c2s := &qotgetoptionvolatility.C2S{
		Security: req.Security,
	}
	if req.QueryTimePeriod != 0 {
		c2s.QueryTimePeriod = &req.QueryTimePeriod
	}
	if req.HvTimePeriod != 0 {
		c2s.HvTimePeriod = &req.HvTimePeriod
	}
	pkt := &qotgetoptionvolatility.Request{C2S: c2s}
	var rsp qotgetoptionvolatility.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionVolatility, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionVolatility", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionVolatility", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetOptionVolatilityResponse{
		ItemList:      s2c.ItemList,
		AverageImpvol: util.ProtoFloat64(s2c.AverageImpvol),
		ImpvolStatus:  int32(s2c.GetImpvolStatus()),
		Analysis:      util.ProtoStr(s2c.Analysis),
	}

	return result, nil
}

type GetOptionExerciseProbabilityRequest struct {
	Security *qotcommon.Security
}

type GetOptionExerciseProbabilityResponse struct {
	ItemList []*qotgetoptionexerciseprobability.StrikeProbabilityItem
}

func GetOptionExerciseProbability(ctx context.Context, c *futuapi.Client, req *GetOptionExerciseProbabilityRequest) (*GetOptionExerciseProbabilityResponse, error) {
	if req == nil {
		return nil, wrapError("GetOptionExerciseProbability", 0, "request is nil")
	}
	if req.Security == nil {
		return nil, wrapError("GetOptionExerciseProbability", 0, "Security is nil")
	}
	c2s := &qotgetoptionexerciseprobability.C2S{
		Security: req.Security,
	}
	pkt := &qotgetoptionexerciseprobability.Request{C2S: c2s}
	var rsp qotgetoptionexerciseprobability.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionExerciseProbability, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionExerciseProbability", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOptionExerciseProbability", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetOptionExerciseProbabilityResponse{
		ItemList: s2c.ItemList,
	}

	return result, nil
}