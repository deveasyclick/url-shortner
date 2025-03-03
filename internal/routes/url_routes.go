package routes

import (
	"net/http"
	url_ui "url-shortner/cmd/web/url"
	"url-shortner/internal/handler"
	"url-shortner/internal/repository"
	"url-shortner/internal/service"

	"github.com/a-h/templ"
	"gorm.io/gorm"
)

func RegisterURLRoutes(mux *http.ServeMux, db *gorm.DB) {
	urlRepo := repository.NewURLRepository(db)
	urlSvc := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlSvc)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handler := templ.Handler(url_ui.URLForm())
			handler.ServeHTTP(w, r)
		case http.MethodPost:
			urlHandler.CreateShortURL(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
