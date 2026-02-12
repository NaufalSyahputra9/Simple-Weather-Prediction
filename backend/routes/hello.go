package routes

import (
	"Weather_Prediction/controller"
	"github.com/gin-gonic/gin"
)

func WeatherRoutes(router *gin.Engine, weatherController *controller.WeatherController) {
	router.GET("/hello/:name", weatherController.Hello)
	router.GET("/weather/:name", weatherController.WeatherApi)


}
