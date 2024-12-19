package varian_repository

import (
	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type Varian struct {
	ProductVarianId           string `gorm:"column:product_varian_id"`
	ProductId                 string `gorm:"column:product_id"`
	ProductName               string `gorm:"column:product_name"`
	VarianName                string `gorm:"column:varian_name"`
	ProductVarianPrice        int    `gorm:"column:product_varian_price"`
	ProductVarianQtyBooth     int    `gorm:"column:product_varian_qty_booth"`
	ProductVarianQtyWarehouse string `gorm:"column:product_varian_qty_warehouse"`
	VarianId                  string `gorm:"column:varian_id"`
	ProductVarianQtyLeft      int    `gorm:"column:product_varian_qty_left"`
}

func (Varian) TableName() string {
	return "v_ms_product_varian"
}

func (Varian) FromEntity(e *entity.Varian) *Varian {
	return &Varian{
		ProductVarianId:           e.ProductVarianId,
		ProductId:                 e.ProductId,
		ProductName:               e.ProductName,
		VarianName:                e.VarianName,
		ProductVarianPrice:        e.ProductVarianPrice,
		ProductVarianQtyBooth:     e.ProductVarianQtyBooth,
		ProductVarianQtyWarehouse: e.ProductVarianQtyWarehouse,
		VarianId:                  e.VarianId,
		ProductVarianQtyLeft:      e.ProductVarianQtyLeft,
	}
}

func (model *Varian) ToEntity() *entity.Varian {
	modelData := &entity.Varian{
		ProductVarianId:           model.ProductVarianId,
		ProductId:                 model.ProductId,
		ProductName:               model.ProductName,
		VarianName:                model.VarianName,
		ProductVarianPrice:        model.ProductVarianPrice,
		ProductVarianQtyBooth:     model.ProductVarianQtyBooth,
		ProductVarianQtyWarehouse: model.ProductVarianQtyWarehouse,
		ProductVarianQtyLeft:      model.ProductVarianQtyLeft,
		VarianId:                  model.VarianId,
	}

	return modelData
}
