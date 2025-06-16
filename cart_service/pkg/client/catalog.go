package client

import (
	"cart_service/pkg/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CatalogClient struct {
	Client proto.CatalogClient
}

func InitCatalogClient(url string) *CatalogClient {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	client := &CatalogClient{
		Client: proto.NewCatalogClient(conn),
	}

	return client
}

func (c *CatalogClient) GetProduct(productId uint64) (*proto.GetProductResponse, error) {
	req := &proto.GetProductRequest{
		Id: productId,
	}

	return c.Client.GetProduct(context.Background(), req)
}

func (c *CatalogClient) GetProductToppings(productId uint64) (*proto.GetProductToppingsResponse, error) {
	req := &proto.GetProductRequest{
		Id: productId,
	}

	return c.Client.GetProductToppings(context.Background(), req)
}

func (c *CatalogClient) GetTopping(toppingId uint64) (*proto.GetToppingResponse, error) {
	req := &proto.GetToppingRequest{
		Id: toppingId,
	}

	return c.Client.GetTopping(context.Background(), req)
}
