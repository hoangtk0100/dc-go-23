package model

import (
	"time"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"
)

type Product struct {
	Base
	BrandID         int64                  `json:"brand_id" gorm:"column:brand_id;"`
	CategoryID      int64                  `json:"category_id" gorm:"column:category_id;"`
	Code            string                 `json:"code,omitempty" gorm:"column:code;"`
	Name            string                 `json:"name" gorm:"column:name;"`
	Quantity        int64                  `json:"quantity" gorm:"column:quantity;"`
	Weight          float64                `json:"weight" gorm:"column:weight;"`
	WeightUnit      constant.WeightUnit    `json:"weight_unit" gorm:"column:weight_unit;default:KG;"`
	Price           float64                `json:"price" gorm:"column:price;"`
	Currency        constant.Currency      `json:"currency" gorm:"column:currency;default:USD;"`
	Description     string                 `json:"description,omitempty" gorm:"column:description;"`
	Slug            string                 `json:"slug" gorm:"column:slug;"`
	Sold            int64                  `json:"sold" gorm:"column:sold;"`
	Rate            int64                  `json:"rate" gorm:"column:rate;"`
	Reviews         int64                  `json:"reviews" gorm:"column:reviews;"`
	Status          constant.ProductStatus `json:"status" gorm:"column:status;default:ACTIVE;"`
	OldStatus       string                 `json:"old_status" gorm:"column:old_status;"`
	CreatorUsername string                 `json:"creator_username" gorm:"column:creator_username;"`
	DeletedAt       *time.Time             `json:"deleted_at" gorm:"column:deleted_at;"`
}

func (Product) TableName() string {
	return "products"
}

type CreateProductParams struct {
	Name        string  `json:"name" gorm:"column:name;" binding:"required"`
	Code        string  `json:"code,omitempty" gorm:"column:code;"`
	BrandID     int64   `json:"brand_id" gorm:"column:brand_id;"`
	CategoryID  int64   `json:"category_id" gorm:"column:category_id;"`
	Quantity    int64   `json:"quantity" gorm:"column:quantity;" binding:"required,min=0"`
	Weight      float64 `json:"weight" gorm:"column:weight;" binding:"required,gt=0"`
	WeightUnit  string  `json:"weight_unit" gorm:"column:weight_unit;default:KG;" binding:"required,weight_unit"`
	Price       float64 `json:"price" gorm:"column:price;" binding:"required,min=0"`
	Currency    string  `json:"currency" gorm:"column:currency;default:USD;" binding:"required,currency"`
	Description string  `json:"description,omitempty" gorm:"column:description;"`
	Slug        string  `json:"slug,omitempty" gorm:"column:slug;"`
}

func (CreateProductParams) TableName() string {
	return Product{}.TableName()
}

type UpdateProductParams struct {
	Name        string  `json:"name" gorm:"column:name;" binding:"required"`
	Code        string  `json:"code,omitempty" gorm:"column:code;"`
	BrandID     int64   `json:"brand_id" gorm:"column:brand_id;"`
	CategoryID  int64   `json:"category_id" gorm:"column:category_id;"`
	Quantity    int64   `json:"quantity" gorm:"column:quantity;" binding:"required,min=0"`
	Weight      float64 `json:"weight" gorm:"column:weight;" binding:"required,gt=0"`
	WeightUnit  string  `json:"weight_unit" gorm:"column:weight_unit;default:KG;" binding:"required,weight_unit"`
	Price       float64 `json:"price" gorm:"column:price;" binding:"required,min=0"`
	Currency    string  `json:"currency" gorm:"column:currency;default:USD;" binding:"required,currency"`
	Description string  `json:"description,omitempty" gorm:"column:description;"`
	Slug        string  `json:"slug,omitempty" gorm:"column:slug;"`
}

func (UpdateProductParams) TableName() string {
	return Product{}.TableName()
}
