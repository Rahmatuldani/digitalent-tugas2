package routers

import (
	"github.com/Rahmatuldani/assignment2/config"
	"github.com/Rahmatuldani/assignment2/controllers"
	"github.com/Rahmatuldani/assignment2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	validate := validator.New()
	models := models.NewModels(db)
	controllers := controllers.NewController(models, validate)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	baseRouter := r.Group("/api")
	{
		baseRouter.GET("/orders", controllers.GetOrders)
		baseRouter.POST("/orders", controllers.CreateOrder)
		baseRouter.PUT("/orders/:id", controllers.UpdateOrder)
		baseRouter.DELETE("/orders/:id", controllers.DeleteOrder)
	}
	return r
}