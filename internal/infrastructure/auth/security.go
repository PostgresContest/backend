package auth

import (
	"context"

	pctx "backend/internal/context"
	"backend/internal/errors"
	"backend/models"
	oapi "github.com/PostgresContest/openapi/gen/v1"
)

type Security struct {
	jwt *AccessTokenProvider
}

func NewSecurityProvider(jwt *AccessTokenProvider) *Security {
	return &Security{jwt: jwt}
}

func (s *Security) HandleBearerAuth(ctx context.Context, _ string, t oapi.BearerAuth) (context.Context, error) {
	token := t.GetToken()

	cl, err := s.jwt.ParseClaims(token)
	if err != nil {
		return ctx, err
	}

	ctx = pctx.WithUserID(ctx, cl.GetUserID())
	ctx = pctx.WithUserRole(ctx, cl.GetUserRole())

	return ctx, nil
}

func (s *Security) HandleBearerAdminAuth(
	ctx context.Context,
	_ string,
	t oapi.BearerAdminAuth,
) (context.Context, error) {
	token := t.GetToken()

	cl, err := s.jwt.ParseClaims(token)
	if err != nil {
		return ctx, err
	}

	if cl.GetUserRole() != models.UserRoleAdmin {
		return nil, errors.UnauthorizedHTTPError
	}

	ctx = pctx.WithUserID(ctx, cl.GetUserID())
	ctx = pctx.WithUserRole(ctx, cl.GetUserRole())

	return ctx, nil
}
