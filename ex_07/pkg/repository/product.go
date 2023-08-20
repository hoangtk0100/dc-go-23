package repository

import (
	"context"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
)

type productRepo struct {
	db *DB
}

func NewProductRepository(db *DB) *productRepo {
	return &productRepo{db}
}

func (p *productRepo) Create(ctx context.Context, data *model.Product) (*model.Product, error) {
	tx := p.db.Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (p *productRepo) Update(ctx context.Context, id int64, data *model.Product) (*model.Product, error) {
	tx := p.db.Begin()
	if err := tx.Where("id = ?", id).Updates(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (p *productRepo) DeleteByID(ctx context.Context, id int64) error {
	if err := p.db.Table(model.Product{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *productRepo) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	var prod model.Product
	if err := p.db.Where("id = ?", id).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &prod, nil
}

func (p *productRepo) GetByCode(ctx context.Context, code string) (*model.Product, error) {
	var prod model.Product
	if err := p.db.Where("code = ?", code).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &prod, nil
}

func (p *productRepo) GetBySlug(ctx context.Context, slug string) (*model.Product, error) {
	var prod model.Product
	if err := p.db.Where("slug = ?", slug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &prod, nil
}

func (p *productRepo) GetAll(ctx context.Context) ([]model.Product, error) {
	var result []model.Product
	if err := p.db.Table(model.Product{}.TableName()).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
