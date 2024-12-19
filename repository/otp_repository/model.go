package otp_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type UserOTP struct {
	OtpId         string    `gorm:"column:otp_id"`
	OtpCustomerId string    `gorm:"column:otp_customer_id"`
	OtpNumber     string    `gorm:"column:otp_number"`
	OtpStatus     int       `gorm:"column:otp_status"`
	OtpExpired    time.Time `gorm:"column:otp_expired"`
}

func (UserOTP) TableName() string {
	return "tb_otp"
}

func (model *UserOTP) BeforeCreate(tx *gorm.DB) (err error) {
	model.OtpId = helpers.GenUUID()
	model.OtpExpired = helpers.CreateDateTime().Add(time.Minute * 15)
	model.OtpStatus = 0
	return
}

func (UserOTP) FromEntity(e *entity.UserOTP) *UserOTP {
	return &UserOTP{
		OtpId:         e.OtpId,
		OtpCustomerId: e.OtpCustomerId,
		OtpNumber:     e.OtpNumber,
		OtpStatus:     e.OtpStatus,
		OtpExpired:    e.OtpExpired,
	}
}

func (model *UserOTP) ToEntity() *entity.UserOTP {
	modelData := &entity.UserOTP{
		OtpId:         model.OtpId,
		OtpCustomerId: model.OtpCustomerId,
		OtpNumber:     model.OtpNumber,
		OtpStatus:     model.OtpStatus,
		OtpExpired:    model.OtpExpired,
	}
	return modelData
}
