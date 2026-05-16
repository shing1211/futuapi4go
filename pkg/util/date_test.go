package util

import (
	"testing"
	"time"
)

func TestTodayDateStr(t *testing.T) {
	got := TodayDateStr()
	expected := time.Now().Format(FutuDateFormat)
	if got != expected {
		t.Errorf("TodayDateStr() = %s, want %s", got, expected)
	}
	// Verify the format is exactly "2006-01-02"
	if len(got) != 10 {
		t.Errorf("TodayDateStr() length = %d, want 10", len(got))
	}
}

func TestNowDateTimeStr(t *testing.T) {
	got := NowDateTimeStr()
	expected := time.Now().Format(FutuDateTimeFormat)
	if got != expected {
		t.Errorf("NowDateTimeStr() = %s, want %s", got, expected)
	}
	// Verify the format is exactly "2006-01-02 15:04:05"
	if len(got) != 19 {
		t.Errorf("NowDateTimeStr() length = %d, want 19", len(got))
	}
}

func TestParseFutuDate(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
		err   bool
	}{
		{"2026-04-08", time.Date(2026, 4, 8, 0, 0, 0, 0, time.UTC), false},
		{"2026-01-01", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"", time.Time{}, true},
		{"invalid", time.Time{}, true},
		{"2026-13-01", time.Time{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseFutuDate(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("ParseFutuDate(%q) error = %v, want err=%v", tt.input, err, tt.err)
				return
			}
			if !tt.err && !got.Equal(tt.want) {
				t.Errorf("ParseFutuDate(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseFutuDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
		err   bool
	}{
		{"2026-04-08 10:00:00", time.Date(2026, 4, 8, 10, 0, 0, 0, time.UTC), false},
		{"2026-01-01 00:00:00", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"", time.Time{}, true},
		{"invalid", time.Time{}, true},
		{"2026-04-08", time.Time{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseFutuDateTime(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("ParseFutuDateTime(%q) error = %v, want err=%v", tt.input, err, tt.err)
				return
			}
			if !tt.err && !got.Equal(tt.want) {
				t.Errorf("ParseFutuDateTime(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormatFutuDate(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2026, 4, 8, 10, 0, 0, 0, time.UTC), "2026-04-08"},
		{time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), "2026-01-01"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatFutuDate(tt.input); got != tt.want {
				t.Errorf("FormatFutuDate(%v) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormatFutuDateTime(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2026, 4, 8, 10, 0, 0, 0, time.UTC), "2026-04-08 10:00:00"},
		{time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), "2026-01-01 00:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatFutuDateTime(tt.input); got != tt.want {
				t.Errorf("FormatFutuDateTime(%v) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidFutuDate(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"2026-04-08", true},
		{"2026-01-01", true},
		{"", false},
		{"invalid", false},
		{"2026-13-01", false},
		{"2026-01-32", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsValidFutuDate(tt.input); got != tt.want {
				t.Errorf("IsValidFutuDate(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidFutuDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"2026-04-08 10:00:00", true},
		{"2026-01-01 00:00:00", true},
		{"", false},
		{"invalid", false},
		{"2026-04-08", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsValidFutuDateTime(tt.input); got != tt.want {
				t.Errorf("IsValidFutuDateTime(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
