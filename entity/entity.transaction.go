package entity

import "time"

type Transaction struct {
	TransId            string               `json:"trans_id"`
	TransUserId        string               `json:"trans_user_id"`
	TransCustomerId    string               `json:"trans_customer_id"`
	TransOrderId       string               `json:"trans_order_id"`
	TransInvoice       string               `json:"trans_invoice"`
	TransQtyTotal      int                  `json:"trans_qty_total"`
	TransProductTotal  int                  `json:"trans_product_total"`
	TransSubtotal      float64              `json:"trans_subtotal"`
	TransDiscount      float64              `json:"trans_discount"`
	TransTotal         float64              `json:"trans_total"`
	TransReceivedTotal float64              `json:"trans_received_total"`
	TransRefundTotal   float64              `json:"trans_refund_total"`
	TransStatus        int                  `json:"trans_status"`
	TransCreateAt      time.Time            `json:"trans_create_at"`
	TransDetail        *[]TransactionDetail `json:"trans_detail"`
	Customer           *CustomerResponse    `json:"customer"`
	User               *User                `json:"user"`
}
