package customer_address_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerAddress struct {
	AddressId         string    `gorm:"column:address_id"`
	AddressCustomerId string    `gorm:"column:address_customer_id"`
	AddressText       string    `gorm:"column:address_text"`
	AddressName       string    `gorm:"column:address_name"`
	AddressProvinceId string    `gorm:"column:address_province_id"`
	AddressProvince   string    `gorm:"column:address_province"`
	AddressCityId     string    `gorm:"column:address_city_id"`
	AddressCity       string    `gorm:"column:address_city"`
	AddressDistrictId string    `gorm:"column:address_district_id"`
	AddressDistrict   string    `gorm:"column:address_district"`
	AddressVillageId  string    `gorm:"column:address_village_id"`
	AddressVillage    string    `gorm:"column:address_village"`
	AddressPostalCode string    `gorm:"column:address_postal_code"`
	AddressMain       int       `gorm:"column:address_main"`
	AddressCreateAt   time.Time `gorm:"column:address_create_at"`
	AddressUpdateAt   time.Time `gorm:"column:address_update_at"`
}

func (CustomerAddress) TableName() string {
	return "tb_customer_address"
}

func (model *CustomerAddress) BeforeCreate(tx *gorm.DB) (err error) {
	model.AddressId = helpers.GenUUID()
	model.AddressCreateAt = helpers.CreateDateTime()
	model.AddressUpdateAt = helpers.CreateDateTime()
	return
}

func (CustomerAddress) FromEntity(e *entity.CustomerAddress) *CustomerAddress {
	return &CustomerAddress{
		AddressId:         e.AddressId,
		AddressCustomerId: e.AddressCustomerId,
		AddressText:       e.AddressText,
		AddressName:       e.AddressName,
		AddressProvince:   e.AddressProvince,
		AddressCity:       e.AddressCity,
		AddressDistrict:   e.AddressDistrict,
		AddressVillage:    e.AddressVillage,
		AddressPostalCode: e.AddressPostalCode,
		AddressMain:       e.AddressMain,
		AddressProvinceId: e.AddressProvinceId,
		AddressCityId:     e.AddressCityId,
		AddressDistrictId: e.AddressDistrictId,
		AddressVillageId:  e.AddressVillageId,
		AddressCreateAt:   e.AddressCreateAt,
		AddressUpdateAt:   e.AddressUpdateAt,
	}
}

func (model *CustomerAddress) ToEntity() *entity.CustomerAddress {
	modelData := &entity.CustomerAddress{
		AddressId:         model.AddressId,
		AddressCustomerId: model.AddressCustomerId,
		AddressText:       model.AddressText,
		AddressName:       model.AddressName,
		AddressProvince:   model.AddressProvince,
		AddressCity:       model.AddressCity,
		AddressDistrict:   model.AddressDistrict,
		AddressVillage:    model.AddressVillage,
		AddressPostalCode: model.AddressPostalCode,
		AddressMain:       model.AddressMain,
		AddressProvinceId: model.AddressProvinceId,
		AddressCityId:     model.AddressCityId,
		AddressDistrictId: model.AddressDistrictId,
		AddressVillageId:  model.AddressVillageId,
		AddressCreateAt:   model.AddressCreateAt,
		AddressUpdateAt:   model.AddressUpdateAt,
	}

	return modelData
}
