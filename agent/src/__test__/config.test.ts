import {describe, expect, it, vi, beforeEach} from "vitest";
import {getConfig, ServerConfig} from "../config";
import {checkInferenceServer} from "../main";

describe("ServerConfig", () => {
  it("returns default values when no env vars are set", () => {
    const config = getConfig();
    expect(config).toBeDefined();
    expect(config.url).toBe("http://localhost:8080");
    expect(config.apiKey).toBeUndefined();
  });

  it("implements ServerConfig interface", () => {
    const config: ServerConfig = getConfig();
    expect(typeof config.url).toBe("string");
  });
});

describe("checkInferenceServer", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it("returns true when inference server responds 200", async () => {
    vi.spyOn(global, "fetch").mockResolvedValueOnce({ok: true} as Response);
    const result = await checkInferenceServer();
    expect(result).toBe(true);
  });

  it("returns false when inference server is unreachable", async () => {
    vi.spyOn(global, "fetch").mockRejectedValueOnce(new Error("network error"));
    const result = await checkInferenceServer();
    expect(result).toBe(false);
  });

  it("includes Authorization header when apiKey is configured", async () => {
    vi.stubEnv("LLAMA_API_KEY", "sk-test-123");
    vi.stubEnv("LLAMA_SERVER_URL", "http://remote.example.com");
    vi.spyOn(global, "fetch").mockResolvedValueOnce({ok: true} as Response);
    await checkInferenceServer();
    expect(fetch).toHaveBeenCalledWith(
      "http://remote.example.com",
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: "Bearer sk-test-123",
        }),
      })
    );
  });
});
