package paris

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	"data-eq1/API/model"
)

// JSONData struct to decode response
type JSONData struct {
	Data temperatureData `json:"main"`
}

type temperatureData struct {
	TemperatureKelvin float64 `json:"temp"`
	Humidity          float64 `json:"humidity"`
}

type weather struct {
	TemperatureCelsius float64 `json:"temperature_celsius"`
	HumidityPercentage float64 `json:"humidity_percentage"`
}

// WeatherGET gets current weather for Paris
func WeatherGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Paris&appid=" + APIkey)
	var res model.Response
	res.StatusCode = 200
	res.Error = ""
	res.Message = "OK"
	if err != nil {
		res.StatusCode = 500
		res.Message = "Internal Server Error"
		res.Error = "OpenWeatherMap Server Error"
		log.Println("Error with OpenWeatherMap HTTP request:", err)
	}
	readResult, _ := ioutil.ReadAll(result.Body)
	defer r.Body.Close()
	var j JSONData
	json.Unmarshal(readResult, &j)
	weather := weather{
		TemperatureCelsius: math.Round((j.Data.TemperatureKelvin-273.15)*10) / 10,
		HumidityPercentage: j.Data.Humidity,
	}
	res.Data = append(res.Data, weather)
	res.Meta.Query = "Température et taux d'humidité à Paris en temps réel"
	json.NewEncoder(w).Encode(res)
}
