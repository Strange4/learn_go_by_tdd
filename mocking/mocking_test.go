package mocking

import (
	"bytes"
	"hello/assertions"
	"reflect"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	callsMade []string
}

const writeCall = "write"
const sleepCall = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.callsMade = append(s.callsMade, sleepCall)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.callsMade = append(s.callsMade, writeCall)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Counts down 3 numbers", func(t *testing.T) {
		spySleeper := SpyCountdownOperations{}
		buffer := bytes.Buffer{}
		Countdown(&buffer, &spySleeper)
		want := `3
2
1
Go!`
		got := buffer.String()
		assertions.AssertString(t, got, want)
	})
	t.Run("Has the right number of sleeps and writes", func(t *testing.T) {
		spySleeper := SpyCountdownOperations{}
		Countdown(&spySleeper, &spySleeper)
		got := spySleeper.callsMade
		want := []string{
			writeCall,
			sleepCall,
			writeCall,
			sleepCall,
			writeCall,
			sleepCall,
			writeCall,
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted the calls %v but got %v", want, got)
		}
	})
}

type SpySleepTime struct {
	durationSlept time.Duration
}

func (s *SpySleepTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spySleepTime := SpySleepTime{}
	sleeper := ConfigurableSleeper{sleepTime, spySleepTime.Sleep}
	sleeper.Sleep()
	if spySleepTime.durationSlept != sleepTime {
		t.Errorf("The sleeper should've slept %v but slept %v", sleepTime, spySleepTime.durationSlept)
	}
}
