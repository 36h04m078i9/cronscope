package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestICalFormatter_Render_ContainsCalendarHeader(t *testing.T) {
	f := NewICalFormatter("UTC")
	result := f.Render("0 9 * * 1", []time.Time{})
	if !strings.Contains(result, "BEGIN:VCALENDAR") {
		t.Error("expected BEGIN:VCALENDAR in output")
	}
	if !strings.Contains(result, "END:VCALENDAR") {
		t.Error("expected END:VCALENDAR in output")
	}
}

func TestICalFormatter_Render_ContainsExpression(t *testing.T) {
	expr := "*/5 * * * *"
	f := NewICalFormatter("UTC")
	result := f.Render(expr, []time.Time{})
	if !strings.Contains(result, expr) {
		t.Errorf("expected expression %q in output", expr)
	}
}

func TestICalFormatter_Render_ContainsTimezone(t *testing.T) {
	f := NewICalFormatter("America/New_York")
	result := f.Render("0 0 * * *", []time.Time{})
	if !strings.Contains(result, "America/New_York") {
		t.Error("expected timezone in output")
	}
}

func TestICalFormatter_Render_DefaultTimezone(t *testing.T) {
	f := NewICalFormatter("")
	result := f.Render("0 0 * * *", []time.Time{})
	if !strings.Contains(result, "UTC") {
		t.Error("expected UTC as default timezone")
	}
}

func TestICalFormatter_Render_ContainsVEvents(t *testing.T) {
	f := NewICalFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 16, 9, 0, 0, 0, time.UTC),
	}
	result := f.Render("0 9 * * *", times)
	count := strings.Count(result, "BEGIN:VEVENT")
	if count != 2 {
		t.Errorf("expected 2 VEVENT blocks, got %d", count)
	}
}

func TestICalFormatter_Render_EmptyTimes(t *testing.T) {
	f := NewICalFormatter("UTC")
	result := f.Render("0 0 * * *", []time.Time{})
	if strings.Contains(result, "BEGIN:VEVENT") {
		t.Error("expected no VEVENT blocks for empty times")
	}
}

func TestICalFormatter_Render_UIDs_Unique(t *testing.T) {
	f := NewICalFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 3, 1, 8, 0, 0, 0, time.UTC),
		time.Date(2024, 3, 1, 9, 0, 0, 0, time.UTC),
	}
	result := f.Render("0 8-9 * * *", times)
	if !strings.Contains(result, "UID:cronscope-1-") {
		t.Error("expected UID for first event")
	}
	if !strings.Contains(result, "UID:cronscope-2-") {
		t.Error("expected UID for second event")
	}
}
