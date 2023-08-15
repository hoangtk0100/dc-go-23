package model

import "time"

type Cart struct {
	Base
	Code       string  `json:"code"`
	Quantity   int64   `json:"quantity" binding:"required,min=0"`
	Weight     float64 `json:"weight" binding:"required,gt=0"`
	WeightUnit string  `json:"weight_unit" binding:"required,weight_unit"`
	Total      float64 `json:"total" binding:"required,min=0"`
	Currency   string  `json:"currency" binding:"required,currency"`
	Note       string  `json:"note,omitempty"`
	Active     bool    `json:"active"`
	OwnerID    int64   `json:"owner_id"`
}

type CartItem struct {
	CartID    int64      `json:"cart_id"`
	ProductID int64      `json:"product_id"`
	Quantity  int64      `json:"quantity" binding:"required,min=0"`
	Price     float64    `json:"price" binding:"required,min=0"`
	Total     float64    `json:"total" binding:"required,min=0"`
	Currency  string     `json:"currency" binding:"required,currency"`
	Note      string     `json:"note,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ModifyCartItemParams struct {
	CartID    int64   `json:"-"`
	ProductID int64   `json:"product_id" binding:"required,min=1"`
	Quantity  int64   `json:"quantity" binding:"required,min=0"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Note      string  `json:"note,omitempty"`
}
