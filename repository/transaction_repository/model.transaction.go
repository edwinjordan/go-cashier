package transaction_repository

import (
	"time"

	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"github.com/jolebo/e-canteen-cashier-api/repository/customer_repository"
	"github.com/jolebo/e-canteen-cashier-api/repository/user_repository"
	"gorm.io/gorm"
)

type Transaction struct {
	TransId            string                        `gorm:"column:trans_id"`
	TransUserId        string                        `gorm:"column:trans_user_id"`
	TransOrderId       string                        `gorm:"column:trans_order_id"`
	TransInvoice       string                        `gorm:"column:trans_invoice"`
	TransCustomerId    string                        `gorm:"column:trans_customer_id"`
	TransQtyTotal      int                           `gorm:"column:trans_qty_total"`
	TransProductTotal  int                           `gorm:"column:trans_product_total"`
	TransSubtotal      float64                       `gorm:"column:trans_subtotal"`
	TransDiscount      float64                       `gorm:"column:trans_discount"`
	TransTotal         float64                       `gorm:"column:trans_total"`
	TransReceivedTotal float64                       `gorm:"column:trans_received_total"`
	TransRefundTotal   float64                       `gorm:"column:trans_refund_total"`
	TransStatus        int                           `gorm:"column:trans_status"`
	TransCreateAt      time.Time                     `gorm:"column:trans_create_at"`
	TransDetail        *[]TransactionDetail          `gorm:"foreignKey:TransDetailId;references:TransId"`
	Customer           *customer_repository.Customer `gorm:"foreignKey:TransCustomerId;references:CustomerId"`
	User               *user_repository.User         `gorm:"foreignKey:TransUserId;references:UserId"`
}

func (Transaction) TableName() string {
	return "tb_transaction"
}

func (model *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	model.TransId = helpers.GenUUID()
	model.TransCreateAt = helpers.CreateDateTime()
	return
}

func (Transaction) FromEntity(e *entity.Transaction) *Transaction {
	return &Transaction{
		TransId:            e.TransId,
		TransUserId:        e.TransUserId,
		TransInvoice:       e.TransInvoice,
		TransOrderId:       e.TransOrderId,
		TransCustomerId:    e.TransCustomerId,
		TransQtyTotal:      e.TransQtyTotal,
		TransProductTotal:  e.TransProductTotal,
		TransSubtotal:      e.TransSubtotal,
		TransDiscount:      e.TransDiscount,
		TransTotal:         e.TransTotal,
		TransReceivedTotal: e.TransReceivedTotal,
		TransRefundTotal:   e.TransRefundTotal,
		TransStatus:        e.TransStatus,
		TransCreateAt:      e.TransCreateAt,
	}
}

func (model *Transaction) ToEntity() *entity.Transaction {
	modelData := &entity.Transaction{
		TransId:            model.TransId,
		TransUserId:        model.TransUserId,
		TransInvoice:       model.TransInvoice,
		TransOrderId:       model.TransOrderId,
		TransCustomerId:    model.TransCustomerId,
		TransQtyTotal:      model.TransQtyTotal,
		TransProductTotal:  model.TransProductTotal,
		TransSubtotal:      model.TransSubtotal,
		TransDiscount:      model.TransDiscount,
		TransTotal:         model.TransTotal,
		TransReceivedTotal: model.TransReceivedTotal,
		TransRefundTotal:   model.TransRefundTotal,
		TransStatus:        model.TransStatus,
		TransCreateAt:      model.TransCreateAt,
	}

	if model.Customer != nil {
		modelData.Customer = model.Customer.ToEntity()
	}

	if model.User != nil {
		modelData.User = model.User.ToEntity()
	}

	if model.TransDetail != nil {
		var tempMenu []entity.TransactionDetail
		for _, v := range *model.TransDetail {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.TransDetail = &tempMenu
	}

	return modelData
}
