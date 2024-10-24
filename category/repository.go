package category

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(category Category) (Category, error)
	FindByID(ID int) (Category, error)
	FindAll() ([]Category, error)
	FindByName(name string) (Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) (Category, error)
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

func (r *repository) Delete(category Category) (Category, error) {
	err := r.db.Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
