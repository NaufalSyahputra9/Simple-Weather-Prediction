package main

import (
	"Weather_Prediction/controller"
	"Weather_Prediction/middleware"
	"Weather_Prediction/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init gin
	router := gin.Default()

	// Init controller
	weatherController := controller.NewWeatherController()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// Register routes
	routes.WeatherRoutes(router, weatherController)

	// Run server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}