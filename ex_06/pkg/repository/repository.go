package repository

import (
	"context"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
)

type Repository interface {
	Product() ProductRepository
	Cart() CartRepository
	Payment() PaymentRepository
	User() UserRepository
}

type ProductRepository interface {
	Create(ctx context.Context, data *model.CreateProductParams) (*model.Product, error)
	Update(ctx context.Context, data *model.UpdateProductParams) (*model.Product, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
}

type CartRepository interface {
	CreateItem(ctx context.Context, data *model.ModifyCartItemParams) (*model.CartItem, error)
	UpdateItem(ctx context.Context, data *model.ModifyCartItemParams) (*model.CartItem, error)
	DeleteItem(ctx context.Context, cartID int64, prodID int64) error
	GetItem(ctx context.Context, cartID int64, prodID int64) (*model.CartItem, error)
	GetItems(ctx context.Context, cartID int64) ([]model.CartItem, error)
	Create(ctx context.Context) (*model.Cart, error)
	Update(ctx context.Context, data *model.Cart) (*model.Cart, error)
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Cart, error)
	GetActiveCart(ctx context.Context) (*model.Cart, error)
}

type PaymentRepository interface {
	Create(ctx context.Context, data *model.Payment) (*model.Payment, error)
}

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, data *model.CreateUserParams) (*model.User, error)
}
