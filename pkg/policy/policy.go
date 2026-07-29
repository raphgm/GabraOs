package policy

import (
	"fmt"

	"github.com/gabraos/gabraos/pkg/artifacts"
)

// PolicyDecision encapsulates an OPA policy evaluation result.
type PolicyDecision struct {
	Allowed          bool      `json:"allowed"`
	ConfidenceScore  float64   `json:"confidenceScore"` // 0.0 - 1.0
	RiskScore        float64   `json:"riskScore"`       // 0.0 - 100.0
	RequiresHumanApp bool      `json:"requiresHumanApproval"`
	Reason           string    `json:"reason"`
	EvaluatedRules   []string  `json:"evaluatedRules"`
}

// GuardrailEngine evaluates autonomous agent proposals against safety policy guardrails.
type GuardrailEngine struct {
	maxAutonomousRiskThreshold float64
	minDeploymentConfidence    float64
}

// NewGuardrailEngine instantiates a GuardrailEngine with standard thresholds.
func NewGuardrailEngine() *GuardrailEngine {
	return &GuardrailEngine{
		maxAutonomousRiskThreshold: 45.0, // Risk > 45 requires human approval
		minDeploymentConfidence:    0.85, // Confidence < 85% triggers caution
	}
}

// EvaluateDeployment assess deployment risk & confidence.
func (ge *GuardrailEngine) EvaluateDeployment(art *artifacts.Artifact, pastIncidents int) PolicyDecision {
	risk := art.RiskScore.Score
	confidence := 0.95

	// Factor past incidents into risk calculation
	if pastIncidents > 0 {
		risk += float64(pastIncidents * 15)
		confidence -= float64(pastIncidents) * 0.10
	}

	requiresHuman := risk > ge.maxAutonomousRiskThreshold || confidence < ge.minDeploymentConfidence

	reason := "Deployment meets all safety guardrails for autonomous execution."
	if requiresHuman {
		reason = fmt.Sprintf("High risk score (%.1f) or low confidence (%.2f) requires human approval.", risk, confidence)
	}

	return PolicyDecision{
		Allowed:          !requiresHuman,
		ConfidenceScore:  confidence,
		RiskScore:        risk,
		RequiresHumanApp: requiresHuman,
		Reason:           reason,
		EvaluatedRules: []string{
			"rule_max_autonomous_risk_check",
			"rule_deployment_confidence_score",
			"rule_incident_history_decay",
		},
	}
}
