package user_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserLog struct {
	LogUserId         string    `gorm:"column:log_user_id"`
	LogUserUserId     string    `gorm:"column:log_user_user_id"`
	LogUserToken      string    `gorm:"column:log_user_token"`
	LogUserMetadata   string    `gorm:"column:log_user_metadata"`
	LogUserLoginDate  time.Time `gorm:"column:log_user_login_date"`
	LogUserLogoutDate time.Time `gorm:"column:log_user_logout_date"`
}

func (UserLog) TableName() string {
	return "tb_user_token"
}

func (model *UserLog) BeforeCreate(tx *gorm.DB) (err error) {
	model.LogUserId = helpers.GenUUID()
	model.LogUserLoginDate = helpers.CreateDateTime()
	return
}

func (UserLog) FromEntity(e *entity.UserLog) *UserLog {
	return &UserLog{
		LogUserId:         e.LogUserId,
		LogUserUserId:     e.LogUserUserId,
		LogUserToken:      e.LogUserToken,
		LogUserMetadata:   e.LogUserMetadata,
		LogUserLoginDate:  e.LogUserLoginDate,
		LogUserLogoutDate: e.LogUserLogoutDate,
	}
}

func (model *UserLog) ToEntity() *entity.UserLog {
	modelData := &entity.UserLog{
		LogUserId:         model.LogUserId,
		LogUserUserId:     model.LogUserUserId,
		LogUserToken:      model.LogUserToken,
		LogUserMetadata:   model.LogUserMetadata,
		LogUserLoginDate:  model.LogUserLoginDate,
		LogUserLogoutDate: model.LogUserLogoutDate,
	}

	return modelData
}
