package repository

import (
	"catalog_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProductByID(id int) (entity.Product, error) {
	var result entity.Product
	err := r.db.QueryRowx("SELECT * FROM products WHERE id = $1 LIMIT 1", id).StructScan(&result)
	return result, err
}
