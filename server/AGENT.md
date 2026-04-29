# Server Service

Go REST API service for web browsing, LLM summarization, graph storage, and plugin management.

## Architecture

```
cmd/
└── main.go              # Entry point — bootstrap, graceful shutdown

internal/
├── config/
│   └── config.go        # Env-var configuration loader
├── health/
│   ├── handler.go       # Extensible health check handler
│   └── handler_test.go  # Health check tests
├── neo4j/
│   ├── client.go        # Neo4j driver wrapper
│   └── client_test.go   # Interface compliance test
└── server/
    ├── server.go        # HTTP server setup, router, middleware, routes
    └── server_test.go   # Server and endpoint tests
```

## Components

### Config (`internal/config/config.go`)
Simple env-var loader. Fields:
- `SERVER_ADDR` — HTTP listen address (default `:9090`)
- `NEO4J_URI`, `NEO4J_USER`, `NEO4J_PASSWORD` — Neo4j connection

### Health Handler (`internal/health/handler.go`)
Interface-driven checker pattern. `Checker` interface with `Check(ctx) error`. Handler runs all checkers, returns JSON with overall status and per-service status.

### Neo4j Client (`internal/neo4j/client.go`)
Wraps `neo4j-go-driver/v6`. Implements `health.Checker` via `Check()` (runs `RETURN 1 AS ping`). Methods:
- `New(uri, user, pass)` — creates driver, verifies connectivity
- `Check(ctx)` — ping query
- `Close(ctx)` — driver cleanup

### Server (`internal/server/server.go`)
Wraps `http.Server` with chi router setup. `New(cfg, checkers)` returns a configured server with:
- Chi router with middleware (logger, request ID, real IP)
- Health endpoint wired via `health.NewHandler`
- Configurable timeouts (10s read/write, 60s idle)
- `registerRoutes()` — single method to add new endpoints as the app grows

### Main (`cmd/main.go`)
Lean bootstrap sequence:
1. Load config
2. Connect to Neo4j
3. Build checkers map
4. Create server via `server.New()`
5. Start HTTP server with graceful shutdown on SIGINT/SIGTERM

## Current API Endpoints

| Method | Endpoint | Status |
|--------|----------|--------|
| `GET` | `/health` | Implemented |
| `POST` | `/api/v1/graph/query` | Not implemented |
| `GET` | `/api/v1/findings` | Not implemented |
| `GET` | `/api/v1/research/:id` | Not implemented |
| `POST` | `/api/v1/research` | Not implemented |
| `GET/POST/DELETE` | `/api/v1/plugins` | Not implemented |

## Dependencies

- `github.com/go-chi/chi/v5` — HTTP router
- `github.com/neo4j/neo4j-go-driver/v6` — Neo4j driver
- Go 1.26.2

## Running

```bash
go run ./cmd/main.go
```

## Tests

```bash
go test ./...
```
