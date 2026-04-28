package constant

import "testing"

func TestLotSize(t *testing.T) {
	tests := []struct {
		name   string
		market TrdMarket
		want   float64
		wantOk bool
	}{
		{"HK", TrdMarket_HK, 100, true},
		{"US", TrdMarket_US, 1, true},
		{"CN", TrdMarket_CN, 100, true},
		{"HKCC", TrdMarket_HKCC, 100, true},
		{"Invalid", TrdMarket(999), 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := LotSize(tt.market)
			if got != tt.want {
				t.Errorf("LotSize() = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("LotSize() ok = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestPriceTick(t *testing.T) {
	tests := []struct {
		name   string
		market TrdMarket
		want   float64
	}{
		{"HK", TrdMarket_HK, 0.01},
		{"US", TrdMarket_US, 0.01},
		{"CN", TrdMarket_CN, 0.01},
		{"HKCC", TrdMarket_HKCC, 0.01},
		{"Unknown", TrdMarket(999), 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PriceTick(tt.market); got != tt.want {
				t.Errorf("PriceTick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateAccID(t *testing.T) {
	if err := ValidateAccID(0); err == nil {
		t.Error("expected error for accID=0")
	}
	if err := ValidateAccID(123456); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateCode(t *testing.T) {
	if err := ValidateCode(""); err == nil {
		t.Error("expected error for empty code")
	}
	if err := ValidateCode("00700.HK"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	longCode := string(make([]byte, 33))
	if err := ValidateCode(longCode); err == nil {
		t.Error("expected error for code > 32 chars")
	}
}

func TestValidateQty(t *testing.T) {
	if err := ValidateQty(0); err == nil {
		t.Error("expected error for qty=0")
	}
	if err := ValidateQty(-1); err == nil {
		t.Error("expected error for negative qty")
	}
	if err := ValidateQty(100); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := ValidateQty(MaxQty + 1); err == nil {
		t.Error("expected error for qty > MaxQty")
	}
}

func TestValidatePrice(t *testing.T) {
	if err := ValidatePrice(-1); err == nil {
		t.Error("expected error for negative price")
	}
	if err := ValidatePrice(100.5); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := ValidatePrice(0); err != nil {
		t.Errorf("unexpected error for price=0: %v", err)
	}
	if err := ValidatePrice(MaxPrice + 1); err == nil {
		t.Error("expected error for price > MaxPrice")
	}
}

func TestValidateRemark(t *testing.T) {
	if err := ValidateRemark(""); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := ValidateRemark("test remark"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	longRemark := string(make([]byte, 257))
	if err := ValidateRemark(longRemark); err == nil {
		t.Error("expected error for remark > 256 chars")
	}
}
