package main

import (
	"hello/mocking"
	"os"
	"time"
)

func main() {
	sleeper := mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, &sleeper)
}
