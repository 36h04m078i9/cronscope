package formatter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/cronscope/internal/formatter"
)

func TestXMLFormatter_Render_ContainsXMLHeader(t *testing.T) {
	f := formatter.NewXMLFormatter("UTC")
	out, err := f.Render("0 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "<?xml") {
		t.Errorf("expected XML header, got: %s", out)
	}
}

func TestXMLFormatter_Render_ContainsExpression(t *testing.T) {
	f := formatter.NewXMLFormatter("UTC")
	expr := "30 8 * * 1-5"
	out, err := f.Render(expr, []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, expr) {
		t.Errorf("expected expression %q in output, got: %s", expr, out)
	}
}

func TestXMLFormatter_Render_ContainsTimezone(t *testing.T) {
	tz := "America/New_York"
	f := formatter.NewXMLFormatter(tz)
	out, err := f.Render("* * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, tz) {
		t.Errorf("expected timezone %q in output, got: %s", tz, out)
	}
}

func TestXMLFormatter_Render_DefaultTimezone(t *testing.T) {
	f := formatter.NewXMLFormatter("")
	out, err := f.Render("* * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected default timezone UTC in output, got: %s", out)
	}
}

func TestXMLFormatter_Render_ContainsEntries(t *testing.T) {
	f := formatter.NewXMLFormatter("UTC")
	times := []time.Time{
		time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 1, 13, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 * * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "<entry") {
		t.Errorf("expected <entry> elements in output, got: %s", out)
	}
	if !strings.Contains(out, "2024-06-01T12:00:00Z") {
		t.Errorf("expected first time in output, got: %s", out)
	}
}

func TestXMLFormatter_Render_EmptyTimes(t *testing.T) {
	f := formatter.NewXMLFormatter("UTC")
	out, err := f.Render("0 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "<schedule") {
		t.Errorf("expected schedule root element, got: %s", out)
	}
}
