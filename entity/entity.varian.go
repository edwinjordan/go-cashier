package entity

type Varian struct {
	ProductVarianId           string `json:"product_varian_id"`
	ProductId                 string `json:"product_id"`
	ProductName               string `json:"product_name"`
	VarianName                string `json:"varian_name"`
	ProductVarianPrice        int    `json:"product_varian_price"`
	ProductVarianQtyBooth     int    `json:"product_varian_qty_booth"`
	ProductVarianQtyWarehouse string `json:"product_varian_qty_warehouse"`
	VarianId                  string `json:"varian_id"`
	ProductVarianQtyLeft      int    `json:"product_varian_qty_left"`
}
