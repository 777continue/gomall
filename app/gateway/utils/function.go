package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	userId, ok := ctx.Value(SessionUserId).(uint32)
	if !ok {
		return 0
	}
	return userId
}
