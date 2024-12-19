package entity

import "time"

type CustomerAddress struct {
	AddressId         string    `json:"address_id"`
	AddressCustomerId string    `json:"address_customer_id"`
	AddressText       string    `json:"address_text"`
	AddressName       string    `json:"address_name"`
	AddressProvinceId string    `json:"address_province_id"`
	AddressProvince   string    `json:"address_province"`
	AddressCityId     string    `json:"address_city_id"`
	AddressCity       string    `json:"address_city"`
	AddressDistrictId string    `json:"address_district_id"`
	AddressDistrict   string    `json:"address_district"`
	AddressVillageId  string    `json:"address_village_id"`
	AddressVillage    string    `json:"address_village"`
	AddressPostalCode string    `json:"address_postal_code"`
	AddressMain       int       `json:"address_main"`
	AddressCreateAt   time.Time `json:"address_create_at"`
	AddressUpdateAt   time.Time `json:"address_update_at"`
}
