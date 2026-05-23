// Package cli provides the command-line interface for cronscope.
//
// It parses flags and positional arguments, delegates to the parser
// and formatter packages, and writes the final output to stdout.
//
// Usage:
//
//	cronscope [flags] "<cron expression>"
//
// Flags:
//
//	-n int        Number of next execution times to display (default 5)
//	-tz string    Timezone name, e.g. "America/New_York" (default "UTC")
//	-format string Output format: table | plain | json (default "table")
//
// Examples:
//
//	cronscope "*/5 * * * *"
//	cronscope -n 10 -tz Europe/London "0 9 * * 1-5"
//	cronscope -format json "0 0 1 * *"
package cli
