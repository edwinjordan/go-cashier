package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type TempCartRepository interface {
	Create(ctx context.Context, tempCart entity.TempCart)
	Update(ctx context.Context, tempCart entity.TempCart, tempCartId string)
	Delete(ctx context.Context, tempCartId string)
	DeleteSpesificData(ctx context.Context, data entity.TempCart)
	FindSpesificData(ctx context.Context, where entity.TempCart) []entity.TempCart
}
