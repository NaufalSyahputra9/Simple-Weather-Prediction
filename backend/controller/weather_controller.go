package controller

import (
	"Weather_Prediction/models"
	apix "Weather_Prediction/utils/api"
	load "Weather_Prediction/utils/load"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)



func query(city string)(models.WeatherData, error){
	
	apiConfig, err := load.LoadApiConfig(`.apiConfig`)
	if err != nil {
		return models.WeatherData{}, err
	}
	
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s&units=metric",
	  apiConfig.OpenWeatherMapApiKey,
		city,
	)

	resp, err := http.Get(url)
	if err != nil {
		return models.WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.WeatherData{}, fmt.Errorf("openweather error: %s", resp.Status)
	}

	var d models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return models.WeatherData{}, err
	}

	return d, nil
}

type WeatherController struct {
	// kalau nanti mau inject service / API key, taruh di sini
}

func NewWeatherController() *WeatherController {
	return &WeatherController{}
}

func (wc *WeatherController) Hello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}

func (wc *WeatherController) WeatherApi(c *gin.Context) {
	city := c.Param("name")
	data, err := query(city)


	if err != nil {
		c.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid city name",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}