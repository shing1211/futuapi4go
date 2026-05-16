package util

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

func TestNewSecurity(t *testing.T) {
	sec := NewSecurity(constant.Market_HK, "00700")
	if sec == nil {
		t.Fatal("NewSecurity returned nil")
	}
	if sec.GetMarket() != constant.Market_HK {
		t.Errorf("market = %d, want %d", sec.GetMarket(), constant.Market_HK)
	}
	if sec.GetCode() != "00700" {
		t.Errorf("code = %s, want 00700", sec.GetCode())
	}
}

func TestNewSecurityList(t *testing.T) {
	codes := []string{"00700", "09988", "03888"}
	secs := NewSecurityList(constant.Market_HK, codes)
	if len(secs) != 3 {
		t.Fatalf("len = %d, want 3", len(secs))
	}
	for i, sec := range secs {
		if sec.GetMarket() != constant.Market_HK {
			t.Errorf("sec[%d] market = %d, want %d", i, sec.GetMarket(), constant.Market_HK)
		}
		if sec.GetCode() != codes[i] {
			t.Errorf("sec[%d] code = %s, want %s", i, sec.GetCode(), codes[i])
		}
	}
}

func TestNewSecurityListEmpty(t *testing.T) {
	secs := NewSecurityList(constant.Market_HK, nil)
	if len(secs) != 0 {
		t.Errorf("expected empty list, got len %d", len(secs))
	}
}

func TestSecurityToString(t *testing.T) {
	tests := []struct {
		name string
		sec  *qotcommon.Security
		want string
	}{
		{"HK stock", NewSecurity(constant.Market_HK, "00700"), "HK.00700"},
		{"US stock", NewSecurity(constant.Market_US, "AAPL"), "US.AAPL"},
		{"nil sec", nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecurityToString(tt.sec); got != tt.want {
				t.Errorf("SecurityToString() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSecurityEqual(t *testing.T) {
	secA := NewSecurity(constant.Market_HK, "00700")
	secB := NewSecurity(constant.Market_HK, "00700")
	secC := NewSecurity(constant.Market_HK, "09988")
	secD := NewSecurity(constant.Market_US, "00700")

	tests := []struct {
		name string
		a, b *qotcommon.Security
		want bool
	}{
		{"same pointer", secA, secA, true},
		{"same values", secA, secB, true},
		{"different code", secA, secC, false},
		{"different market", secA, secD, false},
		{"both nil", nil, nil, true},
		{"a nil", nil, secA, false},
		{"b nil", secA, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecurityEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("SecurityEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityListToStrings(t *testing.T) {
	secs := NewSecurityList(constant.Market_HK, []string{"00700", "09988"})
	strs := SecurityListToStrings(secs)
	if len(strs) != 2 {
		t.Fatalf("len = %d, want 2", len(strs))
	}
	if strs[0] != "HK.00700" {
		t.Errorf("strs[0] = %s, want HK.00700", strs[0])
	}
	if strs[1] != "HK.09988" {
		t.Errorf("strs[1] = %s, want HK.09988", strs[1])
	}

	// Test nil input
	if got := SecurityListToStrings(nil); got != nil {
		t.Errorf("expected nil for nil input, got %v", got)
	}
}

func TestFormatFullCode(t *testing.T) {
	tests := []struct {
		market int32
		code   string
		want   string
	}{
		{constant.Market_HK, "00700", "HK.00700"},
		{constant.Market_US, "AAPL", "US.AAPL"},
		{constant.Market_None, "00700", "00700"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatFullCode(tt.market, tt.code); got != tt.want {
				t.Errorf("FormatFullCode(%d, %s) = %s, want %s", tt.market, tt.code, got, tt.want)
			}
		})
	}
}
