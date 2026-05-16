package util

import (
	"fmt"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

// NewSecurity creates a Security pointer from market and code.
//
// Example:
//
//	sec := NewSecurity(1, "00700")  // HK.00700
func NewSecurity(market int32, code string) *qotcommon.Security {
	m := market
	return &qotcommon.Security{Market: &m, Code: &code}
}

// NewSecurityList creates a slice of Security pointers, one per code, all
// sharing the same market.
//
// Example:
//
//	secs := NewSecurityList(1, []string{"00700", "09988"})
func NewSecurityList(market int32, codes []string) []*qotcommon.Security {
	securities := make([]*qotcommon.Security, len(codes))
	for i, code := range codes {
		securities[i] = NewSecurity(market, code)
	}
	return securities
}

// SecurityToString converts a Security pointer to its "Market.Code" string representation.
//
// Example:
//
//	SecurityToString(NewSecurity(1, "00700")) -> "HK.00700"
func SecurityToString(sec *qotcommon.Security) string {
	if sec == nil {
		return ""
	}
	prefix := marketToPrefix(sec.GetMarket())
	if prefix == "" {
		return sec.GetCode()
	}
	return prefix + "." + sec.GetCode()
}

// SecurityEqual returns true if both Security pointers have the same market and code.
func SecurityEqual(a, b *qotcommon.Security) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.GetMarket() == b.GetMarket() && a.GetCode() == b.GetCode()
}

// SecurityListToStrings converts a slice of Security pointers to their string
// representations.
func SecurityListToStrings(securities []*qotcommon.Security) []string {
	if securities == nil {
		return nil
	}
	strs := make([]string, len(securities))
	for i, sec := range securities {
		if sec != nil {
			strs[i] = SecurityToString(sec)
		}
	}
	return strs
}

// FormatFullCode creates a "Market.Code" string from market and code without
// requiring a Security struct.
func FormatFullCode(market int32, code string) string {
	prefix := marketToPrefix(market)
	if prefix == "" {
		return code
	}
	return fmt.Sprintf("%s.%s", prefix, code)
}
