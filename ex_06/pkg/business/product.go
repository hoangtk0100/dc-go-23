package business

import (
	"context"
	"errors"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/constant"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
)

var (
	ErrCannotCreateProduct  = errors.New("cannot create product")
	ErrCannotUpdateProduct  = errors.New("cannot update product")
	ErrCannotDeleteProduct  = errors.New("cannot delete product")
	ErrCannotGetProduct     = errors.New("cannot get product")
	ErrProductDeleted       = errors.New("product is deleted")
	ErrCannotGetProductList = errors.New("cannot get all products")
)

type productBusiness struct {
	repo repository.Repository
}

func NewProductBusiness(repo repository.Repository) *productBusiness {
	return &productBusiness{repo: repo}
}

func (b *productBusiness) Create(ctx context.Context, data *model.CreateProductParams) (*model.Product, error) {
	// TODO: Check if brand and category exist
	prod, err := b.repo.Product().Create(ctx, data)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotCreateProduct.Error()).
			WithDebug(err.Error())
	}

	return prod, nil
}

func (b *productBusiness) Update(ctx context.Context, data *model.UpdateProductParams) (*model.Product, error) {
	_, err := b.validateExistedProduct(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	// TODO: Check if brand and category exist

	prod, err := b.repo.Product().Update(ctx, data)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotUpdateProduct.Error()).
			WithDebug(err.Error())
	}

	return prod, nil
}

func (b *productBusiness) DeleteByID(ctx context.Context, id int64) error {
	_, err := b.validateExistedProduct(ctx, id)
	if err != nil {
		return err
	}

	err = b.repo.Product().DeleteByID(ctx, id)
	if err != nil {
		return util.ErrInternalServerError.
			WithError(ErrCannotDeleteProduct.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (b *productBusiness) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	prod, err := b.validateExistedProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return prod, nil
}

func (b *productBusiness) GetAll(ctx context.Context) ([]model.Product, error) {
	prods, err := b.repo.Product().GetAll(ctx)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotGetProductList.Error()).
			WithDebug(err.Error())
	}

	return prods, nil
}

func (b *productBusiness) validateExistedProduct(ctx context.Context, id int64) (*model.Product, error) {
	prod, err := b.repo.Product().GetByID(ctx, id)
	if err != nil {
		if util.ErrNotFound.Is(err) {
			return nil, util.ErrNotFound.
				WithError(ErrCannotGetProduct.Error()).
				WithDebug(err.Error())
		}

		return nil, util.ErrInternalServerError.
			WithError(ErrCannotGetProduct.Error()).
			WithDebug(err.Error())
	}

	if constant.ProductStatus(prod.Status) == constant.ProductStatusDeleted {
		return nil, util.ErrBadRequest.
			WithError(ErrProductDeleted.Error()).
			WithDebug(err.Error())
	}

	return prod, nil
}
