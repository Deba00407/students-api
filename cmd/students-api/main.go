package main

import (
	"fmt"
	"log"
	"net/http"

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
	server := http.Server{
		Addr:    config.ServerPath,
		Handler: router,
	}
	fmt.Println("Server running...")
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start the server")
	}
}
