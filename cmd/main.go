package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	get_ip "github.com/giovane-aG/my-current-weather/internal/get-user-ip-address"
	get_user_coordinates "github.com/giovane-aG/my-current-weather/internal/get-user-location"
)

type ForecastApiResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeather       struct {
		Temperature   float64 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Winddirection float64 `json:"winddirection"`
		Weathercode   int     `json:"weathercode"`
		IsDay         int     `json:"is_day"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
}

func main() {
	ip, err := get_ip.GetUserIpAddress()
	if err != nil {
		log.Fatal("Error ->", err)
	}

	user_position, err := get_user_coordinates.GetUserLocation(ip)
	if err != nil {
		log.Fatal("Error2 ->", err)
	}

	lat := strings.Trim(strings.Split(user_position.Loc, ",")[0], `" `)
	lng := strings.Trim(strings.Split(user_position.Loc, ",")[1], `" `)

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current_weather=true&timezone=America/Sao_Paulo", lat, lng)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error3 ->", err)
	}

	defer response.Body.Close()
	var parsedBody = &ForecastApiResponse{}
	err = json.NewDecoder(response.Body).Decode(parsedBody)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current Weather in %v:\n", user_position.City)
	fmt.Printf("Time: %v\n", parsedBody.CurrentWeather.Time)
	fmt.Printf("TimeZone: %v\n", parsedBody.Timezone)
	fmt.Printf("Temperature: %v°C\n", parsedBody.CurrentWeather.Temperature)
	fmt.Printf("Windspeed: %vkm/h\n", parsedBody.CurrentWeather.Windspeed)
}
