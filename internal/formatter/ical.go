package formatter

import (
	"fmt"
	"strings"
	"time"
)

// iCalFormatter formats cron execution times as iCalendar (RFC 5545) VEVENT entries.
type iCalFormatter struct {
	timezone string
}

// NewICalFormatter returns a Formatter that renders times as iCal VEVENT blocks.
func NewICalFormatter(timezone string) Formatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &iCalFormatter{timezone: tz}
}

// Render produces an iCalendar document string for the given expression and times.
func (f *iCalFormatter) Render(expression string, times []time.Time) string {
	var sb strings.Builder

	sb.WriteString("BEGIN:VCALENDAR\r\n")
	sb.WriteString("VERSION:2.0\r\n")
	sb.WriteString("PRODID:-//cronscope//EN\r\n")
	sb.WriteString(fmt.Sprintf("X-WR-CALNAME:cronscope: %s\r\n", expression))
	sb.WriteString(fmt.Sprintf("X-WR-TIMEZONE:%s\r\n", f.timezone))

	for i, t := range times {
		utc := t.UTC()
		stamp := utc.Format("20060102T150405Z")
		sb.WriteString("BEGIN:VEVENT\r\n")
		sb.WriteString(fmt.Sprintf("UID:cronscope-%d-%s@cronscope\r\n", i+1, stamp))
		sb.WriteString(fmt.Sprintf("DTSTAMP:%s\r\n", stamp))
		sb.WriteString(fmt.Sprintf("DTSTART:%s\r\n", stamp))
		sb.WriteString(fmt.Sprintf("DTEND:%s\r\n", stamp))
		sb.WriteString(fmt.Sprintf("SUMMARY:cron: %s\r\n", expression))
		sb.WriteString(fmt.Sprintf("DESCRIPTION:Execution #%d of cron expression %s\r\n", i+1, expression))
		sb.WriteString("END:VEVENT\r\n")
	}

	sb.WriteString("END:VCALENDAR\r\n")
	return sb.String()
}
