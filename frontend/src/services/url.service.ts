import { apiRequest } from "./api";
import type { CreateURLRequest, CreateURLResponse } from "../types/url";

export function createShortURl(request: CreateURLRequest, token: string) {
  return apiRequest<CreateURLResponse>("/v1/api/url/", {
    method: "POST",
    body: JSON.stringify(request),
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

export function fetchAll(token: string) {
  return apiRequest<CreateURLResponse[]>("/v1/api/url", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

export function fetchByID(id: string, token: string) {
  return apiRequest<CreateURLResponse>(`/v1/api/url/${id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

export function deleteById(id: string, token: string) {
  return apiRequest<void>(`/v1/api/url/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}
