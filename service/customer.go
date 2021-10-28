package service

import (
	"tesss/entity"
	"tesss/input"
	"tesss/repository"
)

type CustomerService interface {
	CustomerServiceGetAll() ([]entity.Customer, error)
	CustomerServiceGetByID(inputID input.InputIDCustomer) (entity.Customer, error)
	CustomerServiceCreate(input input.CustomerInput) (entity.Customer, error)
	CustomerServiceUpdate(inputID input.InputIDCustomer, inputData input.CustomerInput) (entity.Customer, error)
	CustomerServiceDeleteByID(inputID input.InputIDCustomer) (bool, error)
}
type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) *customerService {
	return &customerService{repository}
}
func (s *customerService) CustomerServiceCreate(input input.CustomerInput) (entity.Customer, error) {
	customer := entity.Customer{}
	customer.Name = input.Name
	customer.Address = input.Address
	newCustomer, err := s.repository.SaveCustomer(customer)
	if err != nil {
		return newCustomer, err
	}
	return newCustomer, nil
}
func (s *customerService) CustomerServiceUpdate(inputID input.InputIDCustomer, inputData input.CustomerInput) (entity.Customer, error) {
	customer, err := s.repository.FindByIDCustomer(inputID.ID)
	if err != nil {
		return customer, err
	}
	customer.Name = inputData.Name
	customer.Address = inputData.Address

	updatedCustomer, err := s.repository.UpdateCustomer(customer)

	if err != nil {
		return updatedCustomer, err
	}
	return updatedCustomer, nil
}
func (s *customerService) CustomerServiceGetByID(inputID input.InputIDCustomer) (entity.Customer, error) {
	customer, err := s.repository.FindByIDCustomer(inputID.ID)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
func (s *customerService) CustomerServiceGetAll() ([]entity.Customer, error) {
	customers, err := s.repository.FindAllCustomer()
	if err != nil {
		return customers, err
	}
	return customers, nil
}
func (s *customerService) CustomerServiceDeleteByID(inputID input.InputIDCustomer) (bool, error) {
	_, err := s.repository.FindByIDCustomer(inputID.ID)
	if err != nil {
		return false, err
	}
	err = s.repository.DeleteByIDCustomer(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Generated by Micagen at 28 Oktober 2021