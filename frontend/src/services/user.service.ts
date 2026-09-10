import type {
  LoginAndRegisterResponse,
  LoginRequest,
  RegisterRequest,
} from "../types/user";
import { apiRequest } from "./api";

export function Login(request: LoginRequest) {
  return apiRequest<LoginAndRegisterResponse>("/v1/api/user/login", {
    method: "POST",
    body: JSON.stringify(request),
  });
}

export async function Register(request: RegisterRequest) {
  return apiRequest<LoginAndRegisterResponse>("/v1/api/user/register", {
    method: "POST",
    body: JSON.stringify(request),
  });
}
