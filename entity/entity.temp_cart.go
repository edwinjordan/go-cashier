package entity

type TempCart struct {
	TempCartId              string `json:"temp_cart_id"`
	TempCartOrderId         string `json:"temp_cart_order_id"`
	TempCartProductVarianId string `json:"temp_cart_product_varian_id"`
	TempCartUserId          string `json:"temp_cart_user_id"`
	TempCartQty             int    `json:"temp_cart_qty"`
}
