package discount

type Service interface {
	Create(input DiscountInput) (Discount, error)
	GetByID(id int) (Discount, error)
	GetAll() ([]Discount, error)
	Update(ID int, input DiscountInput) (Discount, error)
	Delete(ID int) (Discount, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input DiscountInput) (Discount, error) {
	discount := Discount{
		Name:       input.Name,
		Percentage: input.Percentage,
	}
	newDiscount, err := s.repository.Save(discount)
	if err != nil {
		return newDiscount, err
	}
	return newDiscount, nil
}

func (s *service) GetByID(id int) (Discount, error) {
	discount, err := s.repository.FindByID(id)
	if err != nil {
		return discount, err
	}
	return discount, nil
}

func (s *service) GetAll() ([]Discount, error) {
	discounts, err := s.repository.FindAll()
	if err != nil {
		return discounts, err
	}
	return discounts, nil
}

func (s *service) Update(ID int, input DiscountInput) (Discount, error) {
	discount, err := s.repository.FindByID(ID)
	if err != nil {
		return discount, err
	}

	discount.Name = input.Name
	discount.Percentage = input.Percentage

	updatedDiscount, err := s.repository.Update(ID, discount)
	if err != nil {
		return updatedDiscount, err
	}
	return updatedDiscount, nil
}

func (s *service) Delete(ID int) (Discount, error) {
	discount, err := s.repository.FindByID(ID)
	if err != nil {
		return discount, err
	}
	deletedDiscount, err := s.repository.Delete(ID)
	if err != nil {
		return deletedDiscount, err
	}
	return deletedDiscount, nil
}
