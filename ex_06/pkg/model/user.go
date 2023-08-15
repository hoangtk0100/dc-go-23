package model

import "time"

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Salt              string    `json:"-"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateUserParams struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"-"`
	Salt           string `json:"-"`
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
