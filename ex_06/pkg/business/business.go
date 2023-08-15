package business

import (
	"context"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
)

type Business interface {
	Product() ProductBusiness
	Cart() CartBusiness
}

type ProductBusiness interface {
	Create(ctx context.Context, data *model.CreateProductParams) (*model.Product, error)
	Update(ctx context.Context, data *model.UpdateProductParams) (*model.Product, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
}

type CartBusiness interface {
	AddItem(ctx context.Context, data *model.ModifyCartItemParams) error
	RemoveItem(ctx context.Context, data *model.ModifyCartItemParams) error
	GetByID(ctx context.Context, id int64) (interface{}, error)
	Checkout(ctx context.Context) (*model.Payment, error)
}
