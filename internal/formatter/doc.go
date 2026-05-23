// Package formatter provides multiple output formatters for rendering
// cron expression schedules and their next execution times.
//
// Available formatters:
//
//   - TableFormatter  – ASCII table output (default)
//   - PlainFormatter  – plain text, one time per line
//   - JSONFormatter   – machine-readable JSON
//   - ColorFormatter  – ANSI-colored terminal output
//   - HumanizeFormatter – human-friendly relative durations
//   - ICalFormatter   – iCalendar (.ics) format
//   - CSVFormatter    – comma-separated values
//
// Each formatter implements a Render(expression string, times []time.Time) (string, error)
// method and can be constructed with a timezone string. Passing an empty
// timezone string defaults to UTC.
package formatter
