//go:build !windows
// +build !windows

package glippy

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	osc52DisableEnv = "GLIPPY_DISABLE_OSC52"
	// Upper bound on the plaintext payload. tmux and screen historically cap
	// OSC 52 around 74 KB after base64; keeping the plaintext below that keeps
	// the encoded form under the limit.
	osc52MaxBytes = 74994
)

// osc52Writer is the seam tests override to capture the sequence.
var osc52Writer = openTTY

func openTTY() (io.WriteCloser, error) {
	return os.OpenFile("/dev/tty", os.O_WRONLY, 0)
}

func setOSC52(text string) error {
	if os.Getenv(osc52DisableEnv) != "" {
		return errors.New("OSC 52 clipboard disabled by " + osc52DisableEnv)
	}
	if len(text) > osc52MaxBytes {
		return fmt.Errorf("OSC 52 payload too large: %d bytes (max %d)", len(text), osc52MaxBytes)
	}
	seq := buildOSC52Sequence(text, os.Getenv("TMUX") != "")
	w, err := osc52Writer()
	if err != nil {
		return fmt.Errorf("OSC 52: open TTY: %w", err)
	}
	defer w.Close()
	if _, err := w.Write([]byte(seq)); err != nil {
		return fmt.Errorf("OSC 52: write: %w", err)
	}
	return nil
}

func buildOSC52Sequence(text string, tmux bool) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	seq := "\x1b]52;c;" + encoded + "\x07"
	if tmux {
		// tmux DCS passthrough: ESC P tmux ; <escaped> ESC \
		// Every ESC inside the payload must be doubled.
		return "\x1bPtmux;" + strings.ReplaceAll(seq, "\x1b", "\x1b\x1b") + "\x1b\\"
	}
	return seq
}
