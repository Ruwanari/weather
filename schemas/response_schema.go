package schemas

type GeoResp struct {
	PropertyEntities Property `json:"properties"`
}

type Property struct {
	Meta       MetaEntity         `json:"meta"`
	TimeSeries []TimeSeriesEntity `json:"timeseries"`
}

type MetaEntity struct {
	Units UnitEntity `json:"units"`
}

type UnitEntity struct {
	AirPressureAtSeaLevel string `json:"air_pressure_at_sea_level"`
	AirTemperature        string `json:"air_temperature"`
	CloudAreaFraction     string `json:"cloud_area_fraction"`
	PrecipitationAmount   string `json:"precipitation_amount"`
	RelativeHumidity      string `json:"relative_humidity"`
	WindFromDirection     string `json:"wind_from_direction"`
	WindSpeed             string `json:"wind_speed"`
}

type TimeSeriesEntity struct {
	Time string     `json:"time"`
	Data DataEntity `json:"data"`
}

type DataEntity struct {
	Instant InstantEntity `json:"instant"`
}

type InstantEntity struct {
	Details DetailsEntity `json:"details"`
}

type DetailsEntity struct {
	AirPressureAtSeaLevel float64 `json:"air_pressure_at_sea_level"`
	AirTemperature        float64 `json:"air_temperature"`
	CloudAreaFraction     float64 `json:"cloud_area_fraction"`
	RelativeHumidity      float64 `json:"relative_humidity"`
	WindFromDirection     float64 `json:"wind_from_direction"`
	WindSpeed             float64 `json:"wind_speed"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
