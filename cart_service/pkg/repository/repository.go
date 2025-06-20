package repository

import (
	"cart_service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Cart interface {
	GetCartById(id uint64) (*model.Cart, error)
	GetCartByUserId(userId uint64) (*model.Cart, error)
	GetOrCreateCartIdByUserId(userId uint64) (uint64, error)
	AddItemToCartById(cart_id uint64, item model.CartItem) (uint64, error)
	RemoveCartItemById(id uint64) error
}

type Repository struct {
	Cart
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Cart: NewCartRepository(db),
		db:   db,
	}
}
