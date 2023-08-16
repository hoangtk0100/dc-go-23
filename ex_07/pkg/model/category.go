package model

type Category struct {
	Base
	ParentID        int64  `json:"parent_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatorUsername string `json:"creator_username"`
}
