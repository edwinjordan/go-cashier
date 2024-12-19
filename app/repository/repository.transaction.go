package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction entity.Transaction) entity.Transaction
	Update(ctx context.Context, transaction entity.Transaction, transactionId string) entity.Transaction
	Delete(ctx context.Context, transactionId string)
	FindById(ctx context.Context, transactionId string) (entity.Transaction, error)
	FindAll(ctx context.Context) []entity.Transaction
	FindSpesificData(ctx context.Context, where entity.Transaction, conf map[string]interface{}) []entity.Transaction
	GenInvoice(ctx context.Context) string
	GetTransactionSummary(ctx context.Context, where map[string]interface{}) map[string]interface{}
}
