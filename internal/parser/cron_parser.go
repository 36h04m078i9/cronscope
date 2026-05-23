package parser

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Schedule wraps a parsed cron schedule with its expression.
type Schedule struct {
	Expression string
	spec       cron.Schedule
}

// Parse validates and parses a cron expression string.
// Supports standard 5-field expressions (minute hour dom month dow).
func Parse(expression string) (*Schedule, error) {
	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	spec, err := parser.Parse(expression)
	if err != nil {
		return nil, fmt.Errorf("invalid cron expression %q: %w", expression, err)
	}

	return &Schedule{
		Expression: expression,
		spec:       spec,
	}, nil
}

// NextN returns the next n execution times starting from the given time.
func (s *Schedule) NextN(from time.Time, n int) []time.Time {
	if n <= 0 {
		return nil
	}

	times := make([]time.Time, 0, n)
	current := from

	for i := 0; i < n; i++ {
		next := s.spec.Next(current)
		if next.IsZero() {
			break
		}
		times = append(times, next)
		current = next
	}

	return times
}

// NextNInLocation returns the next n execution times in the specified timezone.
func (s *Schedule) NextNInLocation(from time.Time, n int, loc *time.Location) []time.Time {
	times := s.NextN(from.In(loc), n)
	result := make([]time.Time, len(times))
	for i, t := range times {
		result[i] = t.In(loc)
	}
	return result
}
