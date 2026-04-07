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
	"gitee.com/shing1211/futuapi4go/pb/qotgetcodechange"
	"gitee.com/shing1211/futuapi4go/pb/qotgetfutureinfo"
	"gitee.com/shing1211/futuapi4go/pb/qotgetholdingchangelist"
	"gitee.com/shing1211/futuapi4go/pb/qotgetipolist"
	"gitee.com/shing1211/futuapi4go/pb/qotgetkl"
	// "gitee.com/shing1211/futuapi4go/pb/qotgetoptionchain"
	// "gitee.com/shing1211/futuapi4go/pb/qotgetoptionexpirationdate"
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
	"gitee.com/shing1211/futuapi4go/pb/qotgetusersecuritygroup"
	"gitee.com/shing1211/futuapi4go/pb/qotgetwarrant"
	"gitee.com/shing1211/futuapi4go/pb/qotmodifyusersecurity"
	"gitee.com/shing1211/futuapi4go/pb/qotregqotpush"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesthistorykl"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesthistoryklquota"
	"gitee.com/shing1211/futuapi4go/pb/qotrequestrehab"
	"gitee.com/shing1211/futuapi4go/pb/qotrequesttradedate"
	"gitee.com/shing1211/futuapi4go/pb/qotsetpricereminder"
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
	ProtoID_GetCodeChange           = 2210
	ProtoID_GetFutureInfo           = 2211
	ProtoID_GetIpoList              = 2212
	ProtoID_GetHoldingChangeList    = 2213
	ProtoID_RequestRehab            = 2214
	ProtoID_GetUserSecurityGroup    = 2402
	ProtoID_ModifyUserSecurity      = 2403
	ProtoID_SetPriceReminder        = 2405
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
	ProtoID_RegQotPush              = 3003
	ProtoID_RequestHistoryKLQuota   = 3104
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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("Subscribe failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
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
	return nil, fmt.Errorf("GetOptionExpirationDate: not implemented due to protobuf issues")
}

type GetOptionChainRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
	Type            int32
	Condition       int32
	BeginTime       string
	EndTime         string
	DataFilter      interface{}
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
	return nil, fmt.Errorf("GetOptionChain: not implemented due to protobuf issues")
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

type GetFutureInfoRequest struct {
	SecurityList []*qotcommon.Security
}

type FutureInfo struct {
	Name               string
	Security           *qotcommon.Security
	LastTradeTime      string
	LastTradeTimestamp float64
	Owner              *qotcommon.Security
	OwnerOther         string
	Exchange           string
	ContractType       string
	ContractSize       float64
	ContractSizeUnit   string
	QuoteCurrency      string
	MinVar             float64
	MinVarUnit         string
	QuoteUnit          string
	TradeTimeList      []*qotgetfutureinfo.TradeTime
	TimeZone           string
	ExchangeFormatUrl  string
	Origin             *qotcommon.Security
}

type GetFutureInfoResponse struct {
	FutureInfoList []*FutureInfo
}

func GetFutureInfo(c *futuapi.Client, req *GetFutureInfoRequest) (*GetFutureInfoResponse, error) {
	c2s := &qotgetfutureinfo.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetfutureinfo.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetFutureInfo, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetfutureinfo.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetFutureInfo failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetFutureInfo: s2c is nil")
	}

	result := &GetFutureInfoResponse{
		FutureInfoList: make([]*FutureInfo, 0, len(s2c.GetFutureInfoList())),
	}

	for _, fi := range s2c.GetFutureInfoList() {
		result.FutureInfoList = append(result.FutureInfoList, &FutureInfo{
			Name:               fi.GetName(),
			Security:           fi.GetSecurity(),
			LastTradeTime:      fi.GetLastTradeTime(),
			LastTradeTimestamp: fi.GetLastTradeTimestamp(),
			Owner:              fi.GetOwner(),
			OwnerOther:         fi.GetOwnerOther(),
			Exchange:           fi.GetExchange(),
			ContractType:       fi.GetContractType(),
			ContractSize:       fi.GetContractSize(),
			ContractSizeUnit:   fi.GetContractSizeUnit(),
			QuoteCurrency:      fi.GetQuoteCurrency(),
			MinVar:             fi.GetMinVar(),
			MinVarUnit:         fi.GetMinVarUnit(),
			QuoteUnit:          fi.GetQuoteUnit(),
			TradeTimeList:      fi.GetTradeTime(),
			TimeZone:           fi.GetTimeZone(),
			ExchangeFormatUrl:  fi.GetExchangeFormatUrl(),
			Origin:             fi.GetOrigin(),
		})
	}

	return result, nil
}

type GetCodeChangeRequest struct {
	SecurityList   []*qotcommon.Security
	TimeFilterList []*qotgetcodechange.TimeFilter
	TypeList       []int32
}

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

type GetCodeChangeResponse struct {
	CodeChangeList []*CodeChangeInfo
}

func GetCodeChange(c *futuapi.Client, req *GetCodeChangeRequest) (*GetCodeChangeResponse, error) {
	c2s := &qotgetcodechange.C2S{
		SecurityList:   req.SecurityList,
		TimeFilterList: req.TimeFilterList,
		TypeList:       req.TypeList,
	}

	pkt := &qotgetcodechange.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetCodeChange, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetcodechange.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetCodeChange failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetCodeChange: s2c is nil")
	}

	result := &GetCodeChangeResponse{
		CodeChangeList: make([]*CodeChangeInfo, 0, len(s2c.GetCodeChangeList())),
	}

	for _, cc := range s2c.GetCodeChangeList() {
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

type GetIpoListRequest struct {
	Market int32
}

type BasicIpoData struct {
	Security      *qotcommon.Security
	Name          string
	ListTime      string
	ListTimestamp float64
}

type CNIpoExData struct {
	ApplyCode              string
	IssueSize              int64
	OnlineIssueSize        int64
	ApplyUpperLimit        int64
	ApplyLimitMarketValue  int64
	IsEstimateIpoPrice     bool
	IpoPrice               float64
	IndustryPeRate         float64
	IsEstimateWinningRatio bool
	WinningRatio           float64
	IssuePeRate            float64
	ApplyTime              string
	ApplyTimestamp         float64
	WinningTime            string
	WinningTimestamp       float64
	IsHasWon               bool
	WinningNumDataList     []*qotgetipolist.WinningNumData
}

type HKIpoExData struct {
	IpoPriceMin       float64
	IpoPriceMax       float64
	ListPrice         float64
	LotSize           int32
	EntrancePrice     float64
	IsSubscribeStatus bool
	ApplyEndTime      string
	ApplyEndTimestamp float64
}

type USIpoExData struct {
	IpoPriceMin float64
	IpoPriceMax float64
	IssueSize   int64
}

type IpoData struct {
	Basic    *BasicIpoData
	CnExData *CNIpoExData
	HkExData *HKIpoExData
	UsExData *USIpoExData
}

type GetIpoListResponse struct {
	IpoList []*IpoData
}

func GetIpoList(c *futuapi.Client, req *GetIpoListRequest) (*GetIpoListResponse, error) {
	c2s := &qotgetipolist.C2S{
		Market: &req.Market,
	}

	pkt := &qotgetipolist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetIpoList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetipolist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetIpoList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetIpoList: s2c is nil")
	}

	result := &GetIpoListResponse{
		IpoList: make([]*IpoData, 0, len(s2c.GetIpoList())),
	}

	for _, ipo := range s2c.GetIpoList() {
		ipoData := &IpoData{}

		if basic := ipo.GetBasic(); basic != nil {
			ipoData.Basic = &BasicIpoData{
				Security:      basic.GetSecurity(),
				Name:          basic.GetName(),
				ListTime:      basic.GetListTime(),
				ListTimestamp: basic.GetListTimestamp(),
			}
		}

		if cnEx := ipo.GetCnExData(); cnEx != nil {
			ipoData.CnExData = &CNIpoExData{
				ApplyCode:              cnEx.GetApplyCode(),
				IssueSize:              cnEx.GetIssueSize(),
				OnlineIssueSize:        cnEx.GetOnlineIssueSize(),
				ApplyUpperLimit:        cnEx.GetApplyUpperLimit(),
				ApplyLimitMarketValue:  cnEx.GetApplyLimitMarketValue(),
				IsEstimateIpoPrice:     cnEx.GetIsEstimateIpoPrice(),
				IpoPrice:               cnEx.GetIpoPrice(),
				IndustryPeRate:         cnEx.GetIndustryPeRate(),
				IsEstimateWinningRatio: cnEx.GetIsEstimateWinningRatio(),
				WinningRatio:           cnEx.GetWinningRatio(),
				IssuePeRate:            cnEx.GetIssuePeRate(),
				ApplyTime:              cnEx.GetApplyTime(),
				ApplyTimestamp:         cnEx.GetApplyTimestamp(),
				WinningTime:            cnEx.GetWinningTime(),
				WinningTimestamp:       cnEx.GetWinningTimestamp(),
				IsHasWon:               cnEx.GetIsHasWon(),
				WinningNumDataList:     cnEx.GetWinningNumData(),
			}
		}

		if hkEx := ipo.GetHkExData(); hkEx != nil {
			ipoData.HkExData = &HKIpoExData{
				IpoPriceMin:       hkEx.GetIpoPriceMin(),
				IpoPriceMax:       hkEx.GetIpoPriceMax(),
				ListPrice:         hkEx.GetListPrice(),
				LotSize:           hkEx.GetLotSize(),
				EntrancePrice:     hkEx.GetEntrancePrice(),
				IsSubscribeStatus: hkEx.GetIsSubscribeStatus(),
				ApplyEndTime:      hkEx.GetApplyEndTime(),
				ApplyEndTimestamp: hkEx.GetApplyEndTimestamp(),
			}
		}

		if usEx := ipo.GetUsExData(); usEx != nil {
			ipoData.UsExData = &USIpoExData{
				IpoPriceMin: usEx.GetIpoPriceMin(),
				IpoPriceMax: usEx.GetIpoPriceMax(),
				IssueSize:   usEx.GetIssueSize(),
			}
		}

		result.IpoList = append(result.IpoList, ipoData)
	}

	return result, nil
}

type GetHoldingChangeListRequest struct {
	Security       *qotcommon.Security
	HolderCategory int32
	BeginTime      string
	EndTime        string
}

type GetHoldingChangeListResponse struct {
	Security          *qotcommon.Security
	HoldingChangeList []*qotcommon.ShareHoldingChange
}

func GetHoldingChangeList(c *futuapi.Client, req *GetHoldingChangeListRequest) (*GetHoldingChangeListResponse, error) {
	c2s := &qotgetholdingchangelist.C2S{
		Security:       req.Security,
		HolderCategory: &req.HolderCategory,
		BeginTime:      &req.BeginTime,
		EndTime:        &req.EndTime,
	}

	pkt := &qotgetholdingchangelist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetHoldingChangeList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetholdingchangelist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetHoldingChangeList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHoldingChangeList: s2c is nil")
	}

	return &GetHoldingChangeListResponse{
		Security:          s2c.GetSecurity(),
		HoldingChangeList: s2c.GetHoldingChangeList(),
	}, nil
}

type GetUserSecurityGroupRequest struct {
	GroupType int32
}

type UserSecurityGroupData struct {
	GroupName string
	GroupType int32
}

type GetUserSecurityGroupResponse struct {
	GroupList []*UserSecurityGroupData
}

func GetUserSecurityGroup(c *futuapi.Client, req *GetUserSecurityGroupRequest) (*GetUserSecurityGroupResponse, error) {
	c2s := &qotgetusersecuritygroup.C2S{
		GroupType: &req.GroupType,
	}

	pkt := &qotgetusersecuritygroup.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetUserSecurityGroup, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotgetusersecuritygroup.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetUserSecurityGroup failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetUserSecurityGroup: s2c is nil")
	}

	result := &GetUserSecurityGroupResponse{
		GroupList: make([]*UserSecurityGroupData, 0, len(s2c.GetGroupList())),
	}

	for _, g := range s2c.GetGroupList() {
		result.GroupList = append(result.GroupList, &UserSecurityGroupData{
			GroupName: g.GetGroupName(),
			GroupType: g.GetGroupType(),
		})
	}

	return result, nil
}

type ModifyUserSecurityRequest struct {
	GroupName    string
	Op           int32
	SecurityList []*qotcommon.Security
}

type ModifyUserSecurityResponse struct {
	RetType int32
	RetMsg  string
}

func ModifyUserSecurity(c *futuapi.Client, req *ModifyUserSecurityRequest) (*ModifyUserSecurityResponse, error) {
	c2s := &qotmodifyusersecurity.C2S{
		GroupName:    &req.GroupName,
		Op:           &req.Op,
		SecurityList: req.SecurityList,
	}

	pkt := &qotmodifyusersecurity.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_ModifyUserSecurity, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotmodifyusersecurity.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("ModifyUserSecurity failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return &ModifyUserSecurityResponse{
		RetType: rsp.GetRetType(),
		RetMsg:  rsp.GetRetMsg(),
	}, nil
}

type SetPriceReminderRequest struct {
	Security *qotcommon.Security
	Op       int32
	Key      int64
	Type     int32
	Freq     int32
	Value    float64
	Note     string
}

type SetPriceReminderResponse struct {
	Key int64
}

func SetPriceReminder(c *futuapi.Client, req *SetPriceReminderRequest) (*SetPriceReminderResponse, error) {
	c2s := &qotsetpricereminder.C2S{
		Security: req.Security,
		Op:       &req.Op,
		Key:      &req.Key,
		Type:     &req.Type,
		Freq:     &req.Freq,
		Value:    &req.Value,
		Note:     &req.Note,
	}

	pkt := &qotsetpricereminder.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_SetPriceReminder, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotsetpricereminder.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("SetPriceReminder failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("SetPriceReminder: s2c is nil")
	}

	return &SetPriceReminderResponse{
		Key: s2c.GetKey(),
	}, nil
}

type RegQotPushRequest struct {
	SecurityList  []*qotcommon.Security
	SubTypeList   []int32
	RehabTypeList []int32
	IsRegOrUnReg  bool
	IsFirstPush   bool
}

type RegQotPushResponse struct {
	RetType int32
	RetMsg  string
}

func RegQotPush(c *futuapi.Client, req *RegQotPushRequest) (*RegQotPushResponse, error) {
	c2s := &qotregqotpush.C2S{
		SecurityList:  req.SecurityList,
		SubTypeList:   req.SubTypeList,
		RehabTypeList: req.RehabTypeList,
		IsRegOrUnReg:  &req.IsRegOrUnReg,
		IsFirstPush:   &req.IsFirstPush,
	}

	pkt := &qotregqotpush.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_RegQotPush, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotregqotpush.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("RegQotPush failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return &RegQotPushResponse{
		RetType: rsp.GetRetType(),
		RetMsg:  rsp.GetRetMsg(),
	}, nil
}

type RequestRehabRequest struct {
	Security *qotcommon.Security
}

type RequestRehabResponse struct {
	RehabList []*qotcommon.Rehab
}

func RequestRehab(c *futuapi.Client, req *RequestRehabRequest) (*RequestRehabResponse, error) {
	c2s := &qotrequestrehab.C2S{
		Security: req.Security,
	}

	pkt := &qotrequestrehab.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_RequestRehab, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotrequestrehab.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("RequestRehab failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestRehab: s2c is nil")
	}

	return &RequestRehabResponse{
		RehabList: s2c.GetRehabList(),
	}, nil
}

type RequestHistoryKLQuotaRequest struct {
	GetDetail bool
}

type RequestHistoryKLQuotaResponse struct {
	UsedQuota   int32
	RemainQuota int32
	DetailList  []*qotrequesthistoryklquota.DetailItem
}

func RequestHistoryKLQuota(c *futuapi.Client, req *RequestHistoryKLQuotaRequest) (*RequestHistoryKLQuotaResponse, error) {
	c2s := &qotrequesthistoryklquota.C2S{
		BGetDetail: &req.GetDetail,
	}

	pkt := &qotrequesthistoryklquota.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_RequestHistoryKLQuota, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp qotrequesthistoryklquota.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("RequestHistoryKLQuota failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestHistoryKLQuota: s2c is nil")
	}

	return &RequestHistoryKLQuotaResponse{
		UsedQuota:   s2c.GetUsedQuota(),
		RemainQuota: s2c.GetRemainQuota(),
		DetailList:  s2c.GetDetailList(),
	}, nil
}
