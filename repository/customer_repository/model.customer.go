package customer_repository

import (
	"database/sql"
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_address_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/major_repository"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerId             string                                         `gorm:"column:customer_id"`
	CustomerCode           string                                         `gorm:"column:customer_code"`
	CustomerName           string                                         `gorm:"column:customer_name"`
	CustomerGender         string                                         `gorm:"column:customer_gender"`
	CustomerPhonenumber    string                                         `gorm:"column:customer_phonenumber"`
	CustomerEmail          string                                         `gorm:"column:customer_email"`
	CustomerDob            sql.NullString                                 `gorm:"column:customer_dob"`
	CustomerPassword       string                                         `gorm:"column:customer_password"`
	CustomerProfilePic     string                                         `gorm:"column:customer_profile_pic"`
	CustomerClass          string                                         `gorm:"column:customer_class"`
	CustomerMajorId        string                                         `gorm:"column:customer_major_id"`
	CustomerProfilePicPath string                                         `gorm:"column:customer_profile_pic_path"`
	CustomerStatus         int                                            `gorm:"column:customer_status"`
	CustomerLastStatus     int                                            `gorm:"column:customer_last_status"`
	CustomerCreateAt       time.Time                                      `gorm:"column:customer_create_at"`
	CustomerUpdateAt       time.Time                                      `gorm:"column:customer_update_at"`
	Major                  *major_repository.Major                        `gorm:"foreignKey:CustomerMajorId;references:MajorId"`
	Address                *[]customer_address_repository.CustomerAddress `gorm:"foreignKey:AddressCustomerId;references:CustomerId"`
}

func (Customer) TableName() string {
	return "tb_customer"
}

func (model *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	model.CustomerId = helpers.GenUUID()
	model.CustomerCreateAt = helpers.CreateDateTime()
	model.CustomerUpdateAt = helpers.CreateDateTime()
	model.CustomerLastStatus = 0
	model.CustomerStatus = 0
	return
}

func (Customer) FromEntity(e *entity.Customer) *Customer {
	return &Customer{
		CustomerId:             e.CustomerId,
		CustomerCode:           e.CustomerCode,
		CustomerName:           e.CustomerName,
		CustomerGender:         e.CustomerGender,
		CustomerPhonenumber:    e.CustomerPhonenumber,
		CustomerEmail:          e.CustomerEmail,
		CustomerDob:            e.CustomerDob,
		CustomerPassword:       e.CustomerPassword,
		CustomerProfilePic:     e.CustomerProfilePic,
		CustomerClass:          e.CustomerClass,
		CustomerMajorId:        e.CustomerMajorId,
		CustomerProfilePicPath: e.CustomerProfilePicPath,
		CustomerStatus:         e.CustomerStatus,
		CustomerLastStatus:     e.CustomerLastStatus,
		CustomerCreateAt:       e.CustomerCreateAt,
		CustomerUpdateAt:       e.CustomerUpdateAt,
	}
}

func (model *Customer) ToEntity() *entity.CustomerResponse {
	modelData := &entity.CustomerResponse{
		CustomerId:             model.CustomerId,
		CustomerCode:           model.CustomerCode,
		CustomerName:           model.CustomerName,
		CustomerGender:         model.CustomerGender,
		CustomerPhonenumber:    model.CustomerPhonenumber,
		CustomerEmail:          model.CustomerEmail,
		CustomerDob:            model.CustomerDob,
		CustomerPassword:       model.CustomerPassword,
		CustomerProfilePic:     model.CustomerProfilePic,
		CustomerClass:          model.CustomerClass,
		CustomerMajorId:        model.CustomerMajorId,
		CustomerProfilePicPath: model.CustomerProfilePicPath,
		CustomerStatus:         model.CustomerStatus,
		CustomerLastStatus:     model.CustomerLastStatus,
		CustomerCreateAt:       model.CustomerCreateAt,
		CustomerUpdateAt:       model.CustomerUpdateAt,
	}
	if model.Major != nil {
		modelData.Major = model.Major.ToEntity()
	}

	if model.Address != nil {
		var tempMenu []entity.CustomerAddress
		for _, v := range *model.Address {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.Address = &tempMenu
	}
	return modelData
}
