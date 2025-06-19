package service

import (
	"cart_service/pkg/client"
	"cart_service/pkg/model"
	"cart_service/pkg/proto"
	"cart_service/pkg/repository"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartServer struct {
	proto.UnimplementedCartServer
	repo           *repository.Repository
	catalogService *client.CatalogClient
}

func NewCartServer(repo *repository.Repository, svc *client.CatalogClient) *CartServer {
	return &CartServer{
		repo:           repo,
		catalogService: svc,
	}
}

func (s *CartServer) AddItem(ctx context.Context, req *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	log.Printf("AddItem requested: user_id %v, item_id %v, type %v", req.UserId, req.ItemId, req.ItemType)
	cartId, err := s.repo.GetOrCreateCartIdByUserId(req.UserId)

	if err != nil {
		log.Printf("failed to get cart by user id %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get cart by user id %v", err)
	}

	var price string
	if req.ItemType == "product" {
		product, err := s.catalogService.GetProduct(req.ItemId)
		if err != nil {
			log.Printf("failed to get product %v", err)
			return nil, err
		}
		price = product.Product.Price
	} else if req.ItemType == "topping" {
		topping, err := s.catalogService.GetTopping(req.ItemId)
		if err != nil {
			log.Printf("failed to get topping %v", err)
			return nil, err
		}
		price = topping.Topping.Price
	} else {
		return nil, status.Errorf(codes.Internal, "unknown item type")
	}

	item := model.CartItem{
		CartId:   cartId,
		ItemId:   req.ItemId,
		Type:     req.ItemType,
		Price:    price,
		Quantity: req.Quantity,
	}
	log.Printf("%v", item)
	createdId, err := s.repo.AddItemToCartById(cartId, item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add item: %v", err)
	}
	log.Printf("AddItem response: created_cart_item_id %v", req.UserId)
	return &proto.AddItemResponse{AddedCartItemId: createdId}, nil
}

func (s *CartServer) RemoveCartItem(ctx context.Context, req *proto.RemoveCartItemRequest) (*proto.RemoveCartItemResponse, error) {
	log.Printf("RemoveCartItem requested: cart_item_id %v", req.CartItemId)
	err := s.repo.RemoveCartItemById(req.CartItemId)
	if err != nil {
		log.Printf("failed to remove cart item by id: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to remove cart item by id: %v", err)
	}

	return &proto.RemoveCartItemResponse{}, nil
}

func (s *CartServer) GetCart(ctx context.Context, req *proto.GetCartRequest) (*proto.GetCartResponse, error) {
	log.Printf("GetCart requested: cart_item_id %v", req.CartId)

	cart, err := s.repo.GetCartById(req.CartId)
	if err != nil {
		log.Printf("failed to get cart: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}

	var ptItems []*proto.CartItem

	for _, item := range cart.Items {
		ptItem := &proto.CartItem{
			Id:       item.Id,
			CartId:   item.CartId,
			ItemId:   item.ItemId,
			Type:     item.Type,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
		ptItems = append(ptItems, ptItem)
	}
	log.Printf("GetCart response: cartId: %v userId: %v items: %v", cart.Id, cart.Id, cart.Items)

	return &proto.GetCartResponse{Id: cart.Id, UserId: cart.Id, Items: ptItems}, nil

}
