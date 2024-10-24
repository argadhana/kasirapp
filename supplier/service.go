package supplier

import (
	"api-kasirapp/helper"
	"time"

	"golang.org/x/exp/rand"
)

type Service interface {
	CreateSupplier(Input SupplierInput) (Supplier, error)
	GetByID(ID int) (Supplier, error)
	GetByName(name string) (Supplier, error)
	GetAll(limit int, offset int) ([]Supplier, error)
	Update(ID int, Input SupplierInput) (Supplier, error)
	Delete(ID int) (Supplier, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSupplier(input SupplierInput) (Supplier, error) {
	supplier := Supplier{
		Name:    input.Name,
		Address: input.Address,
		Email:   input.Email,
		Phone:   input.Phone,
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	supplier.Code = rand.Intn(90000) + 10000

	if err := helper.ValidateEmail(supplier.Email); err != nil {
		return Supplier{}, err
	}

	if err := helper.ValidatePhoneNumber(supplier.Phone); err != nil {
		return Supplier{}, err
	}

	newSupplier, err := s.repository.Save(supplier)
	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

func (s *service) GetByID(ID int) (Supplier, error) {
	supplier, err := s.repository.FindByID(ID)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (s *service) GetByName(name string) (Supplier, error) {
	supplier, err := s.repository.FindByName(name)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (s *service) GetAll(limit int, offset int) ([]Supplier, error) {
	suppliers, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return suppliers, err
	}
	return suppliers, nil
}

func (s *service) Update(ID int, input SupplierInput) (Supplier, error) {
	supplier, err := s.repository.FindByID(ID)
	if err != nil {
		return supplier, err
	}

	supplier.Name = input.Name
	supplier.Address = input.Address
	supplier.Email = input.Email
	supplier.Phone = input.Phone

	updatedSupplier, err := s.repository.Update(ID, supplier)
	if err != nil {
		return updatedSupplier, err
	}
	return updatedSupplier, nil
}

func (s *service) Delete(ID int) (Supplier, error) {
	supplier, err := s.repository.FindByID(ID)
	if err != nil {
		return supplier, err
	}
	deletedSupplier, err := s.repository.Delete(ID)
	if err != nil {
		return deletedSupplier, err
	}
	return deletedSupplier, nil
}
