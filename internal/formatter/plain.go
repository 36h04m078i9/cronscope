package formatter

import (
	"fmt"
	"strings"
	"time"
)

// PlainFormatter renders cron execution times as plain text output.
type PlainFormatter struct {
	expression string
	timezone   string
}

// NewPlainFormatter creates a new PlainFormatter for the given cron expression and timezone.
func NewPlainFormatter(expression, timezone string) *PlainFormatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &PlainFormatter{
		expression: expression,
		timezone:   tz,
	}
}

// Render formats the list of times as a plain-text string.
func (f *PlainFormatter) Render(times []time.Time) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Cron Expression : %s\n", f.expression))
	sb.WriteString(fmt.Sprintf("Timezone        : %s\n", f.timezone))
	sb.WriteString(fmt.Sprintf("Next executions : %d\n", len(times)))
	sb.WriteString(strings.Repeat("-", 40) + "\n")

	if len(times) == 0 {
		sb.WriteString("No upcoming executions found.\n")
		return sb.String()
	}

	for i, t := range times {
		sb.WriteString(fmt.Sprintf(" %2d. %s\n", i+1, t.Format(time.RFC3339)))
	}

	return sb.String()
}
