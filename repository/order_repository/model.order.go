package order_repository

import (
	"database/sql"
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_address_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_repository"
	"gorm.io/gorm"
)

type CustomerOrder struct {
	OrderId                string                                       `gorm:"column:order_id"`
	OrderCustomerId        string                                       `gorm:"column:order_customer_id"`
	OrderInvNumber         string                                       `gorm:"column:order_inv_number"`
	OrderAddressId         string                                       `gorm:"column:order_address_id"`
	OrderDeliveryType      string                                       `gorm:"column:order_delivery_type"`
	OrderTotalItem         int                                          `gorm:"column:order_total_item"`
	OrderSubtotal          float64                                      `gorm:"column:order_subtotal"`
	OrderDiscount          float64                                      `gorm:"column:order_discount"`
	OrderTotal             float64                                      `gorm:"column:order_total"`
	OrderNotes             string                                       `gorm:"column:order_notes"`
	OrderStatus            int                                          `gorm:"column:order_status"`
	OrderCancelNotes       string                                       `gorm:"column:order_cancel_notes"`
	OrderProcessedDatetime sql.NullTime                                 `gorm:"column:order_processed_datetime"`
	OrderProcessedBy       string                                       `gorm:"column:order_processed_by"`
	OrderFinishedDatetime  sql.NullTime                                 `gorm:"column:order_finished_datetime"`
	OrderFinishedBy        string                                       `gorm:"column:order_finished_by"`
	OrderCreateAt          time.Time                                    `gorm:"column:order_create_at"`
	OrderDetail            *[]CustomerOrderDetail                       `gorm:"foreignKey:OrderDetailParentId;references:OrderId"`
	Customer               *customer_repository.Customer                `gorm:"foreignKey:OrderCustomerId;references:CustomerId"`
	Address                *customer_address_repository.CustomerAddress `gorm:"foreignKey:OrderAddressId;references:AddressId"`
}

func (CustomerOrder) TableName() string {
	return "tb_customer_order"
}

func (model *CustomerOrder) BeforeCreate(tx *gorm.DB) (err error) {
	model.OrderId = helpers.GenUUID()
	model.OrderCreateAt = helpers.CreateDateTime()
	model.OrderCancelNotes = ""
	model.OrderProcessedBy = ""
	model.OrderFinishedBy = ""
	model.OrderStatus = 4
	return
}

func (CustomerOrder) FromEntity(e *entity.CustomerOrder) *CustomerOrder {
	return &CustomerOrder{
		OrderId:                e.OrderId,
		OrderCustomerId:        e.OrderCustomerId,
		OrderInvNumber:         e.OrderInvNumber,
		OrderAddressId:         e.OrderAddressId,
		OrderDeliveryType:      e.OrderDeliveryType,
		OrderTotalItem:         e.OrderTotalItem,
		OrderSubtotal:          e.OrderSubtotal,
		OrderDiscount:          e.OrderDiscount,
		OrderTotal:             e.OrderTotal,
		OrderNotes:             e.OrderNotes,
		OrderCancelNotes:       e.OrderCancelNotes,
		OrderStatus:            e.OrderStatus,
		OrderProcessedDatetime: e.OrderProcessedDatetime,
		OrderProcessedBy:       e.OrderProcessedBy,
		OrderFinishedBy:        e.OrderFinishedBy,
		OrderFinishedDatetime:  e.OrderFinishedDatetime,
		OrderCreateAt:          e.OrderCreateAt,
	}
}

func (model *CustomerOrder) ToEntity() *entity.CustomerOrder {
	modelData := &entity.CustomerOrder{
		OrderId:                model.OrderId,
		OrderCustomerId:        model.OrderCustomerId,
		OrderInvNumber:         model.OrderInvNumber,
		OrderAddressId:         model.OrderAddressId,
		OrderDeliveryType:      model.OrderDeliveryType,
		OrderTotalItem:         model.OrderTotalItem,
		OrderSubtotal:          model.OrderSubtotal,
		OrderDiscount:          model.OrderDiscount,
		OrderTotal:             model.OrderTotal,
		OrderNotes:             model.OrderNotes,
		OrderCancelNotes:       model.OrderCancelNotes,
		OrderStatus:            model.OrderStatus,
		OrderProcessedDatetime: model.OrderProcessedDatetime,
		OrderFinishedDatetime:  model.OrderFinishedDatetime,
		OrderProcessedBy:       model.OrderProcessedBy,
		OrderFinishedBy:        model.OrderFinishedBy,
		OrderCreateAt:          model.OrderCreateAt,
	}

	if model.OrderDetail != nil {
		var tempMenu []entity.CustomerOrderDetail
		for _, v := range *model.OrderDetail {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.OrderDetail = &tempMenu
	}

	if model.Customer != nil {
		modelData.Customer = model.Customer.ToEntity()
	}
	if model.Address != nil {
		modelData.Address = model.Address.ToEntity()
	}

	return modelData
}
