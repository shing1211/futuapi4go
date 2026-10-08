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
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsecuritysnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotregqotpush"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
	testutil "github.com/shing1211/futuapi4go/test/util"
	"google.golang.org/protobuf/proto"
)

func okRet() *int32 {
	v := int32(common.RetType_RetType_Succeed)
	return &v
}

func strPtr(s string) *string { return &s }

func strContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// =============================================================================
// GetOrderBook (3012)
// =============================================================================

func TestGetOrderBook_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3012, func(reqBody []byte) (proto.Message, error) {
		var req qotgetorderbook.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.C2S == nil || req.C2S.Security == nil {
			t.Fatal("expected C2S.Security")
		}
		price1, price2 := 350.00, 351.00
		vol1, vol2 := int64(5000), int64(5000)
		oc1, oc2 := int32(3), int32(3)
		oid1 := int64(12345)
		detailVol := int64(500)
		svrBid, svrAsk := "10:00:00", "10:00:00"
		tsBid, tsAsk := 1775635200.0, 1775635200.0
		name := "Tencent"

		return &qotgetorderbook.Response{
			RetType: okRet(),
			S2C: &qotgetorderbook.S2C{
				Security:             req.C2S.Security,
				Name:                 &name,
				OrderBookAskList: []*qotcommon.OrderBook{
					{Price: &price1, Volume: &vol1, OrederCount: &oc1,
						DetailList: []*qotcommon.OrderBookDetail{
							{OrderID: &oid1, Volume: &detailVol},
						}},
				},
				OrderBookBidList: []*qotcommon.OrderBook{
					{Price: &price2, Volume: &vol2, OrederCount: &oc2},
				},
				SvrRecvTimeBid:          &svrBid,
				SvrRecvTimeBidTimestamp: &tsBid,
				SvrRecvTimeAsk:          &svrAsk,
				SvrRecvTimeAskTimestamp: &tsAsk,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetOrderBookRequest
		wantErr   bool
		errSubStr string
		wantAsks  int
		wantBids  int
	}{
		{"basic", &GetOrderBookRequest{Security: sec, Num: 10}, false, "", 1, 1},
		{"nil req", nil, true, "request is nil", 0, 0},
		{"nil security", &GetOrderBookRequest{Security: nil, Num: 10}, true, "Security is nil", 0, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetOrderBook(context.Background(), cli, tc.req)
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
			if tc.wantAsks > 0 && len(rsp.OrderBookAskList) != tc.wantAsks {
				t.Errorf("asks: want %d, got %d", tc.wantAsks, len(rsp.OrderBookAskList))
			}
			if tc.wantBids > 0 && len(rsp.OrderBookBidList) != tc.wantBids {
				t.Errorf("bids: want %d, got %d", tc.wantBids, len(rsp.OrderBookBidList))
			}
			srv.AssertProtoID(t, 3012)
		})
	}
}

// =============================================================================
// GetTicker (3010)
// =============================================================================

func TestGetTicker_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3010, func(reqBody []byte) (proto.Message, error) {
		name := "Tencent"
		ts := "10:00:00"
		seq := int64(123456)
		dir := int32(1)
		price := 350.00
		vol := int64(1000)
		turnover := 350000.00
		recvTime := 1775635200.0
		typ, typeSign, pushType := int32(0), int32(1), int32(1)

		return &qotgetticker.Response{
			RetType: okRet(),
			S2C: &qotgetticker.S2C{
				Security:   nil,
				Name:       &name,
				TickerList: []*qotcommon.Ticker{
					{Time: &ts, Sequence: &seq, Dir: &dir, Price: &price, Volume: &vol,
						Turnover: &turnover, RecvTime: &recvTime, Type: &typ, TypeSign: &typeSign,
						Timestamp: &recvTime, PushDataType: &pushType},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetTickerRequest
		wantErr   bool
		errSubStr string
		wantTicks int
	}{
		{"basic", &GetTickerRequest{Security: sec, Num: 100}, false, "", 1},
		{"nil req", nil, true, "request is nil", 0},
		{"nil security", &GetTickerRequest{Security: nil, Num: 100}, true, "Security is nil", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetTicker(context.Background(), cli, tc.req)
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
			if tc.wantTicks > 0 && len(rsp.TickerList) != tc.wantTicks {
				t.Errorf("tickers: want %d, got %d", tc.wantTicks, len(rsp.TickerList))
			}
			srv.AssertProtoID(t, 3010)
		})
	}
}

// =============================================================================
// GetRT (3008)
// =============================================================================

func TestGetRT_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3008, func(reqBody []byte) (proto.Message, error) {
		name := "Tencent"
		ts := "10:00:00"
		minute := int32(600)
		isBlank := false
		price, lastClose, avgPrice := 350.00, 349.00, 349.80
		vol := int64(12345678)
		turnover := 4321098765.00
		tsF := 1744116000.0
		return &qotgetrt.Response{
			RetType: okRet(),
			S2C: &qotgetrt.S2C{
				Security: nil,
				Name:     &name,
				RtList: []*qotcommon.TimeShare{
					{Time: &ts, Minute: &minute, IsBlank: &isBlank, Price: &price,
						LastClosePrice: &lastClose, AvgPrice: &avgPrice, Volume: &vol,
						Turnover: &turnover, Timestamp: &tsF},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name     string
		req      *GetRTRequest
		wantErr  bool
		wantRTs  int
		wantName string
	}{
		{"basic", &GetRTRequest{Security: sec}, false, 1, "Tencent"},
		{"nil req", nil, true, 0, ""},
		{"nil security", &GetRTRequest{Security: nil}, true, 0, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetRT(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.wantName != "" && rsp.Name != tc.wantName {
				t.Errorf("Name: want %q, got %q", tc.wantName, rsp.Name)
			}
			if tc.wantRTs > 0 && len(rsp.RTList) != tc.wantRTs {
				t.Errorf("RTs: want %d, got %d", tc.wantRTs, len(rsp.RTList))
			}
			srv.AssertProtoID(t, 3008)
		})
	}
}

// =============================================================================
// GetBroker (3014)
// =============================================================================

func TestGetBroker_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3014, func(reqBody []byte) (proto.Message, error) {
		name := "Tencent"
		id1, nm1 := int64(1), "Citi"
		pos1, vol1, oid1 := int32(1), int64(5000), int64(67890)
		id2, nm2 := int64(2), "HSBC"
		pos2, vol2, oid2 := int32(1), int64(6000), int64(67891)
		return &qotgetbroker.Response{
			RetType: okRet(),
			S2C: &qotgetbroker.S2C{
				Security:     nil,
				Name:         &name,
				BrokerAskList: []*qotcommon.Broker{
					{Id: &id1, Name: &nm1, Pos: &pos1, Volume: &vol1, OrderID: &oid1},
				},
				BrokerBidList: []*qotcommon.Broker{
					{Id: &id2, Name: &nm2, Pos: &pos2, Volume: &vol2, OrderID: &oid2},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name     string
		req      *GetBrokerRequest
		wantErr  bool
		wantAsks int
		wantBids int
	}{
		{"basic", &GetBrokerRequest{Security: sec}, false, 1, 1},
		{"nil req", nil, true, 0, 0},
		{"nil security", &GetBrokerRequest{Security: nil}, true, 0, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetBroker(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.wantAsks > 0 && len(rsp.AskBrokerList) != tc.wantAsks {
				t.Errorf("asks: want %d, got %d", tc.wantAsks, len(rsp.AskBrokerList))
			}
			if tc.wantBids > 0 && len(rsp.BidBrokerList) != tc.wantBids {
				t.Errorf("bids: want %d, got %d", tc.wantBids, len(rsp.BidBrokerList))
			}
			srv.AssertProtoID(t, 3014)
		})
	}
}

// =============================================================================
// Subscribe (3001)
// =============================================================================

func TestSubscribe_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3001, func(reqBody []byte) (proto.Message, error) {
		return &qotsub.Response{RetType: okRet(), S2C: &qotsub.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *SubscribeRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &SubscribeRequest{SecurityList: []*qotcommon.Security{sec}, SubTypeList: []SubType{SubType_Basic, SubType_KL}, IsSubOrUnSub: true}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"empty seclist", &SubscribeRequest{SecurityList: []*qotcommon.Security{}, SubTypeList: []SubType{SubType_Basic}, IsSubOrUnSub: true}, true, "security list is empty"},
		{"empty subtypelist", &SubscribeRequest{SecurityList: []*qotcommon.Security{sec}, SubTypeList: []SubType{}, IsSubOrUnSub: true}, true, "subtype list is empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			err := Subscribe(context.Background(), cli, tc.req)
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
			srv.AssertProtoID(t, 3001)
		})
	}
}

// =============================================================================
// RegQotPush (3002)
// =============================================================================

func TestRegQotPush_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3002, func(reqBody []byte) (proto.Message, error) {
		return &qotregqotpush.Response{RetType: okRet(), S2C: &qotregqotpush.S2C{}}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *RegQotPushRequest
		wantErr   bool
		errSubStr string
	}{
		{"basic", &RegQotPushRequest{SecurityList: []*qotcommon.Security{sec}, SubTypeList: []int32{1, 2}, RehabTypeList: []int32{0}, IsRegOrUnReg: true, IsFirstPush: true}, false, ""},
		{"nil req", nil, true, "request is nil"},
		{"empty seclist", &RegQotPushRequest{SecurityList: []*qotcommon.Security{}, SubTypeList: []int32{1}, IsRegOrUnReg: true, IsFirstPush: true}, true, "security list is empty"},
		{"empty subtypelist", &RegQotPushRequest{SecurityList: []*qotcommon.Security{sec}, SubTypeList: []int32{}, IsRegOrUnReg: true, IsFirstPush: true}, true, "subtype list is empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			err := RegQotPush(context.Background(), cli, tc.req)
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
			srv.AssertProtoID(t, 3002)
		})
	}
}

// =============================================================================
// GetSubInfo (3003)
// =============================================================================

func TestGetSubInfo_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3003, func(reqBody []byte) (proto.Message, error) {
		totalUsed, remain := int32(10), int32(90)
		subType, usedQuota := int32(1), int32(5)
		isOwnConn := true
		return &qotgetsubinfo.Response{
			RetType: okRet(),
			S2C: &qotgetsubinfo.S2C{
				ConnSubInfoList: []*qotcommon.ConnSubInfo{
					{SubInfoList: []*qotcommon.SubInfo{{SubType: &subType}}, UsedQuota: &usedQuota, IsOwnConnData: &isOwnConn},
				},
				TotalUsedQuota: &totalUsed,
				RemainQuota:    &remain,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()

	rsp, err := GetSubInfo(context.Background(), cli)
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if rsp.TotalUsedQuota != 10 {
		t.Errorf("TotalUsedQuota: want 10, got %d", rsp.TotalUsedQuota)
	}
	if rsp.RemainQuota != 90 {
		t.Errorf("RemainQuota: want 90, got %d", rsp.RemainQuota)
	}
	if len(rsp.ConnSubInfoList) != 1 {
		t.Fatalf("ConnSubInfoList: want 1, got %d", len(rsp.ConnSubInfoList))
	}
	srv.AssertProtoID(t, 3003)
}

// =============================================================================
// GetSecuritySnapshot (3203)
// =============================================================================

func TestGetSecuritySnapshot_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3203, func(reqBody []byte) (proto.Message, error) {
		var req qotgetsecuritysnapshot.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			return nil, err
		}
		sec := req.C2S.SecurityList[0]
		name, secType := "Tencent", int32(1)
		lotSize, listTime := int32(100), "2004-06-16"
		isSuspend, priceSpread := false, 0.1
		updateTime, highPrice := "2024-01-15 14:30:00", 352.0
		openPrice, lowPrice := 348.0, 347.0
		lastClose, curPrice := 349.0, 350.5
		vol, turnover := int64(12345678), 4321098765.0
		turnoverRate := 0.025
		return &qotgetsecuritysnapshot.Response{
			RetType: okRet(),
			S2C: &qotgetsecuritysnapshot.S2C{
				SnapshotList: []*qotgetsecuritysnapshot.Snapshot{
					{Basic: &qotgetsecuritysnapshot.SnapshotBasicData{
						Security: sec, Name: &name, Type: &secType, IsSuspend: &isSuspend,
						ListTime: &listTime, LotSize: &lotSize, PriceSpread: &priceSpread,
						UpdateTime: &updateTime, HighPrice: &highPrice, OpenPrice: &openPrice,
						LowPrice: &lowPrice, LastClosePrice: &lastClose, CurPrice: &curPrice,
						Volume: &vol, Turnover: &turnover, TurnoverRate: &turnoverRate,
					}},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetSecuritySnapshotRequest
		wantErr   bool
		errSubStr string
		wantSnaps int
	}{
		{"basic", &GetSecuritySnapshotRequest{SecurityList: []*qotcommon.Security{sec}}, false, "", 1},
		{"nil req", nil, true, "request is nil", 0},
		{"empty seclist", &GetSecuritySnapshotRequest{SecurityList: []*qotcommon.Security{}}, true, "security list is empty", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetSecuritySnapshot(context.Background(), cli, tc.req)
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
			if tc.wantSnaps > 0 && len(rsp.SnapshotList) != tc.wantSnaps {
				t.Errorf("snaps: want %d, got %d", tc.wantSnaps, len(rsp.SnapshotList))
			}
			srv.AssertProtoID(t, 3203)
		})
	}
}

// =============================================================================
// GetStaticInfo (3202)
// =============================================================================

func TestGetStaticInfo_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3202, func(reqBody []byte) (proto.Message, error) {
		var req qotgetstaticinfo.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			return nil, err
		}
		sec := req.C2S.SecurityList[0]
		name, id := "Tencent", int64(1)
		secType, lotSize := int32(1), int32(100)
		listTime := "2004-06-16"
		return &qotgetstaticinfo.Response{
			RetType: okRet(),
			S2C: &qotgetstaticinfo.S2C{
				StaticInfoList: []*qotcommon.SecurityStaticInfo{
					{Basic: &qotcommon.SecurityStaticBasic{
						Security: sec, Id: &id, Name: &name, SecType: &secType,
						LotSize: &lotSize, ListTime: &listTime,
					}},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetStaticInfoRequest
		wantErr   bool
		errSubStr string
		wantInfos int
	}{
		{"by seclist", &GetStaticInfoRequest{Market: 1, SecType: 1, SecurityList: []*qotcommon.Security{sec}}, false, "", 1},
		{"nil req", nil, true, "request is nil", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetStaticInfo(context.Background(), cli, tc.req)
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
			if tc.wantInfos > 0 && len(rsp.StaticInfoList) != tc.wantInfos {
				t.Errorf("infos: want %d, got %d", tc.wantInfos, len(rsp.StaticInfoList))
			}
			srv.AssertProtoID(t, 3202)
		})
	}
}

// =============================================================================
// GetKL (3006) — qotgetkl, qotcommon.KLine
// =============================================================================

func TestGetKL_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3006, func(reqBody []byte) (proto.Message, error) {
		var req qotgetkl.Request
		if err := proto.Unmarshal(reqBody, &req); err != nil {
			return nil, err
		}
		name := "Tencent"
		time1, time2 := "2024-01-15", "2024-01-16"
		high, open, low, close := 352.0, 348.0, 347.0, 350.5
		vol1, vol2 := int64(12345678), int64(13245678)
		turn1, turn2 := 4321098765.0, 4600000000.0
		ts1, ts2 := 1705312200.0, 1705398600.0
		lastClose := 349.0
		pe, changeRate := 25.5, 0.004
		isBlank := false
		return &qotgetkl.Response{
			RetType: okRet(),
			S2C: &qotgetkl.S2C{
				Security: req.C2S.Security,
				Name:     &name,
				KlList: []*qotcommon.KLine{
					{Time: &time1, HighPrice: &high, OpenPrice: &open, LowPrice: &low, ClosePrice: &close,
						LastClosePrice: &lastClose, Volume: &vol1, Turnover: &turn1, Timestamp: &ts1,
						IsBlank: &isBlank, ChangeRate: &changeRate, Pe: &pe},
					{Time: &time2, HighPrice: &high, OpenPrice: &open, LowPrice: &low, ClosePrice: &close,
						LastClosePrice: &lastClose, Volume: &vol2, Turnover: &turn2, Timestamp: &ts2,
						IsBlank: &isBlank, ChangeRate: &changeRate, Pe: &pe},
				},
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name      string
		req       *GetKLRequest
		wantErr   bool
		errSubStr string
		wantKLs   int
	}{
		{"basic", &GetKLRequest{Security: sec, RehabType: 0, KLType: 0, ReqNum: 10}, false, "", 2},
		{"nil req", nil, true, "request is nil", 0},
		{"nil security", &GetKLRequest{Security: nil, RehabType: 0, KLType: 0, ReqNum: 10}, true, "Security is nil", 0},
		{"zero reqnum", &GetKLRequest{Security: sec, RehabType: 0, KLType: 0, ReqNum: 0}, true, "ReqNum must be positive", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetKL(context.Background(), cli, tc.req)
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
			if tc.wantKLs > 0 && len(rsp.KLList) != tc.wantKLs {
				t.Errorf("KLs: want %d, got %d", tc.wantKLs, len(rsp.KLList))
			}
			srv.AssertProtoID(t, 3006)
		})
	}
}

// =============================================================================
// CapitalFlow (3009) — qotgetcapitalflow, CapitalFlowItem
// =============================================================================

func TestCapitalFlow_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3211, func(reqBody []byte) (proto.Message, error) {
		inFlow := 1234567.0
		time1, time2 := "10:00:00", "10:05:00"
		ts1, ts2 := 1705312200.0, 1705312500.0
		mainInFlow := 800000.0
		lastTime := "15:00:00"
		lastTS := 1705312200.0
		return &qotgetcapitalflow.Response{
			RetType: okRet(),
			S2C: &qotgetcapitalflow.S2C{
				FlowItemList: []*qotgetcapitalflow.CapitalFlowItem{
					{InFlow: &inFlow, Time: &time1, Timestamp: &ts1, MainInFlow: &mainInFlow},
					{InFlow: &inFlow, Time: &time2, Timestamp: &ts2, MainInFlow: &mainInFlow},
				},
				LastValidTime:      &lastTime,
				LastValidTimestamp: &lastTS,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name     string
		req      *GetCapitalFlowRequest
		wantErr  bool
		wantFlow int
	}{
		{"basic", &GetCapitalFlowRequest{Security: sec}, false, 2},
		{"nil req", nil, true, 0},
		{"nil security", &GetCapitalFlowRequest{Security: nil}, true, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetCapitalFlow(context.Background(), cli, tc.req)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.wantFlow > 0 && len(rsp.FlowItemList) != tc.wantFlow {
				t.Errorf("flows: want %d, got %d", tc.wantFlow, len(rsp.FlowItemList))
			}
			srv.AssertProtoID(t, 3211)
		})
	}
}

// =============================================================================
// CapitalDistribution (3011) — qotgetcapitaldistribution
// =============================================================================

func TestCapitalDistribution_API(t *testing.T) {
	srv := testutil.NewMockServer(t)
	srv.Start()
	defer srv.Stop()

	srv.RegisterHandler(3212, func(reqBody []byte) (proto.Message, error) {
		capInBig, capOutBig := 5000000.0, 3000000.0
		capInMid, capOutMid := 2000000.0, 1500000.0
		capInSmall, capOutSmall := 1000000.0, 800000.0
		updateTime := "15:00:00"
		updateTS := 1705312200.0
		return &qotgetcapitaldistribution.Response{
			RetType: okRet(),
			S2C: &qotgetcapitaldistribution.S2C{
				CapitalInBig:    &capInBig,
				CapitalOutBig:   &capOutBig,
				CapitalInMid:    &capInMid,
				CapitalOutMid:   &capOutMid,
				CapitalInSmall:  &capInSmall,
				CapitalOutSmall: &capOutSmall,
				UpdateTime:      &updateTime,
				UpdateTimestamp: &updateTS,
			},
		}, nil
	})

	cli, cleanup := testutil.NewTestClient(t, srv)
	defer cleanup()
	hkMkt := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	sec := &qotcommon.Security{Market: &hkMkt, Code: strPtr("00700")}

	for _, tc := range []struct {
		name        string
		security    *qotcommon.Security
		wantErr     bool
		wantCapIn   float64
	}{
		{"basic", sec, false, 5000000.0},
		{"nil security", nil, true, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv.ClearRequests()
			rsp, err := GetCapitalDistribution(context.Background(), cli, tc.security)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected: %v", err)
			}
			if tc.wantCapIn > 0 && rsp.CapitalDistribution.CapitalInBig != tc.wantCapIn {
				t.Errorf("CapitalInBig: want %.0f, got %.0f", tc.wantCapIn, rsp.CapitalDistribution.CapitalInBig)
			}
			srv.AssertProtoID(t, 3212)
		})
	}
}
