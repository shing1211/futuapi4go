// Package sys provides system-level APIs for the Futu OpenD SDK.
//
// This package covers connection state, user information, delay statistics,
// and verification. These functions work without an active trading account.
//
// For Python SDK migration, use ProtoIDs from the constant package:
//
//	import "github.com/shing1211/futuapi4go/pkg/constant"
//
//	// ProtoIDs for system functions:
//	// constant.ProtoID_GetGlobalState
//	// constant.ProtoID_GetUserInfo
//	// constant.ProtoID_KeepAlive
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
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/getdelaystatistics"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/verification"
)

// wrapError standardizes error messages for proto response failures
func wrapError(funcName string, retType int32, retMsg string) error {
	code := constant.ErrorCode(retType)
	if retType == 0 {
		code = constant.ErrCodeSuccess
	} else if retType < 0 {
		code = constant.ErrorCode(retType)
	} else {
		code = constant.ErrCodeUnknown
	}
	return &constant.FutuError{
		Code:    code,
		Message: retMsg,
		Func:   funcName,
	}
}

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
func GetGlobalState(ctx context.Context, c *futuapi.Client) (*GetGlobalStateResponse, error) {
	c2s := &getglobalstate.C2S{
		UserID: func() *uint64 { v := uint64(0); return &v }(),
	}

	pkt := &getglobalstate.Request{C2S: c2s}
	var rsp getglobalstate.Response

	if err := c.RequestContext(ctx, ProtoID_GetGlobalState, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetGlobalState", rsp.GetRetType(), rsp.GetRetMsg())
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
	ShQotRight            int32 //上海市场行情权限
	SzQotRight            int32 //深圳市场行情权限
	Extra                 int32 //透传信息
}

// GetUserInfo retrieves the current user information including nickname, avatar, and API level.
// Returns the user info or an error if the request fails.
func GetUserInfo(ctx context.Context, c *futuapi.Client) (*GetUserInfoResponse, error) {
	c2s := &getuserinfo.C2S{}

	pkt := &getuserinfo.Request{C2S: c2s}
	var rsp getuserinfo.Response

	if err := c.RequestContext(ctx, ProtoID_GetUserInfo, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUserInfo", rsp.GetRetType(), rsp.GetRetMsg())
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
		ShQotRight:            s2c.GetShQotRight(),
		SzQotRight:            s2c.GetSzQotRight(),
		Extra:                 s2c.GetExtra(),
	}, nil
}

// GetDelayStatisticsResponse represents delay statistics for quote push, request-reply, and order placement.
type GetDelayStatisticsResponse struct {
	QotPushStatisticsList    []*getdelaystatistics.DelayStatistics
	ReqReplyStatisticsList   []*getdelaystatistics.ReqReplyStatisticsItem
	PlaceOrderStatisticsList []*getdelaystatistics.PlaceOrderStatisticsItem
}

// marshalC2SProto2 marshals the C2S message using proto2 wire format.
// This is a workaround for the proto2 vs proto3 wire format incompatibility.
// Proto2 uses non-packed encoding for repeated int32, while proto3 uses packed encoding.
// OpenD's C++ parser expects proto2 non-packed encoding.
func marshalC2SProto2(c2s *getdelaystatistics.C2S) ([]byte, error) {
	buf := make([]byte, 0, 64)

	// Field 1: TypeList (proto2 non-packed encoding)
	// Wire type 0 = varint, field number 1 -> tag = (1 << 3) | 0 = 8
	for _, v := range c2s.GetTypeList() {
		buf = append(buf, 8) // tag for field 1, wire type 0
		buf = appendVarint(buf, uint64(v))
	}

	// Field 2: QotPushStage (optional int32)
	if c2s.QotPushStage != nil {
		// Wire type 0 = varint, field number 2 -> tag = (2 << 3) | 0 = 16
		buf = append(buf, 16)
		buf = appendVarint(buf, uint64(*c2s.QotPushStage))
	}

	// Field 3: SegmentList (proto2 non-packed encoding)
	// Wire type 0 = varint, field number 3 -> tag = (3 << 3) | 0 = 24
	for _, v := range c2s.GetSegmentList() {
		buf = append(buf, 24) // tag for field 3, wire type 0
		buf = appendVarint(buf, uint64(v))
	}

	return buf, nil
}

// appendVarint appends a varint to the buffer.
func appendVarint(buf []byte, v uint64) []byte {
	for v >= 0x80 {
		buf = append(buf, byte(v)|0x80)
		v >>= 7
	}
	buf = append(buf, byte(v))
	return buf
}

// marshalGetDelayStatisticsRequest marshals the GetDelayStatistics request using proto2 wire format.
func marshalGetDelayStatisticsRequest(c2s *getdelaystatistics.C2S) ([]byte, error) {
	c2sBuf, err := marshalC2SProto2(c2s)
	if err != nil {
		return nil, err
	}

	// Wrap C2S in Request message with proto2 length-delimited encoding
	// Field 1, wire type 2 = length-delimited
	buf := make([]byte, 0, len(c2sBuf)+10)
	buf = append(buf, 0x0A) // tag for field 1, wire type 2 (length-delimited)
	buf = appendVarint(buf, uint64(len(c2sBuf)))
	buf = append(buf, c2sBuf...)

	return buf, nil
}

// GetDelayStatistics retrieves performance delay statistics for quote pushes, request-reply, and order placements.
// Returns the delay statistics or an error if the request fails.
//
// Note: This function uses proto2 wire format for compatibility with OpenD's C++ protobuf parser.
func GetDelayStatistics(ctx context.Context, c *futuapi.Client) (*GetDelayStatisticsResponse, error) {
	c2s := &getdelaystatistics.C2S{}

	// Use custom proto2 marshaling to avoid proto3 packed encoding issue
	body, err := marshalGetDelayStatisticsRequest(c2s)
	if err != nil {
		return nil, fmt.Errorf("marshalGetDelayStatisticsRequest failed: %w", err)
	}

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetDelayStatistics, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadResponseContext(ctx, serialNo, apiTimeout)
	if err != nil {
		return nil, err
	}

	var rsp getdelaystatistics.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDelayStatistics", rsp.GetRetType(), rsp.GetRetMsg())
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
func Verification(ctx context.Context, c *futuapi.Client, req *VerificationRequest) error {
	// Input validation
	if req.Code == "" {
		return fmt.Errorf("verification code is required")
	}

	c2s := &verification.C2S{
		Type: func() *int32 { v := int32(req.Type); return &v }(),
		Op:   func() *int32 { v := int32(req.Op); return &v }(),
		Code: &req.Code,
	}

	pkt := &verification.Request{C2S: c2s}
	var rsp verification.Response

	if err := c.RequestContext(ctx, ProtoID_Verification, pkt, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return wrapError("Verification", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}
