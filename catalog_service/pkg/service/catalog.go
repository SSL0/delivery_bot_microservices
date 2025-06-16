package service

import (
	"catalog_service/pkg/proto"
	"catalog_service/pkg/repository"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CatalogService struct {
	proto.UnimplementedCatalogServer
	repo *repository.Repository
}

func NewCatalogService(repo *repository.Repository) *CatalogService {
	return &CatalogService{
		repo: repo,
	}
}

func (s *CatalogService) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	product, err := s.repo.GetProductByID(req.GetId())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get toppings: %v", err)
	}

	pbProduct := &proto.Product{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Type:        product.Price,
	}

	return &proto.GetProductResponse{Product: pbProduct}, nil
}

func (s *CatalogService) GetProductToppings(
	ctx context.Context,
	req *proto.GetProductRequest,
) (*proto.GetProductToppingsResponse, error) {
	toppings, err := s.repo.GetToppingsByProductID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get toppings: %v", err)
	}

	var pbToppings []*proto.Topping
	for _, topping := range toppings {
		pbToppings = append(pbToppings, &proto.Topping{
			Id:    topping.Id,
			Name:  topping.Name,
			Price: topping.Price,
		})
	}

	return &proto.GetProductToppingsResponse{
		Toppings: pbToppings,
	}, nil
}
