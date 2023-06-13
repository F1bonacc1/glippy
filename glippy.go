package glippy

import (
	"context"
	"sync"
	"time"
)

const baseWatchInterval = time.Second * 1

var once sync.Once

func startOnce() {
	once.Do(func() {
		start()
	})
}

// Set set clipboard content
func Set(text string) error {
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
