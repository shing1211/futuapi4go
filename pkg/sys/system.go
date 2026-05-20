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
	"github.com/shing1211/futuapi4go/pkg/pb/testcmd"
	"github.com/shing1211/futuapi4go/pkg/pb/usedquota"
	"github.com/shing1211/futuapi4go/pkg/pb/verification"
	"github.com/shing1211/futuapi4go/pkg/util"
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
	return constant.NewFutuError(code, funcName, retMsg)
}

const (
	ProtoID_GetGlobalState      = 1002
	ProtoID_GetUserInfo        = 1005
	ProtoID_Verification       = 1006
	ProtoID_TestCmd            = 1008
	ProtoID_GetDelayStatistics  = 1007
	ProtoID_UsedQuota        = 1010
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetGlobalState", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetGlobalState", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

return &GetGlobalStateResponse{
		ConnID:         util.ProtoUint64(s2c.ConnID),
		ServerVer:      util.ProtoInt32(s2c.ServerVer),
		ServerBuildNo:  util.ProtoInt32(s2c.ServerBuildNo),
		Time:           util.ProtoInt64(s2c.Time),
		LocalTime:      util.ProtoFloat64(s2c.LocalTime),
		QotLogined:     util.ProtoBool(s2c.QotLogined),
		TrdLogined:     util.ProtoBool(s2c.TrdLogined),
		QotSvrIpAddr:   util.ProtoStr(s2c.QotSvrIpAddr),
		TrdSvrIpAddr:   util.ProtoStr(s2c.TrdSvrIpAddr),
		MarketHK:       util.ProtoInt32(s2c.MarketHK),
		MarketUS:       util.ProtoInt32(s2c.MarketUS),
		MarketSH:       util.ProtoInt32(s2c.MarketSH),
		MarketSZ:       util.ProtoInt32(s2c.MarketSZ),
		MarketHKFuture: util.ProtoInt32(s2c.MarketHKFuture),
		MarketUSFuture: util.ProtoInt32(s2c.MarketUSFuture),
		MarketSGFuture: util.ProtoInt32(s2c.MarketSGFuture),
		MarketJPFuture: util.ProtoInt32(s2c.MarketJPFuture),
		ProgramStatus:  s2c.ProgramStatus,
	}, nil
}

// GetUserInfoRequest defines optional parameters for GetUserInfo.
type GetUserInfoRequest struct {
	Flag int32 // bitmask for selecting specific info fields (see UserInfoField in proto)
}

// GetUserInfoResponse represents the user information including user ID, nickname, avatar, and API level.
type GetUserInfoResponse struct {
	UserID                int64
	NickName              string
	AvatarUrl             string
	ApiLevel              string
	IsNeedAgreeDisclaimer bool
	ShQotRight            int32
	SzQotRight            int32
	Extra                 int32
	HkQotRight            int32
	UsQotRight            int32
	CnQotRight            int32
	SubQuota              int32
	HistoryKLQuota        int32
	HkOptionQotRight      int32
	HasUSOptionQotRight   bool
	HkFutureQotRight      int32
	UsFutureQotRight      int32
	UsOptionQotRight      int32
	WebKey                string
	WebJumpUrlHead        string
	UserAttribution       int32
	UpdateWhatsNew        string
	UpdateType            int32
	UsIndexQotRight       int32
	UsOtcQotRight         int32
	UsCMEFutureQotRight   int32
	UsCBOTFutureQotRight  int32
	UsNYMEXFutureQotRight int32
	UsCOMEXFutureQotRight int32
	UsCBOEFutureQotRight  int32
	SgFutureQotRight      int32
	JpFutureQotRight      int32
	IsAppNNOrMM           bool
}

// GetUserInfo retrieves the current user information including nickname, avatar, and API level.
// If req is nil, sends an empty C2S (backward compatible). Returns the user info or an error.
func GetUserInfo(ctx context.Context, c *futuapi.Client, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	c2s := &getuserinfo.C2S{}
	if req != nil && req.Flag != 0 {
		c2s.Flag = &req.Flag
	}

	pkt := &getuserinfo.Request{C2S: c2s}
	var rsp getuserinfo.Response

	if err := c.RequestContext(ctx, ProtoID_GetUserInfo, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUserInfo", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetUserInfo", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetUserInfoResponse{
		UserID:                util.ProtoInt64(s2c.UserID),
		NickName:              util.ProtoStr(s2c.NickName),
		AvatarUrl:             util.ProtoStr(s2c.AvatarUrl),
		ApiLevel:              util.ProtoStr(s2c.ApiLevel),
		IsNeedAgreeDisclaimer: util.ProtoBool(s2c.IsNeedAgreeDisclaimer),
		ShQotRight:            util.ProtoInt32(s2c.ShQotRight),
		SzQotRight:            util.ProtoInt32(s2c.SzQotRight),
		Extra:                 util.ProtoInt32(s2c.Extra),
		HkQotRight:            util.ProtoInt32(s2c.HkQotRight),
		UsQotRight:            util.ProtoInt32(s2c.UsQotRight),
		CnQotRight:            util.ProtoInt32(s2c.CnQotRight),
		SubQuota:              util.ProtoInt32(s2c.SubQuota),
		HistoryKLQuota:        util.ProtoInt32(s2c.HistoryKLQuota),
		HkOptionQotRight:      util.ProtoInt32(s2c.HkOptionQotRight),
		HasUSOptionQotRight:   util.ProtoBool(s2c.HasUSOptionQotRight),
		HkFutureQotRight:      util.ProtoInt32(s2c.HkFutureQotRight),
		UsFutureQotRight:      util.ProtoInt32(s2c.UsFutureQotRight),
		UsOptionQotRight:      util.ProtoInt32(s2c.UsOptionQotRight),
		WebKey:                util.ProtoStr(s2c.WebKey),
		WebJumpUrlHead:        util.ProtoStr(s2c.WebJumpUrlHead),
		UserAttribution:       util.ProtoInt32(s2c.UserAttribution),
		UpdateWhatsNew:        util.ProtoStr(s2c.UpdateWhatsNew),
		UpdateType:            util.ProtoInt32(s2c.UpdateType),
		UsIndexQotRight:       util.ProtoInt32(s2c.UsIndexQotRight),
		UsOtcQotRight:         util.ProtoInt32(s2c.UsOtcQotRight),
		UsCMEFutureQotRight:   util.ProtoInt32(s2c.UsCMEFutureQotRight),
		UsCBOTFutureQotRight:  util.ProtoInt32(s2c.UsCBOTFutureQotRight),
		UsNYMEXFutureQotRight: util.ProtoInt32(s2c.UsNYMEXFutureQotRight),
		UsCOMEXFutureQotRight: util.ProtoInt32(s2c.UsCOMEXFutureQotRight),
		UsCBOEFutureQotRight:  util.ProtoInt32(s2c.UsCBOEFutureQotRight),
		SgFutureQotRight:      util.ProtoInt32(s2c.SgFutureQotRight),
		JpFutureQotRight:      util.ProtoInt32(s2c.JpFutureQotRight),
		IsAppNNOrMM:           util.ProtoBool(s2c.IsAppNNOrMM),
	}, nil
}

// GetDelayStatisticsRequest defines optional parameters for GetDelayStatistics.
type GetDelayStatisticsRequest struct {
	TypeList     []int32
	QotPushStage int32
	SegmentList  []int32
}

// GetDelayStatisticsResponse represents delay statistics for quote push, request-reply, and order placement.
type GetDelayStatisticsResponse struct {
	QotPushStatisticsList    []*QotPushDelayStatistics
	ReqReplyStatisticsList    []*ReqReplyDelayStatistics
	PlaceOrderStatisticsList  []*PlaceOrderDelayStatistics
}

// QotPushDelayStatistics represents quote push delay statistics.
type QotPushDelayStatistics struct {
	QotPushType int32
	ItemList    []*DelayStatisticsItem
	DelayAvg    float32
	Count       int32
}

// DelayStatisticsItem represents a single delay statistics item.
type DelayStatisticsItem struct {
	Begin          int32
	End            int32
	Count          int32
	Proportion     float32
	CumulativeRatio float32
}

// ReqReplyDelayStatistics represents request-reply delay statistics.
type ReqReplyDelayStatistics struct {
	ProtoID       int32
	Count         int32
	TotalCostAvg  float32
	OpenDCostAvg  float32
	NetDelayAvg   float32
	IsLocalReply  bool
}

// PlaceOrderDelayStatistics represents order placement delay statistics.
type PlaceOrderDelayStatistics struct {
	OrderID    string
	TotalCost  float32
	OpenDCost  float32
	NetDelay   float32
	UpdateCost float32
}

// marshalGetDelayStatisticsRequestProto2 marshals the C2S message using proto2 wire format.
// This is a workaround for the proto2 vs proto3 wire format incompatibility.
// Proto2 uses non-packed encoding for repeated int32, while proto3 uses packed encoding.
// OpenD's C++ parser expects proto2 non-packed encoding.
func marshalGetDelayStatisticsRequestProto2(c2s *getdelaystatistics.C2S) ([]byte, error) {
	buf := make([]byte, 0, 64)

	for _, v := range c2s.TypeList {
		buf = append(buf, 8)
		buf = appendVarint(buf, uint64(v))
	}

	if c2s.QotPushStage != nil {
		buf = append(buf, 16)
		buf = appendVarint(buf, uint64(*c2s.QotPushStage))
	}

	for _, v := range c2s.SegmentList {
		buf = append(buf, 24)
		buf = appendVarint(buf, uint64(v))
	}

	// Wrap C2S in Request message with proto2 length-delimited encoding
	reqBuf := make([]byte, 0, len(buf)+10)
	reqBuf = append(reqBuf, 0x0A)
	reqBuf = appendVarint(reqBuf, uint64(len(buf)))
	reqBuf = append(reqBuf, buf...)

	return reqBuf, nil
}

// appendVarint appends a varint-encoded uint64 to the buffer.
func appendVarint(buf []byte, v uint64) []byte {
	for v >= 0x80 {
		buf = append(buf, byte(v)|0x80)
		v >>= 7
	}
	buf = append(buf, byte(v))
	return buf
}

// GetDelayStatistics retrieves performance delay statistics for quote pushes, request-reply, and order placements.
// Returns the delay statistics or an error if the request fails.
//
// Note: This function uses proto2 wire format for compatibility with OpenD's C++ protobuf parser.
func GetDelayStatistics(ctx context.Context, c *futuapi.Client, req *GetDelayStatisticsRequest) (*GetDelayStatisticsResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, fmt.Errorf("GetDelayStatistics: %w", err)
	}

	c2s := &getdelaystatistics.C2S{}
	if req != nil {
		if len(req.TypeList) > 0 {
			c2s.TypeList = req.TypeList
		}
		if req.QotPushStage != 0 {
			c2s.QotPushStage = &req.QotPushStage
		}
		if len(req.SegmentList) > 0 {
			c2s.SegmentList = req.SegmentList
		}
	}

	body, err := marshalGetDelayStatisticsRequestProto2(c2s)
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDelayStatistics", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetDelayStatistics", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	qotList := make([]*QotPushDelayStatistics, 0, len(s2c.QotPushStatisticsList))
	for _, item := range s2c.QotPushStatisticsList {
		if item == nil {
			continue
		}
		items := make([]*DelayStatisticsItem, 0, len(item.ItemList))
		for _, it := range item.ItemList {
			if it == nil {
				continue
			}
			items = append(items, &DelayStatisticsItem{
				Begin:           util.ProtoInt32(it.Begin),
				End:             util.ProtoInt32(it.End),
				Count:           util.ProtoInt32(it.Count),
				Proportion:      util.ProtoFloat32(it.Proportion),
				CumulativeRatio: util.ProtoFloat32(it.CumulativeRatio),
			})
		}
		qotList = append(qotList, &QotPushDelayStatistics{
			QotPushType: util.ProtoInt32(item.QotPushType),
			ItemList:    items,
			DelayAvg:    util.ProtoFloat32(item.DelayAvg),
			Count:       util.ProtoInt32(item.Count),
		})
	}

	reqReplyList := make([]*ReqReplyDelayStatistics, 0, len(s2c.ReqReplyStatisticsList))
	for _, item := range s2c.ReqReplyStatisticsList {
		if item == nil {
			continue
		}
		reqReplyList = append(reqReplyList, &ReqReplyDelayStatistics{
			ProtoID:      util.ProtoInt32(item.ProtoID),
			Count:        util.ProtoInt32(item.Count),
			TotalCostAvg: util.ProtoFloat32(item.TotalCostAvg),
			OpenDCostAvg: util.ProtoFloat32(item.OpenDCostAvg),
			NetDelayAvg:  util.ProtoFloat32(item.NetDelayAvg),
			IsLocalReply: util.ProtoBool(item.IsLocalReply),
		})
	}

	placeOrderList := make([]*PlaceOrderDelayStatistics, 0, len(s2c.PlaceOrderStatisticsList))
	for _, item := range s2c.PlaceOrderStatisticsList {
		if item == nil {
			continue
		}
		placeOrderList = append(placeOrderList, &PlaceOrderDelayStatistics{
			OrderID:    util.ProtoStr(item.OrderID),
			TotalCost:  util.ProtoFloat32(item.TotalCost),
			OpenDCost:  util.ProtoFloat32(item.OpenDCost),
			NetDelay:   util.ProtoFloat32(item.NetDelay),
			UpdateCost: util.ProtoFloat32(item.UpdateCost),
		})
	}

	return &GetDelayStatisticsResponse{
		QotPushStatisticsList:    qotList,
		ReqReplyStatisticsList:    reqReplyList,
		PlaceOrderStatisticsList:  placeOrderList,
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
	if req == nil {
		return fmt.Errorf("Verification: request is nil")
	}
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return wrapError("Verification", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return nil
}

// GetUsedQuotaResponse represents the quota usage information.
type GetUsedQuotaResponse struct {
	UsedSubQuota   int32 // 已使用订阅额度
	UsedKLineQuota int32 // 已使用历史K线额度
}

// GetUsedQuota retrieves the current quota usage for subscriptions and historical K-line requests.
// Returns the used quota information or an error if the request fails.
func GetUsedQuota(ctx context.Context, c *futuapi.Client) (*GetUsedQuotaResponse, error) {
	c2s := &usedquota.C2S{}
	pkt := &usedquota.Request{C2S: c2s}
	var rsp usedquota.Response

	if err := c.RequestContext(ctx, ProtoID_UsedQuota, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUsedQuota", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetUsedQuota", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetUsedQuotaResponse{
		UsedSubQuota:   util.ProtoInt32(s2c.UsedSubQuota),
		UsedKLineQuota: util.ProtoInt32(s2c.UsedKLineQuota),
	}, nil
}

var _ proto.Message = (*usedquota.Request)(nil)
var _ proto.Message = (*usedquota.Response)(nil)

// TestCmdRequest is the request to send a test command to OpenD for internal diagnostics.
type TestCmdRequest struct {
	Cmd    string
	Params string
}

// TestCmdResponse is the response containing the test command result.
type TestCmdResponse struct {
	Cmd    string
	Result string
}

// TestCmd sends a test command to OpenD for internal diagnostics.
// Returns the command result or an error if the request fails.
func TestCmd(ctx context.Context, c *futuapi.Client, req *TestCmdRequest) (*TestCmdResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("TestCmd: request is nil")
	}
	if req.Cmd == "" {
		return nil, fmt.Errorf("TestCmd: cmd is required")
	}

	c2s := &testcmd.C2S{
		Cmd: &req.Cmd,
	}
	if req.Params != "" {
		c2s.Params = &req.Params
	}

	pkt := &testcmd.Request{C2S: c2s}
	var rsp testcmd.Response

	if err := c.RequestContext(ctx, ProtoID_TestCmd, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("TestCmd", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("TestCmd", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &TestCmdResponse{
		Cmd:    util.ProtoStr(s2c.Cmd),
		Result: util.ProtoStr(s2c.Result),
	}, nil
}
