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

package trd_test

import (
	"context"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
	"github.com/shing1211/futuapi4go/pkg/trd"
	"github.com/shing1211/futuapi4go/test/fixtures"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

func TestGetAccList(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2001, func(req []byte) (proto.Message, error) {
		accID := fixtures.TestAccID
		accType := int32(1 /* TrdType_Security */)
		accStatus := int32(trdcommon.TrdAccStatus_TrdAccStatus_Active)

		s2c := &trdgetacclist.S2C{
			AccList: []*trdcommon.TrdAcc{
				{
					AccID:     &accID,
					TrdEnv:    proto.Int32(fixtures.TestTrdEnv),
					AccType:   &accType,
					AccStatus: &accStatus,
				},
			},
		}

		return &trdgetacclist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	result, err := trd.GetAccList(context.Background(), cli, constant.TrdCategory_Security, false)
	if err != nil {
		t.Fatalf("GetAccList failed: %v", err)
	}

	if len(result.AccList) != 1 {
		t.Errorf("Expected 1 account, got %d", len(result.AccList))
	}

	acc := result.AccList[0]
	if acc.AccID != fixtures.TestAccID {
		t.Errorf("Expected AccID %d, got %d", fixtures.TestAccID, acc.AccID)
	}

	server.AssertProtoID(t, 2001)
}

func TestUnlockTrade(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2005, func(req []byte) (proto.Message, error) {
		s2c := &trdunlocktrade.S2C{}
		return &trdunlocktrade.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// MD5 hash of "test123"
	pwdMD5 := "202cb962ac59075b964b07152d234b70"

	req := &trd.UnlockTradeRequest{
		Unlock: true,
		PwdMD5: constant.SensitiveString(pwdMD5),
	}

	err := trd.UnlockTrade(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("UnlockTrade failed: %v", err)
	}

	server.AssertProtoID(t, 2005)
}

func TestGetFunds_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2101, func(req []byte) (proto.Message, error) {
		funds := fixtures.HSIFunds()
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdgetfunds.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			Funds: funds,
		}

		return &trdgetfunds.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.GetFundsRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
	}

	result, err := trd.GetFunds(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetFunds failed: %v", err)
	}

	if result.Funds.TotalAssets <= 0 {
		t.Error("TotalAssets should be positive")
	}

	if result.Funds.Power <= 0 {
		t.Error("Power should be positive")
	}

	// Validate realistic values
	if result.Funds.TotalAssets != 500000.0 {
		t.Errorf("Expected TotalAssets 500000, got %f", result.Funds.TotalAssets)
	}

	server.AssertProtoID(t, 2101)
}

func TestGetPositionList_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2102, func(req []byte) (proto.Message, error) {
		position := fixtures.HSIPosition()
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdgetpositionlist.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			PositionList: []*trdcommon.Position{position},
		}

		return &trdgetpositionlist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.GetPositionListRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
	}

	result, err := trd.GetPositionList(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetPositionList failed: %v", err)
	}

	if len(result.PositionList) != 1 {
		t.Errorf("Expected 1 position, got %d", len(result.PositionList))
	}

	pos := result.PositionList[0]
	if pos.Code != fixtures.HSIFuturesCode {
		t.Errorf("Expected code %s, got %s", fixtures.HSIFuturesCode, pos.Code)
	}

	if pos.Qty != 2 {
		t.Errorf("Expected qty 2, got %f", pos.Qty)
	}

	// Validate P/L calculation
	expectedPL := (18523.45 - 18480.00) * 2
	if pos.PlVal != expectedPL {
		t.Errorf("Expected P/L %f, got %f", expectedPL, pos.PlVal)
	}

	server.AssertProtoID(t, 2102)
}

func TestPlaceOrder_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2202, func(req []byte) (proto.Message, error) {
		var reqMsg trdplaceorder.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		// Validate request
		if reqMsg.C2S.GetPrice() <= 0 {
			t.Error("Price should be positive")
		}

		if reqMsg.C2S.GetQty() <= 0 {
			t.Error("Qty should be positive")
		}

		orderID := uint64(9876543210)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdplaceorder.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderID: &orderID,
		}

		return &trdplaceorder.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.PlaceOrderRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
		Code:      fixtures.HSIFuturesCode,
		TrdSide:   constant.TrdSide_Buy,
		OrderType: constant.OrderType_Normal,
		Price:     18520.00,
		Qty:       1,
	}

	result, err := trd.PlaceOrder(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("PlaceOrder failed: %v", err)
	}

	if result.OrderID != 9876543210 {
		t.Errorf("Expected OrderID 9876543210, got %d", result.OrderID)
	}

	server.AssertProtoID(t, 2202)
}

func TestPlaceOrder_HSI_Sell(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2202, func(req []byte) (proto.Message, error) {
		orderID := uint64(9876543211)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdplaceorder.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderID: &orderID,
		}

		return &trdplaceorder.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.PlaceOrderRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
		Code:      fixtures.HSIFuturesCode,
		TrdSide:   constant.TrdSide_Sell,
		OrderType: constant.OrderType_Normal,
		Price:     18530.00,
		Qty:       1,
	}

	result, err := trd.PlaceOrder(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("PlaceOrder (sell) failed: %v", err)
	}

	if result.OrderID != 9876543211 {
		t.Errorf("Expected OrderID 9876543211, got %d", result.OrderID)
	}

	server.AssertProtoID(t, 2202)
}

func TestGetOrderList(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2201, func(req []byte) (proto.Message, error) {
		order1 := fixtures.HSIOrder(1001)
		order2 := fixtures.HSIOrder(1002)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdgetorderlist.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderList: []*trdcommon.Order{order1, order2},
		}

		return &trdgetorderlist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.GetOrderListRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
	}

	result, err := trd.GetOrderList(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}

	if len(result.OrderList) != 2 {
		t.Errorf("Expected 2 orders, got %d", len(result.OrderList))
	}

	// Validate order fields
	for _, order := range result.OrderList {
		if order.Price <= 0 {
			t.Errorf("Order %d has invalid price: %f", order.OrderID, order.Price)
		}

		if order.Qty <= 0 {
			t.Errorf("Order %d has invalid qty: %f", order.OrderID, order.Qty)
		}

		if order.Code != fixtures.HSIFuturesCode {
			t.Errorf("Order has wrong code: %s", order.Code)
		}
	}

	server.AssertProtoID(t, 2201)
}

func TestModifyOrder_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2205, func(req []byte) (proto.Message, error) {
		var reqMsg trdmodifyorder.Request
		if err := proto.Unmarshal(req, &reqMsg); err != nil {
			return nil, err
		}

		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)
		orderID := uint64(1001)

		s2c := &trdmodifyorder.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderID: &orderID,
		}
		return &trdmodifyorder.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// Test modify order price
	req := &trd.ModifyOrderRequest{
		AccID:         fixtures.TestAccID,
		TrdMarket:     fixtures.TestTrdMkt,
		OrderID:       1001,
		ModifyOrderOp: constant.ModifyOrderOp_Normal,
		Price:         18525.00,
		Qty:           0, // Keep unchanged
	}

	_, err := trd.ModifyOrder(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("ModifyOrder failed: %v", err)
	}

	server.AssertProtoID(t, 2205)
}

func TestModifyOrder_Cancel(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2205, func(req []byte) (proto.Message, error) {
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)
		orderID := uint64(1001)

		s2c := &trdmodifyorder.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderID: &orderID,
		}
		return &trdmodifyorder.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// Test cancel order
	req := &trd.ModifyOrderRequest{
		AccID:         fixtures.TestAccID,
		TrdMarket:     fixtures.TestTrdMkt,
		OrderID:       1001,
		ModifyOrderOp: constant.ModifyOrderOp_Cancel,
		Price:         0,
		Qty:           0,
	}

	_, err := trd.ModifyOrder(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("CancelOrder failed: %v", err)
	}

	server.AssertProtoID(t, 2205)
}

func TestGetOrderFillList_HSI(t *testing.T) {
	server := testutil.NewMockServer(t)

	server.RegisterHandler(2211, func(req []byte) (proto.Message, error) {
		fill1 := fixtures.HSIOrderFill(5001, 1001)
		fill2 := fixtures.HSIOrderFill(5002, 1002)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)

		s2c := &trdgetorderfilllist.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderFillList: []*trdcommon.OrderFill{fill1, fill2},
		}

		return &trdgetorderfilllist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	req := &trd.GetOrderFillListRequest{
		AccID:     fixtures.TestAccID,
		TrdMarket: fixtures.TestTrdMkt,
	}

	result, err := trd.GetOrderFillList(context.Background(), cli, req)
	if err != nil {
		t.Fatalf("GetOrderFillList failed: %v", err)
	}

	if len(result.OrderFillList) != 2 {
		t.Errorf("Expected 2 fills, got %d", len(result.OrderFillList))
	}

	// Validate fill data
	for _, fill := range result.OrderFillList {
		if fill.Price <= 0 {
			t.Errorf("Fill has invalid price: %f", fill.Price)
		}

		if fill.Qty <= 0 {
			t.Errorf("Fill has invalid qty: %f", fill.Qty)
		}

		if fill.Code != fixtures.HSIFuturesCode {
			t.Errorf("Fill has wrong code: %s", fill.Code)
		}
	}

	server.AssertProtoID(t, 2211)
}

func TestTradingWorkflow_Complete(t *testing.T) {
	server := testutil.NewMockServer(t)

	// Register all handlers
server.RegisterHandler(2001, func(req []byte) (proto.Message, error) {
		accID := fixtures.TestAccID
		s2c := &trdgetacclist.S2C{
			AccList: []*trdcommon.TrdAcc{
				{
					AccID:     &accID,
					TrdEnv:    proto.Int32(fixtures.TestTrdEnv),
				},
			},
		}
		return &trdgetacclist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	server.RegisterHandler(2005, func(req []byte) (proto.Message, error) {
		return &trdunlocktrade.Response{S2C: &trdunlocktrade.S2C{}, RetType: proto.Int32(0)}, nil
	})

	server.RegisterHandler(2101, func(req []byte) (proto.Message, error) {
		funds := fixtures.HSIFunds()
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)
		s2c := &trdgetfunds.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			Funds: funds,
		}
		return &trdgetfunds.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	server.RegisterHandler(2202, func(req []byte) (proto.Message, error) {
		orderID := uint64(9999)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)
		s2c := &trdplaceorder.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderID: &orderID,
		}
		return &trdplaceorder.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	server.RegisterHandler(2201, func(req []byte) (proto.Message, error) {
		order := fixtures.HSIOrder(9999)
		trdEnv := int32(fixtures.TestTrdEnv)
		accID := fixtures.TestAccID
		trdMarket := int32(fixtures.TestTrdMkt)
		s2c := &trdgetorderlist.S2C{
			Header: &trdcommon.TrdHeader{
				TrdEnv:    &trdEnv,
				AccID:     &accID,
				TrdMarket: &trdMarket,
			},
			OrderList: []*trdcommon.Order{order},
		}
		return &trdgetorderlist.Response{S2C: s2c, RetType: proto.Int32(0)}, nil
	})

	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer server.Stop()

	cli, cleanup := testutil.NewTestClient(t, server)
	defer cleanup()

	// Clear requests from registration/setup phase
	server.ClearRequests()

	// Step 1: Get account list
	trdCategory := constant.TrdCategory(1)
	accList, err := trd.GetAccList(context.Background(), cli, trdCategory, false)
	if err != nil {
		t.Fatalf("Step 1 - GetAccList failed: %v", err)
	}
	accID := accList.AccList[0].AccID

	// Step 2: Unlock trade
	err = trd.UnlockTrade(context.Background(), cli, &trd.UnlockTradeRequest{
		Unlock: true,
		PwdMD5: "test123",
	})
	if err != nil {
		t.Fatalf("Step 2 - UnlockTrade failed: %v", err)
	}

	// Step 3: Check funds
	funds, err := trd.GetFunds(context.Background(), cli, &trd.GetFundsRequest{
		AccID:     accID,
		TrdMarket: fixtures.TestTrdMkt,
	})
	if err != nil {
		t.Fatalf("Step 3 - GetFunds failed: %v", err)
	}

	if funds.Funds.Power < 100000 {
		t.Errorf("Insufficient buying power: %f", funds.Funds.Power)
	}

	// Step 4: Place order
	_, err = trd.PlaceOrder(context.Background(), cli, &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: fixtures.TestTrdMkt,
		Code:      fixtures.HSIFuturesCode,
		TrdSide:   constant.TrdSide_Buy,
		OrderType: constant.OrderType_Normal,
		Price:     18500.00,
		Qty:       1,
	})
	if err != nil {
		t.Fatalf("Step 4 - PlaceOrder failed: %v", err)
	}

	// Step 5: Verify order
	orders, err := trd.GetOrderList(context.Background(), cli, &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: fixtures.TestTrdMkt,
	})
	if err != nil {
		t.Fatalf("Step 5 - GetOrderList failed: %v", err)
	}

	if len(orders.OrderList) != 1 {
		t.Errorf("Expected 1 order after placement, got %d", len(orders.OrderList))
	}

	// Verify complete workflow
	server.AssertRequestCount(t, 5)
	if !server.HasProtoID(2001) || !server.HasProtoID(2101) || !server.HasProtoID(2202) || !server.HasProtoID(2201) {
		t.Errorf("Missing expected proto IDs in workflow requests")
	}
	t.Logf("Complete trading workflow executed successfully with %d API calls", len(server.GetRequests()))
}
