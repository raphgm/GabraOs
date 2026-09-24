'use me';
'use client';

import React, { useState } from 'react';
import { 
  ShieldCheck, 
  Activity, 
  Cpu, 
  GitBranch, 
  Zap, 
  Database, 
  BrainCircuit,
  FileCode,
  CheckCircle2,
  AlertTriangle,
  GitPullRequest,
  Lock,
  LineChart,
  Vote
} from 'lucide-react';

export default function EngineeringControlCenter() {
  const [activeTab, setActiveTab] = useState<'overview' | 'graph' | 'agents' | 'testing' | 'patch' | 'consensus' | 'entropy'>('overview');
  const [selectedLang, setSelectedLang] = useState<'go' | 'python' | 'typescript' | 'java'>('go');

  const patchDiffs = {
    go: `--- a/pkg/service/StripeWebhookHandler.go
+++ b/pkg/service/StripeWebhookHandler.go
@@ -140,6 +140,9 @@ func (h *StripeWebhookHandler) HandleEvent(w http.ResponseWriter, r *http.Request
-	customerID := payload["customer"].(map[string]interface{})["id"].(string)
+	customerMap, ok := payload["customer"].(map[string]interface{})
+	if !ok || customerMap["id"] == nil {
+		http.Error(w, "missing required customer ID", http.StatusBadRequest)
+		return
+	}
+	customerID := customerMap["id"].(string)`,
    python: `--- a/services/payment_handler.py
+++ b/services/payment_handler.py
@@ -14,6 +14,8 @@ def process_payload(request_data):
-    customer_id = request_data["customer_id"]
+    if not request_data or "customer_id" not in request_data:
+        raise ValueError("Payload missing required key 'customer_id'")
+    customer_id = request_data["customer_id"]
     return StripeClient.charge(customer_id)`,
    typescript: `--- a/src/handlers/stripeWebhook.ts
+++ b/src/handlers/stripeWebhook.ts
@@ -28,5 +28,8 @@ export async function handleWebhook(event: WebhookEvent) {
-  const customerId = event.payload.customer.id;
+  if (!event?.payload?.customer?.id) {
+    throw new Error('Invalid webhook payload: missing customer.id');
+  }
+  const customerId = event.payload.customer.id;
   await processPayment(customerId);`,
    java: `--- a/src/main/java/com/gabraos/service/StripeWebhookHandler.java
+++ b/src/main/java/com/gabraos/service/StripeWebhookHandler.java
@@ -42,6 +42,9 @@ public class StripeWebhookHandler {
     public Response handle(String payload) {
-        String customerId = jsonNode.get("customer").get("id").asText();
+        if (jsonNode == null || !jsonNode.has("customer") || jsonNode.get("customer").isNull()) {
+            throw new IllegalArgumentException("Payload missing required customer object");
+        }
+        String customerId = jsonNode.get("customer").get("id").asText();
         return Response.ok(service.charge(customerId)).build();
     }`
  };

  return (
    <div className="min-h-screen bg-[#080c14] text-slate-100 flex flex-col">
      {/* Top Navbar */}
      <header className="border-b border-slate-800 bg-[#0d1322]/80 backdrop-blur-md px-6 py-4 flex justify-between items-center sticky top-0 z-50">
        <div className="flex items-center gap-3">
          <div className="h-10 w-10 rounded-xl bg-gradient-to-tr from-indigo-600 to-violet-500 flex items-center justify-center shadow-lg shadow-indigo-500/30">
            <BrainCircuit className="h-6 w-6 text-white" />
          </div>
          <div>
            <h1 className="text-xl font-bold tracking-tight bg-gradient-to-r from-white via-slate-200 to-indigo-300 bg-clip-text text-transparent">
              GabraOS
            </h1>
            <p className="text-xs text-indigo-400 font-mono">Novelty Engine v0.3 — AST & Consensus</p>
          </div>
        </div>

        {/* Global System Health Indicator */}
        <div className="flex items-center gap-6">
          <div className="flex items-center gap-2 bg-emerald-950/60 border border-emerald-500/30 px-3 py-1.5 rounded-full">
            <div className="h-2.5 w-2.5 rounded-full bg-emerald-400 animate-pulse" />
            <span className="text-xs font-semibold text-emerald-300">Production Status: 99.98% HEALTHY</span>
          </div>

          <div className="flex items-center gap-2 bg-slate-800/80 px-3 py-1.5 rounded-lg border border-slate-700 text-xs font-mono text-slate-300">
            <ShieldCheck className="h-4 w-4 text-indigo-400" />
            <span>OPA & BFT Quorum Enforced</span>
          </div>
        </div>
      </header>

      {/* Main Control Center Body */}
      <div className="flex-1 p-6 max-w-7xl mx-auto w-full space-y-6">
        
        {/* Navigation Tabs */}
        <div className="flex flex-wrap gap-2 border-b border-slate-800 pb-2">
          {[
            { id: 'overview', label: 'Overview' },
            { id: 'patch', label: '⚡ AST Code Patch Synthesizer' },
            { id: 'consensus', label: '🗳 Multi-Agent BFT Consensus' },
            { id: 'entropy', label: '📊 AI Semantic Entropy & Drift' },
            { id: 'graph', label: 'Knowledge Graph' },
            { id: 'agents', label: 'Agent Runtime' },
            { id: 'testing', label: 'Continuous Testing' },
          ].map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id as any)}
              className={`px-4 py-2 text-xs font-medium rounded-lg transition-all ${
                activeTab === tab.id
                  ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-600/30 font-semibold'
                  : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800/50'
              }`}
            >
              {tab.label}
            </button>
          ))}
        </div>

        {/* TAB: AST CODE PATCH SYNTHESIZER */}
        {activeTab === 'patch' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <h3 className="text-lg font-bold text-white flex items-center gap-2">
                  <GitPullRequest className="h-5 w-5 text-emerald-400" />
                  AST-Aware Zero-Downtime Patch Synthesizer
                </h3>
                <p className="text-xs text-slate-400 mt-0.5">
                  Generates AST-verified source code fix diffs (.patch) alongside regression tests.
                </p>
              </div>

              {/* Language Selector */}
              <div className="flex gap-2 bg-slate-900 p-1 rounded-xl border border-slate-800">
                {(['go', 'python', 'typescript', 'java'] as const).map((lang) => (
                  <button
                    key={lang}
                    onClick={() => setSelectedLang(lang)}
                    className={`px-3 py-1 text-xs font-semibold rounded-lg capitalize ${
                      selectedLang === lang ? 'bg-indigo-600 text-white' : 'text-slate-400'
                    }`}
                  >
                    {lang}
                  </button>
                ))}
              </div>
            </div>

            <div className="bg-slate-950 p-4 rounded-xl border border-slate-800 font-mono text-xs space-y-3">
              <div className="flex justify-between text-slate-400 border-b border-slate-800 pb-2">
                <span>Incident: <strong className="text-indigo-400">inc_prod_crash_142</strong></span>
                <span>Target: <strong className="text-slate-200">StripeWebhookHandler</strong></span>
                <span>Confidence: <strong className="text-emerald-400">97.0%</strong></span>
              </div>
              <pre className="text-emerald-400 bg-slate-900/90 p-4 rounded-lg border border-slate-800 overflow-x-auto">
                {patchDiffs[selectedLang]}
              </pre>
            </div>
          </div>
        )}

        {/* TAB: MULTI-AGENT BFT CONSENSUS */}
        {activeTab === 'consensus' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <h3 className="text-lg font-bold text-white flex items-center gap-2">
                  <Vote className="h-5 w-5 text-indigo-400" />
                  Multi-Agent BFT Consensus Matrix
                </h3>
                <p className="text-xs text-slate-400">
                  4 specialized agents evaluate risk vectors and vote with 75.0% Quorum requirement.
                </p>
              </div>
              <span className="bg-emerald-950 border border-emerald-500/40 text-emerald-300 px-3 py-1 rounded-full text-xs font-bold">
                VERDICT: APPROVED (87.5% Consensus)
              </span>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
              {[
                { agent: 'Security Agent', vote: 'APPROVE', weight: '30%', risk: '10.0', reason: 'No AST injection vectors detected.' },
                { agent: 'Observability Agent', vote: 'APPROVE', weight: '25%', risk: '15.0', reason: 'Latency impact estimated <2ms.' },
                { agent: 'Testing Agent', vote: 'APPROVE', weight: '25%', risk: '5.0', reason: '100% assertions passed.' },
                { agent: 'Cost Agent', vote: 'APPROVE', weight: '20%', risk: '0.0', reason: 'Zero token cost delta.' },
              ].map((v) => (
                <div key={v.agent} className="bg-slate-900/80 p-4 rounded-xl border border-slate-800 space-y-2">
                  <div className="font-bold text-slate-200 text-xs">{v.agent}</div>
                  <div className="flex justify-between items-center text-xs">
                    <span className="text-emerald-400 font-bold">{v.vote}</span>
                    <span className="text-slate-400 font-mono">Weight: {v.weight}</span>
                  </div>
                  <p className="text-[11px] text-slate-400">{v.reason}</p>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* TAB: SEMANTIC ENTROPY */}
        {activeTab === 'entropy' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <h3 className="text-lg font-bold text-white flex items-center gap-2">
              <LineChart className="h-5 w-5 text-violet-400" />
              AI Semantic Entropy & Vector Drift Telemetry
            </h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div className="bg-slate-900/80 p-4 rounded-xl border border-slate-800">
                <div className="text-xs text-slate-400 uppercase">Semantic Entropy Index</div>
                <div className="text-3xl font-extrabold text-indigo-400 mt-1">0.04</div>
                <div className="text-[11px] text-emerald-400 mt-1">Status: Stable (Threshold &lt; 0.25)</div>
              </div>
              <div className="bg-slate-900/80 p-4 rounded-xl border border-slate-800">
                <div className="text-xs text-slate-400 uppercase">Hallucination Risk</div>
                <div className="text-3xl font-extrabold text-emerald-400 mt-1">1.2%</div>
                <div className="text-[11px] text-slate-400 mt-1">Evaluated across 10,000 requests</div>
              </div>
              <div className="bg-slate-900/80 p-4 rounded-xl border border-slate-800">
                <div className="text-xs text-slate-400 uppercase">RAG Context Precision</div>
                <div className="text-3xl font-extrabold text-violet-400 mt-1">98.6%</div>
                <div className="text-[11px] text-slate-400 mt-1">Embedding Cosine Shift: 0.02</div>
              </div>
            </div>
          </div>
        )}

        {/* TAB OVERVIEW */}
        {activeTab === 'overview' && (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">Production Health</h3>
                <Activity className="h-5 w-5 text-emerald-400" />
              </div>
              <div className="flex items-baseline gap-2">
                <span className="text-4xl font-extrabold text-white">99.98%</span>
                <span className="text-xs text-emerald-400 font-semibold">Healthy</span>
              </div>
            </div>
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">BFT Agent Consensus</h3>
                <Vote className="h-5 w-5 text-indigo-400" />
              </div>
              <div className="flex items-baseline gap-2">
                <span className="text-4xl font-extrabold text-indigo-400">87.5%</span>
                <span className="text-xs text-emerald-400 font-semibold">Quorum Met</span>
              </div>
            </div>
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">Semantic Entropy</h3>
                <LineChart className="h-5 w-5 text-violet-400" />
              </div>
              <div className="flex items-baseline gap-2">
                <span className="text-4xl font-extrabold text-violet-400">0.04</span>
                <span className="text-xs text-emerald-400 font-semibold">Stable</span>
              </div>
            </div>
          </div>
        )}

      </div>
    </div>
  );
}
