package util

import "time"

const (
	FutuDateFormat     = "2006-01-02"
	FutuDateTimeFormat = "2006-01-02 15:04:05"
)

// TodayDateStr returns today's date in Futu date format (2006-01-02).
func TodayDateStr() string {
	return time.Now().Format(FutuDateFormat)
}

// NowDateTimeStr returns the current time in Futu datetime format (2006-01-02 15:04:05).
func NowDateTimeStr() string {
	return time.Now().Format(FutuDateTimeFormat)
}

// ParseFutuDate parses a Futu date string (2006-01-02) into time.Time.
func ParseFutuDate(s string) (time.Time, error) {
	return time.Parse(FutuDateFormat, s)
}

// ParseFutuDateTime parses a Futu datetime string (2006-01-02 15:04:05) into time.Time.
func ParseFutuDateTime(s string) (time.Time, error) {
	return time.Parse(FutuDateTimeFormat, s)
}

// FormatFutuDate formats a time.Time as a Futu date string (2006-01-02).
func FormatFutuDate(t time.Time) string {
	return t.Format(FutuDateFormat)
}

// FormatFutuDateTime formats a time.Time as a Futu datetime string (2006-01-02 15:04:05).
func FormatFutuDateTime(t time.Time) string {
	return t.Format(FutuDateTimeFormat)
}

// IsValidFutuDate returns true if the string is a valid Futu date (2006-01-02).
func IsValidFutuDate(s string) bool {
	_, err := time.Parse(FutuDateFormat, s)
	return err == nil
}

// IsValidFutuDateTime returns true if the string is a valid Futu datetime (2006-01-02 15:04:05).
func IsValidFutuDateTime(s string) bool {
	_, err := time.Parse(FutuDateTimeFormat, s)
	return err == nil
}
