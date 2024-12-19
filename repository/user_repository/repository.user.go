package user_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.UserRepository {

	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repo *UserRepositoryImpl) FindById(ctx context.Context, user entity.User, userId string) (entity.User, error) {
	userData := &User{}
	userData = userData.FromEntity(&user)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Joins("join pegawai on pegawai_id=user_pegawai_id").Preload("Pegawai").Where("user_id = ?", userId).First(&userData).Error
	if err != nil {
		return *userData.ToEntity(), errors.New("data tidak ditemukan")
	}

	return *userData.ToEntity(), nil
}

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	user := []User{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&user).Error
	helpers.PanicIfError(err)

	var tempData []entity.User
	for _, v := range user {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *UserRepositoryImpl) FindSpesificData(ctx context.Context, where entity.User) []entity.User {
	user := []User{}
	whereUser := &User{}
	whereUser = whereUser.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Preload("Pegawai").Where(whereUser).Find(&user).Error
	helpers.PanicIfError(err)

	var tempData []entity.User
	for _, v := range user {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}

func (repo *UserRepositoryImpl) CheckMaintenanceMode(ctx context.Context, where map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Table("ms_config").Where(where).Find(&data).Error
	helpers.PanicIfError(err)

	return data
}
