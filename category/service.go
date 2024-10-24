package category

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Save(input CategoryInput) (Category, error)
	FindByID(ID int) (Category, error)
	FindAll() ([]Category, error)
	Update(ID int, input CategoryInput) (Category, error)
	Delete(ID int) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input CategoryInput) (Category, error) {
	category := Category{}

	category.Name = input.Name

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) FindByID(ID int) (Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return category, errors.New("category not found")
		}
		return category, err
	}

	return category, nil
}

func (s *service) FindAll() ([]Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *service) Update(ID int, input CategoryInput) (Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	category.Name = input.Name

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *service) Delete(ID int) (Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	deletedCategory, err := s.repository.Delete(category)
	if err != nil {
		return deletedCategory, err
	}

	return deletedCategory, nil
}
