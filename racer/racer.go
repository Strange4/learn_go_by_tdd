package racer

import (
	"net/http"
	"time"
)

func Race(link1, link2 string) string {
	link1Time := measureHttpResponseTime(link1)
	link2Time := measureHttpResponseTime(link2)

	if link1Time < link2Time {
		return link1
	}
	return link2
}

func measureHttpResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
