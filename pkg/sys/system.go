// Package sys provides system-level APIs for the Futu OpenD SDK.
//
// This package covers connection state, user information, delay statistics,
// and verification. These functions work without an active trading account.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/sys"
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

// state, err := sys.GetGlobalState(cli)
package sys

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/getdelaystatistics"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/verification"

	"time"
)

const (
	ProtoID_GetGlobalState     = 1002
	ProtoID_GetUserInfo        = 1005
	ProtoID_Verification       = 1006
	ProtoID_GetDelayStatistics = 1007
)

// GetGlobalStateResponse represents the global connection state including server info, login status, and market availability.
type GetGlobalStateResponse struct {
	ConnID         uint64
	ServerVer      int32
	ServerBuildNo  int32
	Time           int64
	LocalTime      float64
	QotLogined     bool
	TrdLogined     bool
	QotSvrIpAddr   string
	TrdSvrIpAddr   string
	MarketHK       int32
	MarketUS       int32
	MarketSH       int32
	MarketSZ       int32
	MarketHKFuture int32
	MarketUSFuture int32
	MarketSGFuture int32
	MarketJPFuture int32
	ProgramStatus  *common.ProgramStatus
}

// GetGlobalState retrieves the global connection state including server version, login status, and market information.
// Returns the global state or an error if the request fails.
func GetGlobalState(c *futuapi.Client) (*GetGlobalStateResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &getglobalstate.C2S{
		UserID: func() *uint64 { v := uint64(0); return &v }(),
	}

	pkt := &getglobalstate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetGlobalState, serialNo, body); err != nil {
		return nil, err
	}

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return nil, err
	}

	var rsp getglobalstate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetGlobalState failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetGlobalState: s2c is nil")
	}

	return &GetGlobalStateResponse{
		ConnID:         s2c.GetConnID(),
		ServerVer:      s2c.GetServerVer(),
		ServerBuildNo:  s2c.GetServerBuildNo(),
		Time:           s2c.GetTime(),
		LocalTime:      s2c.GetLocalTime(),
		QotLogined:     s2c.GetQotLogined(),
		TrdLogined:     s2c.GetTrdLogined(),
		QotSvrIpAddr:   s2c.GetQotSvrIpAddr(),
		TrdSvrIpAddr:   s2c.GetTrdSvrIpAddr(),
		MarketHK:       s2c.GetMarketHK(),
		MarketUS:       s2c.GetMarketUS(),
		MarketSH:       s2c.GetMarketSH(),
		MarketSZ:       s2c.GetMarketSZ(),
		MarketHKFuture: s2c.GetMarketHKFuture(),
		MarketUSFuture: s2c.GetMarketUSFuture(),
		MarketSGFuture: s2c.GetMarketSGFuture(),
		MarketJPFuture: s2c.GetMarketJPFuture(),
		ProgramStatus:  s2c.GetProgramStatus(),
	}, nil
}

// GetUserInfoResponse represents the user information including user ID, nickname, avatar, and API level.
type GetUserInfoResponse struct {
	UserID                int64
	NickName              string
	AvatarUrl             string
	ApiLevel              string
	IsNeedAgreeDisclaimer bool
}

// GetUserInfo retrieves the current user information including nickname, avatar, and API level.
// Returns the user info or an error if the request fails.
func GetUserInfo(c *futuapi.Client) (*GetUserInfoResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &getuserinfo.C2S{}

	pkt := &getuserinfo.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetUserInfo, serialNo, body); err != nil {
		return nil, err
	}

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return nil, err
	}

	var rsp getuserinfo.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetUserInfo failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetUserInfo: s2c is nil")
	}

	return &GetUserInfoResponse{
		UserID:                s2c.GetUserID(),
		NickName:              s2c.GetNickName(),
		AvatarUrl:             s2c.GetAvatarUrl(),
		ApiLevel:              s2c.GetApiLevel(),
		IsNeedAgreeDisclaimer: s2c.GetIsNeedAgreeDisclaimer(),
	}, nil
}

// GetDelayStatisticsResponse represents delay statistics for quote push, request-reply, and order placement.
type GetDelayStatisticsResponse struct {
	QotPushStatisticsList    []*getdelaystatistics.DelayStatistics
	ReqReplyStatisticsList   []*getdelaystatistics.ReqReplyStatisticsItem
	PlaceOrderStatisticsList []*getdelaystatistics.PlaceOrderStatisticsItem
}

// GetDelayStatistics retrieves performance delay statistics for quote pushes, request-reply, and order placements.
// Returns the delay statistics or an error if the request fails.
func GetDelayStatistics(c *futuapi.Client) (*GetDelayStatisticsResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &getdelaystatistics.C2S{}

	pkt := &getdelaystatistics.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetDelayStatistics, serialNo, body); err != nil {
		return nil, err
	}

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return nil, err
	}

	var rsp getdelaystatistics.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetDelayStatistics failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetDelayStatistics: s2c is nil")
	}

	return &GetDelayStatisticsResponse{
		QotPushStatisticsList:    s2c.GetQotPushStatisticsList(),
		ReqReplyStatisticsList:   s2c.GetReqReplyStatisticsList(),
		PlaceOrderStatisticsList: s2c.GetPlaceOrderStatisticsList(),
	}, nil
}

// VerificationRequest is the request to verify a user with a specified code (e.g., SMS or email verification).
type VerificationRequest struct {
	Type verification.VerificationType
	Op   verification.VerificationOp
	Code string
}

// Verification submits a verification request for user authentication.
// Returns an error if the verification fails.
func Verification(c *futuapi.Client, req *VerificationRequest) error {
	if err := c.EnsureConnected(); err != nil {
		return err
	}
	c2s := &verification.C2S{
		Type: func() *int32 { v := int32(req.Type); return &v }(),
		Op:   func() *int32 { v := int32(req.Op); return &v }(),
		Code: &req.Code,
	}

	pkt := &verification.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_Verification, serialNo, body); err != nil {
		return err
	}

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
	if err != nil {
		return err
	}

	var rsp verification.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return fmt.Errorf("Verification failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}
