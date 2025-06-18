package service

import (
	"catalog_service/pkg/proto"
	"catalog_service/pkg/repository"
	"context"
	"log"

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
	log.Printf("GetProduct request: %v", req.Id)

	product, err := s.repo.GetProductByID(req.Id)

	if err != nil {
		log.Printf("failed to get toppings: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get toppings: %v", err)
	}

	pbProduct := &proto.Product{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Type:        product.Type,
	}

	log.Printf("GetProduct response: %v", pbProduct)

	return &proto.GetProductResponse{Product: pbProduct}, nil
}

func (s *CatalogService) GetProductToppings(
	ctx context.Context,
	req *proto.GetProductRequest,
) (*proto.GetProductToppingsResponse, error) {
	log.Printf("GetProductToppings request: %v", req.Id)
	toppings, err := s.repo.GetToppingsByProductID(req.Id)
	if err != nil {
		log.Printf("failed to get toppings: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get toppings: %v", err)
	}
	var pbToppings []*proto.Topping
	for _, topping := range toppings {
		pbToppings = append(pbToppings, &proto.Topping{
			Id:        topping.Id,
			ProductId: topping.ProductId,
			Name:      topping.Name,
			Price:     topping.Price,
		})
	}
	log.Printf("GetProductToppings response: %v", pbToppings)

	return &proto.GetProductToppingsResponse{
		Toppings: pbToppings,
	}, nil
}

func (s *CatalogService) GetTopping(
	ctx context.Context,
	req *proto.GetToppingRequest,
) (*proto.GetToppingResponse, error) {
	log.Printf("GetTopping request: %v", req.Id)
	topping, err := s.repo.GetToppingByID(req.Id)
	if err != nil {
		log.Printf("failed to get topping: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get topping: %v", err)
	}
	pbTopping := &proto.Topping{
		Id:        topping.Id,
		ProductId: topping.ProductId,
		Name:      topping.Name,
		Price:     topping.Price,
	}
	log.Printf("GetProductToppings response: %v", pbTopping)

	return &proto.GetToppingResponse{
		Topping: pbTopping,
	}, nil
}

func (s *CatalogService) GetProductsByType(
	ctx context.Context,
	req *proto.GetProductsByTypeRequest,
) (*proto.GetProductsByTypeResponse, error) {
	log.Printf("GetProductByType request: %v", req.Type)

	products, err := s.repo.GetProductsByType(req.Type)

	if err != nil {
		log.Printf("failed to get products by type: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get toppings: %v", err)
	}

	var pbProducts []*proto.Product
	for _, topping := range products {
		pbProducts = append(pbProducts, &proto.Product{
			Id:          topping.Id,
			Name:        topping.Name,
			Price:       topping.Price,
			Description: topping.Description,
			Type:        topping.Type,
		})
	}

	log.Printf("GetProductsByType response: %v", pbProducts)

	return &proto.GetProductsByTypeResponse{Products: pbProducts}, nil
}
