package context

import (
	"context"
	"hello/assertions"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response      string
	fetchDuration time.Duration
	cancelled     bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(s.fetchDuration)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "cool beans"
	delay := 50 * time.Millisecond
	t.Run("Returns data from the store", func(t *testing.T) {
		store := SpyStore{response: data, fetchDuration: delay}

		server := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertions.AssertString(t, response.Body.String(), data)

		if store.cancelled {
			t.Error("expected the operation to NOT be canceled")
		}
	})
	t.Run("Tries to cancel the work done by the store", func(t *testing.T) {
		timeout := 5 * time.Millisecond
		store := SpyStore{response: data, fetchDuration: delay}

		server := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		time.AfterFunc(timeout, cancel)

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("expected the operation to be canceled")
		}
	})
}
