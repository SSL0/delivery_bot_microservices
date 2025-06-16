package repository

import (
	"catalog_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type Topping interface {
	GetToppingByID(id uint64) (entity.Topping, error)
	GetToppingsByProductID(productID uint64) ([]entity.Topping, error)
}

type Product interface {
	GetProductByID(id uint64) (entity.Product, error)
}

type Repository struct {
	Product
	Topping
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductRepository(db),
		Topping: NewToppingRepository(db),
		db:      db,
	}
}
