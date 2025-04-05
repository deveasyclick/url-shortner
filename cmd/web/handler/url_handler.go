package web_handler

import (
	"net/http"
	"url-shortner/cmd/web/home"
	"url-shortner/internal/service"
	custom_errors "url-shortner/pkg/errors"

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

// Add a handler for redirecting short URLs
func (h *URLHandler) RedirectShortURL(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the request
	shortURL := r.URL.Path[1:]
	if shortURL == "" {
		http.Error(w, "Short URL is missing", http.StatusBadRequest)
		return
	}
	// Redirect to the long URL using the URL service
	longURL, err := h.service.GetOriginalURL(shortURL)
	if err != nil {
		if err == custom_errors.ErrURLNotFound {
			http.Error(w, "Short URL not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
