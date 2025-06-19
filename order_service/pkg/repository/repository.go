package repository

import (
	"order_service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order *model.Order) (uint64, error)
	GetOrderByCartId(cartId uint64) (*model.Order, error)
}

type Repository struct {
	Order
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderRepository(db),
		db:    db,
	}
}
