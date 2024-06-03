package racer

import (
	"hello/assertions"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the fastest url to respond", func(t *testing.T) {
		fastServer := makeServerWithDelay(0 * time.Millisecond)
		slowServer := makeServerWithDelay(50 * time.Millisecond)
		defer fastServer.Close()
		defer slowServer.Close()

		slow := slowServer.URL
		fast := fastServer.URL

		want := fast
		got, err := Racer(slow, fast)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		assertions.AssertString(t, got, want)
	})

	t.Run("Times out after 20 milliseconds and returns errors", func(t *testing.T) {
		verySlowServer := makeServerWithDelay(50 * time.Millisecond)
		_, err := ConfigurableRacer(verySlowServer.URL, verySlowServer.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}

func makeServerWithDelay(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
