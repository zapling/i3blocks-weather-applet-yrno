package internal

var cacheFolder = "weather-applet"

func GetCacheDir(xdg_data_dirs string) string {
	if xdg_data_dirs != "" {
		return xdg_data_dirs + "/" + cacheFolder
	}
	return "/usr/local/share/" + cacheFolder
}
