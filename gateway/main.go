package main

import (
	api "github.com/ZaharBorisenko/z-gateway/gRPC/gen/go"
	"github.com/ZaharBorisenko/z-gateway/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	httpOrdersService := os.Getenv("HTTP_ORDERS")

	conn, err := grpc.Dial(httpOrdersService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failder to dial orders-server %v", err)
	}
	defer conn.Close()

	client := api.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := handlers.NewHandler(client)
	handler.RegisterRoutes(mux)

	log.Printf("Starting HTTP server on address %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
