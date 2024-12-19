package version_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type VersionAdmin struct {
	VersionId       int       `gorm:"column:version_id"`
	VersionNumber   string    `gorm:"column:version_number"`
	VersionCode     int       `gorm:"column:version_code"`
	VersionChagelog string    `gorm:"column:version_chagelog"`
	VersionDatetime time.Time `gorm:"column:version_datetime"`
}

func (VersionAdmin) TableName() string {
	return "tb_version_cashier"
}
func (model *VersionAdmin) ToEntity() *entity.VersionAdmin {
	modelData := &entity.VersionAdmin{
		VersionId:       model.VersionId,
		VersionNumber:   model.VersionNumber,
		VersionCode:     model.VersionCode,
		VersionChagelog: model.VersionChagelog,
		VersionDatetime: model.VersionDatetime,
	}

	return modelData
}
