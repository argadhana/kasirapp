package customers

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateCustomer(input CustomerInput) (Customer, error)
	FindAll() ([]Customer, error)
	FindByID(ID int) (Customer, error)
	Update(ID int, input CustomerInput) (Customer, error)
	Delete(ID int) (Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCustomer(input CustomerInput) (Customer, error) {
	customer := Customer{}

	customer.Name = input.Name
	customer.Address = input.Address
	customer.Phone = input.Phone
	customer.Email = input.Email

	newCustomer, err := s.repository.Save(customer)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil

}

func (s *service) FindAll() ([]Customer, error) {
	customers, err := s.repository.FindAll()
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (s *service) FindByID(ID int) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, errors.New("category not found")
		}
		return customer, err
	}

	return customer, nil
}

func (s *service) Update(ID int, input CustomerInput) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, errors.New("customer not found")
		}
		return customer, err
	}

	customer.Name = input.Name
	customer.Address = input.Address
	customer.Phone = input.Phone
	customer.Email = input.Email

	updatedCustomer, err := s.repository.Update(customer)
	if err != nil {
		return updatedCustomer, err
	}

	return updatedCustomer, nil
}

func (s *service) Delete(ID int) (Customer, error) {
	customer, err := s.repository.FindByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, errors.New("customer not found")
		}
		return customer, err
	}

	deletedCustomer, err := s.repository.Delete(customer.ID)
	if err != nil {
		return deletedCustomer, err
	}

	return deletedCustomer, nil
}
