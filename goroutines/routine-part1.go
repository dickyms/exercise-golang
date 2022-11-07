package main
// jalankan dengan mode unit test

import (
    "testing"
    "fmt"
    "time"
)

func RunHelloWorld() {
    fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
    go RunHelloWorld()
    fmt.Println("UPS")

    time.Sleep(10 * time.Nanosecond)
}

func DisplayNumber(number int ) {
    fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
    for i := 0 ; i < 10000; i++ {
       go DisplayNumber(i)
    }

    time.Sleep(5 * time.Second)
}