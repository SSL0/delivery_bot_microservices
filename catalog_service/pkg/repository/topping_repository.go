package repository

import (
	"catalog_service/pkg/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

type ToppingRepository struct {
	db *sqlx.DB
}

func NewToppingRepository(db *sqlx.DB) *ToppingRepository {
	return &ToppingRepository{db}
}

func (r *ToppingRepository) GetToppingByID(id uint64) (entity.Topping, error) {
	var result entity.Topping
	err := r.db.QueryRowx("SELECT * FROM toppings WHERE id = $1 LIMIT 1", id).StructScan(&result)
	return result, err
}

func (r *ToppingRepository) GetToppingsByProductID(productID uint64) ([]entity.Topping, error) {
	var result []entity.Topping
	rows, err := r.db.Queryx(`
			SELECT * FROM toppings 
			WHERE product_id = $1
		`, productID)
	for rows.Next() {
		var topping entity.Topping
		if err := rows.StructScan(&topping); err != nil {
			return nil, err
		}

		result = append(result, topping)
	}
	log.Println(result)
	return result, err
}
