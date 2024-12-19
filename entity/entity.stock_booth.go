package entity

import "time"

type StockBooth struct {
	ProductStokId              string    `json:"product_stok_id"`
	ProductStokProductVarianId string    `json:"product_stok_product_varian_id"`
	ProductStokFirstQty        int       `json:"product_stok_first_qty"`
	ProductStokQty             int       `json:"product_stok_qty"`
	ProductStokLastQty         int       `json:"product_stok_last_qty"`
	ProductStokJenis           string    `json:"product_stok_jenis"`
	ProductStokDatetime        time.Time `json:"product_stok_datetime"`
	ProductStokPegawaiId       string    `json:"product_stok_pegawai_id"`
}
