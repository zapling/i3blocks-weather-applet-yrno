package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configuration struct {
	SSID      string  `json:"ssid"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ConfigManager struct {
	configPath string
}

func NewConfigManager(configPath string) *ConfigManager {
	return &ConfigManager{configPath: configPath}
}

func (c *ConfigManager) GetConfigBySSID(ssid string) {}

func (c *ConfigManager) getConfigurations() []configuration {
	err := os.MkdirAll(c.getPath(false), 0755)
	if err != nil {
		fmt.Println("Could not create or read config path")
		os.Exit(1)
	}

	file, err := os.Open(c.getPath(true))
	if err != nil {
		file, err = os.OpenFile(c.getPath(true), os.O_RDWR, 0755)
		if err != nil {
			fmt.Println("Could not create default config")
			os.Exit(1)
		}

		defer file.Close()

		defaults := &configuration{SSID: "My WIFI network", Latitude: 57.7, Longitude: 11.9}
		bytes, err := json.Marshal(defaults)
		if err != nil {
			fmt.Println("Could not convert default config to json")
			os.Exit(1)
		}

		file.Write(bytes)

		return []configuration{*defaults}
	}

	defer file.Close()

	var configs []configuration

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Could not read from config file")
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		fmt.Println("Could not read config content, is formatting correct?")
		os.Exit(1)
	}
	return configs
}

func (c *ConfigManager) getPath(fullPath bool) string {
	path := c.configPath + "/weather-applet"
	if fullPath == true {
		path += "/config.json"
	}

	return path
}

// var configFolder = "weather-applet"
// var configFile = "config.json"

// type Configuration struct {
// 	Ssid      string  `json:"ssid"`
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

// // Default configuration with Gothenburg points
// var defaultConfiguration Configuration = Configuration{
// 	Ssid:      "Your wifi name",
// 	Latitude:  57.7,
// 	Longitude: 11.9,
// }

// func GetConfig() []Configuration {
// 	path := getConfigDir()
// 	exists := os.MkdirAll(path, 0755)

// 	// could not create or read path
// 	if exists != nil {
// 		return getDefaultConfig()
// 	}

// 	fullPath := path + "/" + configFile
// 	f, err := os.Open(fullPath)
// 	// could not find config file
// 	if err != nil {
// 		writeDefaultConfig(fullPath)
// 		return getDefaultConfig()
// 	}

// 	defer f.Close()

// 	var config []Configuration
// 	bytes, _ := ioutil.ReadAll(f)
// 	json.Unmarshal(bytes, &config)

// 	return config
// }

// func getDefaultConfig() []Configuration {
// 	var config []Configuration
// 	return append(config, defaultConfiguration)
// }

// func writeDefaultConfig(path string) {
// 	config := getDefaultConfig()
// 	jsonData, _ := json.Marshal(config)
// 	f, _ := os.Create(path)
// 	f.Write(jsonData)
// 	f.Close()
// }

// func getConfigDir() string {
// 	xdg_config_home := os.Getenv("XDG_CONFIG_HOME")
// 	if xdg_config_home != "" {
// 		return xdg_config_home + "/" + configFolder
// 	}
// 	return os.Getenv("HOME") + "/.config/" + configFolder
// }
