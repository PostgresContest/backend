package context

import "context"

type keyUserID struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, keyUserID{}, userID)
}

func UserID(ctx context.Context) int64 {
	val := ctx.Value(keyUserID{})
	if val != nil {
		return val.(int64)
	}
	return 0
}
