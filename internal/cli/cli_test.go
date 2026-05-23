package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_MissingExpression(t *testing.T) {
	err := Run([]string{})
	if err == nil {
		t.Fatal("expected error for missing expression")
	}
	if !strings.Contains(err.Error(), "usage") {
		t.Errorf("expected usage hint, got: %v", err)
	}
}

func TestRun_InvalidExpression(t *testing.T) {
	err := Run([]string{"not-a-cron"})
	if err == nil {
		t.Fatal("expected error for invalid cron expression")
	}
}

func TestRun_InvalidTimezone(t *testing.T) {
	err := Run([]string{"-tz", "Mars/Olympus", "* * * * *"})
	if err == nil {
		t.Fatal("expected error for invalid timezone")
	}
	if !strings.Contains(err.Error(), "invalid timezone") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestExecute_TableFormat(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Expression: "* * * * *",
		Count:      3,
		Timezone:   "UTC",
		Format:     "table",
		Output:     &buf,
	}
	if err := execute(cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "* * * * *") {
		t.Errorf("output missing expression, got:\n%s", buf.String())
	}
}

func TestExecute_JSONFormat(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Expression: "0 9 * * 1",
		Count:      2,
		Timezone:   "UTC",
		Format:     "json",
		Output:     &buf,
	}
	if err := execute(cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "expression") {
		t.Errorf("JSON output missing 'expression' key, got:\n%s", buf.String())
	}
}

func TestExecute_PlainFormat(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Expression: "30 6 * * *",
		Count:      2,
		Timezone:   "UTC",
		Format:     "plain",
		Output:     &buf,
	}
	if err := execute(cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.Len() == 0 {
		t.Error("expected non-empty plain output")
	}
}
