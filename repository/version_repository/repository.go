package version_repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type VersionRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.VersionRepository {
	return &VersionRepositoryImpl{
		DB: db,
	}
}

func (repo *VersionRepositoryImpl) GetVersionAdmin(ctx context.Context) entity.VersionAdmin {
	versionData := &VersionAdmin{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Order("version_code DESC").
		First(&versionData).Error
	helpers.PanicIfError(err)
	return *versionData.ToEntity()
}

func (repo *VersionRepositoryImpl) GetVersionShop(ctx context.Context) entity.VersionShop {
	versionData := &VersionShop{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Order("version_code DESC").
		First(&versionData).Error
	helpers.PanicIfError(err)
	return *versionData.ToEntity()
}
