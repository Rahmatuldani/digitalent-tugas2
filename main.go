package main

import (
	"github.com/Rahmatuldani/assignment2/config"
	"github.com/Rahmatuldani/assignment2/routers"
	docs "github.com/Rahmatuldani/assignment2/docs"
)

// @title Orders API
// @version 1.0
// @description Sample service for orders
// @host localhost:5000
// @BasePath /api

func main() {
	config.DBConnect()
	docs.SwaggerInfo.BasePath = "/api"
	routers.Routers().Run(":5000")
}