export interface CreateURLRequest {
  url: string;
}

export interface CreateURLResponse {
  id: string;
  shortcode: string;
  url: string;
  created_at: string;
}
