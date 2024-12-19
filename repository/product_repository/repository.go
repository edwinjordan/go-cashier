package product_repository

import (
	"context"
	"errors"
	"html"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.ProductRepository {

	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repo *ProductRepositoryImpl) FindById(ctx context.Context, product_id string) (entity.Product, error) {
	productData := &Product{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("product_id = ?", product_id).First(&productData).Error
	if err != nil {
		return *productData.ToEntity(), errors.New("data tidak ditemukan")
	}

	return *productData.ToEntity(), nil
}

func (repo *ProductRepositoryImpl) FindAll(ctx context.Context, where entity.Product, config map[string]interface{}) []entity.Product {
	/* search */
	whereLike := ""
	if config["search"].(string) != "" {
		whereLike = "(product_name LIKE '%" + html.EscapeString(config["search"].(string)) + "%')"
	}

	product := []Product{}
	whereProduct := &Product{}
	whereProduct = whereProduct.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Preload("Varian").
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Where("product_delete_at IS NULL").
		Where(whereProduct).
		Where(whereLike).
		Find(&product).Error
	helpers.PanicIfError(err)

	var tempData []entity.Product
	for _, v := range product {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *ProductRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Product) []entity.Product {
	product := []Product{}
	whereProduct := &Product{}
	whereProduct = whereProduct.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereProduct).Find(&product).Error
	helpers.PanicIfError(err)

	var tempData []entity.Product
	for _, v := range product {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
