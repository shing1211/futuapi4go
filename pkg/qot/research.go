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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchanalystconsensus"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchmorningstarrpt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchratingsummary"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetResearchAnalystConsensus = 3229
	ProtoID_GetResearchRatingSummary    = 3230
	ProtoID_GetResearchMorningstarReport = 3231
)

type GetResearchAnalystConsensusRequest struct {
	Security *qotcommon.Security
}

type GetResearchAnalystConsensusResponse struct {
	Highest       float64
	Average       float64
	Lowest        float64
	Rating        int32
	Total         int32
	UpdateTime    int64
	UpdateTimeStr string
	Buy           float64
	Hold          float64
	Sell          float64
	StrongBuy     float64
	Underperform  float64
}

func GetResearchAnalystConsensus(ctx context.Context, c *futuapi.Client, req *GetResearchAnalystConsensusRequest) (*GetResearchAnalystConsensusResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetResearchAnalystConsensus: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetResearchAnalystConsensus: Security is nil")
	}
	c2s := &qotgetresearchanalystconsensus.C2S{
		Security: req.Security,
	}
	pkt := &qotgetresearchanalystconsensus.Request{C2S: c2s}
	var rsp qotgetresearchanalystconsensus.Response

	if err := c.RequestContext(ctx, ProtoID_GetResearchAnalystConsensus, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetResearchAnalystConsensus", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetResearchAnalystConsensus", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	var rating int32
	if s2c.Rating != nil {
		rating = int32(*s2c.Rating)
	}

	return &GetResearchAnalystConsensusResponse{
		Highest:       util.ProtoFloat64(s2c.Highest),
		Average:       util.ProtoFloat64(s2c.Average),
		Lowest:        util.ProtoFloat64(s2c.Lowest),
		Rating:        rating,
		Total:         util.ProtoInt32(s2c.Total),
		UpdateTime:    util.ProtoInt64(s2c.UpdateTime),
		UpdateTimeStr: util.ProtoStr(s2c.UpdateTimeStr),
		Buy:           util.ProtoFloat64(s2c.Buy),
		Hold:          util.ProtoFloat64(s2c.Hold),
		Sell:          util.ProtoFloat64(s2c.Sell),
		StrongBuy:     util.ProtoFloat64(s2c.StrongBuy),
		Underperform:  util.ProtoFloat64(s2c.Underperform),
	}, nil
}

type GetResearchRatingSummaryRequest struct {
	Security            *qotcommon.Security
	RatingDimensionType *qotcommon.ResearchRatingDimensionType
	Uid                 string
	NextKey             string
	Num                 int32
}

type GetResearchRatingSummaryResponse struct {
	S2C *qotgetresearchratingsummary.S2C
}

func GetResearchRatingSummary(ctx context.Context, c *futuapi.Client, req *GetResearchRatingSummaryRequest) (*GetResearchRatingSummaryResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetResearchRatingSummary: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetResearchRatingSummary: Security is nil")
	}
	c2s := &qotgetresearchratingsummary.C2S{
		Security: req.Security,
	}
	if req.RatingDimensionType != nil {
		c2s.RatingDimensionType = req.RatingDimensionType
	}
	if req.Uid != "" {
		c2s.Uid = &req.Uid
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetresearchratingsummary.Request{C2S: c2s}
	var rsp qotgetresearchratingsummary.Response

	if err := c.RequestContext(ctx, ProtoID_GetResearchRatingSummary, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetResearchRatingSummary", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetResearchRatingSummary", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetResearchRatingSummaryResponse{
		S2C: s2c,
	}, nil
}

type GetResearchMorningstarReportRequest struct {
	Security *qotcommon.Security
}

type GetResearchMorningstarReportResponse struct {
	S2C *qotgetresearchmorningstarrpt.S2C
}

func GetResearchMorningstarReport(ctx context.Context, c *futuapi.Client, req *GetResearchMorningstarReportRequest) (*GetResearchMorningstarReportResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetResearchMorningstarReport: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetResearchMorningstarReport: Security is nil")
	}
	c2s := &qotgetresearchmorningstarrpt.C2S{
		Security: req.Security,
	}
	pkt := &qotgetresearchmorningstarrpt.Request{C2S: c2s}
	var rsp qotgetresearchmorningstarrpt.Response

	if err := c.RequestContext(ctx, ProtoID_GetResearchMorningstarReport, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetResearchMorningstarReport", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetResearchMorningstarReport", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetResearchMorningstarReportResponse{
		S2C: s2c,
	}, nil
}
