package category_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type Category struct {
	CategoryId       string    `json:"category_id"`
	CategoryName     string    `json:"category_name"`
	CategoryDeleteAt time.Time `json:"category_delete_at"`
}

func (Category) TableName() string {
	return "ms_category"
}

func (Category) FromEntity(e *entity.Category) *Category {
	return &Category{
		CategoryId:       e.CategoryId,
		CategoryName:     e.CategoryName,
		CategoryDeleteAt: e.CategoryDeleteAt,
	}
}

func (model *Category) ToEntity() *entity.Category {
	modelData := &entity.Category{
		CategoryId:       model.CategoryId,
		CategoryName:     model.CategoryName,
		CategoryDeleteAt: model.CategoryDeleteAt,
	}
	return modelData
}
