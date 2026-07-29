package artifacts

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ArtifactKind represents one of the 17 core GabraOS artifact kinds.
type ArtifactKind string

const (
	KindApplication    ArtifactKind = "Application"
	KindRepository     ArtifactKind = "Repository"
	KindCommit         ArtifactKind = "Commit"
	KindBranch         ArtifactKind = "Branch"
	KindBuild          ArtifactKind = "Build"
	KindContainer      ArtifactKind = "Container"
	KindDeployment     ArtifactKind = "Deployment"
	KindModel          ArtifactKind = "Model"
	KindPrompt         ArtifactKind = "Prompt"
	KindDataset        ArtifactKind = "Dataset"
	KindEmbedding      ArtifactKind = "Embedding"
	KindWorkflow       ArtifactKind = "Workflow"
	KindInfrastructure ArtifactKind = "Infrastructure"
	KindPolicy         ArtifactKind = "Policy"
	KindIncident       ArtifactKind = "Incident"
	KindKnowledge      ArtifactKind = "Knowledge"
	KindBusinessKPI    ArtifactKind = "BusinessKPI"
)

// HealthStatus represents health status of an artifact.
type HealthStatus string

const (
	HealthHealthy  HealthStatus = "Healthy"
	HealthDegraded HealthStatus = "Degraded"
	HealthCritical HealthStatus = "Critical"
	HealthUnknown  HealthStatus = "Unknown"
)

// RiskLevel represents risk evaluation.
type RiskLevel string

const (
	RiskLow      RiskLevel = "Low"
	RiskMedium   RiskLevel = "Medium"
	RiskHigh     RiskLevel = "High"
	RiskCritical RiskLevel = "Critical"
)

// Lineage defines provenance and parental tracking.
type Lineage struct {
	ParentIDs  []string `json:"parentIds"`
	RootID     string   `json:"rootId"`
	Generation int      `json:"generation"`
}

// Health describes health state and score.
type Health struct {
	Status      HealthStatus `json:"status"`
	Score       float64      `json:"score"` // 0 - 100
	LastChecked time.Time    `json:"lastChecked"`
}

// RiskScore encapsulates calculated risk score and risk factors.
type RiskScore struct {
	Score   float64   `json:"score"` // 0 - 100
	Level   RiskLevel `json:"level"`
	Factors []string  `json:"factors"`
}

// Cost describes financial resource and token utilization metrics.
type Cost struct {
	DailyUSD            float64 `json:"dailyUsd"`
	MonthlyEstimatedUSD float64 `json:"monthlyEstimatedUsd"`
	TokenCostUSD        float64 `json:"tokenCostUsd"`
}

// Relationship defines directed connections between artifacts.
type Relationship struct {
	TargetID     string `json:"targetId"`
	RelationType string `json:"relationType"`
}

// Artifact represents a first-class entity in GabraOS.
type Artifact struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Kind          ArtifactKind           `json:"kind"`
	Version       string                 `json:"version"`
	Owner         string                 `json:"owner"`
	Lineage       Lineage                `json:"lineage"`
	Metadata      map[string]interface{} `json:"metadata"`
	Health        Health                 `json:"health"`
	RiskScore     RiskScore              `json:"riskScore"`
	Cost          Cost                   `json:"cost"`
	Relationships []Relationship         `json:"relationships"`
	CreatedAt     time.Time              `json:"createdAt"`
	UpdatedAt     time.Time              `json:"updatedAt"`
}

// NewArtifact instantiates a new GabraOS Artifact with defaults.
func NewArtifact(name string, kind ArtifactKind, version string, owner string) *Artifact {
	id := fmt.Sprintf("art_%s", uuid.New().String()[:8])
	now := time.Now()
	return &Artifact{
		ID:      id,
		Name:    name,
		Kind:    kind,
		Version: version,
		Owner:   owner,
		Lineage: Lineage{
			ParentIDs:  []string{},
			RootID:     id,
			Generation: 1,
		},
		Metadata: make(map[string]interface{}),
		Health: Health{
			Status:      HealthHealthy,
			Score:       100.0,
			LastChecked: now,
		},
		RiskScore: RiskScore{
			Score:   5.0,
			Level:   RiskLow,
			Factors: []string{"Freshly created artifact"},
		},
		Cost: Cost{
			DailyUSD:            0.0,
			MonthlyEstimatedUSD: 0.0,
			TokenCostUSD:        0.0,
		},
		Relationships: []Relationship{},
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// Validate checks artifact integrity.
func (a *Artifact) Validate() error {
	if a.ID == "" {
		return errors.New("artifact ID cannot be empty")
	}
	if a.Name == "" {
		return errors.New("artifact Name cannot be empty")
	}
	if a.Kind == "" {
		return errors.New("artifact Kind cannot be empty")
	}
	return nil
}

// AddRelationship creates a directed edge to another artifact.
func (a *Artifact) AddRelationship(targetID string, relationType string) {
	a.Relationships = append(a.Relationships, Relationship{
		TargetID:     targetID,
		RelationType: relationType,
	})
	a.UpdatedAt = time.Now()
}
