package entity

import "time"

type User struct {
	UserId              string    `json:"user_id"`
	UserName            string    `json:"user_name"`
	UserEmail           string    `json:"user_email"`
	UserPassword        string    `json:"user_password"`
	UserPegawaiId       string    `json:"user_pegawai_id"`
	UserHasMobileAccess int       `json:"user_has_mobile_access"`
	UserRoleId          string    `json:"user_role_id"`
	UserCreateAt        time.Time `json:"-"`
	UserUpdateAt        time.Time `json:"-"`
	Pegawai             *Pegawai  `json:"pegawai"`
}
