package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
)

func WriteError(w http.ResponseWriter, err error) {

	status := http.StatusInternalServerError
	message := "internal server error"

	switch {
	case errors.Is(err, apperrors.ErrorConflict):
		status = http.StatusConflict
		message = "resource already exists"
	case errors.Is(err, apperrors.ErrorInvalidInput):
		status = http.StatusBadRequest
		message = "invalid input"
	case errors.Is(err, apperrors.ErrorNotFound):
		status = http.StatusNotFound
		message = "resource not found"
	case errors.Is(err, apperrors.ErrorInternal):
		status = http.StatusInternalServerError
		message = "internal server error"
	}

	WriteJsonResponse(w, status, map[string]string{
		"error": message,
	})
}

// Helper function to write json response
func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
