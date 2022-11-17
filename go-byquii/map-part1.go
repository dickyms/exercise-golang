package main

import (
    "testing"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	return d[word], nil
}

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	assertStrings:= func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("know word", func(t *testing.T){
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
}

