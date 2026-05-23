package formatter

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// NewTOMLFormatter returns a Formatter that renders cron execution times as TOML.
func NewTOMLFormatter() Formatter {
	return &tomlFormatter{}
}

type tomlFormatter struct{}

// Render formats the cron expression and its next execution times as a TOML document.
func (f *tomlFormatter) Render(expression string, times []time.Time, timezone string) (string, error) {
	if timezone == "" {
		timezone = "UTC"
	}

	var buf bytes.Buffer

	buf.WriteString("# cronscope output\n")
	buf.WriteString(fmt.Sprintf("expression = %q\n", expression))
	buf.WriteString(fmt.Sprintf("timezone = %q\n", timezone))
	buf.WriteString(fmt.Sprintf("count = %d\n", len(times)))

	if len(times) == 0 {
		buf.WriteString("\n[[executions]]\n")
		buf.WriteString("# no executions\n")
		return buf.String(), nil
	}

	buf.WriteString("\n")
	for i, t := range times {
		buf.WriteString("[[executions]]\n")
		buf.WriteString(fmt.Sprintf("index = %d\n", i+1))
		buf.WriteString(fmt.Sprintf("time = %q\n", t.Format(time.RFC3339)))
		buf.WriteString(fmt.Sprintf("unix = %d\n", t.Unix()))
		if i < len(times)-1 {
			buf.WriteString("\n")
		}
	}

	result := buf.String()
	result = strings.TrimRight(result, "\n") + "\n"
	return result, nil
}
