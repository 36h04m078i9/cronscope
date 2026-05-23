// Package formatter provides multiple output format renderers for cron schedule results.
//
// Each formatter implements the Formatter interface and accepts a cron expression
// string and a slice of scheduled times, returning a formatted string representation.
//
// Available formatters:
//
//   - Table    — ASCII table output (default)
//   - Plain    — Simple line-by-line text output
//   - JSON     — Machine-readable JSON output
//   - Color    — ANSI-colored terminal output
//   - Humanize — Human-friendly relative time output
//   - ICal     — iCalendar (.ics) format for calendar import
//   - CSV      — Comma-separated values for spreadsheet use
//   - Markdown — GitHub-flavored Markdown table
//   - XML      — XML document output for interoperability
//
// Usage:
//
//	f := formatter.NewTableFormatter("UTC")
//	output, err := f.Render("0 * * * *", times)
//
// The Formatter interface:
//
//	type Formatter interface {
//		Render(expression string, times []time.Time) (string, error)
//	}
package formatter

import "time"

// Formatter is the common interface implemented by all output formatters.
type Formatter interface {
	Render(expression string, times []time.Time) (string, error)
}
