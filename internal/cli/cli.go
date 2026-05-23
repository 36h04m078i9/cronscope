package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/cronscope/internal/formatter"
	"github.com/user/cronscope/internal/parser"
)

// Config holds parsed CLI arguments.
type Config struct {
	Expression string
	Count      int
	Timezone   string
	Format     string
	Output     io.Writer
}

// Run parses CLI arguments and executes cronscope.
func Run(args []string) error {
	cfg, err := parseArgs(args)
	if err != nil {
		return err
	}
	return execute(cfg)
}

func parseArgs(args []string) (*Config, error) {
	fs := flag.NewFlagSet("cronscope", flag.ContinueOnError)
	count := fs.Int("n", 5, "number of next execution times to show")
	tz := fs.String("tz", "UTC", "timezone (e.g. America/New_York)")
	fmt_ := fs.String("format", "table", "output format: table, plain, json")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	if fs.NArg() < 1 {
		return nil, errors.New("usage: cronscope [flags] \"<cron expression>\"")
	}
	return &Config{
		Expression: fs.Arg(0),
		Count:      *count,
		Timezone:   *tz,
		Format:     *fmt_,
		Output:     os.Stdout,
	}, nil
}

func execute(cfg *Config) error {
	loc, err := parser.LoadTimezone(cfg.Timezone)
	if err != nil {
		return fmt.Errorf("invalid timezone %q: %w", cfg.Timezone, err)
	}

	schedule, err := parser.Parse(cfg.Expression)
	if err != nil {
		return fmt.Errorf("invalid cron expression: %w", err)
	}

	times := parser.NextNInLocation(schedule, cfg.Count, loc)

	var f formatter.Formatter
	switch cfg.Format {
	case "json":
		f = formatter.NewJSONFormatter(cfg.Expression, cfg.Timezone)
	case "plain":
		f = formatter.NewPlainFormatter(cfg.Expression, cfg.Timezone)
	default:
		f = formatter.NewTableFormatter(cfg.Expression, cfg.Timezone)
	}

	output, err := f.Render(times)
	if err != nil {
		return fmt.Errorf("render error: %w", err)
	}
	fmt.Fprintln(cfg.Output, output)
	return nil
}
