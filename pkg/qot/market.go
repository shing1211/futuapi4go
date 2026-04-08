package qot

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/common"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetmarketstate"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetownerplate"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetreference"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
)

const (
	ProtoID_GetOwnerPlate  = 2204
	ProtoID_GetReference   = 2205
	ProtoID_GetMarketState = 2208
	ProtoID_GetSubInfo     = 3002
)

type GetOwnerPlateRequest struct {
	SecurityList []*qotcommon.Security
}

type GetOwnerPlateResponse struct {
	OwnerPlateList []*qotgetownerplate.SecurityOwnerPlate
}

func GetOwnerPlate(c *futuapi.Client, req *GetOwnerPlateRequest) (*GetOwnerPlateResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetownerplate.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetownerplate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOwnerPlate, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetownerplate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOwnerPlate failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOwnerPlate: s2c is nil")
	}

	return &GetOwnerPlateResponse{
		OwnerPlateList: s2c.GetOwnerPlateList(),
	}, nil
}

type GetReferenceRequest struct {
	Security      *qotcommon.Security
	ReferenceType int32
}

type GetReferenceResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

func GetReference(c *futuapi.Client, req *GetReferenceRequest) (*GetReferenceResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetreference.C2S{
		Security:      req.Security,
		ReferenceType: &req.ReferenceType,
	}

	pkt := &qotgetreference.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetReference, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetreference.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetReference failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetReference: s2c is nil")
	}

	return &GetReferenceResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}

type MarketStateInfo struct {
	Security    *qotcommon.Security
	Name        string
	MarketState int32
}

type GetMarketStateRequest struct {
	SecurityList []*qotcommon.Security
}

type GetMarketStateResponse struct {
	MarketInfoList []*MarketStateInfo
}

func GetMarketState(c *futuapi.Client, req *GetMarketStateRequest) (*GetMarketStateResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetmarketstate.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetmarketstate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetMarketState, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetmarketstate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetMarketState failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetMarketState: s2c is nil")
	}

	result := &GetMarketStateResponse{
		MarketInfoList: make([]*MarketStateInfo, 0, len(s2c.GetMarketInfoList())),
	}

	for _, mi := range s2c.GetMarketInfoList() {
		result.MarketInfoList = append(result.MarketInfoList, &MarketStateInfo{
			Security:    mi.GetSecurity(),
			Name:        mi.GetName(),
			MarketState: mi.GetMarketState(),
		})
	}

	return result, nil
}

type GetSubInfoResponse struct {
	ConnSubInfoList []*qotcommon.ConnSubInfo
	TotalUsedQuota  int32
	RemainQuota     int32
}

func GetSubInfo(c *futuapi.Client) (*GetSubInfoResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetsubinfo.C2S{}

	pkt := &qotgetsubinfo.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetSubInfo, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetsubinfo.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetSubInfo failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetSubInfo: s2c is nil")
	}

	return &GetSubInfoResponse{
		ConnSubInfoList: s2c.GetConnSubInfoList(),
		TotalUsedQuota:  s2c.GetTotalUsedQuota(),
		RemainQuota:     s2c.GetRemainQuota(),
	}, nil
}
