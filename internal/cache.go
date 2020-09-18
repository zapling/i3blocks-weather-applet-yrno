package internal

var cacheFolder = "weather-applet"

func GetCacheDir(home_dir string, xdg_data_home string) string {
	if xdg_data_home != "" {
		return xdg_data_home + "/" + cacheFolder
	}

	return home_dir + "/.local/share/" + cacheFolder
}
