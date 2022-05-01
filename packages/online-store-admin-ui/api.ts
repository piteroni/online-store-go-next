import { apiClient, OpenAPIConfig } from "./generated"

const config: Partial<OpenAPIConfig> = {
  BASE: process.env.API_ENDPOINT,
}

export const api = new apiClient(config).default
