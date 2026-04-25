// Package util provides code parsing, formatting, and market utility helpers
// for working with Futu OpenAPI stock codes.
//
// All functions are pure and have zero dependencies outside the SDK.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/util"
//
//	market, code := util.ParseCode("HK.00700")           // market=1, code="00700"
//	formatted := util.FormatCode(market, code)           // "HK.00700"
//	trdMkt := util.MarketToTrdMarket(market)            // convert qot->trd market
//	detected := util.DetectMarket("US.AAPL")            // market=11
//
// # Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package util

import (
	"strings"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

const (
	codeSep    = "."
	minCodeLen = 1
	maxCodeLen = 32
)

// ParseCode splits a full stock code string into its market prefix and code.
//
// Examples:
//
//	"HK.00700"     -> market=1 (Market_HK),  code="00700"
//	"US.AAPL"      -> market=11 (Market_US), code="AAPL"
//	"SH.600519"    -> market=21 (Market_SH), code="600519"
//	"SZ.000001"    -> market=22 (Market_SZ), code="000001"
//
// Returns Market_None (0) and empty string if the code is malformed.
func ParseCode(fullCode string) (market int32, code string) {
	if fullCode == "" {
		return constant.Market_None, ""
	}

	idx := strings.Index(fullCode, codeSep)
	if idx <= 0 || idx >= len(fullCode)-1 {
		return constant.Market_None, ""
	}

	prefix := fullCode[:idx]
	sym := fullCode[idx+1:]

	if len(sym) < minCodeLen || len(sym) > maxCodeLen {
		return constant.Market_None, ""
	}

	return detectMarketPrefix(prefix), sym
}

// FormatCode joins a market and code into a full stock code string.
//
// Examples:
//
//	FormatCode(1, "00700")    -> "HK.00700"
//	FormatCode(11, "AAPL")   -> "US.AAPL"
//	FormatCode(21, "600519")  -> "SH.600519"
//
// Returns an empty string if market is unknown.
func FormatCode(market int32, code string) string {
	prefix := marketToPrefix(market)
	if prefix == "" {
		return ""
	}
	return prefix + codeSep + code
}

// DetectMarket extracts the market ID from a full stock code string,
// returning the QotMarket constant.
//
// Examples:
//
//	DetectMarket("HK.00700") -> constant.Market_HK
//	DetectMarket("US.AAPL")  -> constant.Market_US
//
// Returns Market_None (0) if the code is malformed.
func DetectMarket(fullCode string) int32 {
	market, _ := ParseCode(fullCode)
	if market != 0 {
		return market
	}

	// Try to detect from code pattern alone
	return detectCodePattern(fullCode)
}

// detectCodePattern tries to detect market from code pattern without prefix.
// Supports: warrants (# prefixed), CBBC (1 prefixed), futures (.HK suffix)
func detectCodePattern(code string) int32 {
	if code == "" {
		return constant.Market_None
	}

	// Warrant: starts with # (e.g., "#12345")
	if strings.HasPrefix(code, "#") && len(code) >= 5 {
		return constant.Market_HK
	}

	// CBBC: starts with 1 and 5 digits (e.g., "12345")
	if len(code) == 5 && code[0] >= '1' && code[0] <= '9' {
		return constant.Market_HK
	}

	// Futures: contains .HK or .US suffix
	if strings.HasSuffix(code, ".HK") || strings.HasSuffix(code, ".US") || strings.HasSuffix(code, ".CN") {
		return constant.Market_HK
	}

	return constant.Market_None
}

// DetectTradingMarkets returns the TrdMarket and TrdSecMarket for a given full code.
// Supports formats: "00700.HK", "HK.00700", "AAPL.US", "US.AAPL"
//
// Example:
//
//	trdMarket, secMarket := util.DetectTradingMarkets("00700.HK")  // HK, HK
//	trdMarket, secMarket := util.DetectTradingMarkets("AAPL.US")  // US, US
func DetectTradingMarkets(fullCode string) (constant.TrdMarket, constant.TrdSecMarket) {
	market := DetectMarket(fullCode)
	if market == 0 {
		return constant.TrdMarket_None, constant.TrdSecMarket_Unknown
	}
	return constant.TrdMarket(market), constant.MarketToTrdSecMarket[int32(market)]
}

// MarketToTrdMarket converts a QotMarket (quote market) to the corresponding
// TrdSecMarket (trading security market).
//
// This is useful when bridging quote functions with trading functions,
// since some APIs require a TrdSecMarket while you may only have a QotMarket.
//
// Example:
//
//	qotMarket := constant.Market_HK  // 1
//	trdMkt := util.MarketToTrdMarket(qotMarket)  // 1 (TrdSecMarket_HK)
//
// Returns TrdSecMarket_Unknown (0) for unknown markets.
func MarketToTrdMarket(qotMarket int32) constant.TrdSecMarket {
	return constant.MarketToTrdSecMarket[qotMarket]
}

// TrdMarketToQotMarket converts a TrdSecMarket (trading security market) to the
// corresponding QotMarket (quote market).
//
// Returns Market_None (0) for unknown markets.
func TrdMarketToQotMarket(trdMkt constant.TrdSecMarket) int32 {
	switch trdMkt {
	case constant.TrdSecMarket_HK:
		return constant.Market_HK
	case constant.TrdSecMarket_US:
		return constant.Market_US
	case constant.TrdSecMarket_CN_SH:
		return constant.Market_SH
	case constant.TrdSecMarket_CN_SZ:
		return constant.Market_SZ
	case constant.TrdSecMarket_SG:
		return constant.Market_SG
	case constant.TrdSecMarket_JP:
		return constant.Market_JP
	case constant.TrdSecMarket_AU:
		return constant.Market_AU
	case constant.TrdSecMarket_MY:
		return constant.Market_MY
	case constant.TrdSecMarket_CA:
		return constant.Market_CA
	case constant.TrdSecMarket_FX:
		return constant.Market_FX
	default:
		return constant.Market_None
	}
}

// IsMarketValid returns true if the given QotMarket is a known, non-None market.
func IsMarketValid(market int32) bool {
	return market != constant.Market_None && marketToPrefix(market) != ""
}

// marketPrefixMap maps QotMarket IDs to their string prefix.
var marketPrefixMap = map[int32]string{
	constant.Market_HK: "HK",
	constant.Market_US: "US",
	constant.Market_SH: "SH",
	constant.Market_SZ: "SZ",
	constant.Market_SG: "SG",
	constant.Market_JP: "JP",
	constant.Market_AU: "AU",
	constant.Market_MY: "MY",
	constant.Market_CA: "CA",
	constant.Market_FX: "FX",
}

func marketToPrefix(market int32) string {
	if p, ok := marketPrefixMap[market]; ok {
		return p
	}
	return ""
}

func detectMarketPrefix(prefix string) int32 {
	switch prefix {
	case "HK":
		return constant.Market_HK
	case "US":
		return constant.Market_US
	case "SH":
		return constant.Market_SH
	case "SZ":
		return constant.Market_SZ
	case "SG":
		return constant.Market_SG
	case "JP":
		return constant.Market_JP
	case "AU":
		return constant.Market_AU
	case "MY":
		return constant.Market_MY
	case "CA":
		return constant.Market_CA
	case "FX":
		return constant.Market_FX
	default:
		return constant.Market_None
	}
}
