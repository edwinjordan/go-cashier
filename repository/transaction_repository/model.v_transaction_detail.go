package transaction_repository

import (
	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type ViewTransactionDetail struct {
	TransDetailId              string  `gorm:"column:trans_detail_id"`
	TransDetailParentId        string  `gorm:"column:trans_detail_parent_id"`
	TransDetailProductVarianId string  `gorm:"column:trans_detail_product_varian_id"`
	TransDetailQty             int     `gorm:"column:trans_detail_qty"`
	TransDetailPrice           float64 `gorm:"column:trans_detail_price"`
	TransDetailSubtotal        float64 `gorm:"column:trans_detail_subtotal"`
	VarianName                 string  `gorm:"column:varian_name"`
	ProductName                string  `gorm:"column:product_name"`
}

func (ViewTransactionDetail) TableName() string {
	return "v_tb_trans_detail"
}

func (ViewTransactionDetail) FromEntity(e *entity.ViewTransactionDetail) *ViewTransactionDetail {
	return &ViewTransactionDetail{
		TransDetailId:              e.TransDetailId,
		TransDetailParentId:        e.TransDetailParentId,
		TransDetailProductVarianId: e.TransDetailProductVarianId,
		TransDetailQty:             e.TransDetailQty,
		TransDetailPrice:           e.TransDetailPrice,
		TransDetailSubtotal:        e.TransDetailSubtotal,
		VarianName:                 e.VarianName,
		ProductName:                e.ProductName,
	}
}

func (model *ViewTransactionDetail) ToEntity() *entity.ViewTransactionDetail {
	modelData := &entity.ViewTransactionDetail{
		TransDetailId:              model.TransDetailId,
		TransDetailParentId:        model.TransDetailParentId,
		TransDetailProductVarianId: model.TransDetailProductVarianId,
		TransDetailQty:             model.TransDetailQty,
		TransDetailPrice:           model.TransDetailPrice,
		TransDetailSubtotal:        model.TransDetailSubtotal,
		VarianName:                 model.VarianName,
		ProductName:                model.ProductName,
	}

	return modelData
}
