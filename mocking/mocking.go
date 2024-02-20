package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	countdownstart = 3
	finalWord = "Go!"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownstart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord) 
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
