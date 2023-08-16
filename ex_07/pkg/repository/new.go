package repository

type repository struct {
	db          *DB
	productRepo ProductRepository
	cartRepo    CartRepository
	paymentRepo PaymentRepository
	userRepo    UserRepository
}

func NewRepository(db *DB) *repository {
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
