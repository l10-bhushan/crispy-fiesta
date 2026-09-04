package handlers

import (
	"net/http"

	"github.com/l10-bhushan/crispy-fiesta/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

func (h *URLHandler) CreateShortCode(w http.ResponseWriter, r *http.Request) {

}

func (h *URLHandler) FetchURLData(w http.ResponseWriter, r *http.Request) {

}
