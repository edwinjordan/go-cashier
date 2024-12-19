package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type UserLogRepository interface {
	Create(ctx context.Context, userLog entity.UserLog)
	Update(ctx context.Context, userLog entity.UserLog, userLogId string) entity.UserLog
	FindSpesificData(ctx context.Context, where entity.UserLog) []entity.UserLog
}
