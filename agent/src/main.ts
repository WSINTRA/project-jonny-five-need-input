import { getConfig } from "./config";

export const checkInferenceServer = async (): Promise<boolean> => {
  try {
    const config = getConfig();
    const headers: Record<string, string> = {};
    if (config.apiKey) {
      headers["Authorization"] = `Bearer ${config.apiKey}`;
    }
    const res = await fetch(config.url, { headers });
    return res.ok;
  } catch {
    return false;
  }
};

export const promptParse = () => {}
export const promptCli = (prompt: string, parser: typeof promptParse) => {}
