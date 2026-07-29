package graph

import (
	"fmt"
	"strings"

	"github.com/gabraos/gabraos/pkg/artifacts"
)

// Neo4jExporter exports GabraOS Knowledge Graph state into Cypher creation scripts for external Neo4j instances.
type Neo4jExporter struct{}

// NewNeo4jExporter initializes Neo4jExporter.
func NewNeo4jExporter() *Neo4jExporter {
	return &Neo4jExporter{}
}

// ExportToCypher converts all graph nodes and edges into executable Cypher queries.
func (e *Neo4jExporter) ExportToCypher(kg *KnowledgeGraph) string {
	kg.mu.RLock()
	defer kg.mu.RUnlock()

	var cypherLines []string
	cypherLines = append(cypherLines, "// Auto-generated Cypher script exported by GabraOS Engine")
	cypherLines = append(cypherLines, "CREATE CONSTRAINT IF NOT EXISTS FOR (a:Artifact) REQUIRE a.id IS UNIQUE;")
	cypherLines = append(cypherLines, "")

	// Node creation statements
	for _, node := range kg.nodes {
		line := fmt.Sprintf(
			"MERGE (a:Artifact {id: '%s'}) SET a.name = '%s', a.kind = '%s', a.version = '%s', a.owner = '%s', a.health = '%s', a.riskScore = %.1f;",
			node.ID,
			escapeString(node.Name),
			string(node.Kind),
			node.Version,
			node.Owner,
			string(node.Health.Status),
			node.RiskScore.Score,
		)
		cypherLines = append(cypherLines, line)
	}

	cypherLines = append(cypherLines, "")
	// Edge creation statements
	for sourceID, edges := range kg.edges {
		for _, edge := range edges {
			line := fmt.Sprintf(
				"MATCH (s:Artifact {id: '%s'}), (t:Artifact {id: '%s'}) MERGE (s)-[:%s]->(t);",
				sourceID,
				edge.TargetID,
				strings.ToUpper(edge.Relation),
			)
			cypherLines = append(cypherLines, line)
		}
	}

	return strings.Join(cypherLines, "\n")
}

func escapeString(s string) string {
	return strings.ReplaceAll(s, "'", "\\'")
}

// ExportSampleNodes populates standard initial sample nodes for graph persistence.
func ExportSampleNodes(kg *KnowledgeGraph) {
	app := artifacts.NewArtifact("checkout-service", artifacts.KindApplication, "v2.1.0", "team-checkout")
	repo := artifacts.NewArtifact("gabraos/checkout-service", artifacts.KindRepository, "main", "team-checkout")
	commit := artifacts.NewArtifact("commit_7a8f9b", artifacts.KindCommit, "7a8f9b", "team-checkout")
	container := artifacts.NewArtifact("checkout-container", artifacts.KindContainer, "sha256:e3b0c442", "team-checkout")
	deployment := artifacts.NewArtifact("checkout-prod-deploy", artifacts.KindDeployment, "v2.1.0-k8s", "team-checkout")

	_ = kg.AddNode(app)
	_ = kg.AddNode(repo)
	_ = kg.AddNode(commit)
	_ = kg.AddNode(container)
	_ = kg.AddNode(deployment)

	_ = kg.AddEdge(app.ID, repo.ID, "HAS_REPOSITORY")
	_ = kg.AddEdge(repo.ID, commit.ID, "CONTAINS_COMMIT")
	_ = kg.AddEdge(commit.ID, container.ID, "PRODUCES_CONTAINER")
	_ = kg.AddEdge(container.ID, deployment.ID, "INSTANTIATED_IN")
}
