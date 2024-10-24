package supplier

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(supplier Supplier) (Supplier, error)
	FindByID(ID int) (Supplier, error)
	FindByName(name string) (Supplier, error)
	FindAll(limit int, offset int) ([]Supplier, error)
	Update(ID int, supplier Supplier) (Supplier, error)
	Delete(ID int) (Supplier, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(supplier Supplier) (Supplier, error) {
	var availableID *int

	if err := r.db.Raw("SELECT MIN(id) FROM suppliers WHERE id NOT IN (SELECT id FROM suppliers)").Scan(&availableID).Error; err != nil {
		return supplier, err
	}

	if availableID != nil {
		supplier.ID = *availableID
	} else {
		var maxID *int
		if err := r.db.Model(&Supplier{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
			return supplier, err
		}
		if maxID != nil {
			supplier.ID = *maxID + 1
		} else {
			supplier.ID = 1
		}
	}

	if err := r.db.Create(&supplier).Error; err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindByID(ID int) (Supplier, error) {
	var supplier Supplier

	err := r.db.Where("id = ?", ID).Find(&supplier).Error
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (r *repository) FindByName(name string) (Supplier, error) {
	var supplier Supplier

	err := r.db.Where("name = ?", name).Find(&supplier).Error
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (r *repository) FindAll(limit int, offset int) ([]Supplier, error) {
	var suppliers []Supplier

	err := r.db.Limit(limit).Offset(offset).Find(&suppliers).Error
	if err != nil {
		return suppliers, err
	}

	return suppliers, nil
}

func (r *repository) Update(ID int, supplier Supplier) (Supplier, error) {
	if err := r.db.Model(&Supplier{}).Where("id = ?", ID).Updates(&supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (r *repository) Delete(ID int) (Supplier, error) {
	var supplier Supplier

	err := r.db.Where("id = ?", ID).First(&supplier).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return supplier, errors.New("supplier not found")
		}
		return supplier, err
	}

	err = r.db.Where("id = ?", ID).Delete(&supplier).Error
	if err != nil {
		return supplier, err
	}

	err = r.db.Exec("UPDATE suppliers SET id = id - 1 WHERE id > ?", ID).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}
