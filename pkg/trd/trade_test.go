package trd

import (
	"testing"

	"gitee.com/shing1211/futuapi4go/pkg/pb/common"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdflowsummary"
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
		AccID:         123456789,
		TrdMarket:     hkMarket,
		OrderID:       9876543210,
		ModifyOrderOp: 1,
		Qty:           200.0,
		Price:         360.00,
	}

	if req.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", req.OrderID)
	}
	if req.Qty != 200.0 {
		t.Errorf("expected Qty 200.0, got %f", req.Qty)
	}
	if req.ModifyOrderOp != 1 {
		t.Errorf("expected ModifyOrderOp 1, got %d", req.ModifyOrderOp)
	}
}

func TestAccFields(t *testing.T) {
	acc := &Acc{
		TrdEnv:    1,
		AccID:     123456789,
		AccType:   1,
		CardNum:   "ABC123",
		AccStatus: 1,
	}

	if acc.TrdEnv != 1 {
		t.Errorf("expected TrdEnv 1, got %d", acc.TrdEnv)
	}
	if acc.AccID != 123456789 {
		t.Errorf("expected AccID 123456789, got %d", acc.AccID)
	}
	if acc.AccType != 1 {
		t.Errorf("expected AccType 1, got %d", acc.AccType)
	}
	if acc.CardNum != "ABC123" {
		t.Errorf("expected CardNum ABC123, got %s", acc.CardNum)
	}
	if acc.AccStatus != 1 {
		t.Errorf("expected AccStatus 1, got %d", acc.AccStatus)
	}
}

func TestGetAccListResponseConstruction(t *testing.T) {
	rsp := &GetAccListResponse{
		AccList: []*Acc{
			{TrdEnv: 1, AccID: 123456789, AccType: 1},
		},
	}

	if len(rsp.AccList) != 1 {
		t.Errorf("expected 1 account, got %d", len(rsp.AccList))
	}
	if rsp.AccList[0].AccID != 123456789 {
		t.Errorf("expected AccID 123456789, got %d", rsp.AccList[0].AccID)
	}
}

func TestFundsFieldsComplete(t *testing.T) {
	funds := &Funds{
		Power:          100000.00,
		TotalAssets:    500000.00,
		Cash:           200000.00,
		MarketVal:      300000.00,
		FrozenCash:     50000.00,
		DebtCash:       10000.00,
		Currency:       1,
		AvailableFunds: 150000.00,
	}

	if funds.Power != 100000.00 {
		t.Errorf("expected Power 100000.00, got %f", funds.Power)
	}
	if funds.Currency != 1 {
		t.Errorf("expected Currency 1, got %d", funds.Currency)
	}
	if funds.DebtCash != 10000.00 {
		t.Errorf("expected DebtCash 10000.00, got %f", funds.DebtCash)
	}
	if funds.AvailableFunds != 150000.00 {
		t.Errorf("expected AvailableFunds 150000.00, got %f", funds.AvailableFunds)
	}
}

func TestGetFundsResponseConstruction(t *testing.T) {
	rsp := &GetFundsResponse{
		Funds: &Funds{
			Power:       100000.00,
			TotalAssets: 500000.00,
		},
	}

	if rsp.Funds == nil {
		t.Fatal("Funds should not be nil")
	}
	if rsp.Funds.TotalAssets != 500000.00 {
		t.Errorf("expected TotalAssets 500000.00, got %f", rsp.Funds.TotalAssets)
	}
}

func TestPositionFieldsComplete(t *testing.T) {
	pos := &Position{
		Code:       "00700",
		Name:       "Tencent",
		Qty:        1000.0,
		CanSellQty: 800.0,
		Price:      350.50,
		CostPrice:  340.00,
		Val:        350500.00,
		PlVal:      10500.00,
		PlRatio:    3.09,
	}

	if pos.CanSellQty != 800.0 {
		t.Errorf("expected CanSellQty 800.0, got %f", pos.CanSellQty)
	}
	if pos.PlRatio != 3.09 {
		t.Errorf("expected PlRatio 3.09, got %f", pos.PlRatio)
	}
}

func TestGetPositionListResponseConstruction(t *testing.T) {
	rsp := &GetPositionListResponse{
		PositionList: []*Position{
			{Code: "00700", Name: "Tencent", Qty: 1000.0},
		},
	}

	if len(rsp.PositionList) != 1 {
		t.Errorf("expected 1 position, got %d", len(rsp.PositionList))
	}
}

func TestOrderFields(t *testing.T) {
	order := &Order{
		OrderID:      9876543210,
		Code:         "00700",
		Name:         "Tencent",
		TrdSide:      1,
		OrderType:    1,
		OrderStatus:  2,
		Price:        350.00,
		Qty:          100.0,
		FillQty:      50.0,
		CreateTime:   "2026-04-08 10:00:00",
		UpdateTime:   "2026-04-08 10:05:00",
		FillAvgPrice: 350.00,
	}

	if order.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", order.OrderID)
	}
	if order.OrderStatus != 2 {
		t.Errorf("expected OrderStatus 2, got %d", order.OrderStatus)
	}
	if order.FillQty != 50.0 {
		t.Errorf("expected FillQty 50.0, got %f", order.FillQty)
	}
	if order.FillAvgPrice != 350.00 {
		t.Errorf("expected FillAvgPrice 350.00, got %f", order.FillAvgPrice)
	}
}

func TestGetOrderListResponseConstruction(t *testing.T) {
	rsp := &GetOrderListResponse{
		OrderList: []*Order{
			{OrderID: 9876543210, Code: "00700", TrdSide: 1},
		},
	}

	if len(rsp.OrderList) != 1 {
		t.Errorf("expected 1 order, got %d", len(rsp.OrderList))
	}
}

func TestOrderFillFields(t *testing.T) {
	fill := &OrderFill{
		OrderID:    9876543210,
		FillID:     1111111111,
		Code:       "00700",
		Name:       "Tencent",
		TrdSide:    1,
		Price:      350.00,
		Qty:        100.0,
		CreateTime: "2026-04-08 10:05:00",
	}

	if fill.FillID != 1111111111 {
		t.Errorf("expected FillID 1111111111, got %d", fill.FillID)
	}
	if fill.TrdSide != 1 {
		t.Errorf("expected TrdSide 1, got %d", fill.TrdSide)
	}
	if fill.CreateTime != "2026-04-08 10:05:00" {
		t.Errorf("expected CreateTime 2026-04-08 10:05:00, got %s", fill.CreateTime)
	}
}

func TestGetOrderFillListResponseConstruction(t *testing.T) {
	rsp := &GetOrderFillListResponse{
		OrderFillList: []*OrderFill{
			{FillID: 1111111111, Code: "00700"},
		},
	}

	if len(rsp.OrderFillList) != 1 {
		t.Errorf("expected 1 fill, got %d", len(rsp.OrderFillList))
	}
}

func TestPlaceOrderResponseConstruction(t *testing.T) {
	rsp := &PlaceOrderResponse{
		OrderID: 9876543210,
	}

	if rsp.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", rsp.OrderID)
	}
}

func TestUnlockTradeRequestFields(t *testing.T) {
	req := &UnlockTradeRequest{
		Unlock:       true,
		PwdMD5:       "abcdef123456",
		SecurityFirm: 1,
	}

	if !req.Unlock {
		t.Error("expected Unlock true")
	}
	if req.PwdMD5 != "abcdef123456" {
		t.Errorf("expected PwdMD5 abcdef123456, got %s", req.PwdMD5)
	}
	if req.SecurityFirm != 1 {
		t.Errorf("expected SecurityFirm 1, got %d", req.SecurityFirm)
	}
}

func TestGetOrderFeeRequestConstruction(t *testing.T) {
	req := &GetOrderFeeRequest{
		AccID:         123456789,
		TrdMarket:     1,
		OrderIDExList: []string{"order1", "order2"},
	}

	if len(req.OrderIDExList) != 2 {
		t.Errorf("expected 2 order IDs, got %d", len(req.OrderIDExList))
	}
}

func TestOrderFeeInfoFields(t *testing.T) {
	feeInfo := &OrderFeeInfo{
		OrderIDEx: "order123",
		FeeAmount: 50.0,
		FeeList: []*OrderFeeItemInfo{
			{Title: "Commission", Value: 25.0},
			{Title: "Platform Fee", Value: 25.0},
		},
	}

	if feeInfo.OrderIDEx != "order123" {
		t.Errorf("expected OrderIDEx order123, got %s", feeInfo.OrderIDEx)
	}
	if feeInfo.FeeAmount != 50.0 {
		t.Errorf("expected FeeAmount 50.0, got %f", feeInfo.FeeAmount)
	}
	if len(feeInfo.FeeList) != 2 {
		t.Errorf("expected 2 fee items, got %d", len(feeInfo.FeeList))
	}
}

func TestOrderFeeItemInfoFields(t *testing.T) {
	item := &OrderFeeItemInfo{
		Title: "Commission",
		Value: 25.0,
	}

	if item.Title != "Commission" {
		t.Errorf("expected Title Commission, got %s", item.Title)
	}
	if item.Value != 25.0 {
		t.Errorf("expected Value 25.0, got %f", item.Value)
	}
}

func TestGetOrderFeeResponseConstruction(t *testing.T) {
	rsp := &GetOrderFeeResponse{
		OrderFeeList: []*OrderFeeInfo{
			{OrderIDEx: "order123", FeeAmount: 50.0},
		},
	}

	if len(rsp.OrderFeeList) != 1 {
		t.Errorf("expected 1 fee entry, got %d", len(rsp.OrderFeeList))
	}
}

func TestGetMarginRatioRequestConstruction(t *testing.T) {
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetMarginRatioRequest{
		AccID:        123456789,
		TrdMarket:    1,
		SecurityList: []*qotcommon.Security{security},
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestMarginRatioInfoFields(t *testing.T) {
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	info := &MarginRatioInfo{
		Security:        &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()},
		IsLongPermit:    true,
		IsShortPermit:   false,
		ShortPoolRemain: 5000.0,
		ShortFeeRate:    0.05,
		AlertLongRatio:  1.0,
		AlertShortRatio: 1.5,
		ImLongRatio:     0.6,
		ImShortRatio:    1.2,
		McmLongRatio:    0.5,
		McmShortRatio:   1.0,
		MmLongRatio:     0.3,
		MmShortRatio:    0.8,
	}

	if !info.IsLongPermit {
		t.Error("expected IsLongPermit true")
	}
	if info.IsShortPermit {
		t.Error("expected IsShortPermit false")
	}
	if info.ShortPoolRemain != 5000.0 {
		t.Errorf("expected ShortPoolRemain 5000.0, got %f", info.ShortPoolRemain)
	}
}

func TestGetMarginRatioResponseConstruction(t *testing.T) {
	rsp := &GetMarginRatioResponse{
		MarginRatioInfoList: []*MarginRatioInfo{
			{IsLongPermit: true, IsShortPermit: false},
		},
	}

	if len(rsp.MarginRatioInfoList) != 1 {
		t.Errorf("expected 1 margin ratio entry, got %d", len(rsp.MarginRatioInfoList))
	}
}

func TestGetMaxTrdQtysRequestConstruction(t *testing.T) {
	req := &GetMaxTrdQtysRequest{
		AccID:              123456789,
		TrdMarket:          1,
		OrderType:          1,
		Code:               "00700",
		Price:              350.0,
		AdjustPrice:        true,
		AdjustSideAndLimit: 0.01,
		SecMarket:          1,
		OrderIDEx:          "ORDER123",
	}

	if req.OrderType != 1 {
		t.Errorf("expected OrderType 1, got %d", req.OrderType)
	}
	if !req.AdjustPrice {
		t.Error("expected AdjustPrice true")
	}
}

func TestMaxTrdQtysInfoFields(t *testing.T) {
	info := &MaxTrdQtysInfo{
		MaxCashBuy:          1000.0,
		MaxCashAndMarginBuy: 2000.0,
		MaxPositionSell:     500.0,
		MaxSellShort:        300.0,
		MaxBuyBack:          200.0,
		LongRequiredIM:      100.0,
		ShortRequiredIM:     150.0,
	}

	if info.MaxCashBuy != 1000.0 {
		t.Errorf("expected MaxCashBuy 1000.0, got %f", info.MaxCashBuy)
	}
	if info.MaxSellShort != 300.0 {
		t.Errorf("expected MaxSellShort 300.0, got %f", info.MaxSellShort)
	}
	if info.ShortRequiredIM != 150.0 {
		t.Errorf("expected ShortRequiredIM 150.0, got %f", info.ShortRequiredIM)
	}
}

func TestGetMaxTrdQtysResponseConstruction(t *testing.T) {
	rsp := &GetMaxTrdQtysResponse{
		MaxTrdQtys: &MaxTrdQtysInfo{
			MaxCashBuy: 1000.0,
		},
	}

	if rsp.MaxTrdQtys == nil {
		t.Fatal("MaxTrdQtys should not be nil")
	}
	if rsp.MaxTrdQtys.MaxCashBuy != 1000.0 {
		t.Errorf("expected MaxCashBuy 1000.0, got %f", rsp.MaxTrdQtys.MaxCashBuy)
	}
}

func TestGetHistoryOrderListRequestConstruction(t *testing.T) {
	req := &GetHistoryOrderListRequest{
		AccID:            123456789,
		TrdMarket:        1,
		FilterStatusList: []int32{2, 5},
	}

	if len(req.FilterStatusList) != 2 {
		t.Errorf("expected 2 status filters, got %d", len(req.FilterStatusList))
	}
}

func TestGetHistoryOrderFillListRequestConstruction(t *testing.T) {
	req := &GetHistoryOrderFillListRequest{
		AccID:     123456789,
		TrdMarket: 1,
	}

	if req.AccID != 123456789 {
		t.Errorf("expected AccID 123456789, got %d", req.AccID)
	}
}

func TestSubAccPushRequestConstruction(t *testing.T) {
	req := &SubAccPushRequest{
		AccIDList: []uint64{123456789, 987654321},
	}

	if len(req.AccIDList) != 2 {
		t.Errorf("expected 2 acc IDs, got %d", len(req.AccIDList))
	}
	if req.AccIDList[0] != 123456789 {
		t.Errorf("expected first AccID 123456789, got %d", req.AccIDList[0])
	}
}

func TestReconfirmOrderRequestConstruction(t *testing.T) {
	accID := uint64(123456789)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	req := &ReconfirmOrderRequest{
		PacketID:        &common.PacketID{ConnID: func() *uint64 { v := uint64(1); return &v }(), SerialNo: func() *uint32 { v := uint32(1); return &v }()},
		Header:          &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
		OrderID:         9876543210,
		ReconfirmReason: 1,
	}

	if req.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", req.OrderID)
	}
	if req.ReconfirmReason != 1 {
		t.Errorf("expected ReconfirmReason 1, got %d", req.ReconfirmReason)
	}
}

func TestReconfirmOrderResponseConstruction(t *testing.T) {
	accID := uint64(123456789)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	rsp := &ReconfirmOrderResponse{
		Header:  &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
		OrderID: 9876543210,
	}

	if rsp.OrderID != 9876543210 {
		t.Errorf("expected OrderID 9876543210, got %d", rsp.OrderID)
	}
}

func TestGetFlowSummaryRequestConstruction(t *testing.T) {
	accID := uint64(123456789)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	req := &GetFlowSummaryRequest{
		Header:            &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
		ClearingDate:      "2026-04-08",
		CashFlowDirection: 1,
	}

	if req.ClearingDate != "2026-04-08" {
		t.Errorf("expected ClearingDate 2026-04-08, got %s", req.ClearingDate)
	}
	if req.CashFlowDirection != 1 {
		t.Errorf("expected CashFlowDirection 1, got %d", req.CashFlowDirection)
	}
}

func TestGetFlowSummaryResponseConstruction(t *testing.T) {
	accID := uint64(123456789)
	trdMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)
	rsp := &GetFlowSummaryResponse{
		Header:          &trdcommon.TrdHeader{AccID: &accID, TrdMarket: &trdMarket},
		FlowSummaryList: []*trdflowsummary.FlowSummaryInfo{},
	}

	if rsp.Header == nil {
		t.Fatal("Header should not be nil")
	}
}
