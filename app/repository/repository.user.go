package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type UserRepository interface {
	FindById(ctx context.Context, user entity.User, userId string) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
	FindSpesificData(ctx context.Context, where entity.User) []entity.User
	CheckMaintenanceMode(ctx context.Context, where map[string]interface{}) map[string]interface{}
}
