package repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type BoothStockRepository interface {
	Create(ctx context.Context, stockBooth entity.StockBooth) entity.StockBooth
}
