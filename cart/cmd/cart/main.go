package main

import (
	"cart/pkg/client"
	"cart/pkg/config"
	"cart/pkg/proto"
	"cart/pkg/repository"
	"cart/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("./cmd/cart/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	postgres, err := repository.NewPostgresdb(config.DBUrl)

	if err != nil {
		log.Fatalf("failed to create db connection: %v", err)
	}
	defer postgres.Close()

	productClient := client.InitCatalogClient(config.CatalogServiceAddress)

	repo := repository.NewRepository(postgres)
	cartServer := service.NewCartServer(repo, productClient)

	grpcServer := grpc.NewServer()
	proto.RegisterCartServer(grpcServer, cartServer)
	lis, err := net.Listen("tcp", config.ListeningAddress)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer grpcServer.GracefulStop()

}
