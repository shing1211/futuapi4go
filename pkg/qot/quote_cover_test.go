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
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

// sp returns a pointer to a string.
func sp(s string) *string { return &s }

// ip returns a pointer to an int32.
func ip(i int32) *int32 { return &i }

func hkSec(code string) *qotcommon.Security {
	m := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	return &qotcommon.Security{Market: &m, Code: &code}
}

// =============================================================================
// BasicQot (quote.go:112)
// =============================================================================

func TestBasicQotStruct(t *testing.T) {
	sec := hkSec("00700")
	bq := &BasicQot{
		Security:        sec,
		Name:            "Tencent",
		IsSuspended:     false,
		UpdateTime:      "2024-01-15 14:30:00",
		HighPrice:       352.0,
		OpenPrice:       348.0,
		LowPrice:        347.0,
		CurPrice:        350.5,
		LastClosePrice:  349.0,
		Volume:          12345678,
		Turnover:        4321098765.0,
		TurnoverRate:    0.025,
		Amplitude:       0.030,
		ListTime:        "2004-06-16",
		PriceSpread:     0.1,
		DarkStatus:      0,
		ListTimestamp:   1087248000.0,
		UpdateTimestamp: 1705312200.0,
		SecStatus:       0,
	}

	if bq.Security.GetCode() != "00700" {
		t.Errorf("Security.Code=%s want 00700", bq.Security.GetCode())
	}
	if bq.Name != "Tencent" {
		t.Errorf("Name=%s want Tencent", bq.Name)
	}
	if bq.CurPrice != 350.5 {
		t.Errorf("CurPrice=%f want 350.5", bq.CurPrice)
	}
	if bq.Volume != 12345678 {
		t.Errorf("Volume=%d want 12345678", bq.Volume)
	}
	if bq.TurnoverRate != 0.025 {
		t.Errorf("TurnoverRate=%f want 0.025", bq.TurnoverRate)
	}
	if bq.IsSuspended {
		t.Error("IsSuspended should be false")
	}
	if bq.HighPrice != 352.0 {
		t.Errorf("HighPrice=%f want 352.0", bq.HighPrice)
	}
	if bq.LowPrice != 347.0 {
		t.Errorf("LowPrice=%f want 347.0", bq.LowPrice)
	}
}

func TestBasicQotStructZero(t *testing.T) {
	bq := &BasicQot{}
	if bq.CurPrice != 0 {
		t.Errorf("CurPrice=%f want 0", bq.CurPrice)
	}
	if bq.Volume != 0 {
		t.Errorf("Volume=%d want 0", bq.Volume)
	}
	if bq.Name != "" {
		t.Errorf("Name=%q want empty", bq.Name)
	}
}

// =============================================================================
// KLine (quote.go:202)
// =============================================================================

func TestKLineStruct(t *testing.T) {
	kl := &KLine{
		Time:           "2024-01-15",
		IsBlank:        false,
		HighPrice:      352.0,
		OpenPrice:      348.0,
		LowPrice:       347.0,
		ClosePrice:     350.5,
		LastClosePrice: 349.0,
		Volume:         12345678,
		Turnover:       4321098765.0,
		TurnoverRate:   0.025,
		Pe:             25.5,
		ChangeRate:     0.004,
		Timestamp:       1705312200.0,
	}

	if kl.ClosePrice != 350.5 {
		t.Errorf("ClosePrice=%f want 350.5", kl.ClosePrice)
	}
	if kl.OpenPrice != 348.0 {
		t.Errorf("OpenPrice=%f want 348.0", kl.OpenPrice)
	}
	if kl.HighPrice != 352.0 {
		t.Errorf("HighPrice=%f want 352.0", kl.HighPrice)
	}
	if kl.LowPrice != 347.0 {
		t.Errorf("LowPrice=%f want 347.0", kl.LowPrice)
	}
	if kl.Volume != 12345678 {
		t.Errorf("Volume=%d want 12345678", kl.Volume)
	}
	if kl.Turnover != 4321098765.0 {
		t.Errorf("Turnover=%f want 4321098765.0", kl.Turnover)
	}
	if kl.IsBlank {
		t.Error("IsBlank should be false")
	}
	if kl.Pe != 25.5 {
		t.Errorf("Pe=%f want 25.5", kl.Pe)
	}
}

func TestKLineStructZero(t *testing.T) {
	kl := &KLine{}
	if kl.ClosePrice != 0 {
		t.Errorf("ClosePrice=%f want 0", kl.ClosePrice)
	}
	if kl.Volume != 0 {
		t.Errorf("Volume=%d want 0", kl.Volume)
	}
	if kl.Time != "" {
		t.Errorf("Time=%q want empty", kl.Time)
	}
}

// =============================================================================
// GetKLRequest (quote.go:447)
// =============================================================================

func TestGetKLRequestStruct(t *testing.T) {
	sec := hkSec("00700")

	tests := []struct {
		name      string
		req       GetKLRequest
		wantNum   int32
		wantRehab int32
		wantKL   int32
	}{
		{
			name:      "day kline",
			req:       GetKLRequest{Security: sec, RehabType: 0, KLType: 0, ReqNum: 10},
			wantNum:   10,
			wantRehab: 0,
			wantKL:   0,
		},
		{
			name:      "1min kline",
			req:       GetKLRequest{Security: sec, RehabType: 1, KLType: 1, ReqNum: 100},
			wantNum:   100,
			wantRehab: 1,
			wantKL:   1,
		},
		{
			name:      "week kline",
			req:       GetKLRequest{Security: sec, RehabType: 1, KLType: 3, ReqNum: 50},
			wantNum:   50,
			wantRehab: 1,
			wantKL:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.req.Security == nil {
				t.Fatal("Security is nil")
			}
			if tt.req.Security.GetCode() != "00700" {
				t.Errorf("Code=%s want 00700", tt.req.Security.GetCode())
			}
			if tt.req.ReqNum != tt.wantNum {
				t.Errorf("ReqNum=%d want %d", tt.req.ReqNum, tt.wantNum)
			}
			if tt.req.RehabType != tt.wantRehab {
				t.Errorf("RehabType=%d want %d", tt.req.RehabType, tt.wantRehab)
			}
			if tt.req.KLType != tt.wantKL {
				t.Errorf("KLType=%d want %d", tt.req.KLType, tt.wantKL)
			}
		})
	}
}

func TestGetKLRequestNilSecurity(t *testing.T) {
	req := &GetKLRequest{}
	if req.Security != nil {
		t.Error("Security should be nil")
	}
	if req.ReqNum != 0 {
		t.Errorf("ReqNum=%d want 0", req.ReqNum)
	}
}

// =============================================================================
// GetKLResponse (quote.go:455)
// =============================================================================

func TestGetKLResponseStruct(t *testing.T) {
	sec := hkSec("00700")
	kl1 := &KLine{Time: "2024-01-15", ClosePrice: 350.5, Volume: 1000, OpenPrice: 348.0, HighPrice: 352.0, LowPrice: 347.0}
	kl2 := &KLine{Time: "2024-01-16", ClosePrice: 351.0, Volume: 1200, OpenPrice: 350.5, HighPrice: 353.0, LowPrice: 350.0}

	rsp := &GetKLResponse{
		Security: sec,
		Name:     "Tencent",
		KLList:   []*KLine{kl1, kl2},
	}

	if rsp.Security == nil {
		t.Fatal("Security should not be nil")
	}
	if rsp.Security.GetCode() != "00700" {
		t.Errorf("Security.Code=%s want 00700", rsp.Security.GetCode())
	}
	if rsp.Name != "Tencent" {
		t.Errorf("Name=%s want Tencent", rsp.Name)
	}
	if len(rsp.KLList) != 2 {
		t.Fatalf("KLList length=%d want 2", len(rsp.KLList))
	}
	if rsp.KLList[0].ClosePrice != 350.5 {
		t.Errorf("KLList[0].ClosePrice=%f want 350.5", rsp.KLList[0].ClosePrice)
	}
	if rsp.KLList[1].Volume != 1200 {
		t.Errorf("KLList[1].Volume=%d want 1200", rsp.KLList[1].Volume)
	}
}

func TestGetKLResponseEmpty(t *testing.T) {
	rsp := &GetKLResponse{KLList: []*KLine{}}
	if len(rsp.KLList) != 0 {
		t.Errorf("KLList length=%d want 0", len(rsp.KLList))
	}
}
