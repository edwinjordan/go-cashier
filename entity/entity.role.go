package entity

import "time"

type Role struct {
	RoleId       string    `json:"role_id"`
	RoleName     string    `json:"role_name"`
	RoleCode     string    `json:"role_code"`
	RoleCreateAt time.Time `json:"-"`
	RoleUpdateAt time.Time `json:"-"`
}
