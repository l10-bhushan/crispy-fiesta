package middleware

import (
	"log"
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {

				// Fetch request ID from context
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}

				// Log panic
				log.Printf(
					"panic recovered: %v, x-request-id: %s",
					err,
					requestID,
				)

				// Send 500 response
				utils.WriteJsonResponse(
					w,
					http.StatusInternalServerError,
					map[string]string{
						"error": "internal server error",
					},
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
