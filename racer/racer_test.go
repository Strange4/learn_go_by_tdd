package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := makeServerWithDelay(0 * time.Millisecond)
	slowServer := makeServerWithDelay(50 * time.Millisecond)
	defer fastServer.Close()
	defer slowServer.Close()

	slow := slowServer.URL
	fast := fastServer.URL

	want := fast
	got := Race(slow, fast)
	if want != got {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}

func makeServerWithDelay(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
