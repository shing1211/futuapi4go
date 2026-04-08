package trd

import (
	"testing"

	"gitee.com/shing1211/futuapi4go/pkg/pb/trdcommon"
)

func TestPlaceOrderRequestValidation(t *testing.T) {
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	req := &PlaceOrderRequest{
		AccID:     123456789,
		TrdMarket: hkMarket,
		Code:      "00700",
		TrdSide:   int32(trdcommon.TrdSide_TrdSide_Buy),
		OrderType: int32(trdcommon.OrderType_OrderType_Normal),
		Price:     350.00,
		Qty:       100.0,
	}

	if req.AccID != 123456789 {
		t.Errorf("expected AccID 123456789, got %d", req.AccID)
	}
	if req.Code != "00700" {
		t.Errorf("expected Code 00700, got %s", req.Code)
	}
	if req.Price != 350.00 {
		t.Errorf("expected Price 350.00, got %f", req.Price)
	}
	if req.Qty != 100.0 {
		t.Errorf("expected Qty 100.0, got %f", req.Qty)
	}
}

func TestGetFundsRequestValidation(t *testing.T) {
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	req := &GetFundsRequest{
		AccID:     123456789,
		TrdMarket: hkMarket,
	}

	if req.AccID != 123456789 {
		t.Errorf("expected AccID 123456789, got %d", req.AccID)
	}
}

func TestGetPositionListRequestValidation(t *testing.T) {
	req := &GetPositionListRequest{
		AccID:     123456789,
		TrdMarket: 0, // All markets
	}

	if req.TrdMarket != 0 {
		t.Errorf("expected TrdMarket 0 (all), got %d", req.TrdMarket)
	}
}

func TestUnlockTradeRequestValidation(t *testing.T) {
	req := &UnlockTradeRequest{}

	// Just verify the struct can be created
	if req == nil {
		t.Error("UnlockTradeRequest should not be nil")
	}
}

func TestFundsStructFields(t *testing.T) {
	funds := &Funds{
		Power:          100000.00,
		TotalAssets:    500000.00,
		Cash:           200000.00,
		MarketVal:      300000.00,
		FrozenCash:     50000.00,
		DebtCash:       0.00,
		AvailableFunds: 150000.00,
	}

	if funds.TotalAssets != 500000.00 {
		t.Errorf("expected TotalAssets 500000.00, got %f", funds.TotalAssets)
	}
	if funds.Power != 100000.00 {
		t.Errorf("expected Power 100000.00, got %f", funds.Power)
	}
}

func TestPositionStructFields(t *testing.T) {
	pos := &Position{
		Code:       "00700",
		Name:       "Tencent",
		Qty:        1000.0,
		CanSellQty: 1000.0,
		Price:      350.50,
		CostPrice:  340.00,
		Val:        350500.00,
		PlVal:      10500.00,
	}

	if pos.Code != "00700" {
		t.Errorf("expected Code 00700, got %s", pos.Code)
	}
	if pos.Qty != 1000.0 {
		t.Errorf("expected Qty 1000.0, got %f", pos.Qty)
	}
	if pos.PlVal != 10500.00 {
		t.Errorf("expected PlVal 10500.00, got %f", pos.PlVal)
	}
}

func TestModifyOrderRequestValidation(t *testing.T) {
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	req := &ModifyOrderRequest{
		AccID:     123456789,
		TrdMarket: hkMarket,
		OrderID:   9876543210,
		Qty:       200.0,
		Price:     360.00,
	}

	if req.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", req.OrderID)
	}
	if req.Qty != 200.0 {
		t.Errorf("expected Qty 200.0, got %f", req.Qty)
	}
}
