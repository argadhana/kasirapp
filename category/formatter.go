package category

type CategoryFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	var categoriesFormatter []CategoryFormatter

	for _, category := range categories {
		formatter := FormatCategory(category)
		categoriesFormatter = append(categoriesFormatter, formatter)
	}

	return categoriesFormatter
}
