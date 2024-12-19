package territory_repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type TerritoryRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.TerritoryRepository {

	return &TerritoryRepositoryImpl{
		DB: db,
	}
}
func (repo *TerritoryRepositoryImpl) FindSpesificDataProvince(ctx context.Context, where map[string]interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Table("ms_provinsi").Where(where).Find(&data).Error
	helpers.PanicIfError(err)

	return data
}
func (repo *TerritoryRepositoryImpl) FindSpesificDataCity(ctx context.Context, where map[string]interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Table("ms_kota").Where(where).Find(&data).Error
	helpers.PanicIfError(err)

	return data
}

func (repo *TerritoryRepositoryImpl) FindSpesificDataSubdistrict(ctx context.Context, where map[string]interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Table("ms_kecamatan").Where(where).Find(&data).Error
	helpers.PanicIfError(err)

	return data
}

func (repo *TerritoryRepositoryImpl) FindSpesificDataVillage(ctx context.Context, where map[string]interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Table("ms_desa").Where(where).Find(&data).Error
	helpers.PanicIfError(err)

	return data
}
