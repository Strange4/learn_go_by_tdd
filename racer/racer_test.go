package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	slow := slowServer.URL
	fast := fastServer.URL

	want := fast
	got := Race(slow, fast)
	if want != got {
		t.Errorf("Wanted %q but got %q", want, got)
	}

	fastServer.Close()
	slowServer.Close()
}
