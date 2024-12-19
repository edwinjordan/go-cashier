package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer entity.Customer) entity.CustomerResponse
	Update(ctx context.Context, selectField interface{}, customer entity.Customer, customerId string) entity.CustomerResponse
	Delete(ctx context.Context, customerId string)
	FindById(ctx context.Context, customerId string) (entity.CustomerResponse, error)
	FindAll(ctx context.Context, conf map[string]interface{}) []entity.CustomerResponse
	FindSpesificData(ctx context.Context, where entity.Customer) []entity.CustomerResponse
	GenCustCode(ctx context.Context) string
}
