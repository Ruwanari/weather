package main

import (
	"fmt"
	"weather/crons"
	"weather/repository"
	"weather/usecase"
)

func main() {

	usecase.Init()
	go crons.RefreshWeatherData()
	for {
		var city string
		fmt.Print("Enter city: ")
		fmt.Scanln(&city)
		response, err := repository.GetWeatherInfoMap().Get(city)
		if err != nil {
			fmt.Println("Error retrieving weather data:", err)
		} else {
			if len(response.PropertyEntities.TimeSeries) != 0 {
				fmt.Printf("City: %s\n", city)
				fmt.Printf("Temperature: %.2f %s\n", response.PropertyEntities.TimeSeries[0].Data.Instant.Details.AirTemperature, response.PropertyEntities.Meta.Units.AirTemperature)
				fmt.Printf("Wind Speed: %.2f %s\n", response.PropertyEntities.TimeSeries[0].Data.Instant.Details.WindSpeed, response.PropertyEntities.Meta.Units.WindSpeed)
				fmt.Printf("Relative Humidity: %.2f %s\n", response.PropertyEntities.TimeSeries[0].Data.Instant.Details.RelativeHumidity, response.PropertyEntities.Meta.Units.RelativeHumidity)
			}

		}
	}
}
