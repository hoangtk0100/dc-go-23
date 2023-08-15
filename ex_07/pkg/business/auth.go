package business

import (
	"context"
	"errors"
	"time"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/repository"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/token"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/util"
	validator "github.com/hoangtk0100/dc-go-23/ex_07/pkg/validation"
)

var (
	ErrEmailOrPasswordInvalid = errors.New("email or password invalid")
)

type authBusiness struct {
	repo       repository.Repository
	tokenMaker token.TokenMaker
}

func NewAuthBusiness(repo repository.Repository, tokenMaker token.TokenMaker) *authBusiness {
	return &authBusiness{
		repo:       repo,
		tokenMaker: tokenMaker,
	}
}

type loginResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (a *authBusiness) Login(ctx context.Context, data *model.LoginParams) (interface{}, error) {
	user, err := a.repo.User().GetByUsername(ctx, data.Username)
	if err != nil {
		if util.ErrNotFound.Is(err) {
			return nil, util.ErrNotFound.
				WithError(ErrEmailOrPasswordInvalid.Error())
		}

		return nil, util.ErrInternalServerError.
			WithError(ErrEmailOrPasswordInvalid.Error()).
			WithDebug(err.Error())
	}

	if err := validator.ValidatePassword(data.Password); err != nil {
		return nil, util.ErrBadRequest.
			WithError(err.Error()).
			WithDebug(err.Error())
	}

	err = util.CheckPassword("", user.HashedPassword, data.Password, user.Salt)
	if err != nil {
		return nil, util.ErrUnauthorized.
			WithError(ErrEmailOrPasswordInvalid.Error()).
			WithDebug(err.Error())
	}

	accessToken, accessPayload, err := a.tokenMaker.CreateToken(token.AccessToken, user.Username)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrEmailOrPasswordInvalid.Error()).
			WithDebug(err.Error())
	}

	refreshToken, refreshPayload, err := a.tokenMaker.CreateToken(token.RefreshToken, user.Username)
	if err != nil {
		return nil, util.ErrInternalServerError.
			WithError(ErrEmailOrPasswordInvalid.Error()).
			WithDebug(err.Error())
	}

	resp := loginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}

	return resp, nil
}
