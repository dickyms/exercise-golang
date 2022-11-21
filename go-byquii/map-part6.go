package main

import (
    "testing"
	//"errors"
)

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
} 
// DictionaryErr ini dia nge-implements interface error yang isinya ada method Error()

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default :
		return err
	}
	return nil
}
// map itu pada dasarnya menggunakan pointer, namun ketika mempraktikannya tidak menggunakan pointer karena map sendiri pointer ke struktur runtime.hmap 

func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _:= dic.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T){
		_, got := dic.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{}

		err := dic.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dic.Update(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, newDefinition)
	})

}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dic := Dictionary{word: definition}
	newDefintion := "new definition"

	dic.Update(word, newDefintion)

	assertDefinition(t, dic, word, newDefintion)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dic Dictionary, word, definition string) {
	t.Helper()
	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}