package auth

import (
	"errors"
	"time"

	"backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrWrongSigningMethod = errors.New("wrong signing method")
	ErrWrongFormat        = errors.New("wrong token format")
	ErrExpiredToken       = errors.New("token expired")
	ErrTokenNotValid      = errors.New("token not valid")
)

const (
	expField    = "exp"
	userIDField = "user_id"
)

type Jwt struct {
	signingString any
	ttlSeconds    int32
}

func NewJwtProvider(cfg *config.Config) (*Jwt, error) {
	secretByte := []byte(cfg.Jwt.Secret)

	return &Jwt{
		signingString: secretByte,
		ttlSeconds:    cfg.Jwt.TTLSeconds,
	}, nil
}

func (t *Jwt) Generate(userID int64) (string, error) {
	payload := jwt.MapClaims{
		expField:    time.Now().Add(time.Duration(t.ttlSeconds) * time.Second).Unix(),
		userIDField: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(t.signingString)
}

func checkTokenExpInfo(tokenMap jwt.MapClaims) error {
	tokenExpTime, ok := tokenMap[expField]

	tokenExpTimeUnix, formatOk := tokenExpTime.(float64)
	if !ok || !formatOk {
		return ErrWrongFormat
	}

	if time.Unix(int64(tokenExpTimeUnix), 0).Before(time.Now()) {
		return ErrExpiredToken
	}

	return ErrExpiredToken
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

		if err := checkTokenExpInfo(tokenMap); err != nil {
			return nil, err
		}

		if _, ok = tokenMap[userIDField]; !ok {
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

	userID, ok := tkn.Claims.(jwt.MapClaims)[userIDField]

	userIDFloat, formatOk := userID.(float64)
	if !ok || !formatOk {
		return false, 0, ErrTokenNotValid
	}

	return true, int64(userIDFloat), nil
}
