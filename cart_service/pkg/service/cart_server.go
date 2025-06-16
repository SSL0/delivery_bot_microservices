package service

import (
	"cart_service/pkg/client"
	"cart_service/pkg/model"
	"cart_service/pkg/proto"
	"cart_service/pkg/repo"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartServer struct {
	proto.UnimplementedCartServer
	repo           *repo.Repository
	catalogService client.CatalogClient
}

func NewCartServer(repo *repo.Repository, svc client.CatalogClient) *CartServer {
	return &CartServer{
		repo:           repo,
		catalogService: svc,
	}
}

func (s *CartServer) AddItem(ctx context.Context, req *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	log.Printf("AddItem requested: user_id %v, item_id %v, type %v", req.UserId, req.ItemId, req.ItemType)
	cartId, err := s.repo.GetCartIdByUserId(req.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			// TODO: create new cart
		} else {
			return nil, status.Errorf(codes.Internal, "failed to get cart by user id %v", err)
		}
	}

	if req.ItemType != "product" && req.ItemType != "topping" {
		return nil, status.Errorf(codes.Internal, "unknown item type")
	}

	item := model.CartItem{
		CartId:   cartId,
		ItemId:   req.ItemId,
		Type:     req.ItemType,
		Quantity: req.Quantity,
	}
	createdId, err := s.repo.AddItemToCartById(cartId, item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add item: %v", err)
	}
	log.Printf("AddItem response: created_cart_item_id %v", req.UserId)
	return &proto.AddItemResponse{AddedCartItemId: createdId}, nil
}
