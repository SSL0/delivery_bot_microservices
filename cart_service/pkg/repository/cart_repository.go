package repository

import (
	"cart_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db}
}

func (r *CartRepository) GetCartById(id uint64) (*entity.Cart, error) {
	var result entity.Cart
	err := r.db.QueryRowx("SELECT * FROM carts WHERE id = $1 LIMIT 1", id).StructScan(&result)
	if err != nil {
		return &entity.Cart{}, err
	}

	rows, err := r.db.Queryx("SELECT * FROM cart_items WHERE cart_id = $1", id)

	var items []entity.CartItem

	for rows.Next() {
		var item entity.CartItem
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	result.Items = items

	return &result, err
}
