package formatter

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

// templateData holds the data passed to user-defined templates.
type templateData struct {
	Expression string
	Timezone   string
	Times      []templateEntry
}

// templateEntry represents a single scheduled time entry.
type templateEntry struct {
	Index     int
	Formatted string
	Unix      int64
	RFC3339   string
}

// TemplateFormatter renders cron schedule output using a Go text/template string.
type TemplateFormatter struct {
	tmpl *template.Template
}

// NewTemplateFormatter creates a TemplateFormatter from the provided template string.
// Returns an error if the template fails to parse.
func NewTemplateFormatter(tmplStr string) (*TemplateFormatter, error) {
	tmpl, err := template.New("cronscope").Parse(tmplStr)
	if err != nil {
		return nil, fmt.Errorf("template parse error: %w", err)
	}
	return &TemplateFormatter{tmpl: tmpl}, nil
}

// Render executes the template with the provided expression, timezone, and times.
func (f *TemplateFormatter) Render(expression, timezone string, times []time.Time) (string, error) {
	if timezone == "" {
		timezone = "UTC"
	}

	entries := make([]templateEntry, len(times))
	for i, t := range times {
		entries[i] = templateEntry{
			Index:     i + 1,
			Formatted: t.Format("2006-01-02 15:04:05 MST"),
			Unix:      t.Unix(),
			RFC3339:   t.Format(time.RFC3339),
		}
	}

	data := templateData{
		Expression: expression,
		Timezone:   timezone,
		Times:      entries,
	}

	var buf bytes.Buffer
	if err := f.tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("template execute error: %w", err)
	}
	return buf.String(), nil
}
