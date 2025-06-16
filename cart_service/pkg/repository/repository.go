package repository

import (
	"cart_service/pkg/entity"

	"github.com/jmoiron/sqlx"
)

type Cart interface {
	GetCartById(id uint64) (*entity.Cart, error)
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
