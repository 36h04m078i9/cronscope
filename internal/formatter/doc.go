// Package formatter provides output renderers for cron execution schedules.
//
// Supported formats:
//   - TableFormatter: renders results as an ASCII table
//   - JSONFormatter:  renders results as a JSON document
//   - PlainFormatter: renders results as human-readable plain text
//
// Each formatter accepts a cron expression string and a timezone name,
// and exposes a Render([]time.Time) string method.
package formatter
