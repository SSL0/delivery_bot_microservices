package model

type Cart struct {
	Id     uint64 `db:"id"`
	UserId uint64 `db:"user_id"`
	Items  []CartItem
}

type CartItem struct {
	Id       uint64 `db:"id"`
	CartId   uint64 `db:"cart_id"`
	ItemId   uint64 `db:"item_id"`
	Type     string `db:"type"` // "product" or "topping"
	Price    string `db:"price"`
	Quantity uint32 `db:"quantity"`
}
