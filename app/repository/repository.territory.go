package repository

import (
	"context"
)

type TerritoryRepository interface {
	FindSpesificDataProvince(ctx context.Context, where map[string]interface{}) []map[string]interface{}
	FindSpesificDataCity(ctx context.Context, where map[string]interface{}) []map[string]interface{}
	FindSpesificDataSubdistrict(ctx context.Context, where map[string]interface{}) []map[string]interface{}
	FindSpesificDataVillage(ctx context.Context, where map[string]interface{}) []map[string]interface{}
}
