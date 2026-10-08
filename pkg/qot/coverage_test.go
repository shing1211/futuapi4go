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
	"encoding/json"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategy"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyanalysis"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyspreads"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsecuritysnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryklquota"
	"github.com/shing1211/futuapi4go/pkg/pb/qotregqotpush"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
	futuapitestutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

func hkSecurity(code string) *qotcommon.Security {
	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	return &qotcommon.Security{Market: &market, Code: &code}
}

func newSuccessResponse(s2c any) proto.Message {
	switch v := s2c.(type) {
	case *qotgetbasicqot.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetbasicqot.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetkl.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetkl.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetorderbook.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetorderbook.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetticker.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetticker.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetrt.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetrt.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetbroker.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetbroker.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetcapitalflow.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetcapitalflow.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetcapitaldistribution.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetcapitaldistribution.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotsub.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotsub.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotregqotpush.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotregqotpush.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetsubinfo.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetsubinfo.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetsecuritysnapshot.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetsecuritysnapshot.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetstaticinfo.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetstaticinfo.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotrequesthistorykl.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotrequesthistorykl.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotrequesthistoryklquota.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotrequesthistoryklquota.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetoptionquote.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetoptionquote.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetoptionstrategy.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetoptionstrategy.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetoptionstrategyanalysis.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetoptionstrategyanalysis.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	case *qotgetoptionstrategyspreads.S2C:
		retType := int32(common.RetType_RetType_Succeed)
		return &qotgetoptionstrategyspreads.Response{RetType: &retType, RetMsg: proto.String("ok"), S2C: v}
	default:
		return nil
	}
}


// =============================================================================
// GetBasicQot
// =============================================================================

func TestGetBasicQot_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	code00700 := "00700"
	nameTencent := "Tencent"
	curPrice := 350.50
	highPrice := 352.00
	openPrice := 348.00
	lowPrice := 347.00
	lastClosePrice := 349.00
	volume := int64(12345678)
	turnover := 4321098765.00
	turnoverRate := 0.025
	amplitude := 0.030
	updateTime := "09:30:00"
	isSuspended := false
	priceSpread := 0.01
	listTime := "2004-06-16"
	listTimestamp := 1087267200.0
	updateTimestamp := 1744162200.0
	secStatus := int32(0)

	server.RegisterHandler(ProtoID_GetBasicQot, func(reqBody []byte) (proto.Message, error) {
		var req qotgetbasicqot.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetBasicQot request failed: %v", err)
		}
		if req.C2S == nil || len(req.C2S.SecurityList) == 0 {
			t.Fatal("expected non-empty security list")
		}
		return newSuccessResponse(&qotgetbasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{
				{
					Security:        &qotcommon.Security{Market: func() *int32 { v := int32(1); return &v }(), Code: &code00700},
					Name:            &nameTencent,
					CurPrice:        &curPrice,
					HighPrice:       &highPrice,
					OpenPrice:       &openPrice,
					LowPrice:        &lowPrice,
					LastClosePrice:  &lastClosePrice,
					Volume:          &volume,
					Turnover:        &turnover,
					TurnoverRate:    &turnoverRate,
					Amplitude:       &amplitude,
					UpdateTime:      &updateTime,
					IsSuspended:     &isSuspended,
					PriceSpread:     &priceSpread,
					ListTime:        &listTime,
					ListTimestamp:   &listTimestamp,
					UpdateTimestamp: &updateTimestamp,
					SecStatus:       &secStatus,
				},
			},
		}), nil
	})

	ctx := context.Background()
	result, err := GetBasicQot(ctx, cli, []*qotcommon.Security{hkSecurity("00700")})
	if err != nil {
		t.Fatalf("GetBasicQot failed: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
	bq := result[0]
	if bq.Name != "Tencent" {
		t.Errorf("expected Name Tencent, got %s", bq.Name)
	}
	if bq.CurPrice != 350.50 {
		t.Errorf("expected CurPrice 350.50, got %f", bq.CurPrice)
	}
	if bq.Volume != 12345678 {
		t.Errorf("expected Volume 12345678, got %d", bq.Volume)
	}
	if bq.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", bq.Security.GetCode())
	}

	server.AssertProtoID(t, ProtoID_GetBasicQot)
}

func TestGetBasicQot_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Failed)
	retMsg := "failed"
	server.RegisterHandler(ProtoID_GetBasicQot, func(reqBody []byte) (proto.Message, error) {
		return &qotgetbasicqot.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetBasicQot(ctx, cli, []*qotcommon.Security{hkSecurity("00700")})
	if err == nil {
		t.Fatal("expected error for failed RetType")
	}
}

func TestGetBasicQot_EmptySecurityList(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetBasicQot(ctx, cli, []*qotcommon.Security{})
	if err == nil {
		t.Error("expected error for empty security list")
	}
}


// =============================================================================
// GetKL
// =============================================================================

func TestGetKL_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	klTime := "2026-04-08 15:00:00"
	klClosePrice := 350.50
	klOpenPrice := 348.00
	klHighPrice := 352.00
	klLowPrice := 347.00
	klLastClosePrice := 349.00
	klVolume := int64(12345678)
	klTurnover := 4321098765.00
	klChangeRate := 0.43
	klTimestamp := 1775635200.0
	klIsBlank := false

	server.RegisterHandler(ProtoID_GetKL, func(reqBody []byte) (proto.Message, error) {
		var req qotgetkl.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetKL request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetkl.S2C{
			Security: hkSecurity("00700"),
			Name:     proto.String("Tencent"),
			KlList: []*qotcommon.KLine{
				{
					Time:           &klTime,
					ClosePrice:     &klClosePrice,
					OpenPrice:      &klOpenPrice,
					HighPrice:      &klHighPrice,
					LowPrice:       &klLowPrice,
					LastClosePrice: &klLastClosePrice,
					Volume:         &klVolume,
					Turnover:       &klTurnover,
					ChangeRate:     &klChangeRate,
					Timestamp:      &klTimestamp,
					IsBlank:        &klIsBlank,
				},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetKLRequest{
		Security:  hkSecurity("00700"),
		RehabType: 0,
		KLType:    6,
		ReqNum:    10,
	}
	result, err := GetKL(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetKL failed: %v", err)
	}

	if len(result.KLList) != 1 {
		t.Fatalf("expected 1 KLine, got %d", len(result.KLList))
	}
	kl := result.KLList[0]
	if kl.ClosePrice != 350.50 {
		t.Errorf("expected ClosePrice 350.50, got %f", kl.ClosePrice)
	}
	if kl.Volume != 12345678 {
		t.Errorf("expected Volume 12345678, got %d", kl.Volume)
	}
	if result.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", result.Security.GetCode())
	}

	server.AssertProtoID(t, ProtoID_GetKL)
}

func TestGetKL_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetKL(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetKL_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetKL(ctx, cli, &GetKLRequest{Security: nil, ReqNum: 10})
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestGetKL_NonPositiveReqNum(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetKL(ctx, cli, &GetKLRequest{Security: hkSecurity("00700"), ReqNum: 0})
	if err == nil {
		t.Error("expected error for ReqNum <= 0")
	}
}

func TestGetKL_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetKL, func(reqBody []byte) (proto.Message, error) {
		return &qotgetkl.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetKL(ctx, cli, &GetKLRequest{Security: hkSecurity("00700"), ReqNum: 10})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetOrderBook
// =============================================================================

func TestGetOrderBook_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	askPrice := 351.00
	askVolume := int64(5000)
	askOrderCount := int32(3)
	bidPrice := 350.00
	bidVolume := int64(5000)
	bidOrderCount := int32(3)
	svrRecvTimeBid := "10:00:00"
	svrRecvTimeBidTs := 1775635200.0
	svrRecvTimeAsk := "10:00:00"
	svrRecvTimeAskTs := 1775635200.0

	server.RegisterHandler(ProtoID_GetOrderBook, func(reqBody []byte) (proto.Message, error) {
		var req qotgetorderbook.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetOrderBook request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetorderbook.S2C{
			Security:             hkSecurity("00700"),
			Name:                 proto.String("Tencent"),
			OrderBookAskList: []*qotcommon.OrderBook{
				{Price: &askPrice, Volume: &askVolume, OrederCount: &askOrderCount},
			},
			OrderBookBidList: []*qotcommon.OrderBook{
				{Price: &bidPrice, Volume: &bidVolume, OrederCount: &bidOrderCount},
			},
			SvrRecvTimeBid:          &svrRecvTimeBid,
			SvrRecvTimeBidTimestamp: &svrRecvTimeBidTs,
			SvrRecvTimeAsk:          &svrRecvTimeAsk,
			SvrRecvTimeAskTimestamp: &svrRecvTimeAskTs,
		}), nil
	})

	ctx := context.Background()
	req := &GetOrderBookRequest{
		Security: hkSecurity("00700"),
		Num:      10,
	}
	result, err := GetOrderBook(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetOrderBook failed: %v", err)
	}

	if len(result.OrderBookAskList) != 1 {
		t.Fatalf("expected 1 ask level, got %d", len(result.OrderBookAskList))
	}
	if len(result.OrderBookBidList) != 1 {
		t.Fatalf("expected 1 bid level, got %d", len(result.OrderBookBidList))
	}
	if result.OrderBookAskList[0].Price != 351.00 {
		t.Errorf("expected ask price 351.00, got %f", result.OrderBookAskList[0].Price)
	}
	if result.OrderBookBidList[0].Price != 350.00 {
		t.Errorf("expected bid price 350.00, got %f", result.OrderBookBidList[0].Price)
	}

	server.AssertProtoID(t, ProtoID_GetOrderBook)
}

func TestGetOrderBook_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOrderBook(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetOrderBook_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOrderBook(ctx, cli, &GetOrderBookRequest{Security: nil})
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestGetOrderBook_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetOrderBook, func(reqBody []byte) (proto.Message, error) {
		return &qotgetorderbook.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetOrderBook(ctx, cli, &GetOrderBookRequest{Security: hkSecurity("00700"), Num: 10})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetTicker
// =============================================================================

func TestGetTicker_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	tickerTime := "10:00:00"
	tickerPrice := 350.0
	tickerVolume := int64(500)
	tickerDir := int32(1)
	tickerTurnover := 175000.0
	tickerSequence := int64(100)
	tickerTimestamp := 1775635200.0
	tickerType := int32(0)
	tickerTypeSign := int32(1)
	tickerRecvTime := 1775635200.0
	tickerPushDataType := int32(1)

	server.RegisterHandler(ProtoID_GetTicker, func(reqBody []byte) (proto.Message, error) {
		var req qotgetticker.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetTicker request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetticker.S2C{
			Security: hkSecurity("00700"),
			Name:     proto.String("Tencent"),
			TickerList: []*qotcommon.Ticker{
				{
					Time:         &tickerTime,
					Sequence:     &tickerSequence,
					Dir:          &tickerDir,
					Price:        &tickerPrice,
					Volume:       &tickerVolume,
					Turnover:     &tickerTurnover,
					RecvTime:     &tickerRecvTime,
					Type:         &tickerType,
					TypeSign:     &tickerTypeSign,
					Timestamp:    &tickerTimestamp,
					PushDataType: &tickerPushDataType,
				},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetTickerRequest{
		Security: hkSecurity("00700"),
		Num:      100,
	}
	result, err := GetTicker(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetTicker failed: %v", err)
	}

	if len(result.TickerList) != 1 {
		t.Fatalf("expected 1 ticker, got %d", len(result.TickerList))
	}
	tk := result.TickerList[0]
	if tk.Price != 350.0 {
		t.Errorf("expected Price 350.0, got %f", tk.Price)
	}
	if tk.Volume != 500 {
		t.Errorf("expected Volume 500, got %d", tk.Volume)
	}
	if tk.Dir != 1 {
		t.Errorf("expected Dir 1, got %d", tk.Dir)
	}

	server.AssertProtoID(t, ProtoID_GetTicker)
}

func TestGetTicker_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetTicker(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetTicker_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetTicker, func(reqBody []byte) (proto.Message, error) {
		return &qotgetticker.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetTicker(ctx, cli, &GetTickerRequest{Security: hkSecurity("00700"), Num: 100})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetRT
// =============================================================================

func TestGetRT_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	rtTime := "10:00:00"
	rtMinute := int32(600)
	rtIsBlank := false
	rtPrice := 350.0
	rtLastClosePrice := 349.0
	rtAvgPrice := 349.8
	rtVolume := int64(12345678)
	rtTurnover := 4321098765.00
	rtTimestamp := 1744116000.0

	server.RegisterHandler(ProtoID_GetRT, func(reqBody []byte) (proto.Message, error) {
		var req qotgetrt.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetRT request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetrt.S2C{
			Security: hkSecurity("00700"),
			Name:     proto.String("Tencent"),
			RtList: []*qotcommon.TimeShare{
				{
					Time:           &rtTime,
					Minute:         &rtMinute,
					IsBlank:        &rtIsBlank,
					Price:          &rtPrice,
					LastClosePrice: &rtLastClosePrice,
					AvgPrice:       &rtAvgPrice,
					Volume:         &rtVolume,
					Turnover:       &rtTurnover,
					Timestamp:      &rtTimestamp,
				},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetRTRequest{
		Security: hkSecurity("00700"),
	}
	result, err := GetRT(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetRT failed: %v", err)
	}

	if len(result.RTList) != 1 {
		t.Fatalf("expected 1 RT, got %d", len(result.RTList))
	}
	rt := result.RTList[0]
	if rt.Price != 350.0 {
		t.Errorf("expected Price 350.0, got %f", rt.Price)
	}
	if rt.Minute != 600 {
		t.Errorf("expected Minute 600, got %d", rt.Minute)
	}
	if rt.AvgPrice != 349.8 {
		t.Errorf("expected AvgPrice 349.8, got %f", rt.AvgPrice)
	}

	server.AssertProtoID(t, ProtoID_GetRT)
}

func TestGetRT_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetRT(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetRT_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetRT(ctx, cli, &GetRTRequest{Security: nil})
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestGetRT_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetRT, func(reqBody []byte) (proto.Message, error) {
		return &qotgetrt.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetRT(ctx, cli, &GetRTRequest{Security: hkSecurity("00700")})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetBroker
// =============================================================================

func TestGetBroker_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	askBrokerID := int64(1)
	askBrokerName := "Citi"
	askBrokerPos := int32(1)
	askBrokerVolume := int64(5000)
	askBrokerOrderID := int64(100)
	bidBrokerID := int64(2)
	bidBrokerName := "HSBC"
	bidBrokerPos := int32(1)
	bidBrokerVolume := int64(6000)
	bidBrokerOrderID := int64(200)

	server.RegisterHandler(ProtoID_GetBroker, func(reqBody []byte) (proto.Message, error) {
		var req qotgetbroker.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetBroker request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetbroker.S2C{
			Security: hkSecurity("00700"),
			Name:     proto.String("Tencent"),
			BrokerAskList: []*qotcommon.Broker{
				{Id: &askBrokerID, Name: &askBrokerName, Pos: &askBrokerPos, Volume: &askBrokerVolume, OrderID: &askBrokerOrderID},
			},
			BrokerBidList: []*qotcommon.Broker{
				{Id: &bidBrokerID, Name: &bidBrokerName, Pos: &bidBrokerPos, Volume: &bidBrokerVolume, OrderID: &bidBrokerOrderID},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetBrokerRequest{
		Security: hkSecurity("00700"),
	}
	result, err := GetBroker(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetBroker failed: %v", err)
	}

	if len(result.AskBrokerList) != 1 {
		t.Fatalf("expected 1 ask broker, got %d", len(result.AskBrokerList))
	}
	if len(result.BidBrokerList) != 1 {
		t.Fatalf("expected 1 bid broker, got %d", len(result.BidBrokerList))
	}
	if result.AskBrokerList[0].Name != "Citi" {
		t.Errorf("expected ask broker Citi, got %s", result.AskBrokerList[0].Name)
	}
	if result.BidBrokerList[0].Name != "HSBC" {
		t.Errorf("expected bid broker HSBC, got %s", result.BidBrokerList[0].Name)
	}

	server.AssertProtoID(t, ProtoID_GetBroker)
}

func TestGetBroker_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetBroker(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetBroker_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetBroker, func(reqBody []byte) (proto.Message, error) {
		return &qotgetbroker.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetBroker(ctx, cli, &GetBrokerRequest{Security: hkSecurity("00700")})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetCapitalFlow
// =============================================================================

func TestGetCapitalFlow_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	inFlow := 1000000.0
	flowTime := "2026-04-08"
	flowTimestamp := 1775635200.0
	mainInFlow := 800000.0
	superInFlow := 200000.0
	bigInFlow := 300000.0
	midInFlow := 250000.0
	smlInFlow := 250000.0
	lastValidTime := "2026-04-08"
	lastValidTs := 1775635200.0

	server.RegisterHandler(ProtoID_GetCapitalFlow, func(reqBody []byte) (proto.Message, error) {
		var req qotgetcapitalflow.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetCapitalFlow request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetcapitalflow.S2C{
			FlowItemList: []*qotgetcapitalflow.CapitalFlowItem{
				{
					InFlow:      &inFlow,
					Time:        &flowTime,
					Timestamp:   &flowTimestamp,
					MainInFlow:  &mainInFlow,
					SuperInFlow: &superInFlow,
					BigInFlow:   &bigInFlow,
					MidInFlow:   &midInFlow,
					SmlInFlow:   &smlInFlow,
				},
			},
			LastValidTime:      &lastValidTime,
			LastValidTimestamp: &lastValidTs,
		}), nil
	})

	ctx := context.Background()
	req := &GetCapitalFlowRequest{
		Security:   hkSecurity("00700"),
		PeriodType: 1,
		BeginTime:  "2026-01-01",
		EndTime:    "2026-04-08",
	}
	result, err := GetCapitalFlow(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetCapitalFlow failed: %v", err)
	}

	if len(result.FlowItemList) != 1 {
		t.Fatalf("expected 1 flow item, got %d", len(result.FlowItemList))
	}
	if result.FlowItemList[0].InFlow != 1000000.0 {
		t.Errorf("expected InFlow 1000000.0, got %f", result.FlowItemList[0].InFlow)
	}
	if result.LastValidTime != "2026-04-08" {
		t.Errorf("expected LastValidTime 2026-04-08, got %s", result.LastValidTime)
	}

	server.AssertProtoID(t, ProtoID_GetCapitalFlow)
}

func TestGetCapitalFlow_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetCapitalFlow(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetCapitalFlow_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetCapitalFlow(ctx, cli, &GetCapitalFlowRequest{})
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestGetCapitalFlow_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetCapitalFlow, func(reqBody []byte) (proto.Message, error) {
		return &qotgetcapitalflow.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetCapitalFlow(ctx, cli, &GetCapitalFlowRequest{Security: hkSecurity("00700")})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// GetCapitalDistribution
// =============================================================================

func TestGetCapitalDistribution_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	capInSuper := 1000000.0
	capInBig := 500000.0
	capInMid := 300000.0
	capInSmall := 200000.0
	capOutSuper := 800000.0
	capOutBig := 400000.0
	capOutMid := 200000.0
	capOutSmall := 100000.0
	updateTime := "2026-04-08 15:00:00"
	updateTimestamp := 1775635200.0

	server.RegisterHandler(ProtoID_GetCapitalDistribution, func(reqBody []byte) (proto.Message, error) {
		var req qotgetcapitaldistribution.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetCapitalDistribution request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotgetcapitaldistribution.S2C{
			CapitalInSuper:  &capInSuper,
			CapitalInBig:    &capInBig,
			CapitalInMid:    &capInMid,
			CapitalInSmall:  &capInSmall,
			CapitalOutSuper: &capOutSuper,
			CapitalOutBig:   &capOutBig,
			CapitalOutMid:   &capOutMid,
			CapitalOutSmall: &capOutSmall,
			UpdateTime:      &updateTime,
			UpdateTimestamp: &updateTimestamp,
		}), nil
	})

	ctx := context.Background()
	result, err := GetCapitalDistribution(ctx, cli, hkSecurity("00700"))
	if err != nil {
		t.Fatalf("GetCapitalDistribution failed: %v", err)
	}

	if result.CapitalDistribution.CapitalInSuper != 1000000.0 {
		t.Errorf("expected CapitalInSuper 1000000.0, got %f", result.CapitalDistribution.CapitalInSuper)
	}
	if result.CapitalDistribution.CapitalOutSmall != 100000.0 {
		t.Errorf("expected CapitalOutSmall 100000.0, got %f", result.CapitalDistribution.CapitalOutSmall)
	}

	server.AssertProtoID(t, ProtoID_GetCapitalDistribution)
}

func TestGetCapitalDistribution_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetCapitalDistribution(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestGetCapitalDistribution_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetCapitalDistribution, func(reqBody []byte) (proto.Message, error) {
		return &qotgetcapitaldistribution.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetCapitalDistribution(ctx, cli, hkSecurity("00700"))
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}


// =============================================================================
// Subscribe
// =============================================================================

func TestSubscribe_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_Subscribe, func(reqBody []byte) (proto.Message, error) {
		var req qotsub.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal Subscribe request failed: %v", err)
		}
		if req.C2S == nil {
			t.Fatal("expected non-nil C2S")
		}
		return newSuccessResponse(&qotsub.S2C{}), nil
	})

	ctx := context.Background()
	req := &SubscribeRequest{
		SecurityList:     []*qotcommon.Security{hkSecurity("00700")},
		SubTypeList:      []SubType{SubType_Basic, SubType_KL},
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}
	err := Subscribe(ctx, cli, req)
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	server.AssertProtoID(t, ProtoID_Subscribe)
}

func TestSubscribe_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	err := Subscribe(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestSubscribe_EmptySecurityList(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	err := Subscribe(ctx, cli, &SubscribeRequest{SecurityList: []*qotcommon.Security{}, SubTypeList: []SubType{SubType_Basic}})
	if err == nil {
		t.Error("expected error for empty security list")
	}
}

func TestSubscribe_EmptySubTypeList(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	err := Subscribe(ctx, cli, &SubscribeRequest{SecurityList: []*qotcommon.Security{hkSecurity("00700")}, SubTypeList: []SubType{}})
	if err == nil {
		t.Error("expected error for empty subtype list")
	}
}

func TestSubscribe_WithAllFields(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_Subscribe, func(reqBody []byte) (proto.Message, error) {
		var req qotsub.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal Subscribe request failed: %v", err)
		}
		if req.C2S.IsUnsubAll != nil && *req.C2S.IsUnsubAll {
			t.Error("expected IsUnsubAll=false")
		}
		if req.C2S.IsFirstPush == nil || !*req.C2S.IsFirstPush {
			t.Error("expected IsFirstPush=true")
		}
		return newSuccessResponse(&qotsub.S2C{}), nil
	})

	ctx := context.Background()
	req := &SubscribeRequest{
		SecurityList:         []*qotcommon.Security{hkSecurity("00700")},
		SubTypeList:          []SubType{SubType_Basic, SubType_KL, SubType_OrderBook},
		IsSubOrUnSub:         true,
		IsRegOrUnRegPush:     true,
		RegPushRehabTypeList: []int32{0},
		IsFirstPush:          true,
		IsUnsubAll:           false,
		IsSubOrderBookDetail: true,
		ExtendedTime:         true,
	}
	err := Subscribe(ctx, cli, req)
	if err != nil {
		t.Fatalf("Subscribe with all fields failed: %v", err)
	}

	server.AssertProtoID(t, ProtoID_Subscribe)
}

func TestSubscribe_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_Subscribe, func(reqBody []byte) (proto.Message, error) {
		return &qotsub.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	err := Subscribe(ctx, cli, &SubscribeRequest{
		SecurityList: []*qotcommon.Security{hkSecurity("00700")},
		SubTypeList:  []SubType{SubType_Basic},
	})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// RegQotPush
// =============================================================================

func TestRegQotPush_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_RegQotPush, func(reqBody []byte) (proto.Message, error) {
		var req qotregqotpush.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal RegQotPush request failed: %v", err)
		}
		if req.C2S == nil {
			t.Fatal("expected non-nil C2S")
		}
		return newSuccessResponse(&qotregqotpush.S2C{}), nil
	})

	ctx := context.Background()
	req := &RegQotPushRequest{
		SecurityList:  []*qotcommon.Security{hkSecurity("00700")},
		SubTypeList:   []int32{1, 2},
		RehabTypeList: []int32{0},
		IsRegOrUnReg:  true,
		IsFirstPush:   true,
	}
	err := RegQotPush(ctx, cli, req)
	if err != nil {
		t.Fatalf("RegQotPush failed: %v", err)
	}

	server.AssertProtoID(t, ProtoID_RegQotPush)
}

func TestRegQotPush_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	err := RegQotPush(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestRegQotPush_EmptySecurityList(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	err := RegQotPush(ctx, cli, &RegQotPushRequest{SecurityList: []*qotcommon.Security{}, SubTypeList: []int32{1}})
	if err == nil {
		t.Error("expected error for empty security list")
	}
}

// =============================================================================
// GetSubInfo
// =============================================================================

func TestGetSubInfo_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	totalUsedQuota := int32(10)
	remainQuota := int32(90)
	subType := int32(1)
	usedQuota := int32(5)
	isOwnConnData := true

	server.RegisterHandler(ProtoID_GetSubInfo, func(reqBody []byte) (proto.Message, error) {
		return newSuccessResponse(&qotgetsubinfo.S2C{
			ConnSubInfoList: []*qotcommon.ConnSubInfo{
				{
					SubInfoList: []*qotcommon.SubInfo{
						{SubType: &subType},
					},
					UsedQuota:     &usedQuota,
					IsOwnConnData: &isOwnConnData,
				},
			},
			TotalUsedQuota: &totalUsedQuota,
			RemainQuota:    &remainQuota,
		}), nil
	})

	ctx := context.Background()
	result, err := GetSubInfo(ctx, cli)
	if err != nil {
		t.Fatalf("GetSubInfo failed: %v", err)
	}

	if result.TotalUsedQuota != 10 {
		t.Errorf("expected TotalUsedQuota 10, got %d", result.TotalUsedQuota)
	}
	if result.RemainQuota != 90 {
		t.Errorf("expected RemainQuota 90, got %d", result.RemainQuota)
	}
	if len(result.ConnSubInfoList) != 1 {
		t.Fatalf("expected 1 conn sub info, got %d", len(result.ConnSubInfoList))
	}

	server.AssertProtoID(t, ProtoID_GetSubInfo)
}

func TestGetSubInfo_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetSubInfo, func(reqBody []byte) (proto.Message, error) {
		return &qotgetsubinfo.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetSubInfo(ctx, cli)
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// GetSecuritySnapshot
// =============================================================================

func TestGetSecuritySnapshot_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_GetSecuritySnapshot, func(reqBody []byte) (proto.Message, error) {
		var req qotgetsecuritysnapshot.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetSecuritySnapshot request failed: %v", err)
		}
		if req.C2S == nil || len(req.C2S.SecurityList) == 0 {
			t.Fatal("expected non-empty security list")
		}
		return newSuccessResponse(&qotgetsecuritysnapshot.S2C{
			SnapshotList: []*qotgetsecuritysnapshot.Snapshot{
				{Security: hkSecurity("00700")},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetSecuritySnapshotRequest{
		SecurityList: []*qotcommon.Security{hkSecurity("00700")},
	}
	result, err := GetSecuritySnapshot(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetSecuritySnapshot failed: %v", err)
	}

	if len(result.SnapshotList) != 1 {
		t.Errorf("expected 1 snapshot, got %d", len(result.SnapshotList))
	}

	server.AssertProtoID(t, ProtoID_GetSecuritySnapshot)
}

func TestGetSecuritySnapshot_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetSecuritySnapshot(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetSecuritySnapshot_EmptySecurityList(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetSecuritySnapshot(ctx, cli, &GetSecuritySnapshotRequest{SecurityList: []*qotcommon.Security{}})
	if err == nil {
		t.Error("expected error for empty security list")
	}
}

func TestGetSecuritySnapshot_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetSecuritySnapshot, func(reqBody []byte) (proto.Message, error) {
		return &qotgetsecuritysnapshot.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetSecuritySnapshot(ctx, cli, &GetSecuritySnapshotRequest{
		SecurityList: []*qotcommon.Security{hkSecurity("00700")},
	})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// GetStaticInfo
// =============================================================================

func TestGetStaticInfo_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_GetStaticInfo, func(reqBody []byte) (proto.Message, error) {
		var req qotgetstaticinfo.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetStaticInfo request failed: %v", err)
		}
		if req.C2S == nil {
			t.Fatal("expected non-nil C2S")
		}
		return newSuccessResponse(&qotgetstaticinfo.S2C{
			StaticInfoList: []*qotcommon.SecurityStaticInfo{
				{Basic: &qotcommon.SecurityStaticBasic{Security: hkSecurity("00700"), Name: proto.String("Tencent")}},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetStaticInfoRequest{
		Market:  1,
		SecType: 1,
		SecurityList: []*qotcommon.Security{
			hkSecurity("00700"),
		},
	}
	result, err := GetStaticInfo(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetStaticInfo failed: %v", err)
	}

	if len(result.StaticInfoList) != 1 {
		t.Errorf("expected 1 static info, got %d", len(result.StaticInfoList))
	}

	server.AssertProtoID(t, ProtoID_GetStaticInfo)
}

func TestGetStaticInfo_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetStaticInfo(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetStaticInfo_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_GetStaticInfo, func(reqBody []byte) (proto.Message, error) {
		return &qotgetstaticinfo.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := GetStaticInfo(ctx, cli, &GetStaticInfoRequest{
		Market:       1,
		SecType:      1,
		SecurityList: []*qotcommon.Security{hkSecurity("00700")},
	})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// RequestHistoryKL
// =============================================================================

func TestRequestHistoryKL_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	klTime := "2026-04-08 15:00:00"
	klClosePrice := 350.50
	klOpenPrice := 348.00
	klHighPrice := 352.00
	klLowPrice := 347.00
	klLastClosePrice := 349.00
	klVolume := int64(12345678)
	klTurnover := 4321098765.00
	klTimestamp := 1775635200.0
	klIsBlank := false

	server.RegisterHandler(ProtoID_RequestHistoryKL, func(reqBody []byte) (proto.Message, error) {
		var req qotrequesthistorykl.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal RequestHistoryKL request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected non-nil security")
		}
		return newSuccessResponse(&qotrequesthistorykl.S2C{
			Security: hkSecurity("00700"),
			Name:     proto.String("Tencent"),
			KlList: []*qotcommon.KLine{
				{
					Time:           &klTime,
					ClosePrice:     &klClosePrice,
					OpenPrice:      &klOpenPrice,
					HighPrice:      &klHighPrice,
					LowPrice:       &klLowPrice,
					LastClosePrice: &klLastClosePrice,
					Volume:         &klVolume,
					Turnover:       &klTurnover,
					Timestamp:      &klTimestamp,
					IsBlank:        &klIsBlank,
				},
			},
			NextReqKey: []byte("nextkey"),
		}), nil
	})

	ctx := context.Background()
	req := &RequestHistoryKLRequest{
		RehabType:   0,
		KlType:      4,
		Security:    hkSecurity("00700"),
		BeginTime:   "2026-01-01",
		EndTime:     "2026-04-08",
		MaxAckKLNum: 100,
	}
	result, err := RequestHistoryKL(ctx, cli, req)
	if err != nil {
		t.Fatalf("RequestHistoryKL failed: %v", err)
	}

	if len(result.KLList) != 1 {
		t.Fatalf("expected 1 KLine, got %d", len(result.KLList))
	}
	if result.KLList[0].ClosePrice != 350.50 {
		t.Errorf("expected ClosePrice 350.50, got %f", result.KLList[0].ClosePrice)
	}
	if string(result.NextReqKey) != "nextkey" {
		t.Errorf("expected NextReqKey 'nextkey', got %s", result.NextReqKey)
	}

	server.AssertProtoID(t, ProtoID_RequestHistoryKL)
}

func TestRequestHistoryKL_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := RequestHistoryKL(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestRequestHistoryKL_NilSecurity(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := RequestHistoryKL(ctx, cli, &RequestHistoryKLRequest{})
	if err == nil {
		t.Error("expected error for nil security")
	}
}

func TestRequestHistoryKL_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_RequestHistoryKL, func(reqBody []byte) (proto.Message, error) {
		return &qotrequesthistorykl.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := RequestHistoryKL(ctx, cli, &RequestHistoryKLRequest{Security: hkSecurity("00700")})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// RequestHistoryKLQuota
// =============================================================================

func TestRequestHistoryKLQuota_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_RequestHistoryKLQuota, func(reqBody []byte) (proto.Message, error) {
		var req qotrequesthistoryklquota.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal RequestHistoryKLQuota request failed: %v", err)
		}
		if req.C2S == nil {
			t.Fatal("expected non-nil C2S")
		}
		usedQuota := int32(50)
		remainQuota := int32(450)
		return newSuccessResponse(&qotrequesthistoryklquota.S2C{
			UsedQuota:   &usedQuota,
			RemainQuota: &remainQuota,
		}), nil
	})

	ctx := context.Background()
	req := &RequestHistoryKLQuotaRequest{GetDetail: true}
	result, err := RequestHistoryKLQuota(ctx, cli, req)
	if err != nil {
		t.Fatalf("RequestHistoryKLQuota failed: %v", err)
	}

	if result.UsedQuota != 50 {
		t.Errorf("expected UsedQuota 50, got %d", result.UsedQuota)
	}
	if result.RemainQuota != 450 {
		t.Errorf("expected RemainQuota 450, got %d", result.RemainQuota)
	}

	server.AssertProtoID(t, ProtoID_RequestHistoryKLQuota)
}

func TestRequestHistoryKLQuota_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := RequestHistoryKLQuota(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestRequestHistoryKLQuota_Error(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	retType := int32(common.RetType_RetType_Unknown)
	retMsg := "server error"
	server.RegisterHandler(ProtoID_RequestHistoryKLQuota, func(reqBody []byte) (proto.Message, error) {
		return &qotrequesthistoryklquota.Response{RetType: &retType, RetMsg: &retMsg}, nil
	})

	ctx := context.Background()
	_, err := RequestHistoryKLQuota(ctx, cli, &RequestHistoryKLQuotaRequest{GetDetail: true})
	if err == nil {
		t.Fatal("expected error for non-success RetType")
	}
}

// =============================================================================
// GetOptionQuote
// =============================================================================

func TestGetOptionQuote_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_GetOptionQuote, func(reqBody []byte) (proto.Message, error) {
		return newSuccessResponse(&qotgetoptionquote.S2C{
			OptionQuoteList: []*qotgetoptionquote.OptionQuote{
				{Code: proto.String("HSI2405C35000")},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetOptionQuoteRequest{
		MultiLegs: []*qotcommon.ComboLeg{
			{Side: proto.Int32(1), Options: &qotcommon.Security{}},
		},
	}
	result, err := GetOptionQuote(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetOptionQuote failed: %v", err)
	}

	if len(result.OptionQuoteList) != 1 {
		t.Errorf("expected 1 option quote, got %d", len(result.OptionQuoteList))
	}

	server.AssertProtoID(t, ProtoID_GetOptionQuote)
}

func TestGetOptionQuote_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionQuote(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetOptionQuote_EmptyMultiLegs(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionQuote(ctx, cli, &GetOptionQuoteRequest{MultiLegs: []*qotcommon.ComboLeg{}})
	if err == nil {
		t.Error("expected error for empty MultiLegs")
	}
}

// =============================================================================
// GetOptionStrategy
// =============================================================================

func TestGetOptionStrategy_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	server.RegisterHandler(ProtoID_GetOptionStrategy, func(reqBody []byte) (proto.Message, error) {
		var req qotgetoptionstrategy.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal GetOptionStrategy request failed: %v", err)
		}
		if req.C2S == nil || req.C2S.Owner == nil {
			t.Fatal("expected non-nil owner")
		}
		return newSuccessResponse(&qotgetoptionstrategy.S2C{
			StrategyList: []*qotgetoptionstrategy.OptionStrategyItem{
				{ExpireDate: proto.String("2026-04-30")},
			},
		}), nil
	})

	ctx := context.Background()
	req := &GetOptionStrategyRequest{
		Owner:          hkSecurity("HSI"),
		OptionStrategy: 1,
		ExpireTime:     "2026-04-30",
	}
	result, err := GetOptionStrategy(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetOptionStrategy failed: %v", err)
	}

	if len(result.StrategyList) != 1 {
		t.Errorf("expected 1 strategy, got %d", len(result.StrategyList))
	}

	server.AssertProtoID(t, ProtoID_GetOptionStrategy)
}

func TestGetOptionStrategy_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategy(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetOptionStrategy_NilOwner(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategy(ctx, cli, &GetOptionStrategyRequest{Owner: nil, OptionStrategy: 1})
	if err == nil {
		t.Error("expected error for nil owner")
	}
}

// =============================================================================
// GetOptionStrategyAnalysis
// =============================================================================

func TestGetOptionStrategyAnalysis_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	code := "STRATEGY001"
	name := "Bull Spread"
	bid1 := 2.5
	ask1 := 2.6
	maxProfit := 1000.0
	maxLoss := -500.0
	breakeven := 350.0
	probProfit := 0.65
	delta := 0.5
	theta := -0.1

	server.RegisterHandler(ProtoID_GetOptionStrategyAnalysis, func(reqBody []byte) (proto.Message, error) {
		return newSuccessResponse(&qotgetoptionstrategyanalysis.S2C{
			Code:            &code,
			Name:            &name,
			OptionStrategy:  proto.Int32(1),
			Bid1:            &bid1,
			Ask1:            &ask1,
			MaxProfit:       &maxProfit,
			MaxLoss:         &maxLoss,
			BreakevenPoints: []float64{breakeven},
			ProbOfProfit:    &probProfit,
			Delta:           &delta,
			Theta:           &theta,
		}), nil
	})

	ctx := context.Background()
	req := &GetOptionStrategyAnalysisRequest{
		MultiLegs: []*qotcommon.ComboLeg{
			{Side: proto.Int32(1), Options: &qotcommon.Security{}},
		},
	}
	result, err := GetOptionStrategyAnalysis(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetOptionStrategyAnalysis failed: %v", err)
	}

	if result.Code != "STRATEGY001" {
		t.Errorf("expected Code STRATEGY001, got %s", result.Code)
	}
	if result.MaxProfit != 1000.0 {
		t.Errorf("expected MaxProfit 1000.0, got %f", result.MaxProfit)
	}
	if len(result.BreakevenPoints) != 1 {
		t.Errorf("expected 1 breakeven point, got %d", len(result.BreakevenPoints))
	}

	server.AssertProtoID(t, ProtoID_GetOptionStrategyAnalysis)
}

func TestGetOptionStrategyAnalysis_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategyAnalysis(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetOptionStrategyAnalysis_EmptyMultiLegs(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategyAnalysis(ctx, cli, &GetOptionStrategyAnalysisRequest{MultiLegs: []*qotcommon.ComboLeg{}})
	if err == nil {
		t.Error("expected error for empty MultiLegs")
	}
}

// =============================================================================
// GetOptionStrategySpread
// =============================================================================

func TestGetOptionStrategySpread_Mock(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	spread1 := 0.5
	spread2 := 1.0
	spread3 := 2.0

	server.RegisterHandler(ProtoID_GetOptionStrategySpread, func(reqBody []byte) (proto.Message, error) {
		return newSuccessResponse(&qotgetoptionstrategyspreads.S2C{
			SpreadList: []float64{spread1, spread2, spread3},
		}), nil
	})

	ctx := context.Background()
	req := &GetOptionStrategySpreadRequest{
		Owner:          hkSecurity("HSI"),
		OptionStrategy: 1,
		ExpireTime:     "2026-04-30",
	}
	result, err := GetOptionStrategySpread(ctx, cli, req)
	if err != nil {
		t.Fatalf("GetOptionStrategySpread failed: %v", err)
	}

	if len(result.SpreadList) != 3 {
		t.Errorf("expected 3 spreads, got %d", len(result.SpreadList))
	}
	if result.SpreadList[0] != 0.5 {
		t.Errorf("expected first spread 0.5, got %f", result.SpreadList[0])
	}

	server.AssertProtoID(t, ProtoID_GetOptionStrategySpread)
}

func TestGetOptionStrategySpread_NilRequest(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategySpread(ctx, cli, nil)
	if err == nil {
		t.Error("expected error for nil request")
	}
}

func TestGetOptionStrategySpread_NilOwner(t *testing.T) {
	server := futuapitestutil.NewMockServer(t)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start mock server: %v", err)
	}
	defer server.Stop()

	cli, cleanup := futuapitestutil.NewTestClient(t, server)
	defer cleanup()

	ctx := context.Background()
	_, err := GetOptionStrategySpread(ctx, cli, &GetOptionStrategySpreadRequest{Owner: nil, OptionStrategy: 1})
	if err == nil {
		t.Error("expected error for nil owner")
	}
}
