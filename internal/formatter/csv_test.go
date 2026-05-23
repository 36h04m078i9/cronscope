package formatter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/cronscope/internal/formatter"
)

func TestCSVFormatter_Render_ContainsHeader(t *testing.T) {
	f := formatter.NewCSVFormatter("UTC")
	times := []time.Time{time.Now().UTC()}
	out, err := f.Render("0 * * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "index") || !strings.Contains(out, "expression") {
		t.Errorf("expected CSV header, got: %s", out)
	}
}

func TestCSVFormatter_Render_ContainsExpression(t *testing.T) {
	f := formatter.NewCSVFormatter("UTC")
	expr := "*/5 * * * *"
	times := []time.Time{time.Now().UTC()}
	out, err := f.Render(expr, times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, expr) {
		t.Errorf("expected expression %q in output, got: %s", expr, out)
	}
}

func TestCSVFormatter_Render_ContainsTimezone(t *testing.T) {
	tz := "America/New_York"
	f := formatter.NewCSVFormatter(tz)
	times := []time.Time{time.Now().UTC()}
	out, err := f.Render("0 0 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, tz) {
		t.Errorf("expected timezone %q in output, got: %s", tz, out)
	}
}

func TestCSVFormatter_Render_DefaultTimezone(t *testing.T) {
	f := formatter.NewCSVFormatter("")
	times := []time.Time{time.Now().UTC()}
	out, err := f.Render("0 0 * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected default UTC timezone in output, got: %s", out)
	}
}

func TestCSVFormatter_Render_EmptyTimes(t *testing.T) {
	f := formatter.NewCSVFormatter("UTC")
	out, err := f.Render("0 * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 1 {
		t.Errorf("expected only header row for empty times, got %d lines", len(lines))
	}
}

func TestCSVFormatter_Render_IndexedRows(t *testing.T) {
	f := formatter.NewCSVFormatter("UTC")
	now := time.Now().UTC()
	times := []time.Time{now, now.Add(time.Hour), now.Add(2 * time.Hour)}
	out, err := f.Render("0 * * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "1,") || !strings.Contains(out, "2,") || !strings.Contains(out, "3,") {
		t.Errorf("expected indexed rows 1-3 in output, got: %s", out)
	}
}
