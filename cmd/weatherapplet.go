package main

import (
	"fmt"
	"os"

	"github.com/zapling/i3blocks-weather-applet-yrno/internal"
)

func main() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Print("Could not get cache dir")
		os.Exit(1)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Could not get config dir")
		os.Exit(1)
	}

	cacheMan := internal.NewCacheManager(cacheDir)
	config, err := internal.LoadConfig(configDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ssid, err := internal.GetCurrentSsid()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if ssid == "" {
		fmt.Println("no ssid")
		os.Exit(1)
	}

	if config.GetConfigBySSID(ssid) == nil {
		os.Exit(0)
	}

	cached := cacheMan.GetCache(ssid)
	if cached != "" {
		fmt.Println(cached)
		os.Exit(0)
	}

	forecast := internal.GetForecast(config, ssid)
	cacheMan.SetCache(ssid, forecast)

	fmt.Println(forecast)
	os.Exit(0)
}
