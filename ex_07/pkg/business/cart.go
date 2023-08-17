package business

import (
	"context"
	"errors"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

var (
	ErrHavenNoActiveCart    = errors.New("haven no active cart")
	ErrCannotGetCart        = errors.New("cannot get cart")
	ErrCannotCheckout       = errors.New("cannot checkout")
	ErrProductNotFound      = errors.New("product not found")
	ErrNegativeQuantity     = errors.New("quantity must be at least 1")
	ErrCannotAddCartItem    = errors.New("cannot add cart item")
	ErrCannotRemoveCartItem = errors.New("cannot remove cart item")
	ErrNoItems              = errors.New("cart has no items")
)

type cartBusiness struct {
	repo repository.Repository
}

func NewCartBusiness(repo repository.Repository) *cartBusiness {
	return &cartBusiness{repo: repo}
}

func (c *cartBusiness) AddItem(ctx context.Context, data *model.ModifyCartItemParams) error {
	prod, err := c.repo.Product().GetByID(ctx, data.ProductID)
	if err != nil {
		return util.ErrBadRequest.
			WithError(ErrProductNotFound.Error()).
			WithDebug(err.Error())
	}

	if data.Quantity < 1 {
		return util.ErrBadRequest.
			WithError(ErrNegativeQuantity.Error())
	}

	if prod.Status == constant.ProductStatusDeleted {
		return util.ErrBadRequest.
			WithError(ErrProductDeleted.Error())
	}

	curUsername := util.GetRequester(ctx).GetUID()
	cart, err := c.repo.Cart().GetActiveCart(ctx)
	if err != nil {
		cart, err = c.repo.Cart().Create(ctx, curUsername)
		if err != nil {
			return util.ErrInternalServerError.
				WithError(ErrCannotAddCartItem.Error()).
				WithDebug(err.Error())
		}
	}

	params := &model.CartItem{
		CartID:    cart.ID,
		ProductID: data.ProductID,
		Quantity:  data.Quantity,
		Price:     prod.Price,
		Currency:  prod.Currency,
		Note:      data.Note,
	}

	quantity := data.Quantity
	item, err := c.repo.Cart().GetItem(ctx, cart.ID, params.ProductID)
	if err != nil && util.ErrNotFound.Is(err) {
		item, err = c.repo.Cart().CreateItem(ctx, params)
		if err != nil {
			return util.ErrInternalServerError.
				WithError(ErrCannotAddCartItem.Error()).
				WithDebug(err.Error())
		}
	} else {
		params.Quantity += item.Quantity
		item, err = c.repo.Cart().UpdateItem(ctx, params)
		if err != nil {
			return util.ErrInternalServerError.
				WithError(ErrCannotAddCartItem.Error()).
				WithDebug(err.Error())
		}
	}

	_, err = c.updateCart(ctx, cart, prod, quantity)
	if err != nil {
		return util.ErrInternalServerError.
			WithError(ErrCannotAddCartItem.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (c *cartBusiness) RemoveItem(ctx context.Context, data *model.ModifyCartItemParams) error {
	cart, err := c.repo.Cart().GetActiveCart(ctx)
	if err != nil {
		return util.ErrBadRequest.
			WithError(ErrHavenNoActiveCart.Error()).
			WithDebug(err.Error())
	}

	prod, err := c.repo.Product().GetByID(ctx, data.ProductID)
	if err != nil {
		return util.ErrBadRequest.
			WithError(ErrProductNotFound.Error()).
			WithDebug(err.Error())
	}

	quantity := data.Quantity
	item, err := c.repo.Cart().GetItem(ctx, cart.ID, data.ProductID)
	if err != nil && util.ErrNotFound.Is(err) {
		return util.ErrBadRequest.
			WithError(ErrCannotRemoveCartItem.Error()).
			WithDebug(err.Error())
	}

	if item.Quantity == 1 {
		err := c.repo.Cart().DeleteItem(ctx, cart.ID, prod.ID)
		if err != nil {
			return util.ErrInternalServerError.
				WithError(ErrCannotRemoveCartItem.Error()).
				WithDebug(err.Error())
		}
	} else {
		params := &model.CartItem{
			CartID:    cart.ID,
			ProductID: data.ProductID,
			Quantity:  item.Quantity - quantity,
			Price:     prod.Price,
			Currency:  prod.Currency,
			Note:      data.Note,
		}

		item, err = c.repo.Cart().UpdateItem(ctx, params)
		if err != nil {
			return util.ErrInternalServerError.
				WithError(ErrCannotAddCartItem.Error()).
				WithDebug(err.Error())
		}
	}

	_, err = c.updateCart(ctx, cart, prod, -quantity)
	if err != nil {
		return util.ErrInternalServerError.
			WithError(ErrCannotAddCartItem.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (c *cartBusiness) GetByID(ctx context.Context, id *int64) (interface{}, error) {
	var cart *model.Cart
	var err error
	if id == nil || *id == 0 {
		cart, err = c.repo.Cart().GetActiveCart(ctx)
		if err != nil {
			return nil, util.ErrBadRequest.
				WithError(ErrHavenNoActiveCart.Error()).
				WithDebug(err.Error())
		}
	} else {
		cart, err = c.repo.Cart().GetByID(ctx, *id)
		if err != nil {
			if util.ErrNotFound.Is(err) {
				return nil, util.ErrNotFound.
					WithError(ErrCannotGetCart.Error()).
					WithDebug(err.Error())
			} else {
				return nil, util.ErrInternalServerError.
					WithError(ErrCannotGetCart.Error()).
					WithDebug(err.Error())
			}
		}
	}

	items, err := c.repo.Cart().GetItems(ctx, cart.ID)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotGetCart.Error()).
			WithDebug(err.Error())
	}

	result := cartResponse{
		Cart:  cart,
		Items: items,
	}

	return result, nil
}

type cartResponse struct {
	Cart  *model.Cart
	Items []model.CartItem
}

func (c *cartBusiness) Checkout(ctx context.Context) (*model.Payment, error) {
	cart, err := c.repo.Cart().GetActiveCart(ctx)
	if err != nil {
		return nil, util.ErrNotFound.
			WithError(ErrHavenNoActiveCart.Error()).
			WithDebug(err.Error())
	}

	items, err := c.repo.Cart().GetItems(ctx, cart.ID)
	if err != nil || len(items) == 0 {
		return nil, util.ErrBadRequest.
			WithError(ErrNoItems.Error()).
			WithDebug(err.Error())
	}

	var total float64
	for idx := range items {
		total += items[idx].Price * float64(items[idx].Quantity)
	}

	curUsername := util.GetRequester(ctx).GetUID()
	params := &model.Payment{
		CartID:        cart.ID,
		Discount:      0,
		Total:         total,
		Currency:      cart.Currency,
		Note:          cart.Note,
		OwnerUsername: curUsername,
	}

	payment, err := c.repo.Payment().Create(ctx, params)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotCheckout.Error()).
			WithDebug(err.Error())
	}

	err = c.repo.Cart().Delete(ctx, cart.ID)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotCheckout.Error()).
			WithDebug(err.Error())
	}

	return payment, nil
}

func (c *cartBusiness) updateCart(ctx context.Context, cart *model.Cart, prod *model.Product, quantity int64) (*model.Cart, error) {
	cart.Quantity += quantity
	cart.Total += prod.Price * float64(quantity)
	cart.Weight += prod.Weight * float64(quantity)
	cart.Currency = prod.Currency

	return c.repo.Cart().Update(ctx, cart)
}
