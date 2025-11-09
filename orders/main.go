package main

import (
	"context"
	"github.com/ZaharBorisenko/z-orders/internal/gRPC"
	"github.com/ZaharBorisenko/z-orders/internal/services"
	"github.com/ZaharBorisenko/z-orders/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	grpcAddr = "localhost:2000"
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	defer l.Close()

	store := storage.NewStore()
	service := services.NewService(store)
	gRPC.NewGrpcHandler(grpcServer)

	err = service.CreateOrder(context.Background())
	if err != nil {
		log.Fatal("service orders error", err)
	}

	log.Println("gRPC server starting!", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
