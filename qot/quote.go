package qot

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"gitee.com/shing1211/futuapi4go/pb/common"
	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pb/qotgetbasicqot"
	"gitee.com/shing1211/futuapi4go/pb/qotgetbroker"
	"gitee.com/shing1211/futuapi4go/pb/qotgetcapitaldistribution"
	"gitee.com/shing1211/futuapi4go/pb/qotgetcapitalflow"
	"gitee.com/shing1211/futuapi4go/pb/qotgetkl"
	"gitee.com/shing1211/futuapi4go/pb/qotgetorderbook"
	"gitee.com/shing1211/futuapi4go/pb/qotgetplatesecurity"
	"gitee.com/shing1211/futuapi4go/pb/qotgetplateset"
	"gitee.com/shing1211/futuapi4go/pb/qotgetpricereminder"
	"gitee.com/shing1211/futuapi4go/pb/qotgetrt"
	"gitee.com/shing1211/futuapi4go/pb/qotgetsecuritysnapshot"
	"gitee.com/shing1211/futuapi4go/pb/qotgetstaticinfo"
	"gitee.com/shing1211/futuapi4go/pb/qotgetticker"
	"gitee.com/shing1211/futuapi4go/pb/qotgettradedate"
	"gitee.com/shing1211/futuapi4go/pb/qotgetusersecurity"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesthistorykl"
	"gitee.com/shing1211/futuapi4go/pb/qotsub"
)

const (
	ProtoID_GetBasicQot            = 2101
	ProtoID_GetKL                  = 2102
	ProtoID_GetHistoryKL           = 2103
	ProtoID_RequestHistoryKL       = 2104
	ProtoID_GetOrderBook           = 2106
	ProtoID_GetTicker              = 2107
	ProtoID_GetRT                  = 2108
	ProtoID_GetMarketSnapshot      = 2109
	ProtoID_GetSecuritySnapshot    = 2110
	ProtoID_GetBroker              = 2111
	ProtoID_GetStaticInfo          = 2201
	ProtoID_GetPlateSet            = 2202
	ProtoID_GetPlateSecurity       = 2203
	ProtoID_GetCapitalFlow         = 2301
	ProtoID_GetCapitalDistribution = 2302
	ProtoID_GetUserSecurity        = 2401
	ProtoID_GetPriceReminder       = 2404
	ProtoID_GetTradeDate           = 2206
	ProtoID_Subscribe              = 3001
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

type RequestHistoryKLRequest struct {
	RehabType    int32
	KlType       int32
	Security     *qotcommon.Security
	BeginTime    string
	EndTime      string
	MaxAckKLNum  int32
	NextReqKey   []byte
	ExtendedTime bool
}

type RequestHistoryKLResponse struct {
	Security   *qotcommon.Security
	Name       string
	KLList     []*qotcommon.KLine
	NextReqKey []byte
}

func RequestHistoryKL(c *futuapi.Client, req *RequestHistoryKLRequest) (*RequestHistoryKLResponse, error) {
	c2s := &qotrequesthistorykl.C2S{
		RehabType:    &req.RehabType,
		KlType:       &req.KlType,
		Security:     req.Security,
		BeginTime:    &req.BeginTime,
		EndTime:      &req.EndTime,
		MaxAckKLNum:  &req.MaxAckKLNum,
		NextReqKey:   req.NextReqKey,
		ExtendedTime: &req.ExtendedTime,
	}

	pkt := &qotrequesthistorykl.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_RequestHistoryKL, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotrequesthistorykl.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("RequestHistoryKL failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestHistoryKL: s2c is nil")
	}

	return &RequestHistoryKLResponse{
		Security:   s2c.GetSecurity(),
		Name:       s2c.GetName(),
		KLList:     s2c.GetKlList(),
		NextReqKey: s2c.GetNextReqKey(),
	}, nil
}

type GetSecuritySnapshotRequest struct {
	SecurityList []*qotcommon.Security
}

type GetSecuritySnapshotResponse struct {
	SnapshotList []*qotgetsecuritysnapshot.Snapshot
}

func GetSecuritySnapshot(c *futuapi.Client, req *GetSecuritySnapshotRequest) (*GetSecuritySnapshotResponse, error) {
	c2s := &qotgetsecuritysnapshot.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetsecuritysnapshot.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetSecuritySnapshot, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetsecuritysnapshot.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetSecuritySnapshot failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetSecuritySnapshot: s2c is nil")
	}

	return &GetSecuritySnapshotResponse{
		SnapshotList: s2c.GetSnapshotList(),
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

type GetCapitalFlowRequest struct {
	Security   *qotcommon.Security
	PeriodType int32
	BeginTime  string
	EndTime    string
}

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

type GetCapitalFlowResponse struct {
	FlowItemList       []*CapitalFlowItem
	LastValidTime      string
	LastValidTimestamp float64
}

func GetCapitalFlow(c *futuapi.Client, req *GetCapitalFlowRequest) (*GetCapitalFlowResponse, error) {
	c2s := &qotgetcapitalflow.C2S{
		Security:   req.Security,
		PeriodType: &req.PeriodType,
		BeginTime:  &req.BeginTime,
		EndTime:    &req.EndTime,
	}

	pkt := &qotgetcapitalflow.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetCapitalFlow, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetcapitalflow.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetCapitalFlow failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetCapitalFlow: s2c is nil")
	}

	result := &GetCapitalFlowResponse{
		FlowItemList:       make([]*CapitalFlowItem, 0, len(s2c.GetFlowItemList())),
		LastValidTime:      s2c.GetLastValidTime(),
		LastValidTimestamp: s2c.GetLastValidTimestamp(),
	}

	for _, f := range s2c.GetFlowItemList() {
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

type GetCapitalDistributionResponse struct {
	CapitalDistribution *CapitalDistribution
}

func GetCapitalDistribution(c *futuapi.Client, security *qotcommon.Security) (*GetCapitalDistributionResponse, error) {
	c2s := &qotgetcapitaldistribution.C2S{
		Security: security,
	}

	pkt := &qotgetcapitaldistribution.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetCapitalDistribution, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetcapitaldistribution.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetCapitalDistribution failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetCapitalDistribution: s2c is nil")
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

type GetUserSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

func GetUserSecurity(c *futuapi.Client, groupName string) (*GetUserSecurityResponse, error) {
	c2s := &qotgetusersecurity.C2S{
		GroupName: &groupName,
	}

	pkt := &qotgetusersecurity.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetUserSecurity, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetusersecurity.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetUserSecurity failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetUserSecurity: s2c is nil")
	}

	return &GetUserSecurityResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}

type PriceReminderItemInfo struct {
	Key      int64
	Type     int32
	Value    float64
	Note     string
	Freq     int32
	IsEnable bool
}

type PriceReminderInfo struct {
	Security *qotcommon.Security
	Name     string
	ItemList []*PriceReminderItemInfo
}

type GetPriceReminderResponse struct {
	PriceReminderList []*PriceReminderInfo
}

func GetPriceReminder(c *futuapi.Client, security *qotcommon.Security, market int32) (*GetPriceReminderResponse, error) {
	c2s := &qotgetpricereminder.C2S{
		Security: security,
		Market:   &market,
	}

	pkt := &qotgetpricereminder.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetPriceReminder, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetpricereminder.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetPriceReminder failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPriceReminder: s2c is nil")
	}

	result := &GetPriceReminderResponse{
		PriceReminderList: make([]*PriceReminderInfo, 0, len(s2c.GetPriceReminderList())),
	}

	for _, pr := range s2c.GetPriceReminderList() {
		info := &PriceReminderInfo{
			Security: pr.GetSecurity(),
			Name:     pr.GetName(),
			ItemList: make([]*PriceReminderItemInfo, 0, len(pr.GetItemList())),
		}
		for _, item := range pr.GetItemList() {
			info.ItemList = append(info.ItemList, &PriceReminderItemInfo{
				Key:      item.GetKey(),
				Type:     item.GetType(),
				Value:    item.GetValue(),
				Note:     item.GetNote(),
				Freq:     item.GetFreq(),
				IsEnable: item.GetIsEnable(),
			})
		}
		result.PriceReminderList = append(result.PriceReminderList, info)
	}

	return result, nil
}
