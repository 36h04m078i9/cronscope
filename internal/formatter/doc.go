// Package formatter provides output rendering utilities for cronscope.
//
// It currently exposes TableFormatter, which prints upcoming cron execution
// times in a human-readable table to any io.Writer (e.g. os.Stdout).
//
// Example usage:
//
//	f := formatter.NewTableFormatter(os.Stdout)
//	f.Render("0 * * * *", times)
package formatter
