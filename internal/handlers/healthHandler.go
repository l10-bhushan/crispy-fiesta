package handlers

import (
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

// Health handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
