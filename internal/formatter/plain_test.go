package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestPlainFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewPlainFormatter("0 9 * * 1-5", "UTC")
	times := []time.Time{time.Now().UTC()}
	output := f.Render(times)

	if !strings.Contains(output, "0 9 * * 1-5") {
		t.Errorf("expected output to contain cron expression, got:\n%s", output)
	}
}

func TestPlainFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewPlainFormatter("*/5 * * * *", "America/New_York")
	output := f.Render([]time.Time{})

	if !strings.Contains(output, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", output)
	}
}

func TestPlainFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewPlainFormatter("*/5 * * * *", "")
	output := f.Render([]time.Time{})

	if !strings.Contains(output, "UTC") {
		t.Errorf("expected default timezone UTC in output, got:\n%s", output)
	}
}

func TestPlainFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewPlainFormatter("0 0 * * *", "UTC")
	output := f.Render([]time.Time{})

	if !strings.Contains(output, "No upcoming executions found.") {
		t.Errorf("expected empty message, got:\n%s", output)
	}
}

func TestPlainFormatter_Render_IndexedRows(t *testing.T) {
	f := NewPlainFormatter("0 12 * * *", "UTC")
	times := []time.Time{
		time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 2, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 3, 12, 0, 0, 0, time.UTC),
	}
	output := f.Render(times)

	for _, want := range []string{" 1.", " 2.", " 3."} {
		if !strings.Contains(output, want) {
			t.Errorf("expected output to contain %q, got:\n%s", want, output)
		}
	}
}

func TestPlainFormatter_Render_ContainsTimes(t *testing.T) {
	f := NewPlainFormatter("0 6 * * *", "UTC")
	ts := time.Date(2024, 3, 15, 6, 0, 0, 0, time.UTC)
	output := f.Render([]time.Time{ts})

	if !strings.Contains(output, "2024-03-15T06:00:00Z") {
		t.Errorf("expected RFC3339 time in output, got:\n%s", output)
	}
}
