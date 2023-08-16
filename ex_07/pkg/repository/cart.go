package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

type cartRepo struct {
	db *DB
}

func NewCartRepository(db *DB) *cartRepo {
	return &cartRepo{db}
}

func (c *cartRepo) CreateItem(ctx context.Context, data *model.CartItem) (*model.CartItem, error) {
	tx := c.db.Begin()
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

func (c *cartRepo) UpdateItem(ctx context.Context, data *model.CartItem) (*model.CartItem, error) {
	tx := c.db.Begin()
	if err := tx.Where("cart_id = ?", data.CartID).
		Where("product_id = ?", data.ProductID).
		Updates(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (c *cartRepo) DeleteItem(ctx context.Context, cartID int64, prodID int64) error {
	if err := c.db.Table(model.CartItem{}.TableName()).
		Where("cart_id = ?", cartID).
		Where("product_id = ?", prodID).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *cartRepo) GetItem(ctx context.Context, cartID int64, prodID int64) (*model.CartItem, error) {
	var item model.CartItem
	if err := c.db.
		Where("cart_id = ?", cartID).
		Where("product_id = ?", prodID).
		First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &item, nil
}

func (c *cartRepo) GetItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
	var items []model.CartItem
	if err := c.db.
		Table(model.CartItem{}.TableName()).
		Where("cart_id = ?", cartID).
		Find(&items).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return items, nil
}

func (c *cartRepo) Create(ctx context.Context) (*model.Cart, error) {
	var cart model.Cart
	tx := c.db.Begin()
	if err := tx.Create(&cart).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return &cart, nil
}

func (c *cartRepo) Update(ctx context.Context, data *model.Cart) (*model.Cart, error) {
	tx := c.db.Begin()
	if err := tx.Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}

func (c *cartRepo) Delete(ctx context.Context, id int64) error {
	if err := c.db.Table(model.Cart{}.TableName()).
		Where("id = ?", id).
		Update("active", false).
		Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *cartRepo) GetByID(ctx context.Context, id int64) (*model.Cart, error) {
	var cart model.Cart
	if err := c.db.
		Where("id = ?", id).
		First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &cart, nil
}

func (c *cartRepo) GetActiveCart(ctx context.Context) (*model.Cart, error) {
	var cart model.Cart
	if err := c.db.Where("active = ?", true).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &cart, nil
}
