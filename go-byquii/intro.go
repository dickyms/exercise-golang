package main

import (
    "testing"
)


func sayHello (name string, lang string) string {
	result := "Hello "
	if lang == "idn" {
		result = "Hai " + name
	} else if lang == "es" {
		result = "Hola " + name
	} else {
		result = "Well " + name
	}
	return result
}

func TestHello (t *testing.T) {
	got := sayHello("Dicky", "idn")
	want := "Ha Dicky"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}