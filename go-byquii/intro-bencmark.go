package main

import (
    "testing"
	"fmt"
)


func sayHello(name string, lang string) string {
	result := "Hello "
	if lang == "idn" {
		result = "Hai " + name
	} else if lang == "es" {
		result = "Hola " + name
	} else {
		result = "Hello World"
	}
	return result
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := sayHello("Dicky", "idn")
		want := "Hai Dicky"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := sayHello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

func calculate(s int) int{
	result := s + 2 
	return result
}

func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
    expected := 4
    result := calculate(2)
    if expected != result {
        t.Error("Failed")
    }
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHello("dicky", "")
	}
}

// untuk menjalankannya go test -run=NamaTestTidakAda -bench=BencmarkNamanya
// kalau untuk semuanya tinggal pake -bench=.
// -run=NamaTestTidakAda untuk mengskip unit test

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}