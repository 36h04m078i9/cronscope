package formatter

import (
	"fmt"
	"strings"
	"time"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorGray   = "\033[90m"
)

// ColorFormatter renders cron execution times with ANSI color highlighting.
type ColorFormatter struct {
	expression string
	timezone   string
}

// NewColorFormatter creates a new ColorFormatter.
func NewColorFormatter(expression, timezone string) *ColorFormatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &ColorFormatter{expression: expression, timezone: tz}
}

// Render outputs colorized cron schedule information.
func (c *ColorFormatter) Render(times []time.Time) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s%sCron Expression:%s %s%s%s\n",
		colorBold, colorCyan, colorReset, colorYellow, c.expression, colorReset))
	sb.WriteString(fmt.Sprintf("%sTimezone:%s %s%s%s\n",
		colorGray, colorReset, colorGreen, c.timezone, colorReset))
	sb.WriteString(fmt.Sprintf("%s%s", colorGray, strings.Repeat("-", 40)+colorReset+"\n"))

	if len(times) == 0 {
		sb.WriteString(fmt.Sprintf("%sNo upcoming executions found.%s\n", colorGray, colorReset))
		return sb.String()
	}

	for i, t := range times {
		sb.WriteString(fmt.Sprintf("%s%3d.%s %s%s%s\n",
			colorGray, i+1, colorReset,
			colorGreen, t.Format("2006-01-02 15:04:05 MST"), colorReset))
	}

	return sb.String()
}
