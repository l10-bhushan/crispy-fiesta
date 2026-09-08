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

// Handler to fetch all the records
func (h *UserHandler) FetchAll(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.FetchAll(r.Context())
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   users,
	})
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

// Handler for user login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	// Creating an instance of LoginRequest
	var loginRequest models.LoginRequest

	// Storing email and password in our instance
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		utils.WriteError(w, apperrors.ErrorInvalidInput)
		return
	}

	// Checking if any of the fields are empty
	if loginRequest.Email == "" || loginRequest.Password == "" {
		utils.WriteError(w, apperrors.ErrorInvalidInput)
		return
	}

	// Calling the login service
	token, err := h.service.Login(r.Context(), loginRequest)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"token":  token,
	})

}

// Handler to delete a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"status": "success",
	})
}
