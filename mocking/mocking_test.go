package mocking

import (
	"bytes"
	"hello/assertions"
	"testing"
)

type SpySleeper struct {
	callsMade int
}

func (s *SpySleeper) Sleep() {
	s.callsMade++
}

func TestCountdown(t *testing.T) {
	spySleeper := SpySleeper{}
	buffer := bytes.Buffer{}
	Countdown(&buffer, &spySleeper)
	want := `3
2
1
Go!`
	got := buffer.String()
	assertions.AssertString(t, got, want)
	assertions.AssertInteger(t, spySleeper.callsMade, 3)
}
