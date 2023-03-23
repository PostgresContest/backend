package auth

import (
	"context"

	pctx "backend/internal/context"
	"backend/internal/errors"

	oapi "github.com/PostgresContest/openapi/gen/v1"
)

type Security struct {
	jwt *Jwt
}

func NewSecurityProvider(jwt *Jwt) *Security {
	return &Security{jwt: jwt}
}

func (s *Security) HandleBearerAuth(ctx context.Context, _ string, t oapi.BearerAuth) (context.Context, error) {
	token := t.GetToken()

	verified, userID, err := s.jwt.Verify(token)
	if err != nil {
		return ctx, err
	}

	if !verified || userID < 1 {
		return ctx, errors.UnauthorizedHTTPError
	}

	return pctx.WithUserID(ctx, userID), nil
}
