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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
)

// GetCapitalFlowRequest defines parameters for GetCapitalFlow.
type GetCapitalFlowRequest struct {
	Security   *qotcommon.Security
	PeriodType int32
	BeginTime  string
	EndTime    string
}

// CapitalFlowItem represents a single capital flow data point.
type CapitalFlowItem struct {
	InFlow      float64
	Time        string
	Timestamp   float64
	MainInFlow  float64
	SuperInFlow float64
	BigInFlow   float64
	MidInFlow   float64
	SmlInFlow   float64
}

// GetCapitalFlowResponse is the response type for GetCapitalFlow.
type GetCapitalFlowResponse struct {
	FlowItemList       []*CapitalFlowItem
	LastValidTime      string
	LastValidTimestamp float64
}

// GetCapitalFlow returns capital flow data for the given security.
func GetCapitalFlow(ctx context.Context, c *futuapi.Client, req *GetCapitalFlowRequest) (*GetCapitalFlowResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCapitalFlow: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetcapitalflow.C2S{
		Security: req.Security,
	}
	if req.PeriodType != 0 {
		c2s.PeriodType = &req.PeriodType
	}
	if req.BeginTime != "" {
		c2s.BeginTime = &req.BeginTime
	}
	if req.EndTime != "" {
		c2s.EndTime = &req.EndTime
	}

	pkt := &qotgetcapitalflow.Request{C2S: c2s}
	var rsp qotgetcapitalflow.Response

	if err := c.RequestContext(ctx, ProtoID_GetCapitalFlow, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCapitalFlow", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetCapitalFlow", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetCapitalFlowResponse{
		FlowItemList:       make([]*CapitalFlowItem, 0, len(s2c.GetFlowItemList())),
		LastValidTime:      s2c.GetLastValidTime(),
		LastValidTimestamp: s2c.GetLastValidTimestamp(),
	}

	for _, f := range s2c.GetFlowItemList() {
		if f == nil {
			continue
		}
		result.FlowItemList = append(result.FlowItemList, &CapitalFlowItem{
			InFlow:      f.GetInFlow(),
			Time:        f.GetTime(),
			Timestamp:   f.GetTimestamp(),
			MainInFlow:  f.GetMainInFlow(),
			SuperInFlow: f.GetSuperInFlow(),
			BigInFlow:   f.GetBigInFlow(),
			MidInFlow:   f.GetMidInFlow(),
			SmlInFlow:   f.GetSmlInFlow(),
		})
	}

	return result, nil
}

// CapitalDistribution represents the distribution of capital across different tiers.
type CapitalDistribution struct {
	CapitalInSuper  float64
	CapitalInBig    float64
	CapitalInMid    float64
	CapitalInSmall  float64
	CapitalOutSuper float64
	CapitalOutBig   float64
	CapitalOutMid   float64
	CapitalOutSmall float64
	UpdateTime      string
	UpdateTimestamp float64
}

// GetCapitalDistributionResponse is the response type for GetCapitalDistribution.
type GetCapitalDistributionResponse struct {
	CapitalDistribution *CapitalDistribution
}

// GetCapitalDistribution returns capital distribution data for the given security.
func GetCapitalDistribution(ctx context.Context, c *futuapi.Client, security *qotcommon.Security) (*GetCapitalDistributionResponse, error) {
	if security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetcapitaldistribution.C2S{
		Security: security,
	}

	pkt := &qotgetcapitaldistribution.Request{C2S: c2s}
	var rsp qotgetcapitaldistribution.Response

	if err := c.RequestContext(ctx, ProtoID_GetCapitalDistribution, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCapitalDistribution", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetCapitalDistribution", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetCapitalDistributionResponse{
		CapitalDistribution: &CapitalDistribution{
			CapitalInSuper:  s2c.GetCapitalInSuper(),
			CapitalInBig:    s2c.GetCapitalInBig(),
			CapitalInMid:    s2c.GetCapitalInMid(),
			CapitalInSmall:  s2c.GetCapitalInSmall(),
			CapitalOutSuper: s2c.GetCapitalOutSuper(),
			CapitalOutBig:   s2c.GetCapitalOutBig(),
			CapitalOutMid:   s2c.GetCapitalOutMid(),
			CapitalOutSmall: s2c.GetCapitalOutSmall(),
			UpdateTime:      s2c.GetUpdateTime(),
			UpdateTimestamp: s2c.GetUpdateTimestamp(),
		},
	}, nil
}
