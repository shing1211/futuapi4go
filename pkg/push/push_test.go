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

package push

import (
	"testing"

	"google.golang.org/protobuf/proto"

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

func TestParseUpdateBasicQotInvalidData(t *testing.T) {
	result, err := ParseUpdateBasicQot([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateBasicQot should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateBasicQot should return nil for empty data")
	}

	_, err = ParseUpdateBasicQot([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("ParseUpdateBasicQot should fail with invalid protobuf data")
	}
}

func TestParseUpdateKLInvalidData(t *testing.T) {
	result, err := ParseUpdateKL([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateKL should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateKL should return nil for empty data")
	}
}

func TestParseUpdateOrderBookInvalidData(t *testing.T) {
	result, err := ParseUpdateOrderBook([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateOrderBook should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateOrderBook should return nil for empty data")
	}
}

func TestParseUpdateTickerInvalidData(t *testing.T) {
	result, err := ParseUpdateTicker([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateTicker should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateTicker should return nil for empty data")
	}
}

func TestParseUpdateRTInvalidData(t *testing.T) {
	result, err := ParseUpdateRT([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateRT should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateRT should return nil for empty data")
	}
}

func TestParseUpdateBrokerInvalidData(t *testing.T) {
	result, err := ParseUpdateBroker([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateBroker should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateBroker should return nil for empty data")
	}
}

func TestParseUpdatePriceReminderInvalidData(t *testing.T) {
	_, err := ParseUpdatePriceReminder([]byte{})
	if err == nil {
		t.Error("ParseUpdatePriceReminder should fail with empty data")
	}
}

func TestParseSystemNotifyInvalidData(t *testing.T) {
	_, err := ParseSystemNotify([]byte{})
	if err == nil {
		t.Error("ParseSystemNotify should fail with empty data")
	}
}

func TestParseUpdateOrderInvalidData(t *testing.T) {
	_, err := ParseUpdateOrder([]byte{})
	if err == nil {
		t.Error("ParseUpdateOrder should fail with empty data")
	}
}

func TestParseUpdateOrderFillInvalidData(t *testing.T) {
	_, err := ParseUpdateOrderFill([]byte{})
	if err == nil {
		t.Error("ParseUpdateOrderFill should fail with empty data")
	}
}

func TestParseTrdNotifyInvalidData(t *testing.T) {
	_, err := ParseTrdNotify([]byte{})
	if err == nil {
		t.Error("ParseTrdNotify should fail with empty data")
	}
}

func TestParseUpdateBasicQotValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"
	curPrice := 350.50
	openPrice := 348.00
	highPrice := 352.00
	lowPrice := 347.00
	volume := int64(12345678)
	turnover := 4321098765.00
	isSuspended := false
	listTime := "2026-01-01"
	priceSpread := 0.01
	updateTime := "15:00:00"
	lastClosePrice := 349.00
	turnoverRate := 0.5
	amplitude := 1.43

	s2c := &qotupdatebasicqot.S2C{
		BasicQotList: []*qotcommon.BasicQot{
			{
				Security:       &qotcommon.Security{Market: &hkMarket, Code: &code},
				Name:           &name,
				IsSuspended:    &isSuspended,
				ListTime:       &listTime,
				PriceSpread:    &priceSpread,
				UpdateTime:     &updateTime,
				CurPrice:       &curPrice,
				OpenPrice:      &openPrice,
				HighPrice:      &highPrice,
				LowPrice:       &lowPrice,
				LastClosePrice: &lastClosePrice,
				Volume:         &volume,
				Turnover:       &turnover,
				TurnoverRate:   &turnoverRate,
				Amplitude:      &amplitude,
			},
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateBasicQot(body)
	if err != nil {
		t.Fatalf("ParseUpdateBasicQot failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", result.Security.GetCode())
	}
	if result.Name != "Tencent" {
		t.Errorf("expected name Tencent, got %s", result.Name)
	}
	if result.CurPrice != 350.50 {
		t.Errorf("expected CurPrice 350.50, got %f", result.CurPrice)
	}
	if result.OpenPrice != 348.00 {
		t.Errorf("expected OpenPrice 348.00, got %f", result.OpenPrice)
	}
	if result.HighPrice != 352.00 {
		t.Errorf("expected HighPrice 352.00, got %f", result.HighPrice)
	}
	if result.LowPrice != 347.00 {
		t.Errorf("expected LowPrice 347.00, got %f", result.LowPrice)
	}
	if result.Volume != 12345678 {
		t.Errorf("expected Volume 12345678, got %d", result.Volume)
	}
	if result.Turnover != 4321098765.00 {
		t.Errorf("expected Turnover 4321098765.00, got %f", result.Turnover)
	}
}

func TestParseUpdateBasicQotZeroLengthList(t *testing.T) {
	s2c := &qotupdatebasicqot.S2C{
		BasicQotList: []*qotcommon.BasicQot{},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateBasicQot(body)
	if err != nil {
		t.Fatalf("ParseUpdateBasicQot should not error on zero-length list, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateBasicQot should return nil for zero-length list")
	}
}

func TestParseUpdateKLValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"
	rehabType := int32(0)
	klType := int32(4)
	timeStr := "2026-04-08 15:00:00"
	closePrice := 350.50
	openPrice := 348.00
	highPrice := 352.00
	lowPrice := 347.00
	lastClosePrice := 349.00
	volume := int64(12345678)
	turnover := 4321098765.00
	changeRate := 0.43
	timestamp := 1775635200.0
	isBlank := false

	s2c := &qotupdatekl.S2C{
		RehabType: &rehabType,
		KlType:    &klType,
		Security:  &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:      &name,
		KlList: []*qotcommon.KLine{
			{
				Time:           &timeStr,
				IsBlank:        &isBlank,
				ClosePrice:     &closePrice,
				OpenPrice:      &openPrice,
				HighPrice:      &highPrice,
				LowPrice:       &lowPrice,
				LastClosePrice: &lastClosePrice,
				Volume:         &volume,
				Turnover:       &turnover,
				ChangeRate:     &changeRate,
				Timestamp:      &timestamp,
			},
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateKL(body)
	if err != nil {
		t.Fatalf("ParseUpdateKL failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.RehabType != 0 {
		t.Errorf("expected RehabType 0, got %d", result.RehabType)
	}
	if result.KlType != 4 {
		t.Errorf("expected KlType 4, got %d", result.KlType)
	}
	if result.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", result.Security.GetCode())
	}
	if len(result.KLList) != 1 {
		t.Fatalf("expected 1 KL entry, got %d", len(result.KLList))
	}
	if result.KLList[0].GetClosePrice() != 350.50 {
		t.Errorf("expected close price 350.50, got %f", result.KLList[0].GetClosePrice())
	}
}

func TestParseUpdateOrderBookValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"
	svrRecvTimeBid := "10:00:00"
	svrRecvTimeBidTs := 1775635200.0
	svrRecvTimeAsk := "10:00:00"
	svrRecvTimeAskTs := 1775635200.0

	askPrice := 351.00
	askVolume := int64(5000)
	askOrederCount := int32(3)
	bidPrice := 350.00
	bidVolume := int64(5000)
	bidOrederCount := int32(3)

	s2c := &qotupdateorderbook.S2C{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:     &name,
		OrderBookAskList: []*qotcommon.OrderBook{
			{Price: &askPrice, Volume: &askVolume, OrederCount: &askOrederCount},
		},
		OrderBookBidList: []*qotcommon.OrderBook{
			{Price: &bidPrice, Volume: &bidVolume, OrederCount: &bidOrederCount},
		},
		SvrRecvTimeBid:          &svrRecvTimeBid,
		SvrRecvTimeBidTimestamp: &svrRecvTimeBidTs,
		SvrRecvTimeAsk:          &svrRecvTimeAsk,
		SvrRecvTimeAskTimestamp: &svrRecvTimeAskTs,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateOrderBook(body)
	if err != nil {
		t.Fatalf("ParseUpdateOrderBook failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", result.Security.GetCode())
	}
	if len(result.OrderBookAskList) != 1 {
		t.Errorf("expected 1 ask level, got %d", len(result.OrderBookAskList))
	}
	if len(result.OrderBookBidList) != 1 {
		t.Errorf("expected 1 bid level, got %d", len(result.OrderBookBidList))
	}
	if result.OrderBookAskList[0].GetPrice() != 351.00 {
		t.Errorf("expected ask price 351.00, got %f", result.OrderBookAskList[0].GetPrice())
	}
	if result.SvrRecvTimeBid != "10:00:00" {
		t.Errorf("expected SvrRecvTimeBid 10:00:00, got %s", result.SvrRecvTimeBid)
	}
}

func TestParseUpdateTickerValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"

	timeStr := "10:00:00"
	sequence := int64(123456)
	dir := int32(1)
	price := 350.50
	volume := int64(1000)
	turnover := 350500.00
	recvTime := 1775635200.0
	typ := int32(0)
	typeSign := int32(1)
	timestamp := 1775635200.0

	s2c := &qotupdateticker.S2C{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:     &name,
		TickerList: []*qotcommon.Ticker{
			{
				Time:      &timeStr,
				Sequence:  &sequence,
				Dir:       &dir,
				Price:     &price,
				Volume:    &volume,
				Turnover:  &turnover,
				RecvTime:  &recvTime,
				Type:      &typ,
				TypeSign:  &typeSign,
				Timestamp: &timestamp,
			},
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateTicker(body)
	if err != nil {
		t.Fatalf("ParseUpdateTicker failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if len(result.TickerList) != 1 {
		t.Errorf("expected 1 ticker, got %d", len(result.TickerList))
	}
	if result.TickerList[0].GetPrice() != 350.50 {
		t.Errorf("expected price 350.50, got %f", result.TickerList[0].GetPrice())
	}
	if result.TickerList[0].GetVolume() != 1000 {
		t.Errorf("expected volume 1000, got %d", result.TickerList[0].GetVolume())
	}
}

func TestParseUpdateRTValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"

	timeStr := "10:00:00"
	minute := int32(600)
	isBlank := false
	price := 350.50
	lastClosePrice := 349.00
	avgPrice := 349.80
	volume := int64(12345678)
	turnover := 4321098765.00

	s2c := &qotupdatert.S2C{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:     &name,
		RtList: []*qotcommon.TimeShare{
			{
				Time:           &timeStr,
				Minute:         &minute,
				IsBlank:        &isBlank,
				Price:          &price,
				LastClosePrice: &lastClosePrice,
				AvgPrice:       &avgPrice,
				Volume:         &volume,
				Turnover:       &turnover,
			},
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateRT(body)
	if err != nil {
		t.Fatalf("ParseUpdateRT failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if len(result.RTList) != 1 {
		t.Errorf("expected 1 RT entry, got %d", len(result.RTList))
	}
	if result.RTList[0].GetPrice() != 350.50 {
		t.Errorf("expected price 350.50, got %f", result.RTList[0].GetPrice())
	}
	if result.RTList[0].GetLastClosePrice() != 349.00 {
		t.Errorf("expected last close price 349.00, got %f", result.RTList[0].GetLastClosePrice())
	}
}

func TestParseUpdateBrokerValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"

	askID := int64(1)
	askName := "Citi"
	askPos := int32(1)
	askVolume := int64(5000)
	bidID := int64(2)
	bidName := "HSBC"
	bidPos := int32(1)
	bidVolume := int64(6000)

	s2c := &qotupdatebroker.S2C{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:     &name,
		BrokerAskList: []*qotcommon.Broker{
			{Id: &askID, Name: &askName, Pos: &askPos, Volume: &askVolume},
		},
		BrokerBidList: []*qotcommon.Broker{
			{Id: &bidID, Name: &bidName, Pos: &bidPos, Volume: &bidVolume},
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateBroker(body)
	if err != nil {
		t.Fatalf("ParseUpdateBroker failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if len(result.AskBrokerList) != 1 {
		t.Errorf("expected 1 ask broker, got %d", len(result.AskBrokerList))
	}
	if len(result.BidBrokerList) != 1 {
		t.Errorf("expected 1 bid broker, got %d", len(result.BidBrokerList))
	}
	if result.AskBrokerList[0].GetName() != "Citi" {
		t.Errorf("expected ask broker name Citi, got %s", result.AskBrokerList[0].GetName())
	}
	if result.BidBrokerList[0].GetName() != "HSBC" {
		t.Errorf("expected bid broker name HSBC, got %s", result.BidBrokerList[0].GetName())
	}
}

func TestParseUpdatePriceReminderValidData(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"
	price := 350.50
	changeRate := 0.43
	marketStatus := int32(1)
	content := "Price reached target"
	note := "Alert note"
	key := int64(123)
	typ := int32(1)
	setValue := 350.0
	curValue := 350.50

	s2c := &qotupdatepricereminder.S2C{
		Security:     &qotcommon.Security{Market: &hkMarket, Code: &code},
		Name:         &name,
		Price:        &price,
		ChangeRate:   &changeRate,
		MarketStatus: &marketStatus,
		Content:      &content,
		Note:         &note,
		Key:          &key,
		Type:         &typ,
		SetValue:     &setValue,
		CurValue:     &curValue,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdatePriceReminder(body)
	if err != nil {
		t.Fatalf("ParseUpdatePriceReminder failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Price != 350.50 {
		t.Errorf("expected price 350.50, got %f", result.Price)
	}
	if result.ChangeRate != 0.43 {
		t.Errorf("expected change rate 0.43, got %f", result.ChangeRate)
	}
	if result.Content != "Price reached target" {
		t.Errorf("expected content 'Price reached target', got %s", result.Content)
	}
	if result.Key != 123 {
		t.Errorf("expected key 123, got %d", result.Key)
	}
}

func TestParseSystemNotifyValidData(t *testing.T) {
	typ := int32(1)
	eventType := int32(1)
	desc := "test event"

	s2c := &notify.S2C{
		Type: &typ,
		Event: &notify.GtwEvent{
			EventType: &eventType,
			Desc:      &desc,
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseSystemNotify(body)
	if err != nil {
		t.Fatalf("ParseSystemNotify failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Type != 1 {
		t.Errorf("expected type 1, got %d", result.Type)
	}
	if result.Event == nil {
		t.Error("Event should not be nil")
	}
	if result.Event.GetEventType() != 1 {
		t.Errorf("expected event type 1, got %d", result.Event.GetEventType())
	}
}

func TestParseUpdateOrderValidData(t *testing.T) {
	accID := uint64(123456789)
	trdEnv := int32(trdcommon.TrdEnv_TrdEnv_Real)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	orderID := uint64(9876543210)
	code := "00700"
	name := "Tencent"
	trdSide := int32(1)
	orderType := int32(1)
	orderStatus := int32(2)
	orderPrice := 350.00
	orderQty := 100.0
	fillQty := 50.0
	createTime := "2026-04-08 10:00:00"
	updateTime := "2026-04-08 10:05:00"
	fillAvgPrice := 350.00

	s2c := &trdupdateorder.S2C{
		Header: &trdcommon.TrdHeader{AccID: &accID, TrdEnv: &trdEnv, TrdMarket: &trdMarket},
		Order: &trdcommon.Order{
			OrderID:      &orderID,
			OrderIDEx:    &code,
			Code:         &code,
			Name:         &name,
			TrdSide:      &trdSide,
			OrderType:    &orderType,
			OrderStatus:  &orderStatus,
			Price:        &orderPrice,
			Qty:          &orderQty,
			FillQty:      &fillQty,
			CreateTime:   &createTime,
			UpdateTime:   &updateTime,
			FillAvgPrice: &fillAvgPrice,
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateOrder(body)
	if err != nil {
		t.Fatalf("ParseUpdateOrder failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Order == nil {
		t.Fatal("Order should not be nil")
	}
	if result.Order.GetOrderID() != 9876543210 {
		t.Errorf("expected order ID 9876543210, got %d", result.Order.GetOrderID())
	}
	if result.Order.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", result.Order.GetCode())
	}
	if result.Order.GetFillQty() != 50.0 {
		t.Errorf("expected fill qty 50.0, got %f", result.Order.GetFillQty())
	}
}

func TestParseUpdateOrderFillValidData(t *testing.T) {
	accID := uint64(123456789)
	trdEnv := int32(trdcommon.TrdEnv_TrdEnv_Real)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	orderID := uint64(9876543210)
	fillID := uint64(1111111111)
	code := "00700"
	name := "Tencent"
	trdSide := int32(1)
	price := 350.00
	qty := 100.0
	createTime := "2026-04-08 10:05:00"

	s2c := &trdupdateorderfill.S2C{
		Header: &trdcommon.TrdHeader{AccID: &accID, TrdEnv: &trdEnv, TrdMarket: &trdMarket},
		OrderFill: &trdcommon.OrderFill{
			OrderID:    &orderID,
			FillID:     &fillID,
			FillIDEx:   &code,
			Code:       &code,
			Name:       &name,
			TrdSide:    &trdSide,
			Price:      &price,
			Qty:        &qty,
			CreateTime: &createTime,
		},
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseUpdateOrderFill(body)
	if err != nil {
		t.Fatalf("ParseUpdateOrderFill failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.OrderFill == nil {
		t.Fatal("OrderFill should not be nil")
	}
	if result.OrderFill.GetFillID() != 1111111111 {
		t.Errorf("expected fill ID 1111111111, got %d", result.OrderFill.GetFillID())
	}
	if result.OrderFill.GetPrice() != 350.00 {
		t.Errorf("expected price 350.00, got %f", result.OrderFill.GetPrice())
	}
}

func TestParseTrdNotifyValidData(t *testing.T) {
	accID := uint64(123456789)
	trdEnv := int32(trdcommon.TrdEnv_TrdEnv_Real)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	typ := int32(1)

	s2c := &trdnotify.S2C{
		Header: &trdcommon.TrdHeader{AccID: &accID, TrdEnv: &trdEnv, TrdMarket: &trdMarket},
		Type:   &typ,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		t.Fatalf("failed to marshal protobuf: %v", err)
	}

	result, err := ParseTrdNotify(body)
	if err != nil {
		t.Fatalf("ParseTrdNotify failed: %v", err)
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Type != 1 {
		t.Errorf("expected type 1, got %d", result.Type)
	}
	if result.Header == nil {
		t.Error("Header should not be nil")
	}
}
