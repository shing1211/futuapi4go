package qot

import (
	"testing"

	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

func TestGetKLRequestValidation(t *testing.T) {
	// Test valid request
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    10,
	}

	if req.Security == nil {
		t.Error("Security should not be nil")
	}
	if req.Security.GetCode() != code {
		t.Errorf("expected code %s, got %s", code, req.Security.GetCode())
	}
	if req.ReqNum != 10 {
		t.Errorf("expected ReqNum 10, got %d", req.ReqNum)
	}
}

func TestGetOrderBookRequestValidation(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &GetOrderBookRequest{
		Security: security,
		Num:      10,
	}

	if req.Num != 10 {
		t.Errorf("expected Num 10, got %d", req.Num)
	}
}

func TestSubscribeRequestValidation(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &SubscribeRequest{
		SecurityList:         []*qotcommon.Security{security},
		SubTypeList:          []SubType{SubType_Basic, SubType_KL},
		IsSubOrUnSub:         true,
		IsRegOrUnRegPush:     true,
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
	if len(req.SubTypeList) != 2 {
		t.Errorf("expected 2 sub types, got %d", len(req.SubTypeList))
	}
	if !req.IsSubOrUnSub {
		t.Error("expected IsSubOrUnSub to be true")
	}
}

func TestBasicQotStructFields(t *testing.T) {
	bq := &BasicQot{
		Security:  &qotcommon.Security{Market: func() *int32 { v := int32(1); return &v }(), Code: func() *string { s := "00700"; return &s }()},
		Name:      "Tencent",
		CurPrice:  350.50,
		OpenPrice: 348.00,
		HighPrice: 352.00,
		LowPrice:  347.00,
		Volume:    12345678,
		Turnover:  4321098765.00,
	}

	if bq.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", bq.Security.GetCode())
	}
	if bq.Name != "Tencent" {
		t.Errorf("expected name Tencent, got %s", bq.Name)
	}
	if bq.CurPrice != 350.50 {
		t.Errorf("expected CurPrice 350.50, got %f", bq.CurPrice)
	}
}

func TestKLineStructFields(t *testing.T) {
	kl := &KLine{
		Time:           "2026-04-08 15:00:00",
		IsBlank:        false,
		HighPrice:      352.00,
		OpenPrice:      348.00,
		LowPrice:       347.00,
		ClosePrice:     350.50,
		LastClosePrice: 349.00,
		Volume:         12345678,
		Turnover:       4321098765.00,
		ChangeRate:     0.43,
		Timestamp:      1775635200.0,
	}

	if kl.Time != "2026-04-08 15:00:00" {
		t.Errorf("unexpected Time: %s", kl.Time)
	}
	if kl.ClosePrice != 350.50 {
		t.Errorf("expected ClosePrice 350.50, got %f", kl.ClosePrice)
	}
}
