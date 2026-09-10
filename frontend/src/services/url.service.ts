import { apiRequest } from "./api";
import type { CreateURLRequest, CreateURLResponse } from "../types/url";

// Request to create short URL
export function createShortURl(request: CreateURLRequest, token: string) {
  return apiRequest<CreateURLResponse>("/v1/api/url/", {
    method: "POST",
    body: JSON.stringify(request),
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

// Request to fetch all the urls created by the respective user
export function fetchAll(token: string) {
  return apiRequest<CreateURLResponse[]>("/v1/api/url", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

// Request to fetch url by id
export function fetchByID(id: string, token: string) {
  return apiRequest<CreateURLResponse>(`/v1/api/url/${id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

// Request to Delete url by id
export function deleteById(id: string, token: string) {
  return apiRequest<void>(`/v1/api/url/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}
