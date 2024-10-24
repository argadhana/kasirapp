package stock

import "time"

type Stock struct {
	ID           int       `json:"id"`
	ProductName  string    `json:"product_name"`
	CodeProduct  string    `json:"code_product"`
	CategoryName string    `json:"category"`
	Stock        int       `json:"stock"`
	SellingPrice float64   `json:"selling_price"`
	BasePrice    float64   `json:"base_price"`
	Date         string    `json:"date"`
	BuyingPrice  float64   `json:"buying_price"`
	Amount       int       `json:"amount"`
	Information  string    `json:"information"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
