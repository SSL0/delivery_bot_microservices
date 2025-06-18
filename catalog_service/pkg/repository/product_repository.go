package repository

import (
	"catalog_service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProductByID(id uint64) (model.Product, error) {
	var result model.Product
	err := r.db.QueryRowx("SELECT * FROM products WHERE id = $1 LIMIT 1", id).StructScan(&result)
	return result, err
}

func (r *ProductRepository) GetProductsByType(productType string) ([]model.Product, error) {
	var result []model.Product
	rows, err := r.db.Queryx("SELECT * FROM products WHERE type = $1", productType)

	for rows.Next() {
		var item model.Product
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	return result, err
}
