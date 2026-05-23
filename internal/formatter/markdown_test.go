package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestMarkdownFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewMarkdownFormatter("UTC")
	out := f.Render("0 9 * * 1-5", []time.Time{})
	if !strings.Contains(out, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestMarkdownFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewMarkdownFormatter("America/New_York")
	out := f.Render("*/5 * * * *", []time.Time{})
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", out)
	}
}

func TestMarkdownFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewMarkdownFormatter("")
	out := f.Render("*/5 * * * *", []time.Time{})
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected default timezone UTC in output, got:\n%s", out)
	}
}

func TestMarkdownFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewMarkdownFormatter("UTC")
	out := f.Render("0 0 * * *", []time.Time{})
	if !strings.Contains(out, "No upcoming executions") {
		t.Errorf("expected empty message, got:\n%s", out)
	}
}

func TestMarkdownFormatter_Render_TableHeader(t *testing.T) {
	f := NewMarkdownFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 4, 9, 0, 0, 0, time.UTC),
	}
	out := f.Render("0 9 * * *", times)
	for _, header := range []string{"#", "Date", "Time", "Weekday"} {
		if !strings.Contains(out, header) {
			t.Errorf("expected header %q in output, got:\n%s", header, out)
		}
	}
}

func TestMarkdownFormatter_Render_IndexedRows(t *testing.T) {
	f := NewMarkdownFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 3, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 4, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 5, 9, 0, 0, 0, time.UTC),
	}
	out := f.Render("0 9 * * *", times)
	for _, idx := range []string{"| 1 |", "| 2 |", "| 3 |"} {
		if !strings.Contains(out, idx) {
			t.Errorf("expected row index %q in output, got:\n%s", idx, out)
		}
	}
}

func TestMarkdownFormatter_Render_ContainsWeekday(t *testing.T) {
	f := NewMarkdownFormatter("UTC")
	monday := time.Date(2024, 6, 3, 9, 0, 0, 0, time.UTC) // Monday
	out := f.Render("0 9 * * 1", []time.Time{monday})
	if !strings.Contains(out, "Monday") {
		t.Errorf("expected weekday Monday in output, got:\n%s", out)
	}
}
