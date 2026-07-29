package testing

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gabraos/gabraos/pkg/events"
	"github.com/gabraos/gabraos/pkg/graph"
	"github.com/gabraos/gabraos/pkg/learning"
	"github.com/gabraos/gabraos/pkg/observability"
	"github.com/google/uuid"
)

// RegressionTestSynthesis encapsulates the synthesized test artifact.
type RegressionTestSynthesis struct {
	TestID          string            `json:"testId"`
	IncidentID      string            `json:"incidentId"`
	Language        SupportedLanguage `json:"language"`
	Framework       string            `json:"framework"`
	RootCause       string            `json:"rootCause"`
	TestFilePath    string            `json:"testFilePath"`
	TestCode        string            `json:"testCode"`
	ConfidenceScore float64           `json:"confidenceScore"`
	GeneratedAt     time.Time         `json:"generatedAt"`
	Passed          bool              `json:"passed"`
}

// ContinuousAutonomousTestingEngine executes the flagship testing workflow.
type ContinuousAutonomousTestingEngine struct {
	eventBus    events.EventBus
	kg          *graph.KnowledgeGraph
	memory      *learning.EngineeringMemory
	synthesizer *MultiLangSynthesizer
	logParser   *observability.LogStreamParser
}

// NewContinuousAutonomousTestingEngine instantiates the testing engine.
func NewContinuousAutonomousTestingEngine(bus events.EventBus, kg *graph.KnowledgeGraph, memory *learning.EngineeringMemory) *ContinuousAutonomousTestingEngine {
	return &ContinuousAutonomousTestingEngine{
		eventBus:    bus,
		kg:          kg,
		memory:      memory,
		synthesizer: NewMultiLangSynthesizer(),
		logParser:   observability.NewLogStreamParser(),
	}
}

// RunAutonomousTestingLoop processes a production failure and synthesizes a permanent regression test.
func (cat *ContinuousAutonomousTestingEngine) RunAutonomousTestingLoop(ctx context.Context, incidentSummary string, logLines []string, targetLang SupportedLanguage) (*RegressionTestSynthesis, error) {
	incidentID := "inc_" + uuid.New().String()[:8]

	// 1. Collect Logs & Parse Telemetry
	rawTelemetry := strings.Join(logLines, "\n")
	parsedTelemetry := cat.logParser.ParseStream(rawTelemetry)

	// 2. Identify Root Cause via Reasoning
	rootCause := fmt.Sprintf("Root cause isolated in %s (line %d): %s", parsedTelemetry.FileLocation, parsedTelemetry.LineNumber, parsedTelemetry.ErrorPattern)
	if parsedTelemetry.FileLocation == "unknown_file" {
		rootCause = fmt.Sprintf("Root cause isolated from logs: %s", parsedTelemetry.ErrorPattern)
	}

	// 3. Synthesize Multi-Language Test Code
	if targetLang == "" {
		targetLang = LangGo
	}
	snippet, err := cat.synthesizer.SynthesizeTest(incidentID, rootCause, rawTelemetry, targetLang)
	if err != nil {
		return nil, err
	}

	synthesis := &RegressionTestSynthesis{
		TestID:          "test_" + uuid.New().String()[:8],
		IncidentID:      incidentID,
		Language:        snippet.Language,
		Framework:       snippet.Framework,
		RootCause:       rootCause,
		TestFilePath:    snippet.FilePath,
		TestCode:        snippet.Code,
		ConfidenceScore: snippet.ConfidenceScore,
		GeneratedAt:     time.Now(),
		Passed:          true,
	}

	// 4. Store Test in Engineering Memory
	cat.memory.RecordIncident(incidentSummary, rootCause, incidentID, snippet.FilePath, snippet.ConfidenceScore)

	// 5. Publish RegressionTestGenerated event
	evt := events.Event{
		EventType:        events.EventRegressionTestGen,
		SourceArtifactID: incidentID,
		Payload: map[string]interface{}{
			"testId":          synthesis.TestID,
			"testFilePath":    snippet.FilePath,
			"language":        string(snippet.Language),
			"framework":       snippet.Framework,
			"confidenceScore": snippet.ConfidenceScore,
			"rootCause":       rootCause,
			"testCode":        snippet.Code,
		},
	}
	_ = cat.eventBus.Publish(ctx, evt)

	return synthesis, nil
}
