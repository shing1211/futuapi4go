// Package market provides market hours and trading session utilities
// for various markets supported by Futu OpenAPI.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/market"
//
//	// Check if market is open
//	if market.IsHKOpen(time.Now()) {
//	    fmt.Println("HK market is open")
//	}
//
//	// Get time until market closes
//	untilClose := market.UntilHKClose(time.Now())
//	fmt.Printf("HK closes in %v\n", untilClose)
package market

import (
	"fmt"
	"time"
)

// Market represents a trading market
type Market string

const (
	MarketHK       Market = "HK"  // Hong Kong
	MarketUS       Market = "US"  // United States
	MarketCN       Market = "CN"  // China (SH/SZ)
	MarketSG       Market = "SG"  // Singapore
	MarketJP       Market = "JP"  // Japan
	MarketAU       Market = "AU"  // Australia
)

// Session represents a trading session within a market day
type Session struct {
	Name  string
	Start time.Duration // Minutes from market midnight
	End   time.Duration // Minutes from market midnight
}

// Standard HK market sessions
var (
	// Hong Kong morning session: 9:15 - 12:00
	HKSessionMorning = Session{Name: "Morning", Start: 9*60 + 15, End: 12 * 60}
	// Hong Kong afternoon session: 13:00 - 16:30 (16:15 for auction)
	HKSessionAfternoon = Session{Name: "Afternoon", Start: 13 * 60, End: 16*60 + 30}
	// Hong Kong pre-market: 9:00 - 9:15 (only for securities)
	HKSessionPreMarket = Session{Name: "PreMarket", Start: 9 * 60, End: 9*60 + 15}
	// Hong Kong after-hours: 16:30 - 17:15 (for securities)
	HKSessionAfterHours = Session{Name: "AfterHours", Start: 16*60 + 30, End: 17*60 + 15}
)

// Standard US market sessions (EST/EDT)
var (
	USSessionPreMarket = Session{Name: "PreMarket", Start: 4 * 60, End: 9*60 + 30}
	USSessionCore = Session{Name: "Core", Start: 9*60 + 30, End: 16 * 60}
	USSessionAfter = Session{Name: "After", Start: 16 * 60, End: 20 * 60}
)

// Standard CN market sessions
var (
 CNSessionMorning = Session{Name: "Morning", Start: 9*60 + 30, End: 11 * 60 + 30}
 CNSessionAfternoon = Session{Name: "Afternoon", Start: 13 * 60, End: 15 * 60}
)

// IsHKOpen checks if the HK market is currently open for trading
func IsHKOpen(t time.Time) bool {
	loc := time.FixedZone("HKST", 8*3600) // Hong Kong is UTC+8
	t = t.In(loc)

	// Weekend check
	weekday := t.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	// Get minutes from midnight in HK time
	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	// Check morning session (9:15 - 12:00)
	if mins >= HKSessionMorning.Start && mins < HKSessionMorning.End {
		return true
	}

	// Check afternoon session (13:00 - 16:30)
	if mins >= HKSessionAfternoon.Start && mins < HKSessionAfternoon.End {
		return true
	}

	return false
}

// IsUSOpen checks if the US market is currently open for trading
// Note: US market operates in EST/EDT (UTC-5/UTC-4)
func IsUSOpen(t time.Time) bool {
	// Determine if DST is in effect (2nd Sunday in March - 1st Sunday in November)
	year := t.Year()
	dstStart := time.Date(year, time.March, 8, 2, 0, 0, 0, time.UTC)
	dstStart = dstStart.AddDate(0, 0, -int(dstStart.Weekday()))

	dstEnd := time.Date(year, time.November, 1, 2, 0, 0, 0, time.UTC)
	dstEnd = dstEnd.AddDate(0, 0, 7-int(dstEnd.Weekday()))

	locName := "EST"
	offset := -5 * 3600
	if t.After(dstStart) && t.Before(dstEnd) {
		locName = "EDT"
		offset = -4 * 3600
	}

	loc := time.FixedZone(locName, offset)
	t = t.In(loc)

	// Weekend check
	weekday := t.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	if mins >= USSessionCore.Start && mins < USSessionCore.End {
		return true
	}

	return false
}

// IsCNOpen checks if the China market (SH/SZ) is open
func IsCNOpen(t time.Time) bool {
	loc := time.FixedZone("CST", 8*3600) // China is UTC+8
	t = t.In(loc)

	weekday := t.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	// Morning session 9:30-11:30
	if mins >= CNSessionMorning.Start && mins < CNSessionMorning.End {
		return true
	}

	// Afternoon session 13:00-15:00
	if mins >= CNSessionAfternoon.Start && mins < CNSessionAfternoon.End {
		return true
	}

	return false
}

// IsOpen checks if a specific market is open
func IsOpen(m Market, t time.Time) bool {
	switch m {
	case MarketHK:
		return IsHKOpen(t)
	case MarketUS:
		return IsUSOpen(t)
	case MarketCN:
		return IsCNOpen(t)
	default:
		return false
	}
}

// UntilClose returns the duration until the market closes
// Returns 0 if market is closed
func UntilClose(m Market, t time.Time) time.Duration {
	switch m {
	case MarketHK:
		return untilHKClose(t)
	case MarketUS:
		return untilUSClose(t)
	case MarketCN:
		return untilCNClose(t)
	default:
		return 0
	}
}

func untilHKClose(t time.Time) time.Duration {
	loc := time.FixedZone("HKST", 8*3600)
	t = t.In(loc)

	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	if mins < HKSessionMorning.Start {
		return (HKSessionMorning.Start - mins) * time.Minute
	}
	if mins < HKSessionMorning.End {
		return (HKSessionMorning.End - mins) * time.Minute
	}
	if mins < HKSessionAfternoon.Start {
		return (HKSessionAfternoon.Start - mins) * time.Minute
	}
	if mins < HKSessionAfternoon.End {
		return (HKSessionAfternoon.End - mins) * time.Minute
	}

	return 0
}

func untilUSClose(t time.Time) time.Duration {
	year := t.Year()
	dstStart := time.Date(year, time.March, 8, 2, 0, 0, 0, time.UTC)
	dstStart = dstStart.AddDate(0, 0, -int(dstStart.Weekday()))

	dstEnd := time.Date(year, time.November, 1, 2, 0, 0, 0, time.UTC)
	dstEnd = dstEnd.AddDate(0, 0, 7-int(dstEnd.Weekday()))

	locName := "EST"
	offset := -5 * 3600
	if t.After(dstStart) && t.Before(dstEnd) {
		locName = "EDT"
		offset = -4 * 3600
	}

	loc := time.FixedZone(locName, offset)
	t = t.In(loc)

	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	if mins < USSessionCore.Start {
		return (USSessionCore.Start - mins) * time.Minute
	}
	if mins < USSessionCore.End {
		return (USSessionCore.End - mins) * time.Minute
	}

	return 0
}

func untilCNClose(t time.Time) time.Duration {
	loc := time.FixedZone("CST", 8*3600)
	t = t.In(loc)

	mins := time.Duration(t.Hour())*60 + time.Duration(t.Minute())

	if mins < CNSessionMorning.Start {
		return (CNSessionMorning.Start - mins) * time.Minute
	}
	if mins < CNSessionMorning.End {
		return (CNSessionMorning.End - mins) * time.Minute
	}
	if mins < CNSessionAfternoon.Start {
		return (CNSessionAfternoon.Start - mins) * time.Minute
	}
	if mins < CNSessionAfternoon.End {
		return (CNSessionAfternoon.End - mins) * time.Minute
	}

	return 0
}

// NextOpen returns the time when the market will next open
func NextOpen(m Market, t time.Time) time.Duration {
	// Simple implementation - returns next day morning session start
	loc := time.FixedZone("HKST", 8*3600)
	next := time.Date(t.Year(), t.Month(), t.Day(), 9, 15, 0, 0, loc)
	
	if t.After(next) {
		next = next.AddDate(1, 0, 0)
	}
	
	// Skip weekends
	for next.Weekday() == time.Saturday || next.Weekday() == time.Sunday {
		next = next.AddDate(1, 0, 0)
	}
	
	return next.Sub(t)
}

// MarketHours returns formatted market hours for display
func MarketHours(m Market) string {
	switch m {
	case MarketHK:
		return fmt.Sprintf("Morning: 9:15-12:00, Afternoon: 13:00-16:30 (HKST)")
	case MarketUS:
		return fmt.Sprintf("Core: 9:30-16:00 (EST/EDT)")
	case MarketCN:
		return fmt.Sprintf("Morning: 9:30-11:30, Afternoon: 13:00-15:00 (CST)")
	default:
		return "Unknown"
	}
}