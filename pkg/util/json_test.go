package util

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	v := map[string]interface{}{"code": "00700", "price": 350.5}
	got := ToJSON(v)
	if !strings.Contains(got, `"code":"00700"`) {
		t.Errorf("ToJSON missing code field: %s", got)
	}
	if !strings.Contains(got, `"price":350.5`) {
		t.Errorf("ToJSON missing price field: %s", got)
	}
}

func TestToJSONError(t *testing.T) {
	// channel cannot be JSON marshalled
	got := ToJSON(make(chan int))
	if got != "null" {
		t.Errorf("ToJSON(channel) = %s, want null", got)
	}
}

func TestToJSONPretty(t *testing.T) {
	v := map[string]string{"code": "00700"}
	got := ToJSONPretty(v)
	if !strings.Contains(got, "\n") {
		t.Errorf("ToJSONPretty should contain newlines: %s", got)
	}
	if !strings.Contains(got, `"code"`) {
		t.Errorf("ToJSONPretty missing code field: %s", got)
	}
}

func TestToJSONPrettyError(t *testing.T) {
	got := ToJSONPretty(make(chan int))
	if got != "null" {
		t.Errorf("ToJSONPretty(channel) = %s, want null", got)
	}
}

func TestToCSV(t *testing.T) {
	headers := []string{"Code", "Price", "Name"}
	rows := [][]string{
		{"00700", "350.50", "Tencent"},
		{"AAPL", "150.25", "Apple, Inc."},
	}
	got := ToCSV(headers, rows)

	if !strings.Contains(got, "Code") {
		t.Errorf("CSV missing Code header")
	}
	if !strings.Contains(got, `"Apple, Inc."`) {
		t.Errorf("CSV should quote comma-containing fields: %s", got)
	}
	if !strings.HasSuffix(got, "\r\n") {
		t.Errorf("CSV should end with CRLF: %q", got)
	}
}

func TestToCSVQuoting(t *testing.T) {
	headers := []string{"A"}
	rows := [][]string{{`contains "quote"`}}
	got := ToCSV(headers, rows)
	if !strings.Contains(got, `""quote""`) {
		t.Errorf("CSV should escape quotes: %s", got)
	}
}

func TestToTable(t *testing.T) {
	headers := []string{"Code", "Price"}
	rows := [][]string{
		{"00700", "350.50"},
		{"AAPL", "150.25"},
	}
	got := ToTable(headers, rows)

	if !strings.Contains(got, "Code") {
		t.Errorf("Table missing Code header")
	}
	if !strings.Contains(got, "00700") {
		t.Errorf("Table missing 00700")
	}
	if !strings.Contains(got, "AAPL") {
		t.Errorf("Table missing AAPL")
	}
	if !strings.Contains(got, "----") {
		t.Errorf("Table missing separator: %s", got)
	}
}

func TestToTableEmpty(t *testing.T) {
	got := ToTable(nil, nil)
	if got != "" {
		t.Errorf("ToTable(nil) = %q, want empty", got)
	}
}
