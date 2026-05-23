package parser

import (
	"fmt"
	"time"
)

// CommonTimezones lists frequently used timezone identifiers for reference.
var CommonTimezones = []string{
	"UTC",
	"America/New_York",
	"America/Chicago",
	"America/Denver",
	"America/Los_Angeles",
	"Europe/London",
	"Europe/Paris",
	"Europe/Berlin",
	"Asia/Tokyo",
	"Asia/Shanghai",
	"Asia/Kolkata",
	"Australia/Sydney",
}

// LoadTimezone loads a time.Location from a timezone string.
// Accepts IANA timezone names (e.g. "America/New_York") or UTC offsets (e.g. "UTC").
func LoadTimezone(tz string) (*time.Location, error) {
	if tz == "" || tz == "UTC" {
		return time.UTC, nil
	}

	if tz == "Local" {
		return time.Local, nil
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("unknown timezone %q: %w", tz, err)
	}

	return loc, nil
}

// FormatWithTimezone formats a time value using the given layout in the specified location.
func FormatWithTimezone(t time.Time, layout string, loc *time.Location) string {
	if loc != nil {
		t = t.In(loc)
	}
	return t.Format(layout)
}

// DefaultTimeLayout is the standard display format for cronscope output.
const DefaultTimeLayout = "2006-01-02 15:04:05 MST"
