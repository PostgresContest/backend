package context

import "context"

type keyUserRole struct{}

func WithUserRole(ctx context.Context, userRole string) context.Context {
	return context.WithValue(ctx, keyUserRole{}, userRole)
}

func UserRole(ctx context.Context) string {
	val := ctx.Value(keyUserRole{})
	if value, ok := val.(string); val != nil && ok {
		return value
	}

	return ""
}
