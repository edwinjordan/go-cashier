package category_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (entity.Category, error) {
	categoryData := &Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("category_id = ?", categoryId).
		First(&categoryData).Error
	if err != nil {
		return *categoryData.ToEntity(), errors.New("data kategori tidak ditemukan")
	}
	return *categoryData.ToEntity(), nil
}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context) []entity.Category {
	category := []Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_delete_at IS NULL").Find(&category).Error
	helpers.PanicIfError(err)

	var tempData []entity.Category
	for _, v := range category {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}
