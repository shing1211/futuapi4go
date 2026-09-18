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
	"fmt"
	"testing"
)

func TestFlowSummaryInfoFields(t *testing.T) {
	info := &FlowSummaryInfo{
		CashFlowID:        12345,
		ClearingDate:      "2026-04-08",
		SettlementDate:    "2026-04-09",
		Currency:          1,
		CashFlowType:      "BUY",
		CashFlowDirection: 1,
		CashFlowAmount:    10000.50,
		CashFlowRemark:    "Test remark",
	}

	if info.CashFlowID != 12345 {
		t.Errorf("expected CashFlowID 12345, got %d", info.CashFlowID)
	}
	if info.ClearingDate != "2026-04-08" {
		t.Errorf("expected ClearingDate 2026-04-08, got %s", info.ClearingDate)
	}
	if info.SettlementDate != "2026-04-09" {
		t.Errorf("expected SettlementDate 2026-04-09, got %s", info.SettlementDate)
	}
	if info.Currency != 1 {
		t.Errorf("expected Currency 1, got %d", info.Currency)
	}
	if info.CashFlowType != "BUY" {
		t.Errorf("expected CashFlowType BUY, got %s", info.CashFlowType)
	}
	if info.CashFlowDirection != 1 {
		t.Errorf("expected CashFlowDirection 1, got %d", info.CashFlowDirection)
	}
	if info.CashFlowAmount != 10000.50 {
		t.Errorf("expected CashFlowAmount 10000.50, got %f", info.CashFlowAmount)
	}
	if info.CashFlowRemark != "Test remark" {
		t.Errorf("expected CashFlowRemark Test remark, got %s", info.CashFlowRemark)
	}
}

func TestFlowSummaryInfoFromProtoNil(t *testing.T) {
	result := flowSummaryInfoFromProto(nil)
	if result != nil {
		t.Errorf("expected nil for nil input, got %v", result)
	}
}

func TestGetFlowSummaryRequestValidation(t *testing.T) {
	tests := []struct {
		name    string
		req     *GetFlowSummaryRequest
		wantErr bool
		errMsg  string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
			errMsg:  "request is nil",
		},
		{
			name: "zero AccID",
			req: &GetFlowSummaryRequest{
				AccID: 0,
			},
			wantErr: true,
			errMsg:  "account ID",
		},
		{
			name: "empty ClearingDate with non-zero CashFlowDirection",
			req: &GetFlowSummaryRequest{
				AccID:             12345,
				ClearingDate:      "",
				CashFlowDirection: 1,
			},
			wantErr: true,
			errMsg:  "clearing date is required",
		},
		{
			name: "valid request with ClearingDate",
			req: &GetFlowSummaryRequest{
				AccID:        12345,
				ClearingDate: "2026-04-08",
			},
			wantErr: false,
		},
		{
			name: "valid request without ClearingDate when direction is zero",
			req: &GetFlowSummaryRequest{
				AccID:             12345,
				CashFlowDirection: 0,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateGetFlowSummaryRequest(tc.req)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tc.errMsg)
				} else if !containsString(err.Error(), tc.errMsg) {
					t.Errorf("expected error containing %q, got %q", tc.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %q", err.Error())
				}
			}
		})
	}
}

func validateGetFlowSummaryRequest(req *GetFlowSummaryRequest) error {
	if req == nil {
		return fmt.Errorf("GetFlowSummary: request is nil")
	}
	if req.AccID == 0 {
		return fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.ClearingDate == "" && req.CashFlowDirection != 0 {
		return fmt.Errorf("clearing date is required when cash flow direction is specified")
	}
	return nil
}

func containsString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
