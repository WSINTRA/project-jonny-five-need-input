import { isHTTPError } from "ky";
import { parseServerSentEvents } from "parse-sse";
import { createKyClient } from "./kyApiClient";
import type { ChatCompletionRequest } from "../inference.types";

export interface ToolCallChunk {
  index: number;
  id?: string;
  type?: string;
  function?: {
    name?: string;
    arguments?: string;
  };
}

export interface ChatChunk {
  content?: string;
  tool_calls?: ToolCallChunk[];
  finish_reason?: string | null;
}

export async function* chatCompletion(
  request: ChatCompletionRequest
): AsyncGenerator<ChatChunk> {
  const client = createKyClient();

  try {
    const response = await client.post("v1/chat/completions", {
      json: { ...request, stream: true },
    });

    const eventStream = parseServerSentEvents(response);
    const reader = eventStream.getReader();

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      if (value.data === "[DONE]") break;

      try {
        const parsed = JSON.parse(value.data);
        const choice = parsed.choices?.[0];
        if (!choice?.delta) continue;

        const chunk: ChatChunk = {
          content: choice.delta.content || undefined,
          tool_calls: choice.delta.tool_calls,
          finish_reason: choice.finish_reason || null,
        };

        yield chunk;
      } catch {
        // Skip malformed SSE data lines
      }
    }
  } catch (error) {
    if (isHTTPError(error)) {
      const body = await error.response.json() as { error?: { message: string } };
      throw new Error(body.error?.message || "Inference request failed");
    }
    throw error;
  }
}
