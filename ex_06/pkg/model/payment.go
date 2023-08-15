package model

type Payment struct {
	Base
	CartID   int64   `json:"cart_id" binding:"required,min=1"`
	Discount float64 `json:"discount" binding:"required,min=0"`
	Total    float64 `json:"total" binding:"required,min=0"`
	Currency string  `json:"currency" binding:"required,currency"`
	Status   string  `json:"status"`
	OwnerID  int64   `json:"owner_id"`
	Note     string  `json:"note,omitempty"`
}
