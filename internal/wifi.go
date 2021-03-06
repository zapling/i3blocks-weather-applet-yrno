package internal

import (
	"bytes"
	"os/exec"
)

func GetCurrentSsid() string {
	cmd := "iw dev | grep ssid | awk '{$1=\"\"; print substr($0,2)'}"
	ssid, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}

	return string(bytes.Trim(ssid, "\n"))
}
