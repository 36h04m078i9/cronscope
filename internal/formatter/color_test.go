package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestColorFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewColorFormatter("0 9 * * 1-5", "UTC")
	times := []time.Time{time.Now().UTC()}
	output := f.Render(times)
	if !strings.Contains(output, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got: %s", output)
	}
}

func TestColorFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewColorFormatter("*/5 * * * *", "America/New_York")
	times := []time.Time{time.Now().UTC()}
	output := f.Render(times)
	if !strings.Contains(output, "America/New_York") {
		t.Errorf("expected output to contain timezone, got: %s", output)
	}
}

func TestColorFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewColorFormatter("*/5 * * * *", "")
	times := []time.Time{time.Now().UTC()}
	output := f.Render(times)
	if !strings.Contains(output, "UTC") {
		t.Errorf("expected default timezone UTC in output, got: %s", output)
	}
}

func TestColorFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewColorFormatter("0 0 31 2 *", "UTC")
	output := f.Render([]time.Time{})
	if !strings.Contains(output, "No upcoming executions") {
		t.Errorf("expected no executions message, got: %s", output)
	}
}

func TestColorFormatter_Render_IndexedRows(t *testing.T) {
	f := NewColorFormatter("0 * * * *", "UTC")
	now := time.Now().UTC()
	times := []time.Time{now, now.Add(time.Hour), now.Add(2 * time.Hour)}
	output := f.Render(times)
	for _, idx := range []string{"1.", "2.", "3."} {
		if !strings.Contains(output, idx) {
			t.Errorf("expected index %q in output, got: %s", idx, output)
		}
	}
}

func TestColorFormatter_Render_ContainsANSI(t *testing.T) {
	f := NewColorFormatter("0 * * * *", "UTC")
	times := []time.Time{time.Now().UTC()}
	output := f.Render(times)
	if !strings.Contains(output, "\033[") {
		t.Errorf("expected ANSI escape codes in colorized output")
	}
}
