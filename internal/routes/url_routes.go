package routes

import (
	"net/http"
	"url-shortner/internal/handler"
	"url-shortner/internal/repository"
	"url-shortner/internal/service"

	"gorm.io/gorm"
)

func RegisterURLRoutes(mux *http.ServeMux, db *gorm.DB) {
	urlRepo := repository.NewURLRepository(db)
	urlSvc := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlSvc)

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			switch r.Method {
			case http.MethodPost:
				urlHandler.CreateShortURL(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
		http.NotFound(w, r)
	})

}
