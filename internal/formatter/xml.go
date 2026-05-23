package formatter

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// XMLEntry represents a single scheduled execution time in XML.
type XMLEntry struct {
	Index     int    `xml:"index,attr"`
	Formatted string `xml:"formatted"`
	Unix      int64  `xml:"unix"`
}

// XMLSchedule is the root XML document structure.
type XMLSchedule struct {
	XMLName    xml.Name   `xml:"schedule"`
	Expression string     `xml:"expression,attr"`
	Timezone   string     `xml:"timezone,attr"`
	Entries    []XMLEntry `xml:"entry"`
}

// xmlFormatter renders cron schedule output as XML.
type xmlFormatter struct {
	timezone string
}

// NewXMLFormatter returns a Formatter that renders output as XML.
func NewXMLFormatter(timezone string) Formatter {
	tz := timezone
	if tz == "" {
		tz = "UTC"
	}
	return &xmlFormatter{timezone: tz}
}

// Render formats the expression and times as an XML document.
func (f *xmlFormatter) Render(expression string, times []time.Time) (string, error) {
	entries := make([]XMLEntry, 0, len(times))
	for i, t := range times {
		entries = append(entries, XMLEntry{
			Index:     i + 1,
			Formatted: t.Format(time.RFC3339),
			Unix:      t.Unix(),
		})
	}

	doc := XMLSchedule{
		Expression: expression,
		Timezone:   f.timezone,
		Entries:    entries,
	}

	data, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", fmt.Errorf("xml formatter: marshal error: %w", err)
	}

	var sb strings.Builder
	sb.WriteString(xml.Header)
	sb.Write(data)
	sb.WriteString("\n")
	return sb.String(), nil
}
