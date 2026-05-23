// Package formatter provides output formatters for cronscope.
//
// It supports multiple output formats for rendering cron schedule results:
//
//   - TableFormatter: renders results as a human-readable terminal table
//   - JSONFormatter: renders results as structured JSON for scripting or piping
//
// All formatters accept a cron expression string and a slice of time.Time values
// representing upcoming execution times.
package formatter
