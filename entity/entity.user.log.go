package entity

import "time"

type UserLog struct {
	LogUserId         string    `json:"log_user_id"`
	LogUserUserId     string    `json:"log_user_user_id"`
	LogUserToken      string    `json:"log_user_token"`
	LogUserMetadata   string    `json:"log_user_metadata"`
	LogUserLoginDate  time.Time `json:"log_user_login_date"`
	LogUserLogoutDate time.Time `json:"log_user_logout_date"`
}
