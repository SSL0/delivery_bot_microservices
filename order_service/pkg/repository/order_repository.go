package repository

import (
	"catalog_service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateOrderWithCart(cart model.Cart) (uint64, error) {
	var createdOrderId uint64

	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	err = tx.QueryRow(`
		INSERT INTO order (
			user_id
		) VALUES (
			$1 
		) RETURNING id
	`, cart.UserId).Scan(&createdOrderId)

	if err != nil {
		return 0, err
	}

	for _, item := range cart.Items {
		var createdOrderItemId uint64
		err := tx.QueryRow(`
		INSERT INTO order_items (
			order_id,
			item_id,
			type,
			price,
			quantity
		)  VALUES (
			$1, $2, $3, $4, $5
		) REUTRNING id
		`, createdOrderId, item.ItemId, item.Type, item.Price, item.Quantity).Scan(createdOrderItemId)
		if err != nil {
			return 0, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return createdOrderId, err
}

func (r *OrderRepository) GetOrderByCartId(cartId uint64) (*model.Order, error) {
	var result model.Order
	err := r.db.QueryRowx("SELECT * FROM order WHERE cart_id = $1 LIMIT 1", cartId).StructScan(&result)
	return &result, err
}
