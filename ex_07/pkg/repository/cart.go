package repository

import (
	"context"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/db"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

type cartRepo struct {
	db *db.DB
}

func NewCartRepository(db *db.DB) *cartRepo {
	return &cartRepo{db}
}

func (c *cartRepo) CreateItem(ctx context.Context, data *model.ModifyCartItemParams) (*model.CartItem, error) {
	_, existed := c.getCartItem(data.ProductID)
	if existed {
		return nil, util.ErrConflict
	}

	item := &model.CartItem{
		CartID:    data.CartID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     data.Price,
		Currency:  data.Currency,
		Note:      data.Note,
	}

	c.db.DBCart.Items[item.ProductID] = *item
	return item, nil
}

func (c *cartRepo) UpdateItem(ctx context.Context, data *model.ModifyCartItemParams) (*model.CartItem, error) {
	item, existed := c.getCartItem(data.ProductID)
	if !existed {
		return nil, util.ErrNotFound
	} else {
		item.Quantity = data.Quantity
		item.Price = data.Price
		item.Currency = data.Currency
		item.Note = data.Note
	}

	c.db.DBCart.Items[item.ProductID] = *item
	return item, nil
}

func (c *cartRepo) DeleteItem(ctx context.Context, cartID int64, prodID int64) error {
	item, existed := c.getCartItem(prodID)
	if !existed {
		return util.ErrNotFound
	}

	delete(c.db.DBCart.Items, item.ProductID)
	return nil
}

func (c *cartRepo) GetItem(ctx context.Context, cartID int64, prodID int64) (*model.CartItem, error) {
	item, existed := c.getCartItem(prodID)
	if !existed {
		return nil, util.ErrNotFound
	}

	return item, nil
}

func (c *cartRepo) GetItems(ctx context.Context, cartID int64) ([]model.CartItem, error) {
	var items []model.CartItem
	for _, itemIdx := range c.db.DBCart.Items {
		items = append(items, itemIdx)
	}

	return items, nil
}

func (c *cartRepo) Create(ctx context.Context) (*model.Cart, error) {
	cart := &model.Cart{
		Active: true,
	}

	cart.ID = 1
	c.db.DBCart.Cart = cart

	return cart, nil
}

func (c *cartRepo) Update(ctx context.Context, data *model.Cart) (*model.Cart, error) {
	c.db.DBCart.Cart = data

	return data, nil
}

func (c *cartRepo) Delete(ctx context.Context, id int64) error {
	c.db.DBCart.Cart = nil
	c.db.DBCart.Items = nil

	return nil
}

func (c *cartRepo) GetByID(ctx context.Context, id int64) (*model.Cart, error) {
	if c.db.DBCart.Cart.ID != id {
		return nil, util.ErrNotFound
	}

	return c.db.DBCart.Cart, nil
}

func (c *cartRepo) GetActiveCart(ctx context.Context) (*model.Cart, error) {
	if c.db.DBCart.Cart == nil {
		return nil, util.ErrNotFound
	}

	return c.db.DBCart.Cart, nil
}

func (c *cartRepo) getCartItem(itemID int64) (*model.CartItem, bool) {
	item, existed := c.db.DBCart.Items[itemID]
	return &item, existed
}
