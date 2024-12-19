package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type VarianRepository interface {
	FindById(ctx context.Context, varianId string) (entity.Varian, error)
	FindSpesificData(ctx context.Context, where entity.Varian) []entity.Varian
	UpdateStock(ctx context.Context, id string, stok int)
}
