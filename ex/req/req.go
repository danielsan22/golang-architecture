package req

import "context"

type ctxKey string

var userNameKey ctxKey = "usr"
var userIdKey ctxKey = "id"

func WithUserId(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIdKey, id)
}

func WithUserName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, userNameKey, name)
}

func UserId(ctx context.Context) *int {
	if v := ctx.Value(userIdKey); v != nil {
		if i, ok := v.(int); ok {
			return &i
		}
	}
	return nil
}

func UserName(ctx context.Context) *string {
	if v := ctx.Value(userNameKey); v != nil {
		if i, ok := v.(string); ok {
			return &i
		}
	}
	return nil
}
