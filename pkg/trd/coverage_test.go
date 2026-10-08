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
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetcombomaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplacecomboorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdreconfirmorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

func okRet() *int32 {
	v := int32(common.RetType_RetType_Succeed)
	return &v
}
func i32Ptr(i int32) *int32   { return &i }
func u64Ptr(u uint64) *uint64 { return &u }
func u32Ptr(u uint32) *uint32 { return &u }
func strPtr(s string) *string { return &s }

// TestUnlockTrade_API exercises UnlockTrade with MockServer.
func TestUnlockTrade_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2005, func(reqBody []byte) (proto.Message, error) {
		var req trdunlocktrade.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetUnlock()
		return &trdunlocktrade.Response{RetType: okRet(), S2C: &trdunlocktrade.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *UnlockTradeRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &UnlockTradeRequest{Unlock: true, PwdMD5: constant.SensitiveString("abc123")}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"empty pwd", &UnlockTradeRequest{Unlock: true, PwdMD5: constant.SensitiveString("")}, true, "password MD5 is required"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			err := UnlockTrade(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2005)
		})
	}
}

// TestSubAccPush_API exercises SubAccPush with MockServer.
func TestSubAccPush_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2008, func(reqBody []byte) (proto.Message, error) {
		return &trdunlocktrade.Response{RetType: okRet(), S2C: &trdunlocktrade.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *SubAccPushRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &SubAccPushRequest{AccIDList: []uint64{111111}}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"empty acclist", &SubAccPushRequest{AccIDList: []uint64{}}, true, "account ID list is empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			err := SubAccPush(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2008)
		})
	}
}

// TestGetAccList_API exercises GetAccList with MockServer.
func TestGetAccList_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2001, func(reqBody []byte) (proto.Message, error) {
		accID := u64Ptr(111111)
		trdEnv := i32Ptr(1)
		accType := i32Ptr(1)
		return &trdgetacclist.Response{
			RetType: okRet(),
			S2C: &trdgetacclist.S2C{
				AccList: []*trdcommon.TrdAcc{
					{AccID: accID, TrdEnv: trdEnv, AccType: accType},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	rsp, err := GetAccList(context.Background(), cli, constant.TrdCategory_Security, false)
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(rsp.AccList) != 1 {
		t.Errorf("AccList: want 1, got %d", len(rsp.AccList))
	}
	srv.AssertProtoID(t, 2001)
}

// TestGetFunds_API exercises GetFunds with MockServer.
func TestGetFunds_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2101, func(reqBody []byte) (proto.Message, error) {
		trdEnv := i32Ptr(1)
		accID := u64Ptr(111111)
		trdMkt := i32Ptr(1)
		hdr := &trdcommon.TrdHeader{TrdEnv: trdEnv, AccID: accID, TrdMarket: trdMkt}
		power, totalAssets := 10000.00, 15000.00
		cash, marketVal := 8000.00, 5000.00
		currency := i32Ptr(1)
		return &trdgetfunds.Response{
			RetType: okRet(),
			S2C: &trdgetfunds.S2C{
				Header: hdr,
				Funds: &trdcommon.Funds{
					Power:       &power,
					TotalAssets: &totalAssets,
					Cash:        &cash,
					MarketVal:   &marketVal,
					Currency:    currency,
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetFundsRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &GetFundsRequest{AccID: 111111, TrdMarket: constant.TrdMarket_HK, TrdEnv: constant.TrdEnv_Simulate}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetFundsRequest{AccID: 0, TrdMarket: constant.TrdMarket_HK, TrdEnv: constant.TrdEnv_Simulate}, true, "account ID is required"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetFunds(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if rsp.Funds == nil {
				t.Fatal("Funds should not be nil")
			}
			srv.AssertProtoID(t, 2101)
		})
	}
}

// TestGetPositionList_API exercises GetPositionList with MockServer.
func TestGetPositionList_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2102, func(reqBody []byte) (proto.Message, error) {
		trdEnv := i32Ptr(1)
		accID := u64Ptr(111111)
		trdMkt := i32Ptr(1)
		code := strPtr("00700")
		name := strPtr("Tencent")
		qty := 100.0
		price := 350.0
		cost := 34000.0
		plRatio := 0.03
		return &trdgetpositionlist.Response{
			RetType: okRet(),
			S2C: &trdgetpositionlist.S2C{
				PositionList: []*trdcommon.Position{
					{Code: code, Name: name, Qty: &qty, Price: &price, CostPrice: &cost, PlRatio: &plRatio},
				},
				Header: &trdcommon.TrdHeader{TrdEnv: trdEnv, AccID: accID, TrdMarket: trdMkt},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetPositionListRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &GetPositionListRequest{AccID: 111111, TrdMarket: constant.TrdMarket_HK, TrdEnv: constant.TrdEnv_Simulate}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"all zero", &GetPositionListRequest{}, true, "at least one of AccID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetPositionList(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.name == "basic" && len(rsp.PositionList) == 0 {
				t.Error("expected non-empty position list")
			}
			srv.AssertProtoID(t, 2102)
		})
	}
}

// TestGetOrderFillList_API exercises GetOrderFillList with MockServer.
func TestGetOrderFillList_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2211, func(reqBody []byte) (proto.Message, error) {
		trdEnv := i32Ptr(1)
		accID := u64Ptr(111111)
		trdMkt := i32Ptr(1)
		fillID, orderID := u64Ptr(1), u64Ptr(99999)
		code := strPtr("00700")
		name := strPtr("Tencent")
		price := 350.00
		qty := 100.0
		trdSide := i32Ptr(1)
		createTime := strPtr("2026-04-08 10:00:00")
		return &trdgetorderfilllist.Response{
			RetType: okRet(),
			S2C: &trdgetorderfilllist.S2C{
				OrderFillList: []*trdcommon.OrderFill{
					{FillID: fillID, OrderID: orderID, Code: code, Name: name, Price: &price, Qty: &qty, TrdSide: trdSide, CreateTime: createTime},
				},
				Header: &trdcommon.TrdHeader{TrdEnv: trdEnv, AccID: accID, TrdMarket: trdMkt},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetOrderFillListRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &GetOrderFillListRequest{AccID: 111111, TrdMarket: constant.TrdMarket_HK, TrdEnv: constant.TrdEnv_Simulate}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetOrderFillListRequest{}, true, "invalid account ID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetOrderFillList(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.name == "basic" && len(rsp.OrderFillList) == 0 {
				t.Error("expected non-empty fill list")
			}
			srv.AssertProtoID(t, 2211)
		})
	}
}

// TestPlaceOrder_API exercises PlaceOrder with MockServer.
func TestPlaceOrder_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2202, func(reqBody []byte) (proto.Message, error) {
		var req trdplaceorder.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		orderID := u64Ptr(99999)
		orderIDStr := strPtr("99999")
		return &trdplaceorder.Response{
			RetType: okRet(),
			S2C: &trdplaceorder.S2C{
				OrderID: orderID, OrderIDEx: orderIDStr,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *PlaceOrderRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := PlaceOrder(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2202)
		})
	}
}

// TestModifyOrder_API exercises ModifyOrder with MockServer.
func TestModifyOrder_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2205, func(reqBody []byte) (proto.Message, error) {
		var req trdmodifyorder.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		orderID := u64Ptr(99999)
		orderIDStr := strPtr("99999")
		return &trdmodifyorder.Response{
			RetType: okRet(),
			S2C: &trdmodifyorder.S2C{
				OrderID: orderID, OrderIDEx: orderIDStr,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *ModifyOrderRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := ModifyOrder(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2205)
		})
	}
}

// TestGetOrderFee_API exercises GetOrderFee with MockServer.
func TestGetOrderFee_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2225, func(reqBody []byte) (proto.Message, error) {
		var req trdgetorderfee.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdgetorderfee.Response{RetType: okRet(), S2C: &trdgetorderfee.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetOrderFeeRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetOrderFeeRequest{AccID: 0, TrdMarket: 1, OrderIDExList: []string{"x"}}, true, "invalid account ID"},
		{"empty orderids", &GetOrderFeeRequest{AccID: 111111, TrdMarket: 1, OrderIDExList: []string{}}, true, "order ID list is empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := GetOrderFee(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2225)
		})
	}
}

// TestGetMarginRatio_API exercises GetMarginRatio with MockServer.
func TestGetMarginRatio_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2223, func(reqBody []byte) (proto.Message, error) {
		var req trdgetmarginratio.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdgetmarginratio.Response{RetType: okRet(), S2C: &trdgetmarginratio.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	hkMkt := int32(constant.TrdMarket_HK)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetMarginRatioRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetMarginRatioRequest{AccID: 0, TrdMarket: constant.TrdMarket_HK, SecurityList: []*qotcommon.Security{sec}}, true, "invalid account ID"},
		{"empty seclist", &GetMarginRatioRequest{AccID: 111111, TrdMarket: constant.TrdMarket_HK, SecurityList: []*qotcommon.Security{}}, true, "security list is empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := GetMarginRatio(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2223)
		})
	}
}

// TestGetMaxTrdQtys_API exercises GetMaxTrdQtys with MockServer.
func TestGetMaxTrdQtys_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2111, func(reqBody []byte) (proto.Message, error) {
		var req trdgetmaxtrdqtys.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		maxBuy := 1000.0
		return &trdgetmaxtrdqtys.Response{
			RetType: okRet(),
			S2C: &trdgetmaxtrdqtys.S2C{
				MaxTrdQtys: &trdcommon.MaxTrdQtys{MaxCashBuy: &maxBuy},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetMaxTrdQtysRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetMaxTrdQtysRequest{AccID: 0, Code: "00700"}, true, "invalid account ID"},
		{"empty code", &GetMaxTrdQtysRequest{AccID: 111111, Code: ""}, true, "security code is required"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetMaxTrdQtys(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if rsp.MaxTrdQtys == nil {
				t.Error("MaxTrdQtys should not be nil")
			}
			srv.AssertProtoID(t, 2111)
		})
	}
}

// TestGetHistoryOrderList_API exercises GetHistoryOrderList with MockServer.
func TestGetHistoryOrderList_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2221, func(reqBody []byte) (proto.Message, error) {
		var req trdgethistoryorderlist.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdgethistoryorderlist.Response{RetType: okRet(), S2C: &trdgethistoryorderlist.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetHistoryOrderListRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetHistoryOrderListRequest{}, true, "invalid account ID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := GetHistoryOrderList(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2221)
		})
	}
}

// TestGetHistoryOrderFillList_API exercises GetHistoryOrderFillList with MockServer.
func TestGetHistoryOrderFillList_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2222, func(reqBody []byte) (proto.Message, error) {
		var req trdgethistoryorderfilllist.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdgethistoryorderfilllist.Response{RetType: okRet(), S2C: &trdgethistoryorderfilllist.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	for _, tc := range []struct {
		name      string
		req       *GetHistoryOrderFillListRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero accid", &GetHistoryOrderFillListRequest{}, true, "invalid account ID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := GetHistoryOrderFillList(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2222)
		})
	}
}

// TestGetComboMaxTrdQtys_API exercises GetComboMaxTrdQtys with MockServer.
func TestGetComboMaxTrdQtys_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2112, func(reqBody []byte) (proto.Message, error) {
		var req trdgetcombomaxtrdqtys.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		maxBuy := 500.0
		return &trdgetcombomaxtrdqtys.Response{
			RetType: okRet(),
			S2C: &trdgetcombomaxtrdqtys.S2C{
				MaxTrdQtys: &trdcommon.ComboMaxTrdQtys{OptionBuyPower: &maxBuy},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	trdEnv := i32Ptr(1)
	accID := u64Ptr(111111)
	trdMkt := i32Ptr(1)
	header := &trdcommon.TrdHeader{TrdEnv: trdEnv, AccID: accID, TrdMarket: trdMkt}

	for _, tc := range []struct {
		name      string
		req       *GetComboMaxTrdQtysRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"nil header", &GetComboMaxTrdQtysRequest{Header: nil, Qty: 1.0}, true, "Header is nil"},
		{"empty legs", &GetComboMaxTrdQtysRequest{Header: header, ComboLegs: []*qotcommon.ComboLeg{}, Qty: 1.0}, true, "ComboLegs is empty"},
		{"zero qty", &GetComboMaxTrdQtysRequest{Header: header, ComboLegs: []*qotcommon.ComboLeg{{}}, Qty: 0}, true, "Qty must be positive"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := GetComboMaxTrdQtys(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2112)
		})
	}
}

// TestPlaceComboOrder_API exercises PlaceComboOrder with MockServer.
func TestPlaceComboOrder_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2227, func(reqBody []byte) (proto.Message, error) {
		var req trdplacecomboorder.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdplacecomboorder.Response{RetType: okRet(), S2C: &trdplacecomboorder.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	trdEnv := i32Ptr(1)
	accID := u64Ptr(111111)
	trdMkt := i32Ptr(1)
	header := &trdcommon.TrdHeader{TrdEnv: trdEnv, AccID: accID, TrdMarket: trdMkt}

	for _, tc := range []struct {
		name      string
		req       *PlaceComboOrderRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"nil header", &PlaceComboOrderRequest{Header: nil, Qty: 1.0}, true, "Header is nil"},
		{"empty legs", &PlaceComboOrderRequest{Header: header, ComboLegs: []*qotcommon.ComboLeg{}, Qty: 1.0}, true, "ComboLegs is empty"},
		{"zero qty", &PlaceComboOrderRequest{Header: header, ComboLegs: []*qotcommon.ComboLeg{{}}, Qty: 0}, true, "Qty must be positive"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := PlaceComboOrder(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2227)
		})
	}
}

// TestReconfirmOrder_API exercises ReconfirmOrder with MockServer.
func TestReconfirmOrder_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(2209, func(reqBody []byte) (proto.Message, error) {
		var req trdreconfirmorder.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		_ = req.C2S.GetHeader()
		return &trdreconfirmorder.Response{RetType: okRet(), S2C: &trdreconfirmorder.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	packetID := &common.PacketID{ConnID: u64Ptr(1), SerialNo: u32Ptr(1)}

	for _, tc := range []struct {
		name      string
		req       *ReconfirmOrderRequest
		wantErr   bool
		errSubStr string
	}{
		{"nil req", nil, true, "request is nil"},
		{"zero orderid", &ReconfirmOrderRequest{PacketID: packetID, AccID: 111111, TrdMarket: constant.TrdMarket_HK, OrderID: 0}, true, "invalid order ID"},
		{"zero accid", &ReconfirmOrderRequest{PacketID: packetID, AccID: 0, TrdMarket: constant.TrdMarket_HK, OrderID: 12345}, true, "invalid account ID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			_, err := ReconfirmOrder(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tc.errSubStr != "" && !strContains(err.Error(), tc.errSubStr) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errSubStr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			srv.AssertProtoID(t, 2209)
		})
	}
}

func strContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
