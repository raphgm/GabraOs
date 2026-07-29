# GabraOS — Knowledge Graph Architecture & Specification

GabraOS stores engineering artifacts not as isolated database tables, but as an interconnected **Engineering Knowledge Graph**. This allows AI agents to reason across line-of-code changes, container builds, model deployments, runtime logs, synthesized tests, and organizational knowledge.

---

## 1. Graph Topology & Edge Relationships

```text
Repository
   │  (CONTAINS_COMMIT)
   ▼
Commit ◄─── (DERIVED_FROM) ─── Prompt / Model / Dataset
   │  (PRODUCES_BUILD)
   ▼
Build
   │  (PACKAGED_INTO)
   ▼
Container
   │  (DEPLOYED_AS)
   ▼
Deployment ─── (PRODUCES_METRICS) ───► Business KPI
   │  (EMITS_LOGS)
   ▼
Incident
   │  (TRIGGERS_ANALYSIS)
   ▼
Root Cause Node
   │  (SYNTHESIZES)
   ▼
Regression Test
   │  (INCORPORATED_INTO)
   ▼
Knowledge (Engineering Memory)
```

---

## 2. Standard Edge Types

| Subject | Edge Type | Object | Description |
| :--- | :--- | :--- | :--- |
| `Repository` | `HAS_COMMIT` | `Commit` | Links repository to a code commit |
| `Commit` | `TRIGGERS_BUILD` | `Build` | Links commit to the generated build |
| `Build` | `PRODUCES_CONTAINER` | `Container` | Links build to the compiled container image |
| `Container` | `INSTANTIATED_IN` | `Deployment` | Links image to active runtime deployment |
| `Deployment` | `USES_MODEL` | `Model` | Links active deployment to AI inference endpoint |
| `Deployment` | `USES_PROMPT` | `Prompt` | Links deployment to system prompt version |
| `Deployment` | `IMPACTS_KPI` | `BusinessKPI` | Connects technical release to business impact |
| `Deployment` | `EXPERIENCED` | `Incident` | Links runtime failure to deployment artifact |
| `Incident` | `RESOLVED_BY` | `Knowledge` | Links outage to synthesized regression test/lesson |

---

## 3. Knowledge Reasoning Cypher Queries

### Example A: Trace Production Incident Back to Code & Prompt Change
```cypher
MATCH (inc:Incident {id: $incidentId})<-[:EXPERIENCED]-(dep:Deployment)
MATCH (dep)<-[:INSTANTIATED_IN]-(cnt:Container)<-[:PRODUCES_CONTAINER]-(bld:Build)
MATCH (bld)<-[:TRIGGERS_BUILD]-(cmt:Commit)
OPTIONAL MATCH (dep)-[:USES_PROMPT]->(prm:Prompt)
RETURN inc, dep, cnt, bld, cmt, prm
```

### Example B: Identify Untested Production Regression Paths
```cypher
MATCH (p:Policy {name: "ContinuousTestingRequired"})
MATCH (inc:Incident)-[:RESOLVED_BY]->(k:Knowledge)
WHERE NOT (k)-[:HAS_REGRESSION_TEST]->()
RETURN inc.id, inc.summary, k.rootCause
```

---

*GabraOS Knowledge Graph Specifications*
