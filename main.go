package main

import (
	routesCars "api-hexagonal-cars/cars/infrastructure/routes"
	routesCustomers "api-hexagonal-cars/customers/infrastructure/routes"
	cars "api-hexagonal-cars/cars/infrastructure"
	customers "api-hexagonal-cars/customers/infrastructure"
	
	

	"github.com/gin-gonic/gin"
)

func main () {
	cars.GoMySQL()
	customers.GoMySQL()

	r := gin.Default()

	routesCars.Routes(r)
	routesCustomers.Routes(r)
	
	r.Run(":8080")
}