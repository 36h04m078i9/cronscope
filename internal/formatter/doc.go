// Package formatter provides multiple output formatters for cron schedule data.
//
// Available formatters:
//
//   - TableFormatter: renders results as an aligned ASCII table.
//   - PlainFormatter: renders results as plain human-readable text.
//   - JSONFormatter:  renders results as structured JSON output.
//   - ColorFormatter: renders results with ANSI color highlighting for
//     interactive terminal use.
//
// All formatters accept a cron expression string, a timezone string, and a
// slice of time.Time values representing upcoming execution times. They each
// expose a Render method that returns the formatted string.
//
// Example usage:
//
//	f := formatter.NewColorFormatter("0 9 * * 1-5", "America/Chicago")
//	fmt.Print(f.Render(times))
package formatter
