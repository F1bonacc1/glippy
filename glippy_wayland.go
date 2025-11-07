//go:build (linux && !android) || freebsd
// +build linux,!android freebsd

package glippy

import (
	"errors"
	"os"
	"os/exec"
)

func isWayland() bool {
	return os.Getenv("XDG_SESSION_TYPE") == "wayland"
}

func getWayland() (string, error) {
	if _, err := exec.LookPath("wl-paste"); err != nil {
		return "", errors.New("wl-clipboard not found, please install it")
	}

	cmd := exec.Command("wl-paste")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func setWayland(text string) error {
	if _, err := exec.LookPath("wl-copy"); err != nil {
		return errors.New("wl-clipboard not found, please install it")
	}

	cmd := exec.Command("wl-copy", text)
	return cmd.Run()
}
