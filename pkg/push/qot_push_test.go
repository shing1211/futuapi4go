package push

import (
	"testing"

	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotupdatebasicqot"
)

func TestParseUpdateBasicQotEmptyBody(t *testing.T) {
	_, err := ParseUpdateBasicQot([]byte{})
	if err == nil {
		t.Error("ParseUpdateBasicQot should fail with empty body")
	}
}

func TestParseUpdateBasicQotValidBody(t *testing.T) {
	// Create a valid BasicQot response
	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	name := "Tencent"
	price := 350.50

	bq := &qotcommon.BasicQot{
		Security: &qotcommon.Security{
			Market: &market,
			Code:   &code,
		},
		Name:     &name,
		CurPrice: &price,
	}

	resp := &qotupdatebasicqot.Response{
		RetType: func() *int32 { v := int32(0); return &v }(),
		S2C: &qotupdatebasicqot.S2C{
			BasicQotList: []*qotcommon.BasicQot{bq},
		},
	}

	body, err := resp.MarshalVT()
	if err != nil {
		t.Skipf("skipping: protobuf marshal error: %v", err)
	}

	result, err := ParseUpdateBasicQot(body)
	if err != nil {
		t.Fatalf("ParseUpdateBasicQot failed: %v", err)
	}

	if len(result.BasicQotList) != 1 {
		t.Errorf("expected 1 BasicQot, got %d", len(result.BasicQotList))
	}
}
