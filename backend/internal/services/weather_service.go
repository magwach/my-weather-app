package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/magwach/my-weather-app/backend/internal/db"
	"github.com/magwach/my-weather-app/backend/internal/models"
)

type WeatherService struct {
	ApiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{ApiKey: apiKey}
}

type owmCurrentResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Cod any `json:"cod"`
}

type owmForecastResponse struct {
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	List []struct {
		DtTxt string `json:"dt_txt"`
		Main  struct {
			TempMax float64 `json:"temp_max"`
			TempMin float64 `json:"temp_min"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
}

func roundCoord(val string) (string, error) {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f", f), nil
}

func (s *WeatherService) GetCurrentWeather(city string) (*models.CurrentWeather, error) {

	key := "weather:current:" + strings.ToLower(city)

	if cached, err := db.GetCache(key); err == nil {
		var w models.CurrentWeather
		if json.Unmarshal([]byte(cached), &w) == nil {
			return &w, nil
		}
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, s.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("weather service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("city not found")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("weather service unavailable")
	}

	var raw owmCurrentResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	mapped := &models.CurrentWeather{
		City:        raw.Name,
		Temperature: raw.Main.Temp,
		Condition:   raw.Weather[0].Description,
		Humidity:    raw.Main.Humidity,
		WindSpeed:   raw.Wind.Speed,
	}

	if data, err := json.Marshal(mapped); err == nil {
		_ = db.SetCache(key, string(data), 10*time.Minute)
	}

	return mapped, nil
}

func (s *WeatherService) GetForecast(city string) (*models.Forecast, error) {

	key := "weather:forecast:" + strings.ToLower(city)

	if cached, err := db.GetCache(key); err == nil {
		var f models.Forecast
		if json.Unmarshal([]byte(cached), &f) == nil {
			return &f, nil
		}
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric",
		city, s.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("weather service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("city not found")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("weather service unavailable")
	}

	var raw owmForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	daily := make(map[string]*models.ForecastDay)

	for _, item := range raw.List {

		date := item.DtTxt[:10]

		if _, exists := daily[date]; !exists {
			daily[date] = &models.ForecastDay{
				Date:      date,
				High:      item.Main.TempMax,
				Low:       item.Main.TempMin,
				Condition: item.Weather[0].Description,
			}
		}

		d := daily[date]

		if item.Main.TempMax > d.High {
			d.High = item.Main.TempMax
		}
		if item.Main.TempMin < d.Low {
			d.Low = item.Main.TempMin
		}

		if strings.Contains(item.DtTxt, "12:00:00") {
			d.Condition = item.Weather[0].Description
		}
	}

	var days []models.ForecastDay
	for _, v := range daily {
		days = append(days, *v)
	}

	result := &models.Forecast{
		City: raw.City.Name,
		Days: days,
	}

	if data, err := json.Marshal(result); err == nil {
		_ = db.SetCache(key, string(data), 30*time.Minute)
	}

	return result, nil
}

func (s *WeatherService) GetCurrentWeatherByCoords(lat, lon string) (*models.CurrentWeather, error) {

	latR, err := roundCoord(lat)
	if err != nil {
		return nil, errors.New("invalid latitude")
	}

	lonR, err := roundCoord(lon)
	if err != nil {
		return nil, errors.New("invalid longitude")
	}

	key := fmt.Sprintf("weather:current:coords:%s:%s", latR, lonR)

	if cached, err := db.GetCache(key); err == nil {
		var w models.CurrentWeather
		if json.Unmarshal([]byte(cached), &w) == nil {
			return &w, nil
		}
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric",
		latR, lonR, s.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("weather service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("weather service unavailable")
	}

	var raw owmCurrentResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	mapped := &models.CurrentWeather{
		City:        raw.Name,
		Temperature: raw.Main.Temp,
		Condition:   raw.Weather[0].Description,
		Humidity:    raw.Main.Humidity,
		WindSpeed:   raw.Wind.Speed,
	}

	if data, err := json.Marshal(mapped); err == nil {
		_ = db.SetCache(key, string(data), 10*time.Minute)
	}

	return mapped, nil
}

func (s *WeatherService) GetForecastByCoords(lat, lon string) (*models.Forecast, error) {

	latR, err := roundCoord(lat)
	if err != nil {
		return nil, errors.New("invalid latitude")
	}

	lonR, err := roundCoord(lon)
	if err != nil {
		return nil, errors.New("invalid longitude")
	}

	key := fmt.Sprintf("weather:forecast:coords:%s:%s", latR, lonR)

	if cached, err := db.GetCache(key); err == nil {
		var f models.Forecast
		if json.Unmarshal([]byte(cached), &f) == nil {
			return &f, nil
		}
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&appid=%s&units=metric",
		latR, lonR, s.ApiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("weather service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("weather service unavailable")
	}

	var raw owmForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	daily := make(map[string]*models.ForecastDay)

	for _, item := range raw.List {

		date := item.DtTxt[:10]

		if _, exists := daily[date]; !exists {
			daily[date] = &models.ForecastDay{
				Date:      date,
				High:      item.Main.TempMax,
				Low:       item.Main.TempMin,
				Condition: item.Weather[0].Description,
			}
		}

		d := daily[date]

		if item.Main.TempMax > d.High {
			d.High = item.Main.TempMax
		}
		if item.Main.TempMin < d.Low {
			d.Low = item.Main.TempMin
		}

		if strings.Contains(item.DtTxt, "12:00:00") {
			d.Condition = item.Weather[0].Description
		}
	}

	var days []models.ForecastDay
	for _, v := range daily {
		days = append(days, *v)
	}

	result := &models.Forecast{
		City: raw.City.Name,
		Days: days,
	}

	if data, err := json.Marshal(result); err == nil {
		_ = db.SetCache(key, string(data), 30*time.Minute)
	}

	return result, nil
}
