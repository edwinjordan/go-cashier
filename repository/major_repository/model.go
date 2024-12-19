package major_repository

import "github.com/jolebo/e-canteen-cashier-api/entity"

type Major struct {
	MajorId       string `gorm:"column:major_id"`
	MajorName     string `gorm:"column:major_name"`
	MajorDeleteAt string `gorm:"column:major_delete_at"`
}

func (Major) TableName() string {
	return "ms_major"
}

func (Major) FromEntity(e *entity.Major) *Major {
	return &Major{
		MajorId:       e.MajorId,
		MajorName:     e.MajorName,
		MajorDeleteAt: e.MajorDeleteAt,
	}
}

func (model *Major) ToEntity() *entity.Major {
	modelData := &entity.Major{
		MajorId:       model.MajorId,
		MajorName:     model.MajorName,
		MajorDeleteAt: model.MajorDeleteAt,
	}

	return modelData
}
