package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type CategoryRepository interface {
	FindById(ctx context.Context, categoryId string) (entity.Category, error)
	FindAll(ctx context.Context) []entity.Category
}
