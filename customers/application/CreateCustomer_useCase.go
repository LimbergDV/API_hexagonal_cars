package application

import "api-hexagonal-cars/customers/domain"

type CreateCustomer struct{
	db domain.ICustomer
}

func NewCreateCustomer (db domain.ICustomer) *CreateCustomer {
	return &CreateCustomer{db: db}
}

func (cc *CreateCustomer) Run (customer domain.Customer) (uint, error) {
	return cc.db.Save(customer)
}