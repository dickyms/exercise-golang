package main

import (
    "testing"
	"fmt"
	"bytes"
)

func Greet(writer *bytes.Buffer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	fmt.Println(buffer)
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}