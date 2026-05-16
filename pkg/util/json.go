package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ToJSON marshals the given value to a compact JSON string.
//
// Returns "null" if marshalling fails.
func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "null"
	}
	return string(b)
}

// ToJSONPretty marshals the given value to an indented JSON string.
//
// Returns "null" if marshalling fails.
func ToJSONPretty(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "null"
	}
	return string(b)
}

// ToCSV converts a header row and data rows into CSV format.
//
// Each inner slice in rows represents one CSV record.
// All values are quoted to handle special characters.
func ToCSV(headers []string, rows [][]string) string {
	var b strings.Builder
	writeCSVLine(&b, headers)
	for _, row := range rows {
		writeCSVLine(&b, row)
	}
	return b.String()
}

func writeCSVLine(b *strings.Builder, fields []string) {
	for i, f := range fields {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteByte('"')
		b.WriteString(strings.ReplaceAll(f, `"`, `""`))
		b.WriteByte('"')
	}
	b.WriteString("\r\n")
}

// ToTable formats data rows as an aligned text table. Useful for CLI output.
//
// Example:
//
//	util.ToTable(
//	  []string{"Code", "Price"},
//	  [][]string{{"00700", "350.50"}, {"AAPL", "150.25"}},
//	)
//
// Output:
//
//	Code   Price
//	00700  350.50
//	AAPL   150.25
func ToTable(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}

	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	var b strings.Builder
	writeRow := func(fields []string) {
		for i, f := range fields {
			if i > 0 {
				b.WriteString("  ")
			}
			fmt.Fprintf(&b, "%-*s", colWidths[i], f)
		}
		b.WriteByte('\n')
	}

	writeRow(headers)
	sep := make([]string, len(colWidths))
	for i, w := range colWidths {
		sep[i] = strings.Repeat("-", w)
	}
	writeRow(sep)
	for _, row := range rows {
		writeRow(row)
	}
	return b.String()
}
