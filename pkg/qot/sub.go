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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsecuritysnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotregqotpush"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
)

const (
	ProtoID_GetSubInfo = 3003
)

// SubType represents the type of market data subscription.
type SubType int32

const (
	SubType_Basic      SubType = 1
	SubType_OrderBook  SubType = 2
	SubType_Ticker     SubType = 4
	SubType_RT         SubType = 5
	SubType_KL_Day     SubType = 6
	SubType_KL_5Min    SubType = 7
	SubType_KL_15Min   SubType = 8
	SubType_KL_30Min   SubType = 9
	SubType_KL_60Min   SubType = 10
	SubType_KL_1Min    SubType = 11
	SubType_KL_Week    SubType = 12
	SubType_KL_Month   SubType = 13
	SubType_Broker     SubType = 14
	SubType_KL_Quarter SubType = 15
	SubType_KL_Year    SubType = 16
	SubType_KL_3Min    SubType = 17
	SubType_KL         SubType = 6
)

// SubscribeRequest defines parameters for Subscribe.
type SubscribeRequest struct {
	SecurityList         []*qotcommon.Security
	SubTypeList          []SubType
	IsSubOrUnSub         bool
	IsRegOrUnRegPush     bool
	RegPushRehabTypeList []int32
	IsFirstPush          bool
	IsUnsubAll           bool
	IsSubOrderBookDetail bool
	ExtendedTime         bool
	Session              int32
}

// Subscribe subscribes to or unsubscribes from real-time market data.
func Subscribe(ctx context.Context, c *futuapi.Client, req *SubscribeRequest) error {
	if req == nil {
		return fmt.Errorf("Subscribe: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return fmt.Errorf("security list is empty")
	}
	if len(req.SubTypeList) == 0 {
		return fmt.Errorf("subtype list is empty")
	}
	subTypeList := make([]int32, len(req.SubTypeList))
	for i, st := range req.SubTypeList {
		subTypeList[i] = int32(st)
	}

	c2s := &qotsub.C2S{
		SecurityList: req.SecurityList,
		SubTypeList:  subTypeList,
		IsSubOrUnSub: &req.IsSubOrUnSub,
	}
	if req.IsRegOrUnRegPush {
		c2s.IsRegOrUnRegPush = &req.IsRegOrUnRegPush
	}
	if len(req.RegPushRehabTypeList) > 0 {
		c2s.RegPushRehabTypeList = req.RegPushRehabTypeList
	}
	c2s.IsFirstPush = &req.IsFirstPush
	if req.IsUnsubAll {
		c2s.IsUnsubAll = &req.IsUnsubAll
	}
	if req.IsSubOrderBookDetail {
		c2s.IsSubOrderBookDetail = &req.IsSubOrderBookDetail
	}
	if req.ExtendedTime {
		c2s.ExtendedTime = &req.ExtendedTime
	}
	if req.Session != 0 {
		c2s.Session = &req.Session
	}

	pkt := &qotsub.Request{C2S: c2s}
	var rsp qotsub.Response

	if err := c.RequestContext(ctx, ProtoID_Subscribe, pkt, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return wrapError("Subscribe", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}

// RegQotPushRequest defines parameters for RegQotPush.
type RegQotPushRequest struct {
	SecurityList  []*qotcommon.Security
	SubTypeList   []int32
	RehabTypeList []int32
	IsRegOrUnReg  bool
	IsFirstPush   bool
}

// RegQotPush registers or unregisters for real-time push notifications.
func RegQotPush(ctx context.Context, c *futuapi.Client, req *RegQotPushRequest) error {
	if req == nil {
		return fmt.Errorf("RegQotPush: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return fmt.Errorf("security list is empty")
	}
	if len(req.SubTypeList) == 0 {
		return fmt.Errorf("subtype list is empty")
	}

	c2s := &qotregqotpush.C2S{
		SecurityList:  req.SecurityList,
		SubTypeList:   req.SubTypeList,
		RehabTypeList: req.RehabTypeList,
		IsRegOrUnReg:  &req.IsRegOrUnReg,
	}
	c2s.IsFirstPush = &req.IsFirstPush

	pkt := &qotregqotpush.Request{C2S: c2s}
	var rsp qotregqotpush.Response

	if err := c.RequestContext(ctx, ProtoID_RegQotPush, pkt, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return wrapError("RegQotPush", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}

// GetSubInfoResponse is the response type for GetSubInfo.
type GetSubInfoResponse struct {
	ConnSubInfoList []*qotcommon.ConnSubInfo
	TotalUsedQuota  int32
	RemainQuota     int32
}

// GetSubInfo returns subscription information and quota usage.
func GetSubInfo(ctx context.Context, c *futuapi.Client) (*GetSubInfoResponse, error) {
	c2s := &qotgetsubinfo.C2S{}

	pkt := &qotgetsubinfo.Request{C2S: c2s}
	var rsp qotgetsubinfo.Response

	if err := c.RequestContext(ctx, ProtoID_GetSubInfo, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetSubInfo", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetSubInfo", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetSubInfoResponse{
		ConnSubInfoList: s2c.GetConnSubInfoList(),
		TotalUsedQuota:  s2c.GetTotalUsedQuota(),
		RemainQuota:     s2c.GetRemainQuota(),
	}, nil
}

// GetSecuritySnapshotRequest defines parameters for GetSecuritySnapshot.
type GetSecuritySnapshotRequest struct {
	SecurityList []*qotcommon.Security
}

// GetSecuritySnapshotResponse is the response type for GetSecuritySnapshot.
type GetSecuritySnapshotResponse struct {
	SnapshotList []*qotgetsecuritysnapshot.Snapshot
}

// GetSecuritySnapshot returns snapshot data for the given securities.
func GetSecuritySnapshot(ctx context.Context, c *futuapi.Client, req *GetSecuritySnapshotRequest) (*GetSecuritySnapshotResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSecuritySnapshot: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetsecuritysnapshot.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetsecuritysnapshot.Request{C2S: c2s}
	var rsp qotgetsecuritysnapshot.Response

	if err := c.RequestContext(ctx, ProtoID_GetSecuritySnapshot, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetSecuritySnapshot", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetSecuritySnapshot", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetSecuritySnapshotResponse{
		SnapshotList: s2c.GetSnapshotList(),
	}, nil
}

// GetStaticInfoRequest defines parameters for GetStaticInfo.
type GetStaticInfoRequest struct {
	Market       int32
	SecType      int32
	SecurityList []*qotcommon.Security
}

// GetStaticInfoResponse is the response type for GetStaticInfo.
type GetStaticInfoResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetStaticInfo returns static info for the given securities.
func GetStaticInfo(ctx context.Context, c *futuapi.Client, req *GetStaticInfoRequest) (*GetStaticInfoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetStaticInfo: request is nil")
	}
	if len(req.SecurityList) == 0 && req.Market == 0 {
		return nil, fmt.Errorf("invalid market: must be non-zero when no securities provided")
	}

	c2s := &qotgetstaticinfo.C2S{
		Market:       &req.Market,
		SecType:      &req.SecType,
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetstaticinfo.Request{C2S: c2s}
	var rsp qotgetstaticinfo.Response

	if err := c.RequestContext(ctx, ProtoID_GetStaticInfo, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetStaticInfo", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetStaticInfo", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetStaticInfoResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}
