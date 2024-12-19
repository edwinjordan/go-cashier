package entity

import "time"

type UserOTP struct {
	OtpId         string    `json:"otp_id"`
	OtpCustomerId string    `json:"otp_customer_id"`
	OtpNumber     string    `json:"otp_number"`
	OtpStatus     int       `json:"-"`
	OtpExpired    time.Time `json:"otp_expired"`
}
