# Server — Plan

## Role

Go microservice providing REST API, web browsing, LLM summarization, and graph storage.

## Current Implementation

### Project Structure

```
server/
├── cmd/main.go              # Entry point, router setup, graceful shutdown
├── internal/
│   ├── config/config.go     # Env-var configuration loader
│   ├── health/
│   │   ├── handler.go       # Extensible health check handler
│   │   └── handler_test.go  # 4 tests (healthy, degraded, mixed, content-type)
│   └── neo4j/
│       ├── client.go        # Neo4j driver wrapper with connectivity check
│       └── client_test.go   # Interface compliance test
├── go.mod                   # Module: github.com/researchbot/server
└── go.sum
```

### Decisions Made

- **HTTP router:** `chi` (confirmed by implementation)
- **Middleware:** `chi/middleware` Logger, RequestID, RealIP
- **Health checks:** Interface-driven `Checker` pattern — extensible for adding new service checks
- **Server timeouts:** 10s read/write, 60s idle, 10s graceful shutdown drain

### Implemented Features

| Feature | Status | Details |
|---------|--------|---------|
| Configuration | Done | Env vars: `SERVER_ADDR`, `NEO4J_URI`, `NEO4J_USER`, `NEO4J_PASSWORD` |
| Neo4j Client | Done | Driver wrapper, startup connectivity verification, `Check()` for health |
| Health Endpoint | Done | `GET /health` returns JSON status with per-service breakdown |
| Graceful Shutdown | Done | SIGINT/SIGTERM handling with 10s drain |
| Docker Compose | Done | Neo4j service (delboy), ports 7474/7687, volume mounts |

### Implemented Routes

| Method | Endpoint | Handler |
|--------|----------|---------|
| `GET` | `/health` | `health.Handler` — Neo4j connectivity check |

## Ongoing Plan

### Phase 2: Graph Layer

- [ ] Schema initialization — Cypher scripts for node/relationship constraints
- [ ] Location model — create, query, CRUD operations
- [ ] Finding model — create, query, link to locations
- [ ] Task model — lifecycle states, create, update, query
- [ ] Source model — track web sources
- [ ] Graph query endpoint — `GET /api/v1/graph/query` for Cypher execution

### Phase 3: Plugin System

- [ ] Plugin directory structure — `server/plugins/`
- [ ] Plugin JSON schema — queries, targets, prompts, schema extensions
- [ ] Plugin registry — discovery, loading, hot-reload
- [ ] Plugin API endpoints — `GET/POST/DELETE /api/v1/plugins`
- [ ] Plugin validation — schema validation on load

### Phase 4: Web Browsing

- [ ] Playwright integration — decision: Go binding vs Python subprocess
- [ ] URL fetching with headless browser
- [ ] Content extraction — text, metadata, links
- [ ] Rate limiting and retry logic
- [ ] Source tracking — store fetched URLs

### Phase 5: AI Pipeline

- [ ] LLM client abstraction — provider-agnostic interface
- [ ] Summarization — extract key findings from web content
- [ ] Entity extraction — locations, organizations, dates
- [ ] Prompt templating — per-plugin extraction prompts
- [ ] Response caching — avoid redundant LLM calls

### Phase 6: Orchestrator

- [ ] Research pipeline — sequential: resolve plugin → browse → extract → store
- [ ] Task creation endpoint — `POST /api/v1/research`
- [ ] Task status endpoint — `GET /api/v1/research/:id`
- [ ] Findings query endpoint — `GET /api/v1/findings`
- [ ] Error handling and retry at pipeline level

### Phase 8: Polish

- [ ] Request filtering and pagination
- [ ] Structured error responses
- [ ] Request validation
- [ ] Comprehensive test coverage
- [ ] Logging improvements (structured logging)
- [ ] Metrics/observability

## Open Decisions

- Playwright: Go binding (`playwright-go`) vs Python subprocess
- LLM provider: OpenAI API vs local model vs abstracted interface
- Plugin hot-reload: file watcher vs API-triggered reload
