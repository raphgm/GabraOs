# GabraOS — Architecture Specification

This document provides the architectural specification for **GabraOS**, the Autonomous Engineering Operating System. 

---

## 1. Overall System Architecture

GabraOS consists of five primary layers:
1. **Control Interface**: Web Control Center & CLI (`gabra`).
2. **Autonomous Agent Runtime**: Event-driven AI agents (Testing, Security, Observability, Incident, Cost).
3. **Core Platform Engine**: Event Bus (NATS), Knowledge Graph (Neo4j), Artifact Store, Policy Engine (OPA), Engineering Memory.
4. **Integration & Plugin SDK**: Ecosystem connectors (Kubernetes, GitHub, AWS, OpenAI, etc.).
5. **Target Infrastructure & Observability**: OpenTelemetry, Prometheus, Loki, Cloud infrastructure.

```mermaid
flowchart TB
    subgraph UI["Control Interface Layer"]
        CLI["gabra CLI"]
        WEB["Engineering Control Center (Web)"]
    end

    subgraph API["API Gateway & Control Plane"]
        GRPC["gRPC Server (:50051)"]
        REST["REST API Server (:8080)"]
    end

    subgraph CORE["Gabra Core Engine"]
        BUS["NATS Event Bus"]
        RUNTIME["Agent Runtime Engine"]
        GRAPH["Knowledge Graph (Neo4j)"]
        STORE["Artifact Registry"]
        OPA["Policy Engine (OPA)"]
        MEMORY["Engineering Memory"]
    end

    subgraph AGENTS["Autonomous Agent Layer"]
        TEST_AGT["Testing Agent"]
        SEC_AGT["Security Agent"]
        OBS_AGT["Observability Agent"]
        INC_AGT["Incident Agent"]
        COST_AGT["Cost Agent"]
    end

    subgraph SDK["Plugin Ecosystem SDK"]
        K8S_PLG["Kubernetes Plugin"]
        GH_PLG["GitHub Plugin"]
        AWS_PLG["AWS/Cloud Plugin"]
        AI_PLG["LLM/OpenAI Plugin"]
    end

    UI --> API
    API --> CORE
    CORE --> AGENTS
    AGENTS --> BUS
    AGENTS --> GRAPH
    AGENTS --> STORE
    AGENTS --> OPA
    AGENTS --> MEMORY
    CORE --> SDK
```

---

## 2. Agent Architecture

Every autonomous agent in GabraOS executes a standardized 5-phase loop:

```mermaid
stateDiagram-v2
    [*] --> ReceiveEvent
    ReceiveEvent --> AnalyzeContext: Fetch Artifact Lineage & Telemetry
    AnalyzeContext --> Reason: Consult Engineering Memory & LLM Engine
    Reason --> RecommendAction: Calculate Risk & Confidence Scores
    RecommendAction --> PublishEvent: Emit Decision / Execute Guarded Action
    PublishEvent --> [*]
```

### Agent Lifecycle
1. **Receive Event**: Agent subscribes to relevant events on the NATS event bus.
2. **Analyze Context**: Queries the Knowledge Graph to extract lineage, dependencies, past incidents, and artifact states.
3. **Reason**: Synthesizes conclusions using AI models combined with historical patterns stored in Engineering Memory.
4. **Recommend Action**: Evaluates risk against OPA policies and decides whether to execute autonomously or request human approval.
5. **Publish Event**: Emits outcome events (e.g., `RegressionTestGenerated`, `DeploymentRolledBack`) to update system state.

---

## 3. Event-Driven Architecture

GabraOS uses an asynchronous event bus (NATS / internal pub-sub) for decoupled agent communication. Agents do not call each other directly; they subscribe to and publish events.

```mermaid
sequenceDiagram
    participant Prod as Production Environment
    participant Bus as NATS Event Bus
    participant IncidentAgt as Incident Agent
    participant TestAgt as Testing Agent
    participant Graph as Knowledge Graph
    participant Dev as Human Developer

    Prod->>Bus: IncidentDetected (Log error trace)
    Bus->>IncidentAgt: Deliver IncidentDetected
    IncidentAgt->>Graph: Query artifact lineage & root cause
    IncidentAgt->>Bus: RootCauseFound
    Bus->>TestAgt: Deliver RootCauseFound
    TestAgt->>TestAgt: Synthesize Regression Test
    TestAgt->>Bus: RegressionTestGenerated
    Bus->>Graph: Record Test in Engineering Memory
    TestAgt-->>Dev: Notify Developer with test code & confidence score
```

---

## 4. Security & Guardrail Architecture

Security and safety operate via Open Policy Agent (OPA) integration:

```mermaid
flowchart LR
    Agent["Autonomous Agent"] -->|Proposed Action| PolicyEngine["OPA Policy Engine"]
    PolicyEngine -->|Check Risk Score| Guardrail{"Risk < Threshold?"}
    Guardrail -->|Yes| AutoExec["Execute Action Autonomously"]
    Guardrail -->|No| ApprovalQueue["Human Approval Queue"]
    ApprovalQueue -->|Human Approved| AutoExec
    ApprovalQueue -->|Human Rejected| Abort["Cancel Action & Log Reason"]
```

---

## 5. Data Flow Architecture

```text
[Telemetry / Git Webhook / Incident]
               │
               ▼
       [NATS Event Bus]
               │
      ┌────────┴────────┐
      ▼                 ▼
[Artifact Store]   [Knowledge Graph]
      │                 │
      └────────┬────────┘
               ▼
     [Agent Runtime Engine]
               │
        (Reason & Test)
               │
               ▼
 [Engineering Memory & UI Console]
```

---

*GabraOS Architecture Specifications*
