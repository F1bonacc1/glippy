//go:build (linux && !android) || freebsd
// +build linux,!android freebsd

package glippy

import (
	"os"
	"os/exec"
	"testing"
)

func TestWaylandClipboardNotFound(t *testing.T) {
	if isWayland() {
		// remove wl-clipboard from path for this test
		path := os.Getenv("PATH")
		os.Setenv("PATH", "")
		defer os.Setenv("PATH", path)

		_, err := getWayland()
		if err == nil {
			t.Errorf("getWayland() error = %v, wantErr %v", err, true)
		}

		err = setWayland("test")
		if err == nil {
			t.Errorf("setWayland() error = %v, wantErr %v", err, true)
		}
	}
}

func TestSetGetWayland(t *testing.T) {
	if isWayland() {
		if _, err := exec.LookPath("wl-copy"); err != nil {
			t.Skip("wl-clipboard not found, skipping wayland tests")
		}
		tests := []struct {
			name    string
			want    string
			wantErr bool
		}{
			{
				name:    "abc",
				want:    "abc",
				wantErr: false,
			},
			{
				name:    "empty",
				want:    "",
				wantErr: false,
			},
			{
				name:    "emoji",
				want:    "🔥",
				wantErr: false,
			},
			{
				name:    "utf8",
				want:    "aš tave myliu",
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := setWayland(tt.want)
				if (err != nil) != tt.wantErr {
					t.Errorf("setWayland() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				got, err := getWayland()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWayland() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("getWayland() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
