package product

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product Product) (Product, error)
	FindByID(ID int) (Product, error)
	FindByName(name string) (Product, error)
	FindAll() ([]Product, error)
	Update(ID int, product Product) (Product, error)
	Delete(ID int) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	var availableID *int // Pointer for checking available ID
	var existingProduct Product

	if err := r.db.Where("code_product = ?", product.CodeProduct).First(&existingProduct).Error; err == nil {
		return product, errors.New("product code already exists") // Return error if product code exists
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Step 1: Check for the first available ID
		if err := tx.Raw("SELECT MIN(id) FROM products WHERE id NOT IN (SELECT id FROM products)").Scan(&availableID).Error; err != nil {
			return err
		}

		// Step 2: If there's an available ID, assign it to the product
		if availableID != nil {
			product.ID = *availableID // Assign available ID
		} else {
			// Use the next sequence number or highest ID + 1
			var maxID *int // Use a pointer to check for NULL
			if err := tx.Model(&Product{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
				return err
			}
			if maxID != nil {
				product.ID = *maxID + 1 // Assign next ID if maxID is not nil
			} else {
				product.ID = 1 // If maxID is nil, start with ID 1 for the first product
			}
		}
		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByName(name string) (Product, error) {
	var product Product

	err := r.db.Where("name = ?", name).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) Update(ID int, product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Delete(ID int) (Product, error) {
	var product Product
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Step 1: Find the product to delete
		if err := tx.Where("id = ?", ID).First(&product).Error; err != nil {
			return err
		}

		// Step 2: Delete the product
		if err := tx.Where("id = ?", ID).Delete(&product).Error; err != nil {
			return err
		}

		// Step 3: Update IDs of all products with ID greater than deleted ID
		if err := tx.Exec("UPDATE products SET id = id - 1 WHERE id > ?", ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return product, err
	}

	return product, nil
}
