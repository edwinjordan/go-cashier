package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type MajorRepository interface {
	Create(ctx context.Context, major entity.Major) entity.Major
	Update(ctx context.Context, major entity.Major, majorId string) entity.Major
	Delete(ctx context.Context, majorId string)
	FindById(ctx context.Context, majorId string) (entity.Major, error)
	FindAll(ctx context.Context) []entity.Major
	FindSpesificData(ctx context.Context, where entity.Major) []entity.Major
}
