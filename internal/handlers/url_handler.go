package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/models"
	"github.com/l10-bhushan/crispy-fiesta/internal/service"
	"github.com/l10-bhushan/crispy-fiesta/internal/utils"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// Handler to create short code
func (h *URLHandler) CreateShortCode(w http.ResponseWriter, r *http.Request) {

	// creating instance of CreateShortURLRequest
	var req models.CreateShortURLRequest

	// Decoding the r.Body into the instance we created
	// This will take help of `json:"url"` we wrote in the struct, the Decode would help us to generate struct.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	// Calling the service to start the process of creation
	urlData, err := h.service.Create(r.Context(), req.Url)
	if err != nil {
		utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// writing the data back to the user
	utils.WriteJsonResponse(w, http.StatusCreated, map[string]any{
		"status": "successfully created",
		"data":   urlData,
	})
}

// Redirect handler to redirect user to original url provided the shortCode
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {

	// Getting the shortCode from path parameter
	shortCode := r.PathValue("shortCode")
	if shortCode == "" {
		utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": "shortCode cannot be empty",
		})
		return
	}

	// Calling the service to retrieve the complete information
	urlData, err := h.service.GetOriginalURL(r.Context(), shortCode)
	if err != nil {
		utils.WriteJsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "URL not found",
		})
		return
	}

	// Redirecting the user to correct site
	http.Redirect(w, r, urlData.URL, http.StatusFound)
}

// Fetch all handler
func (h *URLHandler) FetchAllData(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll(r.Context())
	if err != nil {
		utils.WriteJsonResponse(w, http.StatusNotFound, map[string]string{
			"error": "failed to fetch data",
		})
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   data,
	})
}

// Fetch by id handler
func (h *URLHandler) FetchById(w http.ResponseWriter, r *http.Request) {
	id := string(r.PathValue("id"))

	data, err := h.service.FetchByID(r.Context(), id)
	if err != nil {
		utils.WriteJsonResponse(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   data,
	})
}

// Delete by Id handler
func (h *URLHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := string(r.PathValue("id"))

	err := h.service.DeleteById(r.Context(), id)
	if err != nil {
		utils.WriteJsonResponse(w, http.StatusBadGateway, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, map[string]string{
		"status": "success",
	})
}
