package repositories

import "api-hexagonal-cars/src/cars/domain"

type IRabbit interface {
	SendMessageToBroker(car *domain.Car)
}