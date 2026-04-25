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
