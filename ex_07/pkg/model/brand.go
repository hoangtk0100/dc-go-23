package model

type Brand struct {
	Base
	Code        string `json:"code"`
	Name        string `json:"name"`
	Website     string `json:"website"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
	CreatorID   int64  `json:"creator_id"`
}
