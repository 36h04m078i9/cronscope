package parser_test

import (
	"testing"
	"time"

	"github.com/cronscope/cronscope/internal/parser"
)

func TestLoadTimezone_UTC(t *testing.T) {
	loc, err := parser.LoadTimezone("UTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc != time.UTC {
		t.Errorf("expected time.UTC, got %v", loc)
	}
}

func TestLoadTimezone_Empty(t *testing.T) {
	loc, err := parser.LoadTimezone("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc != time.UTC {
		t.Errorf("expected time.UTC for empty string, got %v", loc)
	}
}

func TestLoadTimezone_Valid(t *testing.T) {
	loc, err := parser.LoadTimezone("America/New_York")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.String() != "America/New_York" {
		t.Errorf("expected America/New_York, got %v", loc)
	}
}

func TestLoadTimezone_Invalid(t *testing.T) {
	_, err := parser.LoadTimezone("Mars/Olympus")
	if err == nil {
		t.Fatal("expected error for invalid timezone, got nil")
	}
}

func TestFormatWithTimezone(t *testing.T) {
	loc, _ := parser.LoadTimezone("America/New_York")
	t0 := time.Date(2024, 6, 15, 17, 0, 0, 0, time.UTC)

	formatted := parser.FormatWithTimezone(t0, parser.DefaultTimeLayout, loc)
	if formatted == "" {
		t.Error("expected non-empty formatted string")
	}

	// 17:00 UTC should be 13:00 EDT (UTC-4)
	expectedHour := "13:00:00"
	locTime := t0.In(loc)
	got := locTime.Format("15:04:05")
	if got != expectedHour {
		t.Errorf("expected %s, got %s", expectedHour, got)
	}
}
