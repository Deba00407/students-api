package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Deba00407/students-api/internal/config"
)

func main() {
	// Load config
	config := config.MustLoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello world !!"))
	})

	// setup server
	serverStatus := make(chan os.Signal, 1)

	server := http.Server{
		Addr:    config.ServerPath,
		Handler: router,
	}

	signal.Notify(serverStatus, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start the server")
		}
	}()

	slog.Info("Server running on:", slog.String("address:", config.ServerPath))

	<-serverStatus

	// Shutdown server gracefully

	slog.Info("Shutting down server...")

	// Incase a process is running, the server waits for the process to end for 5 secs and once it is complete, if any new request appears it is not accepted
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel() // release resources in case the process exits before the time duration set

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server:", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfull")
}
