package model

import "time"

type Product struct {
	Base
	BrandID     int64      `json:"brand_id"`
	CategoryID  int64      `json:"category_id"`
	Code        string     `json:"code,omitempty"`
	Name        string     `json:"name" binding:"required"`
	Quantity    int64      `json:"quantity" binding:"required,min=0"`
	Weight      float64    `json:"weight" binding:"required,gt=0"`
	WeightUnit  string     `json:"weight_unit" binding:"required,weight_unit"`
	Price       float64    `json:"price" binding:"required,min=0"`
	Currency    string     `json:"currency" binding:"required,currency"`
	Description string     `json:"description,omitempty"`
	Slug        string     `json:"slug"`
	Sold        int64      `json:"sold"`
	Rate        int64      `json:"rate"`
	Reviews     int64      `json:"reviews"`
	Status      string     `json:"status"`
	OldStatus   string     `json:"old_status"`
	CreatorID   int64      `json:"creator_id"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CreateProductParams struct {
	Name        string  `json:"name" binding:"required"`
	Code        string  `json:"code,omitempty"`
	BrandID     int64   `json:"brand_id"`
	CategoryID  int64   `json:"category_id"`
	Quantity    int64   `json:"quantity" binding:"required,min=0"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	WeightUnit  string  `json:"weight_unit" binding:"required,weight_unit"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Currency    string  `json:"currency" binding:"required,currency"`
	Description string  `json:"description,omitempty"`
	Slug        string  `json:"slug,omitempty"`
}

type UpdateProductParams struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code,omitempty"`
	BrandID     int64   `json:"brand_id"`
	CategoryID  int64   `json:"category_id"`
	Quantity    int64   `json:"quantity" binding:"required,min=0"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	WeightUnit  string  `json:"weight_unit" binding:"required,weight_unit"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Currency    string  `json:"currency" binding:"required,currency"`
	Description string  `json:"description,omitempty"`
	Slug        string  `json:"slug,omitempty"`
}
