package sys

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/common"
	"gitee.com/shing1211/futuapi4go/pkg/pb/getdelaystatistics"
	"gitee.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"gitee.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"gitee.com/shing1211/futuapi4go/pkg/pb/verification"
)

const (
	ProtoID_GetGlobalState     = 1004
	ProtoID_GetUserInfo        = 1005
	ProtoID_GetDelayStatistics = 1006
	ProtoID_Verification       = 8001
)

type GetGlobalStateResponse struct {
	ConnID        uint64
	ServerVer     int32
	ServerBuildNo int32
	Time          int64
	LocalTime     float64
	QotLogined    bool
	TrdLogined    bool
	QotSvrIpAddr  string
	TrdSvrIpAddr  string
	MarketHK      int32
	MarketUS      int32
	MarketSH      int32
	MarketSZ      int32
}

func GetGlobalState(c *futuapi.Client) (*GetGlobalStateResponse, error) {
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

	pktResp, err := c.Conn().ReadPacket()
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
		ConnID:        s2c.GetConnID(),
		ServerVer:     s2c.GetServerVer(),
		ServerBuildNo: s2c.GetServerBuildNo(),
		Time:          s2c.GetTime(),
		LocalTime:     s2c.GetLocalTime(),
		QotLogined:    s2c.GetQotLogined(),
		TrdLogined:    s2c.GetTrdLogined(),
		QotSvrIpAddr:  s2c.GetQotSvrIpAddr(),
		TrdSvrIpAddr:  s2c.GetTrdSvrIpAddr(),
		MarketHK:      s2c.GetMarketHK(),
		MarketUS:      s2c.GetMarketUS(),
		MarketSH:      s2c.GetMarketSH(),
		MarketSZ:      s2c.GetMarketSZ(),
	}, nil
}

type GetUserInfoResponse struct {
	UserID                int64
	NickName              string
	AvatarUrl             string
	ApiLevel              string
	IsNeedAgreeDisclaimer bool
}

func GetUserInfo(c *futuapi.Client) (*GetUserInfoResponse, error) {
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

	pktResp, err := c.Conn().ReadPacket()
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

type GetDelayStatisticsResponse struct {
	QotPushStatisticsList    []*getdelaystatistics.DelayStatistics
	ReqReplyStatisticsList   []*getdelaystatistics.ReqReplyStatisticsItem
	PlaceOrderStatisticsList []*getdelaystatistics.PlaceOrderStatisticsItem
}

func GetDelayStatistics(c *futuapi.Client) (*GetDelayStatisticsResponse, error) {
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

	pktResp, err := c.Conn().ReadPacket()
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

type VerificationRequest struct {
	Type verification.VerificationType
	Op   verification.VerificationOp
	Code string
}

func Verification(c *futuapi.Client, req *VerificationRequest) error {
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

	pktResp, err := c.Conn().ReadPacket()
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
