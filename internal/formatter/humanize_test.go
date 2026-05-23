package formatter

import (
	"strings"
	"testing"
	"time"
)

func fixedNow(t time.Time) func() time.Time {
	return func() time.Time { return t }
}

func TestHumanizeFormatter_Render_ContainsExpression(t *testing.T) {
	f := NewHumanizeFormatter("0 9 * * 1-5", "UTC").(*HumanizeFormatter)
	f.now = fixedNow(time.Now())
	out := f.Render([]time.Time{})
	if !strings.Contains(out, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestHumanizeFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewHumanizeFormatter("* * * * *", "America/New_York").(*HumanizeFormatter)
	f.now = fixedNow(time.Now())
	out := f.Render([]time.Time{})
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", out)
	}
}

func TestHumanizeFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewHumanizeFormatter("* * * * *", "").(*HumanizeFormatter)
	f.now = fixedNow(time.Now())
	out := f.Render([]time.Time{})
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected UTC default timezone, got:\n%s", out)
	}
}

func TestHumanizeFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewHumanizeFormatter("* * * * *", "UTC").(*HumanizeFormatter)
	f.now = fixedNow(time.Now())
	out := f.Render([]time.Time{})
	if !strings.Contains(out, "No upcoming") {
		t.Errorf("expected empty message, got:\n%s", out)
	}
}

func TestHumanizeFormatter_Render_RelativeLabels(t *testing.T) {
	now := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	times := []time.Time{
		now.Add(30 * time.Second),
		now.Add(5 * time.Minute),
		now.Add(3 * time.Hour),
		now.Add(48 * time.Hour),
	}
	f := NewHumanizeFormatter("* * * * *", "UTC").(*HumanizeFormatter)
	f.now = fixedNow(now)
	out := f.Render(times)

	for _, want := range []string{"in 30 seconds", "in 5 minutes", "in 3 hours", "in 2 days"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in output, got:\n%s", want, out)
		}
	}
}

func TestHumanizeFormatter_Render_IndexedRows(t *testing.T) {
	now := time.Now()
	times := []time.Time{now.Add(time.Hour), now.Add(2 * time.Hour)}
	f := NewHumanizeFormatter("0 * * * *", "UTC").(*HumanizeFormatter)
	f.now = fixedNow(now)
	out := f.Render(times)
	if !strings.Contains(out, "1.") || !strings.Contains(out, "2.") {
		t.Errorf("expected indexed rows, got:\n%s", out)
	}
}

func TestHumanDuration_Singular(t *testing.T) {
	if got := humanDuration(1 * time.Minute); got != "in 1 minute" {
		t.Errorf("expected 'in 1 minute', got %q", got)
	}
	if got := humanDuration(1 * time.Hour); got != "in 1 hour" {
		t.Errorf("expected 'in 1 hour', got %q", got)
	}
}
