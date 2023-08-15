package repository

import (
	"context"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/db"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
)

type userRepo struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) *userRepo {
	return &userRepo{db}
}

func (u *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, existed := u.getUserFromDB(username)
	if !existed {
		return nil, util.ErrNotFound
	}

	return user, nil
}

func (u *userRepo) Create(ctx context.Context, data *model.CreateUserParams) (*model.User, error) {
	_, existed := u.getUserFromDB(data.Username)
	if existed {
		return nil, util.ErrConflict
	}

	user := mapParamsToUser(data)
	u.db.DBUser.Users[data.Username] = *user
	return user, nil
}

func (u *userRepo) getUserFromDB(username string) (*model.User, bool) {
	user, existed := u.db.DBUser.Users[username]
	return &user, existed
}

func mapParamsToUser(params *model.CreateUserParams) *model.User {
	return &model.User{
		Username:       params.Username,
		HashedPassword: params.HashedPassword,
		Salt:           params.Salt,
		Email:          params.Email,
		FullName:       params.FullName,
	}
}
