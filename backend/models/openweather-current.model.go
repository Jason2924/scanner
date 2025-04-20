package models

type OpenWeatherCurrentResp struct {
	Coordinate struct {
		Longitude float64 `json:"lon"`
		Latitude  float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string
	Main struct {
		Temperature    float32 `json:"temp"`
		FeelsLike      float64 `json:"feels_like"`
		TemperatureMin float32 `json:"temp_min"`
		TemperatureMax float32 `json:"temp_max"`
		Pressure       int     `json:"pressure"`
		Humidity       int     `json:"humidity"`
		SeaLevel       int     `json:"sea_level"`
		GroundLevel    int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed  float64 `json:"speed"`
		Degree int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	TimeOfData int64 `json:"dt"`
	System     struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	CityID   int    `json:"id"`
	CityName string `json:"name"`
	Cod      int    `json:"cod"`
}
