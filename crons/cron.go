package crons

import (
	"log"
	"time"
	"weather/usecase"
)

func RefreshWeatherData() {
	log.Println("Refreshing Data")
	t := time.Duration(600)
	for range time.NewTicker(t * time.Second).C {
		usecase.Init()
	}
}
