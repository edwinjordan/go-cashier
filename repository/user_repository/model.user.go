package user_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/repository/pegawai_repository"
)

type User struct {
	UserId              string                     `gorm:"column:user_id"`
	UserName            string                     `gorm:"column:user_name"`
	UserEmail           string                     `gorm:"column:user_email"`
	UserPassword        string                     `gorm:"column:user_password"`
	UserPegawaiId       string                     `gorm:"column:user_pegawai_id"`
	UserRoleId          string                     `gorm:"column:user_role_id"`
	UserHasMobileAccess int                        `gorm:"column:user_has_mobile_access"`
	UserCreateAt        time.Time                  `gorm:"column:user_create_at"`
	UserUpdateAt        time.Time                  `gorm:"column:user_create_at"`
	Pegawai             pegawai_repository.Pegawai `gorm:"foreignKey:PegawaiId;references:UserPegawaiId"`
}

func (User) TableName() string {
	return "tb_user"
}

func (User) FromEntity(e *entity.User) *User {
	return &User{
		UserId:              e.UserId,
		UserName:            e.UserName,
		UserEmail:           e.UserEmail,
		UserPassword:        e.UserPassword,
		UserPegawaiId:       e.UserPegawaiId,
		UserRoleId:          e.UserRoleId,
		UserHasMobileAccess: e.UserHasMobileAccess,
		UserCreateAt:        e.UserCreateAt,
		UserUpdateAt:        e.UserUpdateAt,
	}
}

func (model *User) ToEntity() *entity.User {
	modelData := &entity.User{
		UserId:              model.UserId,
		UserName:            model.UserName,
		UserEmail:           model.UserEmail,
		UserPassword:        model.UserPassword,
		UserPegawaiId:       model.UserPegawaiId,
		UserRoleId:          model.UserRoleId,
		UserHasMobileAccess: model.UserHasMobileAccess,
		UserCreateAt:        model.UserCreateAt,
		UserUpdateAt:        model.UserUpdateAt,
		Pegawai:             model.Pegawai.ToEntity(),
	}

	return modelData
}
