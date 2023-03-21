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
	userIdField           = "user_id"
)

type Jwt struct {
	secret        []byte
	signingString any
}

func NewJwtProvider(cfg *config.Config) (*Jwt, error) {
	secretByte := []byte(cfg.Jwt.Secret)
	rsaSigningString, err := jwt.ParseRSAPrivateKeyFromPEM(secretByte)
	if err != nil {
		return nil, err
	}
	return &Jwt{
		signingString: rsaSigningString,
		secret:        secretByte,
	}, nil
}

func (t *Jwt) Generate(userId int64) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour)
	claims[userIdField] = userId

	return token.SignedString(t.signingString)
}

func (t *Jwt) Verify(token string) (bool, int64, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrWrongSigningMethod
		}

		tokenMap, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, ErrWrongFormat
		}

		tokenExpTime, err := tokenMap.GetExpirationTime()
		if err != nil {
			return nil, ErrExpiredToken
		}
		if tokenExpTime.After(time.Now()) {
			return nil, ErrWrongFormat
		}

		if _, ok = tokenMap[userIdField]; !ok {
			return nil, ErrWrongFormat
		}

		return t.secret, nil
	})

	if err != nil {
		return false, 0, err
	}

	if !tkn.Valid {
		return false, 0, err
	}

	userId, ok := tkn.Claims.(jwt.MapClaims)[userIdField]
	if !ok {
		return false, 0, err
	}

	return true, userId.(int64), nil
}
