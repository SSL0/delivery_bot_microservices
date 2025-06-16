package main

import (
	"cart_service/pkg/config"
	"cart_service/pkg/repository"
	"log"
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
	err = repository.Migrate(config.MigrationsPath, config.DBUrl)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	repo := repository.NewRepository(postgres)
	result, err := repo.GetCartById(1)
	log.Printf("%v", result)
	// service := service.NewCatalogService(repo)
}
