# ResearchBot — Plan

## Intent

A general-purpose AI research platform that takes a location + topic,
browses the web for real-world data, generates summaries, and stores
everything as a knowledge graph. Research types are extensible via a
plugin system.

## Architecture

Two microservices sharing a Neo4j knowledge graph:

| Service | Stack | Role |
|---------|-------|------|
| **Server** | Go | REST API, web browsing (Playwright), LLM summarization, graph storage |
| **Agent** | Python + llama.cpp | NLP intent parsing, task orchestration, natural language responses |

**Shared:** Neo4j Community Edition

## Sub-Plans

Each service has its own detailed plan and status tracking:

| Service | Plan | Status |
|---------|------|--------|
| **Server** | [PLAN_SERVER.md](./PLAN_SERVER.md) | [STATUS_SERVER.md](./STATUS_SERVER.md) |
| **Agent** | [PLAN_AGENT.md](./PLAN_AGENT.md) | [STATUS_AGENT.md](./STATUS_AGENT.md) |


## Knowledge Graph

Free-form knowledge graph storing:
- **Locations** — geographic anchors
- **Findings** — individual research results with summaries
- **Tasks** — research job lifecycle
- **Plugins** — research type definitions
- **Sources** — tracked web sources

Relationships connect findings to locations, types, tasks, and sources.

## Plugin System

JSON-based plugins in `server/plugins/` define:
- Search queries (templated with location/filters)
- Target websites
- LLM extraction prompts
- Graph schema extensions

Plugins are hot-reloadable and discoverable via API.

## Research Pipeline

1. Task created (Agent or direct API call)
2. Plugin resolved by research type
3. Playwright browses web for content
4. LLM extracts entities + generates summaries
5. Findings stored in Neo4j graph
6. Task marked complete

## API Surface

| Method | Endpoint | Purpose |
|--------|----------|---------|
| `POST` | `/api/v1/research` | Trigger research |
| `GET` | `/api/v1/research/:id` | Task status & results |
| `GET` | `/api/v1/findings` | Query findings |
| `GET` | `/api/v1/graph/query` | Cypher query |
| `GET/POST/DELETE` | `/api/v1/plugins` | Plugin management |

## Development Approach

Small, testable increments. Each step produces working, verifiable functionality that builds toward the full system.

## Phases

| # | Phase | Outcome |
|---|-------|---------|
| 1 | Foundation | Project structure, docker-compose, health check |
| 2 | Graph Layer | Neo4j client, schema init, CRUD |
| 3 | Plugin System | Registry, JSON loading, API endpoints |
| 4 | Web Browsing | Playwright scraper, content extraction |
| 5 | AI Pipeline | LLM summarization, entity extraction |
| 6 | Orchestrator | End-to-end research pipeline |
| 7 | NLP Agent | llama.cpp agent, intent parsing, graph I/O |
| 8 | Polish | Filtering, pagination, error handling, tests |

## Open Decisions

- Go HTTP router: `gin` vs `chi`
- Playwright: Go binding vs Python subprocess
- Agent interface: HTTP endpoint vs CLI-only
