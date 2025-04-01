package server

import (
	"net/http"

	"url-shortner/internal/middleware"
	"url-shortner/internal/routes"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register URL routes
	routes.RegisterURLRoutes(mux, s.db.DB)

	// Wrap the mux with CORS middleware
	return middleware.Cors(mux)
}
