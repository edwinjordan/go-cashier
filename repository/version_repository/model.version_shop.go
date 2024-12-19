package version_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type VersionShop struct {
	VersionId       int       `gorm:"column:version_id"`
	VersionNumber   string    `gorm:"column:version_number"`
	VersionCode     int       `gorm:"column:version_code"`
	VersionChagelog string    `gorm:"column:version_chagelog"`
	VersionDatetime time.Time `gorm:"column:version_datetime"`
}

func (VersionShop) TableName() string {
	return "tb_version_shop"
}
func (model *VersionShop) ToEntity() *entity.VersionShop {
	modelData := &entity.VersionShop{
		VersionId:       model.VersionId,
		VersionNumber:   model.VersionNumber,
		VersionCode:     model.VersionCode,
		VersionChagelog: model.VersionChagelog,
		VersionDatetime: model.VersionDatetime,
	}

	return modelData
}
