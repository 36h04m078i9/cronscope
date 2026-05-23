package formatter

import (
	"fmt"
	"strings"
	"time"
)

// markdownFormatter renders cron schedule output as a Markdown table.
type markdownFormatter struct {
	timezone string
}

// NewMarkdownFormatter returns a Formatter that produces Markdown table output.
func NewMarkdownFormatter(timezone string) Formatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &markdownFormatter{timezone: tz}
}

// Render formats the cron expression and its next execution times as a
// Markdown table suitable for embedding in documentation or GitHub issues.
func (m *markdownFormatter) Render(expression string, times []time.Time) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("## Cron Schedule: `%s`\n", expression))
	sb.WriteString(fmt.Sprintf("**Timezone:** %s\n\n", m.timezone))

	if len(times) == 0 {
		sb.WriteString("_No upcoming executions._\n")
		return sb.String()
	}

	sb.WriteString("| # | Date | Time | Weekday |\n")
	sb.WriteString("|---|------|------|---------|\n")

	for i, t := range times {
		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n",
			i+1,
			t.Format("2006-01-02"),
			t.Format("15:04:05"),
			t.Weekday().String(),
		))
	}

	return sb.String()
}
