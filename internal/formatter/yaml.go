package formatter

import (
	"fmt"
	"strings"
	"time"
)

// NewYAMLFormatter returns a Formatter that renders cron schedule output as YAML.
func NewYAMLFormatter(expression, timezone string, times []time.Time) Formatter {
	return &yamlFormatter{
		expression: expression,
		timezone:   timezone,
		times:      times,
	}
}

type yamlFormatter struct {
	expression string
	timezone   string
	times      []time.Time
}

func (f *yamlFormatter) Render() string {
	tz := f.timezone
	if tz == "" {
		tz = "UTC"
	}

	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("expression: %q\n", f.expression))
	sb.WriteString(fmt.Sprintf("timezone: %q\n", tz))
	sb.WriteString("executions:\n")

	if len(f.times) == 0 {
		sb.WriteString("  []\n")
		return sb.String()
	}

	for i, t := range f.times {
		sb.WriteString(fmt.Sprintf("  - index: %d\n", i+1))
		sb.WriteString(fmt.Sprintf("    time: %q\n", t.UTC().Format(time.RFC3339)))
	}

	return sb.String()
}
