package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var EmptyConfigPath = errors.New("Can not instantiate with empty config path")

var defaultConfig = &config{
	SSIDS: map[string]ssid{
		"My WIFI network": {Latitude: 57.7, Longitude: 11.9},
	},
	EmojieOverride: nil,
}

type config struct {
	SSIDS          map[string]ssid    `json:"ssids"`
	EmojieOverride *map[string]string `json:"emojie_override"`
}

func (c *config) GetConfigBySSID(ssid string) *ssid {
	ssidConf, exists := c.SSIDS[ssid]
	if !exists {
		return nil
	}

	return &ssidConf
}

func (c *config) GetEmojie(name string) string {
	if c.EmojieOverride == nil {
		return Emojies[name]
	}

	if emojie, exists := (*c.EmojieOverride)[name]; exists {
		return emojie
	}

	return Emojies[name]
}

type ssid struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func LoadConfig(configPath string) (*config, error) {
	var loadedConfig config

	if configPath == "" {
		return nil, EmptyConfigPath
	}

	path := configPath + "/weather-applet"
	fullPath := path + "/config.json"

	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("Could not create or read config path")
	}

	file, err := os.Open(fullPath)
	if err != nil {
		if err := writeDefaultConfig(fullPath); err != nil {
			return nil, fmt.Errorf("Failed to save default config: %s", err)
		}

		return defaultConfig, nil
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&loadedConfig); err != nil {
		return nil, fmt.Errorf("Failed to load config: %s", err)
	}

	return &loadedConfig, nil
}

func writeDefaultConfig(fullPath string) error {
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(defaultConfig); err != nil {
		return err
	}

	return nil
}
