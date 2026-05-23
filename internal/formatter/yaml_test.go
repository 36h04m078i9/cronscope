package formatter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/cronscope/internal/formatter"
)

func TestYAMLFormatter_Render_ContainsExpression(t *testing.T) {
	times := []time.Time{time.Now().UTC()}
	f := formatter.NewYAMLFormatter("0 9 * * 1-5", "UTC", times)
	out := f.Render()
	if !strings.Contains(out, "0 9 * * 1-5") {
		t.Errorf("expected output to contain expression, got:\n%s", out)
	}
}

func TestYAMLFormatter_Render_ContainsTimezone(t *testing.T) {
	times := []time.Time{time.Now().UTC()}
	f := formatter.NewYAMLFormatter("0 9 * * *", "America/New_York", times)
	out := f.Render()
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected output to contain timezone, got:\n%s", out)
	}
}

func TestYAMLFormatter_Render_DefaultTimezone(t *testing.T) {
	times := []time.Time{time.Now().UTC()}
	f := formatter.NewYAMLFormatter("0 9 * * *", "", times)
	out := f.Render()
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected output to default timezone to UTC, got:\n%s", out)
	}
}

func TestYAMLFormatter_Render_EmptyTimes(t *testing.T) {
	f := formatter.NewYAMLFormatter("0 9 * * *", "UTC", []time.Time{})
	out := f.Render()
	if !strings.Contains(out, "executions:") {
		t.Errorf("expected output to contain executions key, got:\n%s", out)
	}
	if !strings.Contains(out, "[]") {
		t.Errorf("expected output to contain empty list marker, got:\n%s", out)
	}
}

func TestYAMLFormatter_Render_ContainsEntries(t *testing.T) {
	t1 := time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 6, 2, 9, 0, 0, 0, time.UTC)
	f := formatter.NewYAMLFormatter("0 9 * * *", "UTC", []time.Time{t1, t2})
	out := f.Render()
	if !strings.Contains(out, "index: 1") {
		t.Errorf("expected output to contain index 1, got:\n%s", out)
	}
	if !strings.Contains(out, "index: 2") {
		t.Errorf("expected output to contain index 2, got:\n%s", out)
	}
}

func TestYAMLFormatter_Render_YAMLHeader(t *testing.T) {
	f := formatter.NewYAMLFormatter("0 9 * * *", "UTC", []time.Time{})
	out := f.Render()
	if !strings.HasPrefix(out, "---") {
		t.Errorf("expected output to start with YAML document marker, got:\n%s", out)
	}
}
