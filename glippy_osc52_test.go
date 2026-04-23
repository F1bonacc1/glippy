//go:build !windows
// +build !windows

package glippy

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
	"testing"
)

type bufWriter struct{ bytes.Buffer }

func (b *bufWriter) Close() error { return nil }

func withTestWriter(t *testing.T) *bufWriter {
	t.Helper()
	orig := osc52Writer
	buf := &bufWriter{}
	osc52Writer = func() (io.WriteCloser, error) { return buf, nil }
	t.Cleanup(func() { osc52Writer = orig })
	return buf
}

func TestBuildOSC52Sequence_NoTmux(t *testing.T) {
	got := buildOSC52Sequence("hello", false)
	want := "\x1b]52;c;" + base64.StdEncoding.EncodeToString([]byte("hello")) + "\x07"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestBuildOSC52Sequence_Tmux(t *testing.T) {
	got := buildOSC52Sequence("hi", true)
	// Inside the tmux DCS passthrough, each ESC of the inner sequence is doubled.
	inner := "\x1b\x1b]52;c;" + base64.StdEncoding.EncodeToString([]byte("hi")) + "\x07"
	want := "\x1bPtmux;" + inner + "\x1b\\"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	// Sanity: the only double-ESC inside payload is the original ESC at the start.
	payload := strings.TrimPrefix(strings.TrimSuffix(got, "\x1b\\"), "\x1bPtmux;")
	if strings.Count(payload, "\x1b") != 2 {
		t.Errorf("expected exactly 2 ESC bytes in payload, got %d", strings.Count(payload, "\x1b"))
	}
}

func TestSetOSC52_WritesToTTY(t *testing.T) {
	t.Setenv(osc52DisableEnv, "")
	t.Setenv("TMUX", "")
	buf := withTestWriter(t)
	if err := setOSC52("payload"); err != nil {
		t.Fatalf("setOSC52: %v", err)
	}
	want := "\x1b]52;c;" + base64.StdEncoding.EncodeToString([]byte("payload")) + "\x07"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestSetOSC52_RespectsDisableEnv(t *testing.T) {
	t.Setenv(osc52DisableEnv, "1")
	withTestWriter(t)
	err := setOSC52("payload")
	if err == nil {
		t.Fatal("expected error when disabled, got nil")
	}
	if !strings.Contains(err.Error(), osc52DisableEnv) {
		t.Errorf("error should mention %s, got %q", osc52DisableEnv, err.Error())
	}
}

func TestSetOSC52_SizeGuard(t *testing.T) {
	t.Setenv(osc52DisableEnv, "")
	withTestWriter(t)
	big := strings.Repeat("x", osc52MaxBytes+1)
	err := setOSC52(big)
	if err == nil {
		t.Fatal("expected size-guard error, got nil")
	}
	if !strings.Contains(err.Error(), "too large") {
		t.Errorf("error should mention size, got %q", err.Error())
	}
}

func TestSetOSC52_TmuxEnvWrapsPayload(t *testing.T) {
	t.Setenv(osc52DisableEnv, "")
	t.Setenv("TMUX", "/tmp/tmux-1000/default,12345,0")
	buf := withTestWriter(t)
	if err := setOSC52("x"); err != nil {
		t.Fatalf("setOSC52: %v", err)
	}
	if !strings.HasPrefix(buf.String(), "\x1bPtmux;") {
		t.Errorf("expected tmux passthrough prefix, got %q", buf.String())
	}
	if !strings.HasSuffix(buf.String(), "\x1b\\") {
		t.Errorf("expected tmux passthrough terminator, got %q", buf.String())
	}
}
