package formatter

import (
	"encoding/json"
	"time"
)

// JSONFormatter formats cron execution times as JSON output.
type JSONFormatter struct {
	expression string
	timezone   string
}

// JSONOutput represents the JSON structure for cron schedule output.
type JSONOutput struct {
	Expression string   `json:"expression"`
	Timezone   string   `json:"timezone"`
	NextTimes  []string `json:"next_times"`
}

// NewJSONFormatter creates a new JSONFormatter for the given cron expression and timezone.
func NewJSONFormatter(expression, timezone string) *JSONFormatter {
	if timezone == "" {
		timezone = "UTC"
	}
	return &JSONFormatter{
		expression: expression,
		timezone:   timezone,
	}
}

// Render serializes the cron expression and its next execution times to JSON.
func (f *JSONFormatter) Render(times []time.Time) (string, error) {
	formatted := make([]string, 0, len(times))
	for _, t := range times {
		formatted = append(formatted, t.Format(time.RFC3339))
	}

	out := JSONOutput{
		Expression: f.expression,
		Timezone:   f.timezone,
		NextTimes:  formatted,
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
