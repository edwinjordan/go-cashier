package tempcart_repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type TempCartRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.TempCartRepository {
	return &TempCartRepositoryImpl{
		DB: db,
	}
}

func (repo *TempCartRepositoryImpl) Create(ctx context.Context, tempCart entity.TempCart) {
	data := &TempCart{}
	data = data.FromEntity(&tempCart)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&data).Error
	helpers.PanicIfError(err)

}

func (repo *TempCartRepositoryImpl) Update(ctx context.Context, tempCart entity.TempCart, tempCartId string) {
	data := &TempCart{}
	data = data.FromEntity(&tempCart)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("temp_cart_id = ?", tempCartId).Save(&data).Error
	helpers.PanicIfError(err)
}

func (repo *TempCartRepositoryImpl) Delete(ctx context.Context, tempCartId string) {
	tempCart := &TempCart{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("temp_cart_id = ?", tempCartId).Delete(&tempCart).Error
	helpers.PanicIfError(err)
}

func (repo *TempCartRepositoryImpl) DeleteSpesificData(ctx context.Context, data entity.TempCart) {
	tempCart := &TempCart{}
	dataWhere := &TempCart{}
	dataWhere = dataWhere.FromEntity(&data)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(dataWhere).Delete(&tempCart).Error
	helpers.PanicIfError(err)
}

func (repo *TempCartRepositoryImpl) FindSpesificData(ctx context.Context, where entity.TempCart) []entity.TempCart {
	tempCart := []TempCart{}
	tempCartWhere := &TempCart{}
	tempCartWhere = tempCartWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(tempCartWhere).Find(&tempCart).Error
	helpers.PanicIfError(err)

	var tempData []entity.TempCart
	for _, v := range tempCart {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
