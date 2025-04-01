package web_server

import (
	"net/http"
	"url-shortner/cmd/web"
	handler "url-shortner/cmd/web/handler"
	"url-shortner/cmd/web/home"
	"url-shortner/internal/middleware"
	"url-shortner/internal/repository"
	"url-shortner/internal/service"

	"github.com/a-h/templ"
	"gorm.io/gorm"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)

	// Register URL routes
	RegisterURLRoutes(mux, s.db.DB)

	// Wrap the mux with CORS middleware
	return middleware.Cors(mux)
}

func RegisterURLRoutes(mux *http.ServeMux, db *gorm.DB) {
	urlRepo := repository.NewURLRepository(db)
	urlSvc := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlSvc)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			switch r.Method {
			case http.MethodGet:
				handler := templ.Handler(home.URLForm("", ""))
				handler.ServeHTTP(w, r)
			case http.MethodPost:
				urlHandler.CreateShortURL(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		// Handle dynamic path "/:urlParam"
		if len(r.URL.Path) > 1 {
			switch r.Method {
			case http.MethodGet:
				urlHandler.RedirectShortURL(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		// Reject other requests
		http.NotFound(w, r)
	})

}
