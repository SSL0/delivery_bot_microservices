package repo

import (
	"cart_service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db}
}

func (r *CartRepository) GetCartById(id uint64) (*model.Cart, error) {
	var result model.Cart
	err := r.db.QueryRowx("SELECT * FROM carts WHERE id = $1 LIMIT 1", id).StructScan(&result)
	if err != nil {
		return &model.Cart{}, err
	}

	rows, err := r.db.Queryx("SELECT * FROM cart_items WHERE cart_id = $1", id)

	var items []model.CartItem

	for rows.Next() {
		var item model.CartItem
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	result.Items = items

	return &result, err
}

func (r *CartRepository) AddItemToCartById(cart_id uint64, item model.CartItem) (uint64, error) {
	var createdId uint64
	err := r.db.QueryRowx(`
		INSERT INTO cart_items (
			cart_id,
			item_id,
			type,
			price,
			quantity
		)  VALUES (
			:cart_id,
			:item_id,
			:type,
			:price,
			:quantity
		) RETURNING id
	`, item).Scan(&createdId)
	return createdId, err
}

func (r *CartRepository) RemoveCartItemById(id uint64) error {
	_, err := r.db.Exec("DELETE FROM cart_items WHERE id = $1", id)
	return err
}

func (r *CartRepository) GetCartIdByUserId(userId uint64) (uint64, error) {
	var result uint64
	err := r.db.QueryRowx("SELECT id FROM carts WHERE user_id = $1 LIMIT 1", userId).Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
