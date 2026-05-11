import ky from "ky";
import { getConfig } from "../config";

export const createKyClient = () => {
  const config = getConfig();
  const headers: Record<string, string> = {};
  if (config.apiKey) {
    headers["Authorization"] = `Bearer ${config.apiKey}`;
  }

  return ky.create({
    prefix: config.url,
    headers,
    timeout: 60_000,
    retry: 2,
  });
};
