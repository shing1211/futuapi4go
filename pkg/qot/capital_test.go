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

// =============================================================================
// GetCapitalFlowRequest
// =============================================================================

func TestGetCapitalFlowRequestFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	sec := &qotcommon.Security{Market: &hkMarket, Code: &code}

	tests := []struct {
		name      string
		req       *GetCapitalFlowRequest
		wantSec   string
		wantPType int32
	}{
		{
			name:      "basic request",
			req:       &GetCapitalFlowRequest{Security: sec},
			wantSec:   "00700",
			wantPType: 0,
		},
		{
			name: "with period type",
			req: &GetCapitalFlowRequest{
				Security:   sec,
				PeriodType: 1,
			},
			wantSec:   "00700",
			wantPType: 1,
		},
		{
			name: "with time range",
			req: &GetCapitalFlowRequest{
				Security:  sec,
				BeginTime: "2024-01-01 09:00:00",
				EndTime:   "2024-01-31 18:00:00",
			},
			wantSec:   "00700",
			wantPType: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.req.Security == nil {
				t.Fatal("Security should not be nil")
			}
			if tt.req.Security.GetCode() != tt.wantSec {
				t.Errorf("Security.Code = %q, want %q", tt.req.Security.GetCode(), tt.wantSec)
			}
			if tt.req.PeriodType != tt.wantPType {
				t.Errorf("PeriodType = %d, want %d", tt.req.PeriodType, tt.wantPType)
			}
		})
	}
}

func TestGetCapitalFlowRequestNilSecurity(t *testing.T) {
	req := &GetCapitalFlowRequest{}
	if req.Security != nil {
		t.Error("Security should be nil by default")
	}
	if req.PeriodType != 0 {
		t.Errorf("PeriodType = %d, want 0", req.PeriodType)
	}
}

// =============================================================================
// CapitalFlowItem
// =============================================================================

func TestCapitalFlowItemFields_Qot(t *testing.T) {
	item := &CapitalFlowItem{
		InFlow:      1000000.5,
		Time:        "2024-01-15 14:30:00",
		Timestamp:   1705312200.0,
		MainInFlow:  500000.0,
		SuperInFlow: 200000.0,
		BigInFlow:   150000.0,
		MidInFlow:   100000.0,
		SmlInFlow:   50000.0,
	}

	if item.InFlow != 1000000.5 {
		t.Errorf("InFlow = %f, want 1000000.5", item.InFlow)
	}
	if item.Time != "2024-01-15 14:30:00" {
		t.Errorf("Time = %q, want %q", item.Time, "2024-01-15 14:30:00")
	}
	if item.Timestamp != 1705312200.0 {
		t.Errorf("Timestamp = %f, want 1705312200.0", item.Timestamp)
	}
	if item.MainInFlow != 500000.0 {
		t.Errorf("MainInFlow = %f, want 500000.0", item.MainInFlow)
	}
	if item.SmlInFlow != 50000.0 {
		t.Errorf("SmlInFlow = %f, want 50000.0", item.SmlInFlow)
	}
}

func TestCapitalFlowItemZeroValues(t *testing.T) {
	item := &CapitalFlowItem{}
	if item.InFlow != 0 {
		t.Errorf("InFlow = %f, want 0", item.InFlow)
	}
	if item.Time != "" {
		t.Errorf("Time = %q, want empty", item.Time)
	}
}

// =============================================================================
// GetCapitalFlowResponse
// =============================================================================

func TestGetCapitalFlowResponseFields(t *testing.T) {
	item1 := &CapitalFlowItem{InFlow: 1000000.0, Time: "2024-01-15"}
	item2 := &CapitalFlowItem{InFlow: 1200000.0, Time: "2024-01-16"}

	rsp := &GetCapitalFlowResponse{
		FlowItemList:       []*CapitalFlowItem{item1, item2},
		LastValidTime:      "2024-01-16 15:30:00",
		LastValidTimestamp: 1705400000.0,
	}

	if len(rsp.FlowItemList) != 2 {
		t.Fatalf("FlowItemList length = %d, want 2", len(rsp.FlowItemList))
	}
	if rsp.FlowItemList[0].InFlow != 1000000.0 {
		t.Errorf("FlowItemList[0].InFlow = %f, want 1000000.0", rsp.FlowItemList[0].InFlow)
	}
	if rsp.LastValidTime != "2024-01-16 15:30:00" {
		t.Errorf("LastValidTime = %q, want %q", rsp.LastValidTime, "2024-01-16 15:30:00")
	}
	if rsp.LastValidTimestamp != 1705400000.0 {
		t.Errorf("LastValidTimestamp = %f, want 1705400000.0", rsp.LastValidTimestamp)
	}
}

func TestGetCapitalFlowResponseEmpty(t *testing.T) {
	rsp := &GetCapitalFlowResponse{
		FlowItemList: []*CapitalFlowItem{},
	}
	if len(rsp.FlowItemList) != 0 {
		t.Errorf("FlowItemList length = %d, want 0", len(rsp.FlowItemList))
	}
	if rsp.LastValidTime != "" {
		t.Errorf("LastValidTime = %q, want empty", rsp.LastValidTime)
	}
}

// =============================================================================
// CapitalDistribution
// =============================================================================

func TestCapitalDistributionFields_Qot(t *testing.T) {
	cd := &CapitalDistribution{
		CapitalInSuper:  5000000.0,
		CapitalInBig:   3000000.0,
		CapitalInMid:   2000000.0,
		CapitalInSmall: 1000000.0,
		CapitalOutSuper: 4000000.0,
		CapitalOutBig:  2000000.0,
		CapitalOutMid:  1000000.0,
		CapitalOutSmall: 500000.0,
		UpdateTime:     "2024-01-15 16:00:00",
		UpdateTimestamp: 1705310400.0,
	}

	if cd.CapitalInSuper != 5000000.0 {
		t.Errorf("CapitalInSuper = %f, want 5000000.0", cd.CapitalInSuper)
	}
	if cd.CapitalOutSmall != 500000.0 {
		t.Errorf("CapitalOutSmall = %f, want 500000.0", cd.CapitalOutSmall)
	}
	if cd.UpdateTime != "2024-01-15 16:00:00" {
		t.Errorf("UpdateTime = %q, want %q", cd.UpdateTime, "2024-01-15 16:00:00")
	}
}

func TestCapitalDistributionZeroValues(t *testing.T) {
	cd := &CapitalDistribution{}
	if cd.CapitalInSuper != 0 {
		t.Errorf("CapitalInSuper = %f, want 0", cd.CapitalInSuper)
	}
	if cd.CapitalOutBig != 0 {
		t.Errorf("CapitalOutBig = %f, want 0", cd.CapitalOutBig)
	}
	if cd.UpdateTime != "" {
		t.Errorf("UpdateTime = %q, want empty", cd.UpdateTime)
	}
}

// =============================================================================
// GetCapitalDistributionResponse
// =============================================================================

func TestGetCapitalDistributionResponseFields(t *testing.T) {
	cd := &CapitalDistribution{
		CapitalInSuper:  5000000.0,
		CapitalInBig:   3000000.0,
		CapitalInMid:   2000000.0,
		CapitalInSmall: 1000000.0,
		UpdateTime:     "2024-01-15",
	}
	rsp := &GetCapitalDistributionResponse{CapitalDistribution: cd}

	if rsp.CapitalDistribution == nil {
		t.Fatal("CapitalDistribution should not be nil")
	}
	if rsp.CapitalDistribution.CapitalInSuper != 5000000.0 {
		t.Errorf("CapitalInSuper = %f, want 5000000.0", rsp.CapitalDistribution.CapitalInSuper)
	}
}

func TestGetCapitalDistributionResponseNil(t *testing.T) {
	rsp := &GetCapitalDistributionResponse{}
	if rsp.CapitalDistribution != nil {
		t.Error("CapitalDistribution should be nil")
	}
}
