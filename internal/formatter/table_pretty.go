package formatter

import (
	"fmt"
	"strings"
	"time"
)

// PrettyTableFormatter renders cron execution times as a bordered ASCII table
// with aligned columns and a summary footer.
type PrettyTableFormatter struct {
	timezone string
}

// NewPrettyTableFormatter returns a PrettyTableFormatter with the given timezone.
func NewPrettyTableFormatter(timezone string) *PrettyTableFormatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &PrettyTableFormatter{timezone: tz}
}

// Render formats the expression and times into a pretty bordered table string.
func (f *PrettyTableFormatter) Render(expression string, times []time.Time) (string, error) {
	const dateWidth = 30
	const indexWidth = 5

	separator := fmt.Sprintf("+%s+%s+",
		strings.Repeat("-", indexWidth+2),
		strings.Repeat("-", dateWidth+2),
	)

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression : %s\n", expression))
	sb.WriteString(fmt.Sprintf("Timezone   : %s\n", f.timezone))
	sb.WriteString(separator + "\n")
	sb.WriteString(fmt.Sprintf("| %-*s | %-*s |\n", indexWidth, "#", dateWidth, "Scheduled Time"))
	sb.WriteString(separator + "\n")

	if len(times) == 0 {
		sb.WriteString(fmt.Sprintf("| %-*s |\n", indexWidth+dateWidth+3, "No upcoming executions"))
	} else {
		for i, t := range times {
			formatted := t.Format("2006-01-02 15:04:05 MST")
			sb.WriteString(fmt.Sprintf("| %-*d | %-*s |\n", indexWidth, i+1, dateWidth, formatted))
		}
	}

	sb.WriteString(separator + "\n")
	sb.WriteString(fmt.Sprintf("Total: %d execution(s)\n", len(times)))

	return sb.String(), nil
}
