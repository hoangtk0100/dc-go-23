package model

import (
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"
	"time"
)

type Cart struct {
	Base
	Code          string              `json:"code" gorm:"column:code;"`
	Quantity      int64               `json:"quantity" gorm:"column:quantity;" binding:"required,min=0"`
	Weight        float64             `json:"weight" gorm:"column:weight;" binding:"required,gt=0"`
	WeightUnit    constant.WeightUnit `json:"weight_unit" gorm:"column:weight_unit;default:KG;" binding:"required,weight_unit"`
	Total         float64             `json:"total" gorm:"column:total;" binding:"required,min=0"`
	Currency      constant.Currency   `json:"currency" gorm:"column:currency;default:USD;" binding:"required,currency"`
	Note          string              `json:"note,omitempty" gorm:"column:note;"`
	Active        bool                `json:"active" gorm:"column:active;default:true;"`
	OwnerUsername string              `json:"owner_username" gorm:"column:owner_username;"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartItem struct {
	CartID    int64             `json:"cart_id" gorm:"column:cart_id;"`
	ProductID int64             `json:"product_id" gorm:"column:product_id;"`
	Quantity  int64             `json:"quantity" gorm:"column:quantity;" binding:"required,min=0"`
	Price     float64           `json:"price" gorm:"column:price;" binding:"required,min=0"`
	Total     float64           `json:"total" gorm:"column:total;" binding:"required,min=0"`
	Currency  constant.Currency `json:"currency" gorm:"column:currency;default:USD;" binding:"required,currency"`
	Note      string            `json:"note,omitempty" gorm:"column:note;"`
	CreatedAt *time.Time        `json:"created_at" gorm:"column:created_at;"`
}

func (CartItem) TableName() string {
	return "cart_items"
}

type ModifyCartItemParams struct {
	ProductID int64  `json:"product_id" binding:"required,min=1"`
	Quantity  int64  `json:"quantity" gorm:"column:quantity;" binding:"required,min=0"`
	Note      string `json:"note,omitempty" gorm:"column:note;"`
}

func (ModifyCartItemParams) TableName() string {
	return CartItem{}.TableName()
}
