package main

import (
    "testing"
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	assertStrings:= func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("know word", func(t *testing.T){
		_, got := dic.Search("unknown")
		assertStrings(t, got, ErrNotFound)
	})
}

/*Part ini menguji apakan map yang dimasukan yang key yang tidak ada
apakah mengembalikan error atau tidak*/