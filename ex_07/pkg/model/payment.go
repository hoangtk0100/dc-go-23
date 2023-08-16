package model

import "github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"

type Payment struct {
	Base
	CartID        int64                  `json:"cart_id" gorm:"column:cart_id;" binding:"required,min=1"`
	Discount      float64                `json:"discount" gorm:"column:discount;" binding:"required,min=0"`
	Total         float64                `json:"total" gorm:"column:total;" binding:"required,min=0"`
	Currency      constant.Currency      `json:"currency" gorm:"column:currency;default:USD;" binding:"required,currency"`
	Status        constant.PaymentStatus `json:"status" gorm:"column:status;default:PENDING;"`
	OwnerUsername string                 `json:"owner_username" gorm:"column:owner_username;"`
	Note          string                 `json:"note,omitempty" gorm:"column:note;"`
}

func (Payment) TableName() string {
	return "payments"
}
