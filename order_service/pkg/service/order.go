package service

import (
	"context"
	"log"
	"order_service/pkg/client"
	"order_service/pkg/model"
	"order_service/pkg/proto"
	"order_service/pkg/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	proto.UnimplementedOrderServer
	repo   *repository.Repository
	client *client.Client
}

func NewOrderService(repo *repository.Repository, client *client.Client) *OrderService {
	return &OrderService{
		repo:   repo,
		client: client,
	}
}

func (s *OrderService) CreateOrderByCart(
	ctx context.Context,
	req *proto.CreateOrderByCartRequest,
) (*proto.CreateOrderByCartResponse, error) {
	log.Printf("CreateOrderByCart requested: cartId %v", req.CartId)

	pbCart, err := s.client.GetCart(req.CartId)

	if err != nil {
		log.Printf("failed to get cart: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}

	var newOrderItems []model.OrderItem

	for _, pbCartItem := range pbCart.Items {
		var orderItemPrice string
		if pbCartItem.Type == "product" {
			catalogProduct, err := s.client.GetProduct(pbCartItem.ItemId)
			if err != nil {
				log.Printf("failed to get product: %v", err)
				return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
			}
			orderItemPrice = catalogProduct.Price
		} else if pbCartItem.Type == "topping" {
			catalogTopping, err := s.client.GetTopping(pbCartItem.ItemId)
			if err != nil {
				log.Printf("failed to get topping: %v", err)
				return nil, status.Errorf(codes.Internal, "failed to get topping: %v", err)
			}
			orderItemPrice = catalogTopping.Price
		} else {
			log.Printf("failed get item: unknown type")
			return nil, status.Error(codes.Internal, "failed get item: unknown type")
		}

		newOrderItems = append(newOrderItems, model.OrderItem{
			Id:       pbCartItem.Id,
			ItemId:   pbCartItem.ItemId,
			Type:     pbCartItem.Type,
			Price:    orderItemPrice,
			Quantity: pbCartItem.Quantity,
		})
	}

	newOrder := &model.Order{
		UserId: req.CartId,
		Items:  newOrderItems,
	}
	createdOrderId, err := s.repo.CreateOrder(newOrder)
	if err != nil {
		log.Printf("failed to create order: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	log.Printf("CreateOrderByCart complete: order %v", createdOrderId)
	return &proto.CreateOrderByCartResponse{}, nil
}
