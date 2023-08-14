package model

type Category struct {
	Base
	ParentID    int64  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatorID   int64  `json:"creator_id"`
}
