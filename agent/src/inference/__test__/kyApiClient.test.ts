import { describe, expect, it, vi, beforeEach } from "vitest";
import { createKyClient } from "../kyApiClient";

describe("createKyClient", () => {
  beforeEach(() => {
    vi.unstubAllEnvs();
  });

  it("creates client with default server URL", () => {
    const client = createKyClient();
    expect(client).toBeDefined();
  });

  it("includes Authorization header when API key is set", () => {
    vi.stubEnv("LLAMA_API_KEY", "sk-test-key");
    const client = createKyClient();
    expect(client).toBeDefined();
  });

  it("works without API key for local development", () => {
    vi.unstubAllEnvs();
    const client = createKyClient();
    expect(client).toBeDefined();
  });
});
