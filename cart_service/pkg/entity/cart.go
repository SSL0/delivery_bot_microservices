package entity

type Cart struct {
	Id     uint64
	UserId uint64
	Items  []CartItem
}

type CartItem struct {
	Id       uint64
	ItemId   uint64
	Type     string // "product" or "topping"
	Price    string
	Quantity uint
}
