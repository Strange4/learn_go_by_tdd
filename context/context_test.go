package context

import (
	"context"
	"errors"
	"hello/assertions"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	hasBeenWrittenTo bool
}

func (s *SpyResponseWriter) Header() http.Header {
	return nil
}

func (s *SpyResponseWriter) Write(p []byte) (int, error) {
	s.hasBeenWrittenTo = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(int) {
	s.hasBeenWrittenTo = true
}

type SpyStore struct {
	response      string
	fetchDuration time.Duration
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	// simulating data stream
	durationPerIter := time.Duration(int64(s.fetchDuration) / int64(len(s.response)))
	data := make(chan string, 1)
	defer close(data)
	go func() {
		var result string
		for _, char := range s.response {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(durationPerIter)
				result += string(char)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-data:
		return result, nil
	}
}

func TestServer(t *testing.T) {
	data := "cool beans"
	delay := 50 * time.Millisecond
	store := SpyStore{response: data, fetchDuration: delay}
	t.Run("Returns data from the store", func(t *testing.T) {

		server := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertions.AssertString(t, response.Body.String(), data)
	})
	t.Run("Tries to cancel the work done by the store", func(t *testing.T) {
		timeout := 5 * time.Millisecond

		server := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}

		time.AfterFunc(timeout, cancel)

		server.ServeHTTP(response, request)
		if response.hasBeenWrittenTo {
			t.Error("There should have been nothing writen in the response")
		}
	})
}
