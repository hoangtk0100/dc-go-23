package repository

import "github.com/hoangtk0100/dc-go-23/ex_06/pkg/db"

type repository struct {
	db          *db.DB
	productRepo ProductRepository
}

func NewRepository(db *db.DB) *repository {
	return &repository{
		db:          db,
		productRepo: NewProductRepository(db),
	}
}

func (repo *repository) Product() ProductRepository {
	return repo.productRepo
}
