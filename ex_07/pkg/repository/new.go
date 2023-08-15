package repository

import "github.com/hoangtk0100/dc-go-23/ex_07/pkg/db"

type repository struct {
	db          *db.DB
	productRepo ProductRepository
	cartRepo    CartRepository
	paymentRepo PaymentRepository
	userRepo    UserRepository
}

func NewRepository(db *db.DB) *repository {
	return &repository{
		db:          db,
		productRepo: NewProductRepository(db),
		cartRepo:    NewCartRepository(db),
		paymentRepo: NewPaymentRepository(db),
		userRepo:    NewUserRepository(db),
	}
}

func (repo *repository) Product() ProductRepository {
	return repo.productRepo
}

func (repo *repository) Cart() CartRepository {
	return repo.cartRepo
}

func (repo *repository) Payment() PaymentRepository {
	return repo.paymentRepo
}

func (repo *repository) User() UserRepository {
	return repo.userRepo
}
