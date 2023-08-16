package business

import (
	"context"
	"errors"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
	validator "github.com/hoangtk0100/dc-go-23/ex_07/pkg/validation"
)

var (
	ErrUserExisted    = errors.New("user existed")
	ErrCannotRegister = errors.New("cannot register")
)

type userBusiness struct {
	repo repository.Repository
}

func NewUserBusiness(repo repository.Repository) *userBusiness {
	return &userBusiness{repo: repo}
}

func (b *userBusiness) Register(ctx context.Context, data *model.CreateUserParams) (*model.User, error) {
	if err := validator.ValidatePassword(data.Password); err != nil {
		return nil, util.ErrBadRequest.
			WithError(err.Error()).
			WithDebug(err.Error())
	}

	existedUser, err := b.repo.User().GetByUsername(ctx, data.Username)
	if err == nil && existedUser != nil {
		return nil, util.ErrConflict.WithError(ErrUserExisted.Error())
	}

	salt, err := util.RandomString(8)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotRegister.Error()).
			WithDebug(err.Error())
	}

	hashedPassword, err := util.HashPassword("", data.Password, salt)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotRegister.Error()).
			WithDebug(err.Error())
	}

	params := &model.User{
		Username:       data.Username,
		HashedPassword: hashedPassword,
		Salt:           salt,
		Email:          data.Email,
		FullName:       data.FullName,
	}

	user, err := b.repo.User().Create(ctx, params)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrCannotRegister.Error()).
			WithDebug(err.Error())
	}

	return user, nil
}
