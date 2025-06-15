package main

import (
	"catalog_service/pkg/config"
	"catalog_service/pkg/repository"
	"fmt"
)

func main() {
	config, err := config.LoadConfig("./config.json")
	if err != nil {
		panic(err)
	}
	postgres, err := repository.NewPostgresdb(config.DBUrl)

	if err != nil {
		panic(err)
	}
	defer postgres.Close()
	err = repository.Migrate(config.MigrationsPath, config.DBUrl)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(postgres)
	product, err := repo.GetProductInfoByID(1)

	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
