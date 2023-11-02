package repository

import "weather/schemas"

type WeatherInfo interface {
	Get(cityName string) (resp schemas.GeoResp, err error)
	Set(cityName string, resp schemas.GeoResp)
}
