// Importing our variable from .env using VITE
// caution: never save API keys, database wiht the VITE prefix as these keys are expose on the browser
const API_URL = import.meta.env.VITE_API_URL;

// Creating a helper to handle our API requests
export async function apiRequest<T>(
  endpoint: string, // receives our endpoint
  options?: RequestInit, // handles options if any, like auth token
): Promise<T> {
  // Returns a promise

  // Using in-built fetch to trigger request
  const response = await fetch(`${API_URL}${endpoint}`, {
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
    ...options,
  });

  // If no response throw a error
  if (!response.ok) {
    throw new Error(`API request failed: ${response.status}`);
  }

  // return the json response
  return response.json();
}
