package internal

import "testing"

func TestGetCacheDir(t *testing.T) {
	if GetCacheDir("") != "/usr/local/share/weather-applet" {
		t.Error("Default cache directory invalid")
	}

	if GetCacheDir("/home/user/.local/share") != "/home/user/.local/share/weather-applet" {
		t.Error("Supplied xdg_data_dirs directory most be respected")
	}
}
