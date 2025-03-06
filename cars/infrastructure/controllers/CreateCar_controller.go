package controllers

import (
	"api-hexagonal-cars/cars/application"
	"api-hexagonal-cars/cars/domain"
	"api-hexagonal-cars/cars/infrastructure"
	"api-hexagonal-cars/cars/infrastructure/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	app *application.CreateCar
}

func NewCreateCarController() *CreateCarController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreateCar(mysql)
	return &CreateCarController{app: app}
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
	}
}