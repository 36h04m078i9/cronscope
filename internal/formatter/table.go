package formatter

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// TableFormatter renders cron execution times as a formatted table.
type TableFormatter struct {
	Writer     io.Writer
	TimeFormat string
}

// NewTableFormatter creates a TableFormatter writing to w.
func NewTableFormatter(w io.Writer) *TableFormatter {
	return &TableFormatter{
		Writer:     w,
		TimeFormat: "2006-01-02 15:04:05 MST",
	}
}

// Render prints a table of upcoming execution times for the given expression.
func (f *TableFormatter) Render(expression string, times []time.Time) error {
	const colWidth = 30
	separator := strings.Repeat("-", colWidth+4)

	fmt.Fprintf(f.Writer, "\nCron Expression: %s\n", expression)
	fmt.Fprintln(f.Writer, separator)
	fmt.Fprintf(f.Writer, "  %-*s\n", colWidth, "Next Execution Times")
	fmt.Fprintln(f.Writer, separator)

	for i, t := range times {
		formatted := t.Format(f.TimeFormat)
		fmt.Fprintf(f.Writer, "  %2d. %-*s\n", i+1, colWidth, formatted)
	}

	fmt.Fprintln(f.Writer, separator)
	fmt.Fprintf(f.Writer, "  Total: %d execution(s) shown\n\n", len(times))
	return nil
}
