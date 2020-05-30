import config from "../../config.json";

interface Config {
  BACKEND_HOST: string;
}

function loadConfig(): Config {
  if (config && config.BACKEND_HOST) return config;
  throw new Error(`Error reading config.json: ${config}`);
}

export default loadConfig();
