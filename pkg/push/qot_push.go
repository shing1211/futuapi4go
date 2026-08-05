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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractkline"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotpushindicatorcalc"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateeventcontractkline"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateeventcontractorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateeventcontractticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatekl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateoptionevent"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatepricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateticker"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_Qot_UpdateBasicQot      = 3005
	ProtoID_Qot_UpdateKL            = 3007
	ProtoID_Qot_UpdateOrderBook     = 3013
	ProtoID_Qot_UpdateTicker        = 3011
	ProtoID_Qot_UpdateRT            = 3009
	ProtoID_Qot_UpdateBroker        = 3015
	ProtoID_Qot_UpdatePriceReminder = 3019

	// v10.9+ Event Contract push notifications.
	ProtoID_Qot_UpdateEventContractOrderBook = 3450
	ProtoID_Qot_UpdateEventContractKline     = 3451
	ProtoID_Qot_UpdateEventContractTicker    = 3452
	ProtoID_Qot_UpdateOptionEvent            = 3310
	ProtoID_Qot_PushIndicatorCalc            = 3261
)

// UpdateBasicQot represents a real-time basic quote push notification.
type UpdateBasicQot struct {
	Security        *qotcommon.Security
	Name            string
	CurPrice        float64
	OpenPrice       float64
	HighPrice       float64
	LowPrice        float64
	Volume          int64
	Turnover        float64
	IsSuspended     bool
	LastClosePrice  float64
	UpdateTime      string
	UpdateTimestamp float64
	ListTime        string
	PriceSpread     float64
	TurnoverRate    float64
	Amplitude       float64
	DarkStatus      int32
	OptionExData    *qotcommon.OptionBasicQotExData
	ListTimestamp   float64
	PreMarket       *qotcommon.PreAfterMarketData
	AfterMarket     *qotcommon.PreAfterMarketData
	SecStatus       int32
	FutureExData    *qotcommon.FutureBasicQotExData
	WarrantExData   *qotcommon.WarrantBasicQotExData
	Overnight       *qotcommon.PreAfterMarketData
}

// ParseUpdateBasicQot parses a basic quote push notification from a raw protobuf body.
func ParseUpdateBasicQot(body []byte) (*UpdateBasicQot, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatebasicqot.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil || len(s2c.BasicQotList) == 0 {
		return nil, nil
	}
	bq := s2c.BasicQotList[0]
	return &UpdateBasicQot{
		Security:        bq.Security,
		Name:            util.ProtoStr(bq.Name),
		CurPrice:        util.ProtoFloat64(bq.CurPrice),
		OpenPrice:       util.ProtoFloat64(bq.OpenPrice),
		HighPrice:       util.ProtoFloat64(bq.HighPrice),
		LowPrice:        util.ProtoFloat64(bq.LowPrice),
		Volume:          util.ProtoInt64(bq.Volume),
		Turnover:        util.ProtoFloat64(bq.Turnover),
		IsSuspended:     util.ProtoBool(bq.IsSuspended),
		LastClosePrice:  util.ProtoFloat64(bq.LastClosePrice),
		UpdateTime:      util.ProtoStr(bq.UpdateTime),
		UpdateTimestamp: util.ProtoFloat64(bq.UpdateTimestamp),
		ListTime:        util.ProtoStr(bq.ListTime),
		PriceSpread:     util.ProtoFloat64(bq.PriceSpread),
		TurnoverRate:    util.ProtoFloat64(bq.TurnoverRate),
		Amplitude:       util.ProtoFloat64(bq.Amplitude),
		DarkStatus:      util.ProtoInt32(bq.DarkStatus),
		OptionExData:    bq.OptionExData,
		ListTimestamp:   util.ProtoFloat64(bq.ListTimestamp),
		PreMarket:       bq.PreMarket,
		AfterMarket:     bq.AfterMarket,
		SecStatus:       util.ProtoInt32(bq.SecStatus),
		FutureExData:    bq.FutureExData,
		WarrantExData:   bq.WarrantExData,
		Overnight:       bq.Overnight,
	}, nil
}

// PushKLine represents a single K-line bar from push notification.
type PushKLine struct {
	Time           string
	IsBlank        bool
	HighPrice      float64
	OpenPrice      float64
	LowPrice       float64
	ClosePrice     float64
	LastClosePrice float64
	Volume         int64
	Turnover       float64
	TurnoverRate   float64
	Pe             float64
	ChangeRate     float64
	Timestamp      float64
}

// UpdateKL represents a K-line push notification.
type UpdateKL struct {
	RehabType int32
	KlType    int32
	Security  *qotcommon.Security
	Name      string
	KLList    []*PushKLine
}

// ParseUpdateKL parses a K-line push notification from a raw protobuf body.
func ParseUpdateKL(body []byte) (*UpdateKL, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatekl.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil || s2c.KlList == nil {
		return nil, nil
	}
	s2cList := s2c.KlList
	klList := make([]*PushKLine, 0, len(s2cList))
	for _, kl := range s2cList {
		if kl == nil {
			continue
		}
		klList = append(klList, &PushKLine{
			Time:           util.ProtoStr(kl.Time),
			IsBlank:        util.ProtoBool(kl.IsBlank),
			HighPrice:      util.ProtoFloat64(kl.HighPrice),
			OpenPrice:      util.ProtoFloat64(kl.OpenPrice),
			LowPrice:       util.ProtoFloat64(kl.LowPrice),
			ClosePrice:     util.ProtoFloat64(kl.ClosePrice),
			LastClosePrice: util.ProtoFloat64(kl.LastClosePrice),
			Volume:         util.ProtoInt64(kl.Volume),
			Turnover:       util.ProtoFloat64(kl.Turnover),
			TurnoverRate:   util.ProtoFloat64(kl.TurnoverRate),
			Pe:             util.ProtoFloat64(kl.Pe),
			ChangeRate:     util.ProtoFloat64(kl.ChangeRate),
			Timestamp:      util.ProtoFloat64(kl.Timestamp),
		})
	}
	return &UpdateKL{
		RehabType: util.ProtoInt32(s2c.RehabType),
		KlType:    util.ProtoInt32(s2c.KlType),
		Security:  s2c.Security,
		Name:      util.ProtoStr(s2c.Name),
		KLList:    klList,
	}, nil
}

// UpdateOrderBook represents an order book push notification.
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

// ParseUpdateOrderBook parses an order book push notification from a raw protobuf body.
func ParseUpdateOrderBook(body []byte) (*UpdateOrderBook, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateorderbook.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateOrderBook{
		Security:                s2c.Security,
		Name:                    util.ProtoStr(s2c.Name),
		OrderBookAskList:        s2c.OrderBookAskList,
		OrderBookBidList:        s2c.OrderBookBidList,
		SvrRecvTimeBid:          util.ProtoStr(s2c.SvrRecvTimeBid),
		SvrRecvTimeBidTimestamp: util.ProtoFloat64(s2c.SvrRecvTimeBidTimestamp),
		SvrRecvTimeAsk:          util.ProtoStr(s2c.SvrRecvTimeAsk),
		SvrRecvTimeAskTimestamp: util.ProtoFloat64(s2c.SvrRecvTimeAskTimestamp),
	}, nil
}

// UpdateTicker represents a ticker push notification.
type UpdateTicker struct {
	Security   *qotcommon.Security
	Name       string
	TickerList []*qotcommon.Ticker
}

// ParseUpdateTicker parses a ticker push notification from a raw protobuf body.
func ParseUpdateTicker(body []byte) (*UpdateTicker, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateticker.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil || len(s2c.TickerList) == 0 {
		return nil, nil
	}
	return &UpdateTicker{
		Security:   s2c.Security,
		Name:       util.ProtoStr(s2c.Name),
		TickerList: s2c.TickerList,
	}, nil
}

// UpdateRT represents a real-time data push notification.
type UpdateRT struct {
	Security *qotcommon.Security
	Name     string
	RTList   []*qotcommon.TimeShare
}

// ParseUpdateRT parses a real-time data push notification from a raw protobuf body.
func ParseUpdateRT(body []byte) (*UpdateRT, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatert.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil || len(s2c.RtList) == 0 {
		return nil, nil
	}
	return &UpdateRT{
		Security: s2c.Security,
		Name:     util.ProtoStr(s2c.Name),
		RTList:   s2c.RtList,
	}, nil
}

// UpdateBroker represents a broker queue push notification.
type UpdateBroker struct {
	Security      *qotcommon.Security
	Name          string
	AskBrokerList []*qotcommon.Broker
	BidBrokerList []*qotcommon.Broker
}

// ParseUpdateBroker parses a broker queue push notification from a raw protobuf body.
func ParseUpdateBroker(body []byte) (*UpdateBroker, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatebroker.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateBroker{
		Security:      s2c.Security,
		Name:          util.ProtoStr(s2c.Name),
		AskBrokerList: s2c.BrokerAskList,
		BidBrokerList: s2c.BrokerBidList,
	}, nil
}

// UpdatePriceReminder represents a price reminder push notification.
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

// ParseUpdatePriceReminder parses a price reminder push notification from a raw protobuf body.
func ParseUpdatePriceReminder(body []byte) (*UpdatePriceReminder, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdatepricereminder.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdatePriceReminder{
		Security:     s2c.Security,
		Name:         util.ProtoStr(s2c.Name),
		Price:        util.ProtoFloat64(s2c.Price),
		ChangeRate:   util.ProtoFloat64(s2c.ChangeRate),
		MarketStatus: util.ProtoInt32(s2c.MarketStatus),
		Content:      util.ProtoStr(s2c.Content),
		Note:         util.ProtoStr(s2c.Note),
		Key:          util.ProtoInt64(s2c.Key),
		Type:         util.ProtoInt32(s2c.Type),
		SetValue:     util.ProtoFloat64(s2c.SetValue),
		CurValue:     util.ProtoFloat64(s2c.CurValue),
	}, nil
}

// UpdateEventContractOrderBook represents an Event Contract order book push
// notification. Each item contains YES/NO bid/ask levels.
type UpdateEventContractOrderBook struct {
	OrderBookList []*qotgeteventcontractorderbook.OrderBookItem
}

// ParseUpdateEventContractOrderBook parses an EC order book push notification
// from a raw protobuf body.
func ParseUpdateEventContractOrderBook(body []byte) (*UpdateEventContractOrderBook, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateeventcontractorderbook.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateEventContractOrderBook{
		OrderBookList: s2c.OrderBookList,
	}, nil
}

// UpdateEventContractKline represents an Event Contract K-line push
// notification.
type UpdateEventContractKline struct {
	KlineList []*qotgeteventcontractkline.KlineItem
}

// ParseUpdateEventContractKline parses an EC K-line push notification from a
// raw protobuf body.
func ParseUpdateEventContractKline(body []byte) (*UpdateEventContractKline, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateeventcontractkline.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateEventContractKline{
		KlineList: s2c.KlineList,
	}, nil
}

// UpdateEventContractTicker represents an Event Contract tick-by-tick push
// notification.
type UpdateEventContractTicker struct {
	TickerList []*qotgeteventcontractticker.TickerItem
}

// ParseUpdateEventContractTicker parses an EC ticker push notification from a
// raw protobuf body.
func ParseUpdateEventContractTicker(body []byte) (*UpdateEventContractTicker, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateeventcontractticker.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateEventContractTicker{
		TickerList: s2c.TickerList,
	}, nil
}

// UpdateOptionEvent represents an option event alert push notification.
// Triggered server-side when an option event (e.g. listing, expiration,
// assignment) occurs.
type UpdateOptionEvent struct {
	Owner   *qotcommon.Security
	Option  *qotcommon.Security
	Message *string
}

// ParseUpdateOptionEvent parses an option event alert push notification from a
// raw protobuf body.
func ParseUpdateOptionEvent(body []byte) (*UpdateOptionEvent, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotupdateoptionevent.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateOptionEvent{
		Owner:   s2c.Owner,
		Option:  s2c.Option,
		Message: s2c.Message,
	}, nil
}

// PushIndicatorCalc represents an asynchronous indicator calculation result
// push notification, paired with a prior RequestIndicatorCalc(caclId) call.
type PushIndicatorCalc struct {
	CalcId     string
	Outputs    []*qotcommon.IndicatorOutputParam
	OutputRows []*qotpushindicatorcalc.IndicatorOutputRow
}

// ParsePushIndicatorCalc parses an indicator calculation push result from a
// raw protobuf body. The CalcId field corresponds to the value returned by
// RequestIndicatorCalc.
func ParsePushIndicatorCalc(body []byte) (*PushIndicatorCalc, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var resp qotpushindicatorcalc.Response
	if err := proto.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	s2c := resp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &PushIndicatorCalc{
		CalcId:     util.ProtoStr(s2c.CalcId),
		Outputs:    s2c.Outputs,
		OutputRows: s2c.OutputRows,
	}, nil
}
