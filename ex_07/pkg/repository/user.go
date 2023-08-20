package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
)

type userRepo struct {
	db *DB
}

func NewUserRepository(db *DB) *userRepo {
	return &userRepo{db}
}

func (u *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ErrNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (u *userRepo) Create(ctx context.Context, data *model.User) (*model.User, error) {
	tx := u.db.Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}
