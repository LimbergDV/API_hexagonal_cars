package application

import "api-hexagonal-cars/src/cars/domain"



type DeleteCar struct{
	db domain.ICar
}

func NewDeleteCar (db domain.ICar) *DeleteCar {
	return &DeleteCar{db: db}
}

func (dc *DeleteCar) Run (id int) (uint, error) {
	return dc.db.Delete(id)
}