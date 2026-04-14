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
	var rsp qotupdatebasicqot.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.GetS2C()
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
	var rsp qotupdatekl.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.GetS2C()
	if s2c == nil {
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
	var rsp qotupdateorderbook.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateOrderBook{
		Security:                rsp.GetSecurity(),
		Name:                    rsp.GetName(),
		OrderBookAskList:        rsp.GetOrderBookAskList(),
		OrderBookBidList:        rsp.GetOrderBookBidList(),
		SvrRecvTimeBid:          rsp.GetSvrRecvTimeBid(),
		SvrRecvTimeBidTimestamp: rsp.GetSvrRecvTimeBidTimestamp(),
		SvrRecvTimeAsk:          rsp.GetSvrRecvTimeAsk(),
		SvrRecvTimeAskTimestamp: rsp.GetSvrRecvTimeAskTimestamp(),
	}, nil
}

type UpdateTicker struct {
	Security   *qotcommon.Security
	Name       string
	TickerList []*qotcommon.Ticker
}

func ParseUpdateTicker(body []byte) (*UpdateTicker, error) {
	var rsp qotupdateticker.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateTicker{
		Security:   rsp.GetSecurity(),
		Name:       rsp.GetName(),
		TickerList: rsp.GetTickerList(),
	}, nil
}

type UpdateRT struct {
	Security *qotcommon.Security
	Name     string
	RTList   []*qotcommon.TimeShare
}

func ParseUpdateRT(body []byte) (*UpdateRT, error) {
	var rsp qotupdatert.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateRT{
		Security: rsp.GetSecurity(),
		Name:     rsp.GetName(),
		RTList:   rsp.GetRtList(),
	}, nil
}

type UpdateBroker struct {
	Security      *qotcommon.Security
	Name          string
	AskBrokerList []*qotcommon.Broker
	BidBrokerList []*qotcommon.Broker
}

func ParseUpdateBroker(body []byte) (*UpdateBroker, error) {
	var rsp qotupdatebroker.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateBroker{
		Security:      rsp.GetSecurity(),
		Name:          rsp.GetName(),
		AskBrokerList: rsp.GetBrokerAskList(),
		BidBrokerList: rsp.GetBrokerBidList(),
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
	var rsp qotupdatepricereminder.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdatePriceReminder{
		Security:     rsp.GetSecurity(),
		Name:         rsp.GetName(),
		Price:        rsp.GetPrice(),
		ChangeRate:   rsp.GetChangeRate(),
		MarketStatus: rsp.GetMarketStatus(),
		Content:      rsp.GetContent(),
		Note:         rsp.GetNote(),
		Key:          rsp.GetKey(),
		Type:         rsp.GetType(),
		SetValue:     rsp.GetSetValue(),
		CurValue:     rsp.GetCurValue(),
	}, nil
}
