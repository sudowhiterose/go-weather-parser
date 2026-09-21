package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Current   CurrentWeather `json:"current"`
}

type CurrentWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature_2m"`
	WindSpeed   float64 `json:"wind_speed_10m"`
}

func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=55.76&longitude=37.62&hourly=temperature_2m,wind_speed_10m&current=temperature_2m,wind_speed_10m&timezone=Europe%2FMoscow&forecast_days=1"

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Ошибка сети:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка, код:", resp.StatusCode)
		return
	}

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Printf("Ошибка парсинга JSON: %v\n", err)
		return
	}

	fmt.Println("--- Данные о погоде получены ---")
	fmt.Println("Координаты: широта", weather.Latitude, "долгота", weather.Longitude)
	fmt.Println("Время замера:", weather.Current.Time)
	fmt.Println("Температура воздуха:", weather.Current.Temperature, "°C")
	fmt.Println("Скорость ветра:", weather.Current.WindSpeed, "м/с")

}
