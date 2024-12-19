package entity

import (
	"database/sql"
	"time"
)

type CustomerOrder struct {
	OrderId                string       `json:"order_id"`
	OrderCustomerId        string       `json:"order_customer_id"`
	OrderInvNumber         string       `json:"order_inv_number"`
	OrderAddressId         string       `json:"order_address_id"`
	OrderDeliveryType      string       `json:"order_delivery_type"`
	OrderTotalItem         int          `json:"order_total_item"`
	OrderSubtotal          float64      `json:"order_subtotal"`
	OrderDiscount          float64      `json:"order_discount"`
	OrderTotal             float64      `json:"order_total"`
	OrderNotes             string       `json:"order_notes"`
	OrderCancelNotes       string       `json:"order_cancel_notes"`
	OrderStatus            int          `json:"order_status"`
	OrderProcessedDatetime sql.NullTime `json:"order_processed_datetime"`
	OrderProcessedBy       string       `json:"order_processed_by"`
	OrderFinishedDatetime  sql.NullTime `json:"order_finished_datetime"`
	OrderFinishedBy        string       `json:"order_finished_by"`
	OrderCreateAt          time.Time    `json:"order_create_at"`
	OrderDetail            *[]CustomerOrderDetail
	Customer               *CustomerResponse `json:"customer"`
	Address                *CustomerAddress  `json:"address"`
}
