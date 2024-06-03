package main

import (
	"hello/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout, &mocking.DefaultSleeper{})
}
