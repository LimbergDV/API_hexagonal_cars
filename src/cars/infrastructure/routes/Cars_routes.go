package routes

import (
	"api-hexagonal-cars/src/cars/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine) {
	
	carsRoutes := r.Group("/cars") 
	{
		carsRoutes.POST("/", controllers.NewCreateCarController().Run)
		carsRoutes.GET("/", controllers.NewGetAllCarsController().Run)
		carsRoutes.PUT("/:id", controllers.NewUpdateCarByIdController().Run)
		carsRoutes.DELETE("/:id", controllers.NewDeleteCarByIdController().Run)
	}
}