package main

import (
	"context"
	"github.com/ZaharBorisenko/z-orders/internal/services"
	"github.com/ZaharBorisenko/z-orders/internal/storage"
	"log"
)

func main() {
	store := storage.NewStore()
	service := services.NewService(store)

	err := service.CreateOrder(context.Background())
	if err != nil {
		log.Fatal("service orders error", err)
	}

}
