package push

import (
	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pb/qotupdatebasicqot"
	"gitee.com/shing1211/futuapi4go/pb/qotupdatebroker"
	"gitee.com/shing1211/futuapi4go/pb/qotupdatekl"
	"gitee.com/shing1211/futuapi4go/pb/qotupdateorderbook"
	"gitee.com/shing1211/futuapi4go/pb/qotupdatert"
	"gitee.com/shing1211/futuapi4go/pb/qotupdateticker"
)

const (
	ProtoID_Qot_UpdateBasicQot  = 3101
	ProtoID_Qot_UpdateKL        = 3102
	ProtoID_Qot_UpdateOrderBook = 3103
	ProtoID_Qot_UpdateTicker    = 3104
	ProtoID_Qot_UpdateRT        = 3105
	ProtoID_Qot_UpdateBroker    = 3106
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
	var rsp qotupdatebasicqot.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	if len(rsp.GetBasicQotList()) == 0 {
		return nil, nil
	}
	bq := rsp.GetBasicQotList()[0]
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
	var rsp qotupdatekl.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateKL{
		RehabType: rsp.GetRehabType(),
		KlType:    rsp.GetKlType(),
		Security:  rsp.GetSecurity(),
		Name:      rsp.GetName(),
		KLList:    rsp.GetKlList(),
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
