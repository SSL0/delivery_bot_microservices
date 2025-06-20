package client

import (
	"context"
	"log"
	"order_service/pkg/model"
	"order_service/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CartClient struct {
	client proto.CartClient
}

func InitCartClient(url string) *CartClient {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("Could not connect:", err)
	}

	client := &CartClient{
		client: proto.NewCartClient(conn),
	}

	return client
}

func (c *CartClient) GetCartIdByUserId(userId uint64) (uint64, error) {
	req := &proto.GetCartIdByUserIdRequest{
		UserId: userId,
	}

	res, err := c.client.GetCartIdByUserId(context.Background(), req)
	if err != nil {
		return 0, err
	}

	return res.CartId, err
}

func (c *CartClient) GetCart(cartId uint64) (*model.Cart, error) {
	req := &proto.GetCartRequest{
		CartId: cartId,
	}

	res, err := c.client.GetCart(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var cartItems []model.CartItem

	for _, pbCartItem := range res.Items {
		cartItems = append(cartItems, model.CartItem{
			Id:       pbCartItem.Id,
			CartId:   pbCartItem.CartId,
			ItemId:   pbCartItem.ItemId,
			Type:     pbCartItem.Type,
			Quantity: pbCartItem.Quantity,
		})
	}
	return &model.Cart{
		Id:     res.Id,
		UserId: res.UserId,
		Items:  cartItems,
	}, nil
}
