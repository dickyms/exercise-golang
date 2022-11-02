package main

import (
	"fmt"
	"errors"
	"log"
)

var pl = fmt.Println
// how to throw error

func sayHello(name string) (string, error){
	if name == "" {
		pl("ini masuk error")
		return "", errors.New("parameter nama tidak diisi") //tulis error, bisa pakai panic, tapi panic itu menghentikan laju programnya
	}

	msg := fmt.Sprintf("Hello, %v. Welcome!", name)
	return msg, nil
}
func main() {
	msg, err := sayHello("sadsa")
	if err != nil {
		log.Fatal(err) // dari sayHello menerima error terus di console, exit status 1
	}

	pl(msg)
}