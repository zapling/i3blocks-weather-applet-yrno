package internal

import "testing"

func TestGetCacheDir(t *testing.T) {
	if GetCacheDir("/home/test", "") != "/home/test/.local/share/weather-applet" {
		t.Error("Default cache directory invalid")
	}

	if GetCacheDir("/home/test", "/home/user/.local/share") != "/home/user/.local/share/weather-applet" {
		t.Error("Supplied xdg_data_dirs directory most be respected")
	}
}
