package model

type Order struct {
	Id     uint64 `db:"id"`
	UserId uint64 `db:"user_id"`
	Items  []OrderItem
}

type OrderItem struct {
	Id       uint64 `db:"id"`
	OrderId  uint64 `db:"order_id"`
	ItemId   uint64 `db:"item_id"`
	Type     string `db:"type"` // "product" or "topping"
	Price    string `db:"price"`
	Quantity uint32 `db:"quantity"`
}
