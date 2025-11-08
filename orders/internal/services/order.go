package services

import (
	"context"
	"github.com/ZaharBorisenko/z-orders/internal/models"
)

type Service struct {
	store models.OrdersStore
}

func NewService(store models.OrdersStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateOrder(ctx context.Context) error {
	return nil
}
