package repository

import (
	"log"
	"sync"
	"weather/schemas"
)

var (
	weatherInfoSyncMap     sync.Map
	weatherInfoMapInstance = weatherInfo{}
)

type weatherInfo struct{}

func GetWeatherInfoMap() WeatherInfo {
	return weatherInfoMapInstance
}

func (weatherInfo) Get(cityName string) (resp schemas.GeoResp, err error) {
	value, ok := weatherInfoSyncMap.Load(cityName)
	if !ok {
		log.Println("This city is not in the map")
		return
	}

	weatherInfo, ok := value.(schemas.GeoResp)
	if !ok {
		log.Println("Could not parse weather info")
		return
	}

	return weatherInfo, nil
}

func (weatherInfo) Set(cityName string, resp schemas.GeoResp) {
	weatherInfoSyncMap.Store(cityName, resp)
}
