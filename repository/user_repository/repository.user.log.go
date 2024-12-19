package user_repository

import (
	"context"

	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserLogRepositoryImpl struct {
	DB *gorm.DB
}

func NewLog(db *gorm.DB) repository.UserLogRepository {
	return &UserLogRepositoryImpl{
		DB: db,
	}
}

func (repo *UserLogRepositoryImpl) Create(ctx context.Context, userLog entity.UserLog) {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Exec("INSERT INTO tb_user_token(log_user_id,log_user_user_id,log_user_token,log_user_metadata,log_user_login_date) VALUES(?,?,?,?,?)", helpers.GenUUID(), userLog.LogUserUserId, userLog.LogUserToken, userLog.LogUserMetadata, helpers.CreateDateTime()).Error
	helpers.PanicIfError(err)

}

func (repo *UserLogRepositoryImpl) Update(ctx context.Context, userLog entity.UserLog, userLogId string) entity.UserLog {
	userLogData := &UserLog{}
	userLogData = userLogData.FromEntity(&userLog)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("log_user_id = ?", userLogId).Save(&userLogData).Error
	helpers.PanicIfError(err)
	return *userLogData.ToEntity()
}

func (repo *UserLogRepositoryImpl) FindSpesificData(ctx context.Context, where entity.UserLog) []entity.UserLog {
	data := []UserLog{}
	whereData := &UserLog{}
	whereData = whereData.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Find(&data).Error
	helpers.PanicIfError(err)

	var tempData []entity.UserLog
	for _, v := range data {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
