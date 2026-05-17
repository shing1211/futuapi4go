// Package qot provides market data APIs for the Futu OpenD SDK.
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

package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
)

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
func GetOrderBook(ctx context.Context, c *futuapi.Client, req *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOrderBook: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetOrderBook: Security is nil")
	}
	c2s := &qotgetorderbook.C2S{
		Security: req.Security,
		Num:      &req.Num,
	}
	pkt := &qotgetorderbook.Request{C2S: c2s}
	var rsp qotgetorderbook.Response

	if err := c.RequestContext(ctx, ProtoID_GetOrderBook, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderBook", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetOrderBook", int32(common.RetType_RetType_Unknown), "s2c is nil")
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
		if ob == nil {
			continue
		}
		details := make([]*OrderBookDetail, 0, len(ob.GetDetailList()))
		for _, d := range ob.GetDetailList() {
			if d == nil {
				continue
			}
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
		if ob == nil {
			continue
		}
		details := make([]*OrderBookDetail, 0, len(ob.GetDetailList()))
		for _, d := range ob.GetDetailList() {
			if d == nil {
				continue
			}
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
	Timestamp    float64
	PushDataType int32
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
func GetTicker(ctx context.Context, c *futuapi.Client, req *GetTickerRequest) (*GetTickerResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTicker: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetTicker: Security is nil")
	}
	maxRetNum := req.Num
	c2s := &qotgetticker.C2S{
		Security:  req.Security,
		MaxRetNum: &maxRetNum,
	}

	pkt := &qotgetticker.Request{C2S: c2s}
	var rsp qotgetticker.Response

	if err := c.RequestContext(ctx, ProtoID_GetTicker, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTicker", rsp.GetRetType(), rsp.GetRetMsg())
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
		if t == nil {
			continue
		}
		result.TickerList = append(result.TickerList, &Ticker{
			Time:        t.GetTime(),
			Sequence:    t.GetSequence(),
			Dir:         t.GetDir(),
			Price:       t.GetPrice(),
			Volume:      t.GetVolume(),
			Turnover:    t.GetTurnover(),
			RecvTime:    t.GetRecvTime(),
			Type:        t.GetType(),
			TypeSign:    t.GetTypeSign(),
			Timestamp:   t.GetTimestamp(),
			PushDataType: t.GetPushDataType(),
		})
	}

	return result, nil
}

// RT represents a single real-time data point.
type RT struct {
	Time           string
	Minute         int32
	IsBlank        bool
	Price          float64
	LastClosePrice float64
	AvgPrice       float64
	Volume         int64
	Turnover       float64
	Timestamp      float64
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
func GetRT(ctx context.Context, c *futuapi.Client, req *GetRTRequest) (*GetRTResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRT: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetRT: Security is nil")
	}
	c2s := &qotgetrt.C2S{
		Security: req.Security,
	}

	pkt := &qotgetrt.Request{C2S: c2s}
	var rsp qotgetrt.Response

	if err := c.RequestContext(ctx, ProtoID_GetRT, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetRT", rsp.GetRetType(), rsp.GetRetMsg())
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
		if rt == nil {
			continue
		}
		result.RTList = append(result.RTList, &RT{
			Time:           rt.GetTime(),
			Minute:         rt.GetMinute(),
			IsBlank:        rt.GetIsBlank(),
			Price:          rt.GetPrice(),
			LastClosePrice: rt.GetLastClosePrice(),
			AvgPrice:       rt.GetAvgPrice(),
			Volume:         rt.GetVolume(),
			Turnover:       rt.GetTurnover(),
			Timestamp:      rt.GetTimestamp(),
		})
	}

	return result, nil
}

// Broker represents a broker (经纪) in the broker queue.
type Broker struct {
	ID      int64
	Name    string
	Pos     int32
	Volume  int64
	OrderID int64
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
func GetBroker(ctx context.Context, c *futuapi.Client, req *GetBrokerRequest) (*GetBrokerResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetBroker: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetbroker.C2S{
		Security: req.Security,
	}

	pkt := &qotgetbroker.Request{C2S: c2s}
	var rsp qotgetbroker.Response

	if err := c.RequestContext(ctx, ProtoID_GetBroker, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetBroker", rsp.GetRetType(), rsp.GetRetMsg())
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
		if b == nil {
			continue
		}
		result.AskBrokerList = append(result.AskBrokerList, &Broker{
			ID:      b.GetId(),
			Name:    b.GetName(),
			Pos:     b.GetPos(),
			Volume:  b.GetVolume(),
			OrderID: b.GetOrderID(),
		})
	}

	for _, b := range s2c.GetBrokerBidList() {
		if b == nil {
			continue
		}
		result.BidBrokerList = append(result.BidBrokerList, &Broker{
			ID:      b.GetId(),
			Name:    b.GetName(),
			Pos:     b.GetPos(),
			Volume:  b.GetVolume(),
			OrderID: b.GetOrderID(),
		})
	}

	return result, nil
}
