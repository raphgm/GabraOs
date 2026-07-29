package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gabraos/gabraos/pkg/agents"
	"github.com/gabraos/gabraos/pkg/artifacts"
	"github.com/gabraos/gabraos/pkg/events"
	"github.com/gabraos/gabraos/pkg/graph"
	"github.com/gabraos/gabraos/pkg/learning"
)

type Server struct {
	bus     events.EventBus
	kg      *graph.KnowledgeGraph
	runtime *agents.Runtime
	memory  *learning.EngineeringMemory
}

func main() {
	bus := events.NewMemoryEventBus()
	kg := graph.NewKnowledgeGraph()
	rt := agents.NewRuntime(bus, kg)
	mem := learning.NewEngineeringMemory()

	server := &Server{
		bus:     bus,
		kg:      kg,
		runtime: rt,
		memory:  mem,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", server.handleHealth)
	mux.HandleFunc("/api/v1/status", server.handleStatus)
	mux.HandleFunc("/api/v1/agents", server.handleAgents)
	mux.HandleFunc("/api/v1/artifacts", server.handleArtifacts)

	addr := ":8080"
	fmt.Printf("GabraOS Core API Server starting on %s...\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":    "Healthy",
		"version":   "v0.1.0-autonomous",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"system":            "GabraOS",
		"status":            "Operational",
		"activeAgents":      len(s.runtime.ListAgents()),
		"memoryIncidents":   s.memory.MemorySize(),
		"guardrailsEnforced": true,
	})
}

func (s *Server) handleAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.runtime.ListAgents())
}

func (s *Server) handleArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	art1 := artifacts.NewArtifact("payment-service", artifacts.KindApplication, "v1.4.2", "team-checkout")
	art2 := artifacts.NewArtifact("stripe-webhook-handler", artifacts.KindContainer, "sha256:e3b0c442", "team-checkout")
	_ = json.NewEncoder(w).Encode([]*artifacts.Artifact{art1, art2})
}
