package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"url-shortner/internal/config"
	"url-shortner/internal/db"
)

type Server struct {
	port int

	db db.Service
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: config.PORT,

		db: db.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
