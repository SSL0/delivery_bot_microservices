package main

import (
	"catalog/pkg/config"
	"catalog/pkg/proto"
	"catalog/pkg/repository"
	"catalog/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("./cmd/catalog/config.json")
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
	service := service.NewCatalogService(repo)

	proto.RegisterCatalogServer(grpcServer, service)

	lis, err := net.Listen("tcp", config.ListeningAddress)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer grpcServer.GracefulStop()
}
