// Package formatter provides multiple output renderers for cron execution schedules.
//
// Each formatter implements a Render(expression string, times []time.Time) (string, error)
// method and can be selected via the --format CLI flag.
//
// Available formatters:
//
//   - table        : compact tabwriter-aligned table
//   - pretty-table : bordered ASCII table with summary footer
//   - plain        : simple numbered list
//   - text         : plain text with labels
//   - color        : ANSI color-highlighted output
//   - json         : JSON object with expression, timezone, and times array
//   - yaml         : YAML document
//   - toml         : TOML document
//   - csv          : comma-separated values with header row
//   - markdown     : GitHub-flavored Markdown table
//   - xml          : XML document with schedule entries
//   - ical         : iCalendar (RFC 5545) VCALENDAR with VEVENT entries
//   - humanize     : human-readable relative durations (e.g. "in 3 hours")
//   - template     : user-supplied Go text/template string
//
// All formatters default to UTC when no timezone is provided.
package formatter
