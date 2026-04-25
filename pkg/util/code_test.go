package util

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestParseCode(t *testing.T) {
	tests := []struct {
		input   string
		wantMkt int32
		wantSym string
	}{
		{"HK.00700", constant.Market_HK, "00700"},
		{"US.AAPL", constant.Market_US, "AAPL"},
		{"SH.600519", constant.Market_SH, "600519"},
		{"SZ.000001", constant.Market_SZ, "000001"},
		{"SG.CNmain", constant.Market_SG, "CNmain"},
		{"JP.NKmain", constant.Market_JP, "NKmain"},
		{"AU.CBA", constant.Market_AU, "CBA"},
		{"MY.1155", constant.Market_MY, "1155"},
		{"CA.TD", constant.Market_CA, "TD"},
		{"FX.USDHKD", constant.Market_FX, "USDHKD"},
		{"", constant.Market_None, ""},
		{"NOCODE", constant.Market_None, ""},
		{"HK.", constant.Market_None, ""},
		{".AAPL", constant.Market_None, ""},
		{"HK..00700", constant.Market_HK, ".00700"},
		{"XX.00700", constant.Market_None, "00700"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			mkt, sym := ParseCode(tt.input)
			if mkt != tt.wantMkt || sym != tt.wantSym {
				t.Errorf("ParseCode(%q) = (%d, %q), want (%d, %q)",
					tt.input, mkt, sym, tt.wantMkt, tt.wantSym)
			}
		})
	}
}

func TestFormatCode(t *testing.T) {
	tests := []struct {
		market int32
		code   string
		want   string
	}{
		{constant.Market_HK, "00700", "HK.00700"},
		{constant.Market_US, "AAPL", "US.AAPL"},
		{constant.Market_SH, "600519", "SH.600519"},
		{constant.Market_SZ, "000001", "SZ.000001"},
		{constant.Market_SG, "CNmain", "SG.CNmain"},
		{constant.Market_JP, "NKmain", "JP.NKmain"},
		{constant.Market_AU, "CBA", "AU.CBA"},
		{constant.Market_MY, "1155", "MY.1155"},
		{constant.Market_CA, "TD", "CA.TD"},
		{constant.Market_FX, "USDHKD", "FX.USDHKD"},
		{constant.Market_None, "00700", ""},
		{constant.Market_HK, "", "HK."},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FormatCode(tt.market, tt.code)
			if got != tt.want {
				t.Errorf("FormatCode(%d, %q) = %q, want %q",
					tt.market, tt.code, got, tt.want)
			}
		})
	}
}

func TestParseCodeRoundtrip(t *testing.T) {
	codes := []string{
		"HK.00700", "HK.09988", "HK.03888",
		"US.AAPL", "US.TSLA", "US.GOOG",
		"SH.600519", "SH.600036",
		"SZ.000001", "SZ.300750",
		"SG.CNmain", "SG.NKmain",
		"JP.NKmain",
		"AU.CBA",
		"MY.1155",
		"CA.TD",
		"FX.USDHKD",
	}
	for _, code := range codes {
		t.Run(code, func(t *testing.T) {
			mkt, sym := ParseCode(code)
			got := FormatCode(mkt, sym)
			if got != code {
				t.Errorf("ParseCode(%q) -> FormatCode(%d, %q) = %q, want %q",
					code, mkt, sym, got, code)
			}
		})
	}
}

func TestDetectMarket(t *testing.T) {
	tests := []struct {
		code string
		want int32
	}{
		{"HK.00700", constant.Market_HK},
		{"US.AAPL", constant.Market_US},
		{"SH.600519", constant.Market_SH},
		{"SZ.000001", constant.Market_SZ},
		{"SG.CNmain", constant.Market_SG},
		{"JP.NKmain", constant.Market_JP},
		{"AU.CBA", constant.Market_AU},
		{"MY.1155", constant.Market_MY},
		{"CA.TD", constant.Market_CA},
		{"FX.USDHKD", constant.Market_FX},
		{"", constant.Market_None},
		{"XXX.YYY", constant.Market_None},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			if got := DetectMarket(tt.code); got != tt.want {
				t.Errorf("DetectMarket(%q) = %d, want %d", tt.code, got, tt.want)
			}
		})
	}
}

func TestMarketToTrdMarket(t *testing.T) {
	tests := []struct {
		qotMkt int32
		want   constant.TrdSecMarket
	}{
		{constant.Market_HK, constant.TrdSecMarket_HK},
		{constant.Market_US, constant.TrdSecMarket_US},
		{constant.Market_SH, constant.TrdSecMarket_CN_SH},
		{constant.Market_SZ, constant.TrdSecMarket_CN_SZ},
		{constant.Market_SG, constant.TrdSecMarket_SG},
		{constant.Market_JP, constant.TrdSecMarket_JP},
		{constant.Market_AU, constant.TrdSecMarket_AU},
		{constant.Market_MY, constant.TrdSecMarket_MY},
		{constant.Market_CA, constant.TrdSecMarket_CA},
		{constant.Market_FX, constant.TrdSecMarket_FX},
		{constant.Market_None, constant.TrdSecMarket_Unknown},
	}
	for _, tt := range tests {
		t.Run(tt.want.String(), func(t *testing.T) {
			if got := MarketToTrdMarket(tt.qotMkt); got != tt.want {
				t.Errorf("MarketToTrdMarket(%d) = %d, want %d", tt.qotMkt, got, tt.want)
			}
		})
	}
}

func TestTrdMarketToQotMarket(t *testing.T) {
	tests := []struct {
		trdMkt constant.TrdSecMarket
		want   int32
	}{
		{constant.TrdSecMarket_HK, constant.Market_HK},
		{constant.TrdSecMarket_US, constant.Market_US},
		{constant.TrdSecMarket_CN_SH, constant.Market_SH},
		{constant.TrdSecMarket_CN_SZ, constant.Market_SZ},
		{constant.TrdSecMarket_SG, constant.Market_SG},
		{constant.TrdSecMarket_JP, constant.Market_JP},
		{constant.TrdSecMarket_AU, constant.Market_AU},
		{constant.TrdSecMarket_MY, constant.Market_MY},
		{constant.TrdSecMarket_CA, constant.Market_CA},
		{constant.TrdSecMarket_FX, constant.Market_FX},
		{constant.TrdSecMarket_Unknown, constant.Market_None},
	}
	for _, tt := range tests {
		t.Run(tt.trdMkt.String(), func(t *testing.T) {
			if got := TrdMarketToQotMarket(tt.trdMkt); got != tt.want {
				t.Errorf("TrdMarketToQotMarket(%d) = %d, want %d", tt.trdMkt, got, tt.want)
			}
		})
	}
}

func TestIsMarketValid(t *testing.T) {
	tests := []struct {
		market int32
		want   bool
	}{
		{constant.Market_HK, true},
		{constant.Market_US, true},
		{constant.Market_SH, true},
		{constant.Market_SZ, true},
		{constant.Market_SG, true},
		{constant.Market_JP, true},
		{constant.Market_AU, true},
		{constant.Market_MY, true},
		{constant.Market_CA, true},
		{constant.Market_FX, true},
		{constant.Market_None, false},
		{999, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsMarketValid(tt.market); got != tt.want {
				t.Errorf("IsMarketValid(%d) = %v, want %v", tt.market, got, tt.want)
			}
		})
	}
}
