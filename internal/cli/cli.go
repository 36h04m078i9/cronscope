// Package cli provides the command-line interface for cronscope.
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/cronscope/internal/formatter"
	"github.com/user/cronscope/internal/parser"
)

const defaultN = 5

// Run parses CLI arguments and executes the cronscope command.
func Run(args []string, out io.Writer) error {
	expression, opts, err := parseArgs(args)
	if err != nil {
		return err
	}
	return execute(expression, opts, out)
}

type options struct {
	n        int
	timezone string
	format   string
}

func parseArgs(args []string) (string, options, error) {
	fs := flag.NewFlagSet("cronscope", flag.ContinueOnError)
	n := fs.Int("n", defaultN, "number of next execution times to show")
	tz := fs.String("tz", "UTC", "timezone for output (e.g. America/New_York)")
	fmt := fs.String("format", "table", "output format: table, plain, json, color, humanize, ical, csv, markdown, xml")

	if err := fs.Parse(args); err != nil {
		return "", options{}, err
	}

	remaining := fs.Args()
	if len(remaining) == 0 {
		return "", options{}, fmt2.Errorf("missing cron expression")
	}

	return remaining[0], options{n: *n, timezone: *tz, format: *fmt}, nil
}

func execute(expression string, opts options, out io.Writer) error {
	loc, err := parser.LoadTimezone(opts.timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone %q: %w", opts.timezone, err)
	}

	sched, err := parser.Parse(expression)
	if err != nil {
		return fmt.Errorf("invalid cron expression %q: %w", expression, err)
	}

	times := parser.NextN(sched, opts.n, loc)

	f, err := resolveFormatter(opts.format, opts.timezone)
	if err != nil {
		return err
	}

	result, err := f.Render(expression, times)
	if err != nil {
		return fmt.Errorf("render error: %w", err)
	}

	_, err = fmt.Fprint(out, result)
	return err
}

func resolveFormatter(format, timezone string) (formatter.Formatter, error) {
	switch format {
	case "table":
		return formatter.NewTableFormatter(timezone), nil
	case "plain":
		return formatter.NewPlainFormatter(timezone), nil
	case "json":
		return formatter.NewJSONFormatter(timezone), nil
	case "color":
		return formatter.NewColorFormatter(timezone), nil
	case "humanize":
		return formatter.NewHumanizeFormatter(timezone), nil
	case "ical":
		return formatter.NewICalFormatter(timezone), nil
	case "csv":
		return formatter.NewCSVFormatter(timezone), nil
	case "markdown":
		return formatter.NewMarkdownFormatter(timezone), nil
	case "xml":
		return formatter.NewXMLFormatter(timezone), nil
	default:
		return nil, fmt.Errorf("unknown format %q", format)
	}
}

// Main is the entry point called from cmd/cronscope/main.go.
func Main() {
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
