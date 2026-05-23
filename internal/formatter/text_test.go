package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestTextFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewTextFormatter("UTC")
	out, err := f.Render("0 9 * * 1-5", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestTextFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewTextFormatter("America/New_York")
	out, err := f.Render("*/5 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", out)
	}
}

func TestTextFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewTextFormatter("")
	out, err := f.Render("*/5 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected default timezone UTC, got:\n%s", out)
	}
}

func TestTextFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewTextFormatter("UTC")
	out, err := f.Render("0 0 * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "No upcoming executions") {
		t.Errorf("expected empty message, got:\n%s", out)
	}
}

func TestTextFormatter_Render_IndexedRows(t *testing.T) {
	f := NewTextFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 2, 9, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 9 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "  1.") {
		t.Errorf("expected row index 1, got:\n%s", out)
	}
	if !strings.Contains(out, "  2.") {
		t.Errorf("expected row index 2, got:\n%s", out)
	}
}

func TestTextFormatter_Render_ContainsDate(t *testing.T) {
	f := NewTextFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 9 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "2024-06-01") {
		t.Errorf("expected formatted date in output, got:\n%s", out)
	}
}
