package transaction_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type TransactionDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransDetail(db *gorm.DB) repository.TransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{
		DB: db,
	}
}

func (repo *TransactionDetailRepositoryImpl) Create(ctx context.Context, transaksi entity.TransactionDetail) entity.TransactionDetail {
	transData := &TransactionDetail{}
	transData = transData.FromEntity(&transaksi)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&transData).Error
	helpers.PanicIfError(err)

	return *transData.ToEntity()
}

func (repo *TransactionDetailRepositoryImpl) Update(ctx context.Context, transaksi entity.TransactionDetail, transaksiId string) entity.TransactionDetail {
	transData := &TransactionDetail{}
	transData = transData.FromEntity(&transaksi)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("trans_detail_id = ?", transaksiId).Save(&transData).Error
	helpers.PanicIfError(err)
	return *transData.ToEntity()
}

func (repo *TransactionDetailRepositoryImpl) Delete(ctx context.Context, transaksiId string) {
	transaksi := &TransactionDetail{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("trans_detail_id = ?", transaksiId).Delete(&transaksi).Error
	helpers.PanicIfError(err)
}

func (repo *TransactionDetailRepositoryImpl) FindById(ctx context.Context, transaksiId string) (entity.TransactionDetail, error) {
	transData := &TransactionDetail{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("trans_detail_id = ?", transaksiId).
		First(&transData).Error
	if err != nil {
		return *transData.ToEntity(), errors.New("data transaksi tidak ditemukan")
	}
	return *transData.ToEntity(), nil
}

func (repo *TransactionDetailRepositoryImpl) FindAll(ctx context.Context) []entity.TransactionDetail {
	transaksi := []TransactionDetail{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&transaksi).Error
	helpers.PanicIfError(err)

	var tempData []entity.TransactionDetail
	for _, v := range transaksi {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *TransactionDetailRepositoryImpl) FindSpesificData(ctx context.Context, where entity.ViewTransactionDetail) []entity.ViewTransactionDetail {
	transaksi := []ViewTransactionDetail{}
	transaksiWhere := &ViewTransactionDetail{}
	transaksiWhere = transaksiWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(transaksiWhere).Find(&transaksi).Error
	helpers.PanicIfError(err)

	var tempData []entity.ViewTransactionDetail
	for _, v := range transaksi {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
