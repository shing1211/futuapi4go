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

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestPlaceOrderValidation(t *testing.T) {
	tests := []struct {
		name    string
		req     *PlaceOrderRequest
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
			req: &PlaceOrderRequest{
				AccID:     0,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "account ID",
		},
		{
			name: "empty Code",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "stock code",
		},
		{
			name: "zero Qty",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       0,
			},
			wantErr: true,
			errMsg:  "quantity",
		},
		{
			name: "negative Qty",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       -10,
			},
			wantErr: true,
			errMsg:  "quantity",
		},
		{
			name: "invalid OrderType (zero)",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: 0,
				Price:     350.0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "order type",
		},
		{
			name: "invalid TrdSide (zero)",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   0,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "trade side",
		},
		{
			name: "zero price for Normal order type",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "price",
		},
		{
			name: "negative price for Normal order type",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     -100.0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "price",
		},
		{
			name: "zero price for StopLimit order type",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_StopLimit,
				Price:     0,
				Qty:       100,
			},
			wantErr: true,
			errMsg:  "price",
		},
		{
			name: "zero price for Market order type (allowed)",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Market,
				Price:     0,
				Qty:       100,
			},
			wantErr: false,
		},
		{
			name: "zero price for AbsoluteLimit order type (allowed)",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_AbsoluteLimit,
				Price:     0,
				Qty:       100,
			},
			wantErr: false,
		},
		{
			name: "valid request",
			req: &PlaceOrderRequest{
				AccID:     12345,
				Code:      "00700",
				TrdSide:   constant.TrdSide_Buy,
				OrderType: constant.OrderType_Normal,
				Price:     350.0,
				Qty:       100,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validatePlaceOrderRequest(tc.req)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tc.errMsg)
				} else if tc.errMsg != "" && !contains(err.Error(), tc.errMsg) {
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

func TestModifyOrderValidation(t *testing.T) {
	tests := []struct {
		name    string
		req     *ModifyOrderRequest
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
			req: &ModifyOrderRequest{
				AccID:         0,
				ModifyOrderOp: constant.ModifyOrderOp_Enable,
				OrderID:       12345,
			},
			wantErr: true,
			errMsg:  "account ID",
		},
		{
			name: "zero OrderID with empty OrderIDEx and ForAll=false",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Enable,
				OrderID:       0,
				OrderIDEx:     "",
				ForAll:        false,
			},
			wantErr: true,
			errMsg:  "order ID",
		},
		{
			name: "empty OrderIDEx with zero OrderID and ForAll=false",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Enable,
				OrderID:       0,
				OrderIDEx:     "",
				ForAll:        false,
			},
			wantErr: true,
			errMsg:  "order ID",
		},
		{
			name: "ForAll=true with no OrderID (CancelAllOrder case)",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Disable,
				OrderID:       0,
				OrderIDEx:     "",
				ForAll:        true,
			},
			wantErr: false,
		},
		{
			name: "ForAll=true with OrderID provided",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Disable,
				OrderID:       12345,
				ForAll:        true,
			},
			wantErr: false,
		},
		{
			name: "ForAll=true with OrderIDEx provided",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Disable,
				OrderIDEx:     "ORDER_EX_123",
				ForAll:        true,
			},
			wantErr: false,
		},
		{
			name: "valid request with OrderID only",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Enable,
				OrderID:       12345,
				ForAll:        false,
			},
			wantErr: false,
		},
		{
			name: "valid request with OrderIDEx only",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: constant.ModifyOrderOp_Enable,
				OrderIDEx:     "ORDER_EX_123",
				ForAll:        false,
			},
			wantErr: false,
		},
		{
			name: "invalid ModifyOrderOp (zero)",
			req: &ModifyOrderRequest{
				AccID:         12345,
				ModifyOrderOp: 0,
				OrderID:       12345,
			},
			wantErr: true,
			errMsg:  "modify operation",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateModifyOrderRequest(tc.req)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error containing %q, got nil", tc.errMsg)
				} else if tc.errMsg != "" && !contains(err.Error(), tc.errMsg) {
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

func validatePlaceOrderRequest(req *PlaceOrderRequest) error {
	if req == nil {
		return fmt.Errorf("PlaceOrder: request is nil")
	}
	if req.AccID == 0 {
		return fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.Code == "" {
		return fmt.Errorf("stock code is required")
	}
	if req.Qty <= 0 {
		return fmt.Errorf("invalid quantity: must be positive")
	}
	if req.OrderType <= 0 {
		return fmt.Errorf("invalid order type: must be valid order type constant")
	}
	if req.TrdSide <= 0 {
		return fmt.Errorf("invalid trade side: must be buy/sell/other valid type")
	}
	if req.OrderType != constant.OrderType_Market && req.OrderType != constant.OrderType_AbsoluteLimit && req.Price <= 0 {
		return fmt.Errorf("invalid price: must be positive for non-market order types")
	}
	return nil
}

func validateModifyOrderRequest(req *ModifyOrderRequest) error {
	if req == nil {
		return fmt.Errorf("ModifyOrder: request is nil")
	}
	if req.AccID == 0 {
		return fmt.Errorf("invalid account ID: must be non-zero")
	}
	if req.OrderID == 0 && req.OrderIDEx == "" && !req.ForAll {
		return fmt.Errorf("order ID or OrderIDEx must be provided, or ForAll must be true")
	}
	if req.ModifyOrderOp <= 0 {
		return fmt.Errorf("invalid modify operation: must be valid order operation type")
	}
	return nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
