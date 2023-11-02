package external

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weather/schemas"
)

func RetrieveWeatherData(lat float64, lon float64) (schemas.GeoResp, error) {

	var geoResponse schemas.GeoResp

	url := fmt.Sprintf("https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=%v&lon=%v", lat, lon)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "weather/1.0")
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println("Error retrieving HTTP response")
		return geoResponse, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Could not get a success response")
		return geoResponse, err
	}

	err = json.NewDecoder(response.Body).Decode(&geoResponse)
	if err != nil {
		log.Println("Could not get a success response")
		return geoResponse, err
	}

	return geoResponse, nil

}
