package main

import (
	"cart_service/pkg/client"
	"cart_service/pkg/config"
	"cart_service/pkg/proto"
	"cart_service/pkg/repository"
	"cart_service/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("./cmd/cart_service/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	postgres, err := repository.NewPostgresdb(config.DBUrl)

	if err != nil {
		log.Fatalf("failed to create db connection: %v", err)
	}
	defer postgres.Close()

	productClient := client.InitCatalogClient(config.ProductServiceAddress)
	// result, err := productClient.GetProduct(1)

	// if err != nil {
	// 	log.Fatalf("failed to get products: %v", err)
	// }
	// log.Printf("sucessfult get product by id %v", result.Product)

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
