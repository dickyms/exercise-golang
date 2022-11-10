package main

import (
    "testing"
    "net/http"
    "fmt"
)


/*func TestWebServer(t *testing.T) {
    server := http.Server{
        Addr: "localhost:8080",
        Hanlder : handler,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}*/

func TestHandler (t *testing.T) {
    var handler http.HandlerFunc = func (writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hello World")
    }

    server := http.Server{
        Addr: "localhost:8080",
        Handler : handler,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
