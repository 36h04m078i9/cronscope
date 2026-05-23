package formatter

import (
	"fmt"
	"strings"
	"time"
)

// HumanizeFormatter renders cron execution times with human-readable relative
// durations alongside absolute timestamps.
type HumanizeFormatter struct {
	expression string
	timezone   string
	now        func() time.Time
}

// NewHumanizeFormatter returns a Formatter that includes relative time labels
// such as "in 5 minutes" or "in 2 hours" next to each scheduled time.
func NewHumanizeFormatter(expression, timezone string) Formatter {
	return &HumanizeFormatter{
		expression: expression,
		timezone:   timezone,
		now:        time.Now,
	}
}

// Render formats the list of times with relative durations.
func (h *HumanizeFormatter) Render(times []time.Time) string {
	tz := h.timezone
	if tz == "" {
		tz = "UTC"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Expression : %s\n", h.expression))
	sb.WriteString(fmt.Sprintf("Timezone   : %s\n", tz))
	sb.WriteString(strings.Repeat("-", 60) + "\n")

	if len(times) == 0 {
		sb.WriteString("No upcoming executions.\n")
		return sb.String()
	}

	now := h.now()
	for i, t := range times {
		rel := humanDuration(t.Sub(now))
		sb.WriteString(fmt.Sprintf("  %2d. %s  (%s)\n", i+1, t.Format("2006-01-02 15:04:05 MST"), rel))
	}

	return sb.String()
}

// humanDuration converts a duration into a short human-readable string.
func humanDuration(d time.Duration) string {
	if d < 0 {
		return "in the past"
	}
	if d < time.Minute {
		secs := int(d.Seconds())
		return fmt.Sprintf("in %d second%s", secs, plural(secs))
	}
	if d < time.Hour {
		mins := int(d.Minutes())
		return fmt.Sprintf("in %d minute%s", mins, plural(mins))
	}
	if d < 24*time.Hour {
		hrs := int(d.Hours())
		return fmt.Sprintf("in %d hour%s", hrs, plural(hrs))
	}
	days := int(d.Hours() / 24)
	return fmt.Sprintf("in %d day%s", days, plural(days))
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
