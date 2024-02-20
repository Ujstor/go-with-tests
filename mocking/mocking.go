package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord = "Go!"
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperator struct {
	Calls []string 
}

type DefaultSleeper struct{}


func (s *SpyCountdownOperator) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperator) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}


func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord) 
}

func main() {

	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)

}