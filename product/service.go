package product

import (
	"api-kasirapp/category"
	"errors"
)

type Service interface {
	CreateProduct(input ProductInput) (Product, error)
	FindProductByID(ID int) (Product, error)
	FindByName(name string) (Product, error)
	FindAll() ([]Product, error)
	UpdateProduct(ID int, input ProductInput) (Product, error)
	DeleteProduct(ID int) (Product, error)
}

type service struct {
	productRepository  Repository
	categoryRepository category.Repository
}

func NewService(productRepository Repository, categoryRepository category.Repository) *service {
	return &service{
		productRepository: productRepository,

		categoryRepository: categoryRepository,
	}
}

func (s *service) CreateProduct(input ProductInput) (Product, error) {

	category, err := s.categoryRepository.FindByID(input.CategoryID)
	if err != nil {
		return Product{}, err
	}

	product := Product{
		Name:         input.Name,
		ProductType:  input.ProductType,
		BasePrice:    input.BasePrice,
		SellingPrice: input.SellingPrice,
		Stock:        input.Stock,
		CodeProduct:  input.CodeProduct,
		CategoryID:   category.ID,
		MinimumStock: input.MinimumStock,
		Shelf:        input.Shelf,
		Weight:       input.Weight,
		Discount:     input.Discount,
		Information:  input.Information,
	}

	newProduct, err := s.productRepository.Save(product)
	if err != nil {
		if err.Error() == "product code already exists" {
			return Product{}, errors.New("product code already exists") // Return specific error message
		}
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) FindProductByID(ID int) (Product, error) {
	product, err := s.productRepository.FindByID(ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) FindByName(name string) (Product, error) {
	product, err := s.productRepository.FindByName(name)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) FindAll() ([]Product, error) {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) UpdateProduct(ID int, input ProductInput) (Product, error) {
	product, err := s.productRepository.FindByID(ID)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.ProductType = input.ProductType
	product.BasePrice = input.BasePrice
	product.SellingPrice = input.SellingPrice
	product.Stock = input.Stock
	product.CodeProduct = input.CodeProduct
	product.CategoryID = input.CategoryID
	product.MinimumStock = input.MinimumStock
	product.Shelf = input.Shelf
	product.Weight = input.Weight
	product.Discount = input.Discount
	product.Information = input.Information

	updatedProduct, err := s.productRepository.Update(ID, product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(ID int) (Product, error) {
	product, err := s.productRepository.FindByID(ID)
	if err != nil {
		return product, err
	}

	deletedProduct, err := s.productRepository.Delete(ID)
	if err != nil {
		return deletedProduct, err
	}

	return deletedProduct, nil
}
