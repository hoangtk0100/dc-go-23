package db

import "github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"

type DB struct {
	DBProducts DBProducts
}

type DBProducts struct {
	LastIdx  int64
	Products map[int64]model.Product
}

func NewDB() *DB {
	return &DB{
		DBProducts: DBProducts{
			LastIdx:  0,
			Products: make(map[int64]model.Product),
		},
	}
}
