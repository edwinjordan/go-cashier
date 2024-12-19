package otp_repository

import (
	"context"
	"errors"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserOTPImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.UserOTPRepository {
	return &UserOTPImpl{
		DB: db,
	}
}

func (repo *UserOTPImpl) Create(ctx context.Context, otp entity.UserOTP) entity.UserOTP {
	otpData := &UserOTP{}
	otpData = otpData.FromEntity(&otp)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&otpData).Error
	helpers.PanicIfError(err)

	return *otpData.ToEntity()
}

func (repo *UserOTPImpl) Update(ctx context.Context, otp entity.UserOTP, otpId string) entity.UserOTP {
	otpData := &UserOTP{}
	otpData = otpData.FromEntity(&otp)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("otp_id = ?", otpId).Save(&otpData).Error
	helpers.PanicIfError(err)
	return *otpData.ToEntity()
}

func (repo *UserOTPImpl) Delete(ctx context.Context, otpId string) {
	otp := &UserOTP{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("otp_id = ?", otpId).Delete(&otp).Error
	helpers.PanicIfError(err)
}

func (repo *UserOTPImpl) FindById(ctx context.Context, otpId string) (entity.UserOTP, error) {
	otpData := &UserOTP{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("otp_id = ?", otpId).
		First(&otpData).Error
	if err != nil {
		return *otpData.ToEntity(), errors.New("data otp tidak ditemukan")
	}
	return *otpData.ToEntity(), nil
}

func (repo *UserOTPImpl) FindAll(ctx context.Context) []entity.UserOTP {
	otp := []UserOTP{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Find(&otp).Error
	helpers.PanicIfError(err)

	var tempData []entity.UserOTP
	for _, v := range otp {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *UserOTPImpl) FindSpesificData(ctx context.Context, where entity.UserOTP) []entity.UserOTP {
	otp := []UserOTP{}
	otpWhere := &UserOTP{}
	otpWhere = otpWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Order("otp_expired DESC").Where(otpWhere).Find(&otp).Error
	helpers.PanicIfError(err)

	var tempData []entity.UserOTP
	for _, v := range otp {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
