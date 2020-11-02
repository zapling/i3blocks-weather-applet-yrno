package internal

import (
	"encoding/json"
    "errors"
	"fmt"
	"io/ioutil"
	"os"
)

type configuration struct {
	SSID      string  `json:"ssid"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var EmptyConfigPath = errors.New("Can not instantiate with empty config path")

type ConfigManager struct {
	configPath string
}

func NewConfigManager(configPath string) (*ConfigManager, error) {
    if configPath == "" {
        return &ConfigManager{}, EmptyConfigPath
    }

	return &ConfigManager{configPath: configPath}, nil
}

func (c *ConfigManager) GetConfigBySSID(ssid string) *configuration {
	configs := c.getConfigurations()
	for _, config := range configs {
		if config.SSID != ssid {
			continue
		}
		return &config
	}

	return nil
}

func (c *ConfigManager) getConfigurations() []configuration {
	err := os.MkdirAll(c.getPath(false), 0755)
	if err != nil {
		fmt.Println("Could not create or read config path")
		os.Exit(1)
	}

	file, err := os.Open(c.getPath(true))
	if err != nil {
		file, err = os.OpenFile(c.getPath(true), os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Could not create default config")
			os.Exit(1)
		}

		defer file.Close()

		defaults := []configuration{
            {SSID: "My WIFI network", Latitude: 57.7, Longitude: 11.9},
        }
		bytes, err := json.Marshal(defaults)
		if err != nil {
			fmt.Println("Could not convert default config to json")
			os.Exit(1)
		}

		file.Write(bytes)

		return defaults
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
