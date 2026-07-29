package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gabraos/gabraos/pkg/agents"
	"github.com/gabraos/gabraos/pkg/artifacts"
	"github.com/gabraos/gabraos/pkg/events"
	"github.com/gabraos/gabraos/pkg/graph"
	"github.com/gabraos/gabraos/pkg/learning"
	"github.com/gabraos/gabraos/pkg/testing"
)

type Server struct {
	bus     events.EventBus
	kg      *graph.KnowledgeGraph
	runtime *agents.Runtime
	memory  *learning.EngineeringMemory
	engine  *testing.ContinuousAutonomousTestingEngine
}

func main() {
	bus := events.NewMemoryEventBus()
	kg := graph.NewKnowledgeGraph()
	graph.ExportSampleNodes(kg)
	rt := agents.NewRuntime(bus, kg)
	mem := learning.NewEngineeringMemory()
	testEngine := testing.NewContinuousAutonomousTestingEngine(bus, kg, mem)

	server := &Server{
		bus:     bus,
		kg:      kg,
		runtime: rt,
		memory:  mem,
		engine:  testEngine,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", server.handleHealth)
	mux.HandleFunc("/api/v1/status", server.handleStatus)
	mux.HandleFunc("/api/v1/agents", server.handleAgents)
	mux.HandleFunc("/api/v1/artifacts", server.handleArtifacts)
	mux.HandleFunc("/api/v1/testing/synthesize", server.handleSynthesize)
	mux.HandleFunc("/api/v1/graph/export", server.handleGraphExport)

	addr := ":8080"
	fmt.Printf("GabraOS Core API Server v0.2 starting on %s...\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":    "Healthy",
		"version":   "v0.2.0-autonomous",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"system":            "GabraOS",
		"version":           "v0.2.0-autonomous",
		"status":            "Operational",
		"activeAgents":      len(s.runtime.ListAgents()),
		"memoryIncidents":   s.memory.MemorySize(),
		"guardrailsEnforced": true,
		"supportedLanguages": []string{"go", "python", "typescript", "java"},
	})
}

func (s *Server) handleAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.runtime.ListAgents())
}

func (s *Server) handleArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	art1 := artifacts.NewArtifact("payment-service", artifacts.KindApplication, "v2.1.0", "team-checkout")
	art2 := artifacts.NewArtifact("stripe-webhook-handler", artifacts.KindContainer, "sha256:e3b0c442", "team-checkout")
	_ = json.NewEncoder(w).Encode([]*artifacts.Artifact{art1, art2})
}

func (s *Server) handleSynthesize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	langQuery := r.URL.Query().Get("lang")
	targetLang := testing.SupportedLanguage(langQuery)
	if targetLang == "" {
		targetLang = testing.LangGo
	}

	logSnippets := []string{
		"ERROR payment-service: NullPointerException in StripeWebhookHandler.go:142",
		"panic: runtime error: invalid memory address or nil pointer dereference",
	}

	synthesis, err := s.engine.RunAutonomousTestingLoop(context.Background(), "Stripe Webhook Null Pointer Crash", logSnippets, targetLang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(synthesis)
}

func (s *Server) handleGraphExport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	exporter := graph.NewNeo4jExporter()
	cypher := exporter.ExportToCypher(s.kg)
	_, _ = w.Write([]byte(cypher))
}
