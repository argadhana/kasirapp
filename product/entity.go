package product

import (
	"time"
)

type Product struct {
	ID           int
	Name         string
	ProductType  string
	BasePrice    int
	SellingPrice int
	Stock        int
	CodeProduct  string
	CategoryID   int
	MinimumStock int
	Shelf        string
	Weight       int
	Discount     int
	Information  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
