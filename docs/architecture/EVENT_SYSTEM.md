# GabraOS — Event System Specification

In GabraOS, **every action and state change is an event**. Autonomous agents subscribe to NATS event topics rather than calling APIs directly.

---

## Event Catalog

| Event Name | Topic Subject | Emitted By | Primary Consumers |
| :--- | :--- | :--- | :--- |
| `CommitCreated` | `gabra.v1.git.commit.created` | GitHub/GitLab Plugin | Security Agent, Testing Agent |
| `BuildStarted` | `gabra.v1.build.started` | Build Service / CLI | Observability Agent |
| `BuildCompleted` | `gabra.v1.build.completed` | Build Service | Deployment Agent |
| `DeploymentStarted` | `gabra.v1.deploy.started` | Deployment Service | Observability Agent |
| `DeploymentSucceeded` | `gabra.v1.deploy.succeeded` | K8s Plugin / Cloud | Testing Agent, Cost Agent |
| `PerformanceDropped` | `gabra.v1.obs.perf.dropped` | Observability Agent | Incident Agent |
| `IncidentDetected` | `gabra.v1.incident.detected` | Observability / Loki | Incident Agent, Testing Agent |
| `RootCauseFound` | `gabra.v1.incident.rootcause` | Incident Agent | Testing Agent |
| `RegressionTestGenerated` | `gabra.v1.test.generated` | Testing Agent | Engineering Memory, CLI |
| `KnowledgeUpdated` | `gabra.v1.memory.updated` | Memory Module | Reasoning Engine |
| `DeploymentRolledBack` | `gabra.v1.deploy.rolledback` | Guardrail Engine | Notification Service, CLI |

---

## Event Payload Schemas

### `IncidentDetected` Event
```json
{
  "eventId": "evt_77a1b2c3",
  "eventType": "IncidentDetected",
  "timestamp": "2026-07-29T08:00:00Z",
  "sourceArtifactId": "dep_prod_payment_v142",
  "payload": {
    "severity": "CRITICAL",
    "errorPattern": "NullPointerException in StripeWebhookHandler.java:142",
    "affectedService": "payment-service",
    "logSnippets": [
      "2026-07-29T07:59:58Z ERROR payment-service: Failed to parse charge.succeeded webhook payload: key 'customer_id' is null"
    ],
    "metricsTrace": {
      "errorRatePercent": 14.8,
      "latencyP99Ms": 4200
    }
  }
}
```

### `RegressionTestGenerated` Event
```json
{
  "eventId": "evt_88b2c3d4",
  "eventType": "RegressionTestGenerated",
  "timestamp": "2026-07-29T08:02:15Z",
  "sourceArtifactId": "inc_stripe_webhook_null_ptr",
  "payload": {
    "testId": "test_stripe_null_customer_id",
    "testLanguage": "go",
    "testFilePath": "pkg/testing/generated/stripe_webhook_test.go",
    "confidenceScore": 0.98,
    "testCode": "func TestStripeWebhookMissingCustomerId(t *testing.T) { ... }"
  }
}
```

---

*GabraOS Event System Specification*
