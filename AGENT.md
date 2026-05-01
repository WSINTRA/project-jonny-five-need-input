# ResearchBot — Agent Context

## Project Overview

General-purpose AI research platform. Takes a location + topic, browses the web, generates summaries, stores everything in a Neo4j knowledge graph.

**Services:**
- **Agent** (`agent/`) — TypeScript browser app: plan generation (llama-server), plan orchestration, SSE consumption
- **Server** (`server/`) — Go: REST API, tool execution (Playwright, Google Places), graph writes, SSE streaming

**Shared:** Neo4j Community Edition, llama-server (local LLM inference)

## Branch Workflow

- **`main`** — Stable releases only. No direct pushes.
- **`develop`** — Active development branch. All work happens here.

## Key Files

| File | Purpose |
|------|---------|
| `PLAN.md` | Master architecture plan and phase tracking |
| `PLAN_AGENT.md` | Agent-specific plan (knowledge acquirer architecture) |
| `PLAN_SERVER.md` | Server-specific implementation plan |
| `STATUS_AGENT.md` | Agent progress tracker |
| `STATUS_SERVER.md` | Server progress tracker |
| `docker-compose.yml` | Neo4j, server, PlantUML services |
| `.env` | Root environment config |
| `docs/` | PlantUML diagrams (.puml source + rendered PNG) |

## Key Architecture Decisions

| Decision | Choice |
|----------|--------|
| Graph schema | Hybrid: base schema + plugin extensions |
| Research loop | Declarative plan, then execute |
| Runtime | Browser UI (TS), llama-server direct, Go server for tools |
| Stopping condition | Plan scope + diminishing returns |
| Tool selection | Tool registry + LLM routing |
| Agent-to-graph contract | Typed JSON entities → graph writer |
| Schema definition | Plugin-defined schema |
| Tool execution | Go server endpoints |
| Communication | SSE from Go server → browser |
| Entity dedup | Deterministic key merge + LLM fallback |

## Current State

- **Agent**: TS scaffolding created (Vite, TypeScript). Python agent removed.
- **Server**: Phase 1 complete — health check, Neo4j client, graceful shutdown
- **Diagrams**: PlantUML server on port 7070, architecture diagram in `docs/`

## Running Locally

```bash
# Start infrastructure
docker-compose up -d

# Agent dev server
cd agent && npm install && npm run dev

# Run server
cd server && go run ./cmd/main.go

# Render PlantUML diagram
curl -s http://localhost:7070/png -H 'Content-Type: text/plain' \
  --data-binary @docs/your-diagram.puml \
  --output docs/your-diagram.png
```
