// Package qot provides market data APIs for the Futu OpenD SDK.
//
// This package covers real-time quotes, K-lines, order book, tick data,
// broker queue, capital flow, stock screening, options, warrants, and
// historical data requests. All functions require a connected client.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/qot"
//
//	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
//	securities := []*qotcommon.Security{{Market: &hkMarket, Code: ptrStr("00700")}}
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

// quotes, err := qot.GetBasicQot(cli, securities)
package qot

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcodechange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfutureinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetholdingchangelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetipolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionchain"
	qotgetoptionexpirationdate "github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionexpirationdate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplatesecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplateset"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetpricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrehab"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsecuritysnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsuspend"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusersecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusersecuritygroup"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetwarrant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotmodifyusersecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotregqotpush"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryklquota"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestrehab"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesttradedate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetpricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"

	"time"
)

const (
	ProtoID_GetBasicQot             = 3004
	ProtoID_GetKL                   = 3006
	ProtoID_GetOrderBook            = 3012
	ProtoID_GetTicker               = 3010
	ProtoID_GetRT                   = 3008
	ProtoID_GetMarketSnapshot       = 3203
	ProtoID_GetSecuritySnapshot     = 3203
	ProtoID_GetBroker               = 3014
	ProtoID_GetStaticInfo           = 3202
	ProtoID_GetPlateSet             = 3204
	ProtoID_GetPlateSecurity        = 3205
	ProtoID_GetSuspend              = 3201
	ProtoID_GetCodeChange           = 3216
	ProtoID_GetFutureInfo           = 3218
	ProtoID_GetIpoList              = 3217
	ProtoID_GetHoldingChangeList    = 3208
	ProtoID_RequestRehab            = 3105
	ProtoID_GetUserSecurityGroup    = 3222
	ProtoID_ModifyUserSecurity      = 3214
	ProtoID_SetPriceReminder        = 3220
	ProtoID_GetCapitalFlow          = 3211
	ProtoID_GetCapitalDistribution  = 3212
	ProtoID_StockFilter             = 3215
	ProtoID_GetOptionChain          = 3209
	ProtoID_GetRehab                = 3208
	ProtoID_GetOptionExpirationDate = 3224
	ProtoID_GetWarrant              = 3210
	ProtoID_GetUserSecurity         = 3213
	ProtoID_GetPriceReminder        = 3221
	ProtoID_RequestTradeDate        = 3219
	ProtoID_Subscribe               = 3001
	ProtoID_RegQotPush              = 3002
	ProtoID_RequestHistoryKL        = 3103
	ProtoID_RequestHistoryKLQuota   = 3104
)

// BasicQot represents basic quote data for a security.
type BasicQot struct {
	Security       *qotcommon.Security
	Name           string
	IsSuspended    bool
	UpdateTime     string
	HighPrice      float64
	OpenPrice      float64
	LowPrice       float64
	CurPrice       float64
	LastClosePrice float64
	Volume         int64
	Turnover       float64
	TurnoverRate   float64
	Amplitude      float64
}

// GetBasicQot returns basic quote data for the given securities.
func GetBasicQot(ctx context.Context, c *futuapi.Client, securityList []*qotcommon.Security) ([]*BasicQot, error) {
	c2s := &qotgetbasicqot.C2S{
		SecurityList: securityList,
	}
	req := &qotgetbasicqot.Request{C2S: c2s}
	var rsp qotgetbasicqot.Response

	if err := c.RequestContext(ctx, ProtoID_GetBasicQot, req, &rsp); err != nil {
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
			Security:       bq.GetSecurity(),
			Name:           bq.GetName(),
			IsSuspended:    bq.GetIsSuspended(),
			UpdateTime:     bq.GetUpdateTime(),
			HighPrice:      bq.GetHighPrice(),
			OpenPrice:      bq.GetOpenPrice(),
			LowPrice:       bq.GetLowPrice(),
			CurPrice:       bq.GetCurPrice(),
			LastClosePrice: bq.GetLastClosePrice(),
			Volume:         bq.GetVolume(),
			Turnover:       bq.GetTurnover(),
			TurnoverRate:   bq.GetTurnoverRate(),
			Amplitude:      bq.GetAmplitude(),
		})
	}

	return result, nil
}

// KLine represents a single K-line (candlestick) data point.
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

// GetKLRequest defines parameters for GetKL.
type GetKLRequest struct {
	Security  *qotcommon.Security
	RehabType int32
	KLType    int32
	ReqNum    int32
}

// GetKLResponse is the response type for GetKL.
type GetKLResponse struct {
	Security *qotcommon.Security
	Name     string
	KLList   []*KLine
}

// GetKL returns K-line (candlestick) data for the given security.
func GetKL(c *futuapi.Client, req *GetKLRequest) (*GetKLResponse, error) {
	c2s := &qotgetkl.C2S{
		Security:  req.Security,
		RehabType: &req.RehabType,
		KlType:    &req.KLType,
		ReqNum:    &req.ReqNum,
	}
	pkt := &qotgetkl.Request{C2S: c2s}
	var rsp qotgetkl.Response

	if err := c.Request(ProtoID_GetKL, pkt, &rsp); err != nil {
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

// OrderBook represents a price level in the order book.
type OrderBook struct {
	Price      float64
	Volume     int64
	OrderCount int32
	DetailList []*OrderBookDetail
}

// OrderBookDetail represents a single order in the order book.
type OrderBookDetail struct {
	OrderID int64
	Volume  int64
}

// GetOrderBookRequest defines parameters for GetOrderBook.
type GetOrderBookRequest struct {
	Security *qotcommon.Security
	Num      int32
}

// GetOrderBookResponse is the response type for GetOrderBook.
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

// GetOrderBook returns the order book (买卖盘) for the given security.
func GetOrderBook(c *futuapi.Client, req *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	c2s := &qotgetorderbook.C2S{
		Security: req.Security,
		Num:      &req.Num,
	}
	pkt := &qotgetorderbook.Request{C2S: c2s}
	var rsp qotgetorderbook.Response

	if err := c.Request(ProtoID_GetOrderBook, pkt, &rsp); err != nil {
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
		details := make([]*OrderBookDetail, 0, len(ob.GetDetailList()))
		for _, d := range ob.GetDetailList() {
			details = append(details, &OrderBookDetail{
				OrderID: d.GetOrderID(),
				Volume:  d.GetVolume(),
			})
		}
		result.OrderBookAskList = append(result.OrderBookAskList, &OrderBook{
			Price:      ob.GetPrice(),
			Volume:     ob.GetVolume(),
			OrderCount: ob.GetOrederCount(),
			DetailList: details,
		})
	}

	for _, ob := range s2c.GetOrderBookBidList() {
		details := make([]*OrderBookDetail, 0, len(ob.GetDetailList()))
		for _, d := range ob.GetDetailList() {
			details = append(details, &OrderBookDetail{
				OrderID: d.GetOrderID(),
				Volume:  d.GetVolume(),
			})
		}
		result.OrderBookBidList = append(result.OrderBookBidList, &OrderBook{
			Price:      ob.GetPrice(),
			Volume:     ob.GetVolume(),
			OrderCount: ob.GetOrederCount(),
			DetailList: details,
		})
	}

	return result, nil
}

// Ticker represents a single trade tick data point.
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

// GetTickerRequest defines parameters for GetTicker.
type GetTickerRequest struct {
	Security *qotcommon.Security
	Num      int32
}

// GetTickerResponse is the response type for GetTicker.
type GetTickerResponse struct {
	Security   *qotcommon.Security
	Name       string
	TickerList []*Ticker
}

// GetTicker returns recent tick (逐笔成交) data for the given security.
func GetTicker(c *futuapi.Client, req *GetTickerRequest) (*GetTickerResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RT represents a single real-time data point.
type RT struct {
	Time           string
	Price          float64
	LastClosePrice float64
	AvgPrice       float64
	Volume         int64
	Turnover       float64
}

// GetRTRequest defines parameters for GetRT.
type GetRTRequest struct {
	Security *qotcommon.Security
}

// GetRTResponse is the response type for GetRT.
type GetRTResponse struct {
	Security *qotcommon.Security
	Name     string
	RTList   []*RT
}

// GetRT returns real-time (分时) data for the given security.
func GetRT(c *futuapi.Client, req *GetRTRequest) (*GetRTResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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
			Time:           rt.GetTime(),
			Price:          rt.GetPrice(),
			LastClosePrice: rt.GetLastClosePrice(),
			AvgPrice:       rt.GetAvgPrice(),
			Volume:         rt.GetVolume(),
			Turnover:       rt.GetTurnover(),
		})
	}

	return result, nil
}

// Broker represents a broker (经纪) in the broker queue.
type Broker struct {
	ID     int64
	Name   string
	Pos    int32
	Volume int64
}

// GetBrokerRequest defines parameters for GetBroker.
type GetBrokerRequest struct {
	Security *qotcommon.Security
	Num      int32
}

// GetBrokerResponse is the response type for GetBroker.
type GetBrokerResponse struct {
	Security      *qotcommon.Security
	Name          string
	AskBrokerList []*Broker
	BidBrokerList []*Broker
}

// GetBroker returns broker queue data (经纪队列) for the given security.
func GetBroker(c *futuapi.Client, req *GetBrokerRequest) (*GetBrokerResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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
func GetStaticInfo(c *futuapi.Client, req *GetStaticInfoRequest) (*GetStaticInfoResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// Plate represents a market plate (板块).
type Plate struct {
	Plate *qotcommon.Security
	Name  string
}

// GetPlateSetRequest defines parameters for GetPlateSet.
type GetPlateSetRequest struct {
	Market       int32
	PlateSetType int32
}

// GetPlateSetResponse is the response type for GetPlateSet.
type GetPlateSetResponse struct {
	PlateSetList []*Plate
}

// GetPlateSet returns the set of plates for the given market.
func GetPlateSet(c *futuapi.Client, req *GetPlateSetRequest) (*GetPlateSetResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetplateset.C2S{
		Market:       &req.Market,
		PlateSetType: &req.PlateSetType,
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetPlateSecurityRequest defines parameters for GetPlateSecurity.
type GetPlateSecurityRequest struct {
	Plate *qotcommon.Security
}

// GetPlateSecurityResponse is the response type for GetPlateSecurity.
type GetPlateSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetPlateSecurity returns securities belonging to the given plate.
func GetPlateSecurity(c *futuapi.Client, req *GetPlateSecurityRequest) (*GetPlateSecurityResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RequestTradeDateRequest defines parameters for RequestTradeDate.
type RequestTradeDateRequest struct {
	Market    int32
	BeginTime string
	EndTime   string
	Security  *qotcommon.Security
}

// RequestTradeDateResponse is the response type for RequestTradeDate.
type RequestTradeDateResponse struct {
	TradeDateList []*qotrequesttradedate.TradeDate
}

// RequestTradeDate requests trade dates for a specific security.
func RequestTradeDate(c *futuapi.Client, req *RequestTradeDateRequest) (*RequestTradeDateResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RequestHistoryKLRequest defines parameters for RequestHistoryKL.
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

// RequestHistoryKLResponse is the response type for RequestHistoryKL.
type RequestHistoryKLResponse struct {
	Security   *qotcommon.Security
	Name       string
	KLList     []*KLine
	NextReqKey []byte
}

// RequestHistoryKL requests historical K-line (candlestick) data for the given security.
func RequestHistoryKL(c *futuapi.Client, req *RequestHistoryKLRequest) (*RequestHistoryKLResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

	result := &RequestHistoryKLResponse{
		Security:   s2c.GetSecurity(),
		Name:       s2c.GetName(),
		NextReqKey: s2c.GetNextReqKey(),
		KLList:     make([]*KLine, 0, len(s2c.GetKlList())),
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

// GetSecuritySnapshotRequest defines parameters for GetSecuritySnapshot.
type GetSecuritySnapshotRequest struct {
	SecurityList []*qotcommon.Security
}

// GetSecuritySnapshotResponse is the response type for GetSecuritySnapshot.
type GetSecuritySnapshotResponse struct {
	SnapshotList []*qotgetsecuritysnapshot.Snapshot
}

// GetSecuritySnapshot returns snapshot data for the given securities.
func GetSecuritySnapshot(c *futuapi.Client, req *GetSecuritySnapshotRequest) (*GetSecuritySnapshotResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// SubType represents the type of market data subscription.
type SubType int32

// Subscription type constants.
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
	SubType_KL         SubType = 6 // default K-line type (Day)
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
}

// SubscribeResponse is the response type for Subscribe.
type SubscribeResponse struct {
	RetType int32
	RetMsg  string
}

// Subscribe subscribes to or unsubscribes from real-time market data.
func Subscribe(c *futuapi.Client, req *SubscribeRequest) (*SubscribeResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetCapitalFlowRequest defines parameters for GetCapitalFlow.
type GetCapitalFlowRequest struct {
	Security   *qotcommon.Security
	PeriodType int32
	BeginTime  string
	EndTime    string
}

// CapitalFlowItem represents a single capital flow data point.
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

// GetCapitalFlowResponse is the response type for GetCapitalFlow.
type GetCapitalFlowResponse struct {
	FlowItemList       []*CapitalFlowItem
	LastValidTime      string
	LastValidTimestamp float64
}

// GetCapitalFlow returns capital flow data for the given security.
func GetCapitalFlow(c *futuapi.Client, req *GetCapitalFlowRequest) (*GetCapitalFlowResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// CapitalDistribution represents the distribution of capital across different tiers.
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

// GetCapitalDistributionResponse is the response type for GetCapitalDistribution.
type GetCapitalDistributionResponse struct {
	CapitalDistribution *CapitalDistribution
}

// GetCapitalDistribution returns capital distribution data for the given security.
func GetCapitalDistribution(c *futuapi.Client, security *qotcommon.Security) (*GetCapitalDistributionResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetUserSecurityResponse is the response type for GetUserSecurity.
type GetUserSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetUserSecurity returns the list of user-defined securities in the specified group.
func GetUserSecurity(c *futuapi.Client, groupName string) (*GetUserSecurityResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// PriceReminderItemInfo represents a single price reminder item.
type PriceReminderItemInfo struct {
	Key                 int64
	Type                int32
	Value               float64
	Note                string
	Freq                int32
	IsEnable            bool
	ReminderSessionList []int32
}

// PriceReminderInfo represents the price reminder settings for a security.
type PriceReminderInfo struct {
	Security *qotcommon.Security
	Name     string
	ItemList []*PriceReminderItemInfo
}

// GetPriceReminderResponse is the response type for GetPriceReminder.
type GetPriceReminderResponse struct {
	PriceReminderList []*PriceReminderInfo
}

// GetPriceReminder returns price reminder settings for the given security.
func GetPriceReminder(c *futuapi.Client, security *qotcommon.Security, market int32) (*GetPriceReminderResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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
				Key:                 item.GetKey(),
				Type:                item.GetType(),
				Value:               item.GetValue(),
				Note:                item.GetNote(),
				Freq:                item.GetFreq(),
				IsEnable:            item.GetIsEnable(),
				ReminderSessionList: item.GetReminderSessionList(),
			})
		}
		result.PriceReminderList = append(result.PriceReminderList, info)
	}

	return result, nil
}

// GetOptionExpirationDateRequest defines parameters for GetOptionExpirationDate.
type GetOptionExpirationDateRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
}

// OptionExpirationDateInfo represents information about an option expiration date.
type OptionExpirationDateInfo struct {
	StrikeTime               string
	StrikeTimestamp          float64
	OptionExpiryDateDistance int32
	Cycle                    int32
}

// GetOptionExpirationDateResponse is the response type for GetOptionExpirationDate.
type GetOptionExpirationDateResponse struct {
	DateList []*OptionExpirationDateInfo
}

// GetOptionExpirationDate returns the list of option expiration dates for the given underlying.
func GetOptionExpirationDate(c *futuapi.Client, req *GetOptionExpirationDateRequest) (*GetOptionExpirationDateResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}

	c2s := &qotgetoptionexpirationdate.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
	}

	pkt := &qotgetoptionexpirationdate.Request{C2S: c2s}
	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOptionExpirationDate, serialNo, body); err != nil {
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

	var rsp qotgetoptionexpirationdate.Response
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

// GetOptionChainRequest defines parameters for GetOptionChain.
type GetOptionChainRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
	Type            int32
	Condition       int32
	BeginTime       string
	EndTime         string
	DataFilter      interface{}
}

// OptionItem represents a pair of call and put options at the same strike price.
type OptionItem struct {
	Call *qotcommon.SecurityStaticInfo
	Put  *qotcommon.SecurityStaticInfo
}

// OptionChain represents the option chain for a single expiration date.
type OptionChain struct {
	StrikeTime      string
	StrikeTimestamp float64
	Option          []*OptionItem
}

// GetOptionChainResponse is the response type for GetOptionChain.
type GetOptionChainResponse struct {
	OptionChain []*OptionChain
}

// GetOptionChain returns the option chain (期权链) for the given underlying security.
func GetOptionChain(c *futuapi.Client, req *GetOptionChainRequest) (*GetOptionChainResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}

	c2s := &qotgetoptionchain.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
		Type:            &req.Type,
		Condition:       &req.Condition,
		BeginTime:       &req.BeginTime,
		EndTime:         &req.EndTime,
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

	for _, chain := range s2c.GetOptionChain() {
		oc := &OptionChain{
			StrikeTime:      chain.GetStrikeTime(),
			StrikeTimestamp: chain.GetStrikeTimestamp(),
			Option:          make([]*OptionItem, 0, len(chain.GetOption())),
		}

		for _, opt := range chain.GetOption() {
			item := &OptionItem{}
			if opt.GetCall() != nil {
				item.Call = opt.GetCall()
			}
			if opt.GetPut() != nil {
				item.Put = opt.GetPut()
			}
			oc.Option = append(oc.Option, item)
		}

		result.OptionChain = append(result.OptionChain, oc)
	}

	return result, nil
}

// StockFilterRequest defines parameters for StockFilter.
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

// StockFilterData represents a single stock filter result.
type StockFilterData struct {
	Security                *qotcommon.Security
	Name                    string
	BaseDataList            []*qotstockfilter.BaseData
	AccumulateDataList      []*qotstockfilter.AccumulateData
	FinancialDataList       []*qotstockfilter.FinancialData
	CustomIndicatorDataList []*qotstockfilter.CustomIndicatorData
}

// StockFilterResponse is the response type for StockFilter.
type StockFilterResponse struct {
	LastPage bool
	AllCount int32
	DataList []*StockFilterData
}

// StockFilter filters stocks based on various criteria (选股).
func StockFilter(c *futuapi.Client, req *StockFilterRequest) (*StockFilterResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetWarrantRequest defines parameters for GetWarrant.
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

// WarrantData represents detailed data for a single warrant.
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

// GetWarrantResponse is the response type for GetWarrant.
type GetWarrantResponse struct {
	LastPage        bool
	AllCount        int32
	WarrantDataList []*WarrantData
}

// GetWarrant returns warrant data filtered by the given criteria.
func GetWarrant(c *futuapi.Client, req *GetWarrantRequest) (*GetWarrantResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetwarrant.C2S{
		Begin:     &req.Begin,
		Num:       &req.Num,
		SortField: &req.SortField,
		Ascend:    &req.Ascend,
		Owner:     req.Owner,
		TypeList:  req.TypeList,
	}
	if req.MaturityTimeMin != "" {
		c2s.MaturityTimeMin = &req.MaturityTimeMin
	}
	if req.MaturityTimeMax != "" {
		c2s.MaturityTimeMax = &req.MaturityTimeMax
	}
	if req.IpoPeriod != 0 {
		c2s.IpoPeriod = &req.IpoPeriod
	}
	if req.PriceType != 0 {
		c2s.PriceType = &req.PriceType
	}
	if req.Status != 0 {
		c2s.Status = &req.Status
	}
	if req.CurPriceMin != 0 {
		c2s.CurPriceMin = &req.CurPriceMin
	}
	if req.CurPriceMax != 0 {
		c2s.CurPriceMax = &req.CurPriceMax
	}
	if req.StrikePriceMin != 0 {
		c2s.StrikePriceMin = &req.StrikePriceMin
	}
	if req.StrikePriceMax != 0 {
		c2s.StrikePriceMax = &req.StrikePriceMax
	}
	if req.StreetMin != 0 {
		c2s.StreetMin = &req.StreetMin
	}
	if req.StreetMax != 0 {
		c2s.StreetMax = &req.StreetMax
	}
	if req.ConversionMin != 0 {
		c2s.ConversionMin = &req.ConversionMin
	}
	if req.ConversionMax != 0 {
		c2s.ConversionMax = &req.ConversionMax
	}
	if req.VolMin != 0 {
		c2s.VolMin = &req.VolMin
	}
	if req.VolMax != 0 {
		c2s.VolMax = &req.VolMax
	}
	if req.PremiumMin != 0 {
		c2s.PremiumMin = &req.PremiumMin
	}
	if req.PremiumMax != 0 {
		c2s.PremiumMax = &req.PremiumMax
	}
	if req.LeverageRatioMin != 0 {
		c2s.LeverageRatioMin = &req.LeverageRatioMin
	}
	if req.LeverageRatioMax != 0 {
		c2s.LeverageRatioMax = &req.LeverageRatioMax
	}
	if req.DeltaMin != 0 {
		c2s.DeltaMin = &req.DeltaMin
	}
	if req.DeltaMax != 0 {
		c2s.DeltaMax = &req.DeltaMax
	}
	if req.ImpliedMin != 0 {
		c2s.ImpliedMin = &req.ImpliedMin
	}
	if req.ImpliedMax != 0 {
		c2s.ImpliedMax = &req.ImpliedMax
	}
	if req.RecoveryPriceMin != 0 {
		c2s.RecoveryPriceMin = &req.RecoveryPriceMin
	}
	if req.RecoveryPriceMax != 0 {
		c2s.RecoveryPriceMax = &req.RecoveryPriceMax
	}
	if req.PriceRecoveryRatioMin != 0 {
		c2s.PriceRecoveryRatioMin = &req.PriceRecoveryRatioMin
	}
	if req.PriceRecoveryRatioMax != 0 {
		c2s.PriceRecoveryRatioMax = &req.PriceRecoveryRatioMax
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetSuspendRequest defines parameters for GetSuspend.
type GetSuspendRequest struct {
	SecurityList []*qotcommon.Security
	BeginTime    string
	EndTime      string
}

// SuspendInfo represents the suspension time for a security.
type SuspendInfo struct {
	Time      string
	Timestamp float64
}

// SecuritySuspendInfo represents suspension info for a single security.
type SecuritySuspendInfo struct {
	Security    *qotcommon.Security
	SuspendList []*SuspendInfo
}

// GetSuspendResponse is the response type for GetSuspend.
type GetSuspendResponse struct {
	SecuritySuspendList []*SecuritySuspendInfo
}

// GetSuspend returns suspension (停牌) information for the given securities.
func GetSuspend(c *futuapi.Client, req *GetSuspendRequest) (*GetSuspendResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetFutureInfoRequest defines parameters for GetFutureInfo.
type GetFutureInfoRequest struct {
	SecurityList []*qotcommon.Security
}

// FutureInfo represents detailed information about a futures contract.
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

// GetFutureInfoResponse is the response type for GetFutureInfo.
type GetFutureInfoResponse struct {
	FutureInfoList []*FutureInfo
}

// GetFutureInfo returns futures contract information for the given securities.
func GetFutureInfo(c *futuapi.Client, req *GetFutureInfoRequest) (*GetFutureInfoResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetCodeChangeRequest defines parameters for GetCodeChange.
type GetCodeChangeRequest struct {
	SecurityList   []*qotcommon.Security
	TimeFilterList []*qotgetcodechange.TimeFilter
	TypeList       []int32
}

// CodeChangeInfo represents information about a code change ( stock split, merger, etc.).
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

// GetCodeChangeResponse is the response type for GetCodeChange.
type GetCodeChangeResponse struct {
	CodeChangeList []*CodeChangeInfo
}

// GetCodeChange returns code change (股份代号变动) information for the given securities.
func GetCodeChange(c *futuapi.Client, req *GetCodeChangeRequest) (*GetCodeChangeResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetIpoListRequest defines parameters for GetIpoList.
type GetIpoListRequest struct {
	Market int32
}

// BasicIpoData represents basic IPO data.
type BasicIpoData struct {
	Security      *qotcommon.Security
	Name          string
	ListTime      string
	ListTimestamp float64
}

// CNIpoExData represents China A-share IPO extended data.
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

// HKIpoExData represents Hong Kong IPO extended data.
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

// USIpoExData represents US IPO extended data.
type USIpoExData struct {
	IpoPriceMin float64
	IpoPriceMax float64
	IssueSize   int64
}

// IpoData represents complete IPO data including basic and market-specific extended data.
type IpoData struct {
	Basic    *BasicIpoData
	CnExData *CNIpoExData
	HkExData *HKIpoExData
	UsExData *USIpoExData
}

// GetIpoListResponse is the response type for GetIpoList.
type GetIpoListResponse struct {
	IpoList []*IpoData
}

// GetIpoList returns the list of upcoming and recently listed IPOs for the given market.
func GetIpoList(c *futuapi.Client, req *GetIpoListRequest) (*GetIpoListResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetHoldingChangeListRequest defines parameters for GetHoldingChangeList.
type GetHoldingChangeListRequest struct {
	Security       *qotcommon.Security
	HolderCategory int32
	BeginTime      string
	EndTime        string
}

// GetHoldingChangeListResponse is the response type for GetHoldingChangeList.
type GetHoldingChangeListResponse struct {
	Security          *qotcommon.Security
	HoldingChangeList []*qotcommon.ShareHoldingChange
}

// GetHoldingChangeList returns the holding change list for the given security.
func GetHoldingChangeList(c *futuapi.Client, req *GetHoldingChangeListRequest) (*GetHoldingChangeListResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetUserSecurityGroupRequest defines parameters for GetUserSecurityGroup.
type GetUserSecurityGroupRequest struct {
	GroupType int32
}

// UserSecurityGroupData represents a user-defined security group.
type UserSecurityGroupData struct {
	GroupName string
	GroupType int32
}

// GetUserSecurityGroupResponse is the response type for GetUserSecurityGroup.
type GetUserSecurityGroupResponse struct {
	GroupList []*UserSecurityGroupData
}

// GetUserSecurityGroup returns the list of user-defined security groups.
func GetUserSecurityGroup(c *futuapi.Client, req *GetUserSecurityGroupRequest) (*GetUserSecurityGroupResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// ModifyUserSecurityRequest defines parameters for ModifyUserSecurity.
type ModifyUserSecurityRequest struct {
	GroupName    string
	Op           int32
	SecurityList []*qotcommon.Security
}

// ModifyUserSecurityResponse is the response type for ModifyUserSecurity.
type ModifyUserSecurityResponse struct {
	RetType int32
	RetMsg  string
}

// ModifyUserSecurity adds or removes securities from a user-defined group.
func ModifyUserSecurity(c *futuapi.Client, req *ModifyUserSecurityRequest) (*ModifyUserSecurityResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// SetPriceReminderRequest defines parameters for SetPriceReminder.
type SetPriceReminderRequest struct {
	Security *qotcommon.Security
	Op       int32
	Key      int64
	Type     int32
	Freq     int32
	Value    float64
	Note     string
}

// SetPriceReminderResponse is the response type for SetPriceReminder.
type SetPriceReminderResponse struct {
	Key int64
}

// SetPriceReminder creates, updates, or deletes a price reminder.
func SetPriceReminder(c *futuapi.Client, req *SetPriceReminderRequest) (*SetPriceReminderResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RegQotPushRequest defines parameters for RegQotPush.
type RegQotPushRequest struct {
	SecurityList  []*qotcommon.Security
	SubTypeList   []int32
	RehabTypeList []int32
	IsRegOrUnReg  bool
	IsFirstPush   bool
}

// RegQotPushResponse is the response type for RegQotPush.
type RegQotPushResponse struct {
	RetType int32
	RetMsg  string
}

// RegQotPush registers or unregisters for real-time push notifications.
func RegQotPush(c *futuapi.Client, req *RegQotPushRequest) (*RegQotPushResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RequestRehabRequest defines parameters for RequestRehab.
type RequestRehabRequest struct {
	Security *qotcommon.Security
}

// RequestRehabResponse is the response type for RequestRehab.
type RequestRehabResponse struct {
	RehabList []*qotcommon.Rehab
}

// RequestRehab requests rehabilitation (复权) data for the given security.
func RequestRehab(c *futuapi.Client, req *RequestRehabRequest) (*RequestRehabResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// RequestHistoryKLQuotaRequest defines parameters for RequestHistoryKLQuota.
type RequestHistoryKLQuotaRequest struct {
	GetDetail bool
}

// RequestHistoryKLQuotaResponse is the response type for RequestHistoryKLQuota.
type RequestHistoryKLQuotaResponse struct {
	UsedQuota   int32
	RemainQuota int32
	DetailList  []*qotrequesthistoryklquota.DetailItem
}

// RequestHistoryKLQuota returns the quota usage for historical K-line requests.
func RequestHistoryKLQuota(c *futuapi.Client, req *RequestHistoryKLQuotaRequest) (*RequestHistoryKLQuotaResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
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

	apiTimeout := c.Conn().APITimeout()
	if apiTimeout == 0 {
		apiTimeout = 30 * time.Second
	}
	pktResp, err := c.Conn().ReadResponse(serialNo, apiTimeout)
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

// GetRehabRequest defines parameters for GetRehab.
type GetRehabRequest struct {
	SecurityList []*qotcommon.Security
}

// GetRehabResponse is the response type for GetRehab.
type GetRehabResponse struct {
	SecurityRehabList []*qotgetrehab.SecurityRehab
}

// GetRehab returns rehabilitation (复权) data for the given securities.
func GetRehab(c *futuapi.Client, req *GetRehabRequest) (*GetRehabResponse, error) {
	if err := c.EnsureConnected(); err != nil {
		return nil, err
	}
	c2s := &qotgetrehab.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetrehab.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetRehab, serialNo, body); err != nil {
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

	var rsp qotgetrehab.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetRehab failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetRehab: s2c is nil")
	}

	return &GetRehabResponse{
		SecurityRehabList: s2c.GetSecurityRehabList(),
	}, nil
}
