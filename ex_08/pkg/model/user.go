package model

import "time"

type User struct {
	Username          string    `json:"username" gorm:"column:username;"`
	HashedPassword    string    `json:"-" gorm:"column:hashed_password;"`
	Salt              string    `json:"-" gorm:"column:salt;"`
	FullName          string    `json:"full_name" gorm:"column:full_name;"`
	Email             string    `json:"email" gorm:"column:email;"`
	PasswordChangedAt time.Time `json:"password_changed_at" gorm:"column:password_changed_at;"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (User) TableName() string {
	return "users"
}

type CreateUserParams struct {
	Username string `json:"username" gorm:"column:username;" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" gorm:"column:full_name;" binding:"required"`
	Email    string `json:"email" gorm:"column:email;" binding:"required"`
}

func (CreateUserParams) TableName() string {
	return User{}.TableName()
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (LoginParams) TableName() string {
	return User{}.TableName()
}
