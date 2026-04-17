// Package push provides handlers for parsing push notification payloads
// from Futu OpenD. Use RegisterHandler on the client to receive real-time
// market data and order updates.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/push"
//
//	cli.RegisterHandler(push.ProtoID_Qot_UpdateBasicQot, func(protoID uint32, body []byte) {
//	    data, err := push.ParseUpdateBasicQot(body)
//	    // ...
//	})
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

// Note: push constants use the same ProtoID values as the corresponding
// request APIs (e.g., ProtoID_Qot_UpdateBasicQot = 3005). The push
// notification arrives on the same connection after subscribing.
package push

import (
	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatekl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatepricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateticker"
)

const (
	ProtoID_Qot_UpdateBasicQot    = 3005
	ProtoID_Qot_UpdateKL          = 3007
	ProtoID_Qot_UpdateOrderBook   = 3013
	ProtoID_Qot_UpdateTicker      = 3011
	ProtoID_Qot_UpdateRT          = 3009
	ProtoID_Qot_UpdateBroker      = 3015
	ProtoID_Qot_PushPriceReminder = 3107
)

type UpdateBasicQot struct {
	Security  *qotcommon.Security
	Name      string
	CurPrice  float64
	OpenPrice float64
	HighPrice float64
	LowPrice  float64
	Volume    int64
	Turnover  float64
}

func ParseUpdateBasicQot(body []byte) (*UpdateBasicQot, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatebasicqot.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil || len(s2c.GetBasicQotList()) == 0 {
		return nil, nil
	}
	bq := s2c.GetBasicQotList()[0]
	return &UpdateBasicQot{
		Security:  bq.GetSecurity(),
		Name:      bq.GetName(),
		CurPrice:  bq.GetCurPrice(),
		OpenPrice: bq.GetOpenPrice(),
		HighPrice: bq.GetHighPrice(),
		LowPrice:  bq.GetLowPrice(),
		Volume:    bq.GetVolume(),
		Turnover:  bq.GetTurnover(),
	}, nil
}

type UpdateKL struct {
	RehabType int32
	KlType    int32
	Security  *qotcommon.Security
	Name      string
	KLList    []*qotcommon.KLine
}

func ParseUpdateKL(body []byte) (*UpdateKL, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatekl.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil || s2c.KlList == nil {
		return nil, nil
	}
	return &UpdateKL{
		RehabType: s2c.GetRehabType(),
		KlType:    s2c.GetKlType(),
		Security:  s2c.GetSecurity(),
		Name:      s2c.GetName(),
		KLList:    s2c.GetKlList(),
	}, nil
}

type UpdateOrderBook struct {
	Security                *qotcommon.Security
	Name                    string
	OrderBookAskList        []*qotcommon.OrderBook
	OrderBookBidList        []*qotcommon.OrderBook
	SvrRecvTimeBid          string
	SvrRecvTimeBidTimestamp float64
	SvrRecvTimeAsk          string
	SvrRecvTimeAskTimestamp float64
}

func ParseUpdateOrderBook(body []byte) (*UpdateOrderBook, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateorderbook.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil {
		return nil, nil
	}
	return &UpdateOrderBook{
		Security:                s2c.GetSecurity(),
		Name:                    s2c.GetName(),
		OrderBookAskList:        s2c.GetOrderBookAskList(),
		OrderBookBidList:        s2c.GetOrderBookBidList(),
		SvrRecvTimeBid:          s2c.GetSvrRecvTimeBid(),
		SvrRecvTimeBidTimestamp: s2c.GetSvrRecvTimeBidTimestamp(),
		SvrRecvTimeAsk:          s2c.GetSvrRecvTimeAsk(),
		SvrRecvTimeAskTimestamp: s2c.GetSvrRecvTimeAskTimestamp(),
	}, nil
}

type UpdateTicker struct {
	Security   *qotcommon.Security
	Name       string
	TickerList []*qotcommon.Ticker
}

func ParseUpdateTicker(body []byte) (*UpdateTicker, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateticker.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil || len(s2c.GetTickerList()) == 0 {
		return nil, nil
	}
	return &UpdateTicker{
		Security:   s2c.GetSecurity(),
		Name:       s2c.GetName(),
		TickerList: s2c.GetTickerList(),
	}, nil
}

type UpdateRT struct {
	Security *qotcommon.Security
	Name     string
	RTList   []*qotcommon.TimeShare
}

func ParseUpdateRT(body []byte) (*UpdateRT, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatert.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil || len(s2c.GetRtList()) == 0 {
		return nil, nil
	}
	return &UpdateRT{
		Security: s2c.GetSecurity(),
		Name:     s2c.GetName(),
		RTList:   s2c.GetRtList(),
	}, nil
}

type UpdateBroker struct {
	Security      *qotcommon.Security
	Name          string
	AskBrokerList []*qotcommon.Broker
	BidBrokerList []*qotcommon.Broker
}

func ParseUpdateBroker(body []byte) (*UpdateBroker, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatebroker.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil {
		return nil, nil
	}
	return &UpdateBroker{
		Security:      s2c.GetSecurity(),
		Name:          s2c.GetName(),
		AskBrokerList: s2c.GetBrokerAskList(),
		BidBrokerList: s2c.GetBrokerBidList(),
	}, nil
}

type UpdatePriceReminder struct {
	Security     *qotcommon.Security
	Name         string
	Price        float64
	ChangeRate   float64
	MarketStatus int32
	Content      string
	Note         string
	Key          int64
	Type         int32
	SetValue     float64
	CurValue     float64
}

func ParseUpdatePriceReminder(body []byte) (*UpdatePriceReminder, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatepricereminder.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.GetS2C()
	if s2c == nil {
		return nil, nil
	}
	return &UpdatePriceReminder{
		Security:     s2c.GetSecurity(),
		Name:         s2c.GetName(),
		Price:        s2c.GetPrice(),
		ChangeRate:   s2c.GetChangeRate(),
		MarketStatus: s2c.GetMarketStatus(),
		Content:      s2c.GetContent(),
		Note:         s2c.GetNote(),
		Key:          s2c.GetKey(),
		Type:         s2c.GetType(),
		SetValue:     s2c.GetSetValue(),
		CurValue:     s2c.GetCurValue(),
	}, nil
}
