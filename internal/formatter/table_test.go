package formatter_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/user/cronscope/internal/formatter"
)

func TestTableFormatter_Render_ContainsExpression(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.NewTableFormatter(&buf)

	times := []time.Time{
		time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 15, 11, 0, 0, 0, time.UTC),
	}

	err := f.Render("0 * * * *", times)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "0 * * * *") {
		t.Errorf("expected output to contain expression, got:\n%s", output)
	}
}

func TestTableFormatter_Render_ContainsTimes(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.NewTableFormatter(&buf)

	times := []time.Time{
		time.Date(2024, 3, 10, 8, 30, 0, 0, time.UTC),
	}

	_ = f.Render("30 8 * * *", times)
	output := buf.String()

	if !strings.Contains(output, "2024-03-10") {
		t.Errorf("expected output to contain date 2024-03-10, got:\n%s", output)
	}
	if !strings.Contains(output, "08:30:00") {
		t.Errorf("expected output to contain time 08:30:00, got:\n%s", output)
	}
}

func TestTableFormatter_Render_EmptyTimes(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.NewTableFormatter(&buf)

	err := f.Render("* * * * *", []time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Total: 0") {
		t.Errorf("expected 'Total: 0' in output, got:\n%s", output)
	}
}

func TestTableFormatter_Render_IndexedRows(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.NewTableFormatter(&buf)

	times := []time.Time{
		time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 1, 1, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 1, 2, 0, 0, 0, time.UTC),
	}

	_ = f.Render("0 * * * *", times)
	output := buf.String()

	for _, idx := range []string{" 1.", " 2.", " 3."} {
		if !strings.Contains(output, idx) {
			t.Errorf("expected index %q in output, got:\n%s", idx, output)
		}
	}
}
