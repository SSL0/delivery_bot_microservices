package entity

type Topping struct {
	Id        uint64
	ProductId uint64 `db:"product_id"`
	Name      string
	Price     string
}
