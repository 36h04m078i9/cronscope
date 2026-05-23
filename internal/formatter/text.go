package formatter

import (
	"fmt"
	"strings"
	"time"
)

// TextFormatter renders cron schedule output as a simple numbered text list.
type TextFormatter struct {
	timezone string
}

// NewTextFormatter returns a TextFormatter with the given timezone label.
func NewTextFormatter(timezone string) *TextFormatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &TextFormatter{timezone: tz}
}

// Render formats the cron expression and its next execution times as plain
// numbered text, suitable for piping or minimal terminal output.
func (f *TextFormatter) Render(expression string, times []time.Time) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression : %s\n", expression))
	sb.WriteString(fmt.Sprintf("Timezone   : %s\n", f.timezone))
	sb.WriteString(strings.Repeat("-", 40) + "\n")

	if len(times) == 0 {
		sb.WriteString("No upcoming executions found.\n")
		return sb.String(), nil
	}

	for i, t := range times {
		sb.WriteString(fmt.Sprintf("%3d. %s\n", i+1, t.Format("2006-01-02 15:04:05 MST")))
	}

	return sb.String(), nil
}
