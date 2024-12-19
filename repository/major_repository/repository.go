package major_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type MajorRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.MajorRepository {
	return &MajorRepositoryImpl{
		DB: db,
	}
}

func (repo *MajorRepositoryImpl) Create(ctx context.Context, major entity.Major) entity.Major {
	majorData := &Major{}
	majorData = majorData.FromEntity(&major)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&majorData).Error
	helpers.PanicIfError(err)

	return *majorData.ToEntity()
}

func (repo *MajorRepositoryImpl) Update(ctx context.Context, major entity.Major, majorId string) entity.Major {
	majorData := &Major{}
	majorData = majorData.FromEntity(&major)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_id = ?", majorId).Save(&majorData).Error
	helpers.PanicIfError(err)
	return *majorData.ToEntity()
}

func (repo *MajorRepositoryImpl) Delete(ctx context.Context, majorId string) {
	major := &Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_id = ?", majorId).Delete(&major).Error
	helpers.PanicIfError(err)
}

func (repo *MajorRepositoryImpl) FindById(ctx context.Context, majorId string) (entity.Major, error) {
	majorData := &Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("major_id = ?", majorId).
		First(&majorData).Error
	if err != nil {
		return *majorData.ToEntity(), errors.New("data jurusan tidak ditemukan")
	}
	return *majorData.ToEntity(), nil
}

func (repo *MajorRepositoryImpl) FindAll(ctx context.Context) []entity.Major {
	major := []Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_delete_at IS NULL").Find(&major).Error
	helpers.PanicIfError(err)

	var tempData []entity.Major
	for _, v := range major {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *MajorRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Major) []entity.Major {
	major := []Major{}
	majorWhere := &Major{}
	majorWhere = majorWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_delete_at IS NULL").Where(majorWhere).Find(&major).Error
	helpers.PanicIfError(err)

	var tempData []entity.Major
	for _, v := range major {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
