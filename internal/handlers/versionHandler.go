package handlers

import (
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

// Version handler
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"version": "1.0.0",
	})
}
