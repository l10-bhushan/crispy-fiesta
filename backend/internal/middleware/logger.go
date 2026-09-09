package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the strongly-typed context key defined in requestId.go
		requestID, _ := r.Context().Value(requestIDKey).(string)
		if requestID == "" {
			requestID = "unknown"
		}

		fmt.Printf("---------------\nRequest logger:\nMethod: %s \nRequest URL : %s \nRequest Id: %s \n---------------\n", r.Method, r.URL.Path, requestID)
		next.ServeHTTP(w, r)
	})
}
