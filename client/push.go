package client

import (
	"fmt"

	"github.com/shing1211/futuapi4go/pkg/push"
)

// Push ProtoID constants (re-exported from pkg/push for convenience).
const (
	ProtoID_Qot_UpdateBasicQot      = 3005
	ProtoID_Qot_UpdateKL            = 3007
	ProtoID_Qot_UpdateOrderBook     = 3013
	ProtoID_Qot_UpdateTicker        = 3011
	ProtoID_Qot_UpdateRT            = 3009
	ProtoID_Qot_UpdateBroker        = 3015
	ProtoID_Qot_UpdatePriceReminder = 3019
	ProtoID_Trd_UpdateOrder         = 2208
	ProtoID_Trd_UpdateOrderFill     = 2218
	ProtoID_Trd_Notify              = 2207
)

// ParsePushQuote parses a raw push body (ProtoID 3005) into a PushQuote.
func ParsePushQuote(body []byte) (*PushQuote, error) {
	data, err := push.ParseUpdateBasicQot(body)
	if err != nil || data == nil {
		return nil, err
	}
	return &PushQuote{
		Market:       getInt32(data.Security.Market),
		Code:         getStr(data.Security.Code),
		Name:         data.Name,
		CurPrice:     data.CurPrice,
		OpenPrice:    data.OpenPrice,
		HighPrice:    data.HighPrice,
		LowPrice:     data.LowPrice,
		Volume:       data.Volume,
		Turnover:     data.Turnover,
		LastClose:    data.LastClosePrice,
		TurnoverRate: data.TurnoverRate,
		Amplitude:    data.Amplitude,
		IsSuspended:  data.IsSuspended,
		SecStatus:    data.SecStatus,
	}, nil
}

// ParsePushKLine parses a raw push body (ProtoID 3007) into a PushKLine.
func ParsePushKLine(body []byte) (*PushKLine, error) {
	data, err := push.ParseUpdateKL(body)
	if err != nil || data == nil || len(data.KLList) == 0 {
		return nil, err
	}
	kl := data.KLList[0]
	if kl == nil {
		return nil, fmt.Errorf("ParsePushKLine: first KLList element is nil")
	}
	return &PushKLine{
		Market:    getInt32(data.Security.Market),
		Code:      getStr(data.Security.Code),
		Name:      data.Name,
		KLType:    data.KlType,
		RehabType: data.RehabType,
		KLine: KLine{
			Time:         kl.Time,
			IsBlank:      kl.IsBlank,
			Open:         kl.OpenPrice,
			High:         kl.HighPrice,
			Low:          kl.LowPrice,
			Close:        kl.ClosePrice,
			Volume:       kl.Volume,
			LastClose:    kl.LastClosePrice,
			Turnover:     kl.Turnover,
			TurnoverRate: kl.TurnoverRate,
			Pe:           kl.Pe,
			ChangeRate:   kl.ChangeRate,
			Timestamp:    kl.Timestamp,
		},
	}, nil
}

// ParsePushOrderBook parses a raw push body (ProtoID 3013) into a PushOrderBook.
func ParsePushOrderBook(body []byte) (*PushOrderBook, error) {
	data, err := push.ParseUpdateOrderBook(body)
	if err != nil || data == nil {
		return nil, err
	}
	ob := &PushOrderBook{
		Market:         getInt32(data.Security.Market),
		Code:           getStr(data.Security.Code),
		Name:           data.Name,
		SvrRecvTimeBid: data.SvrRecvTimeBid,
		SvrRecvTimeAsk: data.SvrRecvTimeAsk,
	}
	for _, b := range data.OrderBookBidList {
		ob.Bids = append(ob.Bids, OBItem{
			Price:      getFloat64(b.Price),
			Volume:     getInt64(b.Volume),
			OrderCount: int64(getInt32(b.OrederCount)),
		})
	}
	for _, a := range data.OrderBookAskList {
		ob.Asks = append(ob.Asks, OBItem{
			Price:      getFloat64(a.Price),
			Volume:     getInt64(a.Volume),
			OrderCount: int64(getInt32(a.OrederCount)),
		})
	}
	return ob, nil
}

// ParsePushTicker parses a raw push body (ProtoID 3011) into a PushTicker.
func ParsePushTicker(body []byte) (*PushTicker, error) {
	data, err := push.ParseUpdateTicker(body)
	if err != nil || data == nil || len(data.TickerList) == 0 {
		return nil, err
	}
	t := data.TickerList[0]
	return &PushTicker{
		Market:       getInt32(data.Security.Market),
		Code:         getStr(data.Security.Code),
		Name:         data.Name,
		Time:         getStr(t.Time),
		Price:        getFloat64(t.Price),
		Volume:       getInt64(t.Volume),
		Turnover:     getFloat64(t.Turnover),
		Side:         getInt32(t.Dir),
		Sequence:     getInt64(t.Sequence),
		Dir:          getInt32(t.Dir),
		RecvTime:     getFloat64(t.RecvTime),
		Type:         getInt32(t.Type),
		TypeSign:     getInt32(t.TypeSign),
		Timestamp:    getFloat64(t.Timestamp),
		PushDataType: getInt32(t.PushDataType),
	}, nil
}

// ParsePushRT parses a raw push body (ProtoID 3009) into PushRT data.
func ParsePushRT(body []byte) (*PushRT, error) {
	data, err := push.ParseUpdateRT(body)
	if err != nil || data == nil || len(data.RTList) == 0 {
		return nil, err
	}
	rt := data.RTList[0]
	return &PushRT{
		Market:        getInt32(data.Security.Market),
		Code:          getStr(data.Security.Code),
		Name:          data.Name,
		Time:          getStr(rt.Time),
		Price:         getFloat64(rt.Price),
		Volume:        getInt64(rt.Volume),
		AvgPrice:       getFloat64(rt.AvgPrice),
		Turnover:      getFloat64(rt.Turnover),
		Minute:        getInt32(rt.Minute),
		IsBlank:       getBool(rt.IsBlank),
		Timestamp:     getFloat64(rt.Timestamp),
		LastClosePrice: getFloat64(rt.LastClosePrice),
	}, nil
}

// ParsePushBroker parses a raw push body (ProtoID 3015) into PushBroker data.
func ParsePushBroker(body []byte) (*PushBroker, error) {
	data, err := push.ParseUpdateBroker(body)
	if err != nil || data == nil {
		return nil, err
	}
	ob := &PushBroker{
		Market: getInt32(data.Security.Market),
		Code:   getStr(data.Security.Code),
		Name:   data.Name,
	}
	for _, a := range data.AskBrokerList {
		ob.Asks = append(ob.Asks, BrokerItem{
			Volume:   getInt64(a.Volume),
			BrokerID: int32(getInt64(a.Id)),
			Name:     getStr(a.Name),
			Pos:      getInt32(a.Pos),
		})
	}
	for _, b := range data.BidBrokerList {
		ob.Bids = append(ob.Bids, BrokerItem{
			Volume:   getInt64(b.Volume),
			BrokerID: int32(getInt64(b.Id)),
			Name:     getStr(b.Name),
			Pos:      getInt32(b.Pos),
		})
	}
	return ob, nil
}

// ParsePushPriceReminder parses a raw push body (ProtoID 3019) into PushPriceReminder.
func ParsePushPriceReminder(body []byte) (*PushPriceReminder, error) {
	data, err := push.ParseUpdatePriceReminder(body)
	if err != nil || data == nil {
		return nil, err
	}
	return &PushPriceReminder{
		Market:       getInt32(data.Security.Market),
		Code:         getStr(data.Security.Code),
		Name:         data.Name,
		Price:        data.Price,
		ChangeRate:   data.ChangeRate,
		MarketStatus: data.MarketStatus,
		Content:      data.Content,
		Note:         data.Note,
		Key:          data.Key,
		Type:         data.Type,
		SetValue:     data.SetValue,
		CurValue:     data.CurValue,
	}, nil
}

// ParsePushTrdNotify parses a raw push body (ProtoID 2207) into PushTrdNotify.
func ParsePushTrdNotify(body []byte) (*PushTrdNotify, error) {
	data, err := push.ParseTrdNotify(body)
	if err != nil || data == nil {
		return nil, err
	}
	return &PushTrdNotify{
		AccID:     getUint64(data.Header.AccID),
		TrdEnv:    getInt32(data.Header.TrdEnv),
		TrdMarket: getInt32(data.Header.TrdMarket),
		Type:      data.Type,
	}, nil
}

// ParsePushOrderUpdate parses a raw push body (ProtoID 2208) into PushOrderUpdate.
func ParsePushOrderUpdate(body []byte) (*PushOrderUpdate, error) {
	data, err := push.ParseUpdateOrder(body)
	if err != nil || data == nil || data.Order == nil {
		return nil, err
	}
	return &PushOrderUpdate{
		OrderID:     getUint64(data.Order.OrderID),
		OrderIDEx:   getStr(data.Order.OrderIDEx),
		Code:        getStr(data.Order.Code),
		SecMarket:   getInt32(data.Order.SecMarket),
		TrdSide:     getInt32(data.Order.TrdSide),
		Qty:         getFloat64(data.Order.Qty),
		Price:       getFloat64(data.Order.Price),
		OrderStatus: getInt32(data.Order.OrderStatus),
	}, nil
}

// ParsePushOrderFill parses a raw push body (ProtoID 2218) into PushOrderFill.
func ParsePushOrderFill(body []byte) (*PushOrderFill, error) {
	data, err := push.ParseUpdateOrderFill(body)
	if err != nil || data == nil || data.OrderFill == nil {
		return nil, err
	}
	return &PushOrderFill{
		OrderID:        getUint64(data.OrderFill.OrderID),
		OrderIDEx:      getStr(data.OrderFill.OrderIDEx),
		Code:           getStr(data.OrderFill.Code),
		SecMarket:      getInt32(data.OrderFill.SecMarket),
		TrdSide:        getInt32(data.OrderFill.TrdSide),
		Qty:            getFloat64(data.OrderFill.Qty),
		Price:          getFloat64(data.OrderFill.Price),
		FillID:         getUint64(data.OrderFill.FillID),
		FillIDEx:       getStr(data.OrderFill.FillIDEx),
		FillCreateTime: getStr(data.OrderFill.CreateTime),
	}, nil
}
