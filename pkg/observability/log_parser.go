package observability

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ParsedLogTelemetry encapsulates extracted incident diagnostic data.
type ParsedLogTelemetry struct {
	ServiceName    string    `json:"serviceName"`
	Severity       string    `json:"severity"`
	ErrorPattern   string    `json:"errorPattern"`
	FileLocation   string    `json:"fileLocation"`
	LineNumber     int       `json:"lineNumber"`
	LogLines       []string  `json:"logLines"`
	ParsedAt       time.Time `json:"parsedAt"`
	StructuredJSON bool      `json:"structuredJson"`
}

// LogStreamParser parses raw log streams from Loki, OpenTelemetry, or stdin.
type LogStreamParser struct {
	fileLocRegex *regexp.Regexp
}

// NewLogStreamParser initializes the log parser.
func NewLogStreamParser() *LogStreamParser {
	return &LogStreamParser{
		fileLocRegex: regexp.MustCompile(`([a-zA-Z0-9_\-/\.]+\.(go|py|ts|js|java)):(\d+)`),
	}
}

// ParseStream inspects raw log lines and extracts diagnostic indicators.
func (p *LogStreamParser) ParseStream(rawLogs string) ParsedLogTelemetry {
	lines := strings.Split(rawLogs, "\n")
	var cleanLines []string
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed != "" {
			cleanLines = append(cleanLines, trimmed)
		}
	}

	result := ParsedLogTelemetry{
		ServiceName:    "unknown-service",
		Severity:       "ERROR",
		ErrorPattern:   "Production System Error",
		FileLocation:   "unknown_file",
		LineNumber:     0,
		LogLines:       cleanLines,
		ParsedAt:       time.Now(),
		StructuredJSON: false,
	}

	if len(cleanLines) == 0 {
		return result
	}

	// Check if first line is JSON log
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(cleanLines[0]), &jsonMap); err == nil {
		result.StructuredJSON = true
		if svc, ok := jsonMap["service"].(string); ok {
			result.ServiceName = svc
		} else if app, ok := jsonMap["app"].(string); ok {
			result.ServiceName = app
		}
		if level, ok := jsonMap["level"].(string); ok {
			result.Severity = strings.ToUpper(level)
		}
		if msg, ok := jsonMap["message"].(string); ok {
			result.ErrorPattern = msg
		}
	}

	// Regex scan for file location and line numbers in stack trace
	for _, line := range cleanLines {
		if strings.Contains(line, "ERROR") || strings.Contains(line, "panic") || strings.Contains(line, "Exception") {
			result.ErrorPattern = line
		}

		matches := p.fileLocRegex.FindStringSubmatch(line)
		if len(matches) >= 4 {
			result.FileLocation = matches[1]
			fmt.Sscanf(matches[3], "%d", &result.LineNumber)
		}
	}

	return result
}
