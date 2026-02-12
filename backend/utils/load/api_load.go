package load

import (
	"Weather_Prediction/models"
	"encoding/json"
	"os"
)

func LoadApiConfig(filename string) (models.ApiConfigData, error){
	bytes, err := os.ReadFile(filename)

	if err != nil{
		return models.ApiConfigData{}, err
	}

	var c models.ApiConfigData 
	err = json.Unmarshal(bytes, &c)

	if err != nil{
		return models.ApiConfigData{}, err 
	}

	return c, nil

}