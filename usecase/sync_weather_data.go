package usecase

import (
	"log"
	"strconv"
	"weather/external"
	"weather/repository"
)

func Init() {
	cities := external.RetrieveCitiesByCountry()
	for _, city := range cities {
		lat, err := strconv.ParseFloat(city.Latitude, 64)
		if err != nil {
			log.Println("Couldn't parse latitude ")
			return
		}
		lon, err := strconv.ParseFloat(city.Longitude, 64)
		if err != nil {
			log.Println("Couldn't parse longitude ")
			return
		}
		data, err := external.RetrieveWeatherData(lat, lon)
		if err != nil {
			log.Println("Couldn't retrieve weather data ")
			return
		}

		repository.GetWeatherInfoMap().Set(city.Name, data)
	}
}
