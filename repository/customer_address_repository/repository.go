package customer_address_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerAddressRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.CustomerAddressRepository {
	return &CustomerAddressRepositoryImpl{
		DB: db,
	}
}

func (repo *CustomerAddressRepositoryImpl) Create(ctx context.Context, dataEn entity.CustomerAddress) entity.CustomerAddress {
	data := &CustomerAddress{}
	data = data.FromEntity(&dataEn)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&data).Error
	helpers.PanicIfError(err)

	return *data.ToEntity()
}

func (repo *CustomerAddressRepositoryImpl) Update(ctx context.Context, dataEn entity.CustomerAddress, selectField interface{}, where entity.CustomerAddress) entity.CustomerAddress {
	data := &CustomerAddress{}
	data = data.FromEntity(&dataEn)

	dataWhere := &CustomerAddress{}
	dataWhere = data.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Select(selectField).Where(dataWhere).Save(&data).Error
	helpers.PanicIfError(err)
	return *data.ToEntity()
}

func (repo *CustomerAddressRepositoryImpl) Delete(ctx context.Context, id string) {
	data := &CustomerAddress{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("address_id = ?", id).Delete(&data).Error
	helpers.PanicIfError(err)
}

func (repo *CustomerAddressRepositoryImpl) FindById(ctx context.Context, id string) (entity.CustomerAddress, error) {
	data := &CustomerAddress{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("address_id = ?", id).
		First(&data).Error
	if err != nil {
		return *data.ToEntity(), errors.New("data alamat tidak ditemukan")
	}
	return *data.ToEntity(), nil
}

func (repo *CustomerAddressRepositoryImpl) FindAll(ctx context.Context) []entity.CustomerAddress {
	data := []CustomerAddress{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&data).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerAddress
	for _, v := range data {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *CustomerAddressRepositoryImpl) FindSpesificData(ctx context.Context, where entity.CustomerAddress) []entity.CustomerAddress {
	data := []CustomerAddress{}
	dataWhere := &CustomerAddress{}
	dataWhere = dataWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Order("address_main DESC").Where(dataWhere).Find(&data).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerAddress
	for _, v := range data {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
