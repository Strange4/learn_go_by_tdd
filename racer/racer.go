package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("timed out waiting for ping")

func Racer(link1, link2 string) (string, error) {
	const defaultDuration = 10 * time.Second
	return ConfigurableRacer(link1, link2, defaultDuration)
}

func ConfigurableRacer(link1, link2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(link1):
		return link1, nil
	case <-ping(link2):
		return link2, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
