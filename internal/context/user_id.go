package context

import "context"

type keyUserID struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, keyUserID{}, userID)
}

func UserID(ctx context.Context) int64 {
	val := ctx.Value(keyUserID{})
	if value, ok := val.(int64); val != nil && ok {
		return value
	}

	return 0
}
