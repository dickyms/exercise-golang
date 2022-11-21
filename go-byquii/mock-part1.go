package main

import (
    "testing"
	"bytes"
	"fmt"
	"io"
)

func Countdown(out io.Writer){
	for i := 3; i > 0; i-- {
		fmt.Fprint(out, i)
		fmt.Fprint(out, "\n")
	}
	fmt.Fprint(out, "Go!")
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	
	Countdown(buffer)

	got := buffer.String()

	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}