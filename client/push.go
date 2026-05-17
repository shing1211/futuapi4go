package client

import "github.com/shing1211/futuapi4go/pkg/push"

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
		Market:       data.Security.GetMarket(),
		Code:         data.Security.GetCode(),
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
	return &PushKLine{
		Market:    data.Security.GetMarket(),
		Code:      data.Security.GetCode(),
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
		Market:         data.Security.GetMarket(),
		Code:           data.Security.GetCode(),
		Name:           data.Name,
		SvrRecvTimeBid: data.SvrRecvTimeBid,
		SvrRecvTimeAsk: data.SvrRecvTimeAsk,
	}
	for _, b := range data.OrderBookBidList {
		ob.Bids = append(ob.Bids, OBItem{
			Price:      b.GetPrice(),
			Volume:     b.GetVolume(),
			OrderCount: int64(b.GetOrederCount()),
		})
	}
	for _, a := range data.OrderBookAskList {
		ob.Asks = append(ob.Asks, OBItem{
			Price:      a.GetPrice(),
			Volume:     a.GetVolume(),
			OrderCount: int64(a.GetOrederCount()),
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
		Market:       data.Security.GetMarket(),
		Code:         data.Security.GetCode(),
		Name:         data.Name,
		Time:         t.GetTime(),
		Price:        t.GetPrice(),
		Volume:       t.GetVolume(),
		Turnover:     t.GetTurnover(),
		Side:         t.GetDir(),
		Sequence:     t.GetSequence(),
		Dir:          t.GetDir(),
		RecvTime:     t.GetRecvTime(),
		Type:         t.GetType(),
		TypeSign:     t.GetTypeSign(),
		Timestamp:    t.GetTimestamp(),
		PushDataType: t.GetPushDataType(),
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
		Market:        data.Security.GetMarket(),
		Code:          data.Security.GetCode(),
		Name:          data.Name,
		Time:          rt.GetTime(),
		Price:         rt.GetPrice(),
		Volume:        rt.GetVolume(),
		AvgPrice:       rt.GetAvgPrice(),
		Turnover:      rt.GetTurnover(),
		Minute:        rt.GetMinute(),
		IsBlank:       rt.GetIsBlank(),
		Timestamp:     rt.GetTimestamp(),
		LastClosePrice: rt.GetLastClosePrice(),
	}, nil
}

// ParsePushBroker parses a raw push body (ProtoID 3015) into PushBroker data.
func ParsePushBroker(body []byte) (*PushBroker, error) {
	data, err := push.ParseUpdateBroker(body)
	if err != nil || data == nil {
		return nil, err
	}
	ob := &PushBroker{
		Market: data.Security.GetMarket(),
		Code:   data.Security.GetCode(),
		Name:   data.Name,
	}
	for _, a := range data.AskBrokerList {
		ob.Asks = append(ob.Asks, BrokerItem{
			Volume:   a.GetVolume(),
			BrokerID: int32(a.GetId()),
			Name:     a.GetName(),
			Pos:      a.GetPos(),
		})
	}
	for _, b := range data.BidBrokerList {
		ob.Bids = append(ob.Bids, BrokerItem{
			Volume:   b.GetVolume(),
			BrokerID: int32(b.GetId()),
			Name:     b.GetName(),
			Pos:      b.GetPos(),
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
		Market:       data.Security.GetMarket(),
		Code:         data.Security.GetCode(),
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
		AccID:     data.Header.GetAccID(),
		TrdEnv:    data.Header.GetTrdEnv(),
		TrdMarket: data.Header.GetTrdMarket(),
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
		OrderID:     data.Order.GetOrderID(),
		OrderIDEx:   data.Order.GetOrderIDEx(),
		Code:        data.Order.GetCode(),
		SecMarket:   data.Order.GetSecMarket(),
		TrdSide:     data.Order.GetTrdSide(),
		Qty:         data.Order.GetQty(),
		Price:       data.Order.GetPrice(),
		OrderStatus: data.Order.GetOrderStatus(),
	}, nil
}

// ParsePushOrderFill parses a raw push body (ProtoID 2218) into PushOrderFill.
func ParsePushOrderFill(body []byte) (*PushOrderFill, error) {
	data, err := push.ParseUpdateOrderFill(body)
	if err != nil || data == nil || data.OrderFill == nil {
		return nil, err
	}
	return &PushOrderFill{
		OrderID:        data.OrderFill.GetOrderID(),
		OrderIDEx:      data.OrderFill.GetOrderIDEx(),
		Code:           data.OrderFill.GetCode(),
		SecMarket:      data.OrderFill.GetSecMarket(),
		TrdSide:        data.OrderFill.GetTrdSide(),
		Qty:            data.OrderFill.GetQty(),
		Price:          data.OrderFill.GetPrice(),
		FillID:         data.OrderFill.GetFillID(),
		FillIDEx:       data.OrderFill.GetFillIDEx(),
		FillCreateTime: data.OrderFill.GetCreateTime(),
	}, nil
}
