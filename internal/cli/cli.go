package cli

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/cronscope/internal/formatter"
	"github.com/user/cronscope/internal/parser"
)

// Run is the entry point for the CLI. It parses arguments and executes the command.
func Run(args []string, out io.Writer) error {
	opts, err := parseArgs(args)
	if err != nil {
		return err
	}
	return execute(opts, out)
}

type options struct {
	expression string
	n          int
	timezone   string
	format     string
}

func parseArgs(args []string) (*options, error) {
	fs := flag.NewFlagSet("cronscope", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	n := fs.Int("n", 5, "number of next execution times to show")
	tz := fs.String("tz", "UTC", "timezone (e.g. America/New_York)")
	fmt := fs.String("format", "table", "output format: table, plain, json, color, humanize, ical")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	if fs.NArg() < 1 {
		return nil, fmt.Errorf("usage: cronscope [flags] <cron expression>")
	}

	expr := ""
	for i, a := range fs.Args() {
		if i > 0 {
			expr += " "
		}
		expr += a
	}

	return &options{
		expression: expr,
		n:          *n,
		timezone:   *tz,
		format:     *fmt,
	}, nil
}

func execute(opts *options, out io.Writer) error {
	loc, err := parser.LoadTimezone(opts.timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone %q: %w", opts.timezone, err)
	}

	schedule, err := parser.Parse(opts.expression)
	if err != nil {
		return fmt.Errorf("invalid cron expression %q: %w", opts.expression, err)
	}

	times := parser.NextNInLocation(schedule, opts.n, loc)

	var f formatter.Formatter
	switch opts.format {
	case "plain":
		f = formatter.NewPlainFormatter(opts.timezone)
	case "json":
		f = formatter.NewJSONFormatter(opts.timezone)
	case "color":
		f = formatter.NewColorFormatter(opts.timezone)
	case "humanize":
		f = formatter.NewHumanizeFormatter(opts.timezone)
	case "ical":
		f = formatter.NewICalFormatter(opts.timezone)
	default:
		f = formatter.NewTableFormatter(opts.timezone)
	}

	_, err = fmt.Fprint(out, f.Render(opts.expression, times))
	return err
}

func Main() {
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
