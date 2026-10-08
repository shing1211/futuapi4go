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

	"github.com/shing1211/futuapi4go/pkg/pb/trdgetcombomaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplacecomboorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdreconfirmorder"
)

// ============================================================================
// Write/query APIs: full success path via fillNonZero-populated requests
// ============================================================================

func TestCov_ModifyOrder_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_ModifyOrder, &trdmodifyorder.Response{})
	req := &ModifyOrderRequest{}
	fillNonZero(req)
	if _, err := ModifyOrder(context.Background(), cli, req); err != nil {
		t.Fatalf("ModifyOrder: %v", err)
	}
}

func TestCov_ModifyOrder_NoOrderID(t *testing.T) {
	cli := trdTestClient(t, ProtoID_ModifyOrder, &trdmodifyorder.Response{})
	req := &ModifyOrderRequest{AccID: 1, ModifyOrderOp: 1}
	if _, err := ModifyOrder(context.Background(), cli, req); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCov_ReconfirmOrder_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_ReconfirmOrder, &trdreconfirmorder.Response{})
	req := &ReconfirmOrderRequest{}
	fillNonZero(req)
	if _, err := ReconfirmOrder(context.Background(), cli, req); err != nil {
		t.Fatalf("ReconfirmOrder: %v", err)
	}
}

func TestCov_GetMarginRatio_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetMarginRatio, &trdgetmarginratio.Response{})
	req := &GetMarginRatioRequest{}
	fillNonZero(req)
	if _, err := GetMarginRatio(context.Background(), cli, req); err != nil {
		t.Fatalf("GetMarginRatio: %v", err)
	}
}

func TestCov_GetMaxTrdQtys_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetMaxTrdQtys, &trdgetmaxtrdqtys.Response{})
	req := &GetMaxTrdQtysRequest{}
	fillNonZero(req)
	if _, err := GetMaxTrdQtys(context.Background(), cli, req); err != nil {
		t.Fatalf("GetMaxTrdQtys: %v", err)
	}
}

func TestCov_GetOrderFee_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetOrderFee, &trdgetorderfee.Response{})
	req := &GetOrderFeeRequest{}
	fillNonZero(req)
	if _, err := GetOrderFee(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOrderFee: %v", err)
	}
}

func TestCov_GetHistoryOrderList_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetHistoryOrderList, &trdgethistoryorderlist.Response{})
	req := &GetHistoryOrderListRequest{}
	fillNonZero(req)
	if _, err := GetHistoryOrderList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHistoryOrderList: %v", err)
	}
}

func TestCov_GetHistoryOrderFillList_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetHistoryOrderFillList, &trdgethistoryorderfilllist.Response{})
	req := &GetHistoryOrderFillListRequest{}
	fillNonZero(req)
	if _, err := GetHistoryOrderFillList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHistoryOrderFillList: %v", err)
	}
}

func TestCov_PlaceComboOrder_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_PlaceComboOrder, &trdplacecomboorder.Response{})
	req := &PlaceComboOrderRequest{}
	fillNonZero(req)
	if _, err := PlaceComboOrder(context.Background(), cli, req); err != nil {
		t.Fatalf("PlaceComboOrder: %v", err)
	}
}

func TestCov_GetComboMaxTrdQtys_Full(t *testing.T) {
	cli := trdTestClient(t, ProtoID_GetComboMaxTrdQtys, &trdgetcombomaxtrdqtys.Response{})
	req := &GetComboMaxTrdQtysRequest{}
	fillNonZero(req)
	if _, err := GetComboMaxTrdQtys(context.Background(), cli, req); err != nil {
		t.Fatalf("GetComboMaxTrdQtys: %v", err)
	}
}

// ============================================================================
// internal helpers
// ============================================================================

func TestCov_HeaderAndC2SPools(t *testing.T) {
	h := getTrdHeader()
	if h == nil {
		t.Fatal("getTrdHeader returned nil")
	}
	putTrdHeader(h)

	c2s := getPlaceOrderC2S()
	if c2s == nil {
		t.Fatal("getPlaceOrderC2S returned nil")
	}
	putPlaceOrderC2S(c2s)
}

func TestCov_WrapError(t *testing.T) {
	if err := wrapError("PlaceOrder", 1, "boom"); err == nil {
		t.Fatal("wrapError returned nil")
	}
}

func TestCov_ValidationWarningError(t *testing.T) {
	w := ValidationWarning{Field: "price", Severity: SeverityError, Message: "boom"}
	if w.Error() == "" {
		t.Fatal("empty error string")
	}
}
