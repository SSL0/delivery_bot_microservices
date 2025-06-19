package client

import (
	"context"
	"fmt"
	"order_service/pkg/model"
	"order_service/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CatalogClient struct {
	client proto.CatalogClient
}

func InitCatalogClient(url string) *CatalogClient {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	client := &CatalogClient{
		client: proto.NewCatalogClient(conn),
	}

	return client
}

func (c *CatalogClient) GetProduct(productId uint64) (*model.Product, error) {
	req := &proto.GetProductRequest{
		Id: productId,
	}

	res, err := c.client.GetProduct(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return &model.Product{
		Id:          res.Product.Id,
		Name:        res.Product.Name,
		Price:       res.Product.Price,
		Description: res.Product.Description,
		Type:        res.Product.Type,
	}, nil
}

func (c *CatalogClient) GetTopping(toppingId uint64) (*model.Topping, error) {
	req := &proto.GetToppingRequest{
		Id: toppingId,
	}

	res, err := c.client.GetTopping(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return &model.Topping{
		Id:        res.Topping.Id,
		ProductId: res.Topping.ProductId,
		Name:      res.Topping.Name,
		Price:     res.Topping.Price,
	}, nil
}
