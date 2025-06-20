package client

import "order_service/pkg/model"

type Cart interface {
	GetCart(cartId uint64) (*model.Cart, error)
	GetCartIdByUserId(userId uint64) (uint64, error)
	RemoveCart(cartId uint64) error
}

type Catalog interface {
	GetProduct(productId uint64) (*model.Product, error)
	GetTopping(toppingId uint64) (*model.Topping, error)
}

type Client struct {
	Cart
	Catalog
}

func NewClient(cartServerAddress string, catalogServerAddress string) *Client {
	return &Client{
		Cart:    InitCartClient(cartServerAddress),
		Catalog: InitCatalogClient(catalogServerAddress),
	}
}
