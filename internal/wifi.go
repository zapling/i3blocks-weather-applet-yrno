package internal

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetCurrentSsid() (string, error) {
	nmcliPath, err := exec.LookPath("nmcli")
	if err != nil {
		return "", fmt.Errorf("failed to lookup nmcli cmd: %w", err)
	}

	cmd := nmcliPath + " -f IN-USE,SSID device wifi list --rescan no | grep '\\*' | awk '{ print $2 }'"

	ssid, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute cmd: %w", err)
	}

	return string(bytes.Trim(ssid, "\n")), nil
}
