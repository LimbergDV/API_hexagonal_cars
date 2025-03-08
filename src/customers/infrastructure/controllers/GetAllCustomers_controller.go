package controllers

import (
	"api-hexagonal-cars/src/customers/application"
	"api-hexagonal-cars/src/customers/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllCustomersController struct {
	app *application.GetAllCustomers
}

func NewGetAllCustomersController () *GetAllCustomersController{
	mysql := infrastructure.GetMySQL()
	app := application.NewGetAllCustomers(mysql)
	return &GetAllCustomersController{app: app}
}

func (ctrl *GetAllCustomersController) Run (c *gin.Context) {
	res := ctrl.app.Run()

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "No se encontró ningún cliente"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"clientes": res})
	}

	
}