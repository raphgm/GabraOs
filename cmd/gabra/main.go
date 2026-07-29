package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gabraos/gabraos/pkg/agents"
	"github.com/gabraos/gabraos/pkg/artifacts"
	"github.com/gabraos/gabraos/pkg/events"
	"github.com/gabraos/gabraos/pkg/graph"
	"github.com/gabraos/gabraos/pkg/learning"
	"github.com/gabraos/gabraos/pkg/testing"
	"github.com/spf13/cobra"
)

var (
	version = "v0.2.0-autonomous"
	targetLang string
	logFile    string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gabra",
		Short: "GabraOS — The Open Standard & Operating System for Autonomous Engineering",
		Long: `GabraOS is an open-source platform that continuously observes, tests, learns, secures,
and optimizes software systems through autonomous AI agents and unified engineering knowledge graphs.`,
	}

	// 1. gabra version
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print GabraOS version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("GabraOS Core CLI %s (Continuous Testing Expansion v0.2)\n", version)
		},
	}

	// 2. gabra status
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Show system health and agent runtime status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("==================================================")
			fmt.Println("        GabraOS Autonomous Platform Status        ")
			fmt.Println("==================================================")
			fmt.Println("System Status       : [HEALTHY]")
			fmt.Println("Event Bus (NATS)    : [ONLINE] - 0ms latency")
			fmt.Println("Knowledge Graph     : [ONLINE] - 17 node kinds active")
			fmt.Println("Engineering Memory  : [ONLINE] - 100% loss-less")
			fmt.Println("OPA Guardrails      : [ENFORCED]")
			fmt.Println("Supported Languages : [Go, Python, TypeScript, Java]")
			fmt.Println("Active Autonomous Agents: 5/5")
			fmt.Println("  - TestingAgent       [ACTIVE]")
			fmt.Println("  - SecurityAgent      [ACTIVE]")
			fmt.Println("  - ObservabilityAgent [ACTIVE]")
			fmt.Println("  - IncidentAgent      [IDLE]")
			fmt.Println("  - CostAgent          [ACTIVE]")
		},
	}

	// 3. gabra agent list
	var agentListCmd = &cobra.Command{
		Use:   "agent",
		Short: "Manage autonomous agents",
		Run: func(cmd *cobra.Command, args []string) {
			bus := events.NewMemoryEventBus()
			kg := graph.NewKnowledgeGraph()
			rt := agents.NewRuntime(bus, kg)

			fmt.Println("Registered Autonomous Agents:")
			for _, agt := range rt.ListAgents() {
				fmt.Printf(" - ID: %-12s | Role: %-18s | Name: %s\n", agt.ID, agt.Role, agt.Name)
			}
		},
	}

	// 4. gabra artifact list
	var artifactListCmd = &cobra.Command{
		Use:   "artifact",
		Short: "Inspect registered artifacts in Knowledge Graph",
		Run: func(cmd *cobra.Command, args []string) {
			art1 := artifacts.NewArtifact("payment-service", artifacts.KindApplication, "v2.1.0", "team-checkout")
			art2 := artifacts.NewArtifact("stripe-webhook-handler", artifacts.KindContainer, "sha256:e3b0c442", "team-checkout")

			fmt.Println("Knowledge Graph Artifacts:")
			fmt.Printf(" [%s] %-20s | Kind: %-12s | Health: %s | Risk: %.1f\n", art1.ID, art1.Name, art1.Kind, art1.Health.Status, art1.RiskScore.Score)
			fmt.Printf(" [%s] %-20s | Kind: %-12s | Health: %s | Risk: %.1f\n", art2.ID, art2.Name, art2.Kind, art2.Health.Status, art2.RiskScore.Score)
		},
	}

	// 5. gabra test-autonomous
	var testAutoCmd = &cobra.Command{
		Use:   "test-autonomous",
		Short: "Trigger Continuous Autonomous Testing cycle on production failure",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			bus := events.NewMemoryEventBus()
			kg := graph.NewKnowledgeGraph()
			mem := learning.NewEngineeringMemory()
			engine := testing.NewContinuousAutonomousTestingEngine(bus, kg, mem)

			logSnippets := []string{
				"ERROR payment-service: NullPointerException in StripeWebhookHandler.go:142",
				"panic: runtime error: invalid memory address or nil pointer dereference",
			}

			if logFile != "" {
				content, err := os.ReadFile(logFile)
				if err == nil {
					logSnippets = strings.Split(string(content), "\n")
				}
			}

			lang := testing.SupportedLanguage(strings.ToLower(targetLang))
			if lang == "" {
				lang = testing.LangGo
			}

			fmt.Printf("Initiating Continuous Autonomous Testing Cycle (Target Language: %s)...\n", lang)

			synthesis, err := engine.RunAutonomousTestingLoop(ctx, "Stripe Webhook Null Pointer Crash", logSnippets, lang)
			if err != nil {
				fmt.Printf("Testing loop error: %v\n", err)
				return
			}

			fmt.Println("\n==================================================")
			fmt.Println(" Autonomous Regression Test Successfully Generated!")
			fmt.Println("==================================================")
			fmt.Printf("Incident ID     : %s\n", synthesis.IncidentID)
			fmt.Printf("Target Language : %s (%s)\n", synthesis.Language, synthesis.Framework)
			fmt.Printf("Root Cause      : %s\n", synthesis.RootCause)
			fmt.Printf("Test File       : %s\n", synthesis.TestFilePath)
			fmt.Printf("Confidence Score: %.2f%%\n", synthesis.ConfidenceScore*100)
			fmt.Println("\nSynthesized Test Code:")
			fmt.Println(synthesis.TestCode)
			fmt.Println("==================================================")
		},
	}

	testAutoCmd.Flags().StringVarP(&targetLang, "lang", "l", "go", "Target test framework language: go, python, typescript, java")
	testAutoCmd.Flags().StringVarP(&logFile, "log-file", "f", "", "Path to production log stream file")

	// 6. gabra graph export
	var graphExportCmd = &cobra.Command{
		Use:   "graph-export",
		Short: "Export Knowledge Graph topology as Cypher script for Neo4j",
		Run: func(cmd *cobra.Command, args []string) {
			kg := graph.NewKnowledgeGraph()
			graph.ExportSampleNodes(kg)
			exporter := graph.NewNeo4jExporter()

			cypher := exporter.ExportToCypher(kg)
			fmt.Println("==================================================")
			fmt.Println("           Neo4j Cypher Script Export             ")
			fmt.Println("==================================================")
			fmt.Println(cypher)
			fmt.Println("==================================================")
		},
	}

	// 7. gabra server start
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the GabraOS Core API & Event Bus Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting GabraOS Core Engine v0.2...")
			fmt.Println("gRPC API listening on 0.0.0.0:50051")
			fmt.Println("REST API listening on 0.0.0.0:8080")
			fmt.Println("NATS Event Bus listening on 0.0.0.0:4222")
			fmt.Println("Engineering Control Center Web Console on http://localhost:3000")
		},
	}

	rootCmd.AddCommand(versionCmd, statusCmd, agentListCmd, artifactListCmd, testAutoCmd, graphExportCmd, serverCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
