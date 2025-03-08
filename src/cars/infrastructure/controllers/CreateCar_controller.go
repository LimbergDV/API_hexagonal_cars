package controllers

import (
	application "api-hexagonal-cars/src/cars/application/UseCases"
	"api-hexagonal-cars/src/cars/application/service"
	"api-hexagonal-cars/src/cars/domain"
	"api-hexagonal-cars/src/cars/infrastructure"
	"api-hexagonal-cars/src/cars/infrastructure/adapters"
	"api-hexagonal-cars/src/cars/infrastructure/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	app *application.CreateCar
	eventService *service.Event
}

func NewCreateCarController() *CreateCarController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreateCar(mysql)
	rabbit := adapters.NewRabbitMq()
	eventService := service.NewEvent(rabbit)
	return &CreateCarController {app: app, eventService: eventService}
}

func (cc_c *CreateCarController) Run (c *gin.Context){
	var cars domain.Car

	if err := c.ShouldBindJSON(&cars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Datos inválidos" + err.Error()})
		return
	}

	if err := validators.CheckCar(cars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos invalidos" + err.Error()})
	}
	
	rowsAffected, err := cc_c.app.Run(cars)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if rowsAffected == 0{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Carro creado"})
		c.JSON(http.StatusOK, cars)

		cc_c.eventService.Run(&cars)
		if err != nil {
		// Puedes registrar este error en los logs si el evento no se envió correctamente
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending event to RabbitMQ"})
	}

	}
}