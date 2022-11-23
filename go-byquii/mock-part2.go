package main

import (
    "testing"
	"bytes"
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct{}

func (d DefaultSleeper) sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}
//98

type SpySleeper struct {
	Calls int
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(out, i)
		fmt.Fprint(out, "\n")
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	
	Countdown(buffer, spySleeper)

	got := buffer.String()

	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}