package internal

import (
	"fmt"
	"github.com/zapling/yr.no-golang-client/pkg/yr"
	"strings"
    "time"
    "os"
)

func GetForecast(config *configuration) string {
	forecast, err := yr.GetLocationForecast(
		config.Latitude,
		config.Longitude,
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

        emojie := Emojies[symbols[0]]
        return temperature + "°C " + emojie
    }

    return "No forecast for this time. Bug?"
}
