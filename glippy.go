package glippy

import (
	"context"
	"sync"
	"time"
)

const baseWatchInterval = time.Second * 1

// Clipboard method identifiers returned by SetWithMethod.
const (
	// MethodNative indicates the OS clipboard was used (X11, Wayland, pbcopy,
	// Windows clipboard API).
	MethodNative = "native"
	// MethodOSC52 indicates the OSC 52 escape-sequence fallback was used.
	// OSC 52 is fire-and-forget: success means the sequence was written to the
	// terminal, not that the terminal honored it. Many terminal emulators
	// disable OSC 52 writes by default.
	MethodOSC52 = "osc52"
)

var once sync.Once

func startOnce() {
	once.Do(func() {
		start()
	})
}

// Set sets clipboard content.
func Set(text string) error {
	_, err := SetWithMethod(text)
	return err
}

// SetWithMethod sets clipboard content and reports which mechanism was used.
// See the Method* constants for possible values.
func SetWithMethod(text string) (method string, err error) {
	startOnce()
	return set(text)
}

// Get get clipboard content
func Get() (string, error) {
	startOnce()
	return get()
}

// WatchWithInterval watching clipboard content at a specified interval
func WatchWithInterval(ctx context.Context, interval time.Duration) <-chan string {
	recv := make(chan string, 1)
	go func() {
		ticker := time.NewTicker(interval)
		lastData := ""
		for {
			select {
			case <-ctx.Done():
				close(recv)
				return
			case <-ticker.C:
				data, err := Get()
				if err != nil {
					continue
				}

				if data != lastData {
					recv <- data
					lastData = data
				}
			}
		}
	}()

	return recv
}

// Watch watching clipboard content
func Watch(ctx context.Context) <-chan string {
	return WatchWithInterval(ctx, baseWatchInterval)
}
