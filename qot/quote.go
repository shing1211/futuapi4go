package qot

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"github.com/futuopen/ftapi4go/pb/common"
	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"github.com/futuopen/ftapi4go/pb/qotgetbasicqot"
	"github.com/futuopen/ftapi4go/pb/qotgetbroker"
	"github.com/futuopen/ftapi4go/pb/qotgetkl"
	"github.com/futuopen/ftapi4go/pb/qotgetorderbook"
	"github.com/futuopen/ftapi4go/pb/qotgetplatesecurity"
	"github.com/futuopen/ftapi4go/pb/qotgetplateset"
	"github.com/futuopen/ftapi4go/pb/qotgetrt"
	"github.com/futuopen/ftapi4go/pb/qotgetstaticinfo"
	"github.com/futuopen/ftapi4go/pb/qotgetticker"
	"github.com/futuopen/ftapi4go/pb/qotgettradedate"
	"github.com/futuopen/ftapi4go/pb/qotsub"
)

const (
	ProtoID_GetBasicQot       = 2101
	ProtoID_GetKL             = 2102
	ProtoID_GetHistoryKL      = 2103
	ProtoID_GetOrderBook      = 2106
	ProtoID_GetTicker         = 2107
	ProtoID_GetRT             = 2108
	ProtoID_GetMarketSnapshot = 2109
	ProtoID_GetBroker         = 2111
	ProtoID_GetStaticInfo     = 2201
	ProtoID_GetPlateSet       = 2202
	ProtoID_GetPlateSecurity  = 2203
	ProtoID_GetTradeDate      = 2206
	ProtoID_Subscribe         = 3001
)

type BasicQot struct {
	Security  *qotcommon.Security
	Name      string
	CurPrice  float64
	OpenPrice float64
	HighPrice float64
	LowPrice  float64
	Volume    int64
	Turnover  float64
}

func GetBasicQot(c *futuapi.Client, securityList []*qotcommon.Security) ([]*BasicQot, error) {
	c2s := &qotgetbasicqot.C2S{
		SecurityList: securityList,
	}

	pkt := &qotgetbasicqot.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetBasicQot, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetbasicqot.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetBasicQot failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetBasicQot: s2c is nil")
	}

	result := make([]*BasicQot, 0, len(s2c.GetBasicQotList()))
	for _, bq := range s2c.GetBasicQotList() {
		result = append(result, &BasicQot{
			Security:  bq.GetSecurity(),
			Name:      bq.GetName(),
			CurPrice:  bq.GetCurPrice(),
			OpenPrice: bq.GetOpenPrice(),
			HighPrice: bq.GetHighPrice(),
			LowPrice:  bq.GetLowPrice(),
			Volume:    bq.GetVolume(),
			Turnover:  bq.GetTurnover(),
		})
	}

	return result, nil
}

type KLine struct {
	Time           string
	IsBlank        bool
	HighPrice      float64
	OpenPrice      float64
	LowPrice       float64
	ClosePrice     float64
	LastClosePrice float64
	Volume         int64
	Turnover       float64
	ChangeRate     float64
	Timestamp      float64
}

type GetKLRequest struct {
	Security  *qotcommon.Security
	RehabType int32
	KLType    int32
	ReqNum    int32
}

type GetKLResponse struct {
	Security *qotcommon.Security
	Name     string
	KLList   []*KLine
}

func GetKL(c *futuapi.Client, req *GetKLRequest) (*GetKLResponse, error) {
	c2s := &qotgetkl.C2S{
		Security:  req.Security,
		RehabType: &req.RehabType,
		KlType:    &req.KLType,
		ReqNum:    &req.ReqNum,
	}

	pkt := &qotgetkl.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetKL, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetkl.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetKL failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetKL: s2c is nil")
	}

	result := &GetKLResponse{
		Security: s2c.GetSecurity(),
		Name:     s2c.GetName(),
		KLList:   make([]*KLine, 0, len(s2c.GetKlList())),
	}

	for _, kl := range s2c.GetKlList() {
		result.KLList = append(result.KLList, &KLine{
			Time:           kl.GetTime(),
			IsBlank:        kl.GetIsBlank(),
			HighPrice:      kl.GetHighPrice(),
			OpenPrice:      kl.GetOpenPrice(),
			LowPrice:       kl.GetLowPrice(),
			ClosePrice:     kl.GetClosePrice(),
			LastClosePrice: kl.GetLastClosePrice(),
			Volume:         kl.GetVolume(),
			Turnover:       kl.GetTurnover(),
			ChangeRate:     kl.GetChangeRate(),
			Timestamp:      kl.GetTimestamp(),
		})
	}

	return result, nil
}

type OrderBook struct {
	Price      float64
	Volume     int64
	OrderCount int32
	DetailList []*OrderBookDetail
}

type OrderBookDetail struct {
	Price   float64
	Volume  int64
	OrderID int64
}

type GetOrderBookRequest struct {
	Security *qotcommon.Security
	Num      int32
}

type GetOrderBookResponse struct {
	Security                *qotcommon.Security
	Name                    string
	OrderBookAskList        []*OrderBook
	OrderBookBidList        []*OrderBook
	SvrRecvTimeBid          string
	SvrRecvTimeBidTimestamp float64
	SvrRecvTimeAsk          string
	SvrRecvTimeAskTimestamp float64
}

func GetOrderBook(c *futuapi.Client, req *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	c2s := &qotgetorderbook.C2S{
		Security: req.Security,
		Num:      &req.Num,
	}

	pkt := &qotgetorderbook.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOrderBook, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetorderbook.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOrderBook failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderBook: s2c is nil")
	}

	result := &GetOrderBookResponse{
		Security:                s2c.GetSecurity(),
		Name:                    s2c.GetName(),
		SvrRecvTimeBid:          s2c.GetSvrRecvTimeBid(),
		SvrRecvTimeBidTimestamp: s2c.GetSvrRecvTimeBidTimestamp(),
		SvrRecvTimeAsk:          s2c.GetSvrRecvTimeAsk(),
		SvrRecvTimeAskTimestamp: s2c.GetSvrRecvTimeAskTimestamp(),
	}

	for _, ob := range s2c.GetOrderBookAskList() {
		result.OrderBookAskList = append(result.OrderBookAskList, &OrderBook{
			Price:      ob.GetPrice(),
			Volume:     ob.GetVolume(),
			OrderCount: ob.GetOrederCount(),
		})
	}

	for _, ob := range s2c.GetOrderBookBidList() {
		result.OrderBookBidList = append(result.OrderBookBidList, &OrderBook{
			Price:      ob.GetPrice(),
			Volume:     ob.GetVolume(),
			OrderCount: ob.GetOrederCount(),
		})
	}

	return result, nil
}

type Ticker struct {
	Time      string
	Sequence  int64
	Dir       int32
	Price     float64
	Volume    int64
	Turnover  float64
	RecvTime  float64
	Type      int32
	TypeSign  int32
	Timestamp float64
}

type GetTickerRequest struct {
	Security *qotcommon.Security
	Num      int32
}

type GetTickerResponse struct {
	Security   *qotcommon.Security
	Name       string
	TickerList []*Ticker
}

func GetTicker(c *futuapi.Client, req *GetTickerRequest) (*GetTickerResponse, error) {
	maxRetNum := req.Num
	c2s := &qotgetticker.C2S{
		Security:  req.Security,
		MaxRetNum: &maxRetNum,
	}

	pkt := &qotgetticker.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetTicker, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetticker.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetTicker failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetTicker: s2c is nil")
	}

	result := &GetTickerResponse{
		Security:   s2c.GetSecurity(),
		Name:       s2c.GetName(),
		TickerList: make([]*Ticker, 0, len(s2c.GetTickerList())),
	}

	for _, t := range s2c.GetTickerList() {
		result.TickerList = append(result.TickerList, &Ticker{
			Time:      t.GetTime(),
			Sequence:  t.GetSequence(),
			Dir:       t.GetDir(),
			Price:     t.GetPrice(),
			Volume:    t.GetVolume(),
			Turnover:  t.GetTurnover(),
			RecvTime:  t.GetRecvTime(),
			Type:      t.GetType(),
			TypeSign:  t.GetTypeSign(),
			Timestamp: t.GetTimestamp(),
		})
	}

	return result, nil
}

type RT struct {
	Time     string
	Price    float64
	Volume   int64
	Turnover float64
	AvgPrice float64
}

type GetRTRequest struct {
	Security *qotcommon.Security
}

type GetRTResponse struct {
	Security *qotcommon.Security
	Name     string
	RTList   []*RT
}

func GetRT(c *futuapi.Client, req *GetRTRequest) (*GetRTResponse, error) {
	c2s := &qotgetrt.C2S{
		Security: req.Security,
	}

	pkt := &qotgetrt.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetRT, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetrt.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetRT failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetRT: s2c is nil")
	}

	result := &GetRTResponse{
		Security: s2c.GetSecurity(),
		Name:     s2c.GetName(),
		RTList:   make([]*RT, 0, len(s2c.GetRtList())),
	}

	for _, rt := range s2c.GetRtList() {
		result.RTList = append(result.RTList, &RT{
			Time:     rt.GetTime(),
			Price:    rt.GetPrice(),
			Volume:   rt.GetVolume(),
			Turnover: rt.GetTurnover(),
			AvgPrice: rt.GetAvgPrice(),
		})
	}

	return result, nil
}

type Broker struct {
	ID     int64
	Name   string
	Pos    int32
	Volume int64
}

type GetBrokerRequest struct {
	Security *qotcommon.Security
	Num      int32
}

type GetBrokerResponse struct {
	Security      *qotcommon.Security
	Name          string
	AskBrokerList []*Broker
	BidBrokerList []*Broker
}

func GetBroker(c *futuapi.Client, req *GetBrokerRequest) (*GetBrokerResponse, error) {
	c2s := &qotgetbroker.C2S{
		Security: req.Security,
	}

	pkt := &qotgetbroker.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetBroker, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetbroker.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetBroker failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetBroker: s2c is nil")
	}

	result := &GetBrokerResponse{
		Security:      s2c.GetSecurity(),
		Name:          s2c.GetName(),
		AskBrokerList: make([]*Broker, 0, len(s2c.GetBrokerAskList())),
		BidBrokerList: make([]*Broker, 0, len(s2c.GetBrokerBidList())),
	}

	for _, b := range s2c.GetBrokerAskList() {
		result.AskBrokerList = append(result.AskBrokerList, &Broker{
			ID:     b.GetId(),
			Name:   b.GetName(),
			Pos:    b.GetPos(),
			Volume: b.GetVolume(),
		})
	}

	for _, b := range s2c.GetBrokerBidList() {
		result.BidBrokerList = append(result.BidBrokerList, &Broker{
			ID:     b.GetId(),
			Name:   b.GetName(),
			Pos:    b.GetPos(),
			Volume: b.GetVolume(),
		})
	}

	return result, nil
}

type GetStaticInfoRequest struct {
	Market       int32
	SecType      int32
	SecurityList []*qotcommon.Security
}

type GetStaticInfoResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

func GetStaticInfo(c *futuapi.Client, req *GetStaticInfoRequest) (*GetStaticInfoResponse, error) {
	c2s := &qotgetstaticinfo.C2S{
		Market:       &req.Market,
		SecType:      &req.SecType,
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetstaticinfo.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetStaticInfo, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetstaticinfo.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetStaticInfo failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetStaticInfo: s2c is nil")
	}

	return &GetStaticInfoResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}

type Plate struct {
	Plate *qotcommon.Security
	Name  string
}

type GetPlateSetRequest struct {
	Market int32
}

type GetPlateSetResponse struct {
	PlateSetList []*Plate
}

func GetPlateSet(c *futuapi.Client, req *GetPlateSetRequest) (*GetPlateSetResponse, error) {
	c2s := &qotgetplateset.C2S{
		Market: &req.Market,
	}

	pkt := &qotgetplateset.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetPlateSet, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetplateset.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetPlateSet failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPlateSet: s2c is nil")
	}

	result := &GetPlateSetResponse{
		PlateSetList: make([]*Plate, 0, len(s2c.GetPlateInfoList())),
	}

	for _, p := range s2c.GetPlateInfoList() {
		result.PlateSetList = append(result.PlateSetList, &Plate{
			Plate: p.GetPlate(),
			Name:  p.GetName(),
		})
	}

	return result, nil
}

type GetPlateSecurityRequest struct {
	Plate *qotcommon.Security
}

type GetPlateSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

func GetPlateSecurity(c *futuapi.Client, req *GetPlateSecurityRequest) (*GetPlateSecurityResponse, error) {
	c2s := &qotgetplatesecurity.C2S{
		Plate: req.Plate,
	}

	pkt := &qotgetplatesecurity.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetPlateSecurity, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetplatesecurity.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetPlateSecurity failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPlateSecurity: s2c is nil")
	}

	return &GetPlateSecurityResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}

type GetTradeDateRequest struct {
	Market    int32
	BeginTime string
	EndTime   string
}

type GetTradeDateResponse struct {
	TradeDateList []*qotgettradedate.TradeDate
}

func GetTradeDate(c *futuapi.Client, req *GetTradeDateRequest) (*GetTradeDateResponse, error) {
	c2s := &qotgettradedate.C2S{
		Market:    &req.Market,
		BeginTime: &req.BeginTime,
		EndTime:   &req.EndTime,
	}

	pkt := &qotgettradedate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetTradeDate, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgettradedate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetTradeDate failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetTradeDate: s2c is nil")
	}

	return &GetTradeDateResponse{
		TradeDateList: s2c.GetTradeDateList(),
	}, nil
}

type SubType int32

const (
	SubType_Basic     SubType = 1
	SubType_OrderBook SubType = 2
	SubType_Ticker    SubType = 3
	SubType_KL        SubType = 4
	SubType_RT        SubType = 5
	SubType_Broker    SubType = 6
)

type SubscribeRequest struct {
	SecurityList         []*qotcommon.Security
	SubTypeList          []SubType
	IsSubOrUnSub         bool
	IsRegOrUnRegPush     bool
	RegPushRehabTypeList []int32
	IsFirstPush          bool
	IsUnsubAll           bool
}

type SubscribeResponse struct {
	RetType int32
	RetMsg  string
}

func Subscribe(c *futuapi.Client, req *SubscribeRequest) (*SubscribeResponse, error) {
	subTypeList := make([]int32, len(req.SubTypeList))
	for i, st := range req.SubTypeList {
		subTypeList[i] = int32(st)
	}

	c2s := &qotsub.C2S{
		SecurityList:         req.SecurityList,
		SubTypeList:          subTypeList,
		IsSubOrUnSub:         &req.IsSubOrUnSub,
		IsRegOrUnRegPush:     &req.IsRegOrUnRegPush,
		RegPushRehabTypeList: req.RegPushRehabTypeList,
		IsFirstPush:          &req.IsFirstPush,
		IsUnsubAll:           &req.IsUnsubAll,
	}

	pkt := &qotsub.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_Subscribe, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotsub.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	return &SubscribeResponse{
		RetType: rsp.GetRetType(),
		RetMsg:  rsp.GetRetMsg(),
	}, nil
}
