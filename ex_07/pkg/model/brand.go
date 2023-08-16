package model

type Brand struct {
	Base
	Code            string `json:"code"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	Active          bool   `json:"active"`
	Description     string `json:"description"`
	CreatorUsername string `json:"creator_username"`
}
