package supplier

type SupplierFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Code      int    `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FormatSupplier(supplier Supplier) SupplierFormatter {
	formatter := SupplierFormatter{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		CreatedAt: supplier.CreatedAt.String(),
		UpdatedAt: supplier.UpdatedAt.String(),
	}
	return formatter
}

func FormatSuppliers(suppliers []Supplier) []SupplierFormatter {
	formatters := []SupplierFormatter{}
	for _, supplier := range suppliers {
		formatters = append(formatters, FormatSupplier(supplier))
	}
	return formatters
}
