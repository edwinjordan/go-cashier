package entity

import (
	"database/sql"
	"time"
)

type Customer struct {
	CustomerId             string             `json:"customer_id"`
	CustomerCode           string             `json:"customer_code"`
	CustomerName           string             `json:"customer_name"`
	CustomerGender         string             `json:"customer_gender"`
	CustomerPhonenumber    string             `json:"customer_phonenumber"`
	CustomerEmail          string             `json:"customer_email"`
	CustomerDob            sql.NullString     `json:"customer_dob"`
	CustomerPassword       string             `json:"customer_password"`
	CustomerProfilePic     string             `json:"customer_profile_pic"`
	CustomerClass          string             `json:"customer_class"`
	CustomerMajorId        string             `json:"customer_major_id"`
	CustomerProfilePicPath string             `json:"customer_profile_pic_path"`
	CustomerStatus         int                `json:"customer_status"`
	CustomerLastStatus     int                `json:"customer_last_status"`
	CustomerCreateAt       time.Time          `json:"customer_create_at"`
	CustomerUpdateAt       time.Time          `json:"customer_update_at"`
	Major                  *Major             `json:"jurusan"`
	Address                *[]CustomerAddress `json:"alamat"`
}

type CustomerResponse struct {
	CustomerId             string             `json:"customer_id"`
	CustomerCode           string             `json:"customer_code"`
	CustomerName           string             `json:"customer_name"`
	CustomerGender         string             `json:"customer_gender"`
	CustomerPhonenumber    string             `json:"customer_phonenumber"`
	CustomerEmail          string             `json:"customer_email"`
	CustomerDob            sql.NullString     `json:"customer_dob"`
	CustomerPassword       string             `json:"-"`
	CustomerProfilePic     string             `json:"customer_profile_pic"`
	CustomerClass          string             `json:"customer_class"`
	CustomerMajorId        string             `json:"customer_major_id"`
	CustomerProfilePicPath string             `json:"customer_profile_pic_path"`
	CustomerStatus         int                `json:"customer_status"`
	CustomerLastStatus     int                `json:"customer_last_status"`
	CustomerCreateAt       time.Time          `json:"customer_create_at"`
	CustomerUpdateAt       time.Time          `json:"customer_update_at"`
	Major                  *Major             `json:"jurusan"`
	Address                *[]CustomerAddress `json:"alamat"`
}
