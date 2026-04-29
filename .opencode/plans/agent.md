# Agent — Plan (Completed)

## Role

Internal interactive agent tool. Not a public API. Provides a REPL shell with
tools for fetching web content, querying the knowledge graph via the server API,
and running LLM inference through an external llama.cpp server.

## Module Boundaries

| Module | Responsibility | Tests |
|--------|---------------|-------|
| `config` | Env loading, pydantic settings model | 14 |
| `llm` | llama.cpp HTTP client, chat/completion | 11 |
| `fetcher` | Playwright fetch, HTML-to-text cleaning | 14 |
| `prompts` | LLM prompt extraction from HTML/text | 19 |
| `neo4j` | Server API bridge for Neo4j queries | 11 |
| `repl` | Interactive shell, command routing, output | 22 |
| `main` | Entry point, wiring modules together | — |

**Total: 91 tests, all passing**

## Completed Tasks

### Task 1: Project Foundation — DONE
- [x] UV project with pyproject.toml, 7 dependencies
- [x] `src/config.py` — pydantic BaseSettings with validation
- [x] `.env.example` with all agent variables
- [x] 14 config tests

### Task 2: LLM Client — DONE
- [x] `src/llm.py` — `chat()` and `complete()` via httpx
- [x] Error handling (connection, timeout, server errors)
- [x] 11 LLM tests

### Task 3: Fetcher Module — DONE
- [x] `src/fetcher.py` — Playwright async browser, HTML cleaning
- [x] `FetchResult` dataclass (text, html, title, url, links)
- [x] Timeout and retry logic
- [x] 14 fetcher tests

### Task 4: Prompt Parser — DONE
- [x] `src/prompts.py` — Detects markdown blocks, `<prompt>` tags, indented blocks
- [x] `Prompt` dataclass with format and source line tracking
- [x] 19 prompt tests

### Task 5: Neo4j Bridge — DONE
- [x] `src/neo4j.py` — `query_cypher()`, `get_findings()`, `get_task_status()`, `create_research()`
- [x] Server API client with error handling
- [x] 11 Neo4j tests

### Task 6: REPL Module — DONE
- [x] `src/repl.py` — Command table, routing, free-text passthrough
- [x] Commands: fetch, prompt, ask, query, research, status, help, quit
- [x] Rich terminal output
- [x] 22 REPL tests

### Task 7: Integration — DONE
- [x] `src/main.py` launches REPL with all tool dependencies
- [x] `uv run python -m src.main` launches interactive shell

### Task 8: Docker Compose — DONE
- [x] `llamacpp` service in `docker-compose.yml`
- [x] llama.cpp env vars in `.env` / `.env.example`

## Next Steps

- Server Phase 2+ endpoints needed for Neo4j bridge commands to work live
- Download a .gguf model to `$HOME/llama/models/` for local testing
- Consider: conversation history persistence, response streaming, agent autonomy
