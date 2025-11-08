package main

import (
	"github.com/ZaharBorisenko/z-gateway/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	httpAddr := os.Getenv("HTTP_ADDR")

	mux := http.NewServeMux()
	handler := handlers.NewHandler()
	handler.RegisterRoutes(mux)

	log.Printf("Starting HTTP server on address %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
