package routers

import (
	"github.com/Rahmatuldani/assignment2/config"
	"github.com/Rahmatuldani/assignment2/controllers"
	"github.com/Rahmatuldani/assignment2/models"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	models := models.NewModels(db)
	controllers := controllers.NewController(models)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	baseRouter := r.Group("/api")
	{
		baseRouter.GET("/orders", controllers.GetOrders)
	}
	return r
}