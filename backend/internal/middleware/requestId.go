package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fetching the requestId from the Header
		requestId := r.Header.Get("X-Request-ID")

		// Checking if requestId is available, if not we will generate and add it to the context
		if requestId == "" {

			// Reading a random string of bytes
			bytes := make([]byte, 10)

			// rand.Read will fillout bytes array with random byte
			if _, err := rand.Read(bytes); err != nil {
				// Checking if rand.Read returns an error
				fmt.Printf("error while generating requestId: %v", err)
			}

			// using hex.EncodeToString to turn byte array into string, and setting a new values as
			// requestID into request context
			r = r.WithContext(context.WithValue(r.Context(), requestIDKey, hex.EncodeToString(bytes)))

			// Also writing the requestID to response header
			w.Header().Set("requestId", hex.EncodeToString(bytes))
		} else {

			// Doing the same thing here but instead using the requestID we got from
			// request header sent by user
			r = r.WithContext(context.WithValue(r.Context(), requestIDKey, requestId))
			w.Header().Set("requestId", requestId)
		}
		next.ServeHTTP(w, r)
	})
}
