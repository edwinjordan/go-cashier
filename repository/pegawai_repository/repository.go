package pegawai_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type PegawaiRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.PegawaiRepository {
	return &PegawaiRepositoryImpl{
		DB: db,
	}
}

func (repo *PegawaiRepositoryImpl) Create(ctx context.Context, pegawai entity.Pegawai) entity.Pegawai {
	pegawaiData := &Pegawai{}
	pegawaiData = pegawaiData.FromEntity(&pegawai)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&pegawaiData).Error
	helpers.PanicIfError(err)

	return *pegawaiData.ToEntity()
}

func (repo *PegawaiRepositoryImpl) Update(ctx context.Context, pegawai entity.Pegawai, pegawaiId string) entity.Pegawai {
	pegawaiData := &Pegawai{}
	pegawaiData = pegawaiData.FromEntity(&pegawai)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("pegawai_id = ?", pegawaiId).Save(&pegawaiData).Error
	helpers.PanicIfError(err)
	return *pegawaiData.ToEntity()
}

func (repo *PegawaiRepositoryImpl) Delete(ctx context.Context, pegawaiId string) {
	pegawai := &Pegawai{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("pegawai_id = ?", pegawaiId).Delete(&pegawai).Error
	helpers.PanicIfError(err)
}

func (repo *PegawaiRepositoryImpl) FindById(ctx context.Context, pegawaiId string) (entity.Pegawai, error) {
	pegawaiData := &Pegawai{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("pegawai_id = ?", pegawaiId).
		First(&pegawaiData).Error
	if err != nil {
		return *pegawaiData.ToEntity(), errors.New("data pegawai tidak ditemukan")
	}
	return *pegawaiData.ToEntity(), nil
}

func (repo *PegawaiRepositoryImpl) FindAll(ctx context.Context) []entity.Pegawai {
	pegawai := []Pegawai{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&pegawai).Error
	helpers.PanicIfError(err)

	var tempData []entity.Pegawai
	for _, v := range pegawai {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *PegawaiRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Pegawai) []entity.Pegawai {
	pegawai := []Pegawai{}
	pegawaiWhere := &Pegawai{}
	pegawaiWhere = pegawaiWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(pegawaiWhere).Find(&pegawai).Error
	helpers.PanicIfError(err)

	var tempData []entity.Pegawai
	for _, v := range pegawai {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
