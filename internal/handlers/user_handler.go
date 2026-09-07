package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/apperrors"
	"github.com/l10-bhushan/crispy-fiesta/internal/models"
	"github.com/l10-bhushan/crispy-fiesta/internal/service"
	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Handler to find user by email
func (h *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {

	// Creating an instance of request
	var request models.FindByEmailRequest

	// decoding the r.Body into request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.WriteError(w, apperrors.ErrorInvalidInput)
		return
	}

	// Triggering the service for finding user by email
	user, err := h.service.FindByEmail(r.Context(), request.Email)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	// Returning the user
	utils.WriteJsonResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRequest models.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		utils.WriteError(w, apperrors.ErrorInvalidInput)
		return
	}

	userResponse, err := h.service.RegisterUser(r.Context(), userRequest)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusCreated, map[string]any{
		"status": "success",
		"data":   userResponse,
	})
}
