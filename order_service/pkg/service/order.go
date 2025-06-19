package service

import (
	"catalog_service/pkg/proto"
	"catalog_service/pkg/repository"
	"context"
)

type OrderService struct {
	proto.UnimplementedOrderServer
	repo *repository.Repository
	catalogService *client.CatalogClient
	cartService *client.CartClient
	
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrderByCart(
	ctx context.Context,
	req *proto.CreateOrderByCartRequest,
) (*proto.CreateOrderByCartResponse, error) {
	client.
}
