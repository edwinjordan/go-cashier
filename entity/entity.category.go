package entity

import "time"

type Category struct {
	CategoryId       string    `json:"category_id"`
	CategoryName     string    `json:"category_name"`
	CategoryDeleteAt time.Time `json:"category_delete_at"`
}
