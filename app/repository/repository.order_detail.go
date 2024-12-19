package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type CustomerOrderDetailRepository interface {
	Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail
	Update(ctx context.Context, order entity.CustomerOrderDetail, selectField interface{}, where entity.CustomerOrderDetail) entity.CustomerOrderDetail
	Delete(ctx context.Context, where entity.CustomerOrderDetail)
	FindById(ctx context.Context, orderId string) (entity.CustomerOrderDetail, error)
	FindAll(ctx context.Context) []entity.CustomerOrderDetail
	FindSpesificData(ctx context.Context, where entity.ViewOrderDetail) []entity.ViewOrderDetail
}
