# Agent — Status

## Overall Progress

**Phase 1 (Foundation):** Config, server health check, inference streaming complete
**Current Phase:** 1

## Completed

| Task | Status | Description |
|------|--------|-------------|
| 1 | ✅ | Python agent removed |
| 2 | ✅ | TS package.json created |
| 3 | ✅ | Vite config created |
| 4 | ✅ | TypeScript config created |
| 5 | ✅ | PLAN_AGENT.md rewritten with agreed architecture |
| 6 | ✅ | `src/config.ts` — `ServerConfig` interface + `getConfig()` with env var overrides |
| 7 | ✅ | `src/checkInferenceServer` — health check for llama-server (local + remote with API key) |
| 8 | ✅ | `src/inference.types.ts` — TypeScript interfaces for llama-server API (health, errors, chat) |
| 9 | ✅ | `src/inference/kyApiClient.ts` — Ky HTTP client with auth, timeout, retry |
| 10 | ✅ | `src/inference/inference.ts` — Streaming `chatCompletion()` with SSE + tool call support |
| 11 | ✅ | `docs/LLAMA_SERVER_API.md` — API reference doc for llama-server endpoints |

## Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `ky` | 2.0.2 | HTTP client (auth, retry, timeout) |
| `parse-sse` | 0.1.0 | SSE stream parsing |
| `vitest` | ^4.1.6 | Testing framework |

## Project Structure

```
agent/
├── package.json
├── tsconfig.json
├── vite.config.ts
├── yarn.lock
└── src/
    ├── config.ts                    # ServerConfig + getConfig()
    ├── main.ts                      # checkInferenceServer, prompt stubs
    ├── inference.types.ts           # API interfaces (Health, Error, Chat)
    ├── inference/
    │   ├── kyApiClient.ts           # Ky client factory
    │   ├── inference.ts             # chatCompletion() streaming generator
    │   └── __test__/
    │       ├── kyApiClient.test.ts  # 3 tests
    │       └── inference.test.ts    # 3 tests
    └── __test__/
        ├── config.test.ts           # 5 tests (config + health check)
        └── main.test.ts             # 1 test (promptCli stub)
```

## Test Coverage

| File | Tests | Status |
|------|-------|--------|
| `config.test.ts` | 5 | ✅ Passing |
| `kyApiClient.test.ts` | 3 | ✅ Passing |
| `inference.test.ts` | 3 | ✅ Passing |
| `main.test.ts` | 1 | ❌ Pre-existing stub (promptCli) |

**Total: 12 tests, 11 passing**

## Next Steps

- Implement `promptCli` to call `chatCompletion` and yield results
- Build plan generator (LLM → structured JSON research plan)
- Build basic UI with input prompt
- Implement plan executor with SSE consumption from Go server

## Architecture Agreement

All key decisions documented in PLAN_AGENT.md:
- Browser-based TS agent
- Declarative research plans
- Go server handles tools and graph writes
- SSE for real-time streaming
- Plugin-defined schemas
- Typed JSON entity contract
