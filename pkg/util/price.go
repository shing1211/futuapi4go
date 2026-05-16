package util

import "fmt"

// PricePrecision returns the number of decimal places for the given market.
func PricePrecision(market int32) int {
	switch market {
	case 21, 22: // SH, SZ
		return 2
	case 1: // HK
		return 3
	case 11: // US
		return 2
	default:
		return 2
	}
}

// FormatPrice formats a price with the standard number of decimal places for
// the given market.
//
// Examples:
//
//	FormatPrice(350.5, 1)   -> "350.500"   (HK)
//	FormatPrice(150.255, 11) -> "150.26"   (US)
func FormatPrice(price float64, market int32) string {
	return fmt.Sprintf("%."+fmt.Sprintf("%df", PricePrecision(market)), price)
}

// RoundToTickSize rounds a price down to the nearest tick size.
//
// Example:
//
//	RoundToTickSize(350.055, 0.01) -> 350.05
func RoundToTickSize(price float64, tickSize float64) float64 {
	if tickSize <= 0 {
		return price
	}
	return float64(int64(price/tickSize)) * tickSize
}

// PriceToString converts a price to string with the given number of decimal places.
func PriceToString(price float64, precision int) string {
	return fmt.Sprintf("%."+fmt.Sprintf("%df", precision), price)
}

// FormatPercent formats a change rate as a percentage string.
//
// Example:
//
//	FormatPercent(0.4321) -> "43.21%"
func FormatPercent(rate float64) string {
	return fmt.Sprintf("%.2f%%", rate*100)
}
