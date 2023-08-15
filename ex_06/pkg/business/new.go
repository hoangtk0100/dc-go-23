package business

import "github.com/hoangtk0100/dc-go-23/ex_06/pkg/repository"

type business struct {
	repo       repository.Repository
	productBiz ProductBusiness
	cartBiz    CartBusiness
}

func NewBusiness(repo repository.Repository) *business {
	return &business{
		repo:       repo,
		productBiz: NewProductBusiness(repo),
		cartBiz:    NewCartBusiness(repo),
	}
}

func (b *business) Product() ProductBusiness {
	return b.productBiz
}

func (b *business) Cart() CartBusiness {
	return b.cartBiz
}
