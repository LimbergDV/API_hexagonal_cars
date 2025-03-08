package service

import (
	"api-hexagonal-cars/src/cars/application/repositories"
	"api-hexagonal-cars/src/cars/domain"
)


type Event struct {
	rabbit repositories.IRabbit
}

func NewEvent(rabbit repositories.IRabbit) *Event {
	return &Event{rabbit: rabbit}
}

func (e *Event) Run (car *domain.Car){
	e.rabbit.SendMessageToBroker(car)
}