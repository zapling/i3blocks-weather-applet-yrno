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

	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Could not get config dir")
		os.Exit(1)
	}

	cacheMan := internal.NewCacheManager(cacheDir)
	confMan, err := internal.NewConfigManager(configDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ssid := internal.GetCurrentSsid()
	if ssid == "" {
		os.Exit(1)
	}

	config := confMan.GetConfigBySSID(ssid)
	if config == nil {
		os.Exit(0)
	}

	cached := cacheMan.GetCache(ssid)
	if cached != "" {
		fmt.Println(cached)
		os.Exit(0)
	}

	forecast := internal.GetForecast(config)
	cacheMan.SetCache(ssid, forecast)

	fmt.Println(forecast)
	os.Exit(0)
}
