package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type CustomerAddressRepository interface {
	Create(ctx context.Context, dataEn entity.CustomerAddress) entity.CustomerAddress
	Update(ctx context.Context, dataEn entity.CustomerAddress, selectField interface{}, where entity.CustomerAddress) entity.CustomerAddress
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) (entity.CustomerAddress, error)
	FindAll(ctx context.Context) []entity.CustomerAddress
	FindSpesificData(ctx context.Context, where entity.CustomerAddress) []entity.CustomerAddress
}
