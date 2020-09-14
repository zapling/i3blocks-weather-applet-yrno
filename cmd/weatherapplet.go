package main

import (
	"fmt"
	"github.com/zapling/i3blocks-weather-applet-yrno/internal"
	"os"
)

func main() {
	ssid := internal.GetCurrentSsid()
	if ssid == "" {
		os.Exit(1)
	}

	config := internal.GetConfig()
	for i := range config {
		if config[i].Ssid == ssid {
			fmt.Println(internal.GetForecast(config[i]))
		}
	}
}
