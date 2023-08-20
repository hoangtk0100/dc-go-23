package business

import (
	"context"
	"errors"
	"strings"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
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

	code := b.generateCode(ctx, data.Code)
	slug := b.generateSlug(ctx, data.Slug)

	params := &model.Product{
		Name:            data.Name,
		Code:            code,
		Quantity:        data.Quantity,
		Weight:          data.Weight,
		WeightUnit:      constant.WeightUnit(data.WeightUnit),
		Price:           data.Price,
		Currency:        constant.Currency(data.Currency),
		Description:     data.Description,
		Slug:            slug,
		CreatorUsername: util.GetRequester(ctx).GetUID(),
	}

	prod, err := b.repo.Product().Create(ctx, params)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotCreateProduct.Error()).
			WithDebug(err.Error())
	}

	return prod, nil
}

func (b *productBusiness) Update(ctx context.Context, id int64, data *model.UpdateProductParams) (*model.Product, error) {
	prod, err := b.validateExistedProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	// TODO: Check if brand and category exist
	prod.Name = data.Name
	prod.Code = data.Code
	prod.Quantity = data.Quantity
	prod.Weight = data.Weight
	prod.WeightUnit = constant.WeightUnit(data.WeightUnit)
	prod.Price = data.Price
	prod.Currency = constant.Currency(data.Currency)
	prod.Description = data.Description
	prod.Slug = data.Slug

	prod, err = b.repo.Product().Update(ctx, id, prod)
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

	if prod.Status == constant.ProductStatusDeleted {
		return nil, util.ErrBadRequest.
			WithError(ErrProductDeleted.Error()).
			WithDebug(err.Error())
	}

	return prod, nil
}

func (b *productBusiness) generateSlug(ctx context.Context, name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	for {
		_, err := b.repo.Product().GetBySlug(ctx, name)
		if err != nil {
			return name
		}

		randStr, _ := util.RandomString(8)
		name = name + randStr
	}
}

func (b *productBusiness) generateCode(ctx context.Context, code string) string {
	if code == "" {
		randStr, _ := util.RandomString(8)
		code = randStr
	}

	for {
		_, err := b.repo.Product().GetByCode(ctx, code)
		if err != nil {
			return code
		}

		randStr, _ := util.RandomString(8)
		code = randStr
	}
}
