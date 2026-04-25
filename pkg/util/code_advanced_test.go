package util

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestDetectMarketAdvanced(t *testing.T) {
	tests := []struct {
		name    string
		code   string
		want   int32
	}{
		{"warrant", "#12345", constant.Market_HK},
		{"CBBC", "12345", constant.Market_HK},
		{"warrant with dot", "#12345.HK", constant.Market_HK},
		{"empty", "", constant.Market_None},
		{"regular HK", "HK.00700", constant.Market_HK},
		{"regular US", "US.AAPL", constant.Market_US},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectMarket(tt.code)
			if got != tt.want {
				t.Errorf("DetectMarket(%q) = %d, want %d", tt.code, got, tt.want)
			}
		})
	}
}