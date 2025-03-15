package handler

import (
	"net/http"
	"url-shortner/cmd/web/home"
	"url-shortner/internal/service"

	"github.com/a-h/templ"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Read the long URL from the request body
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	// Create a short URL using the URL service
	shortURL, err := h.service.CreateShortURL(url)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	handler := templ.Handler(home.URLForm(url, shortURL))
	handler.ServeHTTP(w, r)
}
