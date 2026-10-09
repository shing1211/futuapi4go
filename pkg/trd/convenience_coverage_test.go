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

package trd

import (
	"context"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/trdflowsummary"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
)

// ============================================================================
// Query APIs not exercised elsewhere
// ============================================================================

func TestCov_GetOrderList(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetOrderList, &trdgetorderlist.Response{})
	req := &GetOrderListRequest{}
	fillNonZero(req)
	if _, err := GetOrderList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOrderList: %v", err)
	}
}

func TestCov_GetOrderList_NilReq(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetOrderList, &trdgetorderlist.Response{})
	if _, err := GetOrderList(context.Background(), cli, nil); err == nil {
		t.Fatal("expected error for nil request")
	}
}

func TestCov_GetFlowSummary(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetFlowSummary, &trdflowsummary.Response{})
	req := &GetFlowSummaryRequest{}
	fillNonZero(req)
	if _, err := GetFlowSummary(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFlowSummary: %v", err)
	}
}

func TestCov_GetFlowSummary_NilReq(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetFlowSummary, &trdflowsummary.Response{})
	if _, err := GetFlowSummary(context.Background(), cli, nil); err == nil {
		t.Fatal("expected error for nil request")
	}
}

// ============================================================================
// convenience.go
// ============================================================================

// NOTE: CancelAllOrders has no success-path test because the implementation
// omits the required Trd_ModifyOrder.C2S.PacketID field; see issue filed by
// the coverage work. The zero-account validation path is covered below.
func TestCov_CancelAllOrders_NoAcc(t *testing.T) {
	cli := trdTestClient(t, ProtoID_ModifyOrder, &trdmodifyorder.Response{})
	if _, err := CancelAllOrders(context.Background(), cli, 0, constant.TrdMarket_HK, constant.TrdEnv_Real); err == nil {
		t.Fatal("expected error for zero accID")
	}
}

func TestCov_QuickBuy(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceOrder, &trdplaceorder.Response{})
	if _, err := QuickBuy(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real, "HK.00700", 100, 350.5); err != nil {
		t.Fatalf("QuickBuy: %v", err)
	}
}

func TestCov_QuickSell(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceOrder, &trdplaceorder.Response{})
	if _, err := QuickSell(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real, "HK.00700", 100, 350.5); err != nil {
		t.Fatalf("QuickSell: %v", err)
	}
}

func TestCov_QuickMarketBuy(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceOrder, &trdplaceorder.Response{})
	if _, err := QuickMarketBuy(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real, "HK.00700", 100); err != nil {
		t.Fatalf("QuickMarketBuy: %v", err)
	}
}

func TestCov_QuickMarketSell(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceOrder, &trdplaceorder.Response{})
	if _, err := QuickMarketSell(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real, "HK.00700", 100); err != nil {
		t.Fatalf("QuickMarketSell: %v", err)
	}
}

func TestCov_QuickBuy_Validation(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceOrder, &trdplaceorder.Response{})
	for _, tc := range []struct {
		name  string
		accID uint64
		code  string
		qty   float64
		price float64
	}{
		{"zero acc", 0, "HK.00700", 100, 350.5},
		{"empty code", 1, "", 100, 350.5},
		{"zero qty", 1, "HK.00700", 0, 350.5},
		{"zero price", 1, "HK.00700", 100, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := QuickBuy(context.Background(), cli, tc.accID, constant.TrdMarket_HK, constant.TrdEnv_Real, tc.code, tc.qty, tc.price); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestCov_GetPositions(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetPositionList, &trdgetpositionlist.Response{})
	if _, err := GetPositions(context.Background(), cli, 1); err != nil {
		t.Fatalf("GetPositions: %v", err)
	}
}

func TestCov_GetTodayFills(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetOrderFillList, &trdgetorderfilllist.Response{})
	if _, err := GetTodayFills(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real); err != nil {
		t.Fatalf("GetTodayFills: %v", err)
	}
}

func TestCov_GetTodayOrders(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetOrderList, &trdgetorderlist.Response{})
	if _, err := GetTodayOrders(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real); err != nil {
		t.Fatalf("GetTodayOrders: %v", err)
	}
}

func TestCov_GetAccountFunds(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetFunds, &trdgetfunds.Response{})
	if _, err := GetAccountFunds(context.Background(), cli, 1, constant.TrdMarket_HK, constant.TrdEnv_Real); err != nil {
		t.Fatalf("GetAccountFunds: %v", err)
	}
}

// ============================================================================
// builder.go methods not covered elsewhere
// ============================================================================

func TestCov_Builder_TrailingAndSecondaryMarket(t *testing.T) {
	b := NewOrder(1, constant.TrdMarket_HK, constant.TrdEnv_Real).
		Buy("HK.00700", 100).
		At(350.5).
		WithSecMarket(constant.TrdSecMarket_HK).
		WithTrailType(constant.TrailType_Ratio).
		WithTrailValue(1.5).
		WithSpread(0.5)

	req, err := b.Build()
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if req.SecMarket != constant.TrdSecMarket_HK {
		t.Errorf("SecMarket = %d, want %d", req.SecMarket, constant.TrdSecMarket_HK)
	}
	if req.TrailType != constant.TrailType_Ratio {
		t.Errorf("TrailType = %d, want %d", req.TrailType, constant.TrailType_Ratio)
	}
	if req.TrailValue != 1.5 {
		t.Errorf("TrailValue = %v, want 1.5", req.TrailValue)
	}
	if req.TrailSpread != 0.5 {
		t.Errorf("TrailSpread = %v, want 0.5", req.TrailSpread)
	}
}
