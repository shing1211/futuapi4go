package simulator

import (
	"time"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/notify"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatekl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatepricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdateticker"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdnotify"
	"github.com/shing1211/futuapi4go/pkg/pb/trdupdateorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdupdateorderfill"
)

func (s *Server) RegisterPushHandlers() {
	s.RegisterHandler(3005, s.handlePushBasicQot)      // UpdateBasicQot
	s.RegisterHandler(3007, s.handlePushKL)            // UpdateKL
	s.RegisterHandler(3013, s.handlePushOrderBook)     // UpdateOrderBook
	s.RegisterHandler(3011, s.handlePushTicker)        // UpdateTicker
	s.RegisterHandler(3009, s.handlePushRT)            // UpdateRT
	s.RegisterHandler(3015, s.handlePushBroker)        // UpdateBroker
	s.RegisterHandler(3107, s.handlePushPriceReminder) // UpdatePriceReminder (unchanged)
	s.RegisterHandler(2208, s.handlePushOrder)         // UpdateOrder
	s.RegisterHandler(2218, s.handlePushOrderFill)     // UpdateOrderFill
	s.RegisterHandler(2207, s.handlePushTrdNotify)     // TrdNotify
	s.RegisterHandler(1003, s.handlePushNotify)
}

func (s *Server) handlePushBasicQot(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"
	price := 350.0
	high := price * 1.02
	low := price * 0.98
	open := price * 0.99
	vol := int64(1000000)
	turn := price * float64(vol)

	basicQot := &qotcommon.BasicQot{
		Security:  security,
		Name:      &name,
		CurPrice:  &price,
		HighPrice: &high,
		LowPrice:  &low,
		OpenPrice: &open,
		Volume:    &vol,
		Turnover:  &turn,
	}

	resp := &qotupdatebasicqot.Response{
		RetType: &retType,
		S2C:     &qotupdatebasicqot.S2C{BasicQotList: []*qotcommon.BasicQot{basicQot}},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushKL(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	ts := float64(now.Unix())
	open := 350.0
	close := 355.0
	high := 360.0
	low := 345.0
	vol := int64(100000)
	turn := 35000000.0

	kl := &qotcommon.KLine{
		Time:       &timeStr,
		Timestamp:  &ts,
		OpenPrice:  &open,
		ClosePrice: &close,
		HighPrice:  &high,
		LowPrice:   &low,
		Volume:     &vol,
		Turnover:   &turn,
	}

	resp := &qotupdatekl.Response{
		RetType: &retType,
		S2C:     &qotupdatekl.S2C{Security: security, Name: &name, KlList: []*qotcommon.KLine{kl}},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushOrderBook(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"

	askList := make([]*qotcommon.OrderBook, 0, 10)
	bidList := make([]*qotcommon.OrderBook, 0, 10)

	basePrice := 350.0
	for i := 0; i < 10; i++ {
		askPrice := basePrice + float64(i+1)*0.01
		bidPrice := basePrice - float64(i+1)*0.01
		askVol := int64(10000 - i*500)
		bidVol := int64(10000 - i*500)
		orderCount := int32(100 - i*5)

		askList = append(askList, &qotcommon.OrderBook{
			Price:       &askPrice,
			Volume:      &askVol,
			OrederCount: &orderCount,
		})
		bidList = append(bidList, &qotcommon.OrderBook{
			Price:       &bidPrice,
			Volume:      &bidVol,
			OrederCount: &orderCount,
		})
	}

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	ts := float64(now.Unix())

	resp := &qotupdateorderbook.Response{
		RetType: &retType,
		S2C: &qotupdateorderbook.S2C{
			Security:                security,
			Name:                    &name,
			OrderBookAskList:        askList,
			OrderBookBidList:        bidList,
			SvrRecvTimeBid:          &timeStr,
			SvrRecvTimeBidTimestamp: &ts,
			SvrRecvTimeAsk:          &timeStr,
			SvrRecvTimeAskTimestamp: &ts,
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushTicker(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	ts := float64(now.Unix())
	price := 350.0
	vol := int64(1000)
	turn := price * float64(vol)
	seq := int64(now.UnixNano())

	ticker := &qotcommon.Ticker{
		Time:      &timeStr,
		Timestamp: &ts,
		Sequence:  &seq,
		Dir:       func() *int32 { v := int32(1); return &v }(),
		Price:     &price,
		Volume:    &vol,
		Turnover:  &turn,
	}

	resp := &qotupdateticker.Response{
		RetType: &retType,
		S2C:     &qotupdateticker.S2C{Security: security, Name: &name, TickerList: []*qotcommon.Ticker{ticker}},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushRT(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	ts := float64(now.Unix())
	price := 350.0
	vol := int64(10000)
	turn := price * float64(vol)
	avgPrice := price * 0.999

	rt := &qotcommon.TimeShare{
		Time:      &timeStr,
		Timestamp: &ts,
		Price:     &price,
		Volume:    &vol,
		Turnover:  &turn,
		AvgPrice:  &avgPrice,
	}

	resp := &qotupdatert.Response{
		RetType: &retType,
		S2C:     &qotupdatert.S2C{Security: security, Name: &name, RtList: []*qotcommon.TimeShare{rt}},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushBroker(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"

	askList := make([]*qotcommon.Broker, 0, 10)
	bidList := make([]*qotcommon.Broker, 0, 10)

	for i := 0; i < 10; i++ {
		id := int64(1000 + i)
		pos := int32(i)
		vol := int64(100000 - i*5000)
		askName := "ASK"
		bidName := "BID"

		askList = append(askList, &qotcommon.Broker{
			Id: &id, Name: &askName, Pos: &pos, Volume: &vol,
		})
		bidList = append(bidList, &qotcommon.Broker{
			Id: &id, Name: &bidName, Pos: &pos, Volume: &vol,
		})
	}

	resp := &qotupdatebroker.Response{
		RetType: &retType,
		S2C: &qotupdatebroker.S2C{
			Security:      security,
			Name:          &name,
			BrokerAskList: askList,
			BrokerBidList: bidList,
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushPriceReminder(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	market := int32(1)
	code := "00700"
	security := &qotcommon.Security{Market: &market, Code: &code}
	name := "腾讯控股"
	price := 350.0
	changeRate := 2.5
	content := "价格达到目标"
	note := "模拟提醒"
	key := int64(123456)
	reminderType := int32(1)
	setValue := 360.0
	curValue := 350.0
	marketStatus := int32(1)

	resp := &qotupdatepricereminder.Response{
		RetType: &retType,
		S2C: &qotupdatepricereminder.S2C{
			Security:     security,
			Name:         &name,
			Price:        &price,
			ChangeRate:   &changeRate,
			MarketStatus: &marketStatus,
			Content:      &content,
			Note:         &note,
			Key:          &key,
			Type:         &reminderType,
			SetValue:     &setValue,
			CurValue:     &curValue,
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushOrder(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	accID := uint64(123456789)
	trdMarket := int32(1)
	orderID := uint64(1234567890)
	price := 350.0
	qty := 100.0
	orderStatus := int32(2)

	resp := &trdupdateorder.Response{
		RetType: &retType,
		S2C: &trdupdateorder.S2C{
			Header: &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
			Order:  &trdcommon.Order{OrderID: &orderID, Price: &price, Qty: &qty, OrderStatus: &orderStatus},
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushOrderFill(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	accID := uint64(123456789)
	trdMarket := int32(1)
	fillID := uint64(9876543210)
	orderID := uint64(1234567890)
	price := 350.0
	qty := 100.0

	resp := &trdupdateorderfill.Response{
		RetType: &retType,
		S2C: &trdupdateorderfill.S2C{
			Header:    &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
			OrderFill: &trdcommon.OrderFill{FillID: &fillID, OrderID: &orderID, Price: &price, Qty: &qty},
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushTrdNotify(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	accID := uint64(123456789)
	trdMarket := int32(1)
	notifyType := int32(1)

	resp := &trdnotify.Response{
		RetType: &retType,
		S2C:     &trdnotify.S2C{Header: &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket}, Type: &notifyType},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handlePushNotify(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	notifyType := int32(1)

	resp := &notify.Response{
		RetType: &retType,
		S2C:     &notify.S2C{Type: &notifyType},
	}

	return s.successResponse(pkt, resp)
}

