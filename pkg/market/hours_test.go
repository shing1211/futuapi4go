package market

import (
	"testing"
	"time"
)

func TestIsHKOpen(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		wantOpen bool
	}{
		{
			name:     "Saturday",
			t:        time.Date(2024, 1, 20, 10, 0, 0, 0, time.UTC), // Saturday
			wantOpen: false,
		},
		{
			name:     "Sunday",
			t:        time.Date(2024, 1, 21, 10, 0, 0, 0, time.UTC), // Sunday
			wantOpen: false,
		},
		{
			name:     "Morning during week",
			t:        time.Date(2024, 1, 22, 2, 0, 0, 0, time.UTC), // Monday 10:00 HKST (2:00 UTC)
			wantOpen: true,
		},
		{
			name:     "After market close",
			t:        time.Date(2024, 1, 22, 18, 0, 0, 0, time.UTC), // Monday 18:00 HKST = 10:00 UTC (after close)
			wantOpen: false,
		},
	}

	loc := time.FixedZone("HKST", 8*3600)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsHKOpen(tt.t.In(loc))
			if got != tt.wantOpen {
				t.Errorf("IsHKOpen() = %v, want %v", got, tt.wantOpen)
			}
		})
	}
}

func TestMarketHours(t *testing.T) {
	tests := []struct {
		market Market
		want   string
	}{
		{MarketHK, "Morning: 9:15-12:00, Afternoon: 13:00-16:30 (HKST)"},
		{MarketUS, "Core: 9:30-16:00 (EST/EDT)"},
		{MarketCN, "Morning: 9:30-11:30, Afternoon: 13:00-15:00 (CST)"},
		{Market("unknown"), "Unknown"},
	}

	for _, tt := range tests {
		t.Run(string(tt.market), func(t *testing.T) {
			got := MarketHours(tt.market)
			if got != tt.want {
				t.Errorf("MarketHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUntilClose(t *testing.T) {
	loc := time.FixedZone("HKST", 8*3600)
	testTime := time.Date(2024, 1, 22, 1, 30, 0, 0, time.UTC).In(loc) // 9:30 HKST

	dur := UntilClose(MarketHK, testTime)
	if dur <= 0 {
		t.Errorf("UntilClose() = %v, want > 0", dur)
	}
}

func TestNextOpen(t *testing.T) {
	loc := time.FixedZone("HKST", 8*3600)
	testTime := time.Date(2024, 1, 22, 2, 0, 0, 0, time.UTC).In(loc) // 10:00 HKST

	dur := NextOpen(MarketHK, testTime)
	if dur <= 0 {
		t.Errorf("NextOpen() = %v, want > 0", dur)
	}
}