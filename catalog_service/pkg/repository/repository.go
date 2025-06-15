package repository

import (
	"catalog_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type Product interface {
	GetProductInfoByID(id int) (entity.Product, error)
}

type Repository struct {
	Product
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductRepository(db),
		db:      db,
	}
}
