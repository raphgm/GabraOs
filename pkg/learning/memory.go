package learning

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// IncidentRecord stores historical incident context and resolutions.
type IncidentRecord struct {
	ID                 string    `json:"id"`
	IncidentSummary    string    `json:"incidentSummary"`
	RootCause          string    `json:"rootCause"`
	AffectedArtifactID string    `json:"affectedArtifactId"`
	GeneratedTestPath  string    `json:"generatedTestPath"`
	ResolvedAt         time.Time `json:"resolvedAt"`
	ConfidenceScore    float64   `json:"confidenceScore"`
}

// EngineeringMemory serves as GabraOS's long-term knowledge repository.
type EngineeringMemory struct {
	mu        sync.RWMutex
	incidents map[string]IncidentRecord
	lessons   []string
}

// NewEngineeringMemory instantiates Engineering Memory.
func NewEngineeringMemory() *EngineeringMemory {
	return &EngineeringMemory{
		incidents: make(map[string]IncidentRecord),
		lessons:   make([]string, 0),
	}
}

// RecordIncident stores a resolved incident and lesson into Engineering Memory.
func (em *EngineeringMemory) RecordIncident(summary, rootCause, artifactID, testPath string, confidence float64) IncidentRecord {
	em.mu.Lock()
	defer em.mu.Unlock()

	id := "mem_" + uuid.New().String()[:8]
	rec := IncidentRecord{
		ID:                 id,
		IncidentSummary:    summary,
		RootCause:          rootCause,
		AffectedArtifactID: artifactID,
		GeneratedTestPath:  testPath,
		ResolvedAt:         time.Time{},
		ConfidenceScore:    confidence,
	}

	em.incidents[id] = rec
	em.lessons = append(em.lessons, summary+": "+rootCause)
	return rec
}

// GetIncidents returns all historical incidents stored in memory.
func (em *EngineeringMemory) GetIncidents() []IncidentRecord {
	em.mu.RLock()
	defer em.mu.RUnlock()

	list := make([]IncidentRecord, 0, len(em.incidents))
	for _, inc := range em.incidents {
		list = append(list, inc)
	}
	return list
}

// MemorySize returns the total count of lessons & incidents.
func (em *EngineeringMemory) MemorySize() int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.incidents)
}
