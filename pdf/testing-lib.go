package main

import (
    "testing"
    "regexp"
)

func sayHello (s string) string {
	return "hello " + s
}

func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := sayHello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`sayHello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := sayHello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}