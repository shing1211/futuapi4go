package util

import "testing"

func TestPricePrecision(t *testing.T) {
	tests := []struct {
		market int32
		want   int
	}{
		{1, 3},  // HK
		{11, 2}, // US
		{21, 2}, // SH
		{22, 2}, // SZ
		{31, 2}, // SG
		{0, 2},  // unknown
		{99, 2}, // unknown
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PricePrecision(tt.market); got != tt.want {
				t.Errorf("PricePrecision(%d) = %d, want %d", tt.market, got, tt.want)
			}
		})
	}
}

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		price  float64
		market int32
		want   string
	}{
		{350.5, 1, "350.500"},
		{150.25, 11, "150.25"},
		{100.0, 21, "100.00"},
		{0.0, 1, "0.000"},
		{99.9999, 1, "100.000"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatPrice(tt.price, tt.market); got != tt.want {
				t.Errorf("FormatPrice(%f, %d) = %s, want %s", tt.price, tt.market, got, tt.want)
			}
		})
	}
}

func TestRoundToTickSize(t *testing.T) {
	tests := []struct {
		price   float64
		tick    float64
		want    float64
	}{
		{350.055, 0.01, 350.05},
		{350.05, 0.01, 350.05},
		{100.999, 0.01, 100.99},
		{100.0, 0.01, 100.0},
		{0.0, 0.01, 0.0},
		{150.25, 0.0, 150.25}, // zero tick returns price unchanged
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RoundToTickSize(tt.price, tt.tick)
			if !floatAlmostEqual(got, tt.want, 1e-9) {
				t.Errorf("RoundToTickSize(%f, %f) = %f, want %f", tt.price, tt.tick, got, tt.want)
			}
		})
	}
}

func floatAlmostEqual(a, b, epsilon float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}

func TestPriceToString(t *testing.T) {
	tests := []struct {
		price     float64
		precision int
		want      string
	}{
		{350.5, 2, "350.50"},
		{350.0, 0, "350"},
		{0.0, 2, "0.00"},
		{99.9999, 3, "100.000"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := PriceToString(tt.price, tt.precision); got != tt.want {
				t.Errorf("PriceToString(%f, %d) = %s, want %s", tt.price, tt.precision, got, tt.want)
			}
		})
	}
}

func TestFormatPercent(t *testing.T) {
	tests := []struct {
		rate float64
		want string
	}{
		{0.4321, "43.21%"},
		{0.0, "0.00%"},
		{1.0, "100.00%"},
		{-0.05, "-5.00%"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatPercent(tt.rate); got != tt.want {
				t.Errorf("FormatPercent(%f) = %s, want %s", tt.rate, got, tt.want)
			}
		})
	}
}
