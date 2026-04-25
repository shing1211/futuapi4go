// Package option provides utilities for working with option contracts
// in the Futu OpenAPI.
//
// This includes parsing option codes, calculating Greeks (if available),
// and finding ATM options.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/option"
//
//	// Parse option code
//	parsed := option.ParseCode("24011900700P")  // 2024-01-19 700P
//	fmt.Println(parsed.Expiry)   // 2024-01-19
//	fmt.Println(parsed.Strike)   // 700.0
//	fmt.Println(parsed.Type)    // Put
//
//	// Find ATM option
//	atms := option.FindAtm("AAPL", "2024-02-16", 180.0)
//	fmt.Println(atms.Call.Code)  // e.g., "24021600180C"
//	fmt.Println(atms.Put.Code)    // e.g., "24021600180P"
package option

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type OptionType string

const (
	OptionTypeCall OptionType = "Call"
	OptionTypePut  OptionType = "Put"
)

type OptionCode struct {
	Code     string    // Full option code (e.g., "24011900700P")
	Expiry   time.Time // Expiration date
	Strike   float64  // Strike price
	Type     OptionType // Call or Put
	Underlying string  // Underlying code (e.g., "00700")
}

// ParseCode parses a Futu option code into its components.
//
// Futu option codes follow the format: YYMMDDStrikeType
// Examples:
//   - "24011900700P" -> 2024-01-19, 700, Put
//   - "24021600180C" -> 2024-02-16, 180, Call
//   - "230930AAPL220P" -> 2023-09-30, 220, Put (US options)
func ParseCode(code string) (*OptionCode, error) {
	if code == "" {
		return nil, fmt.Errorf("option code cannot be empty")
	}

	// Try to detect format based on code length
	// HK options: 11 chars (YYMMDDZZZZZType)
	// US options: variable length with underlying

	var expiry time.Time
	var strike float64
	var optType OptionType
	var underlying string

	code = strings.ToUpper(code)

	// Detect if it's HK or US format
	if len(code) == 11 && isNumeric(code[:7]) {
		// HK format: YYMMDD + 5 digit strike + P/C
		yy := code[0:2]
		mm := code[2:4]
		dd := code[4:6]

		year, err := strconv.Atoi("20" + yy)
		if err != nil {
			return nil, fmt.Errorf("invalid year in option code: %s", yy)
		}

		month, err := strconv.Atoi(mm)
		if err != nil {
			return nil, fmt.Errorf("invalid month in option code: %s", mm)
		}

		day, err := strconv.Atoi(dd)
		if err != nil {
			return nil, fmt.Errorf("invalid day in option code: %s", dd)
		}

		expiry = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		strikeStr := code[6:10]
		var err2 error
		strike, err2 = strconv.ParseFloat(strikeStr, 64)
		if err2 != nil {
			return nil, fmt.Errorf("invalid strike in option code: %s", strikeStr)
		}

		typeChar := code[10:11]
		if typeChar == "P" {
			optType = OptionTypePut
		} else if typeChar == "C" {
			optType = OptionTypeCall
		} else {
			return nil, fmt.Errorf("invalid option type: %s (expected P or C)", typeChar)
		}

	} else if len(code) >= 8 && isNumeric(code[:6]) {
		// Another HK variant (with leading zeros in strike)
		yy := code[0:2]
		mm := code[2:4]
		dd := code[4:6]

		year, err := strconv.Atoi("20" + yy)
		if err != nil {
			return nil, fmt.Errorf("invalid year in option code: %s", yy)
		}

		month, err := strconv.Atoi(mm)
		if err != nil {
			return nil, fmt.Errorf("invalid month in option code: %s", mm)
		}

		day, err := strconv.Atoi(dd)
		if err != nil {
			return nil, fmt.Errorf("invalid day in option code: %s", dd)
		}

		expiry = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		// Find the letter part (P or C) at the end
		lastChar := code[len(code)-1:]
		if lastChar == "P" {
			optType = OptionTypePut
		} else if lastChar == "C" {
			optType = OptionTypeCall
		} else {
			return nil, fmt.Errorf("invalid option type: %s", lastChar)
		}

		// Strike is between date and type
		strikeStr := code[6 : len(code)-1]
		var err3 error
		strike, err3 = strconv.ParseFloat(strikeStr, 64)
		if err3 != nil {
			return nil, fmt.Errorf("invalid strike in option code: %s", strikeStr)
		}

	} else {
		return nil, fmt.Errorf("unrecognized option code format: %s", code)
	}

	return &OptionCode{
		Code:     code,
		Expiry:   expiry,
		Strike:   strike,
		Type:     optType,
		Underlying: underlying,
	}, nil
}

// isNumeric checks if a string contains only numeric characters
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Format returns the option code in standard format
func (o *OptionCode) Format() string {
	if o.Code != "" {
		return o.Code
	}
	return fmt.Sprintf("%s%.2f%s", o.Type, o.Strike, o.Type)
}

// AtmsResult contains ATM call and put options
type AtmsResult struct {
	Call *OptionCode
	Put  *OptionCode
}

// StrikeDistance calculates the absolute distance from strike to spot price
func StrikeDistance(strike, spot float64) float64 {
	if strike > spot {
		return strike - spot
	}
	return spot - strike
}

// FindAtm finds the ATM (At-The-Money) call and put options from a list.
//
// It calculates which strike is closest to the spot price and returns
// both the call and put at that strike.
func FindAtm(options []*OptionCode, spot float64) *AtmsResult {
	if len(options) == 0 {
		return nil
	}

	var calls []*OptionCode
	var puts []*OptionCode

	// Separate calls and puts
	for _, opt := range options {
		if opt.Type == OptionTypeCall {
			calls = append(calls, opt)
		} else {
			puts = append(puts, opt)
		}
	}

	// Find ATM from calls
	var atmCall *OptionCode
	var minCallDist float64 = 1e9
	for _, c := range calls {
		dist := StrikeDistance(c.Strike, spot)
		if dist < minCallDist {
			minCallDist = dist
			atmCall = c
		}
	}

	// Find ATM from puts
	var atmPut *OptionCode
	var minPutDist float64 = 1e9
	for _, p := range puts {
		dist := StrikeDistance(p.Strike, spot)
		if dist < minPutDist {
			minPutDist = dist
			atmPut = p
		}
	}

	return &AtmsResult{
		Call: atmCall,
		Put:  atmPut,
	}
}

// FilterByExpiry returns options that expire on the specified date
func FilterByExpiry(options []*OptionCode, expiry time.Time) []*OptionCode {
	var result []*OptionCode
	for _, opt := range options {
		if opt.Expiry.Year() == expiry.Year() &&
			opt.Expiry.Month() == expiry.Month() &&
			opt.Expiry.Day() == expiry.Day() {
			result = append(result, opt)
		}
	}
	return result
}

// FilterByStrikeRange returns options with strikes between min and max
func FilterByStrikeRange(options []*OptionCode, min, max float64) []*OptionCode {
	var result []*OptionCode
	for _, opt := range options {
		if opt.Strike >= min && opt.Strike <= max {
			result = append(result, opt)
		}
	}
	return result
}