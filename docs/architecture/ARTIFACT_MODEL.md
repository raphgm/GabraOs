# GabraOS — Core Artifact Model Specification

In GabraOS, **everything is an artifact**. This document specifies the standard schema, core metadata fields, and definitions for all 17 artifact types.

---

## Standard Base Schema

All 17 artifacts inherit from the base schema:

```json
{
  "$schema": "https://gabraos.io/schemas/v1/artifact.json",
  "type": "object",
  "required": [
    "id",
    "name",
    "kind",
    "version",
    "owner",
    "lineage",
    "metadata",
    "health",
    "riskScore",
    "cost",
    "relationships",
    "createdAt",
    "updatedAt"
  ],
  "properties": {
    "id": { "type": "string", "example": "art_9x8f2a1b" },
    "name": { "type": "string", "example": "payment-service" },
    "kind": { 
      "type": "string", 
      "enum": [
        "Application", "Repository", "Commit", "Branch", "Build",
        "Container", "Deployment", "Model", "Prompt", "Dataset",
        "Embedding", "Workflow", "Infrastructure", "Policy",
        "Incident", "Knowledge", "BusinessKPI"
      ]
    },
    "version": { "type": "string", "example": "v1.4.2" },
    "owner": { "type": "string", "example": "team-checkout@gabraos.io" },
    "lineage": {
      "type": "object",
      "properties": {
        "parentIds": { "type": "array", "items": { "type": "string" } },
        "rootId": { "type": "string" },
        "generation": { "type": "integer" }
      }
    },
    "metadata": { "type": "object" },
    "health": {
      "type": "object",
      "properties": {
        "status": { "type": "string", "enum": ["Healthy", "Degraded", "Critical", "Unknown"] },
        "score": { "type": "number", "minimum": 0, "maximum": 100 },
        "lastChecked": { "type": "string", "format": "date-time" }
      }
    },
    "riskScore": {
      "type": "object",
      "properties": {
        "score": { "type": "number", "minimum": 0, "maximum": 100 },
        "level": { "type": "string", "enum": ["Low", "Medium", "High", "Critical"] },
        "factors": { "type": "array", "items": { "type": "string" } }
      }
    },
    "cost": {
      "type": "object",
      "properties": {
        "dailyUsd": { "type": "number" },
        "monthlyEstimatedUsd": { "type": "number" },
        "tokenCostUsd": { "type": "number" }
      }
    },
    "relationships": {
      "type": "array",
      "items": {
        "properties": {
          "targetId": { "type": "string" },
          "relationType": { "type": "string" }
        }
      }
    }
  }
}
```

---

## Specification of 17 Artifact Types

| Kind | Primary Purpose | Key Metadata Attributes |
| :--- | :--- | :--- |
| **Application** | Top-level software service or product | `tier`, `serviceSLO`, `environment` |
| **Repository** | Source code repository reference | `gitUrl`, `defaultBranch`, `language` |
| **Commit** | Immutable code commit snapshot | `commitHash`, `author`, `message`, `diffStats` |
| **Branch** | Feature or release branch | `branchName`, `aheadCount`, `behindCount` |
| **Build** | Compiled binary or artifact build job | `buildNumber`, `durationMs`, `compilerFlags` |
| **Container** | OCI/Docker container image | `imageRef`, `digest`, `cveCount`, `sizeBytes` |
| **Deployment** | Running instance in K8s/Cloud | `cluster`, `namespace`, `replicas`, `strategy` |
| **Model** | AI/ML model weights or provider endpoint | `modelName`, `provider`, `contextWindow`, `quantization` |
| **Prompt** | System prompt or agent template | `templateHash`, `variables`, `version` |
| **Dataset** | Fine-tuning, evaluation, or RAG dataset | `recordCount`, `hash`, `sourceStorage` |
| **Embedding** | Vector index or embedding space configuration | `dimensions`, `algorithm`, `vectorStoreType` |
| **Workflow** | Orchestrated Temporal or pipeline execution | `workflowId`, `status`, `stepCount` |
| **Infrastructure** | Terraform / OpenTofu / K8s manifest resource | `provider`, `resourceType`, `region` |
| **Policy** | Open Policy Agent (OPA) safety guardrail rule | `policyPackage`, `enforcementLevel`, `ruleHash` |
| **Incident** | Production outage or anomaly event | `severity`, `rootCauseId`, `timeToResolveMs` |
| **Knowledge** | Engineering Memory lesson or synthesized regression test | `testCode`, `incidentId`, `confidenceScore` |
| **BusinessKPI** | Production business outcome metric | `metricName`, `targetValue`, `currentValue`, `unit` |

---

*GabraOS Artifact Model Specification*
