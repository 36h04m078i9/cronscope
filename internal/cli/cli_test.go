package cli

import (
	"strings"
	"testing"
)

func TestRun_MissingExpression(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{}, &buf)
	if err == nil {
		t.Fatal("expected error for missing expression")
	}
	if !strings.Contains(err.Error(), "usage") {
		t.Errorf("expected usage hint in error, got: %v", err)
	}
}

func TestRun_InvalidExpression(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"not-a-cron"}, &buf)
	if err == nil {
		t.Fatal("expected error for invalid cron expression")
	}
}

func TestRun_InvalidTimezone(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"-tz", "Invalid/Zone", "* * * * *"}, &buf)
	if err == nil {
		t.Fatal("expected error for invalid timezone")
	}
}

func TestExecute_TableFormat(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"-format", "table", "-n", "3", "* * * * *"}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "* * * * *") {
		t.Error("expected expression in table output")
	}
}

func TestExecute_JSONFormat(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"-format", "json", "-n", "2", "0 9 * * 1"}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "{") {
		t.Error("expected JSON output")
	}
}

func TestExecute_ICalFormat(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"-format", "ical", "-n", "2", "0 9 * * *"}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "BEGIN:VCALENDAR") {
		t.Error("expected iCal output")
	}
	if strings.Count(out, "BEGIN:VEVENT") != 2 {
		t.Errorf("expected 2 VEVENT blocks, got output:\n%s", out)
	}
}

func TestExecute_HumanizeFormat(t *testing.T) {
	var buf strings.Builder
	err := Run([]string{"-format", "humanize", "-n", "1", "* * * * *"}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.Len() == 0 {
		t.Error("expected non-empty humanize output")
	}
}
