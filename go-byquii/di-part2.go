package main

import (
    "testing"
	"fmt"
	"io"
	"bytes"
)

// io.Writer digunakan karena bisa menerima interface lain yang meng-implemen interfacenya

func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}