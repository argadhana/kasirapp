package category

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(category Category) (Category, error)
	FindByID(ID int) (Category, error)
	FindAll() ([]Category, error)
	FindByName(name string) (Category, error)
	Update(category Category) (Category, error)
	Delete(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID int) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", ID).First(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *repository) FindByName(name string) (Category, error) {
	var category Category

	err := r.db.Where("name = ?", name).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Delete(ID int) (Category, error) {
	var category Category
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Step 1: Find the product to delete
		if err := tx.Where("id = ?", ID).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("customer not found")
			}
			return err
		}

		// Step 2: Delete the product
		if err := tx.Where("id = ?", ID).Delete(&category).Error; err != nil {
			return err
		}

		// Step 3: Update IDs of all products with ID greater than deleted ID
		if err := tx.Exec("UPDATE products SET id = id - 1 WHERE id > ?", ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return category, err
	}

	return category, nil
}
