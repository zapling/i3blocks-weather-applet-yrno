package main

import (
	"fmt"
	"github.com/zapling/i3blocks-weather-applet-yrno/internal"
	"os"
)

func main() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Print("Could not get cache dir")
		os.Exit(1)
	}

	configs := internal.GetConfig()
	cacheMan := internal.NewCacheManager(cacheDir)

	ssid := internal.GetCurrentSsid()
	if ssid == "" {
		os.Exit(1)
	}

	for i := range configs {
		config := configs[i]
		if ssid != config.Ssid {
			continue
		}

		cached := cacheMan.GetValue(ssid)
		if cached != "" {
			fmt.Println(cached)
			os.Exit(0)
		}

		forecast := internal.GetForecast(config)
		cacheMan.WriteCache(ssid, forecast)

		fmt.Println(forecast)
		os.Exit(0)
	}

	os.Exit(1)
}
