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

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw("SELECT MIN(id) FROM customers WHERE id NOT IN (SELECT id FROM customers)").Scan(&availableID).Error; err != nil {
			return err
		}
		if availableID != nil {
			supplier.ID = *availableID
		} else {
			var maxID *int
			if err := r.db.Model(&Supplier{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
				return err
			}
			if maxID != nil {
				supplier.ID = *maxID + 1
			} else {
				supplier.ID = 1
			}
		}
		if err := tx.Create(&supplier).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return supplier, err
	} else {
		return supplier, nil
	}
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
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", ID).First(&supplier).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("supplier not found")
			}
			return err
		}

		if err := tx.Where("id = ?", ID).Delete(&supplier).Error; err != nil {
			return err
		}

		if err := tx.Exec("UPDATE suppliers SET id = id - 1 WHERE id > ?", ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}
