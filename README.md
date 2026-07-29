# GabraOS — The Open Standard for Autonomous Engineering

<p align="center">
  <img src="https://img.shields.io/badge/GabraOS-Autonomous%20Engineering-6366f1?style=for-the-badge&logo=kubernetes&logoColor=white" alt="GabraOS Badge"/>
  <img src="https://img.shields.io/badge/Release-v0.2.0-10b981?style=for-the-badge&logo=github&logoColor=white" alt="Release Badge"/>
  <img src="https://img.shields.io/badge/Core-Go%201.22-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Badge"/>
  <img src="https://img.shields.io/badge/Event%20Bus-NATS-000000?style=for-the-badge&logo=nats&logoColor=white" alt="NATS Badge"/>
  <img src="https://img.shields.io/badge/Graph-Neo4j-0185DA?style=for-the-badge&logo=neo4j&logoColor=white" alt="Neo4j Badge"/>
  <img src="https://img.shields.io/badge/Policy-OPA-000000?style=for-the-badge&logo=openpolicyagent&logoColor=white" alt="OPA Badge"/>
</p>

> **"Don't build software first—build the foundation for a movement."**

**GabraOS** is the open standard and operating system for **Autonomous Engineering**. It moves beyond traditional static CI/CD pipelines to create an ecosystem of autonomous AI agents that continuously observe application runtime, reason over an integrated Knowledge Graph of engineering artifacts, automatically synthesize multi-language regression tests upon production failure, and enforce human-in-the-loop safety guardrails.

---

## 🌟 Why GabraOS?

Traditional DevOps was engineered for **deterministic software**: static code + predictable inputs = expected outputs. 

Modern AI-native distributed applications introduce non-determinism, model drift, prompt variations, vector embedding updates, and cascading microservice failure surfaces. Standard CI/CD tools (Jenkins, GitHub Actions, GitLab CI, ArgoCD) execute static scripts and forget. Next week, the exact same failure reoccurs.

**GabraOS introduces a self-learning paradigm:**
1. **Continuous Autonomous Testing**: When a production failure occurs, GabraOS parses log telemetry, isolates root causes, synthesizes a deterministic regression test across **Go**, **Python** (`pytest`), **TypeScript** (`Jest`), or **Java** (`JUnit 5`), and updates organizational memory.
2. **Everything is an Artifact**: All 17 entities (*Applications, Repositories, Commits, Branches, Builds, Containers, Deployments, Models, Prompts, Datasets, Embeddings, Workflows, Infrastructure, Policies, Incidents, Knowledge, Business KPIs*) exist in a single unified **Engineering Knowledge Graph**.
3. **Guardrails & Risk Scoring**: Open Policy Agent (OPA) safety guardrails ensure AI agents operate within human-defined boundaries while humans approve high-risk decisions.

---

## 🏗 System Architecture

```text
                  ┌────────────────────────┐
                  │ Human Intent & Policy  │
                  └───────────┬────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────┐
│                     GABRA OS CORE                        │
│                                                          │
│  ┌──────────────┐     ┌──────────────┐    ┌───────────┐  │
│  │ Knowledge    │◄───►│ Agent        │◄──►│ Event     │  │
│  │ Graph        │     │ Runtime      │    │ Bus       │  │
│  └──────────────┘     └──────────────┘    └───────────┘  │
└──────────────┬───────────────────┬──────────────┬────────┘
               │                   │              │
               ▼                   ▼              ▼
       ┌───────────────┐   ┌───────────────┐  ┌───────────────┐
       │ Continuous    │   │ Continuous    │  │ Risk-Based    │
       │ Observation   │   │ Auto-Testing  │  │ Autonomous    │
       │ & Telemetry   │   │ Synthesis     │  │ Deployment    │
       └───────────────┘   └───────────────┘  └───────────────┘
```

Detailed design specifications can be found in the documentation:
- 📖 [Manifesto](docs/MANIFESTO.md)
- 🎯 [Vision Document](docs/VISION.md)
- 📜 [Guiding Principles](docs/PRINCIPLES.md)
- 📐 [Architecture Specs](docs/architecture/ARCHITECTURE.md)
- 📦 [Artifact Model](docs/architecture/ARTIFACT_MODEL.md)
- 🕸 [Knowledge Graph](docs/architecture/KNOWLEDGE_GRAPH.md)
- ⚡ [Event System](docs/architecture/EVENT_SYSTEM.md)

---

## 🔄 The 9-Stage Continuous Autonomous Testing Loop

```text
Production Failure ➔ Parse Logs & Telemetry ➔ Isolate Root Cause ➔ Synthesize Test ➔ Verify Execution ➔ Record Engineering Memory ➔ Emit Knowledge Event ➔ Harden Pipeline ➔ Prevent Future Recurrence
```

---

## 📦 The 17 Core Artifact Kinds

| Kind | Description | Primary Attributes |
| :--- | :--- | :--- |
| **Application** | Top-level software product | `tier`, `serviceSLO`, `environment` |
| **Repository** | Git source code reference | `gitUrl`, `defaultBranch`, `language` |
| **Commit** | Immutable code commit | `commitHash`, `author`, `diffStats` |
| **Branch** | Feature or release branch | `branchName`, `aheadCount`, `behindCount` |
| **Build** | Compiled binary build | `buildNumber`, `durationMs`, `compilerFlags` |
| **Container** | OCI / Docker container image | `imageRef`, `digest`, `cveCount` |
| **Deployment** | Running instance in K8s/Cloud | `cluster`, `namespace`, `replicas` |
| **Model** | AI/ML model or provider endpoint | `modelName`, `provider`, `contextWindow` |
| **Prompt** | System prompt or agent template | `templateHash`, `variables`, `version` |
| **Dataset** | Fine-tuning or evaluation dataset | `recordCount`, `hash`, `sourceStorage` |
| **Embedding** | Vector index configuration | `dimensions`, `algorithm`, `vectorStore` |
| **Workflow** | Orchestrated Temporal execution | `workflowId`, `status`, `stepCount` |
| **Infrastructure** | Terraform / K8s manifest resource | `provider`, `resourceType`, `region` |
| **Policy** | Open Policy Agent (OPA) rule | `package`, `enforcementLevel`, `ruleHash` |
| **Incident** | Production outage or anomaly event | `severity`, `rootCauseId`, `timeToResolve` |
| **Knowledge** | Engineering Memory lesson | `testCode`, `incidentId`, `confidenceScore` |
| **BusinessKPI** | Production business outcome | `metricName`, `targetValue`, `currentValue` |

---

## ⚡ Technology Stack

| Layer | Technology |
| :--- | :--- |
| **Core Language** | Go 1.22+ |
| **AI Services** | Python (for model & embedding integrations) |
| **API** | gRPC + REST |
| **CLI** | Cobra |
| **Event Bus** | NATS |
| **Workflow Engine** | Temporal |
| **Graph Database** | Neo4j |
| **Relational Database** | PostgreSQL |
| **Observability** | OpenTelemetry, Prometheus, Loki |
| **Frontend Web Console** | Next.js, React, TailwindCSS |
| **Policy Engine** | Open Policy Agent (OPA) |

---

## 🚀 Quickstart

### Prerequisites
- Go `1.22+` installed
- Node.js `18+` (optional for Web Console)

### 1. Build Core Binaries
```bash
go build -o bin/gabra ./cmd/gabra
go build -o bin/gabra-api ./cmd/gabra-api
```

### 2. Check System Status & Agents
```bash
./bin/gabra status
./bin/gabra agent list
```

### 3. Run Continuous Autonomous Testing Cycle (Multi-Language)
Synthesize regression test suites across **Go**, **Python** (`pytest`), **TypeScript** (`Jest`), and **Java** (`JUnit 5`):

```bash
# Synthesize Go regression test
./bin/gabra test-autonomous --lang go

# Synthesize Python pytest suite
./bin/gabra test-autonomous --lang python

# Synthesize TypeScript Jest suite
./bin/gabra test-autonomous --lang typescript

# Synthesize Java JUnit 5 suite
./bin/gabra test-autonomous --lang java
```

### 4. Export Knowledge Graph to Neo4j Cypher
```bash
./bin/gabra graph-export
```

### 5. Start Core REST API Server
```bash
./bin/gabra-api
```
Available HTTP API Endpoints:
- `GET /api/v1/health` — Platform health check
- `GET /api/v1/status` — Platform & agent runtime status
- `GET /api/v1/agents` — Registered autonomous agents list
- `GET /api/v1/artifacts` — Knowledge graph artifacts
- `GET /api/v1/testing/synthesize?lang=python` — Trigger multi-language test synthesis API
- `GET /api/v1/graph/export` — Export Cypher script for Neo4j

### 6. Run Engineering Control Center (Web)
```bash
cd web
npm install
npm run dev
```
Navigate to `http://localhost:3000` to inspect production health, multi-language test generation, the knowledge graph, and autonomous agent logs.

---

## 🛣 Release Roadmap

- [x] **v0.1 — Foundation**: CLI (`gabra`), API server (`gabra-api`), Event bus, 17 Artifact schema definitions, Agent runtime lifecycle, Continuous Testing flagship loop, Web control center.
- [x] **v0.2 — Continuous Testing Expansion**: Production Loki & OTel log parser, multi-language regression test synthesizer (Go, Python, TS/Jest, Java), Cypher script generator, and Neo4j graph exporter.
- [ ] **v0.3 — Observability**: OpenTelemetry bridge, prompt quality tracking, dataset freshness & embedding drift detector.
- [ ] **v0.4 — Autonomous Reasoning**: Deployment confidence scoring, risk engine, autonomous canary rollouts.
- [ ] **v0.5 — Ecosystem**: Plugin SDK (AWS, GCP, K8s, GitHub, OpenAI, MLflow), plugin marketplace.
- [ ] **v1.0 — Autonomous Engineering Operating System**: Production-ready platform that continuously observes, tests, learns, secures, and optimizes software systems.

---

## 🤝 Community & Contributing

We welcome contributions from engineers, researchers, and maintainers! Please check out [CONTRIBUTING.md](CONTRIBUTING.md) to get started with design RFCs, architectural decision records (ADRs), and pull requests.

---

*GabraOS — Built for the next decade of Autonomous Engineering.*
