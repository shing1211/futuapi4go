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
	"github.com/shing1211/futuapi4go/pkg/util"
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOrderBook", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetOrderBook", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetOrderBookResponse{
		Security:                s2c.Security,
		Name:                    util.ProtoStr(s2c.Name),
		SvrRecvTimeBid:          util.ProtoStr(s2c.SvrRecvTimeBid),
		SvrRecvTimeBidTimestamp: util.ProtoFloat64(s2c.SvrRecvTimeBidTimestamp),
		SvrRecvTimeAsk:          util.ProtoStr(s2c.SvrRecvTimeAsk),
		SvrRecvTimeAskTimestamp: util.ProtoFloat64(s2c.SvrRecvTimeAskTimestamp),
	}

	for _, ob := range s2c.OrderBookAskList {
		if ob == nil {
			continue
		}
		details := make([]*OrderBookDetail, 0, len(ob.DetailList))
		for _, d := range ob.DetailList {
			if d == nil {
				continue
			}
			details = append(details, &OrderBookDetail{
				OrderID: util.ProtoInt64(d.OrderID),
				Volume:  util.ProtoInt64(d.Volume),
			})
		}
		result.OrderBookAskList = append(result.OrderBookAskList, &OrderBook{
			Price:      util.ProtoFloat64(ob.Price),
			Volume:     util.ProtoInt64(ob.Volume),
			OrderCount: util.ProtoInt32(ob.OrederCount),
			DetailList: details,
		})
	}

	for _, ob := range s2c.OrderBookBidList {
		if ob == nil {
			continue
		}
		details := make([]*OrderBookDetail, 0, len(ob.DetailList))
		for _, d := range ob.DetailList {
			if d == nil {
				continue
			}
			details = append(details, &OrderBookDetail{
				OrderID: util.ProtoInt64(d.OrderID),
				Volume:  util.ProtoInt64(d.Volume),
			})
		}
		result.OrderBookBidList = append(result.OrderBookBidList, &OrderBook{
			Price:      util.ProtoFloat64(ob.Price),
			Volume:     util.ProtoInt64(ob.Volume),
			OrderCount: util.ProtoInt32(ob.OrederCount),
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTicker", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetTicker", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetTickerResponse{
		Security:   s2c.Security,
		Name:       util.ProtoStr(s2c.Name),
		TickerList: make([]*Ticker, 0, len(s2c.TickerList)),
	}

	for _, t := range s2c.TickerList {
		if t == nil {
			continue
		}
		result.TickerList = append(result.TickerList, &Ticker{
			Time:        util.ProtoStr(t.Time),
Sequence:     util.ProtoInt64(t.Sequence),
			Dir:         util.ProtoInt32(t.Dir),
			Price:       util.ProtoFloat64(t.Price),
			Volume:      util.ProtoInt64(t.Volume),
			Turnover:    util.ProtoFloat64(t.Turnover),
			RecvTime:    util.ProtoFloat64(t.RecvTime),
			Type:        util.ProtoInt32(t.Type),
			TypeSign:    util.ProtoInt32(t.TypeSign),
			Timestamp:   util.ProtoFloat64(t.Timestamp),
			PushDataType: util.ProtoInt32(t.PushDataType),
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetRT", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetRT", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetRTResponse{
		Security: s2c.Security,
		Name:     util.ProtoStr(s2c.Name),
		RTList:   make([]*RT, 0, len(s2c.RtList)),
	}

	for _, rt := range s2c.RtList {
		if rt == nil {
			continue
		}
		result.RTList = append(result.RTList, &RT{
			Time:           util.ProtoStr(rt.Time),
			Minute:         util.ProtoInt32(rt.Minute),
			IsBlank:        util.ProtoBool(rt.IsBlank),
			Price:          util.ProtoFloat64(rt.Price),
			LastClosePrice: util.ProtoFloat64(rt.LastClosePrice),
			AvgPrice:       util.ProtoFloat64(rt.AvgPrice),
			Volume:         util.ProtoInt64(rt.Volume),
			Turnover:       util.ProtoFloat64(rt.Turnover),
			Timestamp:      util.ProtoFloat64(rt.Timestamp),
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetBroker", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetBroker", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetBrokerResponse{
		Security:      s2c.Security,
		Name:          util.ProtoStr(s2c.Name),
		AskBrokerList: make([]*Broker, 0, len(s2c.BrokerAskList)),
		BidBrokerList: make([]*Broker, 0, len(s2c.BrokerBidList)),
	}

	for _, b := range s2c.BrokerAskList {
		if b == nil {
			continue
		}
		result.AskBrokerList = append(result.AskBrokerList, &Broker{
			ID:      util.ProtoInt64(b.Id),
			Name:    util.ProtoStr(b.Name),
			Pos:     util.ProtoInt32(b.Pos),
			Volume:  util.ProtoInt64(b.Volume),
			OrderID: util.ProtoInt64(b.OrderID),
		})
	}

	for _, b := range s2c.BrokerBidList {
		if b == nil {
			continue
		}
		result.BidBrokerList = append(result.BidBrokerList, &Broker{
			ID:      util.ProtoInt64(b.Id),
			Name:    util.ProtoStr(b.Name),
			Pos:     util.ProtoInt32(b.Pos),
			Volume:  util.ProtoInt64(b.Volume),
			OrderID: util.ProtoInt64(b.OrderID),
		})
	}

	return result, nil
}
