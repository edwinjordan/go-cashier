package order_repository

import (
	"github.com/jolebo/e-canteen-cashier-api/entity"
)

type ViewOrderDetail struct {
	OrderDetailId              string  `gorm:"column:order_detail_id"`
	OrderDetailParentId        string  `gorm:"column:order_detail_parent_id"`
	OrderDetailProductVarianId string  `gorm:"column:order_detail_product_varian_id"`
	OrderDetailQty             int     `gorm:"column:order_detail_qty"`
	OrderDetailPrice           float64 `gorm:"column:order_detail_price"`
	OrderDetailSubtotal        float64 `gorm:"column:order_detail_subtotal"`
	CustomerName               string  `gorm:"column:customer_name"`
	ProductName                string  `gorm:"column:product_name"`
	VarianName                 string  `gorm:"column:varian_name"`
}

func (ViewOrderDetail) TableName() string {
	return "v_tb_customer_order_detail"
}

func (ViewOrderDetail) FromEntity(e *entity.ViewOrderDetail) *ViewOrderDetail {
	return &ViewOrderDetail{
		OrderDetailId:              e.OrderDetailId,
		OrderDetailParentId:        e.OrderDetailParentId,
		OrderDetailProductVarianId: e.OrderDetailProductVarianId,
		OrderDetailQty:             e.OrderDetailQty,
		OrderDetailPrice:           e.OrderDetailPrice,
		OrderDetailSubtotal:        e.OrderDetailSubtotal,
		CustomerName:               e.CustomerName,
		ProductName:                e.ProductName,
		VarianName:                 e.VarianName,
	}
}

func (model *ViewOrderDetail) ToEntity() *entity.ViewOrderDetail {
	modelData := &entity.ViewOrderDetail{
		OrderDetailId:              model.OrderDetailId,
		OrderDetailParentId:        model.OrderDetailParentId,
		OrderDetailProductVarianId: model.OrderDetailProductVarianId,
		OrderDetailQty:             model.OrderDetailQty,
		OrderDetailPrice:           model.OrderDetailPrice,
		OrderDetailSubtotal:        model.OrderDetailSubtotal,
		CustomerName:               model.CustomerName,
		ProductName:                model.ProductName,
		VarianName:                 model.VarianName,
	}

	return modelData
}
