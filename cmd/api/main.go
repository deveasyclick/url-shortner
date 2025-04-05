package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	web "url-shortner/cmd/web/server"
	"url-shortner/internal/config"
	"url-shortner/internal/server"
	"url-shortner/pkg/logger"
)

/**
// TODO: Replace log with slog
// TODO: Move web outside cmd folder
*/

func main() {
	// Initialize the logger
	logger.Init()
	server := server.NewServer()
	webServer := web.NewServer()

	var wg sync.WaitGroup

	// Start API server
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Server started", "port", config.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("HTTP server error: %s", "error", err)
		}
	}()

	// Start Web server
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Web server started", "port", config.WEB_PORT)
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Web server error: %s", "error", err)
		}
	}()

	// Handle graceful shutdown for both servers
	wg.Add(2)
	go gracefulShutdown(server, &wg)
	go gracefulShutdown(webServer, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	slog.Info("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	slog.Info("Shutting down server on port %s gracefully...", "port", apiServer.Addr)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "port", apiServer.Addr, "error", err)
	} else {
		slog.Info("Server shut down gracefully.", "port", apiServer.Addr)
	}
}
