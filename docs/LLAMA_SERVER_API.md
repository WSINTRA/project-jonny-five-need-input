# Llama-Server API Reference

Endpoints relevant to the ResearchBot agent. Derived from llama.cpp source code.

## Error Response (Universal)

All error responses follow the OpenAI error format:

```json
{
  "error": {
    "code": 401,
    "message": "Invalid API Key",
    "type": "authentication_error"
  }
}
```

Error types: `authentication_error`, `invalid_request_error`, `server_error`, `not_found_error`, `permission_denied_error`, `unavailable_error`, `not_supported_error`, `exceed_context_size_error`

---

## GET /health

Checks if the server is running and ready.

**Response:**
```json
{ "status": "ok" }
```

Or 503 when loading:
```json
{ "error": { "code": 503, "message": "Loading model", "type": "unavailable_error" } }
```

---

## GET /v1/models

Lists available models.

**Response:**
```json
{
  "object": "list",
  "data": [{
    "id": "model-alias",
    "object": "model",
    "created": 1735142223,
    "owned_by": "llamacpp",
    "meta": {
      "vocab_type": 2,
      "n_vocab": 128256,
      "n_ctx_train": 131072,
      "n_embd": 4096,
      "n_params": 8030261312,
      "size": 4912898304
    }
  }]
}
```

---

## POST /v1/chat/completions

Primary endpoint for sending prompts and getting responses. OpenAI-compatible.

**Request:**
```json
{
  "model": "gpt-3.5-turbo",
  "messages": [
    { "role": "system", "content": "You are a helpful assistant" },
    { "role": "user", "content": "What is 2+2?" }
  ],
  "temperature": 0.8,
  "top_p": 0.95,
  "max_tokens": 256,
  "stop": ["string"],
  "stream": false,
  "response_format": { "type": "json_object" },
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "search_venues",
        "description": "Search for venues in a location",
        "parameters": { "type": "object", "properties": { ... } }
      }
    }
  ],
  "tool_choice": "auto"
}
```

**Required fields:** `model`, `messages`

**Optional fields:**
| Field | Type | Default | Purpose |
|-------|------|---------|---------|
| `temperature` | float | 0.8 | Sampling temperature |
| `top_p` | float | 0.95 | Nucleus sampling |
| `top_k` | int | 40 | Top-k sampling |
| `min_p` | float | 0.05 | Min-p sampling |
| `max_tokens` | int | — | Max tokens to generate |
| `stop` | string[] | [] | Stop sequences |
| `stream` | bool | false | SSE streaming |
| `n` | int | 1 | Number of completions |
| `seed` | int | -1 | RNG seed |
| `presence_penalty` | float | 0.0 | Presence penalty |
| `frequency_penalty` | float | 0.0 | Frequency penalty |
| `repetition_penalty` | float | 1.1 | Repetition penalty |
| `response_format` | object | — | JSON output constraint |
| `tools` | array | — | Function calling definitions |
| `tool_choice` | string | "auto" | Tool selection mode |

**Response (non-streaming):**
```json
{
  "id": "chatcmpl-xxxxx",
  "object": "chat.completion",
  "created": 1735142223,
  "model": "model-alias",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "generated text...",
        "tool_calls": [
          {
            "id": "call_abc",
            "type": "function",
            "function": { "name": "search_venues", "arguments": "{\"location\":\"Middlesbrough\"}" }
          }
        ]
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 50,
    "completion_tokens": 200,
    "total_tokens": 250
  },
  "timings": {
    "cache_n": 100,
    "prompt_n": 50,
    "prompt_ms": 120.5,
    "predicted_n": 200,
    "predicted_ms": 3000.0,
    "predicted_per_second": 66.7
  }
}
```

**Message roles:** `system`, `user`, `assistant`, `tool`

**Finish reasons:** `stop`, `length`, `tool_calls`

---

## Streaming (SSE)

Set `stream: true` in the request. Response is a stream of SSE events:

```
data: {"id":"chatcmpl-xxx","object":"chat.completion.chunk","created":123,"model":"...","choices":[{"index":0,"delta":{"role":"assistant","content":""},"finish_reason":null}]}

data: {"id":"chatcmpl-xxx","object":"chat.completion.chunk","created":123,"model":"...","choices":[{"index":0,"delta":{"content":"Hello"},"finish_reason":null}]}

data: {"id":"chatcmpl-xxx","object":"chat.completion.chunk","created":123,"model":"...","choices":[{"index":0,"delta":{},"finish_reason":"stop"}]}

data: [DONE]
```

---

## Authentication

For remote servers, use Bearer token:

```
Authorization: Bearer sk-your-api-key
```

---

## TypeScript Interfaces

Corresponding interfaces are in `agent/src/llama.ts`:

| Interface | Purpose |
|-----------|---------|
| `HealthResponse` | GET /health response |
| `LlamaServerError` | Universal error format |
| `ChatMessage` | Message in conversation |
| `ChatCompletionRequest` | POST /v1/chat/completions body |
| `ChatCompletionResponse` | POST /v1/chat/completions response |
