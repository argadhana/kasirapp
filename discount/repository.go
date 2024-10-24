package discount

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(discount Discount) (Discount, error)
	FindByID(id int) (Discount, error)
	FindAll() ([]Discount, error)
	Update(ID int, discount Discount) (Discount, error)
	Delete(ID int) (Discount, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(discount Discount) (Discount, error) {
	var availableID *int

	if err := r.db.Raw("SELECT MIN(id) FROM discounts WHERE id NOT IN (SELECT id FROM discounts)").Scan(&availableID).Error; err != nil {
		return discount, err
	}

	if availableID != nil {
		discount.ID = *availableID
	} else {
		var maxID *int
		if err := r.db.Model(&Discount{}).Select("MAX(id)").Scan(&maxID).Error; err != nil {
			return discount, err
		}
		if maxID != nil {
			discount.ID = *maxID + 1
		} else {
			discount.ID = 1
		}
	}
	err := r.db.Create(&discount).Error
	if err != nil {
		return discount, err
	}
	return discount, nil
}

func (r *repository) FindByID(id int) (Discount, error) {
	var discount Discount
	err := r.db.Where("id = ?", id).First(&discount).Error
	if err != nil {
		return discount, err
	}
	return discount, nil
}

func (r *repository) FindAll() ([]Discount, error) {
	var discounts []Discount
	err := r.db.Find(&discounts).Error
	if err != nil {
		return discounts, err
	}
	return discounts, nil
}

func (r *repository) Update(ID int, input Discount) (Discount, error) {
	var discount Discount
	if err := r.db.Where("id = ?", ID).First(&discount).Error; err != nil {
		return discount, err
	}
	discount.Name = input.Name
	discount.Percentage = input.Percentage
	err := r.db.Save(&discount).Error
	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) Delete(ID int) (Discount, error) {
	var discount Discount

	err := r.db.Where("id = ?", ID).First(&discount).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discount, errors.New("discount not found")
		}
		return discount, err
	}

	err = r.db.Where("id = ?", ID).Delete(&discount).Error
	if err != nil {
		return discount, err
	}

	err = r.db.Exec("UPDATE discounts SET id = id - 1 WHERE id > ?", ID).Error
	if err != nil {
		return discount, err
	}

	return discount, nil
}
