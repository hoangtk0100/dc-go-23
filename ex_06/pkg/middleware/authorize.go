package middleware

import (
	"context"
	"errors"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/model"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/token"
	"github.com/hoangtk0100/dc-go-23/ex_06/pkg/util"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
)

var (
	ErrAuthHeaderEmpty           = errors.New("authorization header is not provided")
	ErrAuthHeaderInvalidFormat   = errors.New("invalid authorization header format")
	ErrAuthHeaderUnsupportedType = errors.New("unsupported authorization type")
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

func extractTokenFromHeader(input string) (string, error) {
	if len(input) == 0 {
		return "", ErrAuthHeaderEmpty
	}

	parts := strings.Fields(input)
	if len(parts) < 2 {
		return "", ErrAuthHeaderInvalidFormat
	}

	authorizationType := strings.ToLower(parts[0])
	if authorizationType != authorizationTypeBearer {
		return "", ErrAuthHeaderUnsupportedType
	}

	return parts[1], nil
}

func RequireAuth(repo UserRepository, tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := extractTokenFromHeader(ctx.GetHeader(authorizationHeaderKey))
		if err != nil {
			util.ErrorResponse(ctx, util.ErrUnauthorized.WithError(err.Error()))
			ctx.Abort()
			return
		}

		payload, err := tokenMaker.VerifyToken(authToken)
		if err != nil {
			util.ErrorResponse(ctx, util.ErrUnauthorized.WithError(err.Error()).WithDebug(err.Error()))
			ctx.Abort()
			return
		}

		user, err := repo.GetByUsername(ctx.Request.Context(), payload.UID)
		if err != nil {
			util.ErrorResponse(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(util.KeyRequester, util.NewRequester(user.Username))
		ctx.Next()
	}
}
