// Package formatter provides output renderers for cron execution schedules.
//
// Supported formats:
//   - TableFormatter: renders results as an ASCII table
//   - JSONFormatter:  renders results as a JSON document
//   - PlainFormatter: renders results as human-readable plain text
//
// Each formatter accepts a cron expression string and a timezone name,
// and exposes a Render([]time.Time) string method.
//
// Usage example:
//
//	f, err := formatter.New("table", "* * * * *", "UTC")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(f.Render(times))
//
// The format name passed to New is case-insensitive. Valid values are
// "table", "json", and "plain".
package formatter
