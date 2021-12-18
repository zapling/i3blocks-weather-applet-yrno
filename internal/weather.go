package internal

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zapling/yr.no-golang-client/pkg/yr"
)

func GetForecast(config *config, ssid string) string {
	currentCfg := config.GetConfigBySSID(ssid)

	forecast, err := yr.GetLocationForecast(
		currentCfg.Latitude,
		currentCfg.Longitude,
		"WeatherApplet 1.0",
	)

	if err != nil {
		return ""
	}

	now := time.Now()
	for _, data := range forecast.Properties.Timeseries {
		forecastTime, err := time.Parse(time.RFC3339, data.Time)
		if err != nil {
			fmt.Println("Could not decode time")
			os.Exit(1)
		}

		if now.YearDay() != forecastTime.YearDay() || now.Hour() != forecastTime.Hour() {
			continue
		}

		temperature := fmt.Sprintf("%.0f", data.Data.Instant.Details.AirTemperature)
		symbols := strings.Split(
			data.Data.Next1Hours.Summary.SymbolCode,
			"_",
		)

		symbolName := symbols[0]
		emojie := config.GetEmojie(symbolName)
		return temperature + "°C " + emojie
	}

	return "No forecast for this time. Bug?"
}
