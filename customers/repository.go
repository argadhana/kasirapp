package customers

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(customer Customer) (Customer, error)
	FindAll(limit int, offset int) ([]Customer, error)
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

	if err := r.db.Raw("SELECT MIN(id) FROM customers WHERE id NOT IN (SELECT id FROM customers)").Scan(&availableID).Error; err != nil {
		return customer, err
	}

	if availableID != nil {
		customer.ID = *availableID
	} else {
		var maxID *int
		if err := r.db.Model(&Customer{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
			return customer, err
		}
		if maxID != nil {
			customer.ID = *maxID + 1
		} else {
			customer.ID = 1
		}
	}

	if err := r.db.Create(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindAll(limit int, offset int) ([]Customer, error) {
	var customers []Customer

	err := r.db.Limit(limit).Offset(offset).Find(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repository) FindByID(ID int) (Customer, error) {
	var customer Customer
	err := r.db.Where("id = ?", ID).Find(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, errors.New("customer not found")
		}
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
	if err := r.db.Where("id = ?", ID).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, errors.New("customer not found")
		}
		return customer, err
	}

	// Step 2: Delete the customer
	if err := r.db.Delete(&customer).Error; err != nil {
		return customer, err
	}

	// Step 3: Update IDs of all customers with ID greater than deleted ID
	if err := r.db.Exec("UPDATE customers SET id = id - 1 WHERE id > ?", ID).Error; err != nil {
		return customer, err
	}

	return customer, nil
}
