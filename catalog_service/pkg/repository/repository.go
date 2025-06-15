package repository

import (
	"catalog_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type Topping interface {
	GetToppingByID(id int) (entity.Topping, error)
	GetToppingsByProductID(productID int) ([]entity.Topping, error)
}

type Product interface {
	GetProductByID(id int) (entity.Product, error)
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
