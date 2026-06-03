package models

type CurrentWeather struct {
	City        string
	Temperature float64
	Condition   string
	Humidity    int
	WindSpeed   float64
}

type ForecastDay struct {
	Date      string
	High      float64
	Low       float64
	Condition string
}

type Forecast struct {
	City string
	Days []ForecastDay
}


type FavoriteWithWeather struct {
    City    string         `json:"city"`
    Weather *CurrentWeather `json:"weather"`
}