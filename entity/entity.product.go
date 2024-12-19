package entity

import "time"

type Product struct {
	ProductId         string    `json:"product_id"`
	ProductCode       string    `json:"product_code"`
	ProductName       string    `json:"product_name"`
	ProductCategoryId string    `json:"product_category_id"`
	ProductDesc       string    `json:"product_desc"`
	CategoryName      string    `json:"category_name"`
	ProductCreateAt   time.Time `json:"product_create_at"`
	ProductUpdateAt   time.Time `json:"product_update_at"`
	ProductDeleteAt   time.Time `json:"product_delete_at"`
	ProductPhoto      string    `json:"product_photo"`
	Varian            *[]Varian `json:"varian"`
}
