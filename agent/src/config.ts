export interface ServerConfig {
  url: string;
  apiKey?: string;
}

const DEFAULTS: ServerConfig = {
  url: "http://localhost:8080",
};

export const getConfig = (): ServerConfig => {
  return {
    ...DEFAULTS,
    url: process.env.LLAMA_SERVER_URL || DEFAULTS.url,
    apiKey: process.env.LLAMA_API_KEY || undefined,
  };
};
