import { describe, expect, it, vi, beforeEach } from "vitest";
import { chatCompletion, type ChatChunk } from "../inference";

describe("chatCompletion", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it("streams content chunks from SSE response", async () => {
    const mockStream = new ReadableStream({
      start(controller) {
        const encoder = new TextEncoder();
        controller.enqueue(encoder.encode('data: {"id":"1","object":"chat.completion.chunk","choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"}}]}\n\n'));
        controller.enqueue(encoder.encode('data: {"id":"1","object":"chat.completion.chunk","choices":[{"index":0,"delta":{"content":" world"}}]}\n\n'));
        controller.enqueue(encoder.encode('data: [DONE]\n\n'));
        controller.close();
      },
    });

    vi.spyOn(global, "fetch").mockResolvedValueOnce({
      ok: true,
      body: mockStream,
    } as Response);

    const chunks: ChatChunk[] = [];
    for await (const chunk of chatCompletion({
      model: "test-model",
      messages: [{ role: "user", content: "say hello" }],
    })) {
      chunks.push(chunk);
    }

    expect(chunks).toHaveLength(2);
    expect(chunks[0].content).toBe("Hello");
    expect(chunks[1].content).toBe(" world");
  });

  it("throws when server returns error", async () => {
    vi.spyOn(global, "fetch").mockResolvedValueOnce({
      ok: false,
      status: 503,
      json: async () => ({
        error: { code: 503, message: "Loading model", type: "unavailable_error" },
      }),
    } as Response);

    const gen = chatCompletion({
      model: "test-model",
      messages: [{ role: "user", content: "test" }],
    });

    await expect(gen.next()).rejects.toThrow("Loading model");
  });

  it("includes tool_calls in chunks when present", async () => {
    const mockStream = new ReadableStream({
      start(controller) {
        const encoder = new TextEncoder();
        controller.enqueue(encoder.encode('data: {"id":"1","object":"chat.completion.chunk","choices":[{"index":0,"delta":{"tool_calls":[{"index":0,"id":"call_abc","type":"function","function":{"name":"search","arguments":"{}"}}]}}]}\n\n'));
        controller.enqueue(encoder.encode('data: [DONE]\n\n'));
        controller.close();
      },
    });

    vi.spyOn(global, "fetch").mockResolvedValueOnce({
      ok: true,
      body: mockStream,
    } as Response);

    const chunks: ChatChunk[] = [];
    for await (const chunk of chatCompletion({
      model: "test-model",
      messages: [{ role: "user", content: "search venues" }],
    })) {
      chunks.push(chunk);
    }

    expect(chunks).toHaveLength(1);
    expect(chunks[0].tool_calls).toBeDefined();
    expect(chunks[0].tool_calls?.[0].function?.name).toBe("search");
  });
});
