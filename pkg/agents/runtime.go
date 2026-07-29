package agents

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gabraos/gabraos/pkg/events"
	"github.com/gabraos/gabraos/pkg/graph"
	"github.com/gabraos/gabraos/pkg/policy"
)

// AgentRole defines the domain responsibility of an autonomous agent.
type AgentRole string

const (
	RoleTesting       AgentRole = "TestingAgent"
	RoleSecurity      AgentRole = "SecurityAgent"
	RoleObservability AgentRole = "ObservabilityAgent"
	RoleIncident      AgentRole = "IncidentAgent"
	RoleCost          AgentRole = "CostAgent"
)

// AgentLifecycleStage tracks current execution stage.
type AgentLifecycleStage string

const (
	StageReceiveEvent    AgentLifecycleStage = "ReceiveEvent"
	StageAnalyzeContext  AgentLifecycleStage = "AnalyzeContext"
	StageReason          AgentLifecycleStage = "Reason"
	StageRecommendAction AgentLifecycleStage = "RecommendAction"
	StagePublishEvent    AgentLifecycleStage = "PublishEvent"
)

// AgentRecommendation contains reasoned decision output.
type AgentRecommendation struct {
	AgentID         string                 `json:"agentId"`
	Role            AgentRole              `json:"role"`
	Action          string                 `json:"action"`
	Reasoning       string                 `json:"reasoning"`
	ConfidenceScore float64                `json:"confidenceScore"`
	Payload         map[string]interface{} `json:"payload"`
}

// Agent represents an autonomous agent entity.
type Agent struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Role         AgentRole           `json:"role"`
	Status       string              `json:"status"`
	CurrentStage AgentLifecycleStage `json:"currentStage"`
	SubscribedTo []events.EventType  `json:"subscribedTo"`
	LastActive   time.Time           `json:"lastActive"`
}

// Runtime orchestrates all active agents, event bus routing, and policy guardrails.
type Runtime struct {
	mu         sync.RWMutex
	agents     map[string]*Agent
	eventBus   events.EventBus
	kg         *graph.KnowledgeGraph
	guardrails *policy.GuardrailEngine
}

// NewRuntime instantiates the GabraOS Agent Runtime.
func NewRuntime(bus events.EventBus, kg *graph.KnowledgeGraph) *Runtime {
	rt := &Runtime{
		agents:     make(map[string]*Agent),
		eventBus:   bus,
		kg:         kg,
		guardrails: policy.NewGuardrailEngine(),
	}

	// Register default autonomous core agents
	rt.RegisterAgent("agt_test_01", "Continuous Testing Agent", RoleTesting, []events.EventType{events.EventRootCauseFound, events.EventBuildCompleted})
	rt.RegisterAgent("agt_sec_01", "Security & Guardrails Agent", RoleSecurity, []events.EventType{events.EventCommitCreated, events.EventDeploymentStarted})
	rt.RegisterAgent("agt_obs_01", "Observability Telemetry Agent", RoleObservability, []events.EventType{events.EventDeploymentSucceeded, events.EventPerformanceDropped})
	rt.RegisterAgent("agt_inc_01", "Incident & Root Cause Agent", RoleIncident, []events.EventType{events.EventIncidentDetected})
	rt.RegisterAgent("agt_cost_01", "Financial & Token Cost Agent", RoleCost, []events.EventType{events.EventDeploymentSucceeded})

	return rt
}

// RegisterAgent adds a new agent to the runtime and hooks event subscriptions.
func (rt *Runtime) RegisterAgent(id, name string, role AgentRole, topics []events.EventType) *Agent {
	rt.mu.Lock()
	defer rt.mu.Unlock()

	agent := &Agent{
		ID:           id,
		Name:         name,
		Role:         role,
		Status:       "Idle",
		CurrentStage: StageReceiveEvent,
		SubscribedTo: topics,
		LastActive:   time.Now(),
	}

	rt.agents[id] = agent

	// Subscribe agent to topics
	for _, t := range topics {
		topic := t
		rt.eventBus.Subscribe(topic, func(ctx context.Context, evt events.Event) error {
			return rt.ExecuteLifecycle(ctx, agent, evt)
		})
	}

	return agent
}

// ExecuteLifecycle enforces the 5-stage lifecycle for an agent.
func (rt *Runtime) ExecuteLifecycle(ctx context.Context, agent *Agent, evt events.Event) error {
	rt.mu.Lock()
	agent.Status = "Running"
	agent.LastActive = time.Now()
	rt.mu.Unlock()

	// Stage 1: Receive Event
	agent.CurrentStage = StageReceiveEvent

	// Stage 2: Analyze Context
	agent.CurrentStage = StageAnalyzeContext
	_ = fmt.Sprintf("Analyzing context for artifact %s in Knowledge Graph", evt.SourceArtifactID)

	// Stage 3: Reason
	agent.CurrentStage = StageReason
	reasoningMsg := fmt.Sprintf("Agent %s processed event %s and determined system trajectory", agent.Name, evt.EventType)

	// Stage 4: Recommend Action
	agent.CurrentStage = StageRecommendAction
	recommendation := AgentRecommendation{
		AgentID:         agent.ID,
		Role:            agent.Role,
		Action:          "ExecuteAutonomousTask",
		Reasoning:       reasoningMsg,
		ConfidenceScore: 0.96,
		Payload:         evt.Payload,
	}

	// Stage 5: Publish Event
	agent.CurrentStage = StagePublishEvent
	outEvt := events.Event{
		EventType:        events.EventKnowledgeUpdated,
		SourceArtifactID: evt.SourceArtifactID,
		Payload: map[string]interface{}{
			"agentId":        agent.ID,
			"recommendation": recommendation,
		},
	}
	_ = rt.eventBus.Publish(ctx, outEvt)

	rt.mu.Lock()
	agent.Status = "Idle"
	agent.CurrentStage = StageReceiveEvent
	rt.mu.Unlock()

	return nil
}

// ListAgents returns all registered agents.
func (rt *Runtime) ListAgents() []*Agent {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	list := make([]*Agent, 0, len(rt.agents))
	for _, a := range rt.agents {
		list = append(list, a)
	}
	return list
}
