'use me';
'use client';

import React, { useState } from 'react';
import { 
  ShieldCheck, 
  Activity, 
  Cpu, 
  GitBranch, 
  AlertTriangle, 
  Zap, 
  Database, 
  CheckCircle2,
  Terminal,
  RefreshCw,
  BrainCircuit,
  Lock,
  DollarSign
} from 'lucide-react';

export default function EngineeringControlCenter() {
  const [activeTab, setActiveTab] = useState<'overview' | 'graph' | 'agents' | 'testing'>('overview');

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
            <p className="text-xs text-indigo-400 font-mono">Autonomous Engineering OS v0.1</p>
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
            <span>OPA Guardrails Enforced</span>
          </div>
        </div>
      </header>

      {/* Main Control Center Body */}
      <div className="flex-1 p-6 max-w-7xl mx-auto w-full space-y-6">
        
        {/* Navigation Tabs */}
        <div className="flex gap-2 border-b border-slate-800 pb-2">
          {[
            { id: 'overview', label: 'Production Overview & Risk' },
            { id: 'graph', label: 'Artifact & Knowledge Graph' },
            { id: 'agents', label: 'Autonomous Agent Runtime' },
            { id: 'testing', label: 'Continuous Testing Feed' },
          ].map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id as any)}
              className={`px-4 py-2 text-sm font-medium rounded-lg transition-all ${
                activeTab === tab.id
                  ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-600/30'
                  : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800/50'
              }`}
            >
              {tab.label}
            </button>
          ))}
        </div>

        {/* TAB 1: PRODUCTION OVERVIEW */}
        {activeTab === 'overview' && (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            
            {/* 1. Is production healthy? */}
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">Production Health</h3>
                <Activity className="h-5 w-5 text-emerald-400" />
              </div>
              <div className="flex items-baseline gap-2">
                <span className="text-4xl font-extrabold text-white">99.98%</span>
                <span className="text-xs text-emerald-400 font-semibold">Healthy</span>
              </div>
              <div className="space-y-2 text-xs text-slate-400">
                <div className="flex justify-between"><span>P99 Latency</span><span className="font-mono text-slate-200">42ms</span></div>
                <div className="flex justify-between"><span>Error Rate</span><span className="font-mono text-emerald-400">0.001%</span></div>
                <div className="flex justify-between"><span>Model Precision</span><span className="font-mono text-indigo-400">98.4%</span></div>
              </div>
            </div>

            {/* 2. What changed? */}
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">Latest Changes</h3>
                <GitBranch className="h-5 w-5 text-indigo-400" />
              </div>
              <div className="space-y-3 text-xs">
                <div className="bg-slate-900/60 p-2.5 rounded-lg border border-slate-800">
                  <div className="font-semibold text-slate-200">Deployment #142 (payment-service)</div>
                  <div className="text-slate-400 mt-1">Commit: 7a8f9b | Container: sha256:e3b0...</div>
                </div>
                <div className="bg-slate-900/60 p-2.5 rounded-lg border border-slate-800">
                  <div className="font-semibold text-slate-200">Prompt Artifact Update v2.1</div>
                  <div className="text-slate-400 mt-1">System prompt optimized for checkout flow</div>
                </div>
              </div>
            </div>

            {/* 3. AI Risk & Confidence Engine */}
            <div className="glass-panel p-5 rounded-2xl border border-slate-800 space-y-4">
              <div className="flex justify-between items-center">
                <h3 className="text-sm font-semibold text-slate-400 uppercase tracking-wider">Deployment Risk Score</h3>
                <ShieldCheck className="h-5 w-5 text-amber-400" />
              </div>
              <div className="flex items-baseline gap-2">
                <span className="text-4xl font-extrabold text-indigo-400">96%</span>
                <span className="text-xs text-indigo-300 font-semibold">Deployment Confidence</span>
              </div>
              <div className="bg-slate-900/80 p-3 rounded-xl border border-slate-800 text-xs text-slate-300">
                <span className="font-semibold text-emerald-400">AI Recommendation:</span> Proceed with automated canary deployment. Database migration matches 3 historical releases.
              </div>
            </div>

          </div>
        )}

        {/* TAB 2: ARTIFACT & KNOWLEDGE GRAPH */}
        {activeTab === 'graph' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <h3 className="text-lg font-bold text-white flex items-center gap-2">
              <Database className="h-5 w-5 text-indigo-400" />
              Engineering Knowledge Graph (17 Artifact Kinds Connected)
            </h3>
            <p className="text-xs text-slate-400">
              GabraOS connects every Commit, Container, Model, Prompt, Incident, and Test into a unified reasoning graph.
            </p>

            <div className="bg-[#05080f] p-6 rounded-xl border border-slate-800 flex flex-wrap justify-between items-center gap-4 text-xs font-mono">
              {['Repository', 'Commit', 'Build', 'Container', 'Deployment', 'Incident', 'RegressionTest', 'Knowledge'].map((node, idx) => (
                <React.Fragment key={node}>
                  <div className="bg-slate-900 border border-indigo-500/40 hover:border-indigo-400 text-indigo-300 px-4 py-3 rounded-xl shadow-md text-center font-bold">
                    {node}
                  </div>
                  {idx < 7 && <span className="text-slate-600 font-bold">➔</span>}
                </React.Fragment>
              ))}
            </div>
          </div>
        )}

        {/* TAB 3: AGENT RUNTIME */}
        {activeTab === 'agents' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <h3 className="text-lg font-bold text-white flex items-center gap-2">
              <Cpu className="h-5 w-5 text-indigo-400" />
              Autonomous Agent Lifecycle Runtime
            </h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {[
                { name: 'Testing Agent', role: 'Continuous Auto-Testing', status: 'Active', stage: 'Reasoning over root cause' },
                { name: 'Security Agent', role: 'OPA Guardrail Enforcement', status: 'Active', stage: 'Evaluating OPA policy' },
                { name: 'Observability Agent', role: 'Telemetry Analysis', status: 'Idle', stage: 'Subscribed to events' },
                { name: 'Incident Agent', role: 'Log & Trace Diagnosis', status: 'Active', stage: 'Listening for anomalies' },
              ].map((agent) => (
                <div key={agent.name} className="bg-slate-900/60 p-4 rounded-xl border border-slate-800 flex items-center justify-between">
                  <div>
                    <div className="font-bold text-slate-200">{agent.name}</div>
                    <div className="text-xs text-slate-400">{agent.role}</div>
                    <div className="text-xs text-indigo-400 mt-1 font-mono">{agent.stage}</div>
                  </div>
                  <span className="bg-emerald-950 text-emerald-400 border border-emerald-500/30 text-xs px-2.5 py-1 rounded-full font-semibold">
                    {agent.status}
                  </span>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* TAB 4: CONTINUOUS TESTING FEED */}
        {activeTab === 'testing' && (
          <div className="glass-panel p-6 rounded-2xl border border-slate-800 space-y-4">
            <h3 className="text-lg font-bold text-white flex items-center gap-2">
              <Zap className="h-5 w-5 text-amber-400" />
              Continuous Autonomous Testing Engine (Version 1)
            </h3>
            <div className="bg-slate-950 p-4 rounded-xl border border-slate-800 font-mono text-xs text-slate-300 space-y-2">
              <div className="text-emerald-400 font-bold">[AUTONOMOUS LOOP TRIGGERED]</div>
              <div className="text-slate-400">1. Production Failure Received ➔ NullPointer in StripeWebhookHandler.go</div>
              <div className="text-slate-400">2. Log Telemetry Analyzed ➔ Isolated root cause to payload missing 'customer_id'</div>
              <div className="text-indigo-400 font-semibold">3. Synthesized Regression Test ➔ tests/autonomous/test_stripe_null_customer_id.go</div>
              <div className="text-emerald-400">4. Test Verification: PASSED (Confidence Score: 98%)</div>
              <div className="text-slate-300">5. Engineering Memory Updated ➔ Permanent hardening against future releases</div>
            </div>
          </div>
        )}

      </div>
    </div>
  );
}
