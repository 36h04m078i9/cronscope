package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestTOMLFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewTOMLFormatter()
	times := []time.Time{time.Now().Add(time.Minute)}
	out, err := f.Render("0 * * * *", times, "UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `expression = "0 * * * *"`) {
		t.Errorf("expected expression in output, got:\n%s", out)
	}
}

func TestTOMLFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewTOMLFormatter()
	times := []time.Time{time.Now().Add(time.Minute)}
	out, err := f.Render("0 * * * *", times, "America/New_York")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `timezone = "America/New_York"`) {
		t.Errorf("expected timezone in output, got:\n%s", out)
	}
}

func TestTOMLFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewTOMLFormatter()
	times := []time.Time{time.Now().Add(time.Minute)}
	out, err := f.Render("0 * * * *", times, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `timezone = "UTC"`) {
		t.Errorf("expected default UTC timezone, got:\n%s", out)
	}
}

func TestTOMLFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewTOMLFormatter()
	out, err := f.Render("0 * * * *", []time.Time{}, "UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "count = 0") {
		t.Errorf("expected count = 0 in output, got:\n%s", out)
	}
	if !strings.Contains(out, "[[executions]]") {
		t.Errorf("expected executions block in output, got:\n%s", out)
	}
}

func TestTOMLFormatter_Render_ContainsEntries(t *testing.T) {
	f := NewTOMLFormatter()
	now := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	times := []time.Time{now, now.Add(time.Hour)}
	out, err := f.Render("0 * * * *", times, "UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "count = 2") {
		t.Errorf("expected count = 2, got:\n%s", out)
	}
	if strings.Count(out, "[[executions]]") != 2 {
		t.Errorf("expected 2 [[executions]] blocks, got:\n%s", out)
	}
	if !strings.Contains(out, "index = 1") || !strings.Contains(out, "index = 2") {
		t.Errorf("expected indexed entries, got:\n%s", out)
	}
}

func TestTOMLFormatter_Render_ContainsUnix(t *testing.T) {
	f := NewTOMLFormatter()
	now := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	times := []time.Time{now}
	out, err := f.Render("0 * * * *", times, "UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "unix = ") {
		t.Errorf("expected unix timestamp in output, got:\n%s", out)
	}
}
