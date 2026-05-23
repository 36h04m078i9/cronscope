package formatter

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"
)

// CSVFormatter renders cron execution times as CSV output.
type CSVFormatter struct {
	timezone string
}

// NewCSVFormatter creates a new CSVFormatter with the given timezone.
func NewCSVFormatter(timezone string) *CSVFormatter {
	if timezone == "" {
		timezone = "UTC"
	}
	return &CSVFormatter{timezone: timezone}
}

// Render formats the cron expression and its next execution times as CSV.
func (f *CSVFormatter) Render(expression string, times []time.Time) (string, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	// Write header
	if err := w.Write([]string{"index", "expression", "timezone", "datetime"}); err != nil {
		return "", fmt.Errorf("csv: write header: %w", err)
	}

	if len(times) == 0 {
		w.Flush()
		if err := w.Error(); err != nil {
			return "", fmt.Errorf("csv: flush: %w", err)
		}
		return buf.String(), nil
	}

	for i, t := range times {
		row := []string{
			fmt.Sprintf("%d", i+1),
			expression,
			f.timezone,
			t.Format(time.RFC3339),
		}
		if err := w.Write(row); err != nil {
			return "", fmt.Errorf("csv: write row %d: %w", i+1, err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return "", fmt.Errorf("csv: flush: %w", err)
	}

	return buf.String(), nil
}
