package entity

type Login struct {
	UserId             string `json:"user_id"`
	UserEmail          string `json:"user_email" validate:"required"`
	UserPassword       string `json:"user_password" validate:"required"`
	UserFcmToken       string `json:"user_fcmtoken"`
	UserDeviceMetadata string `json:"user_device_metadata"`
}
