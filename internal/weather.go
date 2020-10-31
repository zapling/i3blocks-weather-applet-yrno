package internal

import (
	"fmt"
	"github.com/zapling/yr.no-golang-client/pkg/yr"
	"strings"
)

func GetForecast(config configuration) string {
	forecast, err := yr.GetLocationForecast(
		config.Latitude,
		config.Longitude,
		"WeatherApplet 1.0",
	)

	if err != nil {
		return ""
	}

	temperature := fmt.Sprintf(
		"%.0f",
		forecast.Properties.Timeseries[0].Data.Instant.Details.AirTemperature,
	)
	symbols := strings.Split(
		forecast.Properties.Timeseries[0].Data.Next1Hours.Summary.SymbolCode,
		"_",
	)
	emojie := Emojies[symbols[0]]

	return temperature + "°C " + emojie
}
