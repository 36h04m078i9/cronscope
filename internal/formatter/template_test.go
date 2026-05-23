package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestTemplateFormatter_Render_ContainsExpression(t *testing.T) {
	tmplStr := "Expression: {{.Expression}}"
	f, err := NewTemplateFormatter(tmplStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, err := f.Render("0 * * * *", "UTC", nil)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if !strings.Contains(out, "0 * * * *") {
		t.Errorf("expected expression in output, got: %s", out)
	}
}

func TestTemplateFormatter_Render_ContainsTimezone(t *testing.T) {
	tmplStr := "Timezone: {{.Timezone}}"
	f, err := NewTemplateFormatter(tmplStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, err := f.Render("0 * * * *", "America/New_York", nil)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if !strings.Contains(out, "America/New_York") {
		t.Errorf("expected timezone in output, got: %s", out)
	}
}

func TestTemplateFormatter_Render_DefaultTimezone(t *testing.T) {
	tmplStr := "Timezone: {{.Timezone}}"
	f, err := NewTemplateFormatter(tmplStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, err := f.Render("0 * * * *", "", nil)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if !strings.Contains(out, "UTC") {
		t.Errorf("expected UTC default timezone, got: %s", out)
	}
}

func TestTemplateFormatter_Render_ContainsTimes(t *testing.T) {
	tmplStr := "{{range .Times}}{{.Index}}: {{.Formatted}}\n{{end}}"
	f, err := NewTemplateFormatter(tmplStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	times := []time.Time{
		time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 15, 11, 0, 0, 0, time.UTC),
	}
	out, err := f.Render("0 * * * *", "UTC", times)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if !strings.Contains(out, "1:") || !strings.Contains(out, "2:") {
		t.Errorf("expected indexed entries, got: %s", out)
	}
	if !strings.Contains(out, "2024-01-15") {
		t.Errorf("expected formatted date in output, got: %s", out)
	}
}

func TestTemplateFormatter_InvalidTemplate(t *testing.T) {
	_, err := NewTemplateFormatter("{{.Unclosed")
	if err == nil {
		t.Error("expected error for invalid template, got nil")
	}
}

func TestTemplateFormatter_Render_UnixTimestamp(t *testing.T) {
	tmplStr := "{{range .Times}}{{.Unix}}{{end}}"
	f, err := NewTemplateFormatter(tmplStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ts := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	out, err := f.Render("0 * * * *", "UTC", []time.Time{ts})
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	expected := "1705312800"
	if !strings.Contains(out, expected) {
		t.Errorf("expected unix timestamp %s in output, got: %s", expected, out)
	}
}
