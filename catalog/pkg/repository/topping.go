package repository

import (
	"catalog/pkg/model"
	"log"

	"github.com/jmoiron/sqlx"
)

type ToppingRepository struct {
	db *sqlx.DB
}

func NewToppingRepository(db *sqlx.DB) *ToppingRepository {
	return &ToppingRepository{db}
}

func (r *ToppingRepository) GetToppingByID(id uint64) (model.Topping, error) {
	var result model.Topping
	err := r.db.QueryRowx("SELECT * FROM toppings WHERE id = $1 LIMIT 1", id).StructScan(&result)
	return result, err
}

func (r *ToppingRepository) GetToppingsByProductID(productID uint64) ([]model.Topping, error) {
	var result []model.Topping
	rows, err := r.db.Queryx(`
			SELECT * FROM toppings 
			WHERE product_id = $1
		`, productID)
	for rows.Next() {
		var topping model.Topping
		if err := rows.StructScan(&topping); err != nil {
			return nil, err
		}

		result = append(result, topping)
	}
	log.Println(result)
	return result, err
}
