# ResearchBot — Agent Context

## Project Overview

Two-microservice AI research platform. Takes a location + topic, browses the web, generates summaries, stores everything in a Neo4j knowledge graph. Research types are extensible via plugins.

**Services:**
- **Server** (`server/`) — Go: REST API, web browsing, LLM summarization, graph storage
- **Agent** (`agent/`) — Python + llama.cpp: NLP intent parsing, orchestration, natural language responses

**Shared:** Neo4j Community Edition

## Key Files

| File | Purpose |
|------|---------|
| `PLAN.md` | Master architecture plan and phase tracking |
| `PLAN_AGENT.md` | Agent-specific implementation plan |
| `PLAN_SERVER.md` | Server-specific implementation plan |
| `STATUS_AGENT.md` | Agent progress (all 8 tasks done) |
| `STATUS_SERVER.md` | Server progress (Phase 1 done) |
| `docker-compose.yml` | Neo4j + llama.cpp services |
| `.env` | Root environment config |

## Current State

- **Agent**: Phase 7 complete — 91 tests passing, REPL with fetch/prompt/ask/query/research commands
- **Server**: Phase 1 complete — health check, Neo4j client, graceful shutdown
- **Blocked**: Agent commands (`query`, `research`, `status`) call server API endpoints that don't exist yet (need Phase 2 Graph Layer)

## Phases

| # | Phase | Status |
|---|-------|--------|
| 1 | Foundation | Done |
| 2 | Graph Layer | Not started |
| 3 | Plugin System | Not started |
| 4 | Web Browsing | Not started |
| 5 | AI Pipeline | Not started |
| 6 | Orchestrator | Not started |
| 7 | NLP Agent | Done |
| 8 | Polish | Not started |

## Running Locally

```bash
# Start infrastructure
docker-compose up -d

# Run agent
cd agent && uv run python -m src.main

# Run server
cd server && go run ./cmd/main.go
```
