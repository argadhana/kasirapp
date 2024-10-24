package customers

import (
	"api-kasirapp/helper"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateCustomer(input CustomerInput) (Customer, error)
	FindAll(limit int, offset int) ([]Customer, error)
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

	if err := helper.ValidateEmail(customer.Email); err != nil {
		return Customer{}, err
	}

	if err := helper.ValidatePhoneNumber(customer.Phone); err != nil {
		return Customer{}, err
	}

	newCustomer, err := s.repository.Save(customer)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil

}

func (s *service) FindAll(limit int, offset int) ([]Customer, error) {
	customers, err := s.repository.FindAll(limit, offset)
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
		return customer, err
	}

	deletedCustomer, err := s.repository.Delete(ID)
	if err != nil {
		return deletedCustomer, err
	}

	return deletedCustomer, nil
}
