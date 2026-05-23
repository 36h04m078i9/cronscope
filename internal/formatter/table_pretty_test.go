package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestPrettyTableFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewPrettyTableFormatter("UTC")
	out, err := f.Render("0 9 * * 1", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "0 9 * * 1") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestPrettyTableFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewPrettyTableFormatter("America/New_York")
	out, err := f.Render("*/5 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", out)
	}
}

func TestPrettyTableFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewPrettyTableFormatter("")
	out, err := f.Render("*/5 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected default timezone UTC in output, got:\n%s", out)
	}
}

func TestPrettyTableFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewPrettyTableFormatter("UTC")
	out, err := f.Render("0 0 * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "No upcoming executions") {
		t.Errorf("expected empty message, got:\n%s", out)
	}
	if !strings.Contains(out, "Total: 0") {
		t.Errorf("expected total 0, got:\n%s", out)
	}
}

func TestPrettyTableFormatter_Render_IndexedRows(t *testing.T) {
	f := NewPrettyTableFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 2, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 3, 9, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 9 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 1; i <= 3; i++ {
		if !strings.Contains(out, fmt.Sprintf("%d", i)) {
			t.Errorf("expected row index %d in output, got:\n%s", i, out)
		}
	}
	if !strings.Contains(out, "Total: 3") {
		t.Errorf("expected total 3, got:\n%s", out)
	}
}

func TestPrettyTableFormatter_Render_BorderPresent(t *testing.T) {
	f := NewPrettyTableFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 12 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "+") || !strings.Contains(out, "|") {
		t.Errorf("expected table borders in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Scheduled Time") {
		t.Errorf("expected header 'Scheduled Time', got:\n%s", out)
	}
}
