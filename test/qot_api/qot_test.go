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

package qot_test

import (
	"context"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesttradedate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/test/fixtures"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

func TestGetBasicQot_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	// Register GetBasicQot handler
	server.RegisterHandler(3004, func(req []byte) ([]byte, error) {
		var reqMsg qotgetbasicqot.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		// Build realistic HSI response
		hsiQuote := fixtures.HSIQuote()

		s2c := &qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{hsiQuote},
		}

		return proto.Marshal(&qotgetbasicqot.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// Call API
	result, err := qot.GetBasicQot(context.Background(), cli, []*qotcommon.Security{fixtures.HSISecurity()})
	if err != nil {
		t.Fatalf("GetBasicQot failed: %v", err)
	}

	// Assertions
	if len(result) != 1 {
		t.Fatalf("Expected 1 quote, got %d", len(result))
	}

	quote := result[0]
	if quote.Security.GetCode() != fixtures.HSICode {
		t.Errorf("Expected code %s, got %s", fixtures.HSICode, quote.Security.GetCode())
	}

	if quote.CurPrice != 18523.45 {
		t.Errorf("Expected price 18523.45, got %f", quote.CurPrice)
	}

	server.AssertProtoID(t, 3004)
}

func TestGetKL_HSI_Day(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3006, func(req []byte) ([]byte, error) {
		var reqMsg qotgetkl.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		reqNum := int(reqMsg.C2S.GetReqNum())
		klList := fixtures.HSIKLineData(reqNum, reqMsg.C2S.GetKlType())

		s2c := &qotgetkl.S2C{
			Security: fixtures.HSISecurity(),
			Name:     proto.String(fixtures.HSIName),
			KlList:   klList,
		}

		return proto.Marshal(&qotgetkl.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// Request 10 days of HSI K-line
	req := &qot.GetKLRequest{
		Security:  fixtures.HSISecurity(),
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    10,
	}

	result, err := qot.GetKL(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetKL failed: %v", err)
	}

	if len(result.KLList) != 10 {
		t.Errorf("Expected 10 K-lines, got %d", len(result.KLList))
	}

	// Validate first K-line
	kl := result.KLList[0]
	if kl.OpenPrice < 18000 || kl.OpenPrice > 19000 {
		t.Errorf("HSI open price out of range: %f", kl.OpenPrice)
	}

	if kl.Volume == 0 {
		t.Error("K-line volume should not be zero")
	}

	server.AssertProtoID(t, 3006)
}

func TestGetKL_HSI_Min1(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3006, func(req []byte) ([]byte, error) {
		var reqMsg qotgetkl.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		reqNum := int(reqMsg.C2S.GetReqNum())
		klList := fixtures.HSIKLineData(reqNum, reqMsg.C2S.GetKlType())

		s2c := &qotgetkl.S2C{
			Security: fixtures.HSISecurity(),
			Name:     proto.String(fixtures.HSIName),
			KlList:   klList,
		}

		return proto.Marshal(&qotgetkl.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetKLRequest{
		Security:  fixtures.HSISecurity(),
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_1Min),
		ReqNum:    5,
	}

	result, err := qot.GetKL(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetKL 1min failed: %v", err)
	}

	if len(result.KLList) != 5 {
		t.Errorf("Expected 5 K-lines, got %d", len(result.KLList))
	}

	server.AssertProtoID(t, 3006)
}

func TestGetOrderBook_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3012, func(req []byte) ([]byte, error) {
		var reqMsg qotgetorderbook.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		num := int(reqMsg.C2S.GetNum())
		asks, bids := fixtures.HSIOrderBookLevels(num)

		s2c := &qotgetorderbook.S2C{
			Security:         fixtures.HSISecurity(),
			Name:             proto.String(fixtures.HSIName),
			OrderBookAskList: asks,
			OrderBookBidList: bids,
		}

		return proto.Marshal(&qotgetorderbook.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetOrderBookRequest{
		Security: fixtures.HSISecurity(),
		Num:      10,
	}

	result, err := qot.GetOrderBook(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetOrderBook failed: %v", err)
	}

	if len(result.OrderBookAskList) != 10 {
		t.Errorf("Expected 10 ask levels, got %d", len(result.OrderBookAskList))
	}

	if len(result.OrderBookBidList) != 10 {
		t.Errorf("Expected 10 bid levels, got %d", len(result.OrderBookBidList))
	}

	// Verify ask > bid (no crossed book)
	if len(result.OrderBookAskList) > 0 && len(result.OrderBookBidList) > 0 {
		bestAsk := result.OrderBookAskList[0].Price
		bestBid := result.OrderBookBidList[0].Price
		if bestAsk <= bestBid {
			t.Errorf("Crossed book: ask %f <= bid %f", bestAsk, bestBid)
		}
	}

	server.AssertProtoID(t, 3012)
}

func TestGetTicker_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3010, func(req []byte) ([]byte, error) {
		var reqMsg qotgetticker.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		maxRet := int(reqMsg.C2S.GetMaxRetNum())
		tickers := fixtures.HSITickerData(maxRet)

		s2c := &qotgetticker.S2C{
			Security:   fixtures.HSISecurity(),
			Name:       proto.String(fixtures.HSIName),
			TickerList: tickers,
		}

		return proto.Marshal(&qotgetticker.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetTickerRequest{
		Security: fixtures.HSISecurity(),
		Num:      20,
	}

	result, err := qot.GetTicker(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetTicker failed: %v", err)
	}

	if len(result.TickerList) != 20 {
		t.Errorf("Expected 20 tickers, got %d", len(result.TickerList))
	}

	// Validate ticker data
	for i, ticker := range result.TickerList {
		if ticker.Price <= 0 {
			t.Errorf("Ticker %d has invalid price: %f", i, ticker.Price)
		}
		if ticker.Volume == 0 {
			t.Errorf("Ticker %d has zero volume", i)
		}
	}

	server.AssertProtoID(t, 3010)
}

func TestGetRT_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3008, func(req []byte) ([]byte, error) {
		var reqMsg qotgetrt.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		// Return 240 minutes (full HK trading day)
		rtList := fixtures.HSIRTDData(240)

		s2c := &qotgetrt.S2C{
			Security: fixtures.HSISecurity(),
			Name:     proto.String(fixtures.HSIName),
			RtList:   rtList,
		}

		return proto.Marshal(&qotgetrt.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetRTRequest{
		Security: fixtures.HSISecurity(),
	}

	result, err := qot.GetRT(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetRT failed: %v", err)
	}

	if len(result.RTList) != 240 {
		t.Errorf("Expected 240 RT data points, got %d", len(result.RTList))
	}

	server.AssertProtoID(t, 3008)
}

func TestGetBroker_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3014, func(req []byte) ([]byte, error) {
		num := 10
		askBrokers := make([]*qotcommon.Broker, 0, num)
		bidBrokers := make([]*qotcommon.Broker, 0, num)

		for i := 0; i < num; i++ {
			brokerID := int64(1000 + i)
			volume := int64(10000 + i*1000)
			pos := int32(i)

			askBrokers = append(askBrokers, &qotcommon.Broker{
				Id:     &brokerID,
				Name:   proto.String("Test Broker"),
				Pos:    &pos,
				Volume: &volume,
			})

			bidBrokers = append(bidBrokers, &qotcommon.Broker{
				Id:     &brokerID,
				Name:   proto.String("Test Broker"),
				Pos:    &pos,
				Volume: &volume,
			})
		}

		s2c := &qotgetbroker.S2C{
			Security:      fixtures.HSISecurity(),
			Name:          proto.String(fixtures.HSIName),
			BrokerAskList: askBrokers,
			BrokerBidList: bidBrokers,
		}

		return proto.Marshal(&qotgetbroker.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetBrokerRequest{
		Security: fixtures.HSISecurity(),
		Num:      10,
	}

	result, err := qot.GetBroker(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetBroker failed: %v", err)
	}

	if len(result.AskBrokerList) != 10 {
		t.Errorf("Expected 10 ask brokers, got %d", len(result.AskBrokerList))
	}

	if len(result.BidBrokerList) != 10 {
		t.Errorf("Expected 10 bid brokers, got %d", len(result.BidBrokerList))
	}

	server.AssertProtoID(t, 3014)
}

func TestGetStaticInfo_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3202, func(req []byte) ([]byte, error) {
		sec := fixtures.HSISecurity()
		id := int64(800100)
		lotSize := int32(1)
		secType := int32(qotcommon.SecurityType_SecurityType_Index)
		listTime := "1969-07-31"
		name := fixtures.HSIName

		s2c := &qotgetstaticinfo.S2C{
			StaticInfoList: []*qotcommon.SecurityStaticInfo{
				{
					Basic: &qotcommon.SecurityStaticBasic{
						Security: sec,
						Id:       &id,
						LotSize:  &lotSize,
						SecType:  &secType,
						Name:     &name,
						ListTime: &listTime,
					},
				},
			},
		}

		return proto.Marshal(&qotgetstaticinfo.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetStaticInfoRequest{
		SecurityList: []*qotcommon.Security{fixtures.HSISecurity()},
	}

	result, err := qot.GetStaticInfo(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetStaticInfo failed: %v", err)
	}

	if len(result.StaticInfoList) != 1 {
		t.Errorf("Expected 1 static info, got %d", len(result.StaticInfoList))
	}

	info := result.StaticInfoList[0]
	if info.GetBasic().GetSecurity().GetCode() != fixtures.HSICode {
		t.Errorf("Expected code %s, got %s", fixtures.HSICode, info.GetBasic().GetSecurity().GetCode())
	}

	server.AssertProtoID(t, 3202)
}

func TestRequestTradeDate_HK(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3219, func(req []byte) ([]byte, error) {
		tradeDates := []*qotrequesttradedate.TradeDate{
			{Time: proto.String("2026-04-01")},
			{Time: proto.String("2026-04-02")},
			{Time: proto.String("2026-04-03")},
			{Time: proto.String("2026-04-06")},
			{Time: proto.String("2026-04-07")},
		}

		s2c := &qotrequesttradedate.S2C{
			TradeDateList: tradeDates,
		}

		return proto.Marshal(&qotrequesttradedate.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.RequestTradeDateRequest{
		Market: fixtures.HSIMarket,
	}

	result, err := qot.RequestTradeDate(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("RequestTradeDate failed: %v", err)
	}

	if len(result.TradeDateList) != 5 {
		t.Errorf("Expected 5 trade dates, got %d", len(result.TradeDateList))
	}

	server.AssertProtoID(t, 3219)
}

func TestSubscribe_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3001, func(req []byte) ([]byte, error) {
		s2c := &qotsub.S2C{}
		return proto.Marshal(&qotsub.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.SubscribeRequest{
		SecurityList: []*qotcommon.Security{fixtures.HSISecurity()},
		SubTypeList:  []qot.SubType{qot.SubType_Basic, qot.SubType_KL},
		IsSubOrUnSub: true,
	}

	_, err := qot.Subscribe(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	server.AssertProtoID(t, 3001)
}

func TestGetCapitalFlow_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3211, func(req []byte) ([]byte, error) {
		inFlow := 12345678.90
		mainInFlow := 12345678.90
		bigInFlow := 23456789.01
		midInFlow := 34567890.12
		smlInFlow := 45678901.23
		timeStr := "2026-04-10 10:00:00"

		s2c := &qotgetcapitalflow.S2C{
			FlowItemList: []*qotgetcapitalflow.CapitalFlowItem{
				{
					InFlow:     &inFlow,
					Time:       &timeStr,
					MainInFlow: &mainInFlow,
					BigInFlow:  &bigInFlow,
					MidInFlow:  &midInFlow,
					SmlInFlow:  &smlInFlow,
				},
			},
		}

		return proto.Marshal(&qotgetcapitalflow.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &qot.GetCapitalFlowRequest{
		Security: fixtures.HSISecurity(),
	}

	result, err := qot.GetCapitalFlow(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetCapitalFlow failed: %v", err)
	}

	if len(result.FlowItemList) != 1 {
		t.Errorf("Expected 1 capital flow item, got %d", len(result.FlowItemList))
	}

	server.AssertProtoID(t, 3211)
}

func TestGetCapitalDistribution_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(3212, func(req []byte) ([]byte, error) {
		s2c := &qotgetcapitaldistribution.S2C{
			CapitalInBig:    proto.Float64(100000000.0),
			CapitalInMid:    proto.Float64(50000000.0),
			CapitalInSmall:  proto.Float64(25000000.0),
			CapitalOutBig:   proto.Float64(90000000.0),
			CapitalOutMid:   proto.Float64(55000000.0),
			CapitalOutSmall: proto.Float64(20000000.0),
		}

		return proto.Marshal(&qotgetcapitaldistribution.Response{S2C: s2c})
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	result, err := qot.GetCapitalDistribution(context.Background(), cli, fixtures.HSISecurity())
	if err != nil {
		t.Fatalf("GetCapitalDistribution failed: %v", err)
	}

	if result.CapitalDistribution.CapitalInBig <= 0 {
		t.Error("CapitalInBig should be positive")
	}

	server.AssertProtoID(t, 3212)
}
