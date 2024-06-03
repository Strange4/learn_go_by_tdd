package sync

import (
	"hello/assertions"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increments counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		want := 3
		assertions.AssertInteger(t, got, want)
	})

	t.Run("runs safely concurrently", func(t *testing.T) {
		wantCount := 400
		counter := Counter{}

		var waitGroup sync.WaitGroup
		waitGroup.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
		got := counter.Value()
		assertions.AssertInteger(t, got, wantCount)
	})
}
