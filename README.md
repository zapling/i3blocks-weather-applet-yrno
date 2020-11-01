# i3blocks-weather-applet-yrno
A script that retreives the current temperature based on the current SSID you are connected to.

# Installation

```
go install cmd/weatherapplet.go
```

# Configuration

When first launching the application a config file will be generated and placed in your config
directory.

```
$HOME/.config/weather-applet/config.json
```

You can modify the config and add how many SSIDs you want. Provide latitude and longitude cordinations for your location.

```
[
  {"ssid": "My wifi network name", "latitude": 57.7, "longitude": 11.9}
]
```
