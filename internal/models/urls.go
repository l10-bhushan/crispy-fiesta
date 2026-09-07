package models

import "time"

// Request structure for creating short url
type CreateShortURLRequest struct {
	Url string `json:"url"`
}

// Response structure of our url shortener
type CreateShortURLResponse struct {
	Id        string     `json:"id"`
	ShortCode string     `json:"short_code"`
	URL       string     `json:"original_url"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"` // we use *time.Time here coz a pointer can be nil
}
