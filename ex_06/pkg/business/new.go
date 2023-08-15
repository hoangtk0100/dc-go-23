package business

import (
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/token"
)

type business struct {
	repo       repository.Repository
	productBiz ProductBusiness
	cartBiz    CartBusiness
	userBiz    UserBusiness
	authBiz    AuthBusiness
}

func NewBusiness(repo repository.Repository, tokenMaker token.TokenMaker) *business {
	return &business{
		repo:       repo,
		productBiz: NewProductBusiness(repo),
		cartBiz:    NewCartBusiness(repo),
		userBiz:    NewUserUserBusiness(repo),
		authBiz:    NewAuthBusiness(repo, tokenMaker),
	}
}

func (b *business) Product() ProductBusiness {
	return b.productBiz
}

func (b *business) Cart() CartBusiness {
	return b.cartBiz
}

func (b *business) User() UserBusiness {
	return b.userBiz
}

func (b *business) Auth() AuthBusiness {
	return b.authBiz
}
