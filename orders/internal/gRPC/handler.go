package gRPC

import (
	"context"
	api "github.com/ZaharBorisenko/z-gateway/gRPC/gen/go"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGrpcHandler(server *grpc.Server) {
	hanlder := &grpcHandler{}
	api.RegisterOrderServiceServer(server, hanlder)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, request *api.CreateOrderRequest) (*api.Order, error) {
	log.Println("New order!", request)
	order := &api.Order{
		ID:         "42",
		CustomerID: "",
		Status:     "",
		Items:      nil,
	}

	return order, nil
}
