package racer

import (
	"net/http"
	"time"
)

func Race(link1, link2 string) string {
	start1 := time.Now()
	http.Get(link1)
	link1Time := time.Since(start1)

	start2 := time.Now()
	http.Get(link2)
	link2Time := time.Since(start2)

	if link1Time < link2Time {
		return link1
	}
	return link2
}
