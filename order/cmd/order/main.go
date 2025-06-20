package main

import (
	"log"
	"net"
	"order/pkg/client"
	"order/pkg/config"
	"order/pkg/proto"
	"order/pkg/repository"
	"order/pkg/service"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("./cmd/order/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	postgres, err := repository.NewPostgresdb(config.DBUrl)

	if err != nil {
		log.Fatalf("failed to create db connection: %v", err)
	}
	defer postgres.Close()

	grpcServer := grpc.NewServer()

	repo := repository.NewRepository(postgres)
	client := client.NewClient(config.CartServiceAddress, config.CatalogServiceAddress)
	service := service.NewOrderService(repo, client)

	proto.RegisterOrderServer(grpcServer, service)

	lis, err := net.Listen("tcp", config.ListeningAddress)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer grpcServer.GracefulStop()
}
