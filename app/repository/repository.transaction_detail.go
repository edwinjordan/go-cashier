package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type TransactionDetailRepository interface {
	Create(ctx context.Context, transactionDetail entity.TransactionDetail) entity.TransactionDetail
	Update(ctx context.Context, transactionDetail entity.TransactionDetail, transactionDetailId string) entity.TransactionDetail
	Delete(ctx context.Context, transactionDetailId string)
	FindById(ctx context.Context, transactionDetailId string) (entity.TransactionDetail, error)
	FindAll(ctx context.Context) []entity.TransactionDetail
	FindSpesificData(ctx context.Context, where entity.ViewTransactionDetail) []entity.ViewTransactionDetail
}
