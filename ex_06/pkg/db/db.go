package db

import "github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"

type DB struct {
	DBProducts DBProducts
	DBCart     DBCart
}

type DBProducts struct {
	LastIdx  int64
	Products map[int64]model.Product
}

type DBCart struct {
	Cart  *model.Cart
	Items map[int64]model.CartItem
}

func NewDB() *DB {
	return &DB{
		DBProducts: DBProducts{
			LastIdx:  0,
			Products: make(map[int64]model.Product),
		},
		DBCart: DBCart{
			Cart:  nil,
			Items: make(map[int64]model.CartItem),
		},
	}
}
