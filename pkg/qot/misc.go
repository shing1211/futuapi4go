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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcodechange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmarketstate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsuspend"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesttradedate"
)

const (
	ProtoID_GetMarketState = 3223
)

// RequestTradeDateRequest defines parameters for RequestTradeDate.
type RequestTradeDateRequest struct {
	Market    int32
	BeginTime string
	EndTime   string
	Security  *qotcommon.Security
}

// RequestTradeDateResponse is the response type for RequestTradeDate.
type RequestTradeDateResponse struct {
	TradeDateList []*qotrequesttradedate.TradeDate
}

// RequestTradeDate requests trade dates for a specific security.
func RequestTradeDate(ctx context.Context, c *futuapi.Client, req *RequestTradeDateRequest) (*RequestTradeDateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestTradeDate: request is nil")
	}
	if req.Market == 0 {
		return nil, fmt.Errorf("invalid market: must be non-zero")
	}

	c2s := &qotrequesttradedate.C2S{
		Market:    &req.Market,
		BeginTime: &req.BeginTime,
		EndTime:   &req.EndTime,
		Security:  req.Security,
	}

	pkt := &qotrequesttradedate.Request{C2S: c2s}
	var rsp qotrequesttradedate.Response

	if err := c.RequestContext(ctx, ProtoID_RequestTradeDate, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestTradeDate", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestTradeDate: s2c is nil")
	}

	return &RequestTradeDateResponse{
		TradeDateList: s2c.GetTradeDateList(),
	}, nil
}

// GetSuspendRequest defines parameters for GetSuspend.
type GetSuspendRequest struct {
	SecurityList []*qotcommon.Security
	BeginTime    string
	EndTime      string
}

// SuspendInfo represents the suspension time for a security.
type SuspendInfo struct {
	Time      string
	Timestamp float64
}

// SecuritySuspendInfo represents suspension info for a single security.
type SecuritySuspendInfo struct {
	Security    *qotcommon.Security
	SuspendList []*SuspendInfo
}

// GetSuspendResponse is the response type for GetSuspend.
type GetSuspendResponse struct {
	SecuritySuspendList []*SecuritySuspendInfo
}

// GetSuspend returns suspension (停牌) information for the given securities.
func GetSuspend(ctx context.Context, c *futuapi.Client, req *GetSuspendRequest) (*GetSuspendResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSuspend: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetsuspend.C2S{
		SecurityList: req.SecurityList,
		BeginTime:    &req.BeginTime,
		EndTime:      &req.EndTime,
	}

	pkt := &qotgetsuspend.Request{C2S: c2s}
	var rsp qotgetsuspend.Response

	if err := c.RequestContext(ctx, ProtoID_GetSuspend, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetSuspend", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetSuspend: s2c is nil")
	}

	result := &GetSuspendResponse{
		SecuritySuspendList: make([]*SecuritySuspendInfo, 0, len(s2c.GetSecuritySuspendList())),
	}

	for _, ss := range s2c.GetSecuritySuspendList() {
		if ss == nil {
			continue
		}
		info := &SecuritySuspendInfo{
			Security:    ss.GetSecurity(),
			SuspendList: make([]*SuspendInfo, 0, len(ss.GetSuspendList())),
		}
		for _, s := range ss.GetSuspendList() {
			if s == nil {
				continue
			}
			info.SuspendList = append(info.SuspendList, &SuspendInfo{
				Time:      s.GetTime(),
				Timestamp: s.GetTimestamp(),
			})
		}
		result.SecuritySuspendList = append(result.SecuritySuspendList, info)
	}

	return result, nil
}

// GetCodeChangeRequest defines parameters for GetCodeChange.
type GetCodeChangeRequest struct {
	SecurityList   []*qotcommon.Security
	TimeFilterList []*qotgetcodechange.TimeFilter
	TypeList       []int32
}

// CodeChangeInfo represents information about a code change (stock split, merger, etc.).
type CodeChangeInfo struct {
	Type               int32
	Security           *qotcommon.Security
	RelatedSecurity    *qotcommon.Security
	PublicTime         string
	PublicTimestamp    float64
	EffectiveTime      string
	EffectiveTimestamp float64
	EndTime            string
	EndTimestamp       float64
}

// GetCodeChangeResponse is the response type for GetCodeChange.
type GetCodeChangeResponse struct {
	CodeChangeList []*CodeChangeInfo
}

// GetCodeChange returns code change (股份代号变动) information for the given securities.
func GetCodeChange(ctx context.Context, c *futuapi.Client, req *GetCodeChangeRequest) (*GetCodeChangeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCodeChange: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetcodechange.C2S{
		SecurityList:   req.SecurityList,
		TimeFilterList: req.TimeFilterList,
		TypeList:       req.TypeList,
	}

	pkt := &qotgetcodechange.Request{C2S: c2s}
	var rsp qotgetcodechange.Response

	if err := c.RequestContext(ctx, ProtoID_GetCodeChange, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCodeChange", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetCodeChange: s2c is nil")
	}

	result := &GetCodeChangeResponse{
		CodeChangeList: make([]*CodeChangeInfo, 0, len(s2c.GetCodeChangeList())),
	}

	for _, cc := range s2c.GetCodeChangeList() {
		if cc == nil {
			continue
		}
		result.CodeChangeList = append(result.CodeChangeList, &CodeChangeInfo{
			Type:               cc.GetType(),
			Security:           cc.GetSecurity(),
			RelatedSecurity:    cc.GetRelatedSecurity(),
			PublicTime:         cc.GetPublicTime(),
			PublicTimestamp:    cc.GetPublicTimestamp(),
			EffectiveTime:      cc.GetEffectiveTime(),
			EffectiveTimestamp: cc.GetEffectiveTimestamp(),
			EndTime:            cc.GetEndTime(),
			EndTimestamp:       cc.GetEndTimestamp(),
		})
	}

	return result, nil
}

// MarketStateInfo represents the market state for a single security.
type MarketStateInfo struct {
	Security    *qotcommon.Security
	Name        string
	MarketState int32
}

// GetMarketStateRequest defines parameters for GetMarketState.
type GetMarketStateRequest struct {
	SecurityList []*qotcommon.Security
}

// GetMarketStateResponse is the response type for GetMarketState.
type GetMarketStateResponse struct {
	MarketInfoList []*MarketStateInfo
}

// GetMarketState returns the market state (trading status) for the given securities.
func GetMarketState(ctx context.Context, c *futuapi.Client, req *GetMarketStateRequest) (*GetMarketStateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMarketState: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetmarketstate.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetmarketstate.Request{C2S: c2s}
	var rsp qotgetmarketstate.Response

	if err := c.RequestContext(ctx, ProtoID_GetMarketState, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMarketState", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetMarketState: s2c is nil")
	}

	result := &GetMarketStateResponse{
		MarketInfoList: make([]*MarketStateInfo, 0, len(s2c.GetMarketInfoList())),
	}

	for _, mi := range s2c.GetMarketInfoList() {
		if mi == nil {
			continue
		}
		result.MarketInfoList = append(result.MarketInfoList, &MarketStateInfo{
			Security:    mi.GetSecurity(),
			Name:        mi.GetName(),
			MarketState: mi.GetMarketState(),
		})
	}

	return result, nil
}
