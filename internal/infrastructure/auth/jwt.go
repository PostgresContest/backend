package auth

import (
	"backend/internal/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrWrongSigningMethod = errors.New("wrong signing method")
	ErrWrongFormat        = errors.New("wrong token format")
	ErrExpiredToken       = errors.New("token expired")
	ErrTokenNotValid      = errors.New("token not valid")
)

const (
	expField    = "exp"
	userIdField = "user_id"
)

type Jwt struct {
	secret        []byte
	signingString any
}

func NewJwtProvider(cfg *config.Config) (*Jwt, error) {
	secretByte := []byte(cfg.Jwt.Secret)
	return &Jwt{
		signingString: secretByte,
	}, nil
}

func (t *Jwt) Generate(userID int64) (string, error) {
	payload := jwt.MapClaims{
		expField:    time.Now().Add(24 * time.Hour).Unix(),
		userIdField: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(t.signingString)
}

func (t *Jwt) Verify(token string) (bool, int64, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongSigningMethod
		}

		tokenMap, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, ErrWrongFormat
		}

		tokenExpTime, ok := tokenMap[expField]
		tokenExpTimeUnix, formatOk := tokenExpTime.(float64)
		if !ok || !formatOk {
			return nil, ErrWrongFormat
		}

		if time.Unix(int64(tokenExpTimeUnix), 0).Before(time.Now()) {
			return nil, ErrExpiredToken
		}

		if _, ok = tokenMap[userIdField]; !ok {
			return nil, ErrWrongFormat
		}

		return t.signingString, nil
	})

	if err != nil {
		return false, 0, err
	}

	if !tkn.Valid {
		return false, 0, ErrTokenNotValid
	}

	userId, ok := tkn.Claims.(jwt.MapClaims)[userIdField]
	userIdFloat, formatOk := userId.(float64)
	if !ok || !formatOk {
		return false, 0, ErrTokenNotValid
	}

	return true, int64(userIdFloat), nil
}
