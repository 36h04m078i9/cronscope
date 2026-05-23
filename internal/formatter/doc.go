// Package formatter provides multiple output renderers for cron schedule data.
//
// Each formatter implements the Formatter interface and accepts a cron
// expression along with a slice of upcoming execution times, returning
// a string representation suitable for a specific output format.
//
// Available formatters:
//
//   - Table    – aligned terminal table (via tabwriter)
//   - Plain    – simple numbered list
//   - JSON     – machine-readable JSON array
//   - Color    – ANSI-colored terminal output
//   - Humanize – human-friendly relative durations
//   - ICal     – RFC 5545 iCalendar (.ics) format
//   - CSV      – comma-separated values
//   - Markdown – GitHub-flavoured Markdown table
//
// Usage:
//
//	f := formatter.NewMarkdownFormatter("America/Chicago")
//	fmt.Println(f.Render("0 9 * * 1-5", times))
package formatter

// Formatter is the common interface implemented by all output renderers.
type Formatter interface {
	Render(expression string, times []time.Time) string
}
