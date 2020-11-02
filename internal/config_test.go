package internal

import (
    "testing"
    "errors"
)

func TestNewConfigManager(t *testing.T) {
    tests := []struct{
        configPath string
        expected error
    }{
        {configPath: "", expected: EmptyConfigPath},
        {configPath: "config/", expected: nil},
        {configPath: "/config/path", expected: nil},
    }

    for _, tc := range tests {
        confMan, err := NewConfigManager(tc.configPath)
        if !errors.Is(err, tc.expected) {
            t.Fatalf("got: %v, expected: %v", err, tc.expected)
        } else {
            if confMan.configPath != tc.configPath {
                t.Fatalf("got: %v, expected: %v", confMan.configPath, tc.configPath)
            }
        }
    }
}

func TestGetPath(t *testing.T) {
    tests := []struct{
        configPath string
        fullPath bool
        expected string
    }{
        {configPath: "/config", fullPath: false, expected: "/config/weather-applet"},
        {configPath: "/config", fullPath: true, expected: "/config/weather-applet/config.json"},
    }

    for _, tc := range tests {
        confMan, _ := NewConfigManager(tc.configPath)
        path := confMan.getPath(tc.fullPath)
        if path != tc.expected {
            t.Fatalf("got: %v, expected: %v", path, tc.expected)
        }
    }
}
