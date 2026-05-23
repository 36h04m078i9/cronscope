// Package formatter provides multiple output format renderers for cron execution times.
//
// Each formatter implements the Formatter interface, which accepts a cron expression
// string and a slice of scheduled time.Time values, returning a formatted string
// suitable for terminal output or file export.
//
// Available formatters:
//
//   - Table    — aligned ASCII table (via NewTableFormatter)
//   - Plain    — simple line-by-line text (via NewPlainFormatter)
//   - JSON     — machine-readable JSON (via NewJSONFormatter)
//   - Color    — ANSI-colored terminal output (via NewColorFormatter)
//   - Humanize — human-friendly relative durations (via NewHumanizeFormatter)
//   - ICal     — iCalendar RFC 5545 VEVENT export (via NewICalFormatter)
package formatter

// Formatter is the interface implemented by all output renderers.
type Formatter interface {
	// Render formats the given cron expression and scheduled times into a string.
	Render(expression string, times []time.Time) string
}
