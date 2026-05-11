export interface HealthResponse {
  status: string;
}

export interface InferenceServerError {
  error: {
    code: number;
    message: string;
    type: string;
  };
}

export interface ChatMessage {
  role: "system" | "user" | "assistant" | "tool";
  content: string;
  tool_call_id?: string;
}

export interface ChatCompletionRequest {
  model: string;
  messages: ChatMessage[];
  temperature?: number;
  top_p?: number;
  max_tokens?: number;
  stop?: string[];
  stream?: boolean;
  response_format?: {type: "json_object"};
}

export interface ChatCompletionResponse {
  id: string;
  object: "chat.completion";
  created: number;
  model: string;
  choices: {
    index: number;
    message: {role: "assistant"; content: string};
    finish_reason: "stop" | "length" | "tool_calls";
  }[];
  usage: {
    prompt_tokens: number;
    completion_tokens: number;
    total_tokens: number;
  };
}
