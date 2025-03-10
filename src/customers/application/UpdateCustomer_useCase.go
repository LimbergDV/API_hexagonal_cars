package application

import "api-hexagonal-cars/src/customers/domain"



type UpdateCustomer struct {
	db domain.ICustomer
}

func NewUpdateCustomer( db domain.ICustomer) *UpdateCustomer {
	return &UpdateCustomer{db: db}
}

func (uc *UpdateCustomer) Run (id int, car domain.Customer) (uint, error) {
	return uc.db.Update(id, car)
}