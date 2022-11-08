package main

import (
    "testing"
    "fmt"
    "time"
)


func TestCreateChannel(t *testing.T) {
    channel := make(chan string)
    defer close(channel)
    go func() {
        time.Sleep(2* time.Second)
        channel <- "Dicky M. S."
        fmt.Println("selesai mengirim data")
    }()

    data := <-channel
    fmt.Println(data)
    time.Sleep(5 * time.Second)
}

//23.35