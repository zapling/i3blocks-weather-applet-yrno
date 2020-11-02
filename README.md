# i3blocks-weather-applet-yrno

A weather-applet used with `i3-blocks`. 

Uses your current SSID in order to get defined cordinations of where to get a forecast from. Forecasts are fetched from [Yr.no](yr.no), a norwegian weather forecast website.

# Installation

```
go install cmd/weatherapplet.go
```

# Configuration

When first launching the application a config file will be generated and placed in your config
directory.

```
$XDG_CONFIG_HOME/weather-applet/config.json
```

You can modify the config and add how many SSIDs you want. Provide latitude and longitude cordinations for your location.

```
[
  {"ssid": "My wifi network name", "latitude": 57.7, "longitude": 11.9}
]
```

## Caching

Weather forecasts are cached for 1 hour. The cache is saved to a cache file in your local cache directory.

```
$XDG_CACHE_HOME/weather-applet/cache.json
```

