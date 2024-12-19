package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type ProductRepository interface {
	FindById(ctx context.Context, productId string) (entity.Product, error)
	FindAll(ctx context.Context, where entity.Product, config map[string]interface{}) []entity.Product
	FindSpesificData(ctx context.Context, where entity.Product) []entity.Product
}
