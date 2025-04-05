package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"url-shortner/internal/service"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestData struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		slog.Error("Error decoding request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Validate the URL field
	if requestData.URL == "" {
		slog.Warn("Missing 'url' field in request body", "warn", requestData)
		http.Error(w, "Missing 'url' field in request body", http.StatusBadRequest)
		return
	}

	// Create a short URL using the URL service
	shortURL, err := h.service.CreateShortURL(requestData.URL)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	// Prepare the JSON response
	response := map[string]string{
		"short_url": shortURL,
	}

	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error in createShortUrl", "error", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
