package entity

type ViewTransactionDetail struct {
	TransDetailId              string  `json:"trans_detail_id"`
	TransDetailParentId        string  `json:"trans_detail_parent_id"`
	TransDetailProductVarianId string  `json:"trans_detail_product_varian_id"`
	TransDetailQty             int     `json:"trans_detail_qty"`
	TransDetailPrice           float64 `json:"trans_detail_price"`
	TransDetailSubtotal        float64 `json:"trans_detail_subtotal"`
	VarianName                 string  `json:"varian_name"`
	ProductName                string  `json:"product_name"`
}
