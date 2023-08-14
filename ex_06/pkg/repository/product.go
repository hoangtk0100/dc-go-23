package repository

import (
	"context"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/db"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
)

type productRepo struct {
	db *db.DB
}

func NewProductRepository(db *db.DB) *productRepo {
	return &productRepo{db}
}

func (p *productRepo) Create(ctx context.Context, data *model.CreateProductParams) (*model.Product, error) {
	prod := &model.Product{
		Name:        data.Name,
		Code:        data.Code,
		Quantity:    data.Quantity,
		Weight:      data.Weight,
		WeightUnit:  data.WeightUnit,
		Price:       data.Price,
		Currency:    data.Currency,
		Description: data.Description,
		Slug:        data.Slug,
	}

	prod.ID = p.getLastIdx() + 1

	p.db.DBProducts.Products[prod.ID] = *prod
	p.db.DBProducts.LastIdx = prod.ID
	return prod, nil
}

func (p *productRepo) Update(ctx context.Context, data *model.UpdateProductParams) (*model.Product, error) {
	prod, existed := p.getProduct(data.ID)
	if !existed {
		return nil, util.ErrNotFound
	}

	prod.Name = data.Name
	prod.Code = data.Code
	prod.Quantity = data.Quantity
	prod.Weight = data.Weight
	prod.WeightUnit = data.WeightUnit
	prod.Price = data.Price
	prod.Currency = data.Currency
	prod.Description = data.Description
	prod.Slug = data.Slug

	p.db.DBProducts.Products[prod.ID] = *prod
	return prod, nil
}

func (p *productRepo) DeleteByID(ctx context.Context, id int64) error {
	prod, existed := p.getProduct(id)
	if !existed {
		return util.ErrNotFound
	}

	delete(p.db.DBProducts.Products, prod.ID)
	return nil
}

func (p *productRepo) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	prod, existed := p.getProduct(id)
	if !existed {
		return nil, util.ErrNotFound
	}

	return prod, nil
}

func (p *productRepo) GetAll(ctx context.Context) ([]model.Product, error) {
	var result []model.Product
	for _, prodIdx := range p.db.DBProducts.Products {
		result = append(result, prodIdx)
	}

	return result, nil
}

func (p *productRepo) getProduct(id int64) (*model.Product, bool) {
	prod, existed := p.db.DBProducts.Products[id]
	return &prod, existed
}

func (p *productRepo) getLastIdx() int64 {
	return p.db.DBProducts.LastIdx
}
