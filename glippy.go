package glippy

import "sync"

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
