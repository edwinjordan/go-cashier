package order_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderDetail(db *gorm.DB) repository.CustomerOrderDetailRepository {
	return &CustomerOrderDetailRepositoryImpl{
		DB: db,
	}
}

func (repo *CustomerOrderDetailRepositoryImpl) Create(ctx context.Context, order entity.CustomerOrderDetail) entity.CustomerOrderDetail {
	orderData := &CustomerOrderDetail{}
	orderData = orderData.FromEntity(&order)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&orderData).Error
	helpers.PanicIfError(err)

	return *orderData.ToEntity()
}

func (repo *CustomerOrderDetailRepositoryImpl) Update(ctx context.Context, order entity.CustomerOrderDetail, selectField interface{}, where entity.CustomerOrderDetail) entity.CustomerOrderDetail {
	orderData := &CustomerOrderDetail{}
	orderData = orderData.FromEntity(&order)

	whereData := &CustomerOrderDetail{}
	whereData = whereData.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Select(selectField).Save(&orderData).Error
	helpers.PanicIfError(err)
	return *orderData.ToEntity()
}

func (repo *CustomerOrderDetailRepositoryImpl) Delete(ctx context.Context, where entity.CustomerOrderDetail) {
	order := &CustomerOrderDetail{}
	whereData := order.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Delete(&order).Error
	helpers.PanicIfError(err)
}

func (repo *CustomerOrderDetailRepositoryImpl) FindById(ctx context.Context, orderId string) (entity.CustomerOrderDetail, error) {
	orderData := &CustomerOrderDetail{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("order_id = ?", orderId).
		First(&orderData).Error
	if err != nil {
		return *orderData.ToEntity(), errors.New("data order detail tidak ditemukan")
	}
	return *orderData.ToEntity(), nil
}

func (repo *CustomerOrderDetailRepositoryImpl) FindAll(ctx context.Context) []entity.CustomerOrderDetail {
	order := []CustomerOrderDetail{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrderDetail
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *CustomerOrderDetailRepositoryImpl) FindSpesificData(ctx context.Context, where entity.ViewOrderDetail) []entity.ViewOrderDetail {
	order := []ViewOrderDetail{}
	orderWhere := &ViewOrderDetail{}
	orderWhere = orderWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(orderWhere).Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.ViewOrderDetail
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
