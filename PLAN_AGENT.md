# Agent — Knowledge Acquirer

## Role

TypeScript browser application that orchestrates research tasks. Calls llama-server for LLM inference, delegates tool execution and graph writes to the Go server backend.

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  Browser (TS Agent)                                         │
│  ┌──────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │ Input UI │→│ Plan Generator│→│ Plan Executor         │  │
│  │          │  │ (llama-server)│  │ (SSE from Go server) │  │
│  └──────────┘  └──────────────┘  └──────────────────────┘  │
└──────────┬──────────────────────────────────────────────────┘
           │
    ┌──────┴──────┐
    │             │
    ▼             ▼
llama-server   Go Server (tools + graph writer)
(inference)    ┌─────────────────────────────────┐
               │ Tools: Playwright, Google Places │
               │ Graph Writer: Typed JSON → Neo4j │
               │ SSE: streams execution progress  │
               └─────────────────────────────────┘

## Research Pipeline

1. User submits query → browser sends to llama-server
2. LLM generates structured research plan (JSON)
3. Browser sends plan to Go server for execution
4. Go server executes steps, calls tools, writes to Neo4j
5. Go server streams progress via SSE back to browser
6. Browser displays results as they arrive

## Plan Format

```json
{
  "query": "venues for musicians in Middlesbrough",
  "plugin": "music_venues",
  "goal": "Find live music venues with contact details and recent gig data",
  "steps": [
    {
      "id": 1,
      "action": "tool_call",
      "tool": "google_places",
      "params": { "query": "live music venues", "location": "Middlesbrough" },
      "expected_entities": ["Venue"],
      "depends_on": []
    },
    {
      "id": 2,
      "action": "tool_call",
      "tool": "playwright_scrape",
      "params": { "urls": "{{step_1.results[].website}}" },
      "expected_entities": ["Venue", "Contact", "Gig"],
      "depends_on": [1]
    },
    {
      "id": 3,
      "action": "llm_extract",
      "params": { "schema": "Venue, Contact, Gig", "source": "{{step_2.results}}" },
      "depends_on": [2]
    }
  ],
  "stop_criteria": {
    "target_entities": { "Venue": 15 },
    "max_steps": 10,
    "diminishing_returns_threshold": 3
  }
}
```

## Key Design Decisions

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

## Tool Registry

Tools are registered with descriptions. The planning LLM sees available tools and chooses during plan generation. The Go server executor validates and can retry with alternatives.

Initial toolset:
- `google_places` — Find businesses by type and location
- `playwright_scrape` — Extract content from URLs
- `search_engine` — Find relevant web pages

## Plugin System

Each research domain has a plugin that defines:
- Entity types and their fields (extraction schema)
- Available tools and their capabilities
- Neo4j node labels and relationship rules
- Search query templates

Example: `music_venues` plugin defines `Venue`, `Gig`, `Contact` types with their fields and graph mappings.

## Phases

| # | Phase | Outcome |
|---|-------|---------|
| 1 | Foundation | TS project, Vite, basic UI, llama-server connection |
| 2 | Plan Generator | LLM generates structured research plans |
| 3 | Plan Executor | Go server executes plans, SSE streaming |
| 4 | Tool Integration | Playwright, Google Places, search tools |
| 5 | Graph Writer | Typed JSON → Neo4j, entity dedup |
| 6 | Plugin System | Plugin loading, schema definitions, tool registry |
| 7 | Stopping Logic | Diminishing returns, completeness reports |
| 8 | Polish | Error handling, retry logic, UX refinements |

## Dependencies

- Go server must expose tool endpoints and SSE streaming
- llama-server must be accessible from the browser
- Neo4j schema must support plugin-defined extensions
