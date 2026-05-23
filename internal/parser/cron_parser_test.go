package parser_test

import (
	"testing"
	"time"

	"github.com/cronscope/cronscope/internal/parser"
)

func TestParse_ValidExpression(t *testing.T) {
	tests := []struct {
		name string
		expr string
	}{
		{"every minute", "* * * * *"},
		{"every hour", "0 * * * *"},
		{"daily at midnight", "0 0 * * *"},
		{"weekly on monday", "0 9 * * 1"},
		{"@hourly descriptor", "@hourly"},
		{"@daily descriptor", "@daily"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := parser.Parse(tt.expr)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if s.Expression != tt.expr {
				t.Errorf("expected expression %q, got %q", tt.expr, s.Expression)
			}
		})
	}
}

func TestParse_InvalidExpression(t *testing.T) {
	invalidExprs := []string{
		"not a cron",
		"* * * *",
		"60 * * * *",
		"",
	}

	for _, expr := range invalidExprs {
		t.Run(expr, func(t *testing.T) {
			_, err := parser.Parse(expr)
			if err == nil {
				t.Fatalf("expected error for expression %q, got nil", expr)
			}
		})
	}
}

func TestNextN(t *testing.T) {
	s, err := parser.Parse("0 * * * *")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	from := time.Date(2024, 1, 1, 0, 30, 0, 0, time.UTC)
	times := s.NextN(from, 3)

	if len(times) != 3 {
		t.Fatalf("expected 3 times, got %d", len(times))
	}

	expected := []time.Time{
		time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 2, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC),
	}

	for i, exp := range expected {
		if !times[i].Equal(exp) {
			t.Errorf("time[%d]: expected %v, got %v", i, exp, times[i])
		}
	}
}

func TestNextNInLocation(t *testing.T) {
	s, err := parser.Parse("0 12 * * *")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	loc, _ := time.LoadLocation("America/New_York")
	from := time.Date(2024, 6, 1, 0, 0, 0, 0, loc)
	times := s.NextNInLocation(from, 1, loc)

	if len(times) != 1 {
		t.Fatalf("expected 1 time, got %d", len(times))
	}

	if times[0].Location().String() != loc.String() {
		t.Errorf("expected location %s, got %s", loc, times[0].Location())
	}

	if times[0].Hour() != 12 {
		t.Errorf("expected hour 12, got %d", times[0].Hour())
	}
}
