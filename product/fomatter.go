package product

type ProductFormatter struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProductType  string `json:"product_type"`
	BasePrice    int    `json:"base_price"`
	SellingPrice int    `json:"selling_price"`
	Stock        int    `json:"stock"`
	CodeProduct  string `json:"code_product"`
	CategoryID   int    `json:"category_id"`
	MinimumStock int    `json:"minimum_stock"`
	Shelf        string `json:"shelf"`
	Weight       int    `json:"weight"`
	Discount     int    `json:"discount"`
	Information  string `json:"information"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:           product.ID,
		Name:         product.Name,
		ProductType:  product.ProductType,
		BasePrice:    product.BasePrice,
		SellingPrice: product.SellingPrice,
		Stock:        product.Stock,
		CodeProduct:  product.CodeProduct,
		CategoryID:   product.CategoryID,
		MinimumStock: product.MinimumStock,
		Shelf:        product.Shelf,
		Weight:       product.Weight,
		Discount:     product.Discount,
		Information:  product.Information,
		CreatedAt:    product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	var productsFormatter []ProductFormatter

	for _, product := range products {
		formatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, formatter)
	}

	return productsFormatter
}
