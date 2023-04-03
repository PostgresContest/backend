package auth

import (
	"backend/internal/infrastructure/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrWrongSigningMethod = errors.New("wrong signing method")
	ErrWrongFormat        = errors.New("wrong token format")
	ErrExpiredToken       = errors.New("token expired")
)

type AccessTokenProvider struct {
	signingString any
	ttlSeconds    int32
}

type Claims struct {
	Exp      time.Time
	UserID   int64
	UserRole string
	jwt.RegisteredClaims
}

func (c *Claims) GetExpiration() time.Time {
	return c.Exp
}

func (c *Claims) GetUserID() int64 {
	return c.UserID
}

func (c *Claims) GetUserRole() string {
	return c.UserRole
}

func NewAccessTokenProvider(cfg *config.Config) (*AccessTokenProvider, error) {
	secretByte := []byte(cfg.Jwt.Secret)

	return &AccessTokenProvider{
		signingString: secretByte,
		ttlSeconds:    cfg.Jwt.TTLSeconds,
	}, nil
}

func (t *AccessTokenProvider) Generate(userID int64, userRole string) (string, *Claims, error) {
	payload := Claims{
		Exp:      time.Now().Add(time.Duration(t.ttlSeconds) * time.Second),
		UserID:   userID,
		UserRole: userRole,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedString, err := token.SignedString(t.signingString)
	if err != nil {
		return "", nil, err
	}

	return signedString, &payload, nil
}

func checkTokenCompleteness(cl *Claims) error {
	if !cl.Exp.Before(time.Now()) {
		return ErrExpiredToken
	}

	if cl.UserID == 0 {
		return ErrWrongFormat
	}

	if len(cl.UserRole) == 0 {
		return ErrExpiredToken
	}

	return nil
}

func (t *AccessTokenProvider) ParseClaims(token string) (*Claims, error) {
	cl := new(Claims)
	_, err := jwt.ParseWithClaims(token, cl, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongSigningMethod
		}

		parsedClaims, ok := token.Claims.(*Claims)
		if !ok {
			return nil, ErrWrongFormat
		}

		if err := checkTokenCompleteness(parsedClaims); err != nil {
			return nil, err
		}

		return t.signingString, nil
	})

	if err != nil {
		return nil, err
	}

	return cl, nil
}
