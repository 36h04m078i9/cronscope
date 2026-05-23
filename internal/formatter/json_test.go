package formatter

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestJSONFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewJSONFormatter("0 9 * * 1-5", "UTC")
	times := []time.Time{time.Now().UTC()}

	out, err := f.Render(times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got: %s", out)
	}
}

func TestJSONFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewJSONFormatter("*/5 * * * *", "America/New_York")
	times := []time.Time{time.Now()}

	out, err := f.Render(times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got: %s", out)
	}
}

func TestJSONFormatter_Render_ValidJSON(t *testing.T) {
	f := NewJSONFormatter("0 0 * * *", "UTC")
	times := []time.Time{
		time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC),
	}

	out, err := f.Render(times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		t.Errorf("output is not valid JSON: %v", err)
	}
}

func TestJSONFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewJSONFormatter("0 0 * * *", "UTC")
	out, err := f.Render([]time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "next_times") {
		t.Errorf("expected next_times key in output, got: %s", out)
	}
}

func TestJSONFormatter_DefaultTimezone(t *testing.T) {
	f := NewJSONFormatter("0 0 * * *", "")
	if f.timezone != "UTC" {
		t.Errorf("expected default timezone UTC, got: %s", f.timezone)
	}
}
