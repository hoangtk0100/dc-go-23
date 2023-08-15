package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	minJWTSecretKeySize          = 32
	defaultAccessTokenExpiresIn  = time.Hour * 24 * 7
	defaultRefreshTokenExpiresIn = time.Hour * 24 * 7 * 2
)

var (
	ErrInvalidJWTKeySize     = errors.New(fmt.Sprintf("Invalid key size: must be at least %v characters", minJWTSecretKeySize))
	ErrMissingCustomDuration = errors.New("Duration must be provided for CustomToken")
	ErrTooManyCustomDuration = errors.New("Provide too many durations")
	ErrInvalidTokenType      = errors.New("Invalid token type")
)

// jwtMaker is a JSON Web Token maker
// Symmetric key algorithm to sign the key
type jwtMaker struct {
	secretKey             string
	accessTokenExpiresIn  time.Duration
	refreshTokenExpiresIn time.Duration
}

func NewJWTMaker(secretKey string, accessTokenExpiresIn time.Duration, refreshTokenExpiresIn time.Duration) (*jwtMaker, error) {
	if len(secretKey) < minJWTSecretKeySize {
		return nil, errors.WithStack(ErrInvalidJWTKeySize)
	}

	return &jwtMaker{
		secretKey:             secretKey,
		accessTokenExpiresIn:  accessTokenExpiresIn,
		refreshTokenExpiresIn: refreshTokenExpiresIn,
	}, nil
}

func (maker *jwtMaker) getTokenDuration(tokenType TokenType, duration ...time.Duration) (time.Duration, error) {
	var tokenDuration time.Duration

	switch tokenType {
	case AccessToken:
		if maker.accessTokenExpiresIn == 0 {
			tokenDuration = defaultAccessTokenExpiresIn
		} else {
			tokenDuration = maker.accessTokenExpiresIn
		}
	case RefreshToken:
		if maker.refreshTokenExpiresIn == 0 {
			tokenDuration = defaultRefreshTokenExpiresIn
		} else {
			tokenDuration = maker.refreshTokenExpiresIn
		}
	case CustomToken:
		if len(duration) == 0 {
			return 0, errors.WithStack(ErrMissingCustomDuration)
		} else if len(duration) > 1 {
			return 0, errors.WithStack(ErrTooManyCustomDuration)
		}

		tokenDuration = duration[0]
	default:
		return 0, errors.WithStack(ErrInvalidTokenType)
	}

	return tokenDuration, nil
}

func (maker *jwtMaker) CreateToken(tokenType TokenType, uid string, duration ...time.Duration) (string, *Payload, error) {
	tokenDuration, err := maker.getTokenDuration(tokenType, duration...)
	if err != nil {
		return "", nil, err
	}

	payload, err := NewPayload(uid, tokenDuration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, errors.WithStack(err)
}

func (maker *jwtMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.WithStack(ErrInvalidToken)
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, errors.WithStack(ErrExpiredToken)
		}

		return nil, errors.WithStack(ErrInvalidToken)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.WithStack(ErrInvalidToken)
	}

	return payload, nil
}
