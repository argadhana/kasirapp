package customers

import "gorm.io/gorm"

type Repository interface {
	Save(customer Customer) (Customer, error)
	FindAll() ([]Customer, error)
	FindByID(ID int) (Customer, error)
	Update(customer Customer) (Customer, error)
	Delete(ID int) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(customer Customer) (Customer, error) {
	var availableID *int

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Raw("SELECT MIN(id) FROM customers WHERE id NOT IN (SELECT id FROM customers)").Scan(&availableID).Error; err != nil {
			return err
		}
		if availableID != nil {
			customer.ID = *availableID
		} else {
			var maxID *int
			if err := r.db.Model(&Customer{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
				return err
			}
			if maxID != nil {
				customer.ID = *maxID + 1
			} else {
				customer.ID = 1
			}
		}
		if err := tx.Create(&customer).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return customer, err
	} else {
		return customer, nil
	}
}

func (r *repository) FindAll() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repository) FindByID(ID int) (Customer, error) {
	var customer Customer
	err := r.db.Where("id = ?", ID).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) Update(customer Customer) (Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) Delete(ID int) (Customer, error) {
	var customer Customer
	err := r.db.Where("id = ?", ID).Delete(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}