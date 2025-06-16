package entity

type Cart struct {
	Id     uint64
	UserId uint64 `db:"user_id"`
	Items  []CartItem
}

type CartItem struct {
	Id       uint64
	CartId   uint64 `db:"cart_id"`
	ItemId   uint64 `db:"item_id"`
	Type     string // "product" or "topping"
	Price    string
	Quantity uint
}
