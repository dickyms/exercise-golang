package main

import (
    "testing"
    "net/http"
    "fmt"
)

func TestServeMux (t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hi World")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "images")
	})
	mux.HandleFunc("/images/profil", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "profil")
	})

	//prioritas url pattren dalam servemux

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}