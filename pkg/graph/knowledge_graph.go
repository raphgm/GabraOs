package graph

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gabraos/gabraos/pkg/artifacts"
)

// GraphEdge represents a typed connection between two artifacts.
type GraphEdge struct {
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
	Relation string `json:"relation"`
}

// KnowledgeGraph manages engineering nodes, relationships, and lineage traversal.
type KnowledgeGraph struct {
	mu    sync.RWMutex
	nodes map[string]*artifacts.Artifact
	edges map[string][]GraphEdge
}

// NewKnowledgeGraph initializes a Knowledge Graph instance.
func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		nodes: make(map[string]*artifacts.Artifact),
		edges: make(map[string][]GraphEdge),
	}
}

// AddNode inserts or updates an artifact node in the graph.
func (kg *KnowledgeGraph) AddNode(art *artifacts.Artifact) error {
	if art == nil {
		return errors.New("cannot add nil artifact node")
	}
	kg.mu.Lock()
	defer kg.mu.Unlock()
	kg.nodes[art.ID] = art
	return nil
}

// AddEdge creates a directed relationship between source and target artifact nodes.
func (kg *KnowledgeGraph) AddEdge(sourceID, targetID, relation string) error {
	kg.mu.Lock()
	defer kg.mu.Unlock()

	if _, exists := kg.nodes[sourceID]; !exists {
		return fmt.Errorf("source node %s does not exist", sourceID)
	}
	if _, exists := kg.nodes[targetID]; !exists {
		return fmt.Errorf("target node %s does not exist", targetID)
	}

	edge := GraphEdge{
		SourceID: sourceID,
		TargetID: targetID,
		Relation: relation,
	}
	kg.edges[sourceID] = append(kg.edges[sourceID], edge)
	kg.nodes[sourceID].AddRelationship(targetID, relation)
	return nil
}

// GetNode retrieves an artifact node by ID.
func (kg *KnowledgeGraph) GetNode(id string) (*artifacts.Artifact, bool) {
	kg.mu.RLock()
	defer kg.mu.RUnlock()
	node, exists := kg.nodes[id]
	return node, exists
}

// TraceLineage traverses parental relationships upstream to locate the root cause or commit origin.
func (kg *KnowledgeGraph) TraceLineage(artifactId string) ([]*artifacts.Artifact, error) {
	kg.mu.RLock()
	defer kg.mu.RUnlock()

	var lineagePath []*artifacts.Artifact
	currentID := artifactId

	visited := make(map[string]bool)
	for currentID != "" {
		if visited[currentID] {
			break
		}
		visited[currentID] = true

		node, exists := kg.nodes[currentID]
		if !exists {
			break
		}
		lineagePath = append(lineagePath, node)

		if len(node.Lineage.ParentIDs) > 0 {
			currentID = node.Lineage.ParentIDs[0]
		} else {
			break
		}
	}

	return lineagePath, nil
}

// SearchByKind lists all nodes matching a specific ArtifactKind.
func (kg *KnowledgeGraph) SearchByKind(kind artifacts.ArtifactKind) []*artifacts.Artifact {
	kg.mu.RLock()
	defer kg.mu.RUnlock()

	var results []*artifacts.Artifact
	for _, node := range kg.nodes {
		if node.Kind == kind {
			results = append(results, node)
		}
	}
	return results
}
