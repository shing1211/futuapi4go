package qot

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"gitee.com/shing1211/futuapi4go/pb/common"
	"gitee.com/shing1211/futuapi4go/pb/getoptionexpirationdate"
	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pb/qotgetbasicqot"
	"gitee.com/shing1211/futuapi4go/pb/qotgetbroker"
	"gitee.com/shing1211/futuapi4go/pb/qotgetcapitaldistribution"
	"gitee.com/shing1211/futuapi4go/pb/qotgetcapitalflow"
	"gitee.com/shing1211/futuapi4go/pb/qotgetkl"
	"gitee.com/shing1211/futuapi4go/pb/qotgetoptionchain"
	"gitee.com/shing1211/futuapi4go/pb/qotgetorderbook"
	"gitee.com/shing1211/futuapi4go/pb/qotgetplatesecurity"
	"gitee.com/shing1211/futuapi4go/pb/qotgetplateset"
	"gitee.com/shing1211/futuapi4go/pb/qotgetpricereminder"
	"gitee.com/shing1211/futuapi4go/pb/qotgetrt"
	"gitee.com/shing1211/futuapi4go/pb/qotgetsecuritysnapshot"
	"gitee.com/shing1211/futuapi4go/pb/qotgetstaticinfo"
	"gitee.com/shing1211/futuapi4go/pb/qotgetsuspend"
	"gitee.com/shing1211/futuapi4go/pb/qotgetticker"
	"gitee.com/shing1211/futuapi4go/pb/qotgettradedate"
	"gitee.com/shing1211/futuapi4go/pb/qotgetusersecurity"
	"gitee.com/shing1211/futuapi4go/pb/qotgetwarrant"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesthistorykl"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesttradedate"
	"gitee.com/shing1211/futuapi4go/pb/qotstockfilter"
	"gitee.com/shing1211/futuapi4go/pb/qotsub"
)

const (
	ProtoID_GetBasicQot             = 2101
	ProtoID_GetKL                   = 2102
	ProtoID_GetHistoryKL            = 2103
	ProtoID_RequestHistoryKL        = 2104
	ProtoID_GetOrderBook            = 2106
	ProtoID_GetTicker               = 2107
	ProtoID_GetRT                   = 2108
	ProtoID_GetMarketSnapshot       = 2109
	ProtoID_GetSecuritySnapshot     = 2110
	ProtoID_GetBroker               = 2111
	ProtoID_GetStaticInfo           = 2201
	ProtoID_GetPlateSet             = 2202
	ProtoID_GetPlateSecurity        = 2203
	ProtoID_GetSuspend              = 2209
	ProtoID_GetCapitalFlow          = 2301
	ProtoID_GetCapitalDistribution  = 2302
	ProtoID_StockFilter             = 2303
	ProtoID_GetOptionChain          = 2304
	ProtoID_GetOptionExpirationDate = 2305
	ProtoID_GetWarrant              = 2306
	ProtoID_GetUserSecurity         = 2401
	ProtoID_GetPriceReminder        = 2404
	ProtoID_GetTradeDate            = 2206
	ProtoID_RequestTradeDate        = 2207
	ProtoID_Subscribe               = 3001
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

type RequestTradeDateRequest struct {
	Market    int32
	BeginTime string
	EndTime   string
	Security  *qotcommon.Security
}

type RequestTradeDateResponse struct {
	TradeDateList []*qotrequesttradedate.TradeDate
}

func RequestTradeDate(c *futuapi.Client, req *RequestTradeDateRequest) (*RequestTradeDateResponse, error) {
	c2s := &qotrequesttradedate.C2S{
		Market:    &req.Market,
		BeginTime: &req.BeginTime,
		EndTime:   &req.EndTime,
		Security:  req.Security,
	}

	pkt := &qotrequesttradedate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_RequestTradeDate, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotrequesttradedate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("RequestTradeDate failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestTradeDate: s2c is nil")
	}

	return &RequestTradeDateResponse{
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

type GetOptionExpirationDateRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
}

type OptionExpirationDateInfo struct {
	StrikeTime               string
	StrikeTimestamp          float64
	OptionExpiryDateDistance int32
	Cycle                    int32
}

type GetOptionExpirationDateResponse struct {
	DateList []*OptionExpirationDateInfo
}

func GetOptionExpirationDate(c *futuapi.Client, req *GetOptionExpirationDateRequest) (*GetOptionExpirationDateResponse, error) {
	c2s := &getoptionexpirationdate.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
	}

	pkt := &getoptionexpirationdate.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOptionExpirationDate, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp getoptionexpirationdate.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOptionExpirationDate failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOptionExpirationDate: s2c is nil")
	}

	result := &GetOptionExpirationDateResponse{
		DateList: make([]*OptionExpirationDateInfo, 0, len(s2c.GetDateList())),
	}

	for _, d := range s2c.GetDateList() {
		result.DateList = append(result.DateList, &OptionExpirationDateInfo{
			StrikeTime:               d.GetStrikeTime(),
			StrikeTimestamp:          d.GetStrikeTimestamp(),
			OptionExpiryDateDistance: d.GetOptionExpiryDateDistance(),
			Cycle:                    d.GetCycle(),
		})
	}

	return result, nil
}

type GetOptionChainRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
	Type            int32
	Condition       int32
	BeginTime       string
	EndTime         string
	DataFilter      *qotgetoptionchain.DataFilter
}

type OptionItem struct {
	Call *qotcommon.SecurityStaticInfo
	Put  *qotcommon.SecurityStaticInfo
}

type OptionChain struct {
	StrikeTime      string
	StrikeTimestamp float64
	Option          []*OptionItem
}

type GetOptionChainResponse struct {
	OptionChain []*OptionChain
}

func GetOptionChain(c *futuapi.Client, req *GetOptionChainRequest) (*GetOptionChainResponse, error) {
	c2s := &qotgetoptionchain.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
		Type:            &req.Type,
		Condition:       &req.Condition,
		BeginTime:       &req.BeginTime,
		EndTime:         &req.EndTime,
		DataFilter:      req.DataFilter,
	}

	pkt := &qotgetoptionchain.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOptionChain, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetoptionchain.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOptionChain failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOptionChain: s2c is nil")
	}

	result := &GetOptionChainResponse{
		OptionChain: make([]*OptionChain, 0, len(s2c.GetOptionChain())),
	}

	for _, oc := range s2c.GetOptionChain() {
		chain := &OptionChain{
			StrikeTime:      oc.GetStrikeTime(),
			StrikeTimestamp: oc.GetStrikeTimestamp(),
			Option:          make([]*OptionItem, 0, len(oc.GetOption())),
		}
		for _, opt := range oc.GetOption() {
			chain.Option = append(chain.Option, &OptionItem{
				Call: opt.GetCall(),
				Put:  opt.GetPut(),
			})
		}
		result.OptionChain = append(result.OptionChain, chain)
	}

	return result, nil
}

type StockFilterRequest struct {
	Begin                     int32
	Num                       int32
	Market                    int32
	Plate                     *qotcommon.Security
	BaseFilterList            []*qotstockfilter.BaseFilter
	AccumulateFilterList      []*qotstockfilter.AccumulateFilter
	FinancialFilterList       []*qotstockfilter.FinancialFilter
	PatternFilterList         []*qotstockfilter.PatternFilter
	CustomIndicatorFilterList []*qotstockfilter.CustomIndicatorFilter
}

type StockFilterData struct {
	Security                *qotcommon.Security
	Name                    string
	BaseDataList            []*qotstockfilter.BaseData
	AccumulateDataList      []*qotstockfilter.AccumulateData
	FinancialDataList       []*qotstockfilter.FinancialData
	CustomIndicatorDataList []*qotstockfilter.CustomIndicatorData
}

type StockFilterResponse struct {
	LastPage bool
	AllCount int32
	DataList []*StockFilterData
}

func StockFilter(c *futuapi.Client, req *StockFilterRequest) (*StockFilterResponse, error) {
	c2s := &qotstockfilter.C2S{
		Begin:                     &req.Begin,
		Num:                       &req.Num,
		Market:                    &req.Market,
		Plate:                     req.Plate,
		BaseFilterList:            req.BaseFilterList,
		AccumulateFilterList:      req.AccumulateFilterList,
		FinancialFilterList:       req.FinancialFilterList,
		PatternFilterList:         req.PatternFilterList,
		CustomIndicatorFilterList: req.CustomIndicatorFilterList,
	}

	pkt := &qotstockfilter.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_StockFilter, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotstockfilter.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("StockFilter failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("StockFilter: s2c is nil")
	}

	result := &StockFilterResponse{
		LastPage: s2c.GetLastPage(),
		AllCount: s2c.GetAllCount(),
		DataList: make([]*StockFilterData, 0, len(s2c.GetDataList())),
	}

	for _, d := range s2c.GetDataList() {
		result.DataList = append(result.DataList, &StockFilterData{
			Security:                d.GetSecurity(),
			Name:                    d.GetName(),
			BaseDataList:            d.GetBaseDataList(),
			AccumulateDataList:      d.GetAccumulateDataList(),
			FinancialDataList:       d.GetFinancialDataList(),
			CustomIndicatorDataList: d.GetCustomIndicatorDataList(),
		})
	}

	return result, nil
}

type GetWarrantRequest struct {
	Begin                 int32
	Num                   int32
	SortField             int32
	Ascend                bool
	Owner                 *qotcommon.Security
	TypeList              []int32
	IssuerList            []int32
	MaturityTimeMin       string
	MaturityTimeMax       string
	IpoPeriod             int32
	PriceType             int32
	Status                int32
	CurPriceMin           float64
	CurPriceMax           float64
	StrikePriceMin        float64
	StrikePriceMax        float64
	StreetMin             float64
	StreetMax             float64
	ConversionMin         float64
	ConversionMax         float64
	VolMin                uint64
	VolMax                uint64
	PremiumMin            float64
	PremiumMax            float64
	LeverageRatioMin      float64
	LeverageRatioMax      float64
	DeltaMin              float64
	DeltaMax              float64
	ImpliedMin            float64
	ImpliedMax            float64
	RecoveryPriceMin      float64
	RecoveryPriceMax      float64
	PriceRecoveryRatioMin float64
	PriceRecoveryRatioMax float64
}

type WarrantData struct {
	Stock              *qotcommon.Security
	Owner              *qotcommon.Security
	Type               int32
	Issuer             int32
	MaturityTime       string
	MaturityTimestamp  float64
	ListTime           string
	ListTimestamp      float64
	LastTradeTime      string
	LastTradeTimestamp float64
	RecoveryPrice      float64
	ConversionRatio    float64
	LotSize            int32
	StrikePrice        float64
	LastClosePrice     float64
	Name               string
	CurPrice           float64
	PriceChangeVal     float64
	ChangeRate         float64
	Status             int32
	BidPrice           float64
	AskPrice           float64
	BidVol             int64
	AskVol             int64
	Volume             int64
	Turnover           float64
	Score              float64
	Premium            float64
	BreakEvenPoint     float64
	Leverage           float64
	Ipop               float64
	PriceRecoveryRatio float64
	ConversionPrice    float64
	StreetRate         float64
	StreetVol          int64
	Amplitude          float64
	IssueSize          int64
	HighPrice          float64
	LowPrice           float64
	ImpliedVolatility  float64
	Delta              float64
	EffectiveLeverage  float64
	UpperStrikePrice   float64
	LowerStrikePrice   float64
	InLinePriceStatus  int32
}

type GetWarrantResponse struct {
	LastPage        bool
	AllCount        int32
	WarrantDataList []*WarrantData
}

func GetWarrant(c *futuapi.Client, req *GetWarrantRequest) (*GetWarrantResponse, error) {
	c2s := &qotgetwarrant.C2S{
		Begin:                 &req.Begin,
		Num:                   &req.Num,
		SortField:             &req.SortField,
		Ascend:                &req.Ascend,
		Owner:                 req.Owner,
		TypeList:              req.TypeList,
		IssuerList:            req.IssuerList,
		MaturityTimeMin:       &req.MaturityTimeMin,
		MaturityTimeMax:       &req.MaturityTimeMax,
		IpoPeriod:             &req.IpoPeriod,
		PriceType:             &req.PriceType,
		Status:                &req.Status,
		CurPriceMin:           &req.CurPriceMin,
		CurPriceMax:           &req.CurPriceMax,
		StrikePriceMin:        &req.StrikePriceMin,
		StrikePriceMax:        &req.StrikePriceMax,
		StreetMin:             &req.StreetMin,
		StreetMax:             &req.StreetMax,
		ConversionMin:         &req.ConversionMin,
		ConversionMax:         &req.ConversionMax,
		VolMin:                &req.VolMin,
		VolMax:                &req.VolMax,
		PremiumMin:            &req.PremiumMin,
		PremiumMax:            &req.PremiumMax,
		LeverageRatioMin:      &req.LeverageRatioMin,
		LeverageRatioMax:      &req.LeverageRatioMax,
		DeltaMin:              &req.DeltaMin,
		DeltaMax:              &req.DeltaMax,
		ImpliedMin:            &req.ImpliedMin,
		ImpliedMax:            &req.ImpliedMax,
		RecoveryPriceMin:      &req.RecoveryPriceMin,
		RecoveryPriceMax:      &req.RecoveryPriceMax,
		PriceRecoveryRatioMin: &req.PriceRecoveryRatioMin,
		PriceRecoveryRatioMax: &req.PriceRecoveryRatioMax,
	}

	pkt := &qotgetwarrant.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetWarrant, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetwarrant.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetWarrant failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetWarrant: s2c is nil")
	}

	result := &GetWarrantResponse{
		LastPage:        s2c.GetLastPage(),
		AllCount:        s2c.GetAllCount(),
		WarrantDataList: make([]*WarrantData, 0, len(s2c.GetWarrantDataList())),
	}

	for _, w := range s2c.GetWarrantDataList() {
		result.WarrantDataList = append(result.WarrantDataList, &WarrantData{
			Stock:              w.GetStock(),
			Owner:              w.GetOwner(),
			Type:               w.GetType(),
			Issuer:             w.GetIssuer(),
			MaturityTime:       w.GetMaturityTime(),
			MaturityTimestamp:  w.GetMaturityTimestamp(),
			ListTime:           w.GetListTime(),
			ListTimestamp:      w.GetListTimestamp(),
			LastTradeTime:      w.GetLastTradeTime(),
			LastTradeTimestamp: w.GetLastTradeTimestamp(),
			RecoveryPrice:      w.GetRecoveryPrice(),
			ConversionRatio:    w.GetConversionRatio(),
			LotSize:            w.GetLotSize(),
			StrikePrice:        w.GetStrikePrice(),
			LastClosePrice:     w.GetLastClosePrice(),
			Name:               w.GetName(),
			CurPrice:           w.GetCurPrice(),
			PriceChangeVal:     w.GetPriceChangeVal(),
			ChangeRate:         w.GetChangeRate(),
			Status:             w.GetStatus(),
			BidPrice:           w.GetBidPrice(),
			AskPrice:           w.GetAskPrice(),
			BidVol:             w.GetBidVol(),
			AskVol:             w.GetAskVol(),
			Volume:             w.GetVolume(),
			Turnover:           w.GetTurnover(),
			Score:              w.GetScore(),
			Premium:            w.GetPremium(),
			BreakEvenPoint:     w.GetBreakEvenPoint(),
			Leverage:           w.GetLeverage(),
			Ipop:               w.GetIpop(),
			PriceRecoveryRatio: w.GetPriceRecoveryRatio(),
			ConversionPrice:    w.GetConversionPrice(),
			StreetRate:         w.GetStreetRate(),
			StreetVol:          w.GetStreetVol(),
			Amplitude:          w.GetAmplitude(),
			IssueSize:          w.GetIssueSize(),
			HighPrice:          w.GetHighPrice(),
			LowPrice:           w.GetLowPrice(),
			ImpliedVolatility:  w.GetImpliedVolatility(),
			Delta:              w.GetDelta(),
			EffectiveLeverage:  w.GetEffectiveLeverage(),
			UpperStrikePrice:   w.GetUpperStrikePrice(),
			LowerStrikePrice:   w.GetLowerStrikePrice(),
			InLinePriceStatus:  w.GetInLinePriceStatus(),
		})
	}

	return result, nil
}

type GetSuspendRequest struct {
	SecurityList []*qotcommon.Security
	BeginTime    string
	EndTime      string
}

type SuspendInfo struct {
	Time      string
	Timestamp float64
}

type SecuritySuspendInfo struct {
	Security    *qotcommon.Security
	SuspendList []*SuspendInfo
}

type GetSuspendResponse struct {
	SecuritySuspendList []*SecuritySuspendInfo
}

func GetSuspend(c *futuapi.Client, req *GetSuspendRequest) (*GetSuspendResponse, error) {
	c2s := &qotgetsuspend.C2S{
		SecurityList: req.SecurityList,
		BeginTime:    &req.BeginTime,
		EndTime:      &req.EndTime,
	}

	pkt := &qotgetsuspend.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetSuspend, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetsuspend.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetSuspend failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetSuspend: s2c is nil")
	}

	result := &GetSuspendResponse{
		SecuritySuspendList: make([]*SecuritySuspendInfo, 0, len(s2c.GetSecuritySuspendList())),
	}

	for _, ss := range s2c.GetSecuritySuspendList() {
		info := &SecuritySuspendInfo{
			Security:    ss.GetSecurity(),
			SuspendList: make([]*SuspendInfo, 0, len(ss.GetSuspendList())),
		}
		for _, s := range ss.GetSuspendList() {
			info.SuspendList = append(info.SuspendList, &SuspendInfo{
				Time:      s.GetTime(),
				Timestamp: s.GetTimestamp(),
			})
		}
		result.SecuritySuspendList = append(result.SecuritySuspendList, info)
	}

	return result, nil
}
