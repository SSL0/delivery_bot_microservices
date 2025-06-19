package client

import (
	"catalog_service/pkg/proto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CartClient struct {
	client proto.CatalogClient
}

func InitCartClient(url string) *CatalogClient {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	client := &CatalogClient{
		client: proto.NewCatalogClient(conn),
	}

	return client
}
